// Code generated by go-swagger; DO NOT EDIT.

package autoscale

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewPutStackForAutoscaleParams creates a new PutStackForAutoscaleParams object
// with the default values initialized.
func NewPutStackForAutoscaleParams() *PutStackForAutoscaleParams {
	var ()
	return &PutStackForAutoscaleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutStackForAutoscaleParamsWithTimeout creates a new PutStackForAutoscaleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutStackForAutoscaleParamsWithTimeout(timeout time.Duration) *PutStackForAutoscaleParams {
	var ()
	return &PutStackForAutoscaleParams{

		timeout: timeout,
	}
}

// NewPutStackForAutoscaleParamsWithContext creates a new PutStackForAutoscaleParams object
// with the default values initialized, and the ability to set a context for a request
func NewPutStackForAutoscaleParamsWithContext(ctx context.Context) *PutStackForAutoscaleParams {
	var ()
	return &PutStackForAutoscaleParams{

		Context: ctx,
	}
}

// NewPutStackForAutoscaleParamsWithHTTPClient creates a new PutStackForAutoscaleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutStackForAutoscaleParamsWithHTTPClient(client *http.Client) *PutStackForAutoscaleParams {
	var ()
	return &PutStackForAutoscaleParams{
		HTTPClient: client,
	}
}

/*PutStackForAutoscaleParams contains all the parameters to send to the API endpoint
for the put stack for autoscale operation typically these are written to a http.Request
*/
type PutStackForAutoscaleParams struct {

	/*Body*/
	Body *model.UpdateStack
	/*ID*/
	ID int64
	/*UserID*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put stack for autoscale params
func (o *PutStackForAutoscaleParams) WithTimeout(timeout time.Duration) *PutStackForAutoscaleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put stack for autoscale params
func (o *PutStackForAutoscaleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put stack for autoscale params
func (o *PutStackForAutoscaleParams) WithContext(ctx context.Context) *PutStackForAutoscaleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put stack for autoscale params
func (o *PutStackForAutoscaleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put stack for autoscale params
func (o *PutStackForAutoscaleParams) WithHTTPClient(client *http.Client) *PutStackForAutoscaleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put stack for autoscale params
func (o *PutStackForAutoscaleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put stack for autoscale params
func (o *PutStackForAutoscaleParams) WithBody(body *model.UpdateStack) *PutStackForAutoscaleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put stack for autoscale params
func (o *PutStackForAutoscaleParams) SetBody(body *model.UpdateStack) {
	o.Body = body
}

// WithID adds the id to the put stack for autoscale params
func (o *PutStackForAutoscaleParams) WithID(id int64) *PutStackForAutoscaleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the put stack for autoscale params
func (o *PutStackForAutoscaleParams) SetID(id int64) {
	o.ID = id
}

// WithUserID adds the userID to the put stack for autoscale params
func (o *PutStackForAutoscaleParams) WithUserID(userID string) *PutStackForAutoscaleParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the put stack for autoscale params
func (o *PutStackForAutoscaleParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *PutStackForAutoscaleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	// path param userId
	if err := r.SetPathParam("userId", o.UserID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}