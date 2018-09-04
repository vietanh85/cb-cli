// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_events

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

// NewGetEventsInOrganizationParams creates a new GetEventsInOrganizationParams object
// with the default values initialized.
func NewGetEventsInOrganizationParams() *GetEventsInOrganizationParams {
	var ()
	return &GetEventsInOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEventsInOrganizationParamsWithTimeout creates a new GetEventsInOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEventsInOrganizationParamsWithTimeout(timeout time.Duration) *GetEventsInOrganizationParams {
	var ()
	return &GetEventsInOrganizationParams{

		timeout: timeout,
	}
}

// NewGetEventsInOrganizationParamsWithContext creates a new GetEventsInOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEventsInOrganizationParamsWithContext(ctx context.Context) *GetEventsInOrganizationParams {
	var ()
	return &GetEventsInOrganizationParams{

		Context: ctx,
	}
}

// NewGetEventsInOrganizationParamsWithHTTPClient creates a new GetEventsInOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEventsInOrganizationParamsWithHTTPClient(client *http.Client) *GetEventsInOrganizationParams {
	var ()
	return &GetEventsInOrganizationParams{
		HTTPClient: client,
	}
}

/*GetEventsInOrganizationParams contains all the parameters to send to the API endpoint
for the get events in organization operation typically these are written to a http.Request
*/
type GetEventsInOrganizationParams struct {

	/*OrganizationID*/
	OrganizationID int64
	/*Since*/
	Since *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get events in organization params
func (o *GetEventsInOrganizationParams) WithTimeout(timeout time.Duration) *GetEventsInOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get events in organization params
func (o *GetEventsInOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get events in organization params
func (o *GetEventsInOrganizationParams) WithContext(ctx context.Context) *GetEventsInOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get events in organization params
func (o *GetEventsInOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get events in organization params
func (o *GetEventsInOrganizationParams) WithHTTPClient(client *http.Client) *GetEventsInOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get events in organization params
func (o *GetEventsInOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOrganizationID adds the organizationID to the get events in organization params
func (o *GetEventsInOrganizationParams) WithOrganizationID(organizationID int64) *GetEventsInOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the get events in organization params
func (o *GetEventsInOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WithSince adds the since to the get events in organization params
func (o *GetEventsInOrganizationParams) WithSince(since *int64) *GetEventsInOrganizationParams {
	o.SetSince(since)
	return o
}

// SetSince adds the since to the get events in organization params
func (o *GetEventsInOrganizationParams) SetSince(since *int64) {
	o.Since = since
}

// WriteToRequest writes these params to a swagger request
func (o *GetEventsInOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param organizationId
	if err := r.SetPathParam("organizationId", swag.FormatInt64(o.OrganizationID)); err != nil {
		return err
	}

	if o.Since != nil {

		// query param since
		var qrSince int64
		if o.Since != nil {
			qrSince = *o.Since
		}
		qSince := swag.FormatInt64(qrSince)
		if qSince != "" {
			if err := r.SetQueryParam("since", qSince); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}