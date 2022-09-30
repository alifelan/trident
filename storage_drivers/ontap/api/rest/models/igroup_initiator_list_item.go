// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IgroupInitiatorListItem igroup initiator list item
//
// swagger:model igroup_initiator_list_item
type IgroupInitiatorListItem struct {

	// links
	Links *IgroupInitiatorListItemLinks `json:"_links,omitempty"`

	// A comment available for use by the administrator. Valid in POST and PATCH.
	//
	// Max Length: 254
	// Min Length: 0
	Comment *string `json:"comment,omitempty"`

	// connectivity tracking
	ConnectivityTracking *IgroupInitiatorListItemConnectivityTracking `json:"connectivity_tracking,omitempty"`

	// igroup
	Igroup *IgroupInitiatorListItemIgroup `json:"igroup,omitempty"`

	// The FC WWPN, iSCSI IQN, or iSCSI EUI that identifies the host initiator. Valid in POST only and not allowed when the `records` property is used.<br/>
	// An FC WWPN consists of 16 hexadecimal digits grouped as 8 pairs separated by colons. The format for an iSCSI IQN is _iqn.yyyy-mm.reverse_domain_name:any_. The iSCSI EUI format consists of the _eui._ prefix followed by 16 hexadecimal characters.
	//
	// Example: iqn.1998-01.com.corp.iscsi:name1
	// Max Length: 96
	// Min Length: 1
	Name string `json:"name,omitempty"`
}

// Validate validates this igroup initiator list item
func (m *IgroupInitiatorListItem) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConnectivityTracking(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIgroup(formats); err != nil {
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

func (m *IgroupInitiatorListItem) validateLinks(formats strfmt.Registry) error {
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

func (m *IgroupInitiatorListItem) validateComment(formats strfmt.Registry) error {
	if swag.IsZero(m.Comment) { // not required
		return nil
	}

	if err := validate.MinLength("comment", "body", *m.Comment, 0); err != nil {
		return err
	}

	if err := validate.MaxLength("comment", "body", *m.Comment, 254); err != nil {
		return err
	}

	return nil
}

func (m *IgroupInitiatorListItem) validateConnectivityTracking(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectivityTracking) { // not required
		return nil
	}

	if m.ConnectivityTracking != nil {
		if err := m.ConnectivityTracking.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectivity_tracking")
			}
			return err
		}
	}

	return nil
}

func (m *IgroupInitiatorListItem) validateIgroup(formats strfmt.Registry) error {
	if swag.IsZero(m.Igroup) { // not required
		return nil
	}

	if m.Igroup != nil {
		if err := m.Igroup.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("igroup")
			}
			return err
		}
	}

	return nil
}

func (m *IgroupInitiatorListItem) validateName(formats strfmt.Registry) error {
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

// ContextValidate validate this igroup initiator list item based on the context it is used
func (m *IgroupInitiatorListItem) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateConnectivityTracking(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateIgroup(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItem) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *IgroupInitiatorListItem) contextValidateConnectivityTracking(ctx context.Context, formats strfmt.Registry) error {

	if m.ConnectivityTracking != nil {
		if err := m.ConnectivityTracking.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectivity_tracking")
			}
			return err
		}
	}

	return nil
}

func (m *IgroupInitiatorListItem) contextValidateIgroup(ctx context.Context, formats strfmt.Registry) error {

	if m.Igroup != nil {
		if err := m.Igroup.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("igroup")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IgroupInitiatorListItem) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItem) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItem
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupInitiatorListItemConnectivityTracking Overview of the initiator's connections to ONTAP.
//
// swagger:model IgroupInitiatorListItemConnectivityTracking
type IgroupInitiatorListItemConnectivityTracking struct {

	// Connection state.
	// Read Only: true
	// Enum: [full none partial no_lun_maps]
	ConnectionState string `json:"connection_state,omitempty"`
}

// Validate validates this igroup initiator list item connectivity tracking
func (m *IgroupInitiatorListItemConnectivityTracking) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectionState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var igroupInitiatorListItemConnectivityTrackingTypeConnectionStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["full","none","partial","no_lun_maps"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		igroupInitiatorListItemConnectivityTrackingTypeConnectionStatePropEnum = append(igroupInitiatorListItemConnectivityTrackingTypeConnectionStatePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// IgroupInitiatorListItemConnectivityTracking
	// IgroupInitiatorListItemConnectivityTracking
	// connection_state
	// ConnectionState
	// full
	// END DEBUGGING
	// IgroupInitiatorListItemConnectivityTrackingConnectionStateFull captures enum value "full"
	IgroupInitiatorListItemConnectivityTrackingConnectionStateFull string = "full"

	// BEGIN DEBUGGING
	// IgroupInitiatorListItemConnectivityTracking
	// IgroupInitiatorListItemConnectivityTracking
	// connection_state
	// ConnectionState
	// none
	// END DEBUGGING
	// IgroupInitiatorListItemConnectivityTrackingConnectionStateNone captures enum value "none"
	IgroupInitiatorListItemConnectivityTrackingConnectionStateNone string = "none"

	// BEGIN DEBUGGING
	// IgroupInitiatorListItemConnectivityTracking
	// IgroupInitiatorListItemConnectivityTracking
	// connection_state
	// ConnectionState
	// partial
	// END DEBUGGING
	// IgroupInitiatorListItemConnectivityTrackingConnectionStatePartial captures enum value "partial"
	IgroupInitiatorListItemConnectivityTrackingConnectionStatePartial string = "partial"

	// BEGIN DEBUGGING
	// IgroupInitiatorListItemConnectivityTracking
	// IgroupInitiatorListItemConnectivityTracking
	// connection_state
	// ConnectionState
	// no_lun_maps
	// END DEBUGGING
	// IgroupInitiatorListItemConnectivityTrackingConnectionStateNoLunMaps captures enum value "no_lun_maps"
	IgroupInitiatorListItemConnectivityTrackingConnectionStateNoLunMaps string = "no_lun_maps"
)

// prop value enum
func (m *IgroupInitiatorListItemConnectivityTracking) validateConnectionStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, igroupInitiatorListItemConnectivityTrackingTypeConnectionStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *IgroupInitiatorListItemConnectivityTracking) validateConnectionState(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectionState) { // not required
		return nil
	}

	// value enum
	if err := m.validateConnectionStateEnum("connectivity_tracking"+"."+"connection_state", "body", m.ConnectionState); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this igroup initiator list item connectivity tracking based on the context it is used
func (m *IgroupInitiatorListItemConnectivityTracking) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectionState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemConnectivityTracking) contextValidateConnectionState(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "connectivity_tracking"+"."+"connection_state", "body", string(m.ConnectionState)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IgroupInitiatorListItemConnectivityTracking) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItemConnectivityTracking) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItemConnectivityTracking
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupInitiatorListItemIgroup The initiator group that directly owns the initiator, which is where modification of the initiator is supported. This property will only be populated when the initiator is a member of a nested initiator group.
//
// swagger:model IgroupInitiatorListItemIgroup
type IgroupInitiatorListItemIgroup struct {

	// links
	Links *IgroupInitiatorListItemIgroupLinks `json:"_links,omitempty"`

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

// Validate validates this igroup initiator list item igroup
func (m *IgroupInitiatorListItemIgroup) Validate(formats strfmt.Registry) error {
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

func (m *IgroupInitiatorListItemIgroup) validateLinks(formats strfmt.Registry) error {
	if swag.IsZero(m.Links) { // not required
		return nil
	}

	if m.Links != nil {
		if err := m.Links.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("igroup" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

func (m *IgroupInitiatorListItemIgroup) validateName(formats strfmt.Registry) error {
	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if err := validate.MinLength("igroup"+"."+"name", "body", m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("igroup"+"."+"name", "body", m.Name, 96); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this igroup initiator list item igroup based on the context it is used
func (m *IgroupInitiatorListItemIgroup) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemIgroup) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

	if m.Links != nil {
		if err := m.Links.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("igroup" + "." + "_links")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IgroupInitiatorListItemIgroup) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItemIgroup) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItemIgroup
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupInitiatorListItemIgroupLinks igroup initiator list item igroup links
//
// swagger:model IgroupInitiatorListItemIgroupLinks
type IgroupInitiatorListItemIgroupLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this igroup initiator list item igroup links
func (m *IgroupInitiatorListItemIgroupLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemIgroupLinks) validateSelf(formats strfmt.Registry) error {
	if swag.IsZero(m.Self) { // not required
		return nil
	}

	if m.Self != nil {
		if err := m.Self.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("igroup" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this igroup initiator list item igroup links based on the context it is used
func (m *IgroupInitiatorListItemIgroupLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemIgroupLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

	if m.Self != nil {
		if err := m.Self.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("igroup" + "." + "_links" + "." + "self")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IgroupInitiatorListItemIgroupLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItemIgroupLinks) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItemIgroupLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupInitiatorListItemLinks igroup initiator list item links
//
// swagger:model IgroupInitiatorListItemLinks
type IgroupInitiatorListItemLinks struct {

	// connectivity tracking
	ConnectivityTracking *IgroupInitiatorListItemLinksConnectivityTracking `json:"connectivity_tracking,omitempty"`

	// self
	Self *IgroupInitiatorListItemLinksSelf `json:"self,omitempty"`
}

// Validate validates this igroup initiator list item links
func (m *IgroupInitiatorListItemLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateConnectivityTracking(formats); err != nil {
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

func (m *IgroupInitiatorListItemLinks) validateConnectivityTracking(formats strfmt.Registry) error {
	if swag.IsZero(m.ConnectivityTracking) { // not required
		return nil
	}

	if m.ConnectivityTracking != nil {
		if err := m.ConnectivityTracking.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "connectivity_tracking")
			}
			return err
		}
	}

	return nil
}

func (m *IgroupInitiatorListItemLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this igroup initiator list item links based on the context it is used
func (m *IgroupInitiatorListItemLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateConnectivityTracking(ctx, formats); err != nil {
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

func (m *IgroupInitiatorListItemLinks) contextValidateConnectivityTracking(ctx context.Context, formats strfmt.Registry) error {

	if m.ConnectivityTracking != nil {
		if err := m.ConnectivityTracking.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("_links" + "." + "connectivity_tracking")
			}
			return err
		}
	}

	return nil
}

func (m *IgroupInitiatorListItemLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *IgroupInitiatorListItemLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItemLinks) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItemLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupInitiatorListItemLinksConnectivityTracking A link to the initiator with connectivity information relevant to its membership of this initiator group.
//
// swagger:model IgroupInitiatorListItemLinksConnectivityTracking
type IgroupInitiatorListItemLinksConnectivityTracking struct {

	// href
	// Example: /api/resourcelink
	// Read Only: true
	Href string `json:"href,omitempty"`
}

// Validate validates this igroup initiator list item links connectivity tracking
func (m *IgroupInitiatorListItemLinksConnectivityTracking) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this igroup initiator list item links connectivity tracking based on the context it is used
func (m *IgroupInitiatorListItemLinksConnectivityTracking) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHref(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemLinksConnectivityTracking) contextValidateHref(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_links"+"."+"connectivity_tracking"+"."+"href", "body", string(m.Href)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IgroupInitiatorListItemLinksConnectivityTracking) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItemLinksConnectivityTracking) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItemLinksConnectivityTracking
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// IgroupInitiatorListItemLinksSelf A link to the initiator where mutations can be made. If the initiator is inherited from a nested initiator group, the link refers to the initiator in the nested initiator group. In this case, mutations of the initiator will be applied to all initiator groups referencing the same nested initiator group.
//
// swagger:model IgroupInitiatorListItemLinksSelf
type IgroupInitiatorListItemLinksSelf struct {

	// href
	// Example: /api/resourcelink
	// Read Only: true
	Href string `json:"href,omitempty"`
}

// Validate validates this igroup initiator list item links self
func (m *IgroupInitiatorListItemLinksSelf) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validate this igroup initiator list item links self based on the context it is used
func (m *IgroupInitiatorListItemLinksSelf) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateHref(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IgroupInitiatorListItemLinksSelf) contextValidateHref(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_links"+"."+"self"+"."+"href", "body", string(m.Href)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IgroupInitiatorListItemLinksSelf) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IgroupInitiatorListItemLinksSelf) UnmarshalBinary(b []byte) error {
	var res IgroupInitiatorListItemLinksSelf
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
