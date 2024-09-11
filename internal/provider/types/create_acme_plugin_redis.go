// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package types

import "github.com/hashicorp/terraform-plugin-framework/types"

type CreateAcmePluginRedis struct {
	Database     types.Int64                   `tfsdk:"database"`
	ExtraOptions *CreateAcmePluginExtraOptions `tfsdk:"extra_options"`
	Host         types.String                  `tfsdk:"host"`
	Password     types.String                  `tfsdk:"password"`
	Port         types.Int64                   `tfsdk:"port"`
	ServerName   types.String                  `tfsdk:"server_name"`
	Ssl          types.Bool                    `tfsdk:"ssl"`
	SslVerify    types.Bool                    `tfsdk:"ssl_verify"`
	Timeout      types.Int64                   `tfsdk:"timeout"`
	Username     types.String                  `tfsdk:"username"`
}
