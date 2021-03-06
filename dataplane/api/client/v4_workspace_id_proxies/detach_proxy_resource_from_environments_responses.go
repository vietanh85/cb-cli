// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_proxies

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// DetachProxyResourceFromEnvironmentsReader is a Reader for the DetachProxyResourceFromEnvironments structure.
type DetachProxyResourceFromEnvironmentsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetachProxyResourceFromEnvironmentsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDetachProxyResourceFromEnvironmentsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDetachProxyResourceFromEnvironmentsOK creates a DetachProxyResourceFromEnvironmentsOK with default headers values
func NewDetachProxyResourceFromEnvironmentsOK() *DetachProxyResourceFromEnvironmentsOK {
	return &DetachProxyResourceFromEnvironmentsOK{}
}

/*DetachProxyResourceFromEnvironmentsOK handles this case with default header values.

successful operation
*/
type DetachProxyResourceFromEnvironmentsOK struct {
	Payload *model.ProxyV4Response
}

func (o *DetachProxyResourceFromEnvironmentsOK) Error() string {
	return fmt.Sprintf("[PUT /v4/{workspaceId}/proxies/{name}/detach][%d] detachProxyResourceFromEnvironmentsOK  %+v", 200, o.Payload)
}

func (o *DetachProxyResourceFromEnvironmentsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ProxyV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
