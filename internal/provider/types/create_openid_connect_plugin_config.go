// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateOpenidConnectPluginConfig struct {
	Anonymous                                    types.String                                `tfsdk:"anonymous"`
	Audience                                     []types.String                              `tfsdk:"audience"`
	AudienceClaim                                []types.String                              `tfsdk:"audience_claim"`
	AudienceRequired                             []types.String                              `tfsdk:"audience_required"`
	AuthMethods                                  []types.String                              `tfsdk:"auth_methods"`
	AuthenticatedGroupsClaim                     []types.String                              `tfsdk:"authenticated_groups_claim"`
	AuthorizationCookieDomain                    types.String                                `tfsdk:"authorization_cookie_domain"`
	AuthorizationCookieHTTPOnly                  types.Bool                                  `tfsdk:"authorization_cookie_http_only"`
	AuthorizationCookieName                      types.String                                `tfsdk:"authorization_cookie_name"`
	AuthorizationCookiePath                      types.String                                `tfsdk:"authorization_cookie_path"`
	AuthorizationCookieSameSite                  types.String                                `tfsdk:"authorization_cookie_same_site"`
	AuthorizationCookieSecure                    types.Bool                                  `tfsdk:"authorization_cookie_secure"`
	AuthorizationEndpoint                        types.String                                `tfsdk:"authorization_endpoint"`
	AuthorizationQueryArgsClient                 []types.String                              `tfsdk:"authorization_query_args_client"`
	AuthorizationQueryArgsNames                  []types.String                              `tfsdk:"authorization_query_args_names"`
	AuthorizationQueryArgsValues                 []types.String                              `tfsdk:"authorization_query_args_values"`
	AuthorizationRollingTimeout                  types.Number                                `tfsdk:"authorization_rolling_timeout"`
	BearerTokenCookieName                        types.String                                `tfsdk:"bearer_token_cookie_name"`
	BearerTokenParamType                         []types.String                              `tfsdk:"bearer_token_param_type"`
	ByUsernameIgnoreCase                         types.Bool                                  `tfsdk:"by_username_ignore_case"`
	CacheIntrospection                           types.Bool                                  `tfsdk:"cache_introspection"`
	CacheTokenExchange                           types.Bool                                  `tfsdk:"cache_token_exchange"`
	CacheTokens                                  types.Bool                                  `tfsdk:"cache_tokens"`
	CacheTokensSalt                              types.String                                `tfsdk:"cache_tokens_salt"`
	CacheTTL                                     types.Number                                `tfsdk:"cache_ttl"`
	CacheTTLMax                                  types.Number                                `tfsdk:"cache_ttl_max"`
	CacheTTLMin                                  types.Number                                `tfsdk:"cache_ttl_min"`
	CacheTTLNeg                                  types.Number                                `tfsdk:"cache_ttl_neg"`
	CacheTTLResurrect                            types.Number                                `tfsdk:"cache_ttl_resurrect"`
	CacheUserInfo                                types.Bool                                  `tfsdk:"cache_user_info"`
	ClaimsForbidden                              []types.String                              `tfsdk:"claims_forbidden"`
	ClientAlg                                    []types.String                              `tfsdk:"client_alg"`
	ClientArg                                    types.String                                `tfsdk:"client_arg"`
	ClientAuth                                   []types.String                              `tfsdk:"client_auth"`
	ClientCredentialsParamType                   []types.String                              `tfsdk:"client_credentials_param_type"`
	ClientID                                     []types.String                              `tfsdk:"client_id"`
	ClientJwk                                    []KonnectApplicationAuthPluginClientJwk     `tfsdk:"client_jwk"`
	ClientSecret                                 []types.String                              `tfsdk:"client_secret"`
	ClusterCacheRedis                            *CreateGraphqlProxyCacheAdvancedPluginRedis `tfsdk:"cluster_cache_redis"`
	ClusterCacheStrategy                         types.String                                `tfsdk:"cluster_cache_strategy"`
	ConsumerBy                                   []types.String                              `tfsdk:"consumer_by"`
	ConsumerClaim                                []types.String                              `tfsdk:"consumer_claim"`
	ConsumerOptional                             types.Bool                                  `tfsdk:"consumer_optional"`
	CredentialClaim                              []types.String                              `tfsdk:"credential_claim"`
	DisableSession                               []types.String                              `tfsdk:"disable_session"`
	DiscoveryHeadersNames                        []types.String                              `tfsdk:"discovery_headers_names"`
	DiscoveryHeadersValues                       []types.String                              `tfsdk:"discovery_headers_values"`
	DisplayErrors                                types.Bool                                  `tfsdk:"display_errors"`
	Domains                                      []types.String                              `tfsdk:"domains"`
	DownstreamAccessTokenHeader                  types.String                                `tfsdk:"downstream_access_token_header"`
	DownstreamAccessTokenJwkHeader               types.String                                `tfsdk:"downstream_access_token_jwk_header"`
	DownstreamHeadersClaims                      []types.String                              `tfsdk:"downstream_headers_claims"`
	DownstreamHeadersNames                       []types.String                              `tfsdk:"downstream_headers_names"`
	DownstreamIDTokenHeader                      types.String                                `tfsdk:"downstream_id_token_header"`
	DownstreamIDTokenJwkHeader                   types.String                                `tfsdk:"downstream_id_token_jwk_header"`
	DownstreamIntrospectionHeader                types.String                                `tfsdk:"downstream_introspection_header"`
	DownstreamIntrospectionJwtHeader             types.String                                `tfsdk:"downstream_introspection_jwt_header"`
	DownstreamRefreshTokenHeader                 types.String                                `tfsdk:"downstream_refresh_token_header"`
	DownstreamSessionIDHeader                    types.String                                `tfsdk:"downstream_session_id_header"`
	DownstreamUserInfoHeader                     types.String                                `tfsdk:"downstream_user_info_header"`
	DownstreamUserInfoJwtHeader                  types.String                                `tfsdk:"downstream_user_info_jwt_header"`
	DpopProofLifetime                            types.Number                                `tfsdk:"dpop_proof_lifetime"`
	DpopUseNonce                                 types.Bool                                  `tfsdk:"dpop_use_nonce"`
	EnableHsSignatures                           types.Bool                                  `tfsdk:"enable_hs_signatures"`
	EndSessionEndpoint                           types.String                                `tfsdk:"end_session_endpoint"`
	ExposeErrorCode                              types.Bool                                  `tfsdk:"expose_error_code"`
	ExtraJwksUris                                []types.String                              `tfsdk:"extra_jwks_uris"`
	ForbiddenDestroySession                      types.Bool                                  `tfsdk:"forbidden_destroy_session"`
	ForbiddenErrorMessage                        types.String                                `tfsdk:"forbidden_error_message"`
	ForbiddenRedirectURI                         []types.String                              `tfsdk:"forbidden_redirect_uri"`
	GroupsClaim                                  []types.String                              `tfsdk:"groups_claim"`
	GroupsRequired                               []types.String                              `tfsdk:"groups_required"`
	HideCredentials                              types.Bool                                  `tfsdk:"hide_credentials"`
	HTTPProxy                                    types.String                                `tfsdk:"http_proxy"`
	HTTPProxyAuthorization                       types.String                                `tfsdk:"http_proxy_authorization"`
	HTTPVersion                                  types.Number                                `tfsdk:"http_version"`
	HTTPSProxy                                   types.String                                `tfsdk:"https_proxy"`
	HTTPSProxyAuthorization                      types.String                                `tfsdk:"https_proxy_authorization"`
	IDTokenParamName                             types.String                                `tfsdk:"id_token_param_name"`
	IDTokenParamType                             []types.String                              `tfsdk:"id_token_param_type"`
	IgnoreSignature                              []types.String                              `tfsdk:"ignore_signature"`
	IntrospectJwtTokens                          types.Bool                                  `tfsdk:"introspect_jwt_tokens"`
	IntrospectionAccept                          types.String                                `tfsdk:"introspection_accept"`
	IntrospectionCheckActive                     types.Bool                                  `tfsdk:"introspection_check_active"`
	IntrospectionEndpoint                        types.String                                `tfsdk:"introspection_endpoint"`
	IntrospectionEndpointAuthMethod              types.String                                `tfsdk:"introspection_endpoint_auth_method"`
	IntrospectionHeadersClient                   []types.String                              `tfsdk:"introspection_headers_client"`
	IntrospectionHeadersNames                    []types.String                              `tfsdk:"introspection_headers_names"`
	IntrospectionHeadersValues                   []types.String                              `tfsdk:"introspection_headers_values"`
	IntrospectionHint                            types.String                                `tfsdk:"introspection_hint"`
	IntrospectionPostArgsClient                  []types.String                              `tfsdk:"introspection_post_args_client"`
	IntrospectionPostArgsNames                   []types.String                              `tfsdk:"introspection_post_args_names"`
	IntrospectionPostArgsValues                  []types.String                              `tfsdk:"introspection_post_args_values"`
	IntrospectionTokenParamName                  types.String                                `tfsdk:"introspection_token_param_name"`
	Issuer                                       types.String                                `tfsdk:"issuer"`
	IssuersAllowed                               []types.String                              `tfsdk:"issuers_allowed"`
	JwtSessionClaim                              types.String                                `tfsdk:"jwt_session_claim"`
	JwtSessionCookie                             types.String                                `tfsdk:"jwt_session_cookie"`
	Keepalive                                    types.Bool                                  `tfsdk:"keepalive"`
	Leeway                                       types.Number                                `tfsdk:"leeway"`
	LoginAction                                  types.String                                `tfsdk:"login_action"`
	LoginMethods                                 []types.String                              `tfsdk:"login_methods"`
	LoginRedirectMode                            types.String                                `tfsdk:"login_redirect_mode"`
	LoginRedirectURI                             []types.String                              `tfsdk:"login_redirect_uri"`
	LoginTokens                                  []types.String                              `tfsdk:"login_tokens"`
	LogoutMethods                                []types.String                              `tfsdk:"logout_methods"`
	LogoutPostArg                                types.String                                `tfsdk:"logout_post_arg"`
	LogoutQueryArg                               types.String                                `tfsdk:"logout_query_arg"`
	LogoutRedirectURI                            []types.String                              `tfsdk:"logout_redirect_uri"`
	LogoutRevoke                                 types.Bool                                  `tfsdk:"logout_revoke"`
	LogoutRevokeAccessToken                      types.Bool                                  `tfsdk:"logout_revoke_access_token"`
	LogoutRevokeRefreshToken                     types.Bool                                  `tfsdk:"logout_revoke_refresh_token"`
	LogoutURISuffix                              types.String                                `tfsdk:"logout_uri_suffix"`
	MaxAge                                       types.Number                                `tfsdk:"max_age"`
	MtlsIntrospectionEndpoint                    types.String                                `tfsdk:"mtls_introspection_endpoint"`
	MtlsRevocationEndpoint                       types.String                                `tfsdk:"mtls_revocation_endpoint"`
	MtlsTokenEndpoint                            types.String                                `tfsdk:"mtls_token_endpoint"`
	NoProxy                                      types.String                                `tfsdk:"no_proxy"`
	PasswordParamType                            []types.String                              `tfsdk:"password_param_type"`
	PreserveQueryArgs                            types.Bool                                  `tfsdk:"preserve_query_args"`
	ProofOfPossessionAuthMethodsValidation       types.Bool                                  `tfsdk:"proof_of_possession_auth_methods_validation"`
	ProofOfPossessionDpop                        types.String                                `tfsdk:"proof_of_possession_dpop"`
	ProofOfPossessionMtls                        types.String                                `tfsdk:"proof_of_possession_mtls"`
	PushedAuthorizationRequestEndpoint           types.String                                `tfsdk:"pushed_authorization_request_endpoint"`
	PushedAuthorizationRequestEndpointAuthMethod types.String                                `tfsdk:"pushed_authorization_request_endpoint_auth_method"`
	RedirectURI                                  []types.String                              `tfsdk:"redirect_uri"`
	Redis                                        *CreateKonnectApplicationAuthPluginRedis    `tfsdk:"redis"`
	RediscoveryLifetime                          types.Number                                `tfsdk:"rediscovery_lifetime"`
	RefreshTokenParamName                        types.String                                `tfsdk:"refresh_token_param_name"`
	RefreshTokenParamType                        []types.String                              `tfsdk:"refresh_token_param_type"`
	RefreshTokens                                types.Bool                                  `tfsdk:"refresh_tokens"`
	RequireProofKeyForCodeExchange               types.Bool                                  `tfsdk:"require_proof_key_for_code_exchange"`
	RequirePushedAuthorizationRequests           types.Bool                                  `tfsdk:"require_pushed_authorization_requests"`
	RequireSignedRequestObject                   types.Bool                                  `tfsdk:"require_signed_request_object"`
	ResolveDistributedClaims                     types.Bool                                  `tfsdk:"resolve_distributed_claims"`
	ResponseMode                                 types.String                                `tfsdk:"response_mode"`
	ResponseType                                 []types.String                              `tfsdk:"response_type"`
	Reverify                                     types.Bool                                  `tfsdk:"reverify"`
	RevocationEndpoint                           types.String                                `tfsdk:"revocation_endpoint"`
	RevocationEndpointAuthMethod                 types.String                                `tfsdk:"revocation_endpoint_auth_method"`
	RevocationTokenParamName                     types.String                                `tfsdk:"revocation_token_param_name"`
	RolesClaim                                   []types.String                              `tfsdk:"roles_claim"`
	RolesRequired                                []types.String                              `tfsdk:"roles_required"`
	RunOnPreflight                               types.Bool                                  `tfsdk:"run_on_preflight"`
	Scopes                                       []types.String                              `tfsdk:"scopes"`
	ScopesClaim                                  []types.String                              `tfsdk:"scopes_claim"`
	ScopesRequired                               []types.String                              `tfsdk:"scopes_required"`
	SearchUserInfo                               types.Bool                                  `tfsdk:"search_user_info"`
	SessionAbsoluteTimeout                       types.Number                                `tfsdk:"session_absolute_timeout"`
	SessionAudience                              types.String                                `tfsdk:"session_audience"`
	SessionCookieDomain                          types.String                                `tfsdk:"session_cookie_domain"`
	SessionCookieHTTPOnly                        types.Bool                                  `tfsdk:"session_cookie_http_only"`
	SessionCookieName                            types.String                                `tfsdk:"session_cookie_name"`
	SessionCookiePath                            types.String                                `tfsdk:"session_cookie_path"`
	SessionCookieSameSite                        types.String                                `tfsdk:"session_cookie_same_site"`
	SessionCookieSecure                          types.Bool                                  `tfsdk:"session_cookie_secure"`
	SessionEnforceSameSubject                    types.Bool                                  `tfsdk:"session_enforce_same_subject"`
	SessionHashStorageKey                        types.Bool                                  `tfsdk:"session_hash_storage_key"`
	SessionHashSubject                           types.Bool                                  `tfsdk:"session_hash_subject"`
	SessionIdlingTimeout                         types.Number                                `tfsdk:"session_idling_timeout"`
	SessionMemcachedHost                         types.String                                `tfsdk:"session_memcached_host"`
	SessionMemcachedPort                         types.Int64                                 `tfsdk:"session_memcached_port"`
	SessionMemcachedPrefix                       types.String                                `tfsdk:"session_memcached_prefix"`
	SessionMemcachedSocket                       types.String                                `tfsdk:"session_memcached_socket"`
	SessionRemember                              types.Bool                                  `tfsdk:"session_remember"`
	SessionRememberAbsoluteTimeout               types.Number                                `tfsdk:"session_remember_absolute_timeout"`
	SessionRememberCookieName                    types.String                                `tfsdk:"session_remember_cookie_name"`
	SessionRememberRollingTimeout                types.Number                                `tfsdk:"session_remember_rolling_timeout"`
	SessionRequestHeaders                        []types.String                              `tfsdk:"session_request_headers"`
	SessionResponseHeaders                       []types.String                              `tfsdk:"session_response_headers"`
	SessionRollingTimeout                        types.Number                                `tfsdk:"session_rolling_timeout"`
	SessionSecret                                types.String                                `tfsdk:"session_secret"`
	SessionStorage                               types.String                                `tfsdk:"session_storage"`
	SessionStoreMetadata                         types.Bool                                  `tfsdk:"session_store_metadata"`
	SslVerify                                    types.Bool                                  `tfsdk:"ssl_verify"`
	Timeout                                      types.Number                                `tfsdk:"timeout"`
	TLSClientAuthCertID                          types.String                                `tfsdk:"tls_client_auth_cert_id"`
	TLSClientAuthSslVerify                       types.Bool                                  `tfsdk:"tls_client_auth_ssl_verify"`
	TokenCacheKeyIncludeScope                    types.Bool                                  `tfsdk:"token_cache_key_include_scope"`
	TokenEndpoint                                types.String                                `tfsdk:"token_endpoint"`
	TokenEndpointAuthMethod                      types.String                                `tfsdk:"token_endpoint_auth_method"`
	TokenExchangeEndpoint                        types.String                                `tfsdk:"token_exchange_endpoint"`
	TokenHeadersClient                           []types.String                              `tfsdk:"token_headers_client"`
	TokenHeadersGrants                           []types.String                              `tfsdk:"token_headers_grants"`
	TokenHeadersNames                            []types.String                              `tfsdk:"token_headers_names"`
	TokenHeadersPrefix                           types.String                                `tfsdk:"token_headers_prefix"`
	TokenHeadersReplay                           []types.String                              `tfsdk:"token_headers_replay"`
	TokenHeadersValues                           []types.String                              `tfsdk:"token_headers_values"`
	TokenPostArgsClient                          []types.String                              `tfsdk:"token_post_args_client"`
	TokenPostArgsNames                           []types.String                              `tfsdk:"token_post_args_names"`
	TokenPostArgsValues                          []types.String                              `tfsdk:"token_post_args_values"`
	UnauthorizedDestroySession                   types.Bool                                  `tfsdk:"unauthorized_destroy_session"`
	UnauthorizedErrorMessage                     types.String                                `tfsdk:"unauthorized_error_message"`
	UnauthorizedRedirectURI                      []types.String                              `tfsdk:"unauthorized_redirect_uri"`
	UnexpectedRedirectURI                        []types.String                              `tfsdk:"unexpected_redirect_uri"`
	UpstreamAccessTokenHeader                    types.String                                `tfsdk:"upstream_access_token_header"`
	UpstreamAccessTokenJwkHeader                 types.String                                `tfsdk:"upstream_access_token_jwk_header"`
	UpstreamHeadersClaims                        []types.String                              `tfsdk:"upstream_headers_claims"`
	UpstreamHeadersNames                         []types.String                              `tfsdk:"upstream_headers_names"`
	UpstreamIDTokenHeader                        types.String                                `tfsdk:"upstream_id_token_header"`
	UpstreamIDTokenJwkHeader                     types.String                                `tfsdk:"upstream_id_token_jwk_header"`
	UpstreamIntrospectionHeader                  types.String                                `tfsdk:"upstream_introspection_header"`
	UpstreamIntrospectionJwtHeader               types.String                                `tfsdk:"upstream_introspection_jwt_header"`
	UpstreamRefreshTokenHeader                   types.String                                `tfsdk:"upstream_refresh_token_header"`
	UpstreamSessionIDHeader                      types.String                                `tfsdk:"upstream_session_id_header"`
	UpstreamUserInfoHeader                       types.String                                `tfsdk:"upstream_user_info_header"`
	UpstreamUserInfoJwtHeader                    types.String                                `tfsdk:"upstream_user_info_jwt_header"`
	UserinfoAccept                               types.String                                `tfsdk:"userinfo_accept"`
	UserinfoEndpoint                             types.String                                `tfsdk:"userinfo_endpoint"`
	UserinfoHeadersClient                        []types.String                              `tfsdk:"userinfo_headers_client"`
	UserinfoHeadersNames                         []types.String                              `tfsdk:"userinfo_headers_names"`
	UserinfoHeadersValues                        []types.String                              `tfsdk:"userinfo_headers_values"`
	UserinfoQueryArgsClient                      []types.String                              `tfsdk:"userinfo_query_args_client"`
	UserinfoQueryArgsNames                       []types.String                              `tfsdk:"userinfo_query_args_names"`
	UserinfoQueryArgsValues                      []types.String                              `tfsdk:"userinfo_query_args_values"`
	UsingPseudoIssuer                            types.Bool                                  `tfsdk:"using_pseudo_issuer"`
	VerifyClaims                                 types.Bool                                  `tfsdk:"verify_claims"`
	VerifyNonce                                  types.Bool                                  `tfsdk:"verify_nonce"`
	VerifyParameters                             types.Bool                                  `tfsdk:"verify_parameters"`
	VerifySignature                              types.Bool                                  `tfsdk:"verify_signature"`
}
