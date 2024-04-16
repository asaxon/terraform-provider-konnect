// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
)

type CreateProxyCachePluginProtocols string

const (
	CreateProxyCachePluginProtocolsGrpc           CreateProxyCachePluginProtocols = "grpc"
	CreateProxyCachePluginProtocolsGrpcs          CreateProxyCachePluginProtocols = "grpcs"
	CreateProxyCachePluginProtocolsHTTP           CreateProxyCachePluginProtocols = "http"
	CreateProxyCachePluginProtocolsHTTPS          CreateProxyCachePluginProtocols = "https"
	CreateProxyCachePluginProtocolsTCP            CreateProxyCachePluginProtocols = "tcp"
	CreateProxyCachePluginProtocolsTLS            CreateProxyCachePluginProtocols = "tls"
	CreateProxyCachePluginProtocolsTLSPassthrough CreateProxyCachePluginProtocols = "tls_passthrough"
	CreateProxyCachePluginProtocolsUDP            CreateProxyCachePluginProtocols = "udp"
	CreateProxyCachePluginProtocolsWs             CreateProxyCachePluginProtocols = "ws"
	CreateProxyCachePluginProtocolsWss            CreateProxyCachePluginProtocols = "wss"
)

func (e CreateProxyCachePluginProtocols) ToPointer() *CreateProxyCachePluginProtocols {
	return &e
}

func (e *CreateProxyCachePluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateProxyCachePluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateProxyCachePluginProtocols: %v", v)
	}
}

// CreateProxyCachePluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateProxyCachePluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateProxyCachePluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateProxyCachePluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateProxyCachePluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateProxyCachePluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateProxyCachePluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateProxyCachePluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateProxyCachePluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type RequestMethod string

const (
	RequestMethodHead  RequestMethod = "HEAD"
	RequestMethodGet   RequestMethod = "GET"
	RequestMethodPost  RequestMethod = "POST"
	RequestMethodPatch RequestMethod = "PATCH"
	RequestMethodPut   RequestMethod = "PUT"
)

func (e RequestMethod) ToPointer() *RequestMethod {
	return &e
}

func (e *RequestMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "HEAD":
		fallthrough
	case "GET":
		fallthrough
	case "POST":
		fallthrough
	case "PATCH":
		fallthrough
	case "PUT":
		*e = RequestMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RequestMethod: %v", v)
	}
}

// CreateProxyCachePluginStrategy - The backing data store in which to hold cache entities.
type CreateProxyCachePluginStrategy string

const (
	CreateProxyCachePluginStrategyMemory CreateProxyCachePluginStrategy = "memory"
)

func (e CreateProxyCachePluginStrategy) ToPointer() *CreateProxyCachePluginStrategy {
	return &e
}

func (e *CreateProxyCachePluginStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "memory":
		*e = CreateProxyCachePluginStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateProxyCachePluginStrategy: %v", v)
	}
}

type Memory struct {
	// The name of the shared dictionary in which to hold cache entities when the memory strategy is selected. Note that this dictionary currently must be defined manually in the Kong Nginx template.
	DictionaryName *string `default:"kong_db_cache" json:"dictionary_name"`
}

func (m Memory) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(m, "", false)
}

func (m *Memory) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &m, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *Memory) GetDictionaryName() *string {
	if o == nil {
		return nil
	}
	return o.DictionaryName
}

// ResponseHeaders - Caching related diagnostic headers that should be included in cached responses
type ResponseHeaders struct {
	Age          *bool `default:"true" json:"age"`
	XCacheStatus *bool `default:"true" json:"X-Cache-Status"`
	XCacheKey    *bool `default:"true" json:"X-Cache-Key"`
}

func (r ResponseHeaders) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *ResponseHeaders) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *ResponseHeaders) GetAge() *bool {
	if o == nil {
		return nil
	}
	return o.Age
}

func (o *ResponseHeaders) GetXCacheStatus() *bool {
	if o == nil {
		return nil
	}
	return o.XCacheStatus
}

func (o *ResponseHeaders) GetXCacheKey() *bool {
	if o == nil {
		return nil
	}
	return o.XCacheKey
}

type CreateProxyCachePluginConfig struct {
	// Upstream response status code considered cacheable.
	ResponseCode []int64 `json:"response_code,omitempty"`
	// Downstream request methods considered cacheable.
	RequestMethod []RequestMethod `json:"request_method,omitempty"`
	// Upstream response content types considered cacheable. The plugin performs an **exact match** against each specified value.
	ContentType []string `json:"content_type,omitempty"`
	// TTL, in seconds, of cache entities.
	CacheTTL *int64 `default:"300" json:"cache_ttl"`
	// The backing data store in which to hold cache entities.
	Strategy *CreateProxyCachePluginStrategy `json:"strategy,omitempty"`
	// When enabled, respect the Cache-Control behaviors defined in RFC7234.
	CacheControl  *bool `default:"false" json:"cache_control"`
	IgnoreURICase *bool `default:"false" json:"ignore_uri_case"`
	// Number of seconds to keep resources in the storage backend. This value is independent of `cache_ttl` or resource TTLs defined by Cache-Control behaviors.
	StorageTTL *int64  `json:"storage_ttl,omitempty"`
	Memory     *Memory `json:"memory,omitempty"`
	// Relevant query parameters considered for the cache key. If undefined, all params are taken into consideration.
	VaryQueryParams []string `json:"vary_query_params,omitempty"`
	// Relevant headers considered for the cache key. If undefined, none of the headers are taken into consideration.
	VaryHeaders []string `json:"vary_headers,omitempty"`
	// Caching related diagnostic headers that should be included in cached responses
	ResponseHeaders *ResponseHeaders `json:"response_headers,omitempty"`
}

func (c CreateProxyCachePluginConfig) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateProxyCachePluginConfig) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateProxyCachePluginConfig) GetResponseCode() []int64 {
	if o == nil {
		return nil
	}
	return o.ResponseCode
}

func (o *CreateProxyCachePluginConfig) GetRequestMethod() []RequestMethod {
	if o == nil {
		return nil
	}
	return o.RequestMethod
}

func (o *CreateProxyCachePluginConfig) GetContentType() []string {
	if o == nil {
		return nil
	}
	return o.ContentType
}

func (o *CreateProxyCachePluginConfig) GetCacheTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.CacheTTL
}

func (o *CreateProxyCachePluginConfig) GetStrategy() *CreateProxyCachePluginStrategy {
	if o == nil {
		return nil
	}
	return o.Strategy
}

func (o *CreateProxyCachePluginConfig) GetCacheControl() *bool {
	if o == nil {
		return nil
	}
	return o.CacheControl
}

func (o *CreateProxyCachePluginConfig) GetIgnoreURICase() *bool {
	if o == nil {
		return nil
	}
	return o.IgnoreURICase
}

func (o *CreateProxyCachePluginConfig) GetStorageTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.StorageTTL
}

func (o *CreateProxyCachePluginConfig) GetMemory() *Memory {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *CreateProxyCachePluginConfig) GetVaryQueryParams() []string {
	if o == nil {
		return nil
	}
	return o.VaryQueryParams
}

func (o *CreateProxyCachePluginConfig) GetVaryHeaders() []string {
	if o == nil {
		return nil
	}
	return o.VaryHeaders
}

func (o *CreateProxyCachePluginConfig) GetResponseHeaders() *ResponseHeaders {
	if o == nil {
		return nil
	}
	return o.ResponseHeaders
}

// CreateProxyCachePlugin - A Plugin entity represents a plugin configuration that will be executed during the HTTP request/response lifecycle. It is how you can add functionalities to Services that run behind Kong, like Authentication or Rate Limiting for example. You can find more information about how to install and what values each plugin takes by visiting the [Kong Hub](https://docs.konghq.com/hub/). When adding a Plugin Configuration to a Service, every request made by a client to that Service will run said Plugin. If a Plugin needs to be tuned to different values for some specific Consumers, you can do so by creating a separate plugin instance that specifies both the Service and the Consumer, through the `service` and `consumer` fields.
type CreateProxyCachePlugin struct {
	// Whether the plugin is applied.
	Enabled *bool  `default:"true" json:"enabled"`
	name    string `const:"proxy-cache" json:"name"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateProxyCachePluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer *CreateProxyCachePluginConsumer `json:"consumer,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateProxyCachePluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateProxyCachePluginService `json:"service,omitempty"`
	Config  CreateProxyCachePluginConfig   `json:"config"`
}

func (c CreateProxyCachePlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateProxyCachePlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateProxyCachePlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateProxyCachePlugin) GetName() string {
	return "proxy-cache"
}

func (o *CreateProxyCachePlugin) GetProtocols() []CreateProxyCachePluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateProxyCachePlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateProxyCachePlugin) GetConsumer() *CreateProxyCachePluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateProxyCachePlugin) GetRoute() *CreateProxyCachePluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateProxyCachePlugin) GetService() *CreateProxyCachePluginService {
	if o == nil {
		return nil
	}
	return o.Service
}

func (o *CreateProxyCachePlugin) GetConfig() CreateProxyCachePluginConfig {
	if o == nil {
		return CreateProxyCachePluginConfig{}
	}
	return o.Config
}