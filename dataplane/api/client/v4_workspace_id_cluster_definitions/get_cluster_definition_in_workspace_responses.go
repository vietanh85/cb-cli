// Code generated by go-swagger; DO NOT EDIT.

package v4_workspace_id_cluster_definitions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	model "github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetClusterDefinitionInWorkspaceReader is a Reader for the GetClusterDefinitionInWorkspace structure.
type GetClusterDefinitionInWorkspaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterDefinitionInWorkspaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterDefinitionInWorkspaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetClusterDefinitionInWorkspaceOK creates a GetClusterDefinitionInWorkspaceOK with default headers values
func NewGetClusterDefinitionInWorkspaceOK() *GetClusterDefinitionInWorkspaceOK {
	return &GetClusterDefinitionInWorkspaceOK{}
}

/*GetClusterDefinitionInWorkspaceOK handles this case with default header values.

successful operation
*/
type GetClusterDefinitionInWorkspaceOK struct {
	Payload *model.ClusterDefinitionV4Response
}

func (o *GetClusterDefinitionInWorkspaceOK) Error() string {
	return fmt.Sprintf("[GET /v4/{workspaceId}/cluster_definitions/{name}][%d] getClusterDefinitionInWorkspaceOK  %+v", 200, o.Payload)
}

func (o *GetClusterDefinitionInWorkspaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ClusterDefinitionV4Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
