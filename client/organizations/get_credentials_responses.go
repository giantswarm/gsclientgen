// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/v2/models"
)

// GetCredentialsReader is a Reader for the GetCredentials structure.
type GetCredentialsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetCredentialsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetCredentialsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetCredentialsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetCredentialsOK creates a GetCredentialsOK with default headers values
func NewGetCredentialsOK() *GetCredentialsOK {
	return &GetCredentialsOK{}
}

/*GetCredentialsOK handles this case with default header values.

Credentials
*/
type GetCredentialsOK struct {
	Payload models.V4GetCredentialsResponse
}

func (o *GetCredentialsOK) Error() string {
	return fmt.Sprintf("[GET /v4/organizations/{organization_id}/credentials/][%d] getCredentialsOK  %+v", 200, o.Payload)
}

func (o *GetCredentialsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetCredentialsDefault creates a GetCredentialsDefault with default headers values
func NewGetCredentialsDefault(code int) *GetCredentialsDefault {
	return &GetCredentialsDefault{
		_statusCode: code,
	}
}

/*GetCredentialsDefault handles this case with default header values.

error
*/
type GetCredentialsDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the get credentials default response
func (o *GetCredentialsDefault) Code() int {
	return o._statusCode
}

func (o *GetCredentialsDefault) Error() string {
	return fmt.Sprintf("[GET /v4/organizations/{organization_id}/credentials/][%d] getCredentials default  %+v", o._statusCode, o.Payload)
}

func (o *GetCredentialsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
