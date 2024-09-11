// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginHmacAuthResourceModel) ToSharedCreateHmacAuthPlugin() *shared.CreateHmacAuthPlugin {
	var config *shared.CreateHmacAuthPluginConfig
	if r.Config != nil {
		var algorithms []shared.CreateHmacAuthPluginAlgorithms = []shared.CreateHmacAuthPluginAlgorithms{}
		for _, algorithmsItem := range r.Config.Algorithms {
			algorithms = append(algorithms, shared.CreateHmacAuthPluginAlgorithms(algorithmsItem.ValueString()))
		}
		anonymous := new(string)
		if !r.Config.Anonymous.IsUnknown() && !r.Config.Anonymous.IsNull() {
			*anonymous = r.Config.Anonymous.ValueString()
		} else {
			anonymous = nil
		}
		clockSkew := new(float64)
		if !r.Config.ClockSkew.IsUnknown() && !r.Config.ClockSkew.IsNull() {
			*clockSkew, _ = r.Config.ClockSkew.ValueBigFloat().Float64()
		} else {
			clockSkew = nil
		}
		var enforceHeaders []string = []string{}
		for _, enforceHeadersItem := range r.Config.EnforceHeaders {
			enforceHeaders = append(enforceHeaders, enforceHeadersItem.ValueString())
		}
		hideCredentials := new(bool)
		if !r.Config.HideCredentials.IsUnknown() && !r.Config.HideCredentials.IsNull() {
			*hideCredentials = r.Config.HideCredentials.ValueBool()
		} else {
			hideCredentials = nil
		}
		realm := new(string)
		if !r.Config.Realm.IsUnknown() && !r.Config.Realm.IsNull() {
			*realm = r.Config.Realm.ValueString()
		} else {
			realm = nil
		}
		validateRequestBody := new(bool)
		if !r.Config.ValidateRequestBody.IsUnknown() && !r.Config.ValidateRequestBody.IsNull() {
			*validateRequestBody = r.Config.ValidateRequestBody.ValueBool()
		} else {
			validateRequestBody = nil
		}
		config = &shared.CreateHmacAuthPluginConfig{
			Algorithms:          algorithms,
			Anonymous:           anonymous,
			ClockSkew:           clockSkew,
			EnforceHeaders:      enforceHeaders,
			HideCredentials:     hideCredentials,
			Realm:               realm,
			ValidateRequestBody: validateRequestBody,
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
	var ordering *shared.CreateHmacAuthPluginOrdering
	if r.Ordering != nil {
		var after *shared.CreateHmacAuthPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CreateHmacAuthPluginAfter{
				Access: access,
			}
		}
		var before *shared.CreateHmacAuthPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CreateHmacAuthPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CreateHmacAuthPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CreateHmacAuthPluginProtocols = []shared.CreateHmacAuthPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateHmacAuthPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateHmacAuthPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateHmacAuthPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateHmacAuthPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateHmacAuthPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateHmacAuthPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateHmacAuthPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateHmacAuthPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateHmacAuthPluginService{
			ID: id3,
		}
	}
	out := shared.CreateHmacAuthPlugin{
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

func (r *GatewayPluginHmacAuthResourceModel) RefreshFromSharedHmacAuthPlugin(resp *shared.HmacAuthPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateHmacAuthPluginConfig{}
			r.Config.Algorithms = []types.String{}
			for _, v := range resp.Config.Algorithms {
				r.Config.Algorithms = append(r.Config.Algorithms, types.StringValue(string(v)))
			}
			r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
			if resp.Config.ClockSkew != nil {
				r.Config.ClockSkew = types.NumberValue(big.NewFloat(float64(*resp.Config.ClockSkew)))
			} else {
				r.Config.ClockSkew = types.NumberNull()
			}
			r.Config.EnforceHeaders = []types.String{}
			for _, v := range resp.Config.EnforceHeaders {
				r.Config.EnforceHeaders = append(r.Config.EnforceHeaders, types.StringValue(v))
			}
			r.Config.HideCredentials = types.BoolPointerValue(resp.Config.HideCredentials)
			r.Config.Realm = types.StringPointerValue(resp.Config.Realm)
			r.Config.ValidateRequestBody = types.BoolPointerValue(resp.Config.ValidateRequestBody)
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
