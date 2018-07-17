// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// GetUserOKBody get user o k body
// swagger:model getUserOKBody
type GetUserOKBody struct {

	// The date and time that this account was created
	Created string `json:"created,omitempty"`

	// Email address of the user
	Email string `json:"email,omitempty"`

	// The date and time when this account will expire
	Expiry string `json:"expiry,omitempty"`
}

// Validate validates this get user o k body
func (m *GetUserOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetUserOKBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetUserOKBody) UnmarshalBinary(b []byte) error {
	var res GetUserOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}