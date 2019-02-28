// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PlacementSettingsV4Request placement settings v4 request
// swagger:model PlacementSettingsV4Request
type PlacementSettingsV4Request struct {

	// availability zone of the stack
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// region of the stack
	// Required: true
	Region *string `json:"region"`
}

// Validate validates this placement settings v4 request
func (m *PlacementSettingsV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlacementSettingsV4Request) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlacementSettingsV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlacementSettingsV4Request) UnmarshalBinary(b []byte) error {
	var res PlacementSettingsV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}