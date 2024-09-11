// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type CreateDegraphqlPluginConfig struct {
	// A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).
	GraphqlServerPath *string `json:"graphql_server_path,omitempty"`
}

func (o *CreateDegraphqlPluginConfig) GetGraphqlServerPath() *string {
	if o == nil {
		return nil
	}
	return o.GraphqlServerPath
}

type CreateDegraphqlPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateDegraphqlPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateDegraphqlPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateDegraphqlPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateDegraphqlPluginOrdering struct {
	After  *CreateDegraphqlPluginAfter  `json:"after,omitempty"`
	Before *CreateDegraphqlPluginBefore `json:"before,omitempty"`
}

func (o *CreateDegraphqlPluginOrdering) GetAfter() *CreateDegraphqlPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *CreateDegraphqlPluginOrdering) GetBefore() *CreateDegraphqlPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type CreateDegraphqlPluginProtocols string

const (
	CreateDegraphqlPluginProtocolsGrpc           CreateDegraphqlPluginProtocols = "grpc"
	CreateDegraphqlPluginProtocolsGrpcs          CreateDegraphqlPluginProtocols = "grpcs"
	CreateDegraphqlPluginProtocolsHTTP           CreateDegraphqlPluginProtocols = "http"
	CreateDegraphqlPluginProtocolsHTTPS          CreateDegraphqlPluginProtocols = "https"
	CreateDegraphqlPluginProtocolsTCP            CreateDegraphqlPluginProtocols = "tcp"
	CreateDegraphqlPluginProtocolsTLS            CreateDegraphqlPluginProtocols = "tls"
	CreateDegraphqlPluginProtocolsTLSPassthrough CreateDegraphqlPluginProtocols = "tls_passthrough"
	CreateDegraphqlPluginProtocolsUDP            CreateDegraphqlPluginProtocols = "udp"
	CreateDegraphqlPluginProtocolsWs             CreateDegraphqlPluginProtocols = "ws"
	CreateDegraphqlPluginProtocolsWss            CreateDegraphqlPluginProtocols = "wss"
)

func (e CreateDegraphqlPluginProtocols) ToPointer() *CreateDegraphqlPluginProtocols {
	return &e
}
func (e *CreateDegraphqlPluginProtocols) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "grpc":
		fallthrough
	case "grpcs":
		fallthrough
	case "http":
		fallthrough
	case "https":
		fallthrough
	case "tcp":
		fallthrough
	case "tls":
		fallthrough
	case "tls_passthrough":
		fallthrough
	case "udp":
		fallthrough
	case "ws":
		fallthrough
	case "wss":
		*e = CreateDegraphqlPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDegraphqlPluginProtocols: %v", v)
	}
}

// CreateDegraphqlPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateDegraphqlPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateDegraphqlPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateDegraphqlPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateDegraphqlPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateDegraphqlPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateDegraphqlPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateDegraphqlPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateDegraphqlPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateDegraphqlPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateDegraphqlPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateDegraphqlPlugin struct {
	Config *CreateDegraphqlPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                          `json:"enabled,omitempty"`
	InstanceName *string                        `json:"instance_name,omitempty"`
	name         *string                        `const:"degraphql" json:"name,omitempty"`
	Ordering     *CreateDegraphqlPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateDegraphqlPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateDegraphqlPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateDegraphqlPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateDegraphqlPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateDegraphqlPluginService `json:"service,omitempty"`
}

func (c CreateDegraphqlPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateDegraphqlPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateDegraphqlPlugin) GetConfig() *CreateDegraphqlPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateDegraphqlPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateDegraphqlPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateDegraphqlPlugin) GetName() *string {
	return types.String("degraphql")
}

func (o *CreateDegraphqlPlugin) GetOrdering() *CreateDegraphqlPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateDegraphqlPlugin) GetProtocols() []CreateDegraphqlPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateDegraphqlPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateDegraphqlPlugin) GetConsumer() *CreateDegraphqlPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateDegraphqlPlugin) GetConsumerGroup() *CreateDegraphqlPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateDegraphqlPlugin) GetRoute() *CreateDegraphqlPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateDegraphqlPlugin) GetService() *CreateDegraphqlPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}