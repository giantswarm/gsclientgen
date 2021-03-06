// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V4GetCredentialsResponseItemsAzure Credentials specific to an Azure service principal
// swagger:model v4GetCredentialsResponseItemsAzure
type V4GetCredentialsResponseItemsAzure struct {

	// credential
	Credential *V4GetCredentialsResponseItemsAzureCredential `json:"credential,omitempty"`
}

// Validate validates this v4 get credentials response items azure
func (m *V4GetCredentialsResponseItemsAzure) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCredential(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V4GetCredentialsResponseItemsAzure) validateCredential(formats strfmt.Registry) error {

	if swag.IsZero(m.Credential) { // not required
		return nil
	}

	if m.Credential != nil {
		if err := m.Credential.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("credential")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V4GetCredentialsResponseItemsAzure) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4GetCredentialsResponseItemsAzure) UnmarshalBinary(b []byte) error {
	var res V4GetCredentialsResponseItemsAzure
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
