// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewPutClusterV2Params creates a new PutClusterV2Params object
// with the default values initialized.
func NewPutClusterV2Params() *PutClusterV2Params {
	var ()
	return &PutClusterV2Params{

		timeout: cr.DefaultTimeout,
	}
}

// NewPutClusterV2ParamsWithTimeout creates a new PutClusterV2Params object
// with the default values initialized, and the ability to set a timeout on a request
func NewPutClusterV2ParamsWithTimeout(timeout time.Duration) *PutClusterV2Params {
	var ()
	return &PutClusterV2Params{

		timeout: timeout,
	}
}

// NewPutClusterV2ParamsWithContext creates a new PutClusterV2Params object
// with the default values initialized, and the ability to set a context for a request
func NewPutClusterV2ParamsWithContext(ctx context.Context) *PutClusterV2Params {
	var ()
	return &PutClusterV2Params{

		Context: ctx,
	}
}

// NewPutClusterV2ParamsWithHTTPClient creates a new PutClusterV2Params object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPutClusterV2ParamsWithHTTPClient(client *http.Client) *PutClusterV2Params {
	var ()
	return &PutClusterV2Params{
		HTTPClient: client,
	}
}

/*PutClusterV2Params contains all the parameters to send to the API endpoint
for the put cluster v2 operation typically these are written to a http.Request
*/
type PutClusterV2Params struct {

	/*Body*/
	Body *models_cloudbreak.UpdateClusterV2
	/*Name*/
	Name string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the put cluster v2 params
func (o *PutClusterV2Params) WithTimeout(timeout time.Duration) *PutClusterV2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the put cluster v2 params
func (o *PutClusterV2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the put cluster v2 params
func (o *PutClusterV2Params) WithContext(ctx context.Context) *PutClusterV2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the put cluster v2 params
func (o *PutClusterV2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the put cluster v2 params
func (o *PutClusterV2Params) WithHTTPClient(client *http.Client) *PutClusterV2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the put cluster v2 params
func (o *PutClusterV2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the put cluster v2 params
func (o *PutClusterV2Params) WithBody(body *models_cloudbreak.UpdateClusterV2) *PutClusterV2Params {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the put cluster v2 params
func (o *PutClusterV2Params) SetBody(body *models_cloudbreak.UpdateClusterV2) {
	o.Body = body
}

// WithName adds the name to the put cluster v2 params
func (o *PutClusterV2Params) WithName(name string) *PutClusterV2Params {
	o.SetName(name)
	return o
}

// SetName adds the name to the put cluster v2 params
func (o *PutClusterV2Params) SetName(name string) {
	o.Name = name
}

// WriteToRequest writes these params to a swagger request
func (o *PutClusterV2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.UpdateClusterV2)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}