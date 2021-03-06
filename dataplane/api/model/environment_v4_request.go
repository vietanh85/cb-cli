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

// EnvironmentV4Request environment v4 request
// swagger:model EnvironmentV4Request
type EnvironmentV4Request struct {

	// If credentialName is not specified, the credential is used to create the new credential for the environment.
	Credential *CredentialV4Request `json:"credential,omitempty"`

	// Name of the credential of the environment. If the name is given, the detailed credential is ignored in the request.
	CredentialName string `json:"credentialName,omitempty"`

	// Name of the RDS configurations to be attached to the environment.
	// Unique: true
	Databases []string `json:"databases"`

	// description of the resource
	// Max Length: 1000
	// Min Length: 0
	Description *string `json:"description,omitempty"`

	// Name of Kerberos configs to be attached to the environment.
	// Unique: true
	Kerberoses []string `json:"kerberoses"`

	// Name of the Kubernetes configurations to be attached to the environment.
	// Unique: true
	Kubernetes []string `json:"kubernetes"`

	// Name of the LDAP configurations to be attached to the environment.
	// Unique: true
	Ldaps []string `json:"ldaps"`

	// Location of the environment.
	// Required: true
	Location *LocationV4Request `json:"location"`

	// name of the resource
	// Required: true
	// Max Length: 100
	// Min Length: 5
	// Pattern: (^[a-z][-a-z0-9]*[a-z0-9]$)
	Name *string `json:"name"`

	// Name of the proxy configurations to be attached to the environment.
	// Unique: true
	Proxies []string `json:"proxies"`

	// Regions of the environment.
	// Unique: true
	Regions []string `json:"regions"`
}

// Validate validates this environment v4 request
func (m *EnvironmentV4Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredential(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDatabases(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKerberoses(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKubernetes(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLdaps(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxies(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRegions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *EnvironmentV4Request) validateCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.Credential) { // not required
		return nil
	}

	if m.Credential != nil {
		if err := m.Credential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentV4Request) validateDatabases(formats strfmt.Registry) error {

	if swag.IsZero(m.Databases) { // not required
		return nil
	}

	if err := validate.UniqueItems("databases", "body", m.Databases); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentV4Request) validateDescription(formats strfmt.Registry) error {

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

func (m *EnvironmentV4Request) validateKerberoses(formats strfmt.Registry) error {

	if swag.IsZero(m.Kerberoses) { // not required
		return nil
	}

	if err := validate.UniqueItems("kerberoses", "body", m.Kerberoses); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentV4Request) validateKubernetes(formats strfmt.Registry) error {

	if swag.IsZero(m.Kubernetes) { // not required
		return nil
	}

	if err := validate.UniqueItems("kubernetes", "body", m.Kubernetes); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentV4Request) validateLdaps(formats strfmt.Registry) error {

	if swag.IsZero(m.Ldaps) { // not required
		return nil
	}

	if err := validate.UniqueItems("ldaps", "body", m.Ldaps); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentV4Request) validateLocation(formats strfmt.Registry) error {

	if err := validate.Required("location", "body", m.Location); err != nil {
		return err
	}

	if m.Location != nil {
		if err := m.Location.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("location")
			}
			return err
		}
	}

	return nil
}

func (m *EnvironmentV4Request) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 100); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `(^[a-z][-a-z0-9]*[a-z0-9]$)`); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentV4Request) validateProxies(formats strfmt.Registry) error {

	if swag.IsZero(m.Proxies) { // not required
		return nil
	}

	if err := validate.UniqueItems("proxies", "body", m.Proxies); err != nil {
		return err
	}

	return nil
}

func (m *EnvironmentV4Request) validateRegions(formats strfmt.Registry) error {

	if swag.IsZero(m.Regions) { // not required
		return nil
	}

	if err := validate.UniqueItems("regions", "body", m.Regions); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EnvironmentV4Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EnvironmentV4Request) UnmarshalBinary(b []byte) error {
	var res EnvironmentV4Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
