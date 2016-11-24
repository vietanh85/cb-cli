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

/*UpdateStack update stack

swagger:model UpdateStack
*/
type UpdateStack struct {

	/* instnce group adjustment
	 */
	InstanceGroupAdjustment *InstanceGroupAdjustment `json:"instanceGroupAdjustment,omitempty"`

	/* status of the scale request
	 */
	Status *string `json:"status,omitempty"`
}

// Validate validates this update stack
func (m *UpdateStack) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateStackTypeStatusPropEnum []interface{}

func (m *UpdateStack) validateStatusEnum(path, location string, value string) error {
	if updateStackTypeStatusPropEnum == nil {
		var res []string
		if err := json.Unmarshal([]byte(`["SYNC","FULL_SYNC","REPAIR_FAILED_NODES","STOPPED","STARTED"]`), &res); err != nil {
			return err
		}
		for _, v := range res {
			updateStackTypeStatusPropEnum = append(updateStackTypeStatusPropEnum, v)
		}
	}
	if err := validate.Enum(path, location, value, updateStackTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *UpdateStack) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}
