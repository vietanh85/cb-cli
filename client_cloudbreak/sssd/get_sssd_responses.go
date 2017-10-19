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

// GetSssdReader is a Reader for the GetSssd structure.
type GetSssdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetSssdReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetSssdOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSssdOK creates a GetSssdOK with default headers values
func NewGetSssdOK() *GetSssdOK {
	return &GetSssdOK{}
}

/*GetSssdOK handles this case with default header values.

successful operation
*/
type GetSssdOK struct {
	Payload *models_cloudbreak.SssdConfigResponse
}

func (o *GetSssdOK) Error() string {
	return fmt.Sprintf("[GET /sssd/{id}][%d] getSssdOK  %+v", 200, o.Payload)
}

func (o *GetSssdOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.SssdConfigResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
