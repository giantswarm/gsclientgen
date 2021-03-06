// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V4GetClusterAppsResponseItemsSpec v4 get cluster apps response items spec
// swagger:model v4GetClusterAppsResponseItemsSpec
type V4GetClusterAppsResponseItemsSpec struct {

	// The catalog that this app came from
	Catalog string `json:"catalog,omitempty"`

	// Name of the chart that was used to install this app
	Name string `json:"name,omitempty"`

	// Namespace that this app is installed to
	Namespace string `json:"namespace,omitempty"`

	// user config
	UserConfig *V4GetClusterAppsResponseItemsSpecUserConfig `json:"user_config,omitempty"`

	// Version of the chart that was used to install this app
	Version string `json:"version,omitempty"`
}

// Validate validates this v4 get cluster apps response items spec
func (m *V4GetClusterAppsResponseItemsSpec) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUserConfig(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V4GetClusterAppsResponseItemsSpec) validateUserConfig(formats strfmt.Registry) error {

	if swag.IsZero(m.UserConfig) { // not required
		return nil
	}

	if m.UserConfig != nil {
		if err := m.UserConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("user_config")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V4GetClusterAppsResponseItemsSpec) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4GetClusterAppsResponseItemsSpec) UnmarshalBinary(b []byte) error {
	var res V4GetClusterAppsResponseItemsSpec
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
