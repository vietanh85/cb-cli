// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_flexsubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// DeleteFlexSubscriptionInOrganizationReader is a Reader for the DeleteFlexSubscriptionInOrganization structure.
type DeleteFlexSubscriptionInOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteFlexSubscriptionInOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteFlexSubscriptionInOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteFlexSubscriptionInOrganizationOK creates a DeleteFlexSubscriptionInOrganizationOK with default headers values
func NewDeleteFlexSubscriptionInOrganizationOK() *DeleteFlexSubscriptionInOrganizationOK {
	return &DeleteFlexSubscriptionInOrganizationOK{}
}

/*DeleteFlexSubscriptionInOrganizationOK handles this case with default header values.

successful operation
*/
type DeleteFlexSubscriptionInOrganizationOK struct {
	Payload *models_cloudbreak.FlexSubscriptionResponse
}

func (o *DeleteFlexSubscriptionInOrganizationOK) Error() string {
	return fmt.Sprintf("[DELETE /v3/{organizationId}/flexsubscriptions/{name}][%d] deleteFlexSubscriptionInOrganizationOK  %+v", 200, o.Payload)
}

func (o *DeleteFlexSubscriptionInOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.FlexSubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}