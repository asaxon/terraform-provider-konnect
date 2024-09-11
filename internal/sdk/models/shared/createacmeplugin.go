// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

// CreateAcmePluginAccountKey - The private key associated with the account.
type CreateAcmePluginAccountKey struct {
	// The Key ID.
	KeyID string `json:"key_id"`
	// The ID of the key set to associate the Key ID with.
	KeySet *string `json:"key_set,omitempty"`
}

func (o *CreateAcmePluginAccountKey) GetKeyID() string {
	if o == nil {
		return ""
	}
	return o.KeyID
}

func (o *CreateAcmePluginAccountKey) GetKeySet() *string {
	if o == nil {
		return nil
	}
	return o.KeySet
}

// CreateAcmePluginCertType - The certificate type to create. The possible values are `'rsa'` for RSA certificate or `'ecc'` for EC certificate.
type CreateAcmePluginCertType string

const (
	CreateAcmePluginCertTypeRsa CreateAcmePluginCertType = "rsa"
	CreateAcmePluginCertTypeEcc CreateAcmePluginCertType = "ecc"
)

func (e CreateAcmePluginCertType) ToPointer() *CreateAcmePluginCertType {
	return &e
}
func (e *CreateAcmePluginCertType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "rsa":
		fallthrough
	case "ecc":
		*e = CreateAcmePluginCertType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAcmePluginCertType: %v", v)
	}
}

// CreateAcmePluginRsaKeySize - RSA private key size for the certificate. The possible values are 2048, 3072, or 4096.
type CreateAcmePluginRsaKeySize int64

const (
	CreateAcmePluginRsaKeySizeTwoThousandAndFortyEight   CreateAcmePluginRsaKeySize = 2048
	CreateAcmePluginRsaKeySizeThreeThousandAndSeventyTwo CreateAcmePluginRsaKeySize = 3072
	CreateAcmePluginRsaKeySizeFourThousandAndNinetySix   CreateAcmePluginRsaKeySize = 4096
)

func (e CreateAcmePluginRsaKeySize) ToPointer() *CreateAcmePluginRsaKeySize {
	return &e
}
func (e *CreateAcmePluginRsaKeySize) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case 2048:
		fallthrough
	case 3072:
		fallthrough
	case 4096:
		*e = CreateAcmePluginRsaKeySize(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAcmePluginRsaKeySize: %v", v)
	}
}

// CreateAcmePluginStorage - The backend storage type to use. The possible values are `'kong'`, `'shm'`, `'redis'`, `'consul'`, or `'vault'`. In DB-less mode, `'kong'` storage is unavailable. Note that `'shm'` storage does not persist during Kong restarts and does not work for Kong running on different machines, so consider using one of `'kong'`, `'redis'`, `'consul'`, or `'vault'` in production. Please refer to the Hybrid Mode sections below as well.
type CreateAcmePluginStorage string

const (
	CreateAcmePluginStorageKong   CreateAcmePluginStorage = "kong"
	CreateAcmePluginStorageShm    CreateAcmePluginStorage = "shm"
	CreateAcmePluginStorageRedis  CreateAcmePluginStorage = "redis"
	CreateAcmePluginStorageConsul CreateAcmePluginStorage = "consul"
	CreateAcmePluginStorageVault  CreateAcmePluginStorage = "vault"
)

func (e CreateAcmePluginStorage) ToPointer() *CreateAcmePluginStorage {
	return &e
}
func (e *CreateAcmePluginStorage) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "kong":
		fallthrough
	case "shm":
		fallthrough
	case "redis":
		fallthrough
	case "consul":
		fallthrough
	case "vault":
		*e = CreateAcmePluginStorage(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAcmePluginStorage: %v", v)
	}
}

type CreateAcmePluginConsul struct {
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// Boolean representation of https.
	HTTPS *bool `json:"https,omitempty"`
	// KV prefix path.
	KvPath *string `json:"kv_path,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// Timeout in milliseconds.
	Timeout *float64 `json:"timeout,omitempty"`
	// Consul ACL token.
	Token *string `json:"token,omitempty"`
}

func (o *CreateAcmePluginConsul) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *CreateAcmePluginConsul) GetHTTPS() *bool {
	if o == nil {
		return nil
	}
	return o.HTTPS
}

func (o *CreateAcmePluginConsul) GetKvPath() *string {
	if o == nil {
		return nil
	}
	return o.KvPath
}

func (o *CreateAcmePluginConsul) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *CreateAcmePluginConsul) GetTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *CreateAcmePluginConsul) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

// CreateAcmePluginExtraOptions - Custom ACME Redis options
type CreateAcmePluginExtraOptions struct {
	// A namespace to prepend to all keys stored in Redis.
	Namespace *string `json:"namespace,omitempty"`
	// The number of keys to return in Redis SCAN calls.
	ScanCount *float64 `json:"scan_count,omitempty"`
}

func (o *CreateAcmePluginExtraOptions) GetNamespace() *string {
	if o == nil {
		return nil
	}
	return o.Namespace
}

func (o *CreateAcmePluginExtraOptions) GetScanCount() *float64 {
	if o == nil {
		return nil
	}
	return o.ScanCount
}

type CreateAcmePluginRedis struct {
	// Database to use for the Redis connection when using the `redis` strategy
	Database *int64 `json:"database,omitempty"`
	// Custom ACME Redis options
	ExtraOptions *CreateAcmePluginExtraOptions `json:"extra_options,omitempty"`
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

func (o *CreateAcmePluginRedis) GetDatabase() *int64 {
	if o == nil {
		return nil
	}
	return o.Database
}

func (o *CreateAcmePluginRedis) GetExtraOptions() *CreateAcmePluginExtraOptions {
	if o == nil {
		return nil
	}
	return o.ExtraOptions
}

func (o *CreateAcmePluginRedis) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *CreateAcmePluginRedis) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *CreateAcmePluginRedis) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *CreateAcmePluginRedis) GetServerName() *string {
	if o == nil {
		return nil
	}
	return o.ServerName
}

func (o *CreateAcmePluginRedis) GetSsl() *bool {
	if o == nil {
		return nil
	}
	return o.Ssl
}

func (o *CreateAcmePluginRedis) GetSslVerify() *bool {
	if o == nil {
		return nil
	}
	return o.SslVerify
}

func (o *CreateAcmePluginRedis) GetTimeout() *int64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *CreateAcmePluginRedis) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

type CreateAcmePluginShm struct {
	// Name of shared memory zone used for Kong API gateway storage
	ShmName *string `json:"shm_name,omitempty"`
}

func (o *CreateAcmePluginShm) GetShmName() *string {
	if o == nil {
		return nil
	}
	return o.ShmName
}

// CreateAcmePluginAuthMethod - Auth Method, default to token, can be 'token' or 'kubernetes'.
type CreateAcmePluginAuthMethod string

const (
	CreateAcmePluginAuthMethodToken      CreateAcmePluginAuthMethod = "token"
	CreateAcmePluginAuthMethodKubernetes CreateAcmePluginAuthMethod = "kubernetes"
)

func (e CreateAcmePluginAuthMethod) ToPointer() *CreateAcmePluginAuthMethod {
	return &e
}
func (e *CreateAcmePluginAuthMethod) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "token":
		fallthrough
	case "kubernetes":
		*e = CreateAcmePluginAuthMethod(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAcmePluginAuthMethod: %v", v)
	}
}

type CreateAcmePluginVault struct {
	// Auth Method, default to token, can be 'token' or 'kubernetes'.
	AuthMethod *CreateAcmePluginAuthMethod `json:"auth_method,omitempty"`
	// Vault's authentication path to use.
	AuthPath *string `json:"auth_path,omitempty"`
	// The role to try and assign.
	AuthRole *string `json:"auth_role,omitempty"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// Boolean representation of https.
	HTTPS *bool `json:"https,omitempty"`
	// The path to the JWT.
	JwtPath *string `json:"jwt_path,omitempty"`
	// KV prefix path.
	KvPath *string `json:"kv_path,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// Timeout in milliseconds.
	Timeout *float64 `json:"timeout,omitempty"`
	// SNI used in request, default to host if omitted.
	TLSServerName *string `json:"tls_server_name,omitempty"`
	// Turn on TLS verification.
	TLSVerify *bool `json:"tls_verify,omitempty"`
	// Consul ACL token.
	Token *string `json:"token,omitempty"`
}

func (o *CreateAcmePluginVault) GetAuthMethod() *CreateAcmePluginAuthMethod {
	if o == nil {
		return nil
	}
	return o.AuthMethod
}

func (o *CreateAcmePluginVault) GetAuthPath() *string {
	if o == nil {
		return nil
	}
	return o.AuthPath
}

func (o *CreateAcmePluginVault) GetAuthRole() *string {
	if o == nil {
		return nil
	}
	return o.AuthRole
}

func (o *CreateAcmePluginVault) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *CreateAcmePluginVault) GetHTTPS() *bool {
	if o == nil {
		return nil
	}
	return o.HTTPS
}

func (o *CreateAcmePluginVault) GetJwtPath() *string {
	if o == nil {
		return nil
	}
	return o.JwtPath
}

func (o *CreateAcmePluginVault) GetKvPath() *string {
	if o == nil {
		return nil
	}
	return o.KvPath
}

func (o *CreateAcmePluginVault) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *CreateAcmePluginVault) GetTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.Timeout
}

func (o *CreateAcmePluginVault) GetTLSServerName() *string {
	if o == nil {
		return nil
	}
	return o.TLSServerName
}

func (o *CreateAcmePluginVault) GetTLSVerify() *bool {
	if o == nil {
		return nil
	}
	return o.TLSVerify
}

func (o *CreateAcmePluginVault) GetToken() *string {
	if o == nil {
		return nil
	}
	return o.Token
}

type CreateAcmePluginStorageConfig struct {
	Consul *CreateAcmePluginConsul `json:"consul,omitempty"`
	Kong   map[string]any          `json:"kong,omitempty"`
	Redis  *CreateAcmePluginRedis  `json:"redis,omitempty"`
	Shm    *CreateAcmePluginShm    `json:"shm,omitempty"`
	Vault  *CreateAcmePluginVault  `json:"vault,omitempty"`
}

func (o *CreateAcmePluginStorageConfig) GetConsul() *CreateAcmePluginConsul {
	if o == nil {
		return nil
	}
	return o.Consul
}

func (o *CreateAcmePluginStorageConfig) GetKong() map[string]any {
	if o == nil {
		return nil
	}
	return o.Kong
}

func (o *CreateAcmePluginStorageConfig) GetRedis() *CreateAcmePluginRedis {
	if o == nil {
		return nil
	}
	return o.Redis
}

func (o *CreateAcmePluginStorageConfig) GetShm() *CreateAcmePluginShm {
	if o == nil {
		return nil
	}
	return o.Shm
}

func (o *CreateAcmePluginStorageConfig) GetVault() *CreateAcmePluginVault {
	if o == nil {
		return nil
	}
	return o.Vault
}

type CreateAcmePluginConfig struct {
	// The account identifier. Can be reused in a different plugin instance.
	AccountEmail *string `json:"account_email,omitempty"`
	// The private key associated with the account.
	AccountKey *CreateAcmePluginAccountKey `json:"account_key,omitempty"`
	// If set to `true`, the plugin allows all domains and ignores any values in the `domains` list.
	AllowAnyDomain *bool `json:"allow_any_domain,omitempty"`
	// A string representing a URL, such as https://example.com/path/to/resource?q=search.
	APIURI *string `json:"api_uri,omitempty"`
	// The certificate type to create. The possible values are `'rsa'` for RSA certificate or `'ecc'` for EC certificate.
	CertType *CreateAcmePluginCertType `json:"cert_type,omitempty"`
	// An array of strings representing hosts. A valid host is a string containing one or more labels separated by periods, with at most one wildcard label ('*')
	Domains []string `json:"domains,omitempty"`
	// External account binding (EAB) base64-encoded URL string of the HMAC key. You usually don't need to set this unless it is explicitly required by the CA.
	EabHmacKey *string `json:"eab_hmac_key,omitempty"`
	// External account binding (EAB) key id. You usually don't need to set this unless it is explicitly required by the CA.
	EabKid *string `json:"eab_kid,omitempty"`
	// A boolean value that controls whether to include the IPv4 address in the common name field of generated certificates.
	EnableIpv4CommonName *bool `json:"enable_ipv4_common_name,omitempty"`
	// Minutes to wait for each domain that fails to create a certificate. This applies to both a
	// new certificate and a renewal certificate.
	FailBackoffMinutes *float64 `json:"fail_backoff_minutes,omitempty"`
	// A string value that specifies the preferred certificate chain to use when generating certificates.
	PreferredChain *string `json:"preferred_chain,omitempty"`
	// Days remaining to renew the certificate before it expires.
	RenewThresholdDays *float64 `json:"renew_threshold_days,omitempty"`
	// RSA private key size for the certificate. The possible values are 2048, 3072, or 4096.
	RsaKeySize *CreateAcmePluginRsaKeySize `json:"rsa_key_size,omitempty"`
	// The backend storage type to use. The possible values are `'kong'`, `'shm'`, `'redis'`, `'consul'`, or `'vault'`. In DB-less mode, `'kong'` storage is unavailable. Note that `'shm'` storage does not persist during Kong restarts and does not work for Kong running on different machines, so consider using one of `'kong'`, `'redis'`, `'consul'`, or `'vault'` in production. Please refer to the Hybrid Mode sections below as well.
	Storage       *CreateAcmePluginStorage       `json:"storage,omitempty"`
	StorageConfig *CreateAcmePluginStorageConfig `json:"storage_config,omitempty"`
	// If you are using Let's Encrypt, you must set this to `true` to agree the terms of service.
	TosAccepted *bool `json:"tos_accepted,omitempty"`
}

func (o *CreateAcmePluginConfig) GetAccountEmail() *string {
	if o == nil {
		return nil
	}
	return o.AccountEmail
}

func (o *CreateAcmePluginConfig) GetAccountKey() *CreateAcmePluginAccountKey {
	if o == nil {
		return nil
	}
	return o.AccountKey
}

func (o *CreateAcmePluginConfig) GetAllowAnyDomain() *bool {
	if o == nil {
		return nil
	}
	return o.AllowAnyDomain
}

func (o *CreateAcmePluginConfig) GetAPIURI() *string {
	if o == nil {
		return nil
	}
	return o.APIURI
}

func (o *CreateAcmePluginConfig) GetCertType() *CreateAcmePluginCertType {
	if o == nil {
		return nil
	}
	return o.CertType
}

func (o *CreateAcmePluginConfig) GetDomains() []string {
	if o == nil {
		return nil
	}
	return o.Domains
}

func (o *CreateAcmePluginConfig) GetEabHmacKey() *string {
	if o == nil {
		return nil
	}
	return o.EabHmacKey
}

func (o *CreateAcmePluginConfig) GetEabKid() *string {
	if o == nil {
		return nil
	}
	return o.EabKid
}

func (o *CreateAcmePluginConfig) GetEnableIpv4CommonName() *bool {
	if o == nil {
		return nil
	}
	return o.EnableIpv4CommonName
}

func (o *CreateAcmePluginConfig) GetFailBackoffMinutes() *float64 {
	if o == nil {
		return nil
	}
	return o.FailBackoffMinutes
}

func (o *CreateAcmePluginConfig) GetPreferredChain() *string {
	if o == nil {
		return nil
	}
	return o.PreferredChain
}

func (o *CreateAcmePluginConfig) GetRenewThresholdDays() *float64 {
	if o == nil {
		return nil
	}
	return o.RenewThresholdDays
}

func (o *CreateAcmePluginConfig) GetRsaKeySize() *CreateAcmePluginRsaKeySize {
	if o == nil {
		return nil
	}
	return o.RsaKeySize
}

func (o *CreateAcmePluginConfig) GetStorage() *CreateAcmePluginStorage {
	if o == nil {
		return nil
	}
	return o.Storage
}

func (o *CreateAcmePluginConfig) GetStorageConfig() *CreateAcmePluginStorageConfig {
	if o == nil {
		return nil
	}
	return o.StorageConfig
}

func (o *CreateAcmePluginConfig) GetTosAccepted() *bool {
	if o == nil {
		return nil
	}
	return o.TosAccepted
}

type CreateAcmePluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateAcmePluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateAcmePluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateAcmePluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateAcmePluginOrdering struct {
	After  *CreateAcmePluginAfter  `json:"after,omitempty"`
	Before *CreateAcmePluginBefore `json:"before,omitempty"`
}

func (o *CreateAcmePluginOrdering) GetAfter() *CreateAcmePluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *CreateAcmePluginOrdering) GetBefore() *CreateAcmePluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type CreateAcmePluginProtocols string

const (
	CreateAcmePluginProtocolsGrpc           CreateAcmePluginProtocols = "grpc"
	CreateAcmePluginProtocolsGrpcs          CreateAcmePluginProtocols = "grpcs"
	CreateAcmePluginProtocolsHTTP           CreateAcmePluginProtocols = "http"
	CreateAcmePluginProtocolsHTTPS          CreateAcmePluginProtocols = "https"
	CreateAcmePluginProtocolsTCP            CreateAcmePluginProtocols = "tcp"
	CreateAcmePluginProtocolsTLS            CreateAcmePluginProtocols = "tls"
	CreateAcmePluginProtocolsTLSPassthrough CreateAcmePluginProtocols = "tls_passthrough"
	CreateAcmePluginProtocolsUDP            CreateAcmePluginProtocols = "udp"
	CreateAcmePluginProtocolsWs             CreateAcmePluginProtocols = "ws"
	CreateAcmePluginProtocolsWss            CreateAcmePluginProtocols = "wss"
)

func (e CreateAcmePluginProtocols) ToPointer() *CreateAcmePluginProtocols {
	return &e
}
func (e *CreateAcmePluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateAcmePluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateAcmePluginProtocols: %v", v)
	}
}

// CreateAcmePluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateAcmePluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAcmePluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateAcmePluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAcmePluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateAcmePluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateAcmePluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAcmePluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateAcmePluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateAcmePluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateAcmePluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateAcmePlugin struct {
	Config *CreateAcmePluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                     `json:"enabled,omitempty"`
	InstanceName *string                   `json:"instance_name,omitempty"`
	name         *string                   `const:"acme" json:"name,omitempty"`
	Ordering     *CreateAcmePluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateAcmePluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateAcmePluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateAcmePluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateAcmePluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateAcmePluginService `json:"service,omitempty"`
}

func (c CreateAcmePlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateAcmePlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateAcmePlugin) GetConfig() *CreateAcmePluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateAcmePlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateAcmePlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateAcmePlugin) GetName() *string {
	return types.String("acme")
}

func (o *CreateAcmePlugin) GetOrdering() *CreateAcmePluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateAcmePlugin) GetProtocols() []CreateAcmePluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateAcmePlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateAcmePlugin) GetConsumer() *CreateAcmePluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateAcmePlugin) GetConsumerGroup() *CreateAcmePluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateAcmePlugin) GetRoute() *CreateAcmePluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateAcmePlugin) GetService() *CreateAcmePluginService {
	if o == nil {
		return nil
	}
	return o.Service
}
