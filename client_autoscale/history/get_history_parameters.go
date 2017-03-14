package history

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetHistoryParams creates a new GetHistoryParams object
// with the default values initialized.
func NewGetHistoryParams() *GetHistoryParams {
	var ()
	return &GetHistoryParams{}
}

/*GetHistoryParams contains all the parameters to send to the API endpoint
for the get history operation typically these are written to a http.Request
*/
type GetHistoryParams struct {

	/*ClusterID*/
	ClusterID int64
}

// WithClusterID adds the clusterId to the get history params
func (o *GetHistoryParams) WithClusterID(clusterId int64) *GetHistoryParams {
	o.ClusterID = clusterId
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetHistoryParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param clusterId
	if err := r.SetPathParam("clusterId", swag.FormatInt64(o.ClusterID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
