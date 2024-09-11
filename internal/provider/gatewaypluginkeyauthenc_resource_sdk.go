// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *GatewayPluginKeyAuthEncResourceModel) ToSharedCreateKeyAuthEncPlugin() *shared.CreateKeyAuthEncPlugin {
	var config *shared.CreateKeyAuthEncPluginConfig
	if r.Config != nil {
		anonymous := new(string)
		if !r.Config.Anonymous.IsUnknown() && !r.Config.Anonymous.IsNull() {
			*anonymous = r.Config.Anonymous.ValueString()
		} else {
			anonymous = nil
		}
		hideCredentials := new(bool)
		if !r.Config.HideCredentials.IsUnknown() && !r.Config.HideCredentials.IsNull() {
			*hideCredentials = r.Config.HideCredentials.ValueBool()
		} else {
			hideCredentials = nil
		}
		keyInBody := new(bool)
		if !r.Config.KeyInBody.IsUnknown() && !r.Config.KeyInBody.IsNull() {
			*keyInBody = r.Config.KeyInBody.ValueBool()
		} else {
			keyInBody = nil
		}
		keyInHeader := new(bool)
		if !r.Config.KeyInHeader.IsUnknown() && !r.Config.KeyInHeader.IsNull() {
			*keyInHeader = r.Config.KeyInHeader.ValueBool()
		} else {
			keyInHeader = nil
		}
		keyInQuery := new(bool)
		if !r.Config.KeyInQuery.IsUnknown() && !r.Config.KeyInQuery.IsNull() {
			*keyInQuery = r.Config.KeyInQuery.ValueBool()
		} else {
			keyInQuery = nil
		}
		var keyNames []string = []string{}
		for _, keyNamesItem := range r.Config.KeyNames {
			keyNames = append(keyNames, keyNamesItem.ValueString())
		}
		realm := new(string)
		if !r.Config.Realm.IsUnknown() && !r.Config.Realm.IsNull() {
			*realm = r.Config.Realm.ValueString()
		} else {
			realm = nil
		}
		runOnPreflight := new(bool)
		if !r.Config.RunOnPreflight.IsUnknown() && !r.Config.RunOnPreflight.IsNull() {
			*runOnPreflight = r.Config.RunOnPreflight.ValueBool()
		} else {
			runOnPreflight = nil
		}
		config = &shared.CreateKeyAuthEncPluginConfig{
			Anonymous:       anonymous,
			HideCredentials: hideCredentials,
			KeyInBody:       keyInBody,
			KeyInHeader:     keyInHeader,
			KeyInQuery:      keyInQuery,
			KeyNames:        keyNames,
			Realm:           realm,
			RunOnPreflight:  runOnPreflight,
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
	var ordering *shared.CreateKeyAuthEncPluginOrdering
	if r.Ordering != nil {
		var after *shared.CreateKeyAuthEncPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CreateKeyAuthEncPluginAfter{
				Access: access,
			}
		}
		var before *shared.CreateKeyAuthEncPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CreateKeyAuthEncPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CreateKeyAuthEncPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CreateKeyAuthEncPluginProtocols = []shared.CreateKeyAuthEncPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateKeyAuthEncPluginProtocols(protocolsItem.ValueString()))
	}
	var tags []string = []string{}
	for _, tagsItem := range r.Tags {
		tags = append(tags, tagsItem.ValueString())
	}
	var consumer *shared.CreateKeyAuthEncPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateKeyAuthEncPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateKeyAuthEncPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateKeyAuthEncPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateKeyAuthEncPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateKeyAuthEncPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateKeyAuthEncPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateKeyAuthEncPluginService{
			ID: id3,
		}
	}
	out := shared.CreateKeyAuthEncPlugin{
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

func (r *GatewayPluginKeyAuthEncResourceModel) RefreshFromSharedKeyAuthEncPlugin(resp *shared.KeyAuthEncPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateKeyAuthPluginConfig{}
			r.Config.Anonymous = types.StringPointerValue(resp.Config.Anonymous)
			r.Config.HideCredentials = types.BoolPointerValue(resp.Config.HideCredentials)
			r.Config.KeyInBody = types.BoolPointerValue(resp.Config.KeyInBody)
			r.Config.KeyInHeader = types.BoolPointerValue(resp.Config.KeyInHeader)
			r.Config.KeyInQuery = types.BoolPointerValue(resp.Config.KeyInQuery)
			r.Config.KeyNames = []types.String{}
			for _, v := range resp.Config.KeyNames {
				r.Config.KeyNames = append(r.Config.KeyNames, types.StringValue(v))
			}
			r.Config.Realm = types.StringPointerValue(resp.Config.Realm)
			r.Config.RunOnPreflight = types.BoolPointerValue(resp.Config.RunOnPreflight)
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
