// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_proxies

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

// NewCreateProxyconfigInWorkspaceParams creates a new CreateProxyconfigInWorkspaceParams object
// with the default values initialized.
func NewCreateProxyconfigInWorkspaceParams() *CreateProxyconfigInWorkspaceParams {
	var ()
	return &CreateProxyconfigInWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProxyconfigInWorkspaceParamsWithTimeout creates a new CreateProxyconfigInWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateProxyconfigInWorkspaceParamsWithTimeout(timeout time.Duration) *CreateProxyconfigInWorkspaceParams {
	var ()
	return &CreateProxyconfigInWorkspaceParams{

		timeout: timeout,
	}
}

// NewCreateProxyconfigInWorkspaceParamsWithContext creates a new CreateProxyconfigInWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateProxyconfigInWorkspaceParamsWithContext(ctx context.Context) *CreateProxyconfigInWorkspaceParams {
	var ()
	return &CreateProxyconfigInWorkspaceParams{

		Context: ctx,
	}
}

// NewCreateProxyconfigInWorkspaceParamsWithHTTPClient creates a new CreateProxyconfigInWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateProxyconfigInWorkspaceParamsWithHTTPClient(client *http.Client) *CreateProxyconfigInWorkspaceParams {
	var ()
	return &CreateProxyconfigInWorkspaceParams{
		HTTPClient: client,
	}
}

/*CreateProxyconfigInWorkspaceParams contains all the parameters to send to the API endpoint
for the create proxyconfig in workspace operation typically these are written to a http.Request
*/
type CreateProxyconfigInWorkspaceParams struct {

	/*Body*/
	Body *model.ProxyV4Request
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create proxyconfig in workspace params
func (o *CreateProxyconfigInWorkspaceParams) WithTimeout(timeout time.Duration) *CreateProxyconfigInWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create proxyconfig in workspace params
func (o *CreateProxyconfigInWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create proxyconfig in workspace params
func (o *CreateProxyconfigInWorkspaceParams) WithContext(ctx context.Context) *CreateProxyconfigInWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create proxyconfig in workspace params
func (o *CreateProxyconfigInWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create proxyconfig in workspace params
func (o *CreateProxyconfigInWorkspaceParams) WithHTTPClient(client *http.Client) *CreateProxyconfigInWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create proxyconfig in workspace params
func (o *CreateProxyconfigInWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create proxyconfig in workspace params
func (o *CreateProxyconfigInWorkspaceParams) WithBody(body *model.ProxyV4Request) *CreateProxyconfigInWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create proxyconfig in workspace params
func (o *CreateProxyconfigInWorkspaceParams) SetBody(body *model.ProxyV4Request) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the create proxyconfig in workspace params
func (o *CreateProxyconfigInWorkspaceParams) WithWorkspaceID(workspaceID int64) *CreateProxyconfigInWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the create proxyconfig in workspace params
func (o *CreateProxyconfigInWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProxyconfigInWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
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
