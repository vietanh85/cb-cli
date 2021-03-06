// Code generated by go-swagger; DO NOT EDIT.

package model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GeneratedClusterDefinitionV4Response generated cluster definition v4 response
// swagger:model GeneratedClusterDefinitionV4Response
type GeneratedClusterDefinitionV4Response struct {

	// cluster definition, set this or the url field
	ClusterDefinitionText string `json:"clusterDefinitionText,omitempty"`
}

// Validate validates this generated cluster definition v4 response
func (m *GeneratedClusterDefinitionV4Response) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GeneratedClusterDefinitionV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GeneratedClusterDefinitionV4Response) UnmarshalBinary(b []byte) error {
	var res GeneratedClusterDefinitionV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
