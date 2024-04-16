// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// BadRequestError - standard error
type BadRequestError struct {
	Status   interface{} `json:"status"`
	Title    interface{} `json:"title"`
	Type     interface{} `json:"type,omitempty"`
	Instance interface{} `json:"instance"`
	Detail   interface{} `json:"detail"`
	// invalid parameters
	InvalidParameters []InvalidParameters `json:"invalid_parameters"`
}

func (o *BadRequestError) GetStatus() interface{} {
	if o == nil {
		return nil
	}
	return o.Status
}

func (o *BadRequestError) GetTitle() interface{} {
	if o == nil {
		return nil
	}
	return o.Title
}

func (o *BadRequestError) GetType() interface{} {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *BadRequestError) GetInstance() interface{} {
	if o == nil {
		return nil
	}
	return o.Instance
}

func (o *BadRequestError) GetDetail() interface{} {
	if o == nil {
		return nil
	}
	return o.Detail
}

func (o *BadRequestError) GetInvalidParameters() []InvalidParameters {
	if o == nil {
		return []InvalidParameters{}
	}
	return o.InvalidParameters
}