// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gatewayserver

import (
	"context"
	"fmt"

	"github.com/mohae/deepcopy"
	clusterauth "go.thethings.network/lorawan-stack/v3/pkg/auth/cluster"
	"go.thethings.network/lorawan-stack/v3/pkg/errors"
	"go.thethings.network/lorawan-stack/v3/pkg/events"
	"go.thethings.network/lorawan-stack/v3/pkg/gatewayserver/io"
	"go.thethings.network/lorawan-stack/v3/pkg/log"
	"go.thethings.network/lorawan-stack/v3/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/v3/pkg/unique"
)

var (
	errNotTxRequest = errors.DefineInvalidArgument("not_tx_request", "downlink message is not a Tx request")
	errSchedulePath = errors.Define("schedule_path", "failed to schedule on path `{gateway_uid}`")
	errSchedule     = errors.DefineAborted("schedule", "failed to schedule")
	errUplinkToken  = errors.DefineCorruption("uplink_token", "uplink token is not generated by this server")
)

// ScheduleDownlink instructs the Gateway Server to schedule a downlink message request.
// This method returns an error if the downlink path cannot be found, if the requested parameters are invalid for the
// gateway's frequency plan or if there is no transmission window available because of scheduling conflicts or regional
// limitations such as duty-cycle and dwell time.
func (gs *GatewayServer) ScheduleDownlink(ctx context.Context, down *ttnpb.DownlinkMessage) (*ttnpb.ScheduleDownlinkResponse, error) {
	if err := clusterauth.Authorized(ctx); err != nil {
		return nil, err
	}
	request := down.GetRequest()
	if request == nil {
		return nil, errNotTxRequest.New()
	}

	var pathErrs []errors.ErrorDetails
	logger := log.FromContext(ctx)
	for _, path := range request.DownlinkPaths {
		var ids ttnpb.GatewayIdentifiers
		switch p := path.Path.(type) {
		case *ttnpb.DownlinkPath_Fixed:
			ids = p.Fixed.GatewayIdentifiers
		case *ttnpb.DownlinkPath_UplinkToken:
			token, err := io.ParseUplinkToken(p.UplinkToken)
			if err != nil {
				pathErrs = append(pathErrs, errUplinkToken.New()) // Hide the cause as uplink tokens are opaque to the Network Server.
				continue
			}
			ids = token.GatewayIdentifiers
		default:
			panic(fmt.Sprintf("proto: unexpected type %T in oneof", path.Path))
		}

		uid := unique.ID(ctx, ids)
		conn, ok := gs.GetConnection(ctx, ids)
		if !ok {
			pathErrs = append(pathErrs, errNotConnected.WithAttributes("gateway_uid", uid))
			continue
		}
		down := deepcopy.Copy(down).(*ttnpb.DownlinkMessage) // Let the connection own the DownlinkMessage.
		down.GetRequest().DownlinkPaths = nil                // And do not leak the downlink paths to the gateway.
		delay, err := conn.ScheduleDown(path, down)
		if err != nil {
			logger.WithField("gateway_uid", uid).WithError(err).Debug("Failed to schedule on path")
			pathErrs = append(pathErrs, errSchedulePath.WithCause(err).WithAttributes("gateway_uid", uid))
			continue
		}
		ctx = events.ContextWithCorrelationID(ctx, events.CorrelationIDsFromContext(conn.Context())...)
		down.CorrelationIDs = append(down.CorrelationIDs, events.CorrelationIDsFromContext(ctx)...)
		msg := *down
		registerSendDownlink(ctx, conn.Gateway(), &msg, conn.Frontend().Protocol())
		return &ttnpb.ScheduleDownlinkResponse{
			Delay: delay,
		}, nil
	}

	protoErrs := make([]*ttnpb.ErrorDetails, 0, len(pathErrs))
	for _, pathErr := range pathErrs {
		protoErrs = append(protoErrs, ttnpb.ErrorDetailsToProto(pathErr))
	}
	return nil, errSchedule.WithDetails(&ttnpb.ScheduleDownlinkErrorDetails{
		PathErrors: protoErrs,
	})
}
