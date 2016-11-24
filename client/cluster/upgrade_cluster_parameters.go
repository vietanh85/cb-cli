package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models"
)

// NewUpgradeClusterParams creates a new UpgradeClusterParams object
// with the default values initialized.
func NewUpgradeClusterParams() *UpgradeClusterParams {
	var ()
	return &UpgradeClusterParams{}
}

/*UpgradeClusterParams contains all the parameters to send to the API endpoint
for the upgrade cluster operation typically these are written to a http.Request
*/
type UpgradeClusterParams struct {

	/*Body*/
	Body *models.AmbariRepoDetails
	/*ID*/
	ID int64
}

// WithBody adds the body to the upgrade cluster params
func (o *UpgradeClusterParams) WithBody(body *models.AmbariRepoDetails) *UpgradeClusterParams {
	o.Body = body
	return o
}

// WithID adds the id to the upgrade cluster params
func (o *UpgradeClusterParams) WithID(id int64) *UpgradeClusterParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *UpgradeClusterParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models.AmbariRepoDetails)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
