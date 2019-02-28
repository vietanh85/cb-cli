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

// UserContext user context
// swagger:model UserContext

type UserContext struct {

	// permissions
	// Required: true
	Permissions []string `json:"permissions"`

	// strategy id
	// Required: true
	StrategyID *strfmt.UUID `json:"strategy_id"`

	// tenant id
	// Required: true
	TenantID *strfmt.UUID `json:"tenant_id"`

	// tenant name
	// Required: true
	TenantName *string `json:"tenant_name"`

	// user id
	// Required: true
	UserID *strfmt.UUID `json:"user_id"`

	// username
	// Required: true
	Username *string `json:"username"`
}

/* polymorph UserContext permissions false */

/* polymorph UserContext strategy_id false */

/* polymorph UserContext tenant_id false */

/* polymorph UserContext tenant_name false */

/* polymorph UserContext user_id false */

/* polymorph UserContext username false */

// Validate validates this user context
func (m *UserContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePermissions(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStrategyID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTenantID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTenantName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUsername(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UserContext) validatePermissions(formats strfmt.Registry) error {

	if err := validate.Required("permissions", "body", m.Permissions); err != nil {
		return err
	}

	return nil
}

func (m *UserContext) validateStrategyID(formats strfmt.Registry) error {

	if err := validate.Required("strategy_id", "body", m.StrategyID); err != nil {
		return err
	}

	if err := validate.FormatOf("strategy_id", "body", "uuid", m.StrategyID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserContext) validateTenantID(formats strfmt.Registry) error {

	if err := validate.Required("tenant_id", "body", m.TenantID); err != nil {
		return err
	}

	if err := validate.FormatOf("tenant_id", "body", "uuid", m.TenantID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserContext) validateTenantName(formats strfmt.Registry) error {

	if err := validate.Required("tenant_name", "body", m.TenantName); err != nil {
		return err
	}

	return nil
}

func (m *UserContext) validateUserID(formats strfmt.Registry) error {

	if err := validate.Required("user_id", "body", m.UserID); err != nil {
		return err
	}

	if err := validate.FormatOf("user_id", "body", "uuid", m.UserID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UserContext) validateUsername(formats strfmt.Registry) error {

	if err := validate.Required("username", "body", m.Username); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UserContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UserContext) UnmarshalBinary(b []byte) error {
	var res UserContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}