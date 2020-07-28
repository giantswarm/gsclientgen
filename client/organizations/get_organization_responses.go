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

// GetOrganizationReader is a Reader for the GetOrganization structure.
type GetOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetOrganizationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetOrganizationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetOrganizationDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetOrganizationOK creates a GetOrganizationOK with default headers values
func NewGetOrganizationOK() *GetOrganizationOK {
	return &GetOrganizationOK{}
}

/*GetOrganizationOK handles this case with default header values.

Organization details
*/
type GetOrganizationOK struct {
	Payload *models.V4Organization
}

func (o *GetOrganizationOK) Error() string {
	return fmt.Sprintf("[GET /v4/organizations/{organization_id}/][%d] getOrganizationOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4Organization)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationUnauthorized creates a GetOrganizationUnauthorized with default headers values
func NewGetOrganizationUnauthorized() *GetOrganizationUnauthorized {
	return &GetOrganizationUnauthorized{}
}

/*GetOrganizationUnauthorized handles this case with default header values.

Permission denied
*/
type GetOrganizationUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *GetOrganizationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v4/organizations/{organization_id}/][%d] getOrganizationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetOrganizationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationNotFound creates a GetOrganizationNotFound with default headers values
func NewGetOrganizationNotFound() *GetOrganizationNotFound {
	return &GetOrganizationNotFound{}
}

/*GetOrganizationNotFound handles this case with default header values.

Organization not found
*/
type GetOrganizationNotFound struct {
	Payload *models.V4GenericResponse
}

func (o *GetOrganizationNotFound) Error() string {
	return fmt.Sprintf("[GET /v4/organizations/{organization_id}/][%d] getOrganizationNotFound  %+v", 404, o.Payload)
}

func (o *GetOrganizationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetOrganizationDefault creates a GetOrganizationDefault with default headers values
func NewGetOrganizationDefault(code int) *GetOrganizationDefault {
	return &GetOrganizationDefault{
		_statusCode: code,
	}
}

/*GetOrganizationDefault handles this case with default header values.

Error
*/
type GetOrganizationDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the get organization default response
func (o *GetOrganizationDefault) Code() int {
	return o._statusCode
}

func (o *GetOrganizationDefault) Error() string {
	return fmt.Sprintf("[GET /v4/organizations/{organization_id}/][%d] getOrganization default  %+v", o._statusCode, o.Payload)
}

func (o *GetOrganizationDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
