package recipes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// NewGetPrivateRecipeParams creates a new GetPrivateRecipeParams object
// with the default values initialized.
func NewGetPrivateRecipeParams() *GetPrivateRecipeParams {
	var ()
	return &GetPrivateRecipeParams{}
}

/*GetPrivateRecipeParams contains all the parameters to send to the API endpoint
for the get private recipe operation typically these are written to a http.Request
*/
type GetPrivateRecipeParams struct {

	/*Name*/
	Name string
}

// WithName adds the name to the get private recipe params
func (o *GetPrivateRecipeParams) WithName(name string) *GetPrivateRecipeParams {
	o.Name = name
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *GetPrivateRecipeParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	// path param name
	if err := r.SetPathParam("name", o.Name); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}