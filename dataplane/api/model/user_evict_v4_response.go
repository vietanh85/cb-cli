// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// UserEvictV4Response user evict v4 response
// swagger:model UserEvictV4Response
type UserEvictV4Response struct {

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this user evict v4 response
func (m *UserEvictV4Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *UserEvictV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserEvictV4Response) UnmarshalBinary(b []byte) error {
	var res UserEvictV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
