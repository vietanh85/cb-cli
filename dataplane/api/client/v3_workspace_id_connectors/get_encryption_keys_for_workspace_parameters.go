// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_connectors

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

// NewGetEncryptionKeysForWorkspaceParams creates a new GetEncryptionKeysForWorkspaceParams object
// with the default values initialized.
func NewGetEncryptionKeysForWorkspaceParams() *GetEncryptionKeysForWorkspaceParams {
	var ()
	return &GetEncryptionKeysForWorkspaceParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetEncryptionKeysForWorkspaceParamsWithTimeout creates a new GetEncryptionKeysForWorkspaceParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetEncryptionKeysForWorkspaceParamsWithTimeout(timeout time.Duration) *GetEncryptionKeysForWorkspaceParams {
	var ()
	return &GetEncryptionKeysForWorkspaceParams{

		timeout: timeout,
	}
}

// NewGetEncryptionKeysForWorkspaceParamsWithContext creates a new GetEncryptionKeysForWorkspaceParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetEncryptionKeysForWorkspaceParamsWithContext(ctx context.Context) *GetEncryptionKeysForWorkspaceParams {
	var ()
	return &GetEncryptionKeysForWorkspaceParams{

		Context: ctx,
	}
}

// NewGetEncryptionKeysForWorkspaceParamsWithHTTPClient creates a new GetEncryptionKeysForWorkspaceParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetEncryptionKeysForWorkspaceParamsWithHTTPClient(client *http.Client) *GetEncryptionKeysForWorkspaceParams {
	var ()
	return &GetEncryptionKeysForWorkspaceParams{
		HTTPClient: client,
	}
}

/*GetEncryptionKeysForWorkspaceParams contains all the parameters to send to the API endpoint
for the get encryption keys for workspace operation typically these are written to a http.Request
*/
type GetEncryptionKeysForWorkspaceParams struct {

	/*Body*/
	Body *model.PlatformResourceRequestJSON
	/*WorkspaceID*/
	WorkspaceID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get encryption keys for workspace params
func (o *GetEncryptionKeysForWorkspaceParams) WithTimeout(timeout time.Duration) *GetEncryptionKeysForWorkspaceParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get encryption keys for workspace params
func (o *GetEncryptionKeysForWorkspaceParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get encryption keys for workspace params
func (o *GetEncryptionKeysForWorkspaceParams) WithContext(ctx context.Context) *GetEncryptionKeysForWorkspaceParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get encryption keys for workspace params
func (o *GetEncryptionKeysForWorkspaceParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get encryption keys for workspace params
func (o *GetEncryptionKeysForWorkspaceParams) WithHTTPClient(client *http.Client) *GetEncryptionKeysForWorkspaceParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get encryption keys for workspace params
func (o *GetEncryptionKeysForWorkspaceParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get encryption keys for workspace params
func (o *GetEncryptionKeysForWorkspaceParams) WithBody(body *model.PlatformResourceRequestJSON) *GetEncryptionKeysForWorkspaceParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get encryption keys for workspace params
func (o *GetEncryptionKeysForWorkspaceParams) SetBody(body *model.PlatformResourceRequestJSON) {
	o.Body = body
}

// WithWorkspaceID adds the workspaceID to the get encryption keys for workspace params
func (o *GetEncryptionKeysForWorkspaceParams) WithWorkspaceID(workspaceID int64) *GetEncryptionKeysForWorkspaceParams {
	o.SetWorkspaceID(workspaceID)
	return o
}

// SetWorkspaceID adds the workspaceId to the get encryption keys for workspace params
func (o *GetEncryptionKeysForWorkspaceParams) SetWorkspaceID(workspaceID int64) {
	o.WorkspaceID = workspaceID
}

// WriteToRequest writes these params to a swagger request
func (o *GetEncryptionKeysForWorkspaceParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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