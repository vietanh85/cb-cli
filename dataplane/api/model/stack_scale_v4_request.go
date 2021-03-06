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

// StackScaleV4Request stack scale v4 request
// swagger:model StackScaleV4Request
type StackScaleV4Request struct {

	// scaling adjustment of the instance groups
	// Required: true
	DesiredCount *int32 `json:"desiredCount"`

	// name of the instance group
	// Required: true
	Group *string `json:"group"`
}

// Validate validates this stack scale v4 request
func (m *StackScaleV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDesiredCount(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGroup(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackScaleV4Request) validateDesiredCount(formats strfmt.Registry) error {

	if err := validate.Required("desiredCount", "body", m.DesiredCount); err != nil {
		return err
	}

	return nil
}

func (m *StackScaleV4Request) validateGroup(formats strfmt.Registry) error {

	if err := validate.Required("group", "body", m.Group); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackScaleV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackScaleV4Request) UnmarshalBinary(b []byte) error {
	var res StackScaleV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
