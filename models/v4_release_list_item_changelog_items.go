// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4ReleaseListItemChangelogItems v4 release list item changelog items
// swagger:model v4ReleaseListItemChangelogItems
type V4ReleaseListItemChangelogItems struct {

	// If the changed item was a component, this attribute is the
	// name of the component.
	//
	Component string `json:"component,omitempty"`

	// Human-friendly description of the change
	Description string `json:"description,omitempty"`
}

// Validate validates this v4 release list item changelog items
func (m *V4ReleaseListItemChangelogItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4ReleaseListItemChangelogItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4ReleaseListItemChangelogItems) UnmarshalBinary(b []byte) error {
	var res V4ReleaseListItemChangelogItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}