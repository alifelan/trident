// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ConsistencyGroupNamespaceSubsystemMap The NVMe subsystem with which the NVMe namespace is associated. A namespace can be mapped to zero (0) or one (1) subsystems.<br/>
// There is an added cost to retrieving property values for `subsystem_map`.
// They are not populated for either a collection GET or an instance GET unless explicitly requested using the `fields` query parameter.
//
// swagger:model consistency_group_namespace_subsystem_map
type ConsistencyGroupNamespaceSubsystemMap struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// The Asymmetric Namespace Access Group ID (ANAGRPID) of the NVMe namespace.<br/>
	// The format for an ANAGRPID is 8 hexadecimal digits (zero-filled) followed by a lower case "h".
	//
	// Example: 00103050h
	Anagrpid string `json:"anagrpid,omitempty"`

	// The NVMe namespace identifier. This is an identifier used by an NVMe controller to provide access to the NVMe namespace.<br/>
	// The format for an NVMe namespace identifier is 8 hexadecimal digits (zero-filled) followed by a lower case "h".
	//
	// Example: 00000001h
	Nsid string `json:"nsid,omitempty"`

	// The NVMe subsystem to which the NVMe namespace is mapped.
	//
	Subsystem *NvmeSubsystemReference `json:"subsystem,omitempty"`
}

// Validate validates this consistency group namespace subsystem map
func (m *ConsistencyGroupNamespaceSubsystemMap) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubsystem(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupNamespaceSubsystemMap) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupNamespaceSubsystemMap) validateSubsystem(formats strfmt.Registry) error {
	if swag.IsZero(m.Subsystem) { // not required
		return nil
	}

	if m.Subsystem != nil {
		if err := m.Subsystem.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this consistency group namespace subsystem map based on the context it is used
func (m *ConsistencyGroupNamespaceSubsystemMap) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSubsystem(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ConsistencyGroupNamespaceSubsystemMap) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links")
			}
			return err
		}
	}

	return nil
}

func (m *ConsistencyGroupNamespaceSubsystemMap) contextValidateSubsystem(ctx context.Context, formats strfmt.Registry) error {

	if m.Subsystem != nil {
		if err := m.Subsystem.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("subsystem")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ConsistencyGroupNamespaceSubsystemMap) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ConsistencyGroupNamespaceSubsystemMap) UnmarshalBinary(b []byte) error {
	var res ConsistencyGroupNamespaceSubsystemMap
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
