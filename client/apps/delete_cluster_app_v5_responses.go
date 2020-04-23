// Code generated by go-swagger; DO NOT EDIT.

package apps

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/models"
)

// DeleteClusterAppV5Reader is a Reader for the DeleteClusterAppV5 structure.
type DeleteClusterAppV5Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterAppV5Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteClusterAppV5OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteClusterAppV5Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteClusterAppV5NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteClusterAppV5Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteClusterAppV5OK creates a DeleteClusterAppV5OK with default headers values
func NewDeleteClusterAppV5OK() *DeleteClusterAppV5OK {
	return &DeleteClusterAppV5OK{}
}

/*DeleteClusterAppV5OK handles this case with default header values.

App deleted
*/
type DeleteClusterAppV5OK struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteClusterAppV5OK) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/apps/{app_name}/][%d] deleteClusterAppV5OK  %+v", 200, o.Payload)
}

func (o *DeleteClusterAppV5OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterAppV5Unauthorized creates a DeleteClusterAppV5Unauthorized with default headers values
func NewDeleteClusterAppV5Unauthorized() *DeleteClusterAppV5Unauthorized {
	return &DeleteClusterAppV5Unauthorized{}
}

/*DeleteClusterAppV5Unauthorized handles this case with default header values.

Permission denied
*/
type DeleteClusterAppV5Unauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteClusterAppV5Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/apps/{app_name}/][%d] deleteClusterAppV5Unauthorized  %+v", 401, o.Payload)
}

func (o *DeleteClusterAppV5Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterAppV5NotFound creates a DeleteClusterAppV5NotFound with default headers values
func NewDeleteClusterAppV5NotFound() *DeleteClusterAppV5NotFound {
	return &DeleteClusterAppV5NotFound{}
}

/*DeleteClusterAppV5NotFound handles this case with default header values.

App not found
*/
type DeleteClusterAppV5NotFound struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteClusterAppV5NotFound) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/apps/{app_name}/][%d] deleteClusterAppV5NotFound  %+v", 404, o.Payload)
}

func (o *DeleteClusterAppV5NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterAppV5Default creates a DeleteClusterAppV5Default with default headers values
func NewDeleteClusterAppV5Default(code int) *DeleteClusterAppV5Default {
	return &DeleteClusterAppV5Default{
		_statusCode: code,
	}
}

/*DeleteClusterAppV5Default handles this case with default header values.

Error
*/
type DeleteClusterAppV5Default struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the delete cluster app v5 default response
func (o *DeleteClusterAppV5Default) Code() int {
	return o._statusCode
}

func (o *DeleteClusterAppV5Default) Error() string {
	return fmt.Sprintf("[DELETE /v5/clusters/{cluster_id}/apps/{app_name}/][%d] deleteClusterAppV5 default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteClusterAppV5Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}