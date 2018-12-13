// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_ldapconfigs

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

// NewDetachLdapResourceFromEnvironmentsParams creates a new DetachLdapResourceFromEnvironmentsParams object
// with the default values initialized.
func NewDetachLdapResourceFromEnvironmentsParams() *DetachLdapResourceFromEnvironmentsParams {
	var ()
	return &DetachLdapResourceFromEnvironmentsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDetachLdapResourceFromEnvironmentsParamsWithTimeout creates a new DetachLdapResourceFromEnvironmentsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDetachLdapResourceFromEnvironmentsParamsWithTimeout(timeout time.Duration) *DetachLdapResourceFromEnvironmentsParams {
	var ()
	return &DetachLdapResourceFromEnvironmentsParams{

		timeout: timeout,
	}
}

// NewDetachLdapResourceFromEnvironmentsParamsWithContext creates a new DetachLdapResourceFromEnvironmentsParams object
// with the default values initialized, and the ability to set a context for a request
func NewDetachLdapResourceFromEnvironmentsParamsWithContext(ctx context.Context) *DetachLdapResourceFromEnvironmentsParams {
	var ()
	return &DetachLdapResourceFromEnvironmentsParams{

		Context: ctx,
	}
}

// NewDetachLdapResourceFromEnvironmentsParamsWithHTTPClient creates a new DetachLdapResourceFromEnvironmentsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDetachLdapResourceFromEnvironmentsParamsWithHTTPClient(client *http.Client) *DetachLdapResourceFromEnvironmentsParams {
	var ()
	return &DetachLdapResourceFromEnvironmentsParams{
		HTTPClient: client,
	}
}

/*DetachLdapResourceFromEnvironmentsParams contains all the parameters to send to the API endpoint
for the detach ldap resource from environments operation typically these are written to a http.Request
*/
type DetachLdapResourceFromEnvironmentsParams struct {

	/*Body*/
	Body []string
	/*Name*/
	Name string
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the detach ldap resource from environments params
func (o *DetachLdapResourceFromEnvironmentsParams) WithTimeout(timeout time.Duration) *DetachLdapResourceFromEnvironmentsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the detach ldap resource from environments params
func (o *DetachLdapResourceFromEnvironmentsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the detach ldap resource from environments params
func (o *DetachLdapResourceFromEnvironmentsParams) WithContext(ctx context.Context) *DetachLdapResourceFromEnvironmentsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the detach ldap resource from environments params
func (o *DetachLdapResourceFromEnvironmentsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the detach ldap resource from environments params
func (o *DetachLdapResourceFromEnvironmentsParams) WithHTTPClient(client *http.Client) *DetachLdapResourceFromEnvironmentsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the detach ldap resource from environments params
func (o *DetachLdapResourceFromEnvironmentsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the detach ldap resource from environments params
func (o *DetachLdapResourceFromEnvironmentsParams) WithBody(body []string) *DetachLdapResourceFromEnvironmentsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the detach ldap resource from environments params
func (o *DetachLdapResourceFromEnvironmentsParams) SetBody(body []string) {
	o.Body = body
}

// WithName adds the name to the detach ldap resource from environments params
func (o *DetachLdapResourceFromEnvironmentsParams) WithName(name string) *DetachLdapResourceFromEnvironmentsParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the detach ldap resource from environments params
func (o *DetachLdapResourceFromEnvironmentsParams) SetName(name string) {
	o.Name = name
}

// WithWorkspaceID adds the workspaceID to the detach ldap resource from environments params
func (o *DetachLdapResourceFromEnvironmentsParams) WithWorkspaceID(workspaceID int64) *DetachLdapResourceFromEnvironmentsParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the detach ldap resource from environments params
func (o *DetachLdapResourceFromEnvironmentsParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *DetachLdapResourceFromEnvironmentsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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