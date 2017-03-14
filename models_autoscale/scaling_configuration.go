package models_autoscale

import "github.com/go-swagger/go-swagger/strfmt"

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

/*ScalingConfiguration scaling configuration

swagger:model ScalingConfiguration
*/
type ScalingConfiguration struct {

	/* The time between two scaling activities
	 */
	Cooldown *int32 `json:"cooldown,omitempty"`

	/* The maximum size of the cluster after scaling
	 */
	MaxSize *int32 `json:"maxSize,omitempty"`

	/* The minimum size of the cluster after scaling
	 */
	MinSize *int32 `json:"minSize,omitempty"`
}

// Validate validates this scaling configuration
func (m *ScalingConfiguration) Validate(formats strfmt.Registry) error {
	return nil
}