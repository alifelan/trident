// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// MccipPort Port configuration specification.
// l3_config information is only needed when configuring a MetroCluster IP for use in a layer 3 network.
//
// swagger:model mccip_port
type MccipPort struct {

	// l3 config
	L3Config *MccipPortL3Config `json:"l3_config,omitempty"`

	// Port name
	// Example: e1b
	Name string `json:"name,omitempty"`

	// node
	Node *MccipPortNode `json:"node,omitempty"`

	// Port UUID
	UUID string `json:"uuid,omitempty"`

	// VLAN ID
	// Example: 200
	// Maximum: 4095
	// Minimum: 10
	VlanID int64 `json:"vlan_id,omitempty"`
}

// Validate validates this mccip port
func (m *MccipPort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateL3Config(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPort) validateL3Config(formats strfmt.Registry) error {
	if swag.IsZero(m.L3Config) { // not required
		return nil
	}

	if m.L3Config != nil {
		if err := m.L3Config.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l3_config")
			}
			return err
		}
	}

	return nil
}

func (m *MccipPort) validateNode(formats strfmt.Registry) error {
	if swag.IsZero(m.Node) { // not required
		return nil
	}

	if m.Node != nil {
		if err := m.Node.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

func (m *MccipPort) validateVlanID(formats strfmt.Registry) error {
	if swag.IsZero(m.VlanID) { // not required
		return nil
	}

	if err := validate.MinimumInt("vlan_id", "body", m.VlanID, 10, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("vlan_id", "body", m.VlanID, 4095, false); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this mccip port based on the context it is used
func (m *MccipPort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateL3Config(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPort) contextValidateL3Config(ctx context.Context, formats strfmt.Registry) error {

	if m.L3Config != nil {
		if err := m.L3Config.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l3_config")
			}
			return err
		}
	}

	return nil
}

func (m *MccipPort) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

	if m.Node != nil {
		if err := m.Node.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MccipPort) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MccipPort) UnmarshalBinary(b []byte) error {
	var res MccipPort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MccipPortL3Config mccip port l3 config
//
// swagger:model MccipPortL3Config
type MccipPortL3Config struct {

	// ipv4 interface
	IPV4Interface *MccipPortL3ConfigIPV4Interface `json:"ipv4_interface,omitempty"`
}

// Validate validates this mccip port l3 config
func (m *MccipPortL3Config) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIPV4Interface(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortL3Config) validateIPV4Interface(formats strfmt.Registry) error {
	if swag.IsZero(m.IPV4Interface) { // not required
		return nil
	}

	if m.IPV4Interface != nil {
		if err := m.IPV4Interface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l3_config" + "." + "ipv4_interface")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mccip port l3 config based on the context it is used
func (m *MccipPortL3Config) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateIPV4Interface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortL3Config) contextValidateIPV4Interface(ctx context.Context, formats strfmt.Registry) error {

	if m.IPV4Interface != nil {
		if err := m.IPV4Interface.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("l3_config" + "." + "ipv4_interface")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MccipPortL3Config) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MccipPortL3Config) UnmarshalBinary(b []byte) error {
	var res MccipPortL3Config
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MccipPortL3ConfigIPV4Interface Object to setup an interface along with its default router.
//
// swagger:model MccipPortL3ConfigIPV4Interface
type MccipPortL3ConfigIPV4Interface struct {

	// IPv4 or IPv6 address
	// Example: 10.10.10.7
	Address string `json:"address,omitempty"`

	// The IPv4 or IPv6 address of the default router.
	// Example: 10.1.1.1
	Gateway string `json:"gateway,omitempty"`

	// netmask
	Netmask IPNetmask `json:"netmask,omitempty"`
}

// Validate validates this mccip port l3 config IP v4 interface
func (m *MccipPortL3ConfigIPV4Interface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetmask(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortL3ConfigIPV4Interface) validateNetmask(formats strfmt.Registry) error {
	if swag.IsZero(m.Netmask) { // not required
		return nil
	}

	if err := m.Netmask.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("l3_config" + "." + "ipv4_interface" + "." + "netmask")
		}
		return err
	}

	return nil
}

// ContextValidate validate this mccip port l3 config IP v4 interface based on the context it is used
func (m *MccipPortL3ConfigIPV4Interface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNetmask(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortL3ConfigIPV4Interface) contextValidateNetmask(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Netmask.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("l3_config" + "." + "ipv4_interface" + "." + "netmask")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MccipPortL3ConfigIPV4Interface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MccipPortL3ConfigIPV4Interface) UnmarshalBinary(b []byte) error {
	var res MccipPortL3ConfigIPV4Interface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MccipPortNode Node information
//
// swagger:model MccipPortNode
type MccipPortNode struct {

	// links
	Links *MccipPortNodeLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this mccip port node
func (m *MccipPortNode) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortNode) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mccip port node based on the context it is used
func (m *MccipPortNode) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortNode) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MccipPortNode) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MccipPortNode) UnmarshalBinary(b []byte) error {
	var res MccipPortNode
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// MccipPortNodeLinks mccip port node links
//
// swagger:model MccipPortNodeLinks
type MccipPortNodeLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this mccip port node links
func (m *MccipPortNodeLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortNodeLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this mccip port node links based on the context it is used
func (m *MccipPortNodeLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MccipPortNodeLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("node" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MccipPortNodeLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MccipPortNodeLinks) UnmarshalBinary(b []byte) error {
	var res MccipPortNodeLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
