// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// ConflictError - standard error
type ConflictError struct {
	Status   any `json:"status"`
	Title    any `json:"title"`
	Type     any `json:"type,omitempty"`
	Instance any `json:"instance"`
	Detail   any `json:"detail"`
}

func (o *ConflictError) GetStatus() any {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *ConflictError) GetTitle() any {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *ConflictError) GetType() any {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *ConflictError) GetInstance() any {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *ConflictError) GetDetail() any {
	if o == nil {
		return nil
	}
	return o.Detail
}
