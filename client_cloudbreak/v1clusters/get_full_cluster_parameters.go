// Code generated by go-swagger; DO NOT EDIT.

package v1clusters

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
)

// NewGetFullClusterParams creates a new GetFullClusterParams object
// with the default values initialized.
func NewGetFullClusterParams() *GetFullClusterParams {
	var ()
	return &GetFullClusterParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetFullClusterParamsWithTimeout creates a new GetFullClusterParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetFullClusterParamsWithTimeout(timeout time.Duration) *GetFullClusterParams {
	var ()
	return &GetFullClusterParams{

		timeout: timeout,
	}
}

// NewGetFullClusterParamsWithContext creates a new GetFullClusterParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetFullClusterParamsWithContext(ctx context.Context) *GetFullClusterParams {
	var ()
	return &GetFullClusterParams{

		Context: ctx,
	}
}

// NewGetFullClusterParamsWithHTTPClient creates a new GetFullClusterParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetFullClusterParamsWithHTTPClient(client *http.Client) *GetFullClusterParams {
	var ()
	return &GetFullClusterParams{
		HTTPClient: client,
	}
}

/*GetFullClusterParams contains all the parameters to send to the API endpoint
for the get full cluster operation typically these are written to a http.Request
*/
type GetFullClusterParams struct {

	/*ID*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get full cluster params
func (o *GetFullClusterParams) WithTimeout(timeout time.Duration) *GetFullClusterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get full cluster params
func (o *GetFullClusterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get full cluster params
func (o *GetFullClusterParams) WithContext(ctx context.Context) *GetFullClusterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get full cluster params
func (o *GetFullClusterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get full cluster params
func (o *GetFullClusterParams) WithHTTPClient(client *http.Client) *GetFullClusterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get full cluster params
func (o *GetFullClusterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get full cluster params
func (o *GetFullClusterParams) WithID(id int64) *GetFullClusterParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get full cluster params
func (o *GetFullClusterParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetFullClusterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}