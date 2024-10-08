// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// UnsupportedMediaTypeError - standard error
type UnsupportedMediaTypeError struct {
	Status   any `json:"status"`
	Title    any `json:"title"`
	Type     any `json:"type,omitempty"`
	Instance any `json:"instance"`
	Detail   any `json:"detail"`
}

func (o *UnsupportedMediaTypeError) GetStatus() any {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *UnsupportedMediaTypeError) GetTitle() any {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *UnsupportedMediaTypeError) GetType() any {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *UnsupportedMediaTypeError) GetInstance() any {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *UnsupportedMediaTypeError) GetDetail() any {
	if o == nil {
		return nil
	}
	return o.Detail
}
