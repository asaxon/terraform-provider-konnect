// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
)

func (r *PortalAuthDataSourceModel) RefreshFromSharedPortalAuthenticationSettingsResponse(resp *shared.PortalAuthenticationSettingsResponse) {
	if resp != nil {
		r.BasicAuthEnabled = types.BoolValue(resp.BasicAuthEnabled)
		r.KonnectMappingEnabled = types.BoolValue(resp.KonnectMappingEnabled)
		r.OidcAuthEnabled = types.BoolValue(resp.OidcAuthEnabled)
		if resp.OidcConfig == nil {
			r.OidcConfig = nil
		} else {
			r.OidcConfig = &tfTypes.PortalOIDCConfig{}
			if resp.OidcConfig.ClaimMappings == nil {
				r.OidcConfig.ClaimMappings = nil
			} else {
				r.OidcConfig.ClaimMappings = &tfTypes.PortalClaimMappings{}
				r.OidcConfig.ClaimMappings.Email = types.StringPointerValue(resp.OidcConfig.ClaimMappings.Email)
				r.OidcConfig.ClaimMappings.Groups = types.StringPointerValue(resp.OidcConfig.ClaimMappings.Groups)
				r.OidcConfig.ClaimMappings.Name = types.StringPointerValue(resp.OidcConfig.ClaimMappings.Name)
			}
			r.OidcConfig.ClientID = types.StringValue(resp.OidcConfig.ClientID)
			r.OidcConfig.Issuer = types.StringValue(resp.OidcConfig.Issuer)
			r.OidcConfig.Scopes = []types.String{}
			for _, v := range resp.OidcConfig.Scopes {
				r.OidcConfig.Scopes = append(r.OidcConfig.Scopes, types.StringValue(v))
			}
		}
		r.OidcTeamMappingEnabled = types.BoolValue(resp.OidcTeamMappingEnabled)
	}
}