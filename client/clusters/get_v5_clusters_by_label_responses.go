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

// GetV5ClustersByLabelReader is a Reader for the GetV5ClustersByLabel structure.
type GetV5ClustersByLabelReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetV5ClustersByLabelReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetV5ClustersByLabelOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 401:
		result := NewGetV5ClustersByLabelUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		result := NewGetV5ClustersByLabelDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetV5ClustersByLabelOK creates a GetV5ClustersByLabelOK with default headers values
func NewGetV5ClustersByLabelOK() *GetV5ClustersByLabelOK {
	return &GetV5ClustersByLabelOK{}
}

/*GetV5ClustersByLabelOK handles this case with default header values.

Success
*/
type GetV5ClustersByLabelOK struct {
	Payload []*models.V4ClusterListItem
}

func (o *GetV5ClustersByLabelOK) Error() string {
	return fmt.Sprintf("[POST /v5/clusters/by_label/][%d] getV5ClustersByLabelOK  %+v", 200, o.Payload)
}

func (o *GetV5ClustersByLabelOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV5ClustersByLabelUnauthorized creates a GetV5ClustersByLabelUnauthorized with default headers values
func NewGetV5ClustersByLabelUnauthorized() *GetV5ClustersByLabelUnauthorized {
	return &GetV5ClustersByLabelUnauthorized{}
}

/*GetV5ClustersByLabelUnauthorized handles this case with default header values.

Permission denied
*/
type GetV5ClustersByLabelUnauthorized struct {
	Payload *models.V4GenericResponse
}

func (o *GetV5ClustersByLabelUnauthorized) Error() string {
	return fmt.Sprintf("[POST /v5/clusters/by_label/][%d] getV5ClustersByLabelUnauthorized  %+v", 401, o.Payload)
}

func (o *GetV5ClustersByLabelUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetV5ClustersByLabelDefault creates a GetV5ClustersByLabelDefault with default headers values
func NewGetV5ClustersByLabelDefault(code int) *GetV5ClustersByLabelDefault {
	return &GetV5ClustersByLabelDefault{
		_statusCode: code,
	}
}

/*GetV5ClustersByLabelDefault handles this case with default header values.

Error
*/
type GetV5ClustersByLabelDefault struct {
	_statusCode int

	Payload *models.V4GenericResponse
}

// Code gets the status code for the get v5 clusters by label default response
func (o *GetV5ClustersByLabelDefault) Code() int {
	return o._statusCode
}

func (o *GetV5ClustersByLabelDefault) Error() string {
	return fmt.Sprintf("[POST /v5/clusters/by_label/][%d] getV5ClustersByLabel default  %+v", o._statusCode, o.Payload)
}

func (o *GetV5ClustersByLabelDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V4GenericResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
