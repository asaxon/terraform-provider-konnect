// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"math/big"
)

func (r *GatewayPluginLogglyResourceModel) ToSharedCreateLogglyPlugin() *shared.CreateLogglyPlugin {
	var config *shared.CreateLogglyPluginConfig
	if r.Config != nil {
		clientErrorsSeverity := new(shared.CreateLogglyPluginClientErrorsSeverity)
		if !r.Config.ClientErrorsSeverity.IsUnknown() && !r.Config.ClientErrorsSeverity.IsNull() {
			*clientErrorsSeverity = shared.CreateLogglyPluginClientErrorsSeverity(r.Config.ClientErrorsSeverity.ValueString())
		} else {
			clientErrorsSeverity = nil
		}
		customFieldsByLua := make(map[string]interface{})
		for customFieldsByLuaKey, customFieldsByLuaValue := range r.Config.CustomFieldsByLua {
			var customFieldsByLuaInst interface{}
			_ = json.Unmarshal([]byte(customFieldsByLuaValue.ValueString()), &customFieldsByLuaInst)
			customFieldsByLua[customFieldsByLuaKey] = customFieldsByLuaInst
		}
		host := new(string)
		if !r.Config.Host.IsUnknown() && !r.Config.Host.IsNull() {
			*host = r.Config.Host.ValueString()
		} else {
			host = nil
		}
		key := new(string)
		if !r.Config.Key.IsUnknown() && !r.Config.Key.IsNull() {
			*key = r.Config.Key.ValueString()
		} else {
			key = nil
		}
		logLevel := new(shared.CreateLogglyPluginLogLevel)
		if !r.Config.LogLevel.IsUnknown() && !r.Config.LogLevel.IsNull() {
			*logLevel = shared.CreateLogglyPluginLogLevel(r.Config.LogLevel.ValueString())
		} else {
			logLevel = nil
		}
		port := new(int64)
		if !r.Config.Port.IsUnknown() && !r.Config.Port.IsNull() {
			*port = r.Config.Port.ValueInt64()
		} else {
			port = nil
		}
		serverErrorsSeverity := new(shared.CreateLogglyPluginServerErrorsSeverity)
		if !r.Config.ServerErrorsSeverity.IsUnknown() && !r.Config.ServerErrorsSeverity.IsNull() {
			*serverErrorsSeverity = shared.CreateLogglyPluginServerErrorsSeverity(r.Config.ServerErrorsSeverity.ValueString())
		} else {
			serverErrorsSeverity = nil
		}
		successfulSeverity := new(shared.CreateLogglyPluginSuccessfulSeverity)
		if !r.Config.SuccessfulSeverity.IsUnknown() && !r.Config.SuccessfulSeverity.IsNull() {
			*successfulSeverity = shared.CreateLogglyPluginSuccessfulSeverity(r.Config.SuccessfulSeverity.ValueString())
		} else {
			successfulSeverity = nil
		}
		var tags []string = []string{}
		for _, tagsItem := range r.Config.Tags {
			tags = append(tags, tagsItem.ValueString())
		}
		timeout := new(float64)
		if !r.Config.Timeout.IsUnknown() && !r.Config.Timeout.IsNull() {
			*timeout, _ = r.Config.Timeout.ValueBigFloat().Float64()
		} else {
			timeout = nil
		}
		config = &shared.CreateLogglyPluginConfig{
			ClientErrorsSeverity: clientErrorsSeverity,
			CustomFieldsByLua:    customFieldsByLua,
			Host:                 host,
			Key:                  key,
			LogLevel:             logLevel,
			Port:                 port,
			ServerErrorsSeverity: serverErrorsSeverity,
			SuccessfulSeverity:   successfulSeverity,
			Tags:                 tags,
			Timeout:              timeout,
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
	var ordering *shared.CreateLogglyPluginOrdering
	if r.Ordering != nil {
		var after *shared.CreateLogglyPluginAfter
		if r.Ordering.After != nil {
			var access []string = []string{}
			for _, accessItem := range r.Ordering.After.Access {
				access = append(access, accessItem.ValueString())
			}
			after = &shared.CreateLogglyPluginAfter{
				Access: access,
			}
		}
		var before *shared.CreateLogglyPluginBefore
		if r.Ordering.Before != nil {
			var access1 []string = []string{}
			for _, accessItem1 := range r.Ordering.Before.Access {
				access1 = append(access1, accessItem1.ValueString())
			}
			before = &shared.CreateLogglyPluginBefore{
				Access: access1,
			}
		}
		ordering = &shared.CreateLogglyPluginOrdering{
			After:  after,
			Before: before,
		}
	}
	var protocols []shared.CreateLogglyPluginProtocols = []shared.CreateLogglyPluginProtocols{}
	for _, protocolsItem := range r.Protocols {
		protocols = append(protocols, shared.CreateLogglyPluginProtocols(protocolsItem.ValueString()))
	}
	var tags1 []string = []string{}
	for _, tagsItem1 := range r.Tags {
		tags1 = append(tags1, tagsItem1.ValueString())
	}
	var consumer *shared.CreateLogglyPluginConsumer
	if r.Consumer != nil {
		id := new(string)
		if !r.Consumer.ID.IsUnknown() && !r.Consumer.ID.IsNull() {
			*id = r.Consumer.ID.ValueString()
		} else {
			id = nil
		}
		consumer = &shared.CreateLogglyPluginConsumer{
			ID: id,
		}
	}
	var consumerGroup *shared.CreateLogglyPluginConsumerGroup
	if r.ConsumerGroup != nil {
		id1 := new(string)
		if !r.ConsumerGroup.ID.IsUnknown() && !r.ConsumerGroup.ID.IsNull() {
			*id1 = r.ConsumerGroup.ID.ValueString()
		} else {
			id1 = nil
		}
		consumerGroup = &shared.CreateLogglyPluginConsumerGroup{
			ID: id1,
		}
	}
	var route *shared.CreateLogglyPluginRoute
	if r.Route != nil {
		id2 := new(string)
		if !r.Route.ID.IsUnknown() && !r.Route.ID.IsNull() {
			*id2 = r.Route.ID.ValueString()
		} else {
			id2 = nil
		}
		route = &shared.CreateLogglyPluginRoute{
			ID: id2,
		}
	}
	var service *shared.CreateLogglyPluginService
	if r.Service != nil {
		id3 := new(string)
		if !r.Service.ID.IsUnknown() && !r.Service.ID.IsNull() {
			*id3 = r.Service.ID.ValueString()
		} else {
			id3 = nil
		}
		service = &shared.CreateLogglyPluginService{
			ID: id3,
		}
	}
	out := shared.CreateLogglyPlugin{
		Config:        config,
		Enabled:       enabled,
		InstanceName:  instanceName,
		Ordering:      ordering,
		Protocols:     protocols,
		Tags:          tags1,
		Consumer:      consumer,
		ConsumerGroup: consumerGroup,
		Route:         route,
		Service:       service,
	}
	return &out
}

func (r *GatewayPluginLogglyResourceModel) RefreshFromSharedLogglyPlugin(resp *shared.LogglyPlugin) {
	if resp != nil {
		if resp.Config == nil {
			r.Config = nil
		} else {
			r.Config = &tfTypes.CreateLogglyPluginConfig{}
			if resp.Config.ClientErrorsSeverity != nil {
				r.Config.ClientErrorsSeverity = types.StringValue(string(*resp.Config.ClientErrorsSeverity))
			} else {
				r.Config.ClientErrorsSeverity = types.StringNull()
			}
			if len(resp.Config.CustomFieldsByLua) > 0 {
				r.Config.CustomFieldsByLua = make(map[string]types.String)
				for key, value := range resp.Config.CustomFieldsByLua {
					result, _ := json.Marshal(value)
					r.Config.CustomFieldsByLua[key] = types.StringValue(string(result))
				}
			}
			r.Config.Host = types.StringPointerValue(resp.Config.Host)
			r.Config.Key = types.StringPointerValue(resp.Config.Key)
			if resp.Config.LogLevel != nil {
				r.Config.LogLevel = types.StringValue(string(*resp.Config.LogLevel))
			} else {
				r.Config.LogLevel = types.StringNull()
			}
			r.Config.Port = types.Int64PointerValue(resp.Config.Port)
			if resp.Config.ServerErrorsSeverity != nil {
				r.Config.ServerErrorsSeverity = types.StringValue(string(*resp.Config.ServerErrorsSeverity))
			} else {
				r.Config.ServerErrorsSeverity = types.StringNull()
			}
			if resp.Config.SuccessfulSeverity != nil {
				r.Config.SuccessfulSeverity = types.StringValue(string(*resp.Config.SuccessfulSeverity))
			} else {
				r.Config.SuccessfulSeverity = types.StringNull()
			}
			r.Config.Tags = []types.String{}
			for _, v := range resp.Config.Tags {
				r.Config.Tags = append(r.Config.Tags, types.StringValue(v))
			}
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
