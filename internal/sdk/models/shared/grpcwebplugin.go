// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type GrpcWebPluginConfig struct {
	// The value of the `Access-Control-Allow-Origin` header in the response to the gRPC-Web client.
	AllowOriginHeader *string `json:"allow_origin_header,omitempty"`
	// If set to `true` causes the plugin to pass the stripped request path to the upstream gRPC service.
	PassStrippedPath *bool `json:"pass_stripped_path,omitempty"`
	// If present, describes the gRPC types and methods. Required to support payload transcoding. When absent, the web client must use application/grpw-web+proto content.
	Proto *string `json:"proto,omitempty"`
}

func (o *GrpcWebPluginConfig) GetAllowOriginHeader() *string {
	if o == nil {
		return nil
	}
	return o.AllowOriginHeader
}

func (o *GrpcWebPluginConfig) GetPassStrippedPath() *bool {
	if o == nil {
		return nil
	}
	return o.PassStrippedPath
}

func (o *GrpcWebPluginConfig) GetProto() *string {
	if o == nil {
		return nil
	}
	return o.Proto
}

type GrpcWebPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *GrpcWebPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type GrpcWebPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *GrpcWebPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type GrpcWebPluginOrdering struct {
	After  *GrpcWebPluginAfter  `json:"after,omitempty"`
	Before *GrpcWebPluginBefore `json:"before,omitempty"`
}

func (o *GrpcWebPluginOrdering) GetAfter() *GrpcWebPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *GrpcWebPluginOrdering) GetBefore() *GrpcWebPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type GrpcWebPluginProtocols string

const (
	GrpcWebPluginProtocolsGrpc           GrpcWebPluginProtocols = "grpc"
	GrpcWebPluginProtocolsGrpcs          GrpcWebPluginProtocols = "grpcs"
	GrpcWebPluginProtocolsHTTP           GrpcWebPluginProtocols = "http"
	GrpcWebPluginProtocolsHTTPS          GrpcWebPluginProtocols = "https"
	GrpcWebPluginProtocolsTCP            GrpcWebPluginProtocols = "tcp"
	GrpcWebPluginProtocolsTLS            GrpcWebPluginProtocols = "tls"
	GrpcWebPluginProtocolsTLSPassthrough GrpcWebPluginProtocols = "tls_passthrough"
	GrpcWebPluginProtocolsUDP            GrpcWebPluginProtocols = "udp"
	GrpcWebPluginProtocolsWs             GrpcWebPluginProtocols = "ws"
	GrpcWebPluginProtocolsWss            GrpcWebPluginProtocols = "wss"
)

func (e GrpcWebPluginProtocols) ToPointer() *GrpcWebPluginProtocols {
	return &e
}
func (e *GrpcWebPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = GrpcWebPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for GrpcWebPluginProtocols: %v", v)
	}
}

// GrpcWebPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type GrpcWebPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *GrpcWebPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type GrpcWebPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *GrpcWebPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// GrpcWebPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type GrpcWebPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *GrpcWebPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// GrpcWebPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type GrpcWebPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *GrpcWebPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type GrpcWebPlugin struct {
	Config *GrpcWebPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                  `json:"enabled,omitempty"`
	ID           *string                `json:"id,omitempty"`
	InstanceName *string                `json:"instance_name,omitempty"`
	name         *string                `const:"grpc-web" json:"name,omitempty"`
	Ordering     *GrpcWebPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []GrpcWebPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *GrpcWebPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *GrpcWebPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *GrpcWebPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *GrpcWebPluginService `json:"service,omitempty"`
}

func (g GrpcWebPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GrpcWebPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GrpcWebPlugin) GetConfig() *GrpcWebPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *GrpcWebPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *GrpcWebPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *GrpcWebPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GrpcWebPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *GrpcWebPlugin) GetName() *string {
	return types.String("grpc-web")
}

func (o *GrpcWebPlugin) GetOrdering() *GrpcWebPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *GrpcWebPlugin) GetProtocols() []GrpcWebPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *GrpcWebPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *GrpcWebPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *GrpcWebPlugin) GetConsumer() *GrpcWebPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *GrpcWebPlugin) GetConsumerGroup() *GrpcWebPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *GrpcWebPlugin) GetRoute() *GrpcWebPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *GrpcWebPlugin) GetService() *GrpcWebPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
