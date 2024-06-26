// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateAiprompttemplatePluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID               string                               `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateAIPromptTemplatePlugin *shared.CreateAIPromptTemplatePlugin `request:"mediaType=application/json"`
}

func (o *CreateAiprompttemplatePluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateAiprompttemplatePluginRequest) GetCreateAIPromptTemplatePlugin() *shared.CreateAIPromptTemplatePlugin {
	if o == nil {
		return nil
	}
	return o.CreateAIPromptTemplatePlugin
}

type CreateAiprompttemplatePluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created AIPromptTemplate plugin
	AIPromptTemplatePlugin *shared.AIPromptTemplatePlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateAiprompttemplatePluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateAiprompttemplatePluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateAiprompttemplatePluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateAiprompttemplatePluginResponse) GetAIPromptTemplatePlugin() *shared.AIPromptTemplatePlugin {
	if o == nil {
		return nil
	}
	return o.AIPromptTemplatePlugin
}

func (o *CreateAiprompttemplatePluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
