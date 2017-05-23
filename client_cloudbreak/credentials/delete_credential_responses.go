package credentials

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-swagger/go-swagger/client"
	"github.com/go-swagger/go-swagger/httpkit"

	strfmt "github.com/go-swagger/go-swagger/strfmt"
)

// DeleteCredentialReader is a Reader for the DeleteCredential structure.
type DeleteCredentialReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the recieved o.
func (o *DeleteCredentialReader) ReadResponse(response client.Response, consumer httpkit.Consumer) (interface{}, error) {
	switch response.Code() {

	default:
		result := NewDeleteCredentialDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	}
}

// NewDeleteCredentialDefault creates a DeleteCredentialDefault with default headers values
func NewDeleteCredentialDefault(code int) *DeleteCredentialDefault {
	return &DeleteCredentialDefault{
		_statusCode: code,
	}
}

/*DeleteCredentialDefault handles this case with default header values.

successful operation
*/
type DeleteCredentialDefault struct {
	_statusCode int
}

// Code gets the status code for the delete credential default response
func (o *DeleteCredentialDefault) Code() int {
	return o._statusCode
}

func (o *DeleteCredentialDefault) Error() string {
	return fmt.Sprintf("[DELETE /credentials/{id}][%d] deleteCredential default ", o._statusCode)
}

func (o *DeleteCredentialDefault) readResponse(response client.Response, consumer httpkit.Consumer, formats strfmt.Registry) error {

	return nil
}