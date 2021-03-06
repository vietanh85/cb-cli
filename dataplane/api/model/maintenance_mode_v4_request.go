// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MaintenanceModeV4Request maintenance mode v4 request
// swagger:model MaintenanceModeV4Request
type MaintenanceModeV4Request struct {

	// maintenance mode status
	// Enum: [ENABLED VALIDATION_REQUESTED DISABLED]
	Status string `json:"status,omitempty"`
}

// Validate validates this maintenance mode v4 request
func (m *MaintenanceModeV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var maintenanceModeV4RequestTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ENABLED","VALIDATION_REQUESTED","DISABLED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		maintenanceModeV4RequestTypeStatusPropEnum = append(maintenanceModeV4RequestTypeStatusPropEnum, v)
	}
}

const (

	// MaintenanceModeV4RequestStatusENABLED captures enum value "ENABLED"
	MaintenanceModeV4RequestStatusENABLED string = "ENABLED"

	// MaintenanceModeV4RequestStatusVALIDATIONREQUESTED captures enum value "VALIDATION_REQUESTED"
	MaintenanceModeV4RequestStatusVALIDATIONREQUESTED string = "VALIDATION_REQUESTED"

	// MaintenanceModeV4RequestStatusDISABLED captures enum value "DISABLED"
	MaintenanceModeV4RequestStatusDISABLED string = "DISABLED"
)

// prop value enum
func (m *MaintenanceModeV4Request) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, maintenanceModeV4RequestTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *MaintenanceModeV4Request) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MaintenanceModeV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MaintenanceModeV4Request) UnmarshalBinary(b []byte) error {
	var res MaintenanceModeV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
