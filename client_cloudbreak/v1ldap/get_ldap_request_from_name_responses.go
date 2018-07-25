// Code generated by go-swagger; DO NOT EDIT.

package v1ldap

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetLdapRequestFromNameReader is a Reader for the GetLdapRequestFromName structure.
type GetLdapRequestFromNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetLdapRequestFromNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetLdapRequestFromNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetLdapRequestFromNameOK creates a GetLdapRequestFromNameOK with default headers values
func NewGetLdapRequestFromNameOK() *GetLdapRequestFromNameOK {
	return &GetLdapRequestFromNameOK{}
}

/*GetLdapRequestFromNameOK handles this case with default header values.

successful operation
*/
type GetLdapRequestFromNameOK struct {
	Payload *models_cloudbreak.LdapConfigRequest
}

func (o *GetLdapRequestFromNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/ldap/{name}/request][%d] getLdapRequestFromNameOK  %+v", 200, o.Payload)
}

func (o *GetLdapRequestFromNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.LdapConfigRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
