// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginAzureFunctionsResourceModel) ToSharedCreateAzureFunctionsPlugin() *shared.CreateAzureFunctionsPlugin {
	var config *shared.CreateAzureFunctionsPluginConfig
	if r.Config != nil {
		apikey := new(string)
		if !r.Config.Apikey.IsUnknown() && !r.Config.Apikey.IsNull() {
			*apikey = r.Config.Apikey.ValueString()
		} else {
			apikey = nil
		}
		appname := new(string)
		if !r.Config.Appname.IsUnknown() && !r.Config.Appname.IsNull() {
			*appname = r.Config.Appname.ValueString()
		} else {
			appname = nil
		}
		clientid := new(string)
		if !r.Config.Clientid.IsUnknown() && !r.Config.Clientid.IsNull() {
			*clientid = r.Config.Clientid.ValueString()
		} else {
			clientid = nil
		}
		functionname := new(string)
		if !r.Config.Functionname.IsUnknown() && !r.Config.Functionname.IsNull() {
			*functionname = r.Config.Functionname.ValueString()
		} else {
			functionname = nil
		}
		hostdomain := new(string)
		if !r.Config.Hostdomain.IsUnknown() && !r.Config.Hostdomain.IsNull() {
			*hostdomain = r.Config.Hostdomain.ValueString()
		} else {
			hostdomain = nil
		}
		https := new(bool)
		if !r.Config.HTTPS.IsUnknown() && !r.Config.HTTPS.IsNull() {
			*https = r.Config.HTTPS.ValueBool()
		} else {
			https = nil
		}
		httpsVerify := new(bool)
		if !r.Config.HTTPSVerify.IsUnknown() && !r.Config.HTTPSVerify.IsNull() {
			*httpsVerify = r.Config.HTTPSVerify.ValueBool()
		} else {
			httpsVerify = nil
		}
		keepalive := new(float64)
		if !r.Config.Keepalive.IsUnknown() && !r.Config.Keepalive.IsNull() {
			*keepalive, _ = r.Config.Keepalive.ValueBigFloat().Float64()
		} else {
			keepalive = nil
		}
		routeprefix := new(string)
		if !r.Config.Routeprefix.IsUnknown() && !r.Config.Routeprefix.IsNull() {
			*routeprefix = r.Config.Routeprefix.ValueString()
		} else {
			routeprefix = nil
		}
		timeout := new(float64)
		if !r.Config.Timeout.IsUnknown() && !r.Config.Timeout.IsNull() {
			*timeout, _ = r.Config.Timeout.ValueBigFloat().Float64()
		} else {
			timeout = nil
		}
		config = &shared.CreateAzureFunctionsPluginConfig{
			Apikey:       apikey,
			Appname:      appname,
			Clientid:     clientid,
			Functionname: functionname,
			Hostdomain:   hostdomain,
			HTTPS:        https,
			HTTPSVerify:  httpsVerify,
			Keepalive:    keepalive,
			Routeprefix:  routeprefix,
			Timeout:      timeout,
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
	var ordering *shared.CreateAzureFunctionsPluginOrdering
	if r.Ordering != nil {
		var after *shared.CreateAzureFunctionsPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CreateAzureFunctionsPluginAfter{
				Access: access,
			}
		}
		var before *shared.CreateAzureFunctionsPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CreateAzureFunctionsPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CreateAzureFunctionsPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CreateAzureFunctionsPluginProtocols = []shared.CreateAzureFunctionsPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateAzureFunctionsPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateAzureFunctionsPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateAzureFunctionsPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateAzureFunctionsPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateAzureFunctionsPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateAzureFunctionsPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateAzureFunctionsPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateAzureFunctionsPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateAzureFunctionsPluginService{
			ID: id3,
		}
	}
	out := shared.CreateAzureFunctionsPlugin{
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

func (r *GatewayPluginAzureFunctionsResourceModel) RefreshFromSharedAzureFunctionsPlugin(resp *shared.AzureFunctionsPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateAzureFunctionsPluginConfig{}
			r.Config.Apikey = types.StringPointerValue(resp.Config.Apikey)
			r.Config.Appname = types.StringPointerValue(resp.Config.Appname)
			r.Config.Clientid = types.StringPointerValue(resp.Config.Clientid)
			r.Config.Functionname = types.StringPointerValue(resp.Config.Functionname)
			r.Config.Hostdomain = types.StringPointerValue(resp.Config.Hostdomain)
			r.Config.HTTPS = types.BoolPointerValue(resp.Config.HTTPS)
			r.Config.HTTPSVerify = types.BoolPointerValue(resp.Config.HTTPSVerify)
			if resp.Config.Keepalive != nil {
				r.Config.Keepalive = types.NumberValue(big.NewFloat(float64(*resp.Config.Keepalive)))
			} else {
				r.Config.Keepalive = types.NumberNull()
			}
			r.Config.Routeprefix = types.StringPointerValue(resp.Config.Routeprefix)
			if resp.Config.Timeout != nil {
				r.Config.Timeout = types.NumberValue(big.NewFloat(float64(*resp.Config.Timeout)))
			} else {
				r.Config.Timeout = types.NumberNull()
			}
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