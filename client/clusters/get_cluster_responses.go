// Code generated by go-swagger; DO NOT EDIT.

package clusters

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/giantswarm/gsclientgen/v2/models"
)

// GetClusterReader is a Reader for the GetCluster structure.
type GetClusterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetClusterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetClusterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 301:
		result := NewGetClusterMovedPermanently()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 401:
		result := NewGetClusterUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewGetClusterNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetClusterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetClusterOK creates a GetClusterOK with default headers values
func NewGetClusterOK() *GetClusterOK {
	return &GetClusterOK{}
}

/*GetClusterOK handles this case with default header values.

Cluster details
*/
type GetClusterOK struct {
	Payload *models.V4ClusterDetailsResponse
}

func (o *GetClusterOK) Error() string {
	return fmt.Sprintf("[GET /v4/clusters/{cluster_id}/][%d] getClusterOK  %+v", 200, o.Payload)
}

func (o *GetClusterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4ClusterDetailsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterMovedPermanently creates a GetClusterMovedPermanently with default headers values
func NewGetClusterMovedPermanently() *GetClusterMovedPermanently {
	return &GetClusterMovedPermanently{}
}

/*GetClusterMovedPermanently handles this case with default header values.

Version mismatch
*/
type GetClusterMovedPermanently struct {
	/*URI of the new path to use for retrying the request, as the one
	used is not usable for this cluster.

	*/
	Location string
}

func (o *GetClusterMovedPermanently) Error() string {
	return fmt.Sprintf("[GET /v4/clusters/{cluster_id}/][%d] getClusterMovedPermanently ", 301)
}

func (o *GetClusterMovedPermanently) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header Location
	o.Location = response.GetHeader("Location")

	return nil
}

// NewGetClusterUnauthorized creates a GetClusterUnauthorized with default headers values
func NewGetClusterUnauthorized() *GetClusterUnauthorized {
	return &GetClusterUnauthorized{}
}

/*GetClusterUnauthorized handles this case with default header values.

Permission denied
*/
type GetClusterUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *GetClusterUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v4/clusters/{cluster_id}/][%d] getClusterUnauthorized  %+v", 401, o.Payload)
}

func (o *GetClusterUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterNotFound creates a GetClusterNotFound with default headers values
func NewGetClusterNotFound() *GetClusterNotFound {
	return &GetClusterNotFound{}
}

/*GetClusterNotFound handles this case with default header values.

Cluster not found
*/
type GetClusterNotFound struct {
	Payload *models.V4GenericResponse
}

func (o *GetClusterNotFound) Error() string {
	return fmt.Sprintf("[GET /v4/clusters/{cluster_id}/][%d] getClusterNotFound  %+v", 404, o.Payload)
}

func (o *GetClusterNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetClusterDefault creates a GetClusterDefault with default headers values
func NewGetClusterDefault(code int) *GetClusterDefault {
	return &GetClusterDefault{
		_statusCode: code,
	}
}

/*GetClusterDefault handles this case with default header values.

error
*/
type GetClusterDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the get cluster default response
func (o *GetClusterDefault) Code() int {
	return o._statusCode
}

func (o *GetClusterDefault) Error() string {
	return fmt.Sprintf("[GET /v4/clusters/{cluster_id}/][%d] getCluster default  %+v", o._statusCode, o.Payload)
}

func (o *GetClusterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
