// Code generated by go-swagger; DO NOT EDIT.

package users

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

	"github.com/hortonworks/cb-cli/dataplane/oauthapi/model"
)

// NewAssignRolesToUserParams creates a new AssignRolesToUserParams object
// with the default values initialized.
func NewAssignRolesToUserParams() *AssignRolesToUserParams {
	var ()
	return &AssignRolesToUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAssignRolesToUserParamsWithTimeout creates a new AssignRolesToUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAssignRolesToUserParamsWithTimeout(timeout time.Duration) *AssignRolesToUserParams {
	var ()
	return &AssignRolesToUserParams{

		timeout: timeout,
	}
}

// NewAssignRolesToUserParamsWithContext creates a new AssignRolesToUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewAssignRolesToUserParamsWithContext(ctx context.Context) *AssignRolesToUserParams {
	var ()
	return &AssignRolesToUserParams{

		Context: ctx,
	}
}

// NewAssignRolesToUserParamsWithHTTPClient creates a new AssignRolesToUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAssignRolesToUserParamsWithHTTPClient(client *http.Client) *AssignRolesToUserParams {
	var ()
	return &AssignRolesToUserParams{
		HTTPClient: client,
	}
}

/*AssignRolesToUserParams contains all the parameters to send to the API endpoint
for the assign roles to user operation typically these are written to a http.Request
*/
type AssignRolesToUserParams struct {

	/*RoleIds*/
	RoleIds *model.RolesInput
	/*UserID
	  user IDs

	*/
	UserID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the assign roles to user params
func (o *AssignRolesToUserParams) WithTimeout(timeout time.Duration) *AssignRolesToUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign roles to user params
func (o *AssignRolesToUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign roles to user params
func (o *AssignRolesToUserParams) WithContext(ctx context.Context) *AssignRolesToUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign roles to user params
func (o *AssignRolesToUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign roles to user params
func (o *AssignRolesToUserParams) WithHTTPClient(client *http.Client) *AssignRolesToUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign roles to user params
func (o *AssignRolesToUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithRoleIds adds the roleIds to the assign roles to user params
func (o *AssignRolesToUserParams) WithRoleIds(roleIds *model.RolesInput) *AssignRolesToUserParams {
	o.SetRoleIds(roleIds)
	return o
}

// SetRoleIds adds the roleIds to the assign roles to user params
func (o *AssignRolesToUserParams) SetRoleIds(roleIds *model.RolesInput) {
	o.RoleIds = roleIds
}

// WithUserID adds the userID to the assign roles to user params
func (o *AssignRolesToUserParams) WithUserID(userID string) *AssignRolesToUserParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the assign roles to user params
func (o *AssignRolesToUserParams) SetUserID(userID string) {
	o.UserID = userID
}

// WriteToRequest writes these params to a swagger request
func (o *AssignRolesToUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.RoleIds == nil {
		o.RoleIds = new(model.RolesInput)
	}

	if err := r.SetBodyParam(o.RoleIds); err != nil {
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