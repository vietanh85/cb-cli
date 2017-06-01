package util

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"
	"github.com/go-swagger/go-swagger/swag"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewTestLdapConnectionByIDUtilParams creates a new TestLdapConnectionByIDUtilParams object
// with the default values initialized.
func NewTestLdapConnectionByIDUtilParams() *TestLdapConnectionByIDUtilParams {
	var ()
	return &TestLdapConnectionByIDUtilParams{}
}

/*TestLdapConnectionByIDUtilParams contains all the parameters to send to the API endpoint
for the test ldap connection by id util operation typically these are written to a http.Request
*/
type TestLdapConnectionByIDUtilParams struct {

	/*ID*/
	ID int64
}

// WithID adds the id to the test ldap connection by id util params
func (o *TestLdapConnectionByIDUtilParams) WithID(id int64) *TestLdapConnectionByIDUtilParams {
	o.ID = id
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *TestLdapConnectionByIDUtilParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

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