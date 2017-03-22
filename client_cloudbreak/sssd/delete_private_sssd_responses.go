package sssd

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// DeletePrivateSssdReader is a Reader for the DeletePrivateSssd structure.
type DeletePrivateSssdReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeletePrivateSssdReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewDeletePrivateSssdDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeletePrivateSssdDefault creates a DeletePrivateSssdDefault with default headers values
func NewDeletePrivateSssdDefault(code int) *DeletePrivateSssdDefault {
	return &DeletePrivateSssdDefault{
		_statusCode: code,
	}
}

/*DeletePrivateSssdDefault handles this case with default header values.

successful operation
*/
type DeletePrivateSssdDefault struct {
	_statusCode int
}

// Code gets the status code for the delete private sssd default response
func (o *DeletePrivateSssdDefault) Code() int {
	return o._statusCode
}

func (o *DeletePrivateSssdDefault) Error() string {
	return fmt.Sprintf("[DELETE /sssd/user/{name}][%d] deletePrivateSssd default ", o._statusCode)
}

func (o *DeletePrivateSssdDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}