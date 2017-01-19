package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
	"github.com/go-swagger/go-swagger/swag"

	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/httpkit/validate"
)

/*StackRequest stack request

swagger:model StackRequest
*/
type StackRequest struct {

	/* specific version of ambari
	 */
	AmbariVersion *string `json:"ambariVersion,omitempty"`

	/* availability zone of the stack
	 */
	AvailabilityZone *string `json:"availabilityZone,omitempty"`

	/* type of cloud provider

	Read Only: true
	*/
	CloudPlatform *string `json:"cloudPlatform,omitempty"`

	/* stack related credential
	 */
	Credential *CredentialRequest `json:"credential,omitempty"`

	/* credential resource id for the stack
	 */
	CredentialID *int64 `json:"credentialId,omitempty"`

	/* failure policy in case of failures
	 */
	FailurePolicy *FailurePolicyRequest `json:"failurePolicy,omitempty"`

	/* specific version of HDP
	 */
	HdpVersion *string `json:"hdpVersion,omitempty"`

	/* custom image catalog URL
	 */
	ImageCatalog *string `json:"imageCatalog,omitempty"`

	/* collection of instance groupst

	Required: true
	*/
	InstanceGroups []*InstanceGroups `json:"instanceGroups"`

	/* name of the stack

	Required: true
	Max Length: 40
	Min Length: 5
	Pattern: ([a-z][-a-z0-9]*[a-z0-9])
	*/
	Name string `json:"name"`

	/* stack related network
	 */
	Network *NetworkRequest `json:"network,omitempty"`

	/* network resource id for the stack
	 */
	NetworkID *int64 `json:"networkId,omitempty"`

	/* action on failure
	 */
	OnFailureAction *string `json:"onFailureAction,omitempty"`

	/* the details of the container orchestrator api to use
	 */
	Orchestrator *OrchestratorRequest `json:"orchestrator,omitempty"`

	/* additional cloud specific parameters for stack
	 */
	Parameters map[string]string `json:"parameters,omitempty"`

	/* cloud provider api variant
	 */
	PlatformVariant *string `json:"platformVariant,omitempty"`

	/* region of the stack
	 */
	Region *string `json:"region,omitempty"`

	/* relocate the docker service in startup time
	 */
	RelocateDocker *bool `json:"relocateDocker,omitempty"`

	/* stack related tags
	 */
	Tags map[string]interface{} `json:"tags,omitempty"`
}

// Validate validates this stack request
func (m *StackRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInstanceGroups(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateOnFailureAction(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *StackRequest) validateInstanceGroups(formats strfmt.Registry) error {

	if err := validate.Required("instanceGroups", "body", m.InstanceGroups); err != nil {
		return err
	}

	for i := 0; i < len(m.InstanceGroups); i++ {

		if m.InstanceGroups[i] != nil {

			if err := m.InstanceGroups[i].Validate(formats); err != nil {
				return err
			}
		}

	}

	return nil
}

func (m *StackRequest) validateName(formats strfmt.Registry) error {

	if err := validate.RequiredString("name", "body", string(m.Name)); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", string(m.Name), 5); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", string(m.Name), 40); err != nil {
		return err
	}

	if err := validate.Pattern("name", "body", string(m.Name), `([a-z][-a-z0-9]*[a-z0-9])`); err != nil {
		return err
	}

	return nil
}

var stackRequestTypeOnFailureActionPropEnum []interface{}

func (m *StackRequest) validateOnFailureActionEnum(path, location string, value string) error {
	if stackRequestTypeOnFailureActionPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["ROLLBACK","DO_NOTHING"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			stackRequestTypeOnFailureActionPropEnum = append(stackRequestTypeOnFailureActionPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, stackRequestTypeOnFailureActionPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *StackRequest) validateOnFailureAction(formats strfmt.Registry) error {

	if swag.IsZero(m.OnFailureAction) { // not required
		return nil
	}

	if err := m.validateOnFailureActionEnum("onFailureAction", "body", *m.OnFailureAction); err != nil {
		return err
	}

	return nil
}

func (m *StackRequest) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	if err := validate.Required("parameters", "body", m.Parameters); err != nil {
		return err
	}

	return nil
}
