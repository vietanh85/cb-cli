package sssd

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/cb-cli/models_cloudbreak"
)

// GetPrivateSssdReader is a Reader for the GetPrivateSssd structure.
type GetPrivateSssdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPrivateSssdReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrivateSssdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrivateSssdOK creates a GetPrivateSssdOK with default headers values
func NewGetPrivateSssdOK() *GetPrivateSssdOK {
	return &GetPrivateSssdOK{}
}

/*GetPrivateSssdOK handles this case with default header values.

successful operation
*/
type GetPrivateSssdOK struct {
	Payload *models_cloudbreak.SssdConfigResponse
}

func (o *GetPrivateSssdOK) Error() string {
	return fmt.Sprintf("[GET /sssd/user/{name}][%d] getPrivateSssdOK  %+v", 200, o.Payload)
}

func (o *GetPrivateSssdOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.SssdConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
