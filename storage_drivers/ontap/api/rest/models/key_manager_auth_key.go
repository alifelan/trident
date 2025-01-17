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

// KeyManagerAuthKey key manager auth key
//
// swagger:model key_manager_auth_key
type KeyManagerAuthKey struct {

	// links
	Links *SelfLink `json:"_links,omitempty"`

	// Key identifier.
	// Example: 00000000000000000100000000000123456789asdfghjklqwertyuio123456780000000000000090
	KeyID *string `json:"key_id,omitempty"`

	// Optional parameter to define key-tag for the authentication key, length 0-32 characters.
	// Example: Authentication Key-01
	KeyTag *string `json:"key_tag,omitempty"`

	// Authentication passphrase, length 20-32 characters. May contain the '=' character.
	// Example: AuthenticationKey_01
	// Format: password
	Passphrase *strfmt.Password `json:"passphrase,omitempty"`
}

// Validate validates this key manager auth key
func (m *KeyManagerAuthKey) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePassphrase(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyManagerAuthKey) validateLinks(formats strfmt.Registry) error {
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

func (m *KeyManagerAuthKey) validatePassphrase(formats strfmt.Registry) error {
	if swag.IsZero(m.Passphrase) { // not required
		return nil
	}

	if err := validate.FormatOf("passphrase", "body", "password", m.Passphrase.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this key manager auth key based on the context it is used
func (m *KeyManagerAuthKey) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *KeyManagerAuthKey) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *KeyManagerAuthKey) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *KeyManagerAuthKey) UnmarshalBinary(b []byte) error {
	var res KeyManagerAuthKey
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
