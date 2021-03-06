// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_databases

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

// NewDetachDatabaseFromEnvironmentsParams creates a new DetachDatabaseFromEnvironmentsParams object
// with the default values initialized.
func NewDetachDatabaseFromEnvironmentsParams() *DetachDatabaseFromEnvironmentsParams {
	var ()
	return &DetachDatabaseFromEnvironmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDetachDatabaseFromEnvironmentsParamsWithTimeout creates a new DetachDatabaseFromEnvironmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDetachDatabaseFromEnvironmentsParamsWithTimeout(timeout time.Duration) *DetachDatabaseFromEnvironmentsParams {
	var ()
	return &DetachDatabaseFromEnvironmentsParams{

		timeout: timeout,
	}
}

// NewDetachDatabaseFromEnvironmentsParamsWithContext creates a new DetachDatabaseFromEnvironmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDetachDatabaseFromEnvironmentsParamsWithContext(ctx context.Context) *DetachDatabaseFromEnvironmentsParams {
	var ()
	return &DetachDatabaseFromEnvironmentsParams{

		Context: ctx,
	}
}

// NewDetachDatabaseFromEnvironmentsParamsWithHTTPClient creates a new DetachDatabaseFromEnvironmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDetachDatabaseFromEnvironmentsParamsWithHTTPClient(client *http.Client) *DetachDatabaseFromEnvironmentsParams {
	var ()
	return &DetachDatabaseFromEnvironmentsParams{
		HTTPClient: client,
	}
}

/*DetachDatabaseFromEnvironmentsParams contains all the parameters to send to the API endpoint
for the detach database from environments operation typically these are written to a http.Request
*/
type DetachDatabaseFromEnvironmentsParams struct {

	/*Body*/
	Body *model.EnvironmentNames
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the detach database from environments params
func (o *DetachDatabaseFromEnvironmentsParams) WithTimeout(timeout time.Duration) *DetachDatabaseFromEnvironmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detach database from environments params
func (o *DetachDatabaseFromEnvironmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detach database from environments params
func (o *DetachDatabaseFromEnvironmentsParams) WithContext(ctx context.Context) *DetachDatabaseFromEnvironmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detach database from environments params
func (o *DetachDatabaseFromEnvironmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detach database from environments params
func (o *DetachDatabaseFromEnvironmentsParams) WithHTTPClient(client *http.Client) *DetachDatabaseFromEnvironmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detach database from environments params
func (o *DetachDatabaseFromEnvironmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the detach database from environments params
func (o *DetachDatabaseFromEnvironmentsParams) WithBody(body *model.EnvironmentNames) *DetachDatabaseFromEnvironmentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the detach database from environments params
func (o *DetachDatabaseFromEnvironmentsParams) SetBody(body *model.EnvironmentNames) {
	o.Body = body
}

// WithName adds the name to the detach database from environments params
func (o *DetachDatabaseFromEnvironmentsParams) WithName(name string) *DetachDatabaseFromEnvironmentsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the detach database from environments params
func (o *DetachDatabaseFromEnvironmentsParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the detach database from environments params
func (o *DetachDatabaseFromEnvironmentsParams) WithWorkspaceID(workspaceID int64) *DetachDatabaseFromEnvironmentsParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the detach database from environments params
func (o *DetachDatabaseFromEnvironmentsParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DetachDatabaseFromEnvironmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	// path param workspaceId
	if err := r.SetPathParam("workspaceId", swag.FormatInt64(o.WorkspaceID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
