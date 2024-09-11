// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

// LimitBy - The entity that is used when aggregating the limits.
type LimitBy string

const (
	LimitByConsumer      LimitBy = "consumer"
	LimitByCredential    LimitBy = "credential"
	LimitByIP            LimitBy = "ip"
	LimitByService       LimitBy = "service"
	LimitByHeader        LimitBy = "header"
	LimitByPath          LimitBy = "path"
	LimitByConsumerGroup LimitBy = "consumer-group"
)

func (e LimitBy) ToPointer() *LimitBy {
	return &e
}
func (e *LimitBy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consumer":
		fallthrough
	case "credential":
		fallthrough
	case "ip":
		fallthrough
	case "service":
		fallthrough
	case "header":
		fallthrough
	case "path":
		fallthrough
	case "consumer-group":
		*e = LimitBy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for LimitBy: %v", v)
	}
}

// Policy - The rate-limiting policies to use for retrieving and incrementing the limits.
type Policy string

const (
	PolicyLocal   Policy = "local"
	PolicyCluster Policy = "cluster"
	PolicyRedis   Policy = "redis"
)

func (e Policy) ToPointer() *Policy {
	return &e
}
func (e *Policy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "local":
		fallthrough
	case "cluster":
		fallthrough
	case "redis":
		*e = Policy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Policy: %v", v)
	}
}

// RateLimitingPluginRedis - Redis configuration
type RateLimitingPluginRedis struct {
	// Database to use for the Redis connection when using the `redis` strategy
	Database *int64 `json:"database,omitempty"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.
	Password *string `json:"password,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// A string representing an SNI (server name indication) value for TLS.
	ServerName *string `json:"server_name,omitempty"`
	// If set to true, uses SSL to connect to Redis.
	Ssl *bool `json:"ssl,omitempty"`
	// If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure `lua_ssl_trusted_certificate` in `kong.conf` to specify the CA (or server) certificate used by your Redis server. You may also need to configure `lua_ssl_verify_depth` accordingly.
	SslVerify *bool `json:"ssl_verify,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	Timeout *int64 `json:"timeout,omitempty"`
	// Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to `default`.
	Username *string `json:"username,omitempty"`
}

func (o *RateLimitingPluginRedis) GetDatabase() *int64 {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *RateLimitingPluginRedis) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *RateLimitingPluginRedis) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *RateLimitingPluginRedis) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *RateLimitingPluginRedis) GetServerName() *string {
	if o == nil {
		return nil
	}
	return o.ServerName
}

func (o *RateLimitingPluginRedis) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *RateLimitingPluginRedis) GetSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SslVerify
}

func (o *RateLimitingPluginRedis) GetTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *RateLimitingPluginRedis) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

type RateLimitingPluginConfig struct {
	// The number of HTTP requests that can be made per day.
	Day *float64 `json:"day,omitempty"`
	// Set a custom error code to return when the rate limit is exceeded.
	ErrorCode *float64 `json:"error_code,omitempty"`
	// Set a custom error message to return when the rate limit is exceeded.
	ErrorMessage *string `json:"error_message,omitempty"`
	// A boolean value that determines if the requests should be proxied even if Kong has troubles connecting a third-party data store. If `true`, requests will be proxied anyway, effectively disabling the rate-limiting function until the data store is working again. If `false`, then the clients will see `500` errors.
	FaultTolerant *bool `json:"fault_tolerant,omitempty"`
	// A string representing an HTTP header name.
	HeaderName *string `json:"header_name,omitempty"`
	// Optionally hide informative response headers.
	HideClientHeaders *bool `json:"hide_client_headers,omitempty"`
	// The number of HTTP requests that can be made per hour.
	Hour *float64 `json:"hour,omitempty"`
	// The entity that is used when aggregating the limits.
	LimitBy *LimitBy `json:"limit_by,omitempty"`
	// The number of HTTP requests that can be made per minute.
	Minute *float64 `json:"minute,omitempty"`
	// The number of HTTP requests that can be made per month.
	Month *float64 `json:"month,omitempty"`
	// A string representing a URL path, such as /path/to/resource. Must start with a forward slash (/) and must not contain empty segments (i.e., two consecutive forward slashes).
	Path *string `json:"path,omitempty"`
	// The rate-limiting policies to use for retrieving and incrementing the limits.
	Policy *Policy `json:"policy,omitempty"`
	// Redis configuration
	Redis *RateLimitingPluginRedis `json:"redis,omitempty"`
	// The number of HTTP requests that can be made per second.
	Second *float64 `json:"second,omitempty"`
	// How often to sync counter data to the central data store. A value of -1 results in synchronous behavior.
	SyncRate *float64 `json:"sync_rate,omitempty"`
	// The number of HTTP requests that can be made per year.
	Year *float64 `json:"year,omitempty"`
}

func (o *RateLimitingPluginConfig) GetDay() *float64 {
	if o == nil {
		return nil
	}
	return o.Day
}

func (o *RateLimitingPluginConfig) GetErrorCode() *float64 {
	if o == nil {
		return nil
	}
	return o.ErrorCode
}

func (o *RateLimitingPluginConfig) GetErrorMessage() *string {
	if o == nil {
		return nil
	}
	return o.ErrorMessage
}

func (o *RateLimitingPluginConfig) GetFaultTolerant() *bool {
	if o == nil {
		return nil
	}
	return o.FaultTolerant
}

func (o *RateLimitingPluginConfig) GetHeaderName() *string {
	if o == nil {
		return nil
	}
	return o.HeaderName
}

func (o *RateLimitingPluginConfig) GetHideClientHeaders() *bool {
	if o == nil {
		return nil
	}
	return o.HideClientHeaders
}

func (o *RateLimitingPluginConfig) GetHour() *float64 {
	if o == nil {
		return nil
	}
	return o.Hour
}

func (o *RateLimitingPluginConfig) GetLimitBy() *LimitBy {
	if o == nil {
		return nil
	}
	return o.LimitBy
}

func (o *RateLimitingPluginConfig) GetMinute() *float64 {
	if o == nil {
		return nil
	}
	return o.Minute
}

func (o *RateLimitingPluginConfig) GetMonth() *float64 {
	if o == nil {
		return nil
	}
	return o.Month
}

func (o *RateLimitingPluginConfig) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *RateLimitingPluginConfig) GetPolicy() *Policy {
	if o == nil {
		return nil
	}
	return o.Policy
}

func (o *RateLimitingPluginConfig) GetRedis() *RateLimitingPluginRedis {
	if o == nil {
		return nil
	}
	return o.Redis
}

func (o *RateLimitingPluginConfig) GetSecond() *float64 {
	if o == nil {
		return nil
	}
	return o.Second
}

func (o *RateLimitingPluginConfig) GetSyncRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SyncRate
}

func (o *RateLimitingPluginConfig) GetYear() *float64 {
	if o == nil {
		return nil
	}
	return o.Year
}

type RateLimitingPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *RateLimitingPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type RateLimitingPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *RateLimitingPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type RateLimitingPluginOrdering struct {
	After  *RateLimitingPluginAfter  `json:"after,omitempty"`
	Before *RateLimitingPluginBefore `json:"before,omitempty"`
}

func (o *RateLimitingPluginOrdering) GetAfter() *RateLimitingPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *RateLimitingPluginOrdering) GetBefore() *RateLimitingPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type RateLimitingPluginProtocols string

const (
	RateLimitingPluginProtocolsGrpc           RateLimitingPluginProtocols = "grpc"
	RateLimitingPluginProtocolsGrpcs          RateLimitingPluginProtocols = "grpcs"
	RateLimitingPluginProtocolsHTTP           RateLimitingPluginProtocols = "http"
	RateLimitingPluginProtocolsHTTPS          RateLimitingPluginProtocols = "https"
	RateLimitingPluginProtocolsTCP            RateLimitingPluginProtocols = "tcp"
	RateLimitingPluginProtocolsTLS            RateLimitingPluginProtocols = "tls"
	RateLimitingPluginProtocolsTLSPassthrough RateLimitingPluginProtocols = "tls_passthrough"
	RateLimitingPluginProtocolsUDP            RateLimitingPluginProtocols = "udp"
	RateLimitingPluginProtocolsWs             RateLimitingPluginProtocols = "ws"
	RateLimitingPluginProtocolsWss            RateLimitingPluginProtocols = "wss"
)

func (e RateLimitingPluginProtocols) ToPointer() *RateLimitingPluginProtocols {
	return &e
}
func (e *RateLimitingPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = RateLimitingPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for RateLimitingPluginProtocols: %v", v)
	}
}

// RateLimitingPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type RateLimitingPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *RateLimitingPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type RateLimitingPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *RateLimitingPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RateLimitingPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type RateLimitingPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *RateLimitingPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// RateLimitingPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type RateLimitingPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *RateLimitingPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type RateLimitingPlugin struct {
	Config *RateLimitingPluginConfig `json:"config,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64 `json:"created_at,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                       `json:"enabled,omitempty"`
	ID           *string                     `json:"id,omitempty"`
	InstanceName *string                     `json:"instance_name,omitempty"`
	name         *string                     `const:"rate-limiting" json:"name,omitempty"`
	Ordering     *RateLimitingPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []RateLimitingPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// Unix epoch when the resource was last updated.
	UpdatedAt *int64 `json:"updated_at,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *RateLimitingPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *RateLimitingPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *RateLimitingPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *RateLimitingPluginService `json:"service,omitempty"`
}

func (r RateLimitingPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(r, "", false)
}

func (r *RateLimitingPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &r, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *RateLimitingPlugin) GetConfig() *RateLimitingPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *RateLimitingPlugin) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *RateLimitingPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *RateLimitingPlugin) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *RateLimitingPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *RateLimitingPlugin) GetName() *string {
	return types.String("rate-limiting")
}

func (o *RateLimitingPlugin) GetOrdering() *RateLimitingPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *RateLimitingPlugin) GetProtocols() []RateLimitingPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *RateLimitingPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *RateLimitingPlugin) GetUpdatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.UpdatedAt
}

func (o *RateLimitingPlugin) GetConsumer() *RateLimitingPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *RateLimitingPlugin) GetConsumerGroup() *RateLimitingPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *RateLimitingPlugin) GetRoute() *RateLimitingPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *RateLimitingPlugin) GetService() *RateLimitingPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
