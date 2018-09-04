// Code generated by go-swagger; DO NOT EDIT.

package v3_organization_id_imagecatalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// PutPublicImageCatalogInOrganizationReader is a Reader for the PutPublicImageCatalogInOrganization structure.
type PutPublicImageCatalogInOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutPublicImageCatalogInOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPutPublicImageCatalogInOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutPublicImageCatalogInOrganizationOK creates a PutPublicImageCatalogInOrganizationOK with default headers values
func NewPutPublicImageCatalogInOrganizationOK() *PutPublicImageCatalogInOrganizationOK {
	return &PutPublicImageCatalogInOrganizationOK{}
}

/*PutPublicImageCatalogInOrganizationOK handles this case with default header values.

successful operation
*/
type PutPublicImageCatalogInOrganizationOK struct {
	Payload *models_cloudbreak.ImageCatalogResponse
}

func (o *PutPublicImageCatalogInOrganizationOK) Error() string {
	return fmt.Sprintf("[PUT /v3/{organizationId}/imagecatalogs][%d] putPublicImageCatalogInOrganizationOK  %+v", 200, o.Payload)
}

func (o *PutPublicImageCatalogInOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.ImageCatalogResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}