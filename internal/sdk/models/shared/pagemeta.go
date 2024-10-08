// Code generated by Speakeasy (https://speakeasy.com). DO NOT EDIT.

package shared

// PageMeta - Contains pagination query parameters and the total number of objects returned.
type PageMeta struct {
	Number float64 `json:"number"`
	Size   float64 `json:"size"`
	Total  float64 `json:"total"`
}

func (o *PageMeta) GetNumber() float64 {
	if o == nil {
		return 0.0
	}
	return o.Number
}

func (o *PageMeta) GetSize() float64 {
	if o == nil {
		return 0.0
	}
	return o.Size
}

func (o *PageMeta) GetTotal() float64 {
	if o == nil {
		return 0.0
	}
	return o.Total
}
