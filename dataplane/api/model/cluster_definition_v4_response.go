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

// ClusterDefinitionV4Response cluster definition v4 response
// swagger:model ClusterDefinitionV4Response
type ClusterDefinitionV4Response struct {

	// cluster definition, set this or the url field
	ClusterDefinition string `json:"clusterDefinition,omitempty"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// number of host groups
	HostGroupCount int32 `json:"hostGroupCount,omitempty"`

	// id of the resource
	ID int64 `json:"id,omitempty"`

	// name of the resource
	// Required: true
	// Max Length: 100
	// Min Length: 1
	// Pattern: ^[^;\/%]*$
	Name *string `json:"name"`

	// status of the cluster definition
	// Enum: [DEFAULT DEFAULT_DELETED USER_MANAGED OUTDATED]
	Status string `json:"status,omitempty"`

	// user defined tags for cluster definition
	Tags map[string]interface{} `json:"tags,omitempty"`
}

// Validate validates this cluster definition v4 response
func (m *ClusterDefinitionV4Response) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ClusterDefinitionV4Response) validateDescription(formats strfmt.Registry) error {

	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MinLength("description", "body", string(*m.Description), 0); err != nil {
		return err
	}

	if err := validate.MaxLength("description", "body", string(*m.Description), 1000); err != nil {
		return err
	}

	return nil
}

func (m *ClusterDefinitionV4Response) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `^[^;\/%]*$`); err != nil {
		return err
	}

	return nil
}

var clusterDefinitionV4ResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","DEFAULT_DELETED","USER_MANAGED","OUTDATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterDefinitionV4ResponseTypeStatusPropEnum = append(clusterDefinitionV4ResponseTypeStatusPropEnum, v)
	}
}

const (

	// ClusterDefinitionV4ResponseStatusDEFAULT captures enum value "DEFAULT"
	ClusterDefinitionV4ResponseStatusDEFAULT string = "DEFAULT"

	// ClusterDefinitionV4ResponseStatusDEFAULTDELETED captures enum value "DEFAULT_DELETED"
	ClusterDefinitionV4ResponseStatusDEFAULTDELETED string = "DEFAULT_DELETED"

	// ClusterDefinitionV4ResponseStatusUSERMANAGED captures enum value "USER_MANAGED"
	ClusterDefinitionV4ResponseStatusUSERMANAGED string = "USER_MANAGED"

	// ClusterDefinitionV4ResponseStatusOUTDATED captures enum value "OUTDATED"
	ClusterDefinitionV4ResponseStatusOUTDATED string = "OUTDATED"
)

// prop value enum
func (m *ClusterDefinitionV4Response) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterDefinitionV4ResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterDefinitionV4Response) validateStatus(formats strfmt.Registry) error {

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
func (m *ClusterDefinitionV4Response) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDefinitionV4Response) UnmarshalBinary(b []byte) error {
	var res ClusterDefinitionV4Response
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
