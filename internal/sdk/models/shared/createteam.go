// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

type CreateTeam struct {
	// A name for the team being created.
	Name string `json:"name"`
	// The description of the new team.
	Description *string `json:"description,omitempty"`
	// Labels store metadata of an entity that can be used for filtering an entity list or for searching across entity types.
	//
	// Keys must be of length 1-63 characters, and cannot start with "kong", "konnect", "mesh", "kic", or "_".
	//
	Labels map[string]string `json:"labels,omitempty"`
}

func (o *CreateTeam) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *CreateTeam) GetDescription() *string {
	if o == nil {
		return nil
	}
	return o.Description
}

func (o *CreateTeam) GetLabels() map[string]string {
	if o == nil {
		return nil
	}
	return o.Labels
}
