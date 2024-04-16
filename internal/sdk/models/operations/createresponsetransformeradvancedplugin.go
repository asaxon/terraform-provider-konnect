// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateResponsetransformeradvancedPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Create a new ResponseTransformerAdvanced plugin
	CreateResponseTransformerAdvancedPlugin shared.CreateResponseTransformerAdvancedPlugin `request:"mediaType=application/json"`
}

func (o *CreateResponsetransformeradvancedPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateResponsetransformeradvancedPluginRequest) GetCreateResponseTransformerAdvancedPlugin() shared.CreateResponseTransformerAdvancedPlugin {
	if o == nil {
		return shared.CreateResponseTransformerAdvancedPlugin{}
	}
	return o.CreateResponseTransformerAdvancedPlugin
}

type CreateResponsetransformeradvancedPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created Plugin
	ResponseTransformerAdvancedPlugin *shared.ResponseTransformerAdvancedPlugin
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *CreateResponsetransformeradvancedPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateResponsetransformeradvancedPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateResponsetransformeradvancedPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateResponsetransformeradvancedPluginResponse) GetResponseTransformerAdvancedPlugin() *shared.ResponseTransformerAdvancedPlugin {
	if o == nil {
		return nil
	}
	return o.ResponseTransformerAdvancedPlugin
}

func (o *CreateResponsetransformeradvancedPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}