// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type CreateRequestTerminationPluginConfig struct {
	// The raw response body to send. This is mutually exclusive with the `config.message` field.
	Body *string `json:"body,omitempty"`
	// Content type of the raw response configured with `config.body`.
	ContentType *string `json:"content_type,omitempty"`
	// When set, the plugin will echo a copy of the request back to the client. The main usecase for this is debugging. It can be combined with `trigger` in order to debug requests on live systems without disturbing real traffic.
	Echo *bool `json:"echo,omitempty"`
	// The message to send, if using the default response generator.
	Message *string `json:"message,omitempty"`
	// The response code to send. Must be an integer between 100 and 599.
	StatusCode *int64 `json:"status_code,omitempty"`
	// A string representing an HTTP header name.
	Trigger *string `json:"trigger,omitempty"`
}

func (o *CreateRequestTerminationPluginConfig) GetBody() *string {
	if o == nil {
		return nil
	}
	return o.Body
}

func (o *CreateRequestTerminationPluginConfig) GetContentType() *string {
	if o == nil {
		return nil
	}
	return o.ContentType
}

func (o *CreateRequestTerminationPluginConfig) GetEcho() *bool {
	if o == nil {
		return nil
	}
	return o.Echo
}

func (o *CreateRequestTerminationPluginConfig) GetMessage() *string {
	if o == nil {
		return nil
	}
	return o.Message
}

func (o *CreateRequestTerminationPluginConfig) GetStatusCode() *int64 {
	if o == nil {
		return nil
	}
	return o.StatusCode
}

func (o *CreateRequestTerminationPluginConfig) GetTrigger() *string {
	if o == nil {
		return nil
	}
	return o.Trigger
}

type CreateRequestTerminationPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateRequestTerminationPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateRequestTerminationPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateRequestTerminationPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateRequestTerminationPluginOrdering struct {
	After  *CreateRequestTerminationPluginAfter  `json:"after,omitempty"`
	Before *CreateRequestTerminationPluginBefore `json:"before,omitempty"`
}

func (o *CreateRequestTerminationPluginOrdering) GetAfter() *CreateRequestTerminationPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *CreateRequestTerminationPluginOrdering) GetBefore() *CreateRequestTerminationPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type CreateRequestTerminationPluginProtocols string

const (
	CreateRequestTerminationPluginProtocolsGrpc           CreateRequestTerminationPluginProtocols = "grpc"
	CreateRequestTerminationPluginProtocolsGrpcs          CreateRequestTerminationPluginProtocols = "grpcs"
	CreateRequestTerminationPluginProtocolsHTTP           CreateRequestTerminationPluginProtocols = "http"
	CreateRequestTerminationPluginProtocolsHTTPS          CreateRequestTerminationPluginProtocols = "https"
	CreateRequestTerminationPluginProtocolsTCP            CreateRequestTerminationPluginProtocols = "tcp"
	CreateRequestTerminationPluginProtocolsTLS            CreateRequestTerminationPluginProtocols = "tls"
	CreateRequestTerminationPluginProtocolsTLSPassthrough CreateRequestTerminationPluginProtocols = "tls_passthrough"
	CreateRequestTerminationPluginProtocolsUDP            CreateRequestTerminationPluginProtocols = "udp"
	CreateRequestTerminationPluginProtocolsWs             CreateRequestTerminationPluginProtocols = "ws"
	CreateRequestTerminationPluginProtocolsWss            CreateRequestTerminationPluginProtocols = "wss"
)

func (e CreateRequestTerminationPluginProtocols) ToPointer() *CreateRequestTerminationPluginProtocols {
	return &e
}
func (e *CreateRequestTerminationPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateRequestTerminationPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateRequestTerminationPluginProtocols: %v", v)
	}
}

// CreateRequestTerminationPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateRequestTerminationPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestTerminationPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateRequestTerminationPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestTerminationPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateRequestTerminationPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateRequestTerminationPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestTerminationPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateRequestTerminationPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateRequestTerminationPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateRequestTerminationPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateRequestTerminationPlugin struct {
	Config *CreateRequestTerminationPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                                   `json:"enabled,omitempty"`
	InstanceName *string                                 `json:"instance_name,omitempty"`
	name         *string                                 `const:"request-termination" json:"name,omitempty"`
	Ordering     *CreateRequestTerminationPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateRequestTerminationPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateRequestTerminationPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateRequestTerminationPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateRequestTerminationPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateRequestTerminationPluginService `json:"service,omitempty"`
}

func (c CreateRequestTerminationPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateRequestTerminationPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateRequestTerminationPlugin) GetConfig() *CreateRequestTerminationPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateRequestTerminationPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateRequestTerminationPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateRequestTerminationPlugin) GetName() *string {
	return types.String("request-termination")
}

func (o *CreateRequestTerminationPlugin) GetOrdering() *CreateRequestTerminationPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateRequestTerminationPlugin) GetProtocols() []CreateRequestTerminationPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateRequestTerminationPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateRequestTerminationPlugin) GetConsumer() *CreateRequestTerminationPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateRequestTerminationPlugin) GetConsumerGroup() *CreateRequestTerminationPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateRequestTerminationPlugin) GetRoute() *CreateRequestTerminationPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateRequestTerminationPlugin) GetService() *CreateRequestTerminationPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
