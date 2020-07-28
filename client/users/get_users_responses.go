// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/v2/models"
)

// GetUsersReader is a Reader for the GetUsers structure.
type GetUsersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUsersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetUsersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetUsersUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetUsersDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetUsersOK creates a GetUsersOK with default headers values
func NewGetUsersOK() *GetUsersOK {
	return &GetUsersOK{}
}

/*GetUsersOK handles this case with default header values.

Success
*/
type GetUsersOK struct {
	Payload []*models.V4UserListItem
}

func (o *GetUsersOK) Error() string {
	return fmt.Sprintf("[GET /v4/users/][%d] getUsersOK  %+v", 200, o.Payload)
}

func (o *GetUsersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersUnauthorized creates a GetUsersUnauthorized with default headers values
func NewGetUsersUnauthorized() *GetUsersUnauthorized {
	return &GetUsersUnauthorized{}
}

/*GetUsersUnauthorized handles this case with default header values.

Permission denied
*/
type GetUsersUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *GetUsersUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v4/users/][%d] getUsersUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUsersUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUsersDefault creates a GetUsersDefault with default headers values
func NewGetUsersDefault(code int) *GetUsersDefault {
	return &GetUsersDefault{
		_statusCode: code,
	}
}

/*GetUsersDefault handles this case with default header values.

Error
*/
type GetUsersDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the get users default response
func (o *GetUsersDefault) Code() int {
	return o._statusCode
}

func (o *GetUsersDefault) Error() string {
	return fmt.Sprintf("[GET /v4/users/][%d] getUsers default  %+v", o._statusCode, o.Payload)
}

func (o *GetUsersDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
