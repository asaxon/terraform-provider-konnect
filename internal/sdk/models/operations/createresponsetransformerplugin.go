// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateResponsetransformerPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID string `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	// Create a new ResponseTransformer plugin
	CreateResponseTransformerPlugin shared.CreateResponseTransformerPlugin `request:"mediaType=application/json"`
}

func (o *CreateResponsetransformerPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateResponsetransformerPluginRequest) GetCreateResponseTransformerPlugin() shared.CreateResponseTransformerPlugin {
	if o == nil {
		return shared.CreateResponseTransformerPlugin{}
	}
	return o.CreateResponseTransformerPlugin
}

type CreateResponsetransformerPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Successfully created Plugin
	ResponseTransformerPlugin *shared.ResponseTransformerPlugin
	// Unauthorized
	UnauthorizedError *shared.UnauthorizedError
}

func (o *CreateResponsetransformerPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateResponsetransformerPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateResponsetransformerPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateResponsetransformerPluginResponse) GetResponseTransformerPlugin() *shared.ResponseTransformerPlugin {
	if o == nil {
		return nil
	}
	return o.ResponseTransformerPlugin
}

func (o *CreateResponsetransformerPluginResponse) GetUnauthorizedError() *shared.UnauthorizedError {
	if o == nil {
		return nil
	}
	return o.UnauthorizedError
}