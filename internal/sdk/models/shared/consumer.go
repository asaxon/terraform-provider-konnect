// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// The Consumer object represents a consumer - or a user - of a Service. You can either rely on Kong as the primary datastore, or you can map the consumer list with your database to keep consistency between Kong and your existing primary datastore.
type Consumer struct {
	// Field for storing an existing unique ID for the Consumer - useful for mapping Kong with users in your existing database. You must send either this field or `username` with the request.
	CustomID *string `json:"custom_id,omitempty"`
	// An optional set of strings associated with the Consumer for grouping and filtering.
	Tags []string `json:"tags,omitempty"`
	// The unique username of the Consumer. You must send either this field or `custom_id` with the request.
	Username *string `json:"username,omitempty"`
	// Unix epoch when the resource was created.
	CreatedAt *int64  `json:"created_at,omitempty"`
	ID        *string `json:"id,omitempty"`
}

func (o *Consumer) GetCustomID() *string {
	if o == nil {
		return nil
	}
	return o.CustomID
}

func (o *Consumer) GetTags() []string {
	if o == nil {
		return nil
	}
	return o.Tags
}

func (o *Consumer) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}

func (o *Consumer) GetCreatedAt() *int64 {
	if o == nil {
		return nil
	}
	return o.CreatedAt
}

func (o *Consumer) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}
