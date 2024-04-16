// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package provider

import (
	"encoding/json"
	"github.com/hashicorp/terraform-plugin-framework/types"
	tfTypes "github.com/kong/terraform-provider-konnect/internal/provider/types"
	"github.com/kong/terraform-provider-konnect/internal/sdk/models/shared"
	"time"
)

func (r *ApplicationAuthStrategyResourceModel) ToSharedCreateAppAuthStrategyRequest() *shared.CreateAppAuthStrategyRequest {
	var out shared.CreateAppAuthStrategyRequest
	var appAuthStrategyKeyAuthRequest *shared.AppAuthStrategyKeyAuthRequest
	if r.KeyAuth != nil {
		name := r.KeyAuth.Name.ValueString()
		displayName := r.KeyAuth.DisplayName.ValueString()
		strategyType := shared.StrategyType(r.KeyAuth.StrategyType.ValueString())
		var keyNames []string = []string{}
		for _, keyNamesItem := range r.KeyAuth.Configs.KeyAuth.KeyNames {
			keyNames = append(keyNames, keyNamesItem.ValueString())
		}
		keyAuth := shared.AppAuthStrategyConfigKeyAuth{
			KeyNames: keyNames,
		}
		configs := shared.AppAuthStrategyKeyAuthRequestConfigs{
			KeyAuth: keyAuth,
		}
		appAuthStrategyKeyAuthRequest = &shared.AppAuthStrategyKeyAuthRequest{
			Name:         name,
			DisplayName:  displayName,
			StrategyType: strategyType,
			Configs:      configs,
		}
	}
	if appAuthStrategyKeyAuthRequest != nil {
		out = shared.CreateAppAuthStrategyRequest{
			AppAuthStrategyKeyAuthRequest: appAuthStrategyKeyAuthRequest,
		}
	}
	var appAuthStrategyOpenIDConnectRequest *shared.AppAuthStrategyOpenIDConnectRequest
	if r.OpenidConnect != nil {
		name1 := r.OpenidConnect.Name.ValueString()
		displayName1 := r.OpenidConnect.DisplayName.ValueString()
		strategyType1 := shared.AppAuthStrategyOpenIDConnectRequestStrategyType(r.OpenidConnect.StrategyType.ValueString())
		issuer := r.OpenidConnect.Configs.OpenidConnect.Issuer.ValueString()
		var credentialClaim []string = []string{}
		for _, credentialClaimItem := range r.OpenidConnect.Configs.OpenidConnect.CredentialClaim {
			credentialClaim = append(credentialClaim, credentialClaimItem.ValueString())
		}
		var scopes []string = []string{}
		for _, scopesItem := range r.OpenidConnect.Configs.OpenidConnect.Scopes {
			scopes = append(scopes, scopesItem.ValueString())
		}
		var authMethods []string = []string{}
		for _, authMethodsItem := range r.OpenidConnect.Configs.OpenidConnect.AuthMethods {
			authMethods = append(authMethods, authMethodsItem.ValueString())
		}
		var additionalProperties interface{}
		if !r.OpenidConnect.Configs.OpenidConnect.AdditionalProperties.IsUnknown() && !r.OpenidConnect.Configs.OpenidConnect.AdditionalProperties.IsNull() {
			_ = json.Unmarshal([]byte(r.OpenidConnect.Configs.OpenidConnect.AdditionalProperties.ValueString()), &additionalProperties)
		}
		openidConnect := shared.AppAuthStrategyConfigOpenIDConnect{
			Issuer:               issuer,
			CredentialClaim:      credentialClaim,
			Scopes:               scopes,
			AuthMethods:          authMethods,
			AdditionalProperties: additionalProperties,
		}
		configs1 := shared.AppAuthStrategyOpenIDConnectRequestConfigs{
			OpenidConnect: openidConnect,
		}
		dcrProviderID := new(string)
		if !r.OpenidConnect.DcrProviderID.IsUnknown() && !r.OpenidConnect.DcrProviderID.IsNull() {
			*dcrProviderID = r.OpenidConnect.DcrProviderID.ValueString()
		} else {
			dcrProviderID = nil
		}
		appAuthStrategyOpenIDConnectRequest = &shared.AppAuthStrategyOpenIDConnectRequest{
			Name:          name1,
			DisplayName:   displayName1,
			StrategyType:  strategyType1,
			Configs:       configs1,
			DcrProviderID: dcrProviderID,
		}
	}
	if appAuthStrategyOpenIDConnectRequest != nil {
		out = shared.CreateAppAuthStrategyRequest{
			AppAuthStrategyOpenIDConnectRequest: appAuthStrategyOpenIDConnectRequest,
		}
	}
	return &out
}

func (r *ApplicationAuthStrategyResourceModel) RefreshFromSharedCreateAppAuthStrategyResponse(resp *shared.CreateAppAuthStrategyResponse) {
	if resp == nil {
	} else {
		if resp.AppAuthStrategyKeyAuthResponse != nil {
			r.KeyAuth = &tfTypes.AppAuthStrategyKeyAuthRequest{}
			r.KeyAuth.Active = types.BoolValue(resp.AppAuthStrategyKeyAuthResponse.Active)
			r.Active = r.KeyAuth.Active
			r.KeyAuth.Configs.KeyAuth.KeyNames = []types.String{}
			for _, v := range resp.AppAuthStrategyKeyAuthResponse.Configs.KeyAuth.KeyNames {
				r.KeyAuth.Configs.KeyAuth.KeyNames = append(r.KeyAuth.Configs.KeyAuth.KeyNames, types.StringValue(v))
			}
			r.KeyAuth.CreatedAt = types.StringValue(resp.AppAuthStrategyKeyAuthResponse.CreatedAt.Format(time.RFC3339Nano))
			if resp.AppAuthStrategyKeyAuthResponse.DcrProvider == nil {
				r.KeyAuth.DcrProvider = nil
			} else {
				r.KeyAuth.DcrProvider = &tfTypes.DcrProvider{}
				r.KeyAuth.DcrProvider.DisplayName = types.StringPointerValue(resp.AppAuthStrategyKeyAuthResponse.DcrProvider.DisplayName)
				r.KeyAuth.DcrProvider.ID = types.StringValue(resp.AppAuthStrategyKeyAuthResponse.DcrProvider.ID)
				r.KeyAuth.DcrProvider.Name = types.StringValue(resp.AppAuthStrategyKeyAuthResponse.DcrProvider.Name)
				r.KeyAuth.DcrProvider.ProviderType = types.StringValue(string(resp.AppAuthStrategyKeyAuthResponse.DcrProvider.ProviderType))
			}
			r.KeyAuth.DisplayName = types.StringValue(resp.AppAuthStrategyKeyAuthResponse.DisplayName)
			r.DisplayName = r.KeyAuth.DisplayName
			r.KeyAuth.ID = types.StringValue(resp.AppAuthStrategyKeyAuthResponse.ID)
			r.ID = r.KeyAuth.ID
			r.KeyAuth.Name = types.StringValue(resp.AppAuthStrategyKeyAuthResponse.Name)
			r.Name = r.KeyAuth.Name
			r.KeyAuth.StrategyType = types.StringValue(string(resp.AppAuthStrategyKeyAuthResponse.StrategyType))
			r.KeyAuth.UpdatedAt = types.StringValue(resp.AppAuthStrategyKeyAuthResponse.UpdatedAt.Format(time.RFC3339Nano))
		}
		if resp.AppAuthStrategyOpenIDConnectResponse != nil {
			r.OpenidConnect = &tfTypes.AppAuthStrategyOpenIDConnectRequest{}
			r.OpenidConnect.Active = types.BoolValue(resp.AppAuthStrategyOpenIDConnectResponse.Active)
			r.Active = r.OpenidConnect.Active
			if resp.AppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.AdditionalProperties == nil {
				r.OpenidConnect.Configs.OpenidConnect.AdditionalProperties = types.StringNull()
			} else {
				additionalPropertiesResult, _ := json.Marshal(resp.AppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.AdditionalProperties)
				r.OpenidConnect.Configs.OpenidConnect.AdditionalProperties = types.StringValue(string(additionalPropertiesResult))
			}
			r.OpenidConnect.Configs.OpenidConnect.AuthMethods = []types.String{}
			for _, v := range resp.AppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.AuthMethods {
				r.OpenidConnect.Configs.OpenidConnect.AuthMethods = append(r.OpenidConnect.Configs.OpenidConnect.AuthMethods, types.StringValue(v))
			}
			r.OpenidConnect.Configs.OpenidConnect.CredentialClaim = []types.String{}
			for _, v := range resp.AppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.CredentialClaim {
				r.OpenidConnect.Configs.OpenidConnect.CredentialClaim = append(r.OpenidConnect.Configs.OpenidConnect.CredentialClaim, types.StringValue(v))
			}
			r.OpenidConnect.Configs.OpenidConnect.Issuer = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.Issuer)
			r.OpenidConnect.Configs.OpenidConnect.Scopes = []types.String{}
			for _, v := range resp.AppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.Scopes {
				r.OpenidConnect.Configs.OpenidConnect.Scopes = append(r.OpenidConnect.Configs.OpenidConnect.Scopes, types.StringValue(v))
			}
			r.OpenidConnect.CreatedAt = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponse.CreatedAt.Format(time.RFC3339Nano))
			if resp.AppAuthStrategyOpenIDConnectResponse.DcrProvider == nil {
				r.OpenidConnect.DcrProvider = nil
			} else {
				r.OpenidConnect.DcrProvider = &tfTypes.DcrProvider{}
				r.OpenidConnect.DcrProvider.DisplayName = types.StringPointerValue(resp.AppAuthStrategyOpenIDConnectResponse.DcrProvider.DisplayName)
				r.OpenidConnect.DcrProvider.ID = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponse.DcrProvider.ID)
				r.OpenidConnect.DcrProvider.Name = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponse.DcrProvider.Name)
				r.OpenidConnect.DcrProvider.ProviderType = types.StringValue(string(resp.AppAuthStrategyOpenIDConnectResponse.DcrProvider.ProviderType))
			}
			r.OpenidConnect.DisplayName = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponse.DisplayName)
			r.DisplayName = r.OpenidConnect.DisplayName
			r.OpenidConnect.ID = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponse.ID)
			r.ID = r.OpenidConnect.ID
			r.OpenidConnect.Name = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponse.Name)
			r.Name = r.OpenidConnect.Name
			r.OpenidConnect.StrategyType = types.StringValue(string(resp.AppAuthStrategyOpenIDConnectResponse.StrategyType))
			r.OpenidConnect.UpdatedAt = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponse.UpdatedAt.Format(time.RFC3339Nano))
		}
	}
}

func (r *ApplicationAuthStrategyResourceModel) RefreshFromSharedGetAppAuthStrategyResponse(resp *shared.GetAppAuthStrategyResponse) {
	if resp == nil {
	} else {
		if resp.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse != nil {
			r.KeyAuth = &tfTypes.AppAuthStrategyKeyAuthRequest{}
			r.KeyAuth.Active = types.BoolValue(resp.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse.Active)
			r.Active = r.KeyAuth.Active
			r.KeyAuth.Configs.KeyAuth.KeyNames = []types.String{}
			for _, v := range resp.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse.Configs.KeyAuth.KeyNames {
				r.KeyAuth.Configs.KeyAuth.KeyNames = append(r.KeyAuth.Configs.KeyAuth.KeyNames, types.StringValue(v))
			}
			r.KeyAuth.CreatedAt = types.StringValue(resp.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse.CreatedAt.Format(time.RFC3339Nano))
			if resp.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse.DcrProvider == nil {
				r.KeyAuth.DcrProvider = nil
			} else {
				r.KeyAuth.DcrProvider = &tfTypes.DcrProvider{}
				r.KeyAuth.DcrProvider.DisplayName = types.StringPointerValue(resp.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse.DcrProvider.DisplayName)
				r.KeyAuth.DcrProvider.ID = types.StringValue(resp.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse.DcrProvider.ID)
				r.KeyAuth.DcrProvider.Name = types.StringValue(resp.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse.DcrProvider.Name)
				r.KeyAuth.DcrProvider.ProviderType = types.StringValue(string(resp.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse.DcrProvider.ProviderType))
			}
			r.KeyAuth.DisplayName = types.StringValue(resp.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse.DisplayName)
			r.DisplayName = r.KeyAuth.DisplayName
			r.KeyAuth.ID = types.StringValue(resp.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse.ID)
			r.ID = r.KeyAuth.ID
			r.KeyAuth.Name = types.StringValue(resp.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse.Name)
			r.Name = r.KeyAuth.Name
			r.KeyAuth.StrategyType = types.StringValue(string(resp.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse.StrategyType))
			r.KeyAuth.UpdatedAt = types.StringValue(resp.AppAuthStrategyKeyAuthResponseAppAuthStrategyKeyAuthResponse.UpdatedAt.Format(time.RFC3339Nano))
		}
		if resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse != nil {
			r.OpenidConnect = &tfTypes.AppAuthStrategyOpenIDConnectRequest{}
			r.OpenidConnect.Active = types.BoolValue(resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.Active)
			r.Active = r.OpenidConnect.Active
			if resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.AdditionalProperties == nil {
				r.OpenidConnect.Configs.OpenidConnect.AdditionalProperties = types.StringNull()
			} else {
				additionalPropertiesResult, _ := json.Marshal(resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.AdditionalProperties)
				r.OpenidConnect.Configs.OpenidConnect.AdditionalProperties = types.StringValue(string(additionalPropertiesResult))
			}
			r.OpenidConnect.Configs.OpenidConnect.AuthMethods = []types.String{}
			for _, v := range resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.AuthMethods {
				r.OpenidConnect.Configs.OpenidConnect.AuthMethods = append(r.OpenidConnect.Configs.OpenidConnect.AuthMethods, types.StringValue(v))
			}
			r.OpenidConnect.Configs.OpenidConnect.CredentialClaim = []types.String{}
			for _, v := range resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.CredentialClaim {
				r.OpenidConnect.Configs.OpenidConnect.CredentialClaim = append(r.OpenidConnect.Configs.OpenidConnect.CredentialClaim, types.StringValue(v))
			}
			r.OpenidConnect.Configs.OpenidConnect.Issuer = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.Issuer)
			r.OpenidConnect.Configs.OpenidConnect.Scopes = []types.String{}
			for _, v := range resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.Scopes {
				r.OpenidConnect.Configs.OpenidConnect.Scopes = append(r.OpenidConnect.Configs.OpenidConnect.Scopes, types.StringValue(v))
			}
			r.OpenidConnect.CreatedAt = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.CreatedAt.Format(time.RFC3339Nano))
			if resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.DcrProvider == nil {
				r.OpenidConnect.DcrProvider = nil
			} else {
				r.OpenidConnect.DcrProvider = &tfTypes.DcrProvider{}
				r.OpenidConnect.DcrProvider.DisplayName = types.StringPointerValue(resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.DcrProvider.DisplayName)
				r.OpenidConnect.DcrProvider.ID = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.DcrProvider.ID)
				r.OpenidConnect.DcrProvider.Name = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.DcrProvider.Name)
				r.OpenidConnect.DcrProvider.ProviderType = types.StringValue(string(resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.DcrProvider.ProviderType))
			}
			r.OpenidConnect.DisplayName = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.DisplayName)
			r.DisplayName = r.OpenidConnect.DisplayName
			r.OpenidConnect.ID = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.ID)
			r.ID = r.OpenidConnect.ID
			r.OpenidConnect.Name = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.Name)
			r.Name = r.OpenidConnect.Name
			r.OpenidConnect.StrategyType = types.StringValue(string(resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.StrategyType))
			r.OpenidConnect.UpdatedAt = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseAppAuthStrategyOpenIDConnectResponse.UpdatedAt.Format(time.RFC3339Nano))
		}
	}
}

func (r *ApplicationAuthStrategyResourceModel) ToSharedUpdateAppAuthStrategyRequest() *shared.UpdateAppAuthStrategyRequest {
	name := new(string)
	if !r.Name.IsUnknown() && !r.Name.IsNull() {
		*name = r.Name.ValueString()
	} else {
		name = nil
	}
	displayName := new(string)
	if !r.DisplayName.IsUnknown() && !r.DisplayName.IsNull() {
		*displayName = r.DisplayName.ValueString()
	} else {
		displayName = nil
	}
	out := shared.UpdateAppAuthStrategyRequest{
		Name:        name,
		DisplayName: displayName,
	}
	return &out
}

func (r *ApplicationAuthStrategyResourceModel) RefreshFromSharedUpdateAppAuthStrategyResponse(resp *shared.UpdateAppAuthStrategyResponse) {
	if resp == nil {
	} else {
		if resp.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse != nil {
			r.KeyAuth = &tfTypes.AppAuthStrategyKeyAuthRequest{}
			r.KeyAuth.Active = types.BoolValue(resp.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse.Active)
			r.Active = r.KeyAuth.Active
			r.KeyAuth.Configs.KeyAuth.KeyNames = []types.String{}
			for _, v := range resp.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse.Configs.KeyAuth.KeyNames {
				r.KeyAuth.Configs.KeyAuth.KeyNames = append(r.KeyAuth.Configs.KeyAuth.KeyNames, types.StringValue(v))
			}
			r.KeyAuth.CreatedAt = types.StringValue(resp.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse.CreatedAt.Format(time.RFC3339Nano))
			if resp.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse.DcrProvider == nil {
				r.KeyAuth.DcrProvider = nil
			} else {
				r.KeyAuth.DcrProvider = &tfTypes.DcrProvider{}
				r.KeyAuth.DcrProvider.DisplayName = types.StringPointerValue(resp.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse.DcrProvider.DisplayName)
				r.KeyAuth.DcrProvider.ID = types.StringValue(resp.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse.DcrProvider.ID)
				r.KeyAuth.DcrProvider.Name = types.StringValue(resp.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse.DcrProvider.Name)
				r.KeyAuth.DcrProvider.ProviderType = types.StringValue(string(resp.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse.DcrProvider.ProviderType))
			}
			r.KeyAuth.DisplayName = types.StringValue(resp.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse.DisplayName)
			r.DisplayName = r.KeyAuth.DisplayName
			r.KeyAuth.ID = types.StringValue(resp.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse.ID)
			r.ID = r.KeyAuth.ID
			r.KeyAuth.Name = types.StringValue(resp.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse.Name)
			r.Name = r.KeyAuth.Name
			r.KeyAuth.StrategyType = types.StringValue(string(resp.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse.StrategyType))
			r.KeyAuth.UpdatedAt = types.StringValue(resp.AppAuthStrategyKeyAuthResponseUpdateAppAuthStrategyResponseAppAuthStrategyKeyAuthResponse.UpdatedAt.Format(time.RFC3339Nano))
		}
		if resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse != nil {
			r.OpenidConnect = &tfTypes.AppAuthStrategyOpenIDConnectRequest{}
			r.OpenidConnect.Active = types.BoolValue(resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.Active)
			r.Active = r.OpenidConnect.Active
			if resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.AdditionalProperties == nil {
				r.OpenidConnect.Configs.OpenidConnect.AdditionalProperties = types.StringNull()
			} else {
				additionalPropertiesResult, _ := json.Marshal(resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.AdditionalProperties)
				r.OpenidConnect.Configs.OpenidConnect.AdditionalProperties = types.StringValue(string(additionalPropertiesResult))
			}
			r.OpenidConnect.Configs.OpenidConnect.AuthMethods = []types.String{}
			for _, v := range resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.AuthMethods {
				r.OpenidConnect.Configs.OpenidConnect.AuthMethods = append(r.OpenidConnect.Configs.OpenidConnect.AuthMethods, types.StringValue(v))
			}
			r.OpenidConnect.Configs.OpenidConnect.CredentialClaim = []types.String{}
			for _, v := range resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.CredentialClaim {
				r.OpenidConnect.Configs.OpenidConnect.CredentialClaim = append(r.OpenidConnect.Configs.OpenidConnect.CredentialClaim, types.StringValue(v))
			}
			r.OpenidConnect.Configs.OpenidConnect.Issuer = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.Issuer)
			r.OpenidConnect.Configs.OpenidConnect.Scopes = []types.String{}
			for _, v := range resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.Configs.OpenidConnect.Scopes {
				r.OpenidConnect.Configs.OpenidConnect.Scopes = append(r.OpenidConnect.Configs.OpenidConnect.Scopes, types.StringValue(v))
			}
			r.OpenidConnect.CreatedAt = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.CreatedAt.Format(time.RFC3339Nano))
			if resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.DcrProvider == nil {
				r.OpenidConnect.DcrProvider = nil
			} else {
				r.OpenidConnect.DcrProvider = &tfTypes.DcrProvider{}
				r.OpenidConnect.DcrProvider.DisplayName = types.StringPointerValue(resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.DcrProvider.DisplayName)
				r.OpenidConnect.DcrProvider.ID = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.DcrProvider.ID)
				r.OpenidConnect.DcrProvider.Name = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.DcrProvider.Name)
				r.OpenidConnect.DcrProvider.ProviderType = types.StringValue(string(resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.DcrProvider.ProviderType))
			}
			r.OpenidConnect.DisplayName = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.DisplayName)
			r.DisplayName = r.OpenidConnect.DisplayName
			r.OpenidConnect.ID = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.ID)
			r.ID = r.OpenidConnect.ID
			r.OpenidConnect.Name = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.Name)
			r.Name = r.OpenidConnect.Name
			r.OpenidConnect.StrategyType = types.StringValue(string(resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.StrategyType))
			r.OpenidConnect.UpdatedAt = types.StringValue(resp.AppAuthStrategyOpenIDConnectResponseUpdateAppAuthStrategyResponseAppAuthStrategyOpenIDConnectResponse.UpdatedAt.Format(time.RFC3339Nano))
		}
	}
}