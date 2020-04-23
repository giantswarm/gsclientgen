// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// V5GetNodePoolResponseNodeSpecAwsInstanceDistribution Attributes defining the instance distribution in a node pool.
//
// swagger:model v5GetNodePoolResponseNodeSpecAwsInstanceDistribution
type V5GetNodePoolResponseNodeSpecAwsInstanceDistribution struct {

	// Base capacity of on-demand EC2 instances to use for worker nodes in this pools.
	// Find details on this attribute in the [addNodePool](#operation/addNodePool) operation.
	//
	OnDemandBaseCapacity int64 `json:"on_demand_base_capacity,omitempty"`

	// Percentage of on-demand EC2 instances to use for worker nodes, instead of spot
	// instances, for instances exceeding `on_demand_base_capacity`.
	// Find details on this attribute in the [addNodePool](#operation/addNodePool) operation.
	//
	OnDemandPercentageAboveBaseCapacity int64 `json:"on_demand_percentage_above_base_capacity,omitempty"`
}

// Validate validates this v5 get node pool response node spec aws instance distribution
func (m *V5GetNodePoolResponseNodeSpecAwsInstanceDistribution) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *V5GetNodePoolResponseNodeSpecAwsInstanceDistribution) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V5GetNodePoolResponseNodeSpecAwsInstanceDistribution) UnmarshalBinary(b []byte) error {
	var res V5GetNodePoolResponseNodeSpecAwsInstanceDistribution
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}