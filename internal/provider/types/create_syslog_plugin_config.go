// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateSyslogPluginConfig struct {
	ClientErrorsSeverity types.String            `tfsdk:"client_errors_severity"`
	CustomFieldsByLua    map[string]types.String `tfsdk:"custom_fields_by_lua"`
	Facility             types.String            `tfsdk:"facility"`
	LogLevel             types.String            `tfsdk:"log_level"`
	ServerErrorsSeverity types.String            `tfsdk:"server_errors_severity"`
	SuccessfulSeverity   types.String            `tfsdk:"successful_severity"`
}