package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_autoscale"
)

// NewSetStateParams creates a new SetStateParams object
// with the default values initialized.
func NewSetStateParams() *SetStateParams {
	var ()
	return &SetStateParams{}
}

/*SetStateParams contains all the parameters to send to the API endpoint
for the set state operation typically these are written to a http.Request
*/
type SetStateParams struct {

	/*Body*/
	Body *models_autoscale.ClusterState
	/*ClusterID*/
	ClusterID int64
}

// WithBody adds the body to the set state params
func (o *SetStateParams) WithBody(body *models_autoscale.ClusterState) *SetStateParams {
	o.Body = body
	return o
}

// WithClusterID adds the clusterId to the set state params
func (o *SetStateParams) WithClusterID(clusterId int64) *SetStateParams {
	o.ClusterID = clusterId
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *SetStateParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models_autoscale.ClusterState)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param clusterId
	if err := r.SetPathParam("clusterId", swag.FormatInt64(o.ClusterID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
