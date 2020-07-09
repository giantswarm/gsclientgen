// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V5SetClusterLabelsRequest Cluster labels
// swagger:model v5SetClusterLabelsRequest
type V5SetClusterLabelsRequest struct {

	// Labels object
	//
	// Object containing keys with string values representing label changes
	Labels map[string]*string `json:"labels,omitempty"`
}

// Validate validates this v5 set cluster labels request
func (m *V5SetClusterLabelsRequest) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V5SetClusterLabelsRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5SetClusterLabelsRequest) UnmarshalBinary(b []byte) error {
	var res V5SetClusterLabelsRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}