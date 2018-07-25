// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// AzureParameters azure parameters
// swagger:model AzureParameters

type AzureParameters struct {

	// private id for azure
	PrivateID string `json:"privateId,omitempty"`
}

/* polymorph AzureParameters privateId false */

// Validate validates this azure parameters
func (m *AzureParameters) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// MarshalBinary interface implementation
func (m *AzureParameters) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AzureParameters) UnmarshalBinary(b []byte) error {
	var res AzureParameters
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
