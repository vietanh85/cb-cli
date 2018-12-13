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

// PlacementSettings placement settings
// swagger:model PlacementSettings
type PlacementSettings struct {

	// availability zone of the stack
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// region of the stack
	// Required: true
	Region *string `json:"region"`
}

// Validate validates this placement settings
func (m *PlacementSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PlacementSettings) validateRegion(formats strfmt.Registry) error {

	if err := validate.Required("region", "body", m.Region); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PlacementSettings) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PlacementSettings) UnmarshalBinary(b []byte) error {
	var res PlacementSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}