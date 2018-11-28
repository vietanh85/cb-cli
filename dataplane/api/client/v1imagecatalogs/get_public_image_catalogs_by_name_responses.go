// Code generated by go-swagger; DO NOT EDIT.

package v1imagecatalogs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/dataplane/api/model"
)

// GetPublicImageCatalogsByNameReader is a Reader for the GetPublicImageCatalogsByName structure.
type GetPublicImageCatalogsByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPublicImageCatalogsByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPublicImageCatalogsByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPublicImageCatalogsByNameOK creates a GetPublicImageCatalogsByNameOK with default headers values
func NewGetPublicImageCatalogsByNameOK() *GetPublicImageCatalogsByNameOK {
	return &GetPublicImageCatalogsByNameOK{}
}

/*GetPublicImageCatalogsByNameOK handles this case with default header values.

successful operation
*/
type GetPublicImageCatalogsByNameOK struct {
	Payload *model.ImageCatalogResponse
}

func (o *GetPublicImageCatalogsByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/imagecatalogs/account/{name}][%d] getPublicImageCatalogsByNameOK  %+v", 200, o.Payload)
}

func (o *GetPublicImageCatalogsByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.ImageCatalogResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}