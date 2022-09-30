// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FcLogin A Fibre Channel (FC) login represents a connection formed by an FC initiator that has successfully logged in to ONTAP. This represents the FC login on which higher-level protocols such as Fibre Channel Protocol and NVMe over Fibre Channel (NVMe/FC) rely.
//
// swagger:model fc_login
type FcLogin struct {

	// links
	Links *FcLoginLinks `json:"_links,omitempty"`

	// The initiator groups in which the initiator is a member.
	//
	// Read Only: true
	Igroups []*FcLoginIgroupsItems0 `json:"igroups,omitempty"`

	// initiator
	Initiator *FcLoginInitiator `json:"initiator,omitempty"`

	// interface
	Interface *FcLoginInterface `json:"interface,omitempty"`

	// The data protocol used to perform the login.
	//
	// Read Only: true
	// Enum: [fc_nvme fcp]
	Protocol string `json:"protocol,omitempty"`

	// svm
	Svm *FcLoginSvm `json:"svm,omitempty"`
}

// Validate validates this fc login
func (m *FcLogin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIgroups(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInitiator(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInterface(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSvm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLogin) validateLinks(formats strfmt.Registry) error {
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

func (m *FcLogin) validateIgroups(formats strfmt.Registry) error {
	if swag.IsZero(m.Igroups) { // not required
		return nil
	}

	for i := 0; i < len(m.Igroups); i++ {
		if swag.IsZero(m.Igroups[i]) { // not required
			continue
		}

		if m.Igroups[i] != nil {
			if err := m.Igroups[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FcLogin) validateInitiator(formats strfmt.Registry) error {
	if swag.IsZero(m.Initiator) { // not required
		return nil
	}

	if m.Initiator != nil {
		if err := m.Initiator.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiator")
			}
			return err
		}
	}

	return nil
}

func (m *FcLogin) validateInterface(formats strfmt.Registry) error {
	if swag.IsZero(m.Interface) { // not required
		return nil
	}

	if m.Interface != nil {
		if err := m.Interface.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface")
			}
			return err
		}
	}

	return nil
}

var fcLoginTypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["fc_nvme","fcp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		fcLoginTypeProtocolPropEnum = append(fcLoginTypeProtocolPropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// fc_login
	// FcLogin
	// protocol
	// Protocol
	// fc_nvme
	// END DEBUGGING
	// FcLoginProtocolFcNvme captures enum value "fc_nvme"
	FcLoginProtocolFcNvme string = "fc_nvme"

	// BEGIN DEBUGGING
	// fc_login
	// FcLogin
	// protocol
	// Protocol
	// fcp
	// END DEBUGGING
	// FcLoginProtocolFcp captures enum value "fcp"
	FcLoginProtocolFcp string = "fcp"
)

// prop value enum
func (m *FcLogin) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, fcLoginTypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *FcLogin) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(m.Protocol) { // not required
		return nil
	}

	// value enum
	if err := m.validateProtocolEnum("protocol", "body", m.Protocol); err != nil {
		return err
	}

	return nil
}

func (m *FcLogin) validateSvm(formats strfmt.Registry) error {
	if swag.IsZero(m.Svm) { // not required
		return nil
	}

	if m.Svm != nil {
		if err := m.Svm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc login based on the context it is used
func (m *FcLogin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIgroups(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInitiator(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterface(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateProtocol(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLogin) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *FcLogin) contextValidateIgroups(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "igroups", "body", []*FcLoginIgroupsItems0(m.Igroups)); err != nil {
		return err
	}

	for i := 0; i < len(m.Igroups); i++ {

		if m.Igroups[i] != nil {
			if err := m.Igroups[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("igroups" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FcLogin) contextValidateInitiator(ctx context.Context, formats strfmt.Registry) error {

	if m.Initiator != nil {
		if err := m.Initiator.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("initiator")
			}
			return err
		}
	}

	return nil
}

func (m *FcLogin) contextValidateInterface(ctx context.Context, formats strfmt.Registry) error {

	if m.Interface != nil {
		if err := m.Interface.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface")
			}
			return err
		}
	}

	return nil
}

func (m *FcLogin) contextValidateProtocol(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "protocol", "body", string(m.Protocol)); err != nil {
		return err
	}

	return nil
}

func (m *FcLogin) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

	if m.Svm != nil {
		if err := m.Svm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcLogin) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcLogin) UnmarshalBinary(b []byte) error {
	var res FcLogin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcLoginIgroupsItems0 fc login igroups items0
//
// swagger:model FcLoginIgroupsItems0
type FcLoginIgroupsItems0 struct {

	// links
	Links *FcLoginIgroupsItems0Links `json:"_links,omitempty"`

	// The name of the initiator group.
	//
	// Example: igroup1
	// Max Length: 96
	// Min Length: 1
	Name string `json:"name,omitempty"`

	// The unique identifier of the initiator group.
	//
	// Example: 4ea7a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this fc login igroups items0
func (m *FcLoginIgroupsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLoginIgroupsItems0) validateLinks(formats strfmt.Registry) error {
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

func (m *FcLoginIgroupsItems0) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", m.Name, 96); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this fc login igroups items0 based on the context it is used
func (m *FcLoginIgroupsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLoginIgroupsItems0) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *FcLoginIgroupsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcLoginIgroupsItems0) UnmarshalBinary(b []byte) error {
	var res FcLoginIgroupsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcLoginIgroupsItems0Links fc login igroups items0 links
//
// swagger:model FcLoginIgroupsItems0Links
type FcLoginIgroupsItems0Links struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fc login igroups items0 links
func (m *FcLoginIgroupsItems0Links) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLoginIgroupsItems0Links) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc login igroups items0 links based on the context it is used
func (m *FcLoginIgroupsItems0Links) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLoginIgroupsItems0Links) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcLoginIgroupsItems0Links) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcLoginIgroupsItems0Links) UnmarshalBinary(b []byte) error {
	var res FcLoginIgroupsItems0Links
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcLoginInitiator Information about the logged in FC initiator.
//
// swagger:model FcLoginInitiator
type FcLoginInitiator struct {

	// The logged in initiator world wide port name (WWPN) aliases.
	//
	// Read Only: true
	Aliases []string `json:"aliases,omitempty"`

	// A comment available for use by the administrator. This is modifiable from the initiator REST endpoint directly. See [`PATCH /protocols/san/igroups/{igroup.uuid}/initiators/{name}`](#/SAN/igroup_initiator_modify).
	//
	// Example: This is an FC initiator for host 5
	// Read Only: true
	Comment string `json:"comment,omitempty"`

	// The port address of the initiator's FC port.<br/>
	// Each port in an FC switched fabric has its own unique port address for routing purposes. The port address is assigned by a switch in the fabric when that port logs in to the fabric. This property refers to the address given by a switch to the initiator port.<br/>
	// This is useful for obtaining statistics and diagnostic information from FC switches.<br/>
	// This is a hexadecimal encoded numeric value.
	//
	// Example: 5060A
	// Read Only: true
	PortAddress string `json:"port_address,omitempty"`

	// The logged in initiator world wide node name (WWNN).
	//
	// Example: 2f:a0:00:a0:98:0b:56:13
	// Read Only: true
	Wwnn string `json:"wwnn,omitempty"`

	// The logged in initiator WWPN.
	//
	// Example: 2f:a0:00:a0:98:0b:56:13
	// Read Only: true
	Wwpn string `json:"wwpn,omitempty"`
}

// Validate validates this fc login initiator
func (m *FcLoginInitiator) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this fc login initiator based on the context it is used
func (m *FcLoginInitiator) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateAliases(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateComment(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePortAddress(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwnn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLoginInitiator) contextValidateAliases(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "initiator"+"."+"aliases", "body", []string(m.Aliases)); err != nil {
		return err
	}

	for i := 0; i < len(m.Aliases); i++ {

		if err := validate.ReadOnly(ctx, "initiator"+"."+"aliases"+"."+strconv.Itoa(i), "body", string(m.Aliases[i])); err != nil {
			return err
		}

	}

	return nil
}

func (m *FcLoginInitiator) contextValidateComment(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "initiator"+"."+"comment", "body", string(m.Comment)); err != nil {
		return err
	}

	return nil
}

func (m *FcLoginInitiator) contextValidatePortAddress(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "initiator"+"."+"port_address", "body", string(m.PortAddress)); err != nil {
		return err
	}

	return nil
}

func (m *FcLoginInitiator) contextValidateWwnn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "initiator"+"."+"wwnn", "body", string(m.Wwnn)); err != nil {
		return err
	}

	return nil
}

func (m *FcLoginInitiator) contextValidateWwpn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "initiator"+"."+"wwpn", "body", string(m.Wwpn)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcLoginInitiator) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcLoginInitiator) UnmarshalBinary(b []byte) error {
	var res FcLoginInitiator
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcLoginInterface An FC interface.
//
// swagger:model FcLoginInterface
type FcLoginInterface struct {

	// links
	Links *FcLoginInterfaceLinks `json:"_links,omitempty"`

	// The name of the FC interface.
	//
	// Example: fc_lif1
	Name string `json:"name,omitempty"`

	// The unique identifier of the FC interface.
	//
	// Example: 3a09ab42-4da1-32cf-9d35-3385a6101a0b
	UUID string `json:"uuid,omitempty"`

	// The WWPN of the FC interface.
	//
	// Example: 20:00:00:50:56:b4:13:a8
	// Read Only: true
	Wwpn string `json:"wwpn,omitempty"`
}

// Validate validates this fc login interface
func (m *FcLoginInterface) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLoginInterface) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc login interface based on the context it is used
func (m *FcLoginInterface) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateWwpn(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLoginInterface) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *FcLoginInterface) contextValidateWwpn(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "interface"+"."+"wwpn", "body", string(m.Wwpn)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcLoginInterface) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcLoginInterface) UnmarshalBinary(b []byte) error {
	var res FcLoginInterface
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcLoginInterfaceLinks fc login interface links
//
// swagger:model FcLoginInterfaceLinks
type FcLoginInterfaceLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fc login interface links
func (m *FcLoginInterfaceLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLoginInterfaceLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc login interface links based on the context it is used
func (m *FcLoginInterfaceLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLoginInterfaceLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("interface" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcLoginInterfaceLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcLoginInterfaceLinks) UnmarshalBinary(b []byte) error {
	var res FcLoginInterfaceLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcLoginLinks fc login links
//
// swagger:model FcLoginLinks
type FcLoginLinks struct {

	// next
	Next *Href `json:"next,omitempty"`

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fc login links
func (m *FcLoginLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNext(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLoginLinks) validateNext(formats strfmt.Registry) error {
	if swag.IsZero(m.Next) { // not required
		return nil
	}

	if m.Next != nil {
		if err := m.Next.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *FcLoginLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc login links based on the context it is used
func (m *FcLoginLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNext(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLoginLinks) contextValidateNext(ctx context.Context, formats strfmt.Registry) error {

	if m.Next != nil {
		if err := m.Next.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "next")
			}
			return err
		}
	}

	return nil
}

func (m *FcLoginLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcLoginLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcLoginLinks) UnmarshalBinary(b []byte) error {
	var res FcLoginLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcLoginSvm fc login svm
//
// swagger:model FcLoginSvm
type FcLoginSvm struct {

	// links
	Links *FcLoginSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this fc login svm
func (m *FcLoginSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLoginSvm) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc login svm based on the context it is used
func (m *FcLoginSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLoginSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcLoginSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcLoginSvm) UnmarshalBinary(b []byte) error {
	var res FcLoginSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// FcLoginSvmLinks fc login svm links
//
// swagger:model FcLoginSvmLinks
type FcLoginSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this fc login svm links
func (m *FcLoginSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLoginSvmLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this fc login svm links based on the context it is used
func (m *FcLoginSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FcLoginSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("svm" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FcLoginSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FcLoginSvmLinks) UnmarshalBinary(b []byte) error {
	var res FcLoginSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
