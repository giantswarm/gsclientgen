// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V5AddNodePoolRequestNodeSpec v5 add node pool request node spec
// swagger:model v5AddNodePoolRequestNodeSpec
type V5AddNodePoolRequestNodeSpec struct {

	// aws
	Aws *V5AddNodePoolRequestNodeSpecAws `json:"aws,omitempty"`

	// azure
	Azure *V5AddNodePoolRequestNodeSpecAzure `json:"azure,omitempty"`
}

// Validate validates this v5 add node pool request node spec
func (m *V5AddNodePoolRequestNodeSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAws(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAzure(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V5AddNodePoolRequestNodeSpec) validateAws(formats strfmt.Registry) error {

	if swag.IsZero(m.Aws) { // not required
		return nil
	}

	if m.Aws != nil {
		if err := m.Aws.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("aws")
			}
			return err
		}
	}

	return nil
}

func (m *V5AddNodePoolRequestNodeSpec) validateAzure(formats strfmt.Registry) error {

	if swag.IsZero(m.Azure) { // not required
		return nil
	}

	if m.Azure != nil {
		if err := m.Azure.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("azure")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V5AddNodePoolRequestNodeSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5AddNodePoolRequestNodeSpec) UnmarshalBinary(b []byte) error {
	var res V5AddNodePoolRequestNodeSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
