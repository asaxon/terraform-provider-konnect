// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreatePrometheusPluginConfig struct {
	AiMetrics             types.Bool `tfsdk:"ai_metrics"`
	BandwidthMetrics      types.Bool `tfsdk:"bandwidth_metrics"`
	LatencyMetrics        types.Bool `tfsdk:"latency_metrics"`
	PerConsumer           types.Bool `tfsdk:"per_consumer"`
	StatusCodeMetrics     types.Bool `tfsdk:"status_code_metrics"`
	UpstreamHealthMetrics types.Bool `tfsdk:"upstream_health_metrics"`
}
