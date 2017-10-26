// Code generated by go-swagger; DO NOT EDIT.

package models_cloudbreak

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// StackV2Request stack v2 request
// swagger:model StackV2Request

type StackV2Request struct {

	// specific version of ambari
	AmbariVersion string `json:"ambariVersion,omitempty"`

	// stack related application tags
	ApplicationTags map[string]string `json:"applicationTags,omitempty"`

	// availability zone of the stack
	AvailabilityZone string `json:"availabilityZone,omitempty"`

	// using the cluster name to create subdomain
	ClusterNameAsSubdomain *bool `json:"clusterNameAsSubdomain,omitempty"`

	// cluster request object on stack
	ClusterRequest *ClusterV2Request `json:"clusterRequest,omitempty"`

	// credential resource name for the stack
	CredentialName string `json:"credentialName,omitempty"`

	// custom domain name for the nodes in the stack
	CustomDomain string `json:"customDomain,omitempty"`

	// custom hostname for nodes in the stack
	CustomHostname string `json:"customHostname,omitempty"`

	// custom image for the cluster
	CustomImage string `json:"customImage,omitempty"`

	// stack related default tags
	DefaultTags map[string]string `json:"defaultTags,omitempty"`

	// failure policy in case of failures
	FailurePolicy *FailurePolicyRequest `json:"failurePolicy,omitempty"`

	// id of the related flex subscription
	FlexID int64 `json:"flexId,omitempty"`

	// specific version of HDP
	HdpVersion string `json:"hdpVersion,omitempty"`

	// using the hostgroup names to create hostnames
	HostgroupNameAsHostname *bool `json:"hostgroupNameAsHostname,omitempty"`

	// custom image catalog URL
	ImageCatalog string `json:"imageCatalog,omitempty"`

	// collection of instance groupst
	// Required: true
	InstanceGroups []*InstanceGroupsV2 `json:"instanceGroups"`

	// name of the stack
	// Required: true
	// Max Length: 40
	// Min Length: 5
	// Pattern: ([a-z][-a-z0-9]*[a-z0-9])
	Name *string `json:"name"`

	// stack related network
	Network *NetworkV2Request `json:"network,omitempty"`

	// action on failure
	OnFailureAction string `json:"onFailureAction,omitempty"`

	// the details of the container orchestrator api to use
	Orchestrator *OrchestratorRequest `json:"orchestrator,omitempty"`

	// additional cloud specific parameters for stack
	Parameters map[string]string `json:"parameters,omitempty"`

	// cloud provider api variant
	PlatformVariant string `json:"platformVariant,omitempty"`

	// region of the stack
	Region string `json:"region,omitempty"`

	// stack related authentication
	StackAuthentication *StackAuthentication `json:"stackAuthentication,omitempty"`

	// stack related userdefined tags
	UserDefinedTags map[string]string `json:"userDefinedTags,omitempty"`
}

/* polymorph StackV2Request ambariVersion false */

/* polymorph StackV2Request applicationTags false */

/* polymorph StackV2Request availabilityZone false */

/* polymorph StackV2Request clusterNameAsSubdomain false */

/* polymorph StackV2Request clusterRequest false */

/* polymorph StackV2Request credentialName false */

/* polymorph StackV2Request customDomain false */

/* polymorph StackV2Request customHostname false */

/* polymorph StackV2Request customImage false */

/* polymorph StackV2Request defaultTags false */

/* polymorph StackV2Request failurePolicy false */

/* polymorph StackV2Request flexId false */

/* polymorph StackV2Request hdpVersion false */

/* polymorph StackV2Request hostgroupNameAsHostname false */

/* polymorph StackV2Request imageCatalog false */

/* polymorph StackV2Request instanceGroups false */

/* polymorph StackV2Request name false */

/* polymorph StackV2Request network false */

/* polymorph StackV2Request onFailureAction false */

/* polymorph StackV2Request orchestrator false */

/* polymorph StackV2Request parameters false */

/* polymorph StackV2Request platformVariant false */

/* polymorph StackV2Request region false */

/* polymorph StackV2Request stackAuthentication false */

/* polymorph StackV2Request userDefinedTags false */

// Validate validates this stack v2 request
func (m *StackV2Request) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateClusterRequest(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateFailurePolicy(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateInstanceGroups(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateNetwork(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOnFailureAction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOrchestrator(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateStackAuthentication(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackV2Request) validateClusterRequest(formats strfmt.Registry) error {

	if swag.IsZero(m.ClusterRequest) { // not required
		return nil
	}

	if m.ClusterRequest != nil {

		if err := m.ClusterRequest.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clusterRequest")
			}
			return err
		}
	}

	return nil
}

func (m *StackV2Request) validateFailurePolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.FailurePolicy) { // not required
		return nil
	}

	if m.FailurePolicy != nil {

		if err := m.FailurePolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("failurePolicy")
			}
			return err
		}
	}

	return nil
}

func (m *StackV2Request) validateInstanceGroups(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroups", "body", m.InstanceGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.InstanceGroups); i++ {

		if swag.IsZero(m.InstanceGroups[i]) { // not required
			continue
		}

		if m.InstanceGroups[i] != nil {

			if err := m.InstanceGroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("instanceGroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *StackV2Request) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(*m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(*m.Name), 40); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(*m.Name), `([a-z][-a-z0-9]*[a-z0-9])`); err != nil {
		return err
	}

	return nil
}

func (m *StackV2Request) validateNetwork(formats strfmt.Registry) error {

	if swag.IsZero(m.Network) { // not required
		return nil
	}

	if m.Network != nil {

		if err := m.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

var stackV2RequestTypeOnFailureActionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ROLLBACK","DO_NOTHING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		stackV2RequestTypeOnFailureActionPropEnum = append(stackV2RequestTypeOnFailureActionPropEnum, v)
	}
}

const (
	// StackV2RequestOnFailureActionROLLBACK captures enum value "ROLLBACK"
	StackV2RequestOnFailureActionROLLBACK string = "ROLLBACK"
	// StackV2RequestOnFailureActionDONOTHING captures enum value "DO_NOTHING"
	StackV2RequestOnFailureActionDONOTHING string = "DO_NOTHING"
)

// prop value enum
func (m *StackV2Request) validateOnFailureActionEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, stackV2RequestTypeOnFailureActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackV2Request) validateOnFailureAction(formats strfmt.Registry) error {

	if swag.IsZero(m.OnFailureAction) { // not required
		return nil
	}

	// value enum
	if err := m.validateOnFailureActionEnum("onFailureAction", "body", m.OnFailureAction); err != nil {
		return err
	}

	return nil
}

func (m *StackV2Request) validateOrchestrator(formats strfmt.Registry) error {

	if swag.IsZero(m.Orchestrator) { // not required
		return nil
	}

	if m.Orchestrator != nil {

		if err := m.Orchestrator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("orchestrator")
			}
			return err
		}
	}

	return nil
}

func (m *StackV2Request) validateStackAuthentication(formats strfmt.Registry) error {

	if swag.IsZero(m.StackAuthentication) { // not required
		return nil
	}

	if m.StackAuthentication != nil {

		if err := m.StackAuthentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("stackAuthentication")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *StackV2Request) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *StackV2Request) UnmarshalBinary(b []byte) error {
	var res StackV2Request
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}