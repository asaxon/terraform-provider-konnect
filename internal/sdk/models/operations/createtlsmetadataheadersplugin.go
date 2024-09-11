// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package operations

import (
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"net/http"
)

type CreateTlsmetadataheadersPluginRequest struct {
	// The UUID of your control plane. This variable is available in the Konnect manager.
	ControlPlaneID                 string                                 `pathParam:"style=simple,explode=false,name=controlPlaneId"`
	CreateTLSMetadataHeadersPlugin *shared.CreateTLSMetadataHeadersPlugin `request:"mediaType=application/json"`
}

func (o *CreateTlsmetadataheadersPluginRequest) GetControlPlaneID() string {
	if o == nil {
		return ""
	}
	return o.ControlPlaneID
}

func (o *CreateTlsmetadataheadersPluginRequest) GetCreateTLSMetadataHeadersPlugin() *shared.CreateTLSMetadataHeadersPlugin {
	if o == nil {
		return nil
	}
	return o.CreateTLSMetadataHeadersPlugin
}

type CreateTlsmetadataheadersPluginResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Created TlsMetadataHeaders plugin
	TLSMetadataHeadersPlugin *shared.TLSMetadataHeadersPlugin
	// Unauthorized
	GatewayUnauthorizedError *shared.GatewayUnauthorizedError
}

func (o *CreateTlsmetadataheadersPluginResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *CreateTlsmetadataheadersPluginResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *CreateTlsmetadataheadersPluginResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *CreateTlsmetadataheadersPluginResponse) GetTLSMetadataHeadersPlugin() *shared.TLSMetadataHeadersPlugin {
	if o == nil {
		return nil
	}
	return o.TLSMetadataHeadersPlugin
}

func (o *CreateTlsmetadataheadersPluginResponse) GetGatewayUnauthorizedError() *shared.GatewayUnauthorizedError {
	if o == nil {
		return nil
	}
	return o.GatewayUnauthorizedError
}
