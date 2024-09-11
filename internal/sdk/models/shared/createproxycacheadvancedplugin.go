// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

type CreateProxyCacheAdvancedPluginMemory struct {
	// The name of the shared dictionary in which to hold cache entities when the memory strategy is selected. Note that this dictionary currently must be defined manually in the Kong Nginx template.
	DictionaryName *string `json:"dictionary_name,omitempty"`
}

func (o *CreateProxyCacheAdvancedPluginMemory) GetDictionaryName() *string {
	if o == nil {
		return nil
	}
	return o.DictionaryName
}

type CreateProxyCacheAdvancedPluginClusterNodes struct {
	// A string representing a host name, such as example.com.
	IP *string `json:"ip,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
}

func (o *CreateProxyCacheAdvancedPluginClusterNodes) GetIP() *string {
	if o == nil {
		return nil
	}
	return o.IP
}

func (o *CreateProxyCacheAdvancedPluginClusterNodes) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

type CreateProxyCacheAdvancedPluginSentinelNodes struct {
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
}

func (o *CreateProxyCacheAdvancedPluginSentinelNodes) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *CreateProxyCacheAdvancedPluginSentinelNodes) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

// CreateProxyCacheAdvancedPluginSentinelRole - Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.
type CreateProxyCacheAdvancedPluginSentinelRole string

const (
	CreateProxyCacheAdvancedPluginSentinelRoleMaster CreateProxyCacheAdvancedPluginSentinelRole = "master"
	CreateProxyCacheAdvancedPluginSentinelRoleSlave  CreateProxyCacheAdvancedPluginSentinelRole = "slave"
	CreateProxyCacheAdvancedPluginSentinelRoleAny    CreateProxyCacheAdvancedPluginSentinelRole = "any"
)

func (e CreateProxyCacheAdvancedPluginSentinelRole) ToPointer() *CreateProxyCacheAdvancedPluginSentinelRole {
	return &e
}
func (e *CreateProxyCacheAdvancedPluginSentinelRole) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "master":
		fallthrough
	case "slave":
		fallthrough
	case "any":
		*e = CreateProxyCacheAdvancedPluginSentinelRole(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateProxyCacheAdvancedPluginSentinelRole: %v", v)
	}
}

type CreateProxyCacheAdvancedPluginRedis struct {
	// Maximum retry attempts for redirection.
	ClusterMaxRedirections *int64 `json:"cluster_max_redirections,omitempty"`
	// Cluster addresses to use for Redis connections when the `redis` strategy is defined. Defining this field implies using a Redis Cluster. The minimum length of the array is 1 element.
	ClusterNodes []CreateProxyCacheAdvancedPluginClusterNodes `json:"cluster_nodes,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ConnectTimeout *int64 `json:"connect_timeout,omitempty"`
	// If the connection to Redis is proxied (e.g. Envoy), set it `true`. Set the `host` and `port` to point to the proxy address.
	ConnectionIsProxied *bool `json:"connection_is_proxied,omitempty"`
	// Database to use for the Redis connection when using the `redis` strategy
	Database *int64 `json:"database,omitempty"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// Limits the total number of opened connections for a pool. If the connection pool is full, connection queues above the limit go into the backlog queue. If the backlog queue is full, subsequent connect operations fail and return `nil`. Queued operations (subject to set timeouts) resume once the number of connections in the pool is less than `keepalive_pool_size`. If latency is high or throughput is low, try increasing this value. Empirically, this value is larger than `keepalive_pool_size`.
	KeepaliveBacklog *int64 `json:"keepalive_backlog,omitempty"`
	// The size limit for every cosocket connection pool associated with every remote server, per worker process. If neither `keepalive_pool_size` nor `keepalive_backlog` is specified, no pool is created. If `keepalive_pool_size` isn't specified but `keepalive_backlog` is specified, then the pool uses the default value. Try to increase (e.g. 512) this value if latency is high or throughput is low.
	KeepalivePoolSize *int64 `json:"keepalive_pool_size,omitempty"`
	// Password to use for Redis connections. If undefined, no AUTH commands are sent to Redis.
	Password *string `json:"password,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	ReadTimeout *int64 `json:"read_timeout,omitempty"`
	// An integer representing a timeout in milliseconds. Must be between 0 and 2^31-2.
	SendTimeout *int64 `json:"send_timeout,omitempty"`
	// Sentinel master to use for Redis connections. Defining this value implies using Redis Sentinel.
	SentinelMaster *string `json:"sentinel_master,omitempty"`
	// Sentinel node addresses to use for Redis connections when the `redis` strategy is defined. Defining this field implies using a Redis Sentinel. The minimum length of the array is 1 element.
	SentinelNodes []CreateProxyCacheAdvancedPluginSentinelNodes `json:"sentinel_nodes,omitempty"`
	// Sentinel password to authenticate with a Redis Sentinel instance. If undefined, no AUTH commands are sent to Redis Sentinels.
	SentinelPassword *string `json:"sentinel_password,omitempty"`
	// Sentinel role to use for Redis connections when the `redis` strategy is defined. Defining this value implies using Redis Sentinel.
	SentinelRole *CreateProxyCacheAdvancedPluginSentinelRole `json:"sentinel_role,omitempty"`
	// Sentinel username to authenticate with a Redis Sentinel instance. If undefined, ACL authentication won't be performed. This requires Redis v6.2.0+.
	SentinelUsername *string `json:"sentinel_username,omitempty"`
	// A string representing an SNI (server name indication) value for TLS.
	ServerName *string `json:"server_name,omitempty"`
	// If set to true, uses SSL to connect to Redis.
	Ssl *bool `json:"ssl,omitempty"`
	// If set to true, verifies the validity of the server SSL certificate. If setting this parameter, also configure `lua_ssl_trusted_certificate` in `kong.conf` to specify the CA (or server) certificate used by your Redis server. You may also need to configure `lua_ssl_verify_depth` accordingly.
	SslVerify *bool `json:"ssl_verify,omitempty"`
	// Username to use for Redis connections. If undefined, ACL authentication won't be performed. This requires Redis v6.0.0+. To be compatible with Redis v5.x.y, you can set it to `default`.
	Username *string `json:"username,omitempty"`
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetClusterMaxRedirections() *int64 {
	if o == nil {
		return nil
	}
	return o.ClusterMaxRedirections
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetClusterNodes() []CreateProxyCacheAdvancedPluginClusterNodes {
	if o == nil {
		return nil
	}
	return o.ClusterNodes
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetConnectTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ConnectTimeout
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetConnectionIsProxied() *bool {
	if o == nil {
		return nil
	}
	return o.ConnectionIsProxied
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetDatabase() *int64 {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetKeepaliveBacklog() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepaliveBacklog
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetKeepalivePoolSize() *int64 {
	if o == nil {
		return nil
	}
	return o.KeepalivePoolSize
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetReadTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.ReadTimeout
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetSendTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.SendTimeout
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetSentinelMaster() *string {
	if o == nil {
		return nil
	}
	return o.SentinelMaster
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetSentinelNodes() []CreateProxyCacheAdvancedPluginSentinelNodes {
	if o == nil {
		return nil
	}
	return o.SentinelNodes
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetSentinelPassword() *string {
	if o == nil {
		return nil
	}
	return o.SentinelPassword
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetSentinelRole() *CreateProxyCacheAdvancedPluginSentinelRole {
	if o == nil {
		return nil
	}
	return o.SentinelRole
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetSentinelUsername() *string {
	if o == nil {
		return nil
	}
	return o.SentinelUsername
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetServerName() *string {
	if o == nil {
		return nil
	}
	return o.ServerName
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SslVerify
}

func (o *CreateProxyCacheAdvancedPluginRedis) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

type CreateProxyCacheAdvancedPluginRequestMethod string

const (
	CreateProxyCacheAdvancedPluginRequestMethodHead  CreateProxyCacheAdvancedPluginRequestMethod = "HEAD"
	CreateProxyCacheAdvancedPluginRequestMethodGet   CreateProxyCacheAdvancedPluginRequestMethod = "GET"
	CreateProxyCacheAdvancedPluginRequestMethodPost  CreateProxyCacheAdvancedPluginRequestMethod = "POST"
	CreateProxyCacheAdvancedPluginRequestMethodPatch CreateProxyCacheAdvancedPluginRequestMethod = "PATCH"
	CreateProxyCacheAdvancedPluginRequestMethodPut   CreateProxyCacheAdvancedPluginRequestMethod = "PUT"
)

func (e CreateProxyCacheAdvancedPluginRequestMethod) ToPointer() *CreateProxyCacheAdvancedPluginRequestMethod {
	return &e
}
func (e *CreateProxyCacheAdvancedPluginRequestMethod) UnmarshalJSON(data []byte) error {
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
		*e = CreateProxyCacheAdvancedPluginRequestMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateProxyCacheAdvancedPluginRequestMethod: %v", v)
	}
}

// CreateProxyCacheAdvancedPluginResponseHeaders - Caching related diagnostic headers that should be included in cached responses
type CreateProxyCacheAdvancedPluginResponseHeaders struct {
	XCacheKey    *bool `json:"X-Cache-Key,omitempty"`
	XCacheStatus *bool `json:"X-Cache-Status,omitempty"`
	Age          *bool `json:"age,omitempty"`
}

func (o *CreateProxyCacheAdvancedPluginResponseHeaders) GetXCacheKey() *bool {
	if o == nil {
		return nil
	}
	return o.XCacheKey
}

func (o *CreateProxyCacheAdvancedPluginResponseHeaders) GetXCacheStatus() *bool {
	if o == nil {
		return nil
	}
	return o.XCacheStatus
}

func (o *CreateProxyCacheAdvancedPluginResponseHeaders) GetAge() *bool {
	if o == nil {
		return nil
	}
	return o.Age
}

// CreateProxyCacheAdvancedPluginStrategy - The backing data store in which to hold cache entities. Accepted values are: `memory` and `redis`.
type CreateProxyCacheAdvancedPluginStrategy string

const (
	CreateProxyCacheAdvancedPluginStrategyMemory CreateProxyCacheAdvancedPluginStrategy = "memory"
	CreateProxyCacheAdvancedPluginStrategyRedis  CreateProxyCacheAdvancedPluginStrategy = "redis"
)

func (e CreateProxyCacheAdvancedPluginStrategy) ToPointer() *CreateProxyCacheAdvancedPluginStrategy {
	return &e
}
func (e *CreateProxyCacheAdvancedPluginStrategy) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "memory":
		fallthrough
	case "redis":
		*e = CreateProxyCacheAdvancedPluginStrategy(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateProxyCacheAdvancedPluginStrategy: %v", v)
	}
}

type CreateProxyCacheAdvancedPluginConfig struct {
	// Unhandled errors while trying to retrieve a cache entry (such as redis down) are resolved with `Bypass`, with the request going upstream.
	BypassOnErr *bool `json:"bypass_on_err,omitempty"`
	// When enabled, respect the Cache-Control behaviors defined in RFC7234.
	CacheControl *bool `json:"cache_control,omitempty"`
	// TTL in seconds of cache entities.
	CacheTTL *int64 `json:"cache_ttl,omitempty"`
	// Upstream response content types considered cacheable. The plugin performs an **exact match** against each specified value; for example, if the upstream is expected to respond with a `application/json; charset=utf-8` content-type, the plugin configuration must contain said value or a `Bypass` cache status is returned.
	ContentType []string `json:"content_type,omitempty"`
	// Determines whether to treat URIs as case sensitive. By default, case sensitivity is enabled. If set to true, requests are cached while ignoring case sensitivity in the URI.
	IgnoreURICase *bool                                 `json:"ignore_uri_case,omitempty"`
	Memory        *CreateProxyCacheAdvancedPluginMemory `json:"memory,omitempty"`
	Redis         *CreateProxyCacheAdvancedPluginRedis  `json:"redis,omitempty"`
	// Downstream request methods considered cacheable. Available options: `HEAD`, `GET`, `POST`, `PATCH`, `PUT`.
	RequestMethod []CreateProxyCacheAdvancedPluginRequestMethod `json:"request_method,omitempty"`
	// Upstream response status code considered cacheable. The integers must be a value between 100 and 900.
	ResponseCode []int64 `json:"response_code,omitempty"`
	// Caching related diagnostic headers that should be included in cached responses
	ResponseHeaders *CreateProxyCacheAdvancedPluginResponseHeaders `json:"response_headers,omitempty"`
	// Number of seconds to keep resources in the storage backend. This value is independent of `cache_ttl` or resource TTLs defined by Cache-Control behaviors.
	StorageTTL *int64 `json:"storage_ttl,omitempty"`
	// The backing data store in which to hold cache entities. Accepted values are: `memory` and `redis`.
	Strategy *CreateProxyCacheAdvancedPluginStrategy `json:"strategy,omitempty"`
	// Relevant headers considered for the cache key. If undefined, none of the headers are taken into consideration.
	VaryHeaders []string `json:"vary_headers,omitempty"`
	// Relevant query parameters considered for the cache key. If undefined, all params are taken into consideration.
	VaryQueryParams []string `json:"vary_query_params,omitempty"`
}

func (o *CreateProxyCacheAdvancedPluginConfig) GetBypassOnErr() *bool {
	if o == nil {
		return nil
	}
	return o.BypassOnErr
}

func (o *CreateProxyCacheAdvancedPluginConfig) GetCacheControl() *bool {
	if o == nil {
		return nil
	}
	return o.CacheControl
}

func (o *CreateProxyCacheAdvancedPluginConfig) GetCacheTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.CacheTTL
}

func (o *CreateProxyCacheAdvancedPluginConfig) GetContentType() []string {
	if o == nil {
		return nil
	}
	return o.ContentType
}

func (o *CreateProxyCacheAdvancedPluginConfig) GetIgnoreURICase() *bool {
	if o == nil {
		return nil
	}
	return o.IgnoreURICase
}

func (o *CreateProxyCacheAdvancedPluginConfig) GetMemory() *CreateProxyCacheAdvancedPluginMemory {
	if o == nil {
		return nil
	}
	return o.Memory
}

func (o *CreateProxyCacheAdvancedPluginConfig) GetRedis() *CreateProxyCacheAdvancedPluginRedis {
	if o == nil {
		return nil
	}
	return o.Redis
}

func (o *CreateProxyCacheAdvancedPluginConfig) GetRequestMethod() []CreateProxyCacheAdvancedPluginRequestMethod {
	if o == nil {
		return nil
	}
	return o.RequestMethod
}

func (o *CreateProxyCacheAdvancedPluginConfig) GetResponseCode() []int64 {
	if o == nil {
		return nil
	}
	return o.ResponseCode
}

func (o *CreateProxyCacheAdvancedPluginConfig) GetResponseHeaders() *CreateProxyCacheAdvancedPluginResponseHeaders {
	if o == nil {
		return nil
	}
	return o.ResponseHeaders
}

func (o *CreateProxyCacheAdvancedPluginConfig) GetStorageTTL() *int64 {
	if o == nil {
		return nil
	}
	return o.StorageTTL
}

func (o *CreateProxyCacheAdvancedPluginConfig) GetStrategy() *CreateProxyCacheAdvancedPluginStrategy {
	if o == nil {
		return nil
	}
	return o.Strategy
}

func (o *CreateProxyCacheAdvancedPluginConfig) GetVaryHeaders() []string {
	if o == nil {
		return nil
	}
	return o.VaryHeaders
}

func (o *CreateProxyCacheAdvancedPluginConfig) GetVaryQueryParams() []string {
	if o == nil {
		return nil
	}
	return o.VaryQueryParams
}

type CreateProxyCacheAdvancedPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateProxyCacheAdvancedPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateProxyCacheAdvancedPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateProxyCacheAdvancedPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateProxyCacheAdvancedPluginOrdering struct {
	After  *CreateProxyCacheAdvancedPluginAfter  `json:"after,omitempty"`
	Before *CreateProxyCacheAdvancedPluginBefore `json:"before,omitempty"`
}

func (o *CreateProxyCacheAdvancedPluginOrdering) GetAfter() *CreateProxyCacheAdvancedPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *CreateProxyCacheAdvancedPluginOrdering) GetBefore() *CreateProxyCacheAdvancedPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type CreateProxyCacheAdvancedPluginProtocols string

const (
	CreateProxyCacheAdvancedPluginProtocolsGrpc           CreateProxyCacheAdvancedPluginProtocols = "grpc"
	CreateProxyCacheAdvancedPluginProtocolsGrpcs          CreateProxyCacheAdvancedPluginProtocols = "grpcs"
	CreateProxyCacheAdvancedPluginProtocolsHTTP           CreateProxyCacheAdvancedPluginProtocols = "http"
	CreateProxyCacheAdvancedPluginProtocolsHTTPS          CreateProxyCacheAdvancedPluginProtocols = "https"
	CreateProxyCacheAdvancedPluginProtocolsTCP            CreateProxyCacheAdvancedPluginProtocols = "tcp"
	CreateProxyCacheAdvancedPluginProtocolsTLS            CreateProxyCacheAdvancedPluginProtocols = "tls"
	CreateProxyCacheAdvancedPluginProtocolsTLSPassthrough CreateProxyCacheAdvancedPluginProtocols = "tls_passthrough"
	CreateProxyCacheAdvancedPluginProtocolsUDP            CreateProxyCacheAdvancedPluginProtocols = "udp"
	CreateProxyCacheAdvancedPluginProtocolsWs             CreateProxyCacheAdvancedPluginProtocols = "ws"
	CreateProxyCacheAdvancedPluginProtocolsWss            CreateProxyCacheAdvancedPluginProtocols = "wss"
)

func (e CreateProxyCacheAdvancedPluginProtocols) ToPointer() *CreateProxyCacheAdvancedPluginProtocols {
	return &e
}
func (e *CreateProxyCacheAdvancedPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateProxyCacheAdvancedPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateProxyCacheAdvancedPluginProtocols: %v", v)
	}
}

// CreateProxyCacheAdvancedPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateProxyCacheAdvancedPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateProxyCacheAdvancedPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateProxyCacheAdvancedPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateProxyCacheAdvancedPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateProxyCacheAdvancedPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateProxyCacheAdvancedPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateProxyCacheAdvancedPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateProxyCacheAdvancedPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateProxyCacheAdvancedPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateProxyCacheAdvancedPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateProxyCacheAdvancedPlugin struct {
	Config *CreateProxyCacheAdvancedPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                                   `json:"enabled,omitempty"`
	InstanceName *string                                 `json:"instance_name,omitempty"`
	name         *string                                 `const:"proxy-cache-advanced" json:"name,omitempty"`
	Ordering     *CreateProxyCacheAdvancedPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateProxyCacheAdvancedPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateProxyCacheAdvancedPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateProxyCacheAdvancedPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateProxyCacheAdvancedPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateProxyCacheAdvancedPluginService `json:"service,omitempty"`
}

func (c CreateProxyCacheAdvancedPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateProxyCacheAdvancedPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateProxyCacheAdvancedPlugin) GetConfig() *CreateProxyCacheAdvancedPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateProxyCacheAdvancedPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateProxyCacheAdvancedPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateProxyCacheAdvancedPlugin) GetName() *string {
	return types.String("proxy-cache-advanced")
}

func (o *CreateProxyCacheAdvancedPlugin) GetOrdering() *CreateProxyCacheAdvancedPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateProxyCacheAdvancedPlugin) GetProtocols() []CreateProxyCacheAdvancedPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateProxyCacheAdvancedPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateProxyCacheAdvancedPlugin) GetConsumer() *CreateProxyCacheAdvancedPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateProxyCacheAdvancedPlugin) GetConsumerGroup() *CreateProxyCacheAdvancedPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateProxyCacheAdvancedPlugin) GetRoute() *CreateProxyCacheAdvancedPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateProxyCacheAdvancedPlugin) GetService() *CreateProxyCacheAdvancedPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
