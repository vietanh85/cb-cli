// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/hortonworks/cb-cli/dataplane/oauthapi/model"
)

// GetAllUsersReader is a Reader for the GetAllUsers structure.
type GetAllUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAllUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAllUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewGetAllUsersBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetAllUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 403:
		result := NewGetAllUsersForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 500:
		result := NewGetAllUsersInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetAllUsersOK creates a GetAllUsersOK with default headers values
func NewGetAllUsersOK() *GetAllUsersOK {
	return &GetAllUsersOK{}
}

/*GetAllUsersOK handles this case with default header values.

successful operation
*/
type GetAllUsersOK struct {
	Payload *model.UserList
}

func (o *GetAllUsersOK) Error() string {
	return fmt.Sprintf("[GET /caas/api/users][%d] getAllUsersOK  %+v", 200, o.Payload)
}

func (o *GetAllUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(model.UserList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAllUsersBadRequest creates a GetAllUsersBadRequest with default headers values
func NewGetAllUsersBadRequest() *GetAllUsersBadRequest {
	return &GetAllUsersBadRequest{}
}

/*GetAllUsersBadRequest handles this case with default header values.

Bad Request
*/
type GetAllUsersBadRequest struct {
}

func (o *GetAllUsersBadRequest) Error() string {
	return fmt.Sprintf("[GET /caas/api/users][%d] getAllUsersBadRequest ", 400)
}

func (o *GetAllUsersBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUsersUnauthorized creates a GetAllUsersUnauthorized with default headers values
func NewGetAllUsersUnauthorized() *GetAllUsersUnauthorized {
	return &GetAllUsersUnauthorized{}
}

/*GetAllUsersUnauthorized handles this case with default header values.

Unauthorized
*/
type GetAllUsersUnauthorized struct {
}

func (o *GetAllUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /caas/api/users][%d] getAllUsersUnauthorized ", 401)
}

func (o *GetAllUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUsersForbidden creates a GetAllUsersForbidden with default headers values
func NewGetAllUsersForbidden() *GetAllUsersForbidden {
	return &GetAllUsersForbidden{}
}

/*GetAllUsersForbidden handles this case with default header values.

Forbidden
*/
type GetAllUsersForbidden struct {
}

func (o *GetAllUsersForbidden) Error() string {
	return fmt.Sprintf("[GET /caas/api/users][%d] getAllUsersForbidden ", 403)
}

func (o *GetAllUsersForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetAllUsersInternalServerError creates a GetAllUsersInternalServerError with default headers values
func NewGetAllUsersInternalServerError() *GetAllUsersInternalServerError {
	return &GetAllUsersInternalServerError{}
}

/*GetAllUsersInternalServerError handles this case with default header values.

Internal server error
*/
type GetAllUsersInternalServerError struct {
}

func (o *GetAllUsersInternalServerError) Error() string {
	return fmt.Sprintf("[GET /caas/api/users][%d] getAllUsersInternalServerError ", 500)
}

func (o *GetAllUsersInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}