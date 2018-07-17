// Code generated by go-swagger; DO NOT EDIT.

package auth_tokens

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/models"
)

// CreateAuthTokenReader is a Reader for the CreateAuthToken structure.
type CreateAuthTokenReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAuthTokenReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreateAuthTokenOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewCreateAuthTokenUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAuthTokenOK creates a CreateAuthTokenOK with default headers values
func NewCreateAuthTokenOK() *CreateAuthTokenOK {
	return &CreateAuthTokenOK{}
}

/*CreateAuthTokenOK handles this case with default header values.

Success
*/
type CreateAuthTokenOK struct {
	Payload *models.V4CreateAuthTokenResponse
}

func (o *CreateAuthTokenOK) Error() string {
	return fmt.Sprintf("[POST /v4/auth-tokens/][%d] createAuthTokenOK  %+v", 200, o.Payload)
}

func (o *CreateAuthTokenOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4CreateAuthTokenResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAuthTokenUnauthorized creates a CreateAuthTokenUnauthorized with default headers values
func NewCreateAuthTokenUnauthorized() *CreateAuthTokenUnauthorized {
	return &CreateAuthTokenUnauthorized{}
}

/*CreateAuthTokenUnauthorized handles this case with default header values.

Permission denied
*/
type CreateAuthTokenUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *CreateAuthTokenUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v4/auth-tokens/][%d] createAuthTokenUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAuthTokenUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
