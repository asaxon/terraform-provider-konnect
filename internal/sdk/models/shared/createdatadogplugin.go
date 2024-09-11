// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
	"github.com/kong/terraform-provider-konnect/internal/sdk/internal/utils"
	"github.com/kong/terraform-provider-konnect/internal/sdk/types"
)

// CreateDatadogPluginConsumerIdentifier - Authenticated user detail
type CreateDatadogPluginConsumerIdentifier string

const (
	CreateDatadogPluginConsumerIdentifierConsumerID CreateDatadogPluginConsumerIdentifier = "consumer_id"
	CreateDatadogPluginConsumerIdentifierCustomID   CreateDatadogPluginConsumerIdentifier = "custom_id"
	CreateDatadogPluginConsumerIdentifierUsername   CreateDatadogPluginConsumerIdentifier = "username"
)

func (e CreateDatadogPluginConsumerIdentifier) ToPointer() *CreateDatadogPluginConsumerIdentifier {
	return &e
}
func (e *CreateDatadogPluginConsumerIdentifier) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "consumer_id":
		fallthrough
	case "custom_id":
		fallthrough
	case "username":
		*e = CreateDatadogPluginConsumerIdentifier(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDatadogPluginConsumerIdentifier: %v", v)
	}
}

// CreateDatadogPluginName - Datadog metric’s name
type CreateDatadogPluginName string

const (
	CreateDatadogPluginNameKongLatency     CreateDatadogPluginName = "kong_latency"
	CreateDatadogPluginNameLatency         CreateDatadogPluginName = "latency"
	CreateDatadogPluginNameRequestCount    CreateDatadogPluginName = "request_count"
	CreateDatadogPluginNameRequestSize     CreateDatadogPluginName = "request_size"
	CreateDatadogPluginNameResponseSize    CreateDatadogPluginName = "response_size"
	CreateDatadogPluginNameUpstreamLatency CreateDatadogPluginName = "upstream_latency"
)

func (e CreateDatadogPluginName) ToPointer() *CreateDatadogPluginName {
	return &e
}
func (e *CreateDatadogPluginName) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "kong_latency":
		fallthrough
	case "latency":
		fallthrough
	case "request_count":
		fallthrough
	case "request_size":
		fallthrough
	case "response_size":
		fallthrough
	case "upstream_latency":
		*e = CreateDatadogPluginName(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDatadogPluginName: %v", v)
	}
}

// CreateDatadogPluginStatType - Determines what sort of event the metric represents
type CreateDatadogPluginStatType string

const (
	CreateDatadogPluginStatTypeCounter      CreateDatadogPluginStatType = "counter"
	CreateDatadogPluginStatTypeGauge        CreateDatadogPluginStatType = "gauge"
	CreateDatadogPluginStatTypeHistogram    CreateDatadogPluginStatType = "histogram"
	CreateDatadogPluginStatTypeMeter        CreateDatadogPluginStatType = "meter"
	CreateDatadogPluginStatTypeSet          CreateDatadogPluginStatType = "set"
	CreateDatadogPluginStatTypeTimer        CreateDatadogPluginStatType = "timer"
	CreateDatadogPluginStatTypeDistribution CreateDatadogPluginStatType = "distribution"
)

func (e CreateDatadogPluginStatType) ToPointer() *CreateDatadogPluginStatType {
	return &e
}
func (e *CreateDatadogPluginStatType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "counter":
		fallthrough
	case "gauge":
		fallthrough
	case "histogram":
		fallthrough
	case "meter":
		fallthrough
	case "set":
		fallthrough
	case "timer":
		fallthrough
	case "distribution":
		*e = CreateDatadogPluginStatType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDatadogPluginStatType: %v", v)
	}
}

type CreateDatadogPluginMetrics struct {
	// Authenticated user detail
	ConsumerIdentifier *CreateDatadogPluginConsumerIdentifier `json:"consumer_identifier,omitempty"`
	// Datadog metric’s name
	Name CreateDatadogPluginName `json:"name"`
	// Sampling rate
	SampleRate *float64 `json:"sample_rate,omitempty"`
	// Determines what sort of event the metric represents
	StatType CreateDatadogPluginStatType `json:"stat_type"`
	// List of tags
	Tags []string `json:"tags,omitempty"`
}

func (o *CreateDatadogPluginMetrics) GetConsumerIdentifier() *CreateDatadogPluginConsumerIdentifier {
	if o == nil {
		return nil
	}
	return o.ConsumerIdentifier
}

func (o *CreateDatadogPluginMetrics) GetName() CreateDatadogPluginName {
	if o == nil {
		return CreateDatadogPluginName("")
	}
	return o.Name
}

func (o *CreateDatadogPluginMetrics) GetSampleRate() *float64 {
	if o == nil {
		return nil
	}
	return o.SampleRate
}

func (o *CreateDatadogPluginMetrics) GetStatType() CreateDatadogPluginStatType {
	if o == nil {
		return CreateDatadogPluginStatType("")
	}
	return o.StatType
}

func (o *CreateDatadogPluginMetrics) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

// CreateDatadogPluginConcurrencyLimit - The number of of queue delivery timers. -1 indicates unlimited.
type CreateDatadogPluginConcurrencyLimit int64

const (
	CreateDatadogPluginConcurrencyLimitMinus1 CreateDatadogPluginConcurrencyLimit = -1
	CreateDatadogPluginConcurrencyLimitOne    CreateDatadogPluginConcurrencyLimit = 1
)

func (e CreateDatadogPluginConcurrencyLimit) ToPointer() *CreateDatadogPluginConcurrencyLimit {
	return &e
}
func (e *CreateDatadogPluginConcurrencyLimit) UnmarshalJSON(data []byte) error {
	var v int64
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case -1:
		fallthrough
	case 1:
		*e = CreateDatadogPluginConcurrencyLimit(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDatadogPluginConcurrencyLimit: %v", v)
	}
}

type CreateDatadogPluginQueue struct {
	// The number of of queue delivery timers. -1 indicates unlimited.
	ConcurrencyLimit *CreateDatadogPluginConcurrencyLimit `json:"concurrency_limit,omitempty"`
	// Time in seconds before the initial retry is made for a failing batch.
	InitialRetryDelay *float64 `json:"initial_retry_delay,omitempty"`
	// Maximum number of entries that can be processed at a time.
	MaxBatchSize *int64 `json:"max_batch_size,omitempty"`
	// Maximum number of bytes that can be waiting on a queue, requires string content.
	MaxBytes *int64 `json:"max_bytes,omitempty"`
	// Maximum number of (fractional) seconds to elapse after the first entry was queued before the queue starts calling the handler.
	MaxCoalescingDelay *float64 `json:"max_coalescing_delay,omitempty"`
	// Maximum number of entries that can be waiting on the queue.
	MaxEntries *int64 `json:"max_entries,omitempty"`
	// Maximum time in seconds between retries, caps exponential backoff.
	MaxRetryDelay *float64 `json:"max_retry_delay,omitempty"`
	// Time in seconds before the queue gives up calling a failed handler for a batch.
	MaxRetryTime *float64 `json:"max_retry_time,omitempty"`
}

func (o *CreateDatadogPluginQueue) GetConcurrencyLimit() *CreateDatadogPluginConcurrencyLimit {
	if o == nil {
		return nil
	}
	return o.ConcurrencyLimit
}

func (o *CreateDatadogPluginQueue) GetInitialRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.InitialRetryDelay
}

func (o *CreateDatadogPluginQueue) GetMaxBatchSize() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBatchSize
}

func (o *CreateDatadogPluginQueue) GetMaxBytes() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxBytes
}

func (o *CreateDatadogPluginQueue) GetMaxCoalescingDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxCoalescingDelay
}

func (o *CreateDatadogPluginQueue) GetMaxEntries() *int64 {
	if o == nil {
		return nil
	}
	return o.MaxEntries
}

func (o *CreateDatadogPluginQueue) GetMaxRetryDelay() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryDelay
}

func (o *CreateDatadogPluginQueue) GetMaxRetryTime() *float64 {
	if o == nil {
		return nil
	}
	return o.MaxRetryTime
}

type CreateDatadogPluginConfig struct {
	// String to be attached as tag of the consumer.
	ConsumerTag *string `json:"consumer_tag,omitempty"`
	// Optional time in seconds. If `queue_size` > 1, this is the max idle time before sending a log with less than `queue_size` records.
	FlushTimeout *float64 `json:"flush_timeout,omitempty"`
	// A string representing a host name, such as example.com.
	Host *string `json:"host,omitempty"`
	// List of metrics to be logged.
	Metrics []CreateDatadogPluginMetrics `json:"metrics,omitempty"`
	// An integer representing a port number between 0 and 65535, inclusive.
	Port *int64 `json:"port,omitempty"`
	// String to be attached as a prefix to a metric's name.
	Prefix *string                   `json:"prefix,omitempty"`
	Queue  *CreateDatadogPluginQueue `json:"queue,omitempty"`
	// Maximum number of log entries to be sent on each message to the upstream server.
	QueueSize *int64 `json:"queue_size,omitempty"`
	// Number of times to retry when sending data to the upstream server.
	RetryCount *int64 `json:"retry_count,omitempty"`
	// String to be attached as the name of the service.
	ServiceNameTag *string `json:"service_name_tag,omitempty"`
	// String to be attached as the tag of the HTTP status.
	StatusTag *string `json:"status_tag,omitempty"`
}

func (o *CreateDatadogPluginConfig) GetConsumerTag() *string {
	if o == nil {
		return nil
	}
	return o.ConsumerTag
}

func (o *CreateDatadogPluginConfig) GetFlushTimeout() *float64 {
	if o == nil {
		return nil
	}
	return o.FlushTimeout
}

func (o *CreateDatadogPluginConfig) GetHost() *string {
	if o == nil {
		return nil
	}
	return o.Host
}

func (o *CreateDatadogPluginConfig) GetMetrics() []CreateDatadogPluginMetrics {
	if o == nil {
		return nil
	}
	return o.Metrics
}

func (o *CreateDatadogPluginConfig) GetPort() *int64 {
	if o == nil {
		return nil
	}
	return o.Port
}

func (o *CreateDatadogPluginConfig) GetPrefix() *string {
	if o == nil {
		return nil
	}
	return o.Prefix
}

func (o *CreateDatadogPluginConfig) GetQueue() *CreateDatadogPluginQueue {
	if o == nil {
		return nil
	}
	return o.Queue
}

func (o *CreateDatadogPluginConfig) GetQueueSize() *int64 {
	if o == nil {
		return nil
	}
	return o.QueueSize
}

func (o *CreateDatadogPluginConfig) GetRetryCount() *int64 {
	if o == nil {
		return nil
	}
	return o.RetryCount
}

func (o *CreateDatadogPluginConfig) GetServiceNameTag() *string {
	if o == nil {
		return nil
	}
	return o.ServiceNameTag
}

func (o *CreateDatadogPluginConfig) GetStatusTag() *string {
	if o == nil {
		return nil
	}
	return o.StatusTag
}

type CreateDatadogPluginAfter struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateDatadogPluginAfter) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateDatadogPluginBefore struct {
	Access []string `json:"access,omitempty"`
}

func (o *CreateDatadogPluginBefore) GetAccess() []string {
	if o == nil {
		return nil
	}
	return o.Access
}

type CreateDatadogPluginOrdering struct {
	After  *CreateDatadogPluginAfter  `json:"after,omitempty"`
	Before *CreateDatadogPluginBefore `json:"before,omitempty"`
}

func (o *CreateDatadogPluginOrdering) GetAfter() *CreateDatadogPluginAfter {
	if o == nil {
		return nil
	}
	return o.After
}

func (o *CreateDatadogPluginOrdering) GetBefore() *CreateDatadogPluginBefore {
	if o == nil {
		return nil
	}
	return o.Before
}

type CreateDatadogPluginProtocols string

const (
	CreateDatadogPluginProtocolsGrpc           CreateDatadogPluginProtocols = "grpc"
	CreateDatadogPluginProtocolsGrpcs          CreateDatadogPluginProtocols = "grpcs"
	CreateDatadogPluginProtocolsHTTP           CreateDatadogPluginProtocols = "http"
	CreateDatadogPluginProtocolsHTTPS          CreateDatadogPluginProtocols = "https"
	CreateDatadogPluginProtocolsTCP            CreateDatadogPluginProtocols = "tcp"
	CreateDatadogPluginProtocolsTLS            CreateDatadogPluginProtocols = "tls"
	CreateDatadogPluginProtocolsTLSPassthrough CreateDatadogPluginProtocols = "tls_passthrough"
	CreateDatadogPluginProtocolsUDP            CreateDatadogPluginProtocols = "udp"
	CreateDatadogPluginProtocolsWs             CreateDatadogPluginProtocols = "ws"
	CreateDatadogPluginProtocolsWss            CreateDatadogPluginProtocols = "wss"
)

func (e CreateDatadogPluginProtocols) ToPointer() *CreateDatadogPluginProtocols {
	return &e
}
func (e *CreateDatadogPluginProtocols) UnmarshalJSON(data []byte) error {
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
		*e = CreateDatadogPluginProtocols(v)
		return nil
	default:
		return fmt.Errorf("invalid value for CreateDatadogPluginProtocols: %v", v)
	}
}

// CreateDatadogPluginConsumer - If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
type CreateDatadogPluginConsumer struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateDatadogPluginConsumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateDatadogPluginConsumerGroup struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateDatadogPluginConsumerGroup) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateDatadogPluginRoute - If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
type CreateDatadogPluginRoute struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateDatadogPluginRoute) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

// CreateDatadogPluginService - If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
type CreateDatadogPluginService struct {
	ID *string `json:"id,omitempty"`
}

func (o *CreateDatadogPluginService) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

type CreateDatadogPlugin struct {
	Config *CreateDatadogPluginConfig `json:"config,omitempty"`
	// Whether the plugin is applied.
	Enabled      *bool                        `json:"enabled,omitempty"`
	InstanceName *string                      `json:"instance_name,omitempty"`
	name         *string                      `const:"datadog" json:"name,omitempty"`
	Ordering     *CreateDatadogPluginOrdering `json:"ordering,omitempty"`
	// A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support `"tcp"` and `"tls"`.
	Protocols []CreateDatadogPluginProtocols `json:"protocols,omitempty"`
	// An optional set of strings associated with the Plugin for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.
	Consumer      *CreateDatadogPluginConsumer      `json:"consumer,omitempty"`
	ConsumerGroup *CreateDatadogPluginConsumerGroup `json:"consumer_group,omitempty"`
	// If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.
	Route *CreateDatadogPluginRoute `json:"route,omitempty"`
	// If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.
	Service *CreateDatadogPluginService `json:"service,omitempty"`
}

func (c CreateDatadogPlugin) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(c, "", false)
}

func (c *CreateDatadogPlugin) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &c, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *CreateDatadogPlugin) GetConfig() *CreateDatadogPluginConfig {
	if o == nil {
		return nil
	}
	return o.Config
}

func (o *CreateDatadogPlugin) GetEnabled() *bool {
	if o == nil {
		return nil
	}
	return o.Enabled
}

func (o *CreateDatadogPlugin) GetInstanceName() *string {
	if o == nil {
		return nil
	}
	return o.InstanceName
}

func (o *CreateDatadogPlugin) GetName() *string {
	return types.String("datadog")
}

func (o *CreateDatadogPlugin) GetOrdering() *CreateDatadogPluginOrdering {
	if o == nil {
		return nil
	}
	return o.Ordering
}

func (o *CreateDatadogPlugin) GetProtocols() []CreateDatadogPluginProtocols {
	if o == nil {
		return nil
	}
	return o.Protocols
}

func (o *CreateDatadogPlugin) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *CreateDatadogPlugin) GetConsumer() *CreateDatadogPluginConsumer {
	if o == nil {
		return nil
	}
	return o.Consumer
}

func (o *CreateDatadogPlugin) GetConsumerGroup() *CreateDatadogPluginConsumerGroup {
	if o == nil {
		return nil
	}
	return o.ConsumerGroup
}

func (o *CreateDatadogPlugin) GetRoute() *CreateDatadogPluginRoute {
	if o == nil {
		return nil
	}
	return o.Route
}

func (o *CreateDatadogPlugin) GetService() *CreateDatadogPluginService {
	if o == nil {
		return nil
	}
	return o.Service
}