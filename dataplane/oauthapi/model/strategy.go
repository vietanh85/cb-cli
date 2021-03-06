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

// Strategy strategy
// swagger:model Strategy

type Strategy struct {

	// config
	// Required: true
	Config interface{} `json:"config"`

	// enabled
	// Required: true
	Enabled *bool `json:"enabled"`

	// id
	ID strfmt.UUID `json:"id,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// strategy type id
	// Required: true
	StrategyTypeID *strfmt.UUID `json:"strategy_type_id"`

	// tenant id
	// Required: true
	TenantID *strfmt.UUID `json:"tenant_id"`
}

/* polymorph Strategy config false */

/* polymorph Strategy enabled false */

/* polymorph Strategy id false */

/* polymorph Strategy name false */

/* polymorph Strategy strategy_type_id false */

/* polymorph Strategy tenant_id false */

// Validate validates this strategy
func (m *Strategy) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConfig(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEnabled(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStrategyTypeID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Strategy) validateConfig(formats strfmt.Registry) error {

	return nil
}

func (m *Strategy) validateEnabled(formats strfmt.Registry) error {

	if err := validate.Required("enabled", "body", m.Enabled); err != nil {
		return err
	}

	return nil
}

func (m *Strategy) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Strategy) validateStrategyTypeID(formats strfmt.Registry) error {

	if err := validate.Required("strategy_type_id", "body", m.StrategyTypeID); err != nil {
		return err
	}

	if err := validate.FormatOf("strategy_type_id", "body", "uuid", m.StrategyTypeID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Strategy) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenant_id", "body", m.TenantID); err != nil {
		return err
	}

	if err := validate.FormatOf("tenant_id", "body", "uuid", m.TenantID.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Strategy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Strategy) UnmarshalBinary(b []byte) error {
	var res Strategy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
