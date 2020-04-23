// Code generated by go-swagger; DO NOT EDIT.

package node_pools

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/models"
)

// GetNodePoolReader is a Reader for the GetNodePool structure.
type GetNodePoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNodePoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNodePoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetNodePoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetNodePoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetNodePoolDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNodePoolOK creates a GetNodePoolOK with default headers values
func NewGetNodePoolOK() *GetNodePoolOK {
	return &GetNodePoolOK{}
}

/*GetNodePoolOK handles this case with default header values.

Node pool details
*/
type GetNodePoolOK struct {
	Payload *models.V5GetNodePoolResponse
}

func (o *GetNodePoolOK) Error() string {
	return fmt.Sprintf("[GET /v5/clusters/{cluster_id}/nodepools/{nodepool_id}/][%d] getNodePoolOK  %+v", 200, o.Payload)
}

func (o *GetNodePoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V5GetNodePoolResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodePoolUnauthorized creates a GetNodePoolUnauthorized with default headers values
func NewGetNodePoolUnauthorized() *GetNodePoolUnauthorized {
	return &GetNodePoolUnauthorized{}
}

/*GetNodePoolUnauthorized handles this case with default header values.

Permission denied
*/
type GetNodePoolUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *GetNodePoolUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v5/clusters/{cluster_id}/nodepools/{nodepool_id}/][%d] getNodePoolUnauthorized  %+v", 401, o.Payload)
}

func (o *GetNodePoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodePoolNotFound creates a GetNodePoolNotFound with default headers values
func NewGetNodePoolNotFound() *GetNodePoolNotFound {
	return &GetNodePoolNotFound{}
}

/*GetNodePoolNotFound handles this case with default header values.

Cluster not found
*/
type GetNodePoolNotFound struct {
	Payload *models.V4GenericResponse
}

func (o *GetNodePoolNotFound) Error() string {
	return fmt.Sprintf("[GET /v5/clusters/{cluster_id}/nodepools/{nodepool_id}/][%d] getNodePoolNotFound  %+v", 404, o.Payload)
}

func (o *GetNodePoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNodePoolDefault creates a GetNodePoolDefault with default headers values
func NewGetNodePoolDefault(code int) *GetNodePoolDefault {
	return &GetNodePoolDefault{
		_statusCode: code,
	}
}

/*GetNodePoolDefault handles this case with default header values.

error
*/
type GetNodePoolDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the get node pool default response
func (o *GetNodePoolDefault) Code() int {
	return o._statusCode
}

func (o *GetNodePoolDefault) Error() string {
	return fmt.Sprintf("[GET /v5/clusters/{cluster_id}/nodepools/{nodepool_id}/][%d] getNodePool default  %+v", o._statusCode, o.Payload)
}

func (o *GetNodePoolDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}