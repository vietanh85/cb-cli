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

// AppBased app based
// swagger:model AppBased
type AppBased struct {

	// access key
	// Required: true
	AccessKey *string `json:"accessKey"`

	// secret key
	// Required: true
	SecretKey *string `json:"secretKey"`
}

// Validate validates this app based
func (m *AppBased) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccessKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AppBased) validateAccessKey(formats strfmt.Registry) error {

	if err := validate.Required("accessKey", "body", m.AccessKey); err != nil {
		return err
	}

	return nil
}

func (m *AppBased) validateSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("secretKey", "body", m.SecretKey); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AppBased) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AppBased) UnmarshalBinary(b []byte) error {
	var res AppBased
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
