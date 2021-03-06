// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_kerberos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetKerberosConfigInWorkspaceReader is a Reader for the GetKerberosConfigInWorkspace structure.
type GetKerberosConfigInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetKerberosConfigInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetKerberosConfigInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetKerberosConfigInWorkspaceOK creates a GetKerberosConfigInWorkspaceOK with default headers values
func NewGetKerberosConfigInWorkspaceOK() *GetKerberosConfigInWorkspaceOK {
	return &GetKerberosConfigInWorkspaceOK{}
}

/*GetKerberosConfigInWorkspaceOK handles this case with default header values.

successful operation
*/
type GetKerberosConfigInWorkspaceOK struct {
	Payload *model.KerberosV4Response
}

func (o *GetKerberosConfigInWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/kerberos/{name}][%d] getKerberosConfigInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *GetKerberosConfigInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.KerberosV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
