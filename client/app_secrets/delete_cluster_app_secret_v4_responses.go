// Code generated by go-swagger; DO NOT EDIT.

package app_secrets

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/models"
)

// DeleteClusterAppSecretV4Reader is a Reader for the DeleteClusterAppSecretV4 structure.
type DeleteClusterAppSecretV4Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteClusterAppSecretV4Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteClusterAppSecretV4OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewDeleteClusterAppSecretV4Unauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewDeleteClusterAppSecretV4NotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewDeleteClusterAppSecretV4Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteClusterAppSecretV4OK creates a DeleteClusterAppSecretV4OK with default headers values
func NewDeleteClusterAppSecretV4OK() *DeleteClusterAppSecretV4OK {
	return &DeleteClusterAppSecretV4OK{}
}

/*DeleteClusterAppSecretV4OK handles this case with default header values.

Secret deleted
*/
type DeleteClusterAppSecretV4OK struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteClusterAppSecretV4OK) Error() string {
	return fmt.Sprintf("[DELETE /v4/clusters/{cluster_id}/apps/{app_name}/secret/][%d] deleteClusterAppSecretV4OK  %+v", 200, o.Payload)
}

func (o *DeleteClusterAppSecretV4OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterAppSecretV4Unauthorized creates a DeleteClusterAppSecretV4Unauthorized with default headers values
func NewDeleteClusterAppSecretV4Unauthorized() *DeleteClusterAppSecretV4Unauthorized {
	return &DeleteClusterAppSecretV4Unauthorized{}
}

/*DeleteClusterAppSecretV4Unauthorized handles this case with default header values.

Permission denied
*/
type DeleteClusterAppSecretV4Unauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteClusterAppSecretV4Unauthorized) Error() string {
	return fmt.Sprintf("[DELETE /v4/clusters/{cluster_id}/apps/{app_name}/secret/][%d] deleteClusterAppSecretV4Unauthorized  %+v", 401, o.Payload)
}

func (o *DeleteClusterAppSecretV4Unauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterAppSecretV4NotFound creates a DeleteClusterAppSecretV4NotFound with default headers values
func NewDeleteClusterAppSecretV4NotFound() *DeleteClusterAppSecretV4NotFound {
	return &DeleteClusterAppSecretV4NotFound{}
}

/*DeleteClusterAppSecretV4NotFound handles this case with default header values.

Secret not found
*/
type DeleteClusterAppSecretV4NotFound struct {
	Payload *models.V4GenericResponse
}

func (o *DeleteClusterAppSecretV4NotFound) Error() string {
	return fmt.Sprintf("[DELETE /v4/clusters/{cluster_id}/apps/{app_name}/secret/][%d] deleteClusterAppSecretV4NotFound  %+v", 404, o.Payload)
}

func (o *DeleteClusterAppSecretV4NotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteClusterAppSecretV4Default creates a DeleteClusterAppSecretV4Default with default headers values
func NewDeleteClusterAppSecretV4Default(code int) *DeleteClusterAppSecretV4Default {
	return &DeleteClusterAppSecretV4Default{
		_statusCode: code,
	}
}

/*DeleteClusterAppSecretV4Default handles this case with default header values.

Error
*/
type DeleteClusterAppSecretV4Default struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the delete cluster app secret v4 default response
func (o *DeleteClusterAppSecretV4Default) Code() int {
	return o._statusCode
}

func (o *DeleteClusterAppSecretV4Default) Error() string {
	return fmt.Sprintf("[DELETE /v4/clusters/{cluster_id}/apps/{app_name}/secret/][%d] deleteClusterAppSecretV4 default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteClusterAppSecretV4Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}