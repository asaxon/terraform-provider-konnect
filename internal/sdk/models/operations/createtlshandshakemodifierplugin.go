// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateTlshandshakemodifierPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID                   string                                   `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateTLSHandshakeModifierPlugin *shared.CreateTLSHandshakeModifierPlugin `request:"mediaType=application/json"`
}

func (o *CreateTlshandshakemodifierPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateTlshandshakemodifierPluginRequest) GetCreateTLSHandshakeModifierPlugin() *shared.CreateTLSHandshakeModifierPlugin {
	if o == nil {
		return nil
	}
	return o.CreateTLSHandshakeModifierPlugin
}

type CreateTlshandshakemodifierPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created TlsHandshakeModifier plugin
	TLSHandshakeModifierPlugin *shared.TLSHandshakeModifierPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateTlshandshakemodifierPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTlshandshakemodifierPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTlshandshakemodifierPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTlshandshakemodifierPluginResponse) GetTLSHandshakeModifierPlugin() *shared.TLSHandshakeModifierPlugin {
	if o == nil {
		return nil
	}
	return o.TLSHandshakeModifierPlugin
}

func (o *CreateTlshandshakemodifierPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
