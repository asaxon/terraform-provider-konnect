// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// CreateTransitGatewayRequest - Request schema for creating a transit gateway.
type CreateTransitGatewayRequest struct {
	// Human-readable name of the transit gateway.
	Name string `json:"name"`
	// CIDR blocks for constructing a route table for the transit gateway, when attaching to the owning
	// network.
	//
	CidrBlocks                     []string                       `json:"cidr_blocks"`
	TransitGatewayAttachmentConfig TransitGatewayAttachmentConfig `json:"transit_gateway_attachment_config"`
	// List of mappings from remote DNS server IP address sets to proxied internal domains, for a transit gateway
	// attachment.
	//
	DNSConfig []TransitGatewayDNSConfig `json:"dns_config,omitempty"`
}

func (o *CreateTransitGatewayRequest) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateTransitGatewayRequest) GetCidrBlocks() []string {
	if o == nil {
		return []string{}
	}
	return o.CidrBlocks
}

func (o *CreateTransitGatewayRequest) GetTransitGatewayAttachmentConfig() TransitGatewayAttachmentConfig {
	if o == nil {
		return TransitGatewayAttachmentConfig{}
	}
	return o.TransitGatewayAttachmentConfig
}

func (o *CreateTransitGatewayRequest) GetDNSConfig() []TransitGatewayDNSConfig {
	if o == nil {
		return nil
	}
	return o.DNSConfig
}