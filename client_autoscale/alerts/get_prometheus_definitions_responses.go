package alerts

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"

	"github.com/hortonworks/cb-cli/models_autoscale"
)

// GetPrometheusDefinitionsReader is a Reader for the GetPrometheusDefinitions structure.
type GetPrometheusDefinitionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *GetPrometheusDefinitionsReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetPrometheusDefinitionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, client.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPrometheusDefinitionsOK creates a GetPrometheusDefinitionsOK with default headers values
func NewGetPrometheusDefinitionsOK() *GetPrometheusDefinitionsOK {
	return &GetPrometheusDefinitionsOK{}
}

/*GetPrometheusDefinitionsOK handles this case with default header values.

successful operation
*/
type GetPrometheusDefinitionsOK struct {
	Payload []*models_autoscale.AlertRuleDefinitionEntry
}

func (o *GetPrometheusDefinitionsOK) Error() string {
	return fmt.Sprintf("[GET /clusters/{clusterId}/alerts/prometheus/definitions][%d] getPrometheusDefinitionsOK  %+v", 200, o.Payload)
}

func (o *GetPrometheusDefinitionsOK) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
