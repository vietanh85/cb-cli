package constraints

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/errors"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// NewPostConstraintsAccountParams creates a new PostConstraintsAccountParams object
// with the default values initialized.
func NewPostConstraintsAccountParams() *PostConstraintsAccountParams {
	var ()
	return &PostConstraintsAccountParams{}
}

/*PostConstraintsAccountParams contains all the parameters to send to the API endpoint
for the post constraints account operation typically these are written to a http.Request
*/
type PostConstraintsAccountParams struct {

	/*Body*/
	Body *models_cloudbreak.ConstraintTemplateRequest
}

// WithBody adds the body to the post constraints account params
func (o *PostConstraintsAccountParams) WithBody(body *models_cloudbreak.ConstraintTemplateRequest) *PostConstraintsAccountParams {
	o.Body = body
	return o
}

// WriteToRequest writes these params to a swagger request
func (o *PostConstraintsAccountParams) WriteToRequest(r client.Request, reg strfmt.Registry) error {

	var res []error

	if o.Body == nil {
		o.Body = new(models_cloudbreak.ConstraintTemplateRequest)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}