package configurations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/hdc-cli/models_autoscale"
)

// GetScalingConfigurationReader is a Reader for the GetScalingConfiguration structure.
type GetScalingConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetScalingConfigurationReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetScalingConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetScalingConfigurationOK creates a GetScalingConfigurationOK with default headers values
func NewGetScalingConfigurationOK() *GetScalingConfigurationOK {
	return &GetScalingConfigurationOK{}
}

/*GetScalingConfigurationOK handles this case with default header values.

successful operation
*/
type GetScalingConfigurationOK struct {
	Payload *models_autoscale.ScalingConfiguration
}

func (o *GetScalingConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{clusterId}/configurations/scaling][%d] getScalingConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetScalingConfigurationOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models_autoscale.ScalingConfiguration)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
