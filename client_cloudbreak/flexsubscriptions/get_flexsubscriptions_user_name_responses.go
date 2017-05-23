package flexsubscriptions

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_cloudbreak"
)

// GetFlexsubscriptionsUserNameReader is a Reader for the GetFlexsubscriptionsUserName structure.
type GetFlexsubscriptionsUserNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetFlexsubscriptionsUserNameReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetFlexsubscriptionsUserNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetFlexsubscriptionsUserNameOK creates a GetFlexsubscriptionsUserNameOK with default headers values
func NewGetFlexsubscriptionsUserNameOK() *GetFlexsubscriptionsUserNameOK {
	return &GetFlexsubscriptionsUserNameOK{}
}

/*GetFlexsubscriptionsUserNameOK handles this case with default header values.

successful operation
*/
type GetFlexsubscriptionsUserNameOK struct {
	Payload *models_cloudbreak.FlexSubscriptionResponse
}

func (o *GetFlexsubscriptionsUserNameOK) Error() string {
	return fmt.Sprintf("[GET /flexsubscriptions/user/{name}][%d] getFlexsubscriptionsUserNameOK  %+v", 200, o.Payload)
}

func (o *GetFlexsubscriptionsUserNameOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_cloudbreak.FlexSubscriptionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}