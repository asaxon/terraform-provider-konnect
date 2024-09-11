// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginLdapAuthResourceModel) ToSharedCreateLdapAuthPlugin() *shared.CreateLdapAuthPlugin {
	var config *shared.CreateLdapAuthPluginConfig
	if r.Config != nil {
		anonymous := new(string)
		if !r.Config.Anonymous.IsUnknown() && !r.Config.Anonymous.IsNull() {
			*anonymous = r.Config.Anonymous.ValueString()
		} else {
			anonymous = nil
		}
		attribute := new(string)
		if !r.Config.Attribute.IsUnknown() && !r.Config.Attribute.IsNull() {
			*attribute = r.Config.Attribute.ValueString()
		} else {
			attribute = nil
		}
		baseDn := new(string)
		if !r.Config.BaseDn.IsUnknown() && !r.Config.BaseDn.IsNull() {
			*baseDn = r.Config.BaseDn.ValueString()
		} else {
			baseDn = nil
		}
		cacheTTL := new(float64)
		if !r.Config.CacheTTL.IsUnknown() && !r.Config.CacheTTL.IsNull() {
			*cacheTTL, _ = r.Config.CacheTTL.ValueBigFloat().Float64()
		} else {
			cacheTTL = nil
		}
		headerType := new(string)
		if !r.Config.HeaderType.IsUnknown() && !r.Config.HeaderType.IsNull() {
			*headerType = r.Config.HeaderType.ValueString()
		} else {
			headerType = nil
		}
		hideCredentials := new(bool)
		if !r.Config.HideCredentials.IsUnknown() && !r.Config.HideCredentials.IsNull() {
			*hideCredentials = r.Config.HideCredentials.ValueBool()
		} else {
			hideCredentials = nil
		}
		keepalive := new(float64)
		if !r.Config.Keepalive.IsUnknown() && !r.Config.Keepalive.IsNull() {
			*keepalive, _ = r.Config.Keepalive.ValueBigFloat().Float64()
		} else {
			keepalive = nil
		}
		ldapHost := new(string)
		if !r.Config.LdapHost.IsUnknown() && !r.Config.LdapHost.IsNull() {
			*ldapHost = r.Config.LdapHost.ValueString()
		} else {
			ldapHost = nil
		}
		ldapPort := new(int64)
		if !r.Config.LdapPort.IsUnknown() && !r.Config.LdapPort.IsNull() {
			*ldapPort = r.Config.LdapPort.ValueInt64()
		} else {
			ldapPort = nil
		}
		ldaps := new(bool)
		if !r.Config.Ldaps.IsUnknown() && !r.Config.Ldaps.IsNull() {
			*ldaps = r.Config.Ldaps.ValueBool()
		} else {
			ldaps = nil
		}
		realm := new(string)
		if !r.Config.Realm.IsUnknown() && !r.Config.Realm.IsNull() {
			*realm = r.Config.Realm.ValueString()
		} else {
			realm = nil
		}
		startTLS := new(bool)
		if !r.Config.StartTLS.IsUnknown() && !r.Config.StartTLS.IsNull() {
			*startTLS = r.Config.StartTLS.ValueBool()
		} else {
			startTLS = nil
		}
		timeout := new(float64)
		if !r.Config.Timeout.IsUnknown() && !r.Config.Timeout.IsNull() {
			*timeout, _ = r.Config.Timeout.ValueBigFloat().Float64()
		} else {
			timeout = nil
		}
		verifyLdapHost := new(bool)
		if !r.Config.VerifyLdapHost.IsUnknown() && !r.Config.VerifyLdapHost.IsNull() {
			*verifyLdapHost = r.Config.VerifyLdapHost.ValueBool()
		} else {
			verifyLdapHost = nil
		}
		config = &shared.CreateLdapAuthPluginConfig{
			Anonymous:       anonymous,
			Attribute:       attribute,
			BaseDn:          baseDn,
			CacheTTL:        cacheTTL,
			HeaderType:      headerType,
			HideCredentials: hideCredentials,
			Keepalive:       keepalive,
			LdapHost:        ldapHost,
			LdapPort:        ldapPort,
			Ldaps:           ldaps,
			Realm:           realm,
			StartTLS:        startTLS,
			Timeout:         timeout,
			VerifyLdapHost:  verifyLdapHost,
		}
	}
	enabled := new(bool)
	if !r.Enabled.IsUnknown() && !r.Enabled.IsNull() {
		*enabled = r.Enabled.ValueBool()
	} else {
		enabled = nil
	}
	instanceName := new(string)
	if !r.InstanceName.IsUnknown() && !r.InstanceName.IsNull() {
		*instanceName = r.InstanceName.ValueString()
	} else {
		instanceName = nil
	}
	var ordering *shared.CreateLdapAuthPluginOrdering
	if r.Ordering != nil {
		var after *shared.CreateLdapAuthPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CreateLdapAuthPluginAfter{
				Access: access,
			}
		}
		var before *shared.CreateLdapAuthPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CreateLdapAuthPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CreateLdapAuthPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CreateLdapAuthPluginProtocols = []shared.CreateLdapAuthPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateLdapAuthPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateLdapAuthPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateLdapAuthPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateLdapAuthPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateLdapAuthPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateLdapAuthPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateLdapAuthPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateLdapAuthPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateLdapAuthPluginService{
			ID: id3,
		}
	}
	out := shared.CreateLdapAuthPlugin{
		Config:        config,
		Enabled:       enabled,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Protocols:     protocols,
		Tags:          tags,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *GatewayPluginLdapAuthResourceModel) RefreshFromSharedLdapAuthPlugin(resp *shared.LdapAuthPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateLdapAuthPluginConfig{}
			r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
			r.Config.Attribute = types.StringPointerValue(resp.Config.Attribute)
			r.Config.BaseDn = types.StringPointerValue(resp.Config.BaseDn)
			if resp.Config.CacheTTL != nil {
				r.Config.CacheTTL = types.NumberValue(big.NewFloat(float64(*resp.Config.CacheTTL)))
			} else {
				r.Config.CacheTTL = types.NumberNull()
			}
			r.Config.HeaderType = types.StringPointerValue(resp.Config.HeaderType)
			r.Config.HideCredentials = types.BoolPointerValue(resp.Config.HideCredentials)
			if resp.Config.Keepalive != nil {
				r.Config.Keepalive = types.NumberValue(big.NewFloat(float64(*resp.Config.Keepalive)))
			} else {
				r.Config.Keepalive = types.NumberNull()
			}
			r.Config.LdapHost = types.StringPointerValue(resp.Config.LdapHost)
			r.Config.LdapPort = types.Int64PointerValue(resp.Config.LdapPort)
			r.Config.Ldaps = types.BoolPointerValue(resp.Config.Ldaps)
			r.Config.Realm = types.StringPointerValue(resp.Config.Realm)
			r.Config.StartTLS = types.BoolPointerValue(resp.Config.StartTLS)
			if resp.Config.Timeout != nil {
				r.Config.Timeout = types.NumberValue(big.NewFloat(float64(*resp.Config.Timeout)))
			} else {
				r.Config.Timeout = types.NumberNull()
			}
			r.Config.VerifyLdapHost = types.BoolPointerValue(resp.Config.VerifyLdapHost)
		}
		if resp.Consumer == nil {
			r.Consumer = nil
		} else {
			r.Consumer = &tfTypes.ACLConsumer{}
			r.Consumer.ID = types.StringPointerValue(resp.Consumer.ID)
		}
		if resp.ConsumerGroup == nil {
			r.ConsumerGroup = nil
		} else {
			r.ConsumerGroup = &tfTypes.ACLConsumer{}
			r.ConsumerGroup.ID = types.StringPointerValue(resp.ConsumerGroup.ID)
		}
		r.CreatedAt = types.Int64PointerValue(resp.CreatedAt)
		r.Enabled = types.BoolPointerValue(resp.Enabled)
		r.ID = types.StringPointerValue(resp.ID)
		r.InstanceName = types.StringPointerValue(resp.InstanceName)
		if resp.Ordering == nil {
			r.Ordering = nil
		} else {
			r.Ordering = &tfTypes.CreateACLPluginOrdering{}
			if resp.Ordering.After == nil {
				r.Ordering.After = nil
			} else {
				r.Ordering.After = &tfTypes.CreateACLPluginAfter{}
				r.Ordering.After.Access = []types.String{}
				for _, v := range resp.Ordering.After.Access {
					r.Ordering.After.Access = append(r.Ordering.After.Access, types.StringValue(v))
				}
			}
			if resp.Ordering.Before == nil {
				r.Ordering.Before = nil
			} else {
				r.Ordering.Before = &tfTypes.CreateACLPluginAfter{}
				r.Ordering.Before.Access = []types.String{}
				for _, v := range resp.Ordering.Before.Access {
					r.Ordering.Before.Access = append(r.Ordering.Before.Access, types.StringValue(v))
				}
			}
		}
		r.Protocols = []types.String{}
		for _, v := range resp.Protocols {
			r.Protocols = append(r.Protocols, types.StringValue(string(v)))
		}
		if resp.Route == nil {
			r.Route = nil
		} else {
			r.Route = &tfTypes.ACLConsumer{}
			r.Route.ID = types.StringPointerValue(resp.Route.ID)
		}
		if resp.Service == nil {
			r.Service = nil
		} else {
			r.Service = &tfTypes.ACLConsumer{}
			r.Service.ID = types.StringPointerValue(resp.Service.ID)
		}
		r.Tags = []types.String{}
		for _, v := range resp.Tags {
			r.Tags = append(r.Tags, types.StringValue(v))
		}
		r.UpdatedAt = types.Int64PointerValue(resp.UpdatedAt)
	}
}
