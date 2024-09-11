// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type UpdateCanaryPluginRequest struct {
	// ID of the Plugin to lookup
	PluginID string `pathParam:"style=simple,explode=false,name=PluginId"`
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID     string                     `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateCanaryPlugin *shared.CreateCanaryPlugin `request:"mediaType=application/json"`
}

func (o *UpdateCanaryPluginRequest) GetPluginID() string {
	if o == nil {
		return ""
	}
	return o.PluginID
}

func (o *UpdateCanaryPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *UpdateCanaryPluginRequest) GetCreateCanaryPlugin() *shared.CreateCanaryPlugin {
	if o == nil {
		return nil
	}
	return o.CreateCanaryPlugin
}

type UpdateCanaryPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Canary plugin
	CanaryPlugin *shared.CanaryPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *UpdateCanaryPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateCanaryPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateCanaryPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdateCanaryPluginResponse) GetCanaryPlugin() *shared.CanaryPlugin {
	if o == nil {
		return nil
	}
	return o.CanaryPlugin
}

func (o *UpdateCanaryPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
