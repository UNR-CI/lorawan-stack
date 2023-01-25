// Code generated by protoc-gen-go-json. DO NOT EDIT.
// versions:
// - protoc-gen-go-json v1.4.2
// - protoc             v3.21.1
// source: lorawan-stack/api/identifiers.proto

package ttnpb

import (
	golang "github.com/TheThingsIndustries/protoc-gen-go-json/golang"
	jsonplugin "github.com/TheThingsIndustries/protoc-gen-go-json/jsonplugin"
	types "go.thethings.network/lorawan-stack/v3/pkg/types"
)

// MarshalProtoJSON marshals the EndDeviceIdentifiers message to JSON.
func (x *EndDeviceIdentifiers) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.DeviceId != "" || s.HasField("device_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("device_id")
		s.WriteString(x.DeviceId)
	}
	if x.ApplicationIds != nil || s.HasField("application_ids") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("application_ids")
		// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
		golang.MarshalMessage(s, x.ApplicationIds)
	}
	if len(x.DevEui) > 0 || s.HasField("dev_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_eui")
		types.MarshalHEXBytes(s.WithField("dev_eui"), x.DevEui)
	}
	if len(x.JoinEui) > 0 || s.HasField("join_eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("join_eui")
		types.MarshalHEXBytes(s.WithField("join_eui"), x.JoinEui)
	}
	if len(x.DevAddr) > 0 || s.HasField("dev_addr") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("dev_addr")
		types.MarshalHEXBytes(s.WithField("dev_addr"), x.DevAddr)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the EndDeviceIdentifiers to JSON.
func (x *EndDeviceIdentifiers) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the EndDeviceIdentifiers message from JSON.
func (x *EndDeviceIdentifiers) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "device_id", "deviceId":
			s.AddField("device_id")
			x.DeviceId = s.ReadString()
		case "application_ids", "applicationIds":
			s.AddField("application_ids")
			if s.ReadNil() {
				x.ApplicationIds = nil
				return
			}
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			golang.UnmarshalMessage(s, &v)
			x.ApplicationIds = &v
		case "dev_eui", "devEui":
			s.AddField("dev_eui")
			x.DevEui = types.Unmarshal8Bytes(s.WithField("dev_eui", false))
		case "join_eui", "joinEui":
			s.AddField("join_eui")
			x.JoinEui = types.Unmarshal8Bytes(s.WithField("join_eui", false))
		case "dev_addr", "devAddr":
			s.AddField("dev_addr")
			x.DevAddr = types.Unmarshal4Bytes(s.WithField("dev_addr", false))
		}
	})
}

// UnmarshalJSON unmarshals the EndDeviceIdentifiers from JSON.
func (x *EndDeviceIdentifiers) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the GatewayIdentifiers message to JSON.
func (x *GatewayIdentifiers) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.GatewayId != "" || s.HasField("gateway_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("gateway_id")
		s.WriteString(x.GatewayId)
	}
	if len(x.Eui) > 0 || s.HasField("eui") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("eui")
		types.MarshalHEXBytes(s.WithField("eui"), x.Eui)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the GatewayIdentifiers to JSON.
func (x *GatewayIdentifiers) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the GatewayIdentifiers message from JSON.
func (x *GatewayIdentifiers) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "gateway_id", "gatewayId":
			s.AddField("gateway_id")
			x.GatewayId = s.ReadString()
		case "eui":
			s.AddField("eui")
			x.Eui = types.Unmarshal8Bytes(s.WithField("eui", false))
		}
	})
}

// UnmarshalJSON unmarshals the GatewayIdentifiers from JSON.
func (x *GatewayIdentifiers) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the EntityIdentifiers message to JSON.
func (x *EntityIdentifiers) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if x.Ids != nil {
		switch ov := x.Ids.(type) {
		case *EntityIdentifiers_ApplicationIds:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("application_ids")
			// NOTE: ApplicationIdentifiers does not seem to implement MarshalProtoJSON.
			golang.MarshalMessage(s, ov.ApplicationIds)
		case *EntityIdentifiers_ClientIds:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("client_ids")
			// NOTE: ClientIdentifiers does not seem to implement MarshalProtoJSON.
			golang.MarshalMessage(s, ov.ClientIds)
		case *EntityIdentifiers_DeviceIds:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("device_ids")
			ov.DeviceIds.MarshalProtoJSON(s.WithField("device_ids"))
		case *EntityIdentifiers_GatewayIds:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("gateway_ids")
			ov.GatewayIds.MarshalProtoJSON(s.WithField("gateway_ids"))
		case *EntityIdentifiers_OrganizationIds:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("organization_ids")
			// NOTE: OrganizationIdentifiers does not seem to implement MarshalProtoJSON.
			golang.MarshalMessage(s, ov.OrganizationIds)
		case *EntityIdentifiers_UserIds:
			s.WriteMoreIf(&wroteField)
			s.WriteObjectField("user_ids")
			// NOTE: UserIdentifiers does not seem to implement MarshalProtoJSON.
			golang.MarshalMessage(s, ov.UserIds)
		}
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the EntityIdentifiers to JSON.
func (x *EntityIdentifiers) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the EntityIdentifiers message from JSON.
func (x *EntityIdentifiers) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "application_ids", "applicationIds":
			s.AddField("application_ids")
			ov := &EntityIdentifiers_ApplicationIds{}
			x.Ids = ov
			if s.ReadNil() {
				ov.ApplicationIds = nil
				return
			}
			// NOTE: ApplicationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ApplicationIdentifiers
			golang.UnmarshalMessage(s, &v)
			ov.ApplicationIds = &v
		case "client_ids", "clientIds":
			s.AddField("client_ids")
			ov := &EntityIdentifiers_ClientIds{}
			x.Ids = ov
			if s.ReadNil() {
				ov.ClientIds = nil
				return
			}
			// NOTE: ClientIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v ClientIdentifiers
			golang.UnmarshalMessage(s, &v)
			ov.ClientIds = &v
		case "device_ids", "deviceIds":
			ov := &EntityIdentifiers_DeviceIds{}
			x.Ids = ov
			if s.ReadNil() {
				ov.DeviceIds = nil
				return
			}
			ov.DeviceIds = &EndDeviceIdentifiers{}
			ov.DeviceIds.UnmarshalProtoJSON(s.WithField("device_ids", true))
		case "gateway_ids", "gatewayIds":
			ov := &EntityIdentifiers_GatewayIds{}
			x.Ids = ov
			if s.ReadNil() {
				ov.GatewayIds = nil
				return
			}
			ov.GatewayIds = &GatewayIdentifiers{}
			ov.GatewayIds.UnmarshalProtoJSON(s.WithField("gateway_ids", true))
		case "organization_ids", "organizationIds":
			s.AddField("organization_ids")
			ov := &EntityIdentifiers_OrganizationIds{}
			x.Ids = ov
			if s.ReadNil() {
				ov.OrganizationIds = nil
				return
			}
			// NOTE: OrganizationIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v OrganizationIdentifiers
			golang.UnmarshalMessage(s, &v)
			ov.OrganizationIds = &v
		case "user_ids", "userIds":
			s.AddField("user_ids")
			ov := &EntityIdentifiers_UserIds{}
			x.Ids = ov
			if s.ReadNil() {
				ov.UserIds = nil
				return
			}
			// NOTE: UserIdentifiers does not seem to implement UnmarshalProtoJSON.
			var v UserIdentifiers
			golang.UnmarshalMessage(s, &v)
			ov.UserIds = &v
		}
	})
}

// UnmarshalJSON unmarshals the EntityIdentifiers from JSON.
func (x *EntityIdentifiers) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}

// MarshalProtoJSON marshals the NetworkIdentifiers message to JSON.
func (x *NetworkIdentifiers) MarshalProtoJSON(s *jsonplugin.MarshalState) {
	if x == nil {
		s.WriteNil()
		return
	}
	s.WriteObjectStart()
	var wroteField bool
	if len(x.NetId) > 0 || s.HasField("net_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("net_id")
		types.MarshalHEXBytes(s.WithField("net_id"), x.NetId)
	}
	if x.TenantId != "" || s.HasField("tenant_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("tenant_id")
		s.WriteString(x.TenantId)
	}
	if x.ClusterId != "" || s.HasField("cluster_id") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("cluster_id")
		s.WriteString(x.ClusterId)
	}
	if x.ClusterAddress != "" || s.HasField("cluster_address") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("cluster_address")
		s.WriteString(x.ClusterAddress)
	}
	if x.TenantAddress != "" || s.HasField("tenant_address") {
		s.WriteMoreIf(&wroteField)
		s.WriteObjectField("tenant_address")
		s.WriteString(x.TenantAddress)
	}
	s.WriteObjectEnd()
}

// MarshalJSON marshals the NetworkIdentifiers to JSON.
func (x *NetworkIdentifiers) MarshalJSON() ([]byte, error) {
	return jsonplugin.DefaultMarshalerConfig.Marshal(x)
}

// UnmarshalProtoJSON unmarshals the NetworkIdentifiers message from JSON.
func (x *NetworkIdentifiers) UnmarshalProtoJSON(s *jsonplugin.UnmarshalState) {
	if s.ReadNil() {
		return
	}
	s.ReadObject(func(key string) {
		switch key {
		default:
			s.ReadAny() // ignore unknown field
		case "net_id", "netId":
			s.AddField("net_id")
			x.NetId = types.Unmarshal3Bytes(s.WithField("net_id", false))
		case "tenant_id", "tenantId":
			s.AddField("tenant_id")
			x.TenantId = s.ReadString()
		case "cluster_id", "clusterId":
			s.AddField("cluster_id")
			x.ClusterId = s.ReadString()
		case "cluster_address", "clusterAddress":
			s.AddField("cluster_address")
			x.ClusterAddress = s.ReadString()
		case "tenant_address", "tenantAddress":
			s.AddField("tenant_address")
			x.TenantAddress = s.ReadString()
		}
	})
}

// UnmarshalJSON unmarshals the NetworkIdentifiers from JSON.
func (x *NetworkIdentifiers) UnmarshalJSON(b []byte) error {
	return jsonplugin.DefaultUnmarshalerConfig.Unmarshal(b, x)
}
