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

// AdlsGen2CloudStorageV4Parameters adls gen2 cloud storage v4 parameters
// swagger:model AdlsGen2CloudStorageV4Parameters
type AdlsGen2CloudStorageV4Parameters struct {

	// account key
	// Required: true
	AccountKey *string `json:"accountKey"`

	// account name
	// Required: true
	AccountName *string `json:"accountName"`
}

// Validate validates this adls gen2 cloud storage v4 parameters
func (m *AdlsGen2CloudStorageV4Parameters) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountKey(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAccountName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AdlsGen2CloudStorageV4Parameters) validateAccountKey(formats strfmt.Registry) error {

	if err := validate.Required("accountKey", "body", m.AccountKey); err != nil {
		return err
	}

	return nil
}

func (m *AdlsGen2CloudStorageV4Parameters) validateAccountName(formats strfmt.Registry) error {

	if err := validate.Required("accountName", "body", m.AccountName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AdlsGen2CloudStorageV4Parameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AdlsGen2CloudStorageV4Parameters) UnmarshalBinary(b []byte) error {
	var res AdlsGen2CloudStorageV4Parameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
