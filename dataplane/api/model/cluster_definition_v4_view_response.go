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

// ClusterDefinitionV4ViewResponse cluster definition v4 view response
// swagger:model ClusterDefinitionV4ViewResponse
type ClusterDefinitionV4ViewResponse struct {

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

	// The type of the stack: for example HDP or HDF
	StackType string `json:"stackType,omitempty"`

	// The version of the stack
	StackVersion string `json:"stackVersion,omitempty"`

	// status of the cluster definition
	// Enum: [DEFAULT DEFAULT_DELETED USER_MANAGED OUTDATED]
	Status string `json:"status,omitempty"`

	// user defined tags for cluster definition
	Tags map[string]interface{} `json:"tags,omitempty"`
}

// Validate validates this cluster definition v4 view response
func (m *ClusterDefinitionV4ViewResponse) Validate(formats strfmt.Registry) error {
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

func (m *ClusterDefinitionV4ViewResponse) validateDescription(formats strfmt.Registry) error {

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

func (m *ClusterDefinitionV4ViewResponse) validateName(formats strfmt.Registry) error {

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

var clusterDefinitionV4ViewResponseTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["DEFAULT","DEFAULT_DELETED","USER_MANAGED","OUTDATED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		clusterDefinitionV4ViewResponseTypeStatusPropEnum = append(clusterDefinitionV4ViewResponseTypeStatusPropEnum, v)
	}
}

const (

	// ClusterDefinitionV4ViewResponseStatusDEFAULT captures enum value "DEFAULT"
	ClusterDefinitionV4ViewResponseStatusDEFAULT string = "DEFAULT"

	// ClusterDefinitionV4ViewResponseStatusDEFAULTDELETED captures enum value "DEFAULT_DELETED"
	ClusterDefinitionV4ViewResponseStatusDEFAULTDELETED string = "DEFAULT_DELETED"

	// ClusterDefinitionV4ViewResponseStatusUSERMANAGED captures enum value "USER_MANAGED"
	ClusterDefinitionV4ViewResponseStatusUSERMANAGED string = "USER_MANAGED"

	// ClusterDefinitionV4ViewResponseStatusOUTDATED captures enum value "OUTDATED"
	ClusterDefinitionV4ViewResponseStatusOUTDATED string = "OUTDATED"
)

// prop value enum
func (m *ClusterDefinitionV4ViewResponse) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, clusterDefinitionV4ViewResponseTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ClusterDefinitionV4ViewResponse) validateStatus(formats strfmt.Registry) error {

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
func (m *ClusterDefinitionV4ViewResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ClusterDefinitionV4ViewResponse) UnmarshalBinary(b []byte) error {
	var res ClusterDefinitionV4ViewResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
