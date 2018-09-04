// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_ldapconfigs

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

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// NewCreateLdapConfigsInOrganizationParams creates a new CreateLdapConfigsInOrganizationParams object
// with the default values initialized.
func NewCreateLdapConfigsInOrganizationParams() *CreateLdapConfigsInOrganizationParams {
	var ()
	return &CreateLdapConfigsInOrganizationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateLdapConfigsInOrganizationParamsWithTimeout creates a new CreateLdapConfigsInOrganizationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateLdapConfigsInOrganizationParamsWithTimeout(timeout time.Duration) *CreateLdapConfigsInOrganizationParams {
	var ()
	return &CreateLdapConfigsInOrganizationParams{

		timeout: timeout,
	}
}

// NewCreateLdapConfigsInOrganizationParamsWithContext creates a new CreateLdapConfigsInOrganizationParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateLdapConfigsInOrganizationParamsWithContext(ctx context.Context) *CreateLdapConfigsInOrganizationParams {
	var ()
	return &CreateLdapConfigsInOrganizationParams{

		Context: ctx,
	}
}

// NewCreateLdapConfigsInOrganizationParamsWithHTTPClient creates a new CreateLdapConfigsInOrganizationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateLdapConfigsInOrganizationParamsWithHTTPClient(client *http.Client) *CreateLdapConfigsInOrganizationParams {
	var ()
	return &CreateLdapConfigsInOrganizationParams{
		HTTPClient: client,
	}
}

/*CreateLdapConfigsInOrganizationParams contains all the parameters to send to the API endpoint
for the create ldap configs in organization operation typically these are written to a http.Request
*/
type CreateLdapConfigsInOrganizationParams struct {

	/*Body*/
	Body *models_cloudbreak.LdapConfigRequest
	/*OrganizationID*/
	OrganizationID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create ldap configs in organization params
func (o *CreateLdapConfigsInOrganizationParams) WithTimeout(timeout time.Duration) *CreateLdapConfigsInOrganizationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create ldap configs in organization params
func (o *CreateLdapConfigsInOrganizationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create ldap configs in organization params
func (o *CreateLdapConfigsInOrganizationParams) WithContext(ctx context.Context) *CreateLdapConfigsInOrganizationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create ldap configs in organization params
func (o *CreateLdapConfigsInOrganizationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create ldap configs in organization params
func (o *CreateLdapConfigsInOrganizationParams) WithHTTPClient(client *http.Client) *CreateLdapConfigsInOrganizationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create ldap configs in organization params
func (o *CreateLdapConfigsInOrganizationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create ldap configs in organization params
func (o *CreateLdapConfigsInOrganizationParams) WithBody(body *models_cloudbreak.LdapConfigRequest) *CreateLdapConfigsInOrganizationParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create ldap configs in organization params
func (o *CreateLdapConfigsInOrganizationParams) SetBody(body *models_cloudbreak.LdapConfigRequest) {
	o.Body = body
}

// WithOrganizationID adds the organizationID to the create ldap configs in organization params
func (o *CreateLdapConfigsInOrganizationParams) WithOrganizationID(organizationID int64) *CreateLdapConfigsInOrganizationParams {
	o.SetOrganizationID(organizationID)
	return o
}

// SetOrganizationID adds the organizationId to the create ldap configs in organization params
func (o *CreateLdapConfigsInOrganizationParams) SetOrganizationID(organizationID int64) {
	o.OrganizationID = organizationID
}

// WriteToRequest writes these params to a swagger request
func (o *CreateLdapConfigsInOrganizationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.LdapConfigRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// path param organizationId
	if err := r.SetPathParam("organizationId", swag.FormatInt64(o.OrganizationID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}