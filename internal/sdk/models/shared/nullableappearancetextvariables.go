// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type Catalog struct {
	PrimaryHeader  string `json:"primary_header"`
	WelcomeMessage string `json:"welcome_message"`
}

func (o *Catalog) GetPrimaryHeader() string {
	if o == nil {
		return ""
	}
	return o.PrimaryHeader
}

func (o *Catalog) GetWelcomeMessage() string {
	if o == nil {
		return ""
	}
	return o.WelcomeMessage
}

// NullableAppearanceTextVariables - Values to display for customizable text in the portal user interface
type NullableAppearanceTextVariables struct {
	Catalog Catalog `json:"catalog"`
}

func (o *NullableAppearanceTextVariables) GetCatalog() Catalog {
	if o == nil {
		return Catalog{}
	}
	return o.Catalog
}