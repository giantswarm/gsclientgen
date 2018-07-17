// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V4ClusterListItem v4 cluster list item
// swagger:model v4ClusterListItem
type V4ClusterListItem struct {

	// Date/time of cluster creation
	CreateDate string `json:"create_date,omitempty"`

	// Unique cluster identifier
	ID string `json:"id,omitempty"`

	// Cluster name
	Name string `json:"name,omitempty"`

	// Name of the organization owning the cluster
	Owner string `json:"owner,omitempty"`

	// The semantic version number of this cluster
	ReleaseVersion string `json:"release_version,omitempty"`
}

// Validate validates this v4 cluster list item
func (m *V4ClusterListItem) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V4ClusterListItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V4ClusterListItem) UnmarshalBinary(b []byte) error {
	var res V4ClusterListItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
