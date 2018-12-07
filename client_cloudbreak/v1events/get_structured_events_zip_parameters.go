// Code generated by go-swagger; DO NOT EDIT.

package v1events

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

// NewGetStructuredEventsZipParams creates a new GetStructuredEventsZipParams object
// with the default values initialized.
func NewGetStructuredEventsZipParams() *GetStructuredEventsZipParams {
	var ()
	return &GetStructuredEventsZipParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetStructuredEventsZipParamsWithTimeout creates a new GetStructuredEventsZipParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetStructuredEventsZipParamsWithTimeout(timeout time.Duration) *GetStructuredEventsZipParams {
	var ()
	return &GetStructuredEventsZipParams{

		timeout: timeout,
	}
}

// NewGetStructuredEventsZipParamsWithContext creates a new GetStructuredEventsZipParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetStructuredEventsZipParamsWithContext(ctx context.Context) *GetStructuredEventsZipParams {
	var ()
	return &GetStructuredEventsZipParams{

		Context: ctx,
	}
}

// NewGetStructuredEventsZipParamsWithHTTPClient creates a new GetStructuredEventsZipParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetStructuredEventsZipParamsWithHTTPClient(client *http.Client) *GetStructuredEventsZipParams {
	var ()
	return &GetStructuredEventsZipParams{
		HTTPClient: client,
	}
}

/*GetStructuredEventsZipParams contains all the parameters to send to the API endpoint
for the get structured events zip operation typically these are written to a http.Request
*/
type GetStructuredEventsZipParams struct {

	/*StackID*/
	StackID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get structured events zip params
func (o *GetStructuredEventsZipParams) WithTimeout(timeout time.Duration) *GetStructuredEventsZipParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get structured events zip params
func (o *GetStructuredEventsZipParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get structured events zip params
func (o *GetStructuredEventsZipParams) WithContext(ctx context.Context) *GetStructuredEventsZipParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get structured events zip params
func (o *GetStructuredEventsZipParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get structured events zip params
func (o *GetStructuredEventsZipParams) WithHTTPClient(client *http.Client) *GetStructuredEventsZipParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get structured events zip params
func (o *GetStructuredEventsZipParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithStackID adds the stackID to the get structured events zip params
func (o *GetStructuredEventsZipParams) WithStackID(stackID int64) *GetStructuredEventsZipParams {
	o.SetStackID(stackID)
	return o
}

// SetStackID adds the stackId to the get structured events zip params
func (o *GetStructuredEventsZipParams) SetStackID(stackID int64) {
	o.StackID = stackID
}

// WriteToRequest writes these params to a swagger request
func (o *GetStructuredEventsZipParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param stackId
	if err := r.SetPathParam("stackId", swag.FormatInt64(o.StackID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}