// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/operations"
)

// Ensure provider defined types fully satisfy framework interfaces.
var _ datasource.DataSource = &GatewayPluginCorsDataSource{}
var _ datasource.DataSourceWithConfigure = &GatewayPluginCorsDataSource{}

func NewGatewayPluginCorsDataSource() datasource.DataSource {
	return &GatewayPluginCorsDataSource{}
}

// GatewayPluginCorsDataSource is the data source implementation.
type GatewayPluginCorsDataSource struct {
	client *sdk.Konnect
}

// GatewayPluginCorsDataSourceModel describes the data model.
type GatewayPluginCorsDataSourceModel struct {
	Config         *tfTypes.CreateCorsPluginConfig  `tfsdk:"config"`
	Consumer       *tfTypes.ACLConsumer             `tfsdk:"consumer"`
	ConsumerGroup  *tfTypes.ACLConsumer             `tfsdk:"consumer_group"`
	ControlPlaneID types.String                     `tfsdk:"control_plane_id"`
	CreatedAt      types.Int64                      `tfsdk:"created_at"`
	Enabled        types.Bool                       `tfsdk:"enabled"`
	ID             types.String                     `tfsdk:"id"`
	InstanceName   types.String                     `tfsdk:"instance_name"`
	Ordering       *tfTypes.CreateACLPluginOrdering `tfsdk:"ordering"`
	Protocols      []types.String                   `tfsdk:"protocols"`
	Route          *tfTypes.ACLConsumer             `tfsdk:"route"`
	Service        *tfTypes.ACLConsumer             `tfsdk:"service"`
	Tags           []types.String                   `tfsdk:"tags"`
	UpdatedAt      types.Int64                      `tfsdk:"updated_at"`
}

// Metadata returns the data source type name.
func (r *GatewayPluginCorsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_gateway_plugin_cors"
}

// Schema defines the schema for the data source.
func (r *GatewayPluginCorsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		MarkdownDescription: "GatewayPluginCors DataSource",

		Attributes: map[string]schema.Attribute{
			"config": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"credentials": schema.BoolAttribute{
						Computed:    true,
						Description: `Flag to determine whether the ` + "`" + `Access-Control-Allow-Credentials` + "`" + ` header should be sent with ` + "`" + `true` + "`" + ` as the value.`,
					},
					"exposed_headers": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Value for the ` + "`" + `Access-Control-Expose-Headers` + "`" + ` header. If not specified, no custom headers are exposed.`,
					},
					"headers": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `Value for the ` + "`" + `Access-Control-Allow-Headers` + "`" + ` header.`,
					},
					"max_age": schema.NumberAttribute{
						Computed:    true,
						Description: `Indicates how long the results of the preflight request can be cached, in ` + "`" + `seconds` + "`" + `.`,
					},
					"methods": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `'Value for the ` + "`" + `Access-Control-Allow-Methods` + "`" + ` header. Available options include ` + "`" + `GET` + "`" + `, ` + "`" + `HEAD` + "`" + `, ` + "`" + `PUT` + "`" + `, ` + "`" + `PATCH` + "`" + `, ` + "`" + `POST` + "`" + `, ` + "`" + `DELETE` + "`" + `, ` + "`" + `OPTIONS` + "`" + `, ` + "`" + `TRACE` + "`" + `, ` + "`" + `CONNECT` + "`" + `. By default, all options are allowed.'`,
					},
					"origins": schema.ListAttribute{
						Computed:    true,
						ElementType: types.StringType,
						Description: `List of allowed domains for the ` + "`" + `Access-Control-Allow-Origin` + "`" + ` header. If you want to allow all origins, add ` + "`" + `*` + "`" + ` as a single value to this configuration field. The accepted values can either be flat strings or PCRE regexes.`,
					},
					"preflight_continue": schema.BoolAttribute{
						Computed:    true,
						Description: `A boolean value that instructs the plugin to proxy the ` + "`" + `OPTIONS` + "`" + ` preflight request to the Upstream service.`,
					},
					"private_network": schema.BoolAttribute{
						Computed:    true,
						Description: `Flag to determine whether the ` + "`" + `Access-Control-Allow-Private-Network` + "`" + ` header should be sent with ` + "`" + `true` + "`" + ` as the value.`,
					},
				},
			},
			"consumer": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will activate only for requests where the specified has been authenticated. (Note that some plugins can not be restricted to consumers this way.). Leave unset for the plugin to activate regardless of the authenticated Consumer.`,
			},
			"consumer_group": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
			},
			"control_plane_id": schema.StringAttribute{
				Required:    true,
				Description: `The UUID of your control plane. This variable is available in the Konnect manager. Requires replacement if changed. `,
			},
			"created_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was created.`,
			},
			"enabled": schema.BoolAttribute{
				Computed:    true,
				Description: `Whether the plugin is applied.`,
			},
			"id": schema.StringAttribute{
				Computed: true,
			},
			"instance_name": schema.StringAttribute{
				Computed: true,
			},
			"ordering": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"after": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
					"before": schema.SingleNestedAttribute{
						Computed: true,
						Attributes: map[string]schema.Attribute{
							"access": schema.ListAttribute{
								Computed:    true,
								ElementType: types.StringType,
							},
						},
					},
				},
			},
			"protocols": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `A list of the request protocols that will trigger this plugin. The default value, as well as the possible values allowed on this field, may change depending on the plugin type. For example, plugins that only work in stream mode will only support ` + "`" + `"tcp"` + "`" + ` and ` + "`" + `"tls"` + "`" + `.`,
			},
			"route": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via the specified route. Leave unset for the plugin to activate regardless of the Route being used.`,
			},
			"service": schema.SingleNestedAttribute{
				Computed: true,
				Attributes: map[string]schema.Attribute{
					"id": schema.StringAttribute{
						Computed: true,
					},
				},
				Description: `If set, the plugin will only activate when receiving requests via one of the routes belonging to the specified Service. Leave unset for the plugin to activate regardless of the Service being matched.`,
			},
			"tags": schema.ListAttribute{
				Computed:    true,
				ElementType: types.StringType,
				Description: `An optional set of strings associated with the Plugin for grouping and filtering.`,
			},
			"updated_at": schema.Int64Attribute{
				Computed:    true,
				Description: `Unix epoch when the resource was last updated.`,
			},
		},
	}
}

func (r *GatewayPluginCorsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	// Prevent panic if the provider has not been configured.
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*sdk.Konnect)

	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected DataSource Configure Type",
			fmt.Sprintf("Expected *sdk.Konnect, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	r.client = client
}

func (r *GatewayPluginCorsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data *GatewayPluginCorsDataSourceModel
	var item types.Object

	resp.Diagnostics.Append(req.Config.Get(ctx, &item)...)
	if resp.Diagnostics.HasError() {
		return
	}

	resp.Diagnostics.Append(item.As(ctx, &data, basetypes.ObjectAsOptions{
		UnhandledNullAsEmpty:    true,
		UnhandledUnknownAsEmpty: true,
	})...)

	if resp.Diagnostics.HasError() {
		return
	}

	var pluginID string
	pluginID = data.ID.ValueString()

	var controlPlaneID string
	controlPlaneID = data.ControlPlaneID.ValueString()

	request := operations.GetCorsPluginRequest{
		PluginID:       pluginID,
		ControlPlaneID: controlPlaneID,
	}
	res, err := r.client.Plugins.GetCorsPlugin(ctx, request)
	if err != nil {
		resp.Diagnostics.AddError("failure to invoke API", err.Error())
		if res != nil && res.RawResponse != nil {
			resp.Diagnostics.AddError("unexpected http request/response", debugResponse(res.RawResponse))
		}
		return
	}
	if res == nil {
		resp.Diagnostics.AddError("unexpected response from API", fmt.Sprintf("%v", res))
		return
	}
	if res.StatusCode == 404 {
		resp.State.RemoveResource(ctx)
		return
	}
	if res.StatusCode != 200 {
		resp.Diagnostics.AddError(fmt.Sprintf("unexpected response from API. Got an unexpected response code %v", res.StatusCode), debugResponse(res.RawResponse))
		return
	}
	if !(res.CorsPlugin != nil) {
		resp.Diagnostics.AddError("unexpected response from API. Got an unexpected response body", debugResponse(res.RawResponse))
		return
	}
	data.RefreshFromSharedCorsPlugin(res.CorsPlugin)

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
