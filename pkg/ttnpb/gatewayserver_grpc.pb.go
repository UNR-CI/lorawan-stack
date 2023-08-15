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

// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.22.2
// source: ttn/lorawan/v3/gatewayserver.proto

package ttnpb

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

const (
	GtwGs_LinkGateway_FullMethodName             = "/ttn.lorawan.v3.GtwGs/LinkGateway"
	GtwGs_GetConcentratorConfig_FullMethodName   = "/ttn.lorawan.v3.GtwGs/GetConcentratorConfig"
	GtwGs_GetMQTTConnectionInfo_FullMethodName   = "/ttn.lorawan.v3.GtwGs/GetMQTTConnectionInfo"
	GtwGs_GetMQTTV2ConnectionInfo_FullMethodName = "/ttn.lorawan.v3.GtwGs/GetMQTTV2ConnectionInfo"
)

// GtwGsClient is the client API for GtwGs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GtwGsClient interface {
	// Link a gateway to the Gateway Server for streaming upstream messages and downstream messages.
	LinkGateway(ctx context.Context, opts ...grpc.CallOption) (GtwGs_LinkGatewayClient, error)
	// Get configuration for the concentrator.
	GetConcentratorConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ConcentratorConfig, error)
	// Get connection information to connect an MQTT gateway.
	GetMQTTConnectionInfo(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*MQTTConnectionInfo, error)
	// Get legacy connection information to connect a The Things Network Stack V2 MQTT gateway.
	GetMQTTV2ConnectionInfo(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*MQTTConnectionInfo, error)
}

type gtwGsClient struct {
	cc grpc.ClientConnInterface
}

func NewGtwGsClient(cc grpc.ClientConnInterface) GtwGsClient {
	return &gtwGsClient{cc}
}

func (c *gtwGsClient) LinkGateway(ctx context.Context, opts ...grpc.CallOption) (GtwGs_LinkGatewayClient, error) {
	stream, err := c.cc.NewStream(ctx, &GtwGs_ServiceDesc.Streams[0], GtwGs_LinkGateway_FullMethodName, opts...)
	if err != nil {
		return nil, err
	}
	x := &gtwGsLinkGatewayClient{stream}
	return x, nil
}

type GtwGs_LinkGatewayClient interface {
	Send(*GatewayUp) error
	Recv() (*GatewayDown, error)
	grpc.ClientStream
}

type gtwGsLinkGatewayClient struct {
	grpc.ClientStream
}

func (x *gtwGsLinkGatewayClient) Send(m *GatewayUp) error {
	return x.ClientStream.SendMsg(m)
}

func (x *gtwGsLinkGatewayClient) Recv() (*GatewayDown, error) {
	m := new(GatewayDown)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *gtwGsClient) GetConcentratorConfig(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*ConcentratorConfig, error) {
	out := new(ConcentratorConfig)
	err := c.cc.Invoke(ctx, GtwGs_GetConcentratorConfig_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gtwGsClient) GetMQTTConnectionInfo(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*MQTTConnectionInfo, error) {
	out := new(MQTTConnectionInfo)
	err := c.cc.Invoke(ctx, GtwGs_GetMQTTConnectionInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gtwGsClient) GetMQTTV2ConnectionInfo(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*MQTTConnectionInfo, error) {
	out := new(MQTTConnectionInfo)
	err := c.cc.Invoke(ctx, GtwGs_GetMQTTV2ConnectionInfo_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GtwGsServer is the server API for GtwGs service.
// All implementations must embed UnimplementedGtwGsServer
// for forward compatibility
type GtwGsServer interface {
	// Link a gateway to the Gateway Server for streaming upstream messages and downstream messages.
	LinkGateway(GtwGs_LinkGatewayServer) error
	// Get configuration for the concentrator.
	GetConcentratorConfig(context.Context, *emptypb.Empty) (*ConcentratorConfig, error)
	// Get connection information to connect an MQTT gateway.
	GetMQTTConnectionInfo(context.Context, *GatewayIdentifiers) (*MQTTConnectionInfo, error)
	// Get legacy connection information to connect a The Things Network Stack V2 MQTT gateway.
	GetMQTTV2ConnectionInfo(context.Context, *GatewayIdentifiers) (*MQTTConnectionInfo, error)
	mustEmbedUnimplementedGtwGsServer()
}

// UnimplementedGtwGsServer must be embedded to have forward compatible implementations.
type UnimplementedGtwGsServer struct {
}

func (UnimplementedGtwGsServer) LinkGateway(GtwGs_LinkGatewayServer) error {
	return status.Errorf(codes.Unimplemented, "method LinkGateway not implemented")
}
func (UnimplementedGtwGsServer) GetConcentratorConfig(context.Context, *emptypb.Empty) (*ConcentratorConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetConcentratorConfig not implemented")
}
func (UnimplementedGtwGsServer) GetMQTTConnectionInfo(context.Context, *GatewayIdentifiers) (*MQTTConnectionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMQTTConnectionInfo not implemented")
}
func (UnimplementedGtwGsServer) GetMQTTV2ConnectionInfo(context.Context, *GatewayIdentifiers) (*MQTTConnectionInfo, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetMQTTV2ConnectionInfo not implemented")
}
func (UnimplementedGtwGsServer) mustEmbedUnimplementedGtwGsServer() {}

// UnsafeGtwGsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GtwGsServer will
// result in compilation errors.
type UnsafeGtwGsServer interface {
	mustEmbedUnimplementedGtwGsServer()
}

func RegisterGtwGsServer(s grpc.ServiceRegistrar, srv GtwGsServer) {
	s.RegisterService(&GtwGs_ServiceDesc, srv)
}

func _GtwGs_LinkGateway_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(GtwGsServer).LinkGateway(&gtwGsLinkGatewayServer{stream})
}

type GtwGs_LinkGatewayServer interface {
	Send(*GatewayDown) error
	Recv() (*GatewayUp, error)
	grpc.ServerStream
}

type gtwGsLinkGatewayServer struct {
	grpc.ServerStream
}

func (x *gtwGsLinkGatewayServer) Send(m *GatewayDown) error {
	return x.ServerStream.SendMsg(m)
}

func (x *gtwGsLinkGatewayServer) Recv() (*GatewayUp, error) {
	m := new(GatewayUp)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func _GtwGs_GetConcentratorConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GtwGsServer).GetConcentratorConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GtwGs_GetConcentratorConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GtwGsServer).GetConcentratorConfig(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GtwGs_GetMQTTConnectionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GtwGsServer).GetMQTTConnectionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GtwGs_GetMQTTConnectionInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GtwGsServer).GetMQTTConnectionInfo(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _GtwGs_GetMQTTV2ConnectionInfo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GtwGsServer).GetMQTTV2ConnectionInfo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GtwGs_GetMQTTV2ConnectionInfo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GtwGsServer).GetMQTTV2ConnectionInfo(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

// GtwGs_ServiceDesc is the grpc.ServiceDesc for GtwGs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GtwGs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.GtwGs",
	HandlerType: (*GtwGsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetConcentratorConfig",
			Handler:    _GtwGs_GetConcentratorConfig_Handler,
		},
		{
			MethodName: "GetMQTTConnectionInfo",
			Handler:    _GtwGs_GetMQTTConnectionInfo_Handler,
		},
		{
			MethodName: "GetMQTTV2ConnectionInfo",
			Handler:    _GtwGs_GetMQTTV2ConnectionInfo_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "LinkGateway",
			Handler:       _GtwGs_LinkGateway_Handler,
			ServerStreams: true,
			ClientStreams: true,
		},
	},
	Metadata: "ttn/lorawan/v3/gatewayserver.proto",
}

const (
	NsGs_ScheduleDownlink_FullMethodName = "/ttn.lorawan.v3.NsGs/ScheduleDownlink"
)

// NsGsClient is the client API for NsGs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type NsGsClient interface {
	// Instructs the Gateway Server to schedule a downlink message.
	// The Gateway Server may refuse if there are any conflicts in the schedule or
	// if a duty cycle prevents the gateway from transmitting.
	ScheduleDownlink(ctx context.Context, in *DownlinkMessage, opts ...grpc.CallOption) (*ScheduleDownlinkResponse, error)
}

type nsGsClient struct {
	cc grpc.ClientConnInterface
}

func NewNsGsClient(cc grpc.ClientConnInterface) NsGsClient {
	return &nsGsClient{cc}
}

func (c *nsGsClient) ScheduleDownlink(ctx context.Context, in *DownlinkMessage, opts ...grpc.CallOption) (*ScheduleDownlinkResponse, error) {
	out := new(ScheduleDownlinkResponse)
	err := c.cc.Invoke(ctx, NsGs_ScheduleDownlink_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// NsGsServer is the server API for NsGs service.
// All implementations must embed UnimplementedNsGsServer
// for forward compatibility
type NsGsServer interface {
	// Instructs the Gateway Server to schedule a downlink message.
	// The Gateway Server may refuse if there are any conflicts in the schedule or
	// if a duty cycle prevents the gateway from transmitting.
	ScheduleDownlink(context.Context, *DownlinkMessage) (*ScheduleDownlinkResponse, error)
	mustEmbedUnimplementedNsGsServer()
}

// UnimplementedNsGsServer must be embedded to have forward compatible implementations.
type UnimplementedNsGsServer struct {
}

func (UnimplementedNsGsServer) ScheduleDownlink(context.Context, *DownlinkMessage) (*ScheduleDownlinkResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ScheduleDownlink not implemented")
}
func (UnimplementedNsGsServer) mustEmbedUnimplementedNsGsServer() {}

// UnsafeNsGsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to NsGsServer will
// result in compilation errors.
type UnsafeNsGsServer interface {
	mustEmbedUnimplementedNsGsServer()
}

func RegisterNsGsServer(s grpc.ServiceRegistrar, srv NsGsServer) {
	s.RegisterService(&NsGs_ServiceDesc, srv)
}

func _NsGs_ScheduleDownlink_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DownlinkMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(NsGsServer).ScheduleDownlink(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: NsGs_ScheduleDownlink_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(NsGsServer).ScheduleDownlink(ctx, req.(*DownlinkMessage))
	}
	return interceptor(ctx, in, info, handler)
}

// NsGs_ServiceDesc is the grpc.ServiceDesc for NsGs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var NsGs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.NsGs",
	HandlerType: (*NsGsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ScheduleDownlink",
			Handler:    _NsGs_ScheduleDownlink_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ttn/lorawan/v3/gatewayserver.proto",
}

const (
	Gs_GetGatewayConnectionStats_FullMethodName      = "/ttn.lorawan.v3.Gs/GetGatewayConnectionStats"
	Gs_BatchGetGatewayConnectionStats_FullMethodName = "/ttn.lorawan.v3.Gs/BatchGetGatewayConnectionStats"
)

// GsClient is the client API for Gs service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GsClient interface {
	// Get statistics about the current gateway connection to the Gateway Server.
	// This is not persisted between reconnects.
	GetGatewayConnectionStats(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*GatewayConnectionStats, error)
	// Get statistics about gateway connections to the Gateway Server of a batch of gateways.
	// - Statistics are not persisted between reconnects.
	// - Gateways that are not connected or are part of a different cluster are ignored.
	// - The client should ensure that the requested gateways are in the requested cluster.
	// - The client should have the right to get the gateway connection stats on all requested gateways.
	BatchGetGatewayConnectionStats(ctx context.Context, in *BatchGetGatewayConnectionStatsRequest, opts ...grpc.CallOption) (*BatchGetGatewayConnectionStatsResponse, error)
}

type gsClient struct {
	cc grpc.ClientConnInterface
}

func NewGsClient(cc grpc.ClientConnInterface) GsClient {
	return &gsClient{cc}
}

func (c *gsClient) GetGatewayConnectionStats(ctx context.Context, in *GatewayIdentifiers, opts ...grpc.CallOption) (*GatewayConnectionStats, error) {
	out := new(GatewayConnectionStats)
	err := c.cc.Invoke(ctx, Gs_GetGatewayConnectionStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gsClient) BatchGetGatewayConnectionStats(ctx context.Context, in *BatchGetGatewayConnectionStatsRequest, opts ...grpc.CallOption) (*BatchGetGatewayConnectionStatsResponse, error) {
	out := new(BatchGetGatewayConnectionStatsResponse)
	err := c.cc.Invoke(ctx, Gs_BatchGetGatewayConnectionStats_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GsServer is the server API for Gs service.
// All implementations must embed UnimplementedGsServer
// for forward compatibility
type GsServer interface {
	// Get statistics about the current gateway connection to the Gateway Server.
	// This is not persisted between reconnects.
	GetGatewayConnectionStats(context.Context, *GatewayIdentifiers) (*GatewayConnectionStats, error)
	// Get statistics about gateway connections to the Gateway Server of a batch of gateways.
	// - Statistics are not persisted between reconnects.
	// - Gateways that are not connected or are part of a different cluster are ignored.
	// - The client should ensure that the requested gateways are in the requested cluster.
	// - The client should have the right to get the gateway connection stats on all requested gateways.
	BatchGetGatewayConnectionStats(context.Context, *BatchGetGatewayConnectionStatsRequest) (*BatchGetGatewayConnectionStatsResponse, error)
	mustEmbedUnimplementedGsServer()
}

// UnimplementedGsServer must be embedded to have forward compatible implementations.
type UnimplementedGsServer struct {
}

func (UnimplementedGsServer) GetGatewayConnectionStats(context.Context, *GatewayIdentifiers) (*GatewayConnectionStats, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetGatewayConnectionStats not implemented")
}
func (UnimplementedGsServer) BatchGetGatewayConnectionStats(context.Context, *BatchGetGatewayConnectionStatsRequest) (*BatchGetGatewayConnectionStatsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BatchGetGatewayConnectionStats not implemented")
}
func (UnimplementedGsServer) mustEmbedUnimplementedGsServer() {}

// UnsafeGsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GsServer will
// result in compilation errors.
type UnsafeGsServer interface {
	mustEmbedUnimplementedGsServer()
}

func RegisterGsServer(s grpc.ServiceRegistrar, srv GsServer) {
	s.RegisterService(&Gs_ServiceDesc, srv)
}

func _Gs_GetGatewayConnectionStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GatewayIdentifiers)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GsServer).GetGatewayConnectionStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gs_GetGatewayConnectionStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GsServer).GetGatewayConnectionStats(ctx, req.(*GatewayIdentifiers))
	}
	return interceptor(ctx, in, info, handler)
}

func _Gs_BatchGetGatewayConnectionStats_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BatchGetGatewayConnectionStatsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GsServer).BatchGetGatewayConnectionStats(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Gs_BatchGetGatewayConnectionStats_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GsServer).BatchGetGatewayConnectionStats(ctx, req.(*BatchGetGatewayConnectionStatsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Gs_ServiceDesc is the grpc.ServiceDesc for Gs service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Gs_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ttn.lorawan.v3.Gs",
	HandlerType: (*GsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetGatewayConnectionStats",
			Handler:    _Gs_GetGatewayConnectionStats_Handler,
		},
		{
			MethodName: "BatchGetGatewayConnectionStats",
			Handler:    _Gs_BatchGetGatewayConnectionStats_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ttn/lorawan/v3/gatewayserver.proto",
}
