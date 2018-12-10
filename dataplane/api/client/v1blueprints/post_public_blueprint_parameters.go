// Code generated by go-swagger; DO NOT EDIT.

package v1blueprints

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

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// NewPostPublicBlueprintParams creates a new PostPublicBlueprintParams object
// with the default values initialized.
func NewPostPublicBlueprintParams() *PostPublicBlueprintParams {
	var ()
	return &PostPublicBlueprintParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostPublicBlueprintParamsWithTimeout creates a new PostPublicBlueprintParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostPublicBlueprintParamsWithTimeout(timeout time.Duration) *PostPublicBlueprintParams {
	var ()
	return &PostPublicBlueprintParams{

		timeout: timeout,
	}
}

// NewPostPublicBlueprintParamsWithContext creates a new PostPublicBlueprintParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostPublicBlueprintParamsWithContext(ctx context.Context) *PostPublicBlueprintParams {
	var ()
	return &PostPublicBlueprintParams{

		Context: ctx,
	}
}

// NewPostPublicBlueprintParamsWithHTTPClient creates a new PostPublicBlueprintParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostPublicBlueprintParamsWithHTTPClient(client *http.Client) *PostPublicBlueprintParams {
	var ()
	return &PostPublicBlueprintParams{
		HTTPClient: client,
	}
}

/*PostPublicBlueprintParams contains all the parameters to send to the API endpoint
for the post public blueprint operation typically these are written to a http.Request
*/
type PostPublicBlueprintParams struct {

	/*Body*/
	Body *model.BlueprintRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post public blueprint params
func (o *PostPublicBlueprintParams) WithTimeout(timeout time.Duration) *PostPublicBlueprintParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post public blueprint params
func (o *PostPublicBlueprintParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post public blueprint params
func (o *PostPublicBlueprintParams) WithContext(ctx context.Context) *PostPublicBlueprintParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post public blueprint params
func (o *PostPublicBlueprintParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post public blueprint params
func (o *PostPublicBlueprintParams) WithHTTPClient(client *http.Client) *PostPublicBlueprintParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post public blueprint params
func (o *PostPublicBlueprintParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post public blueprint params
func (o *PostPublicBlueprintParams) WithBody(body *model.BlueprintRequest) *PostPublicBlueprintParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post public blueprint params
func (o *PostPublicBlueprintParams) SetBody(body *model.BlueprintRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *PostPublicBlueprintParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}