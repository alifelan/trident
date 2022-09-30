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

// GcpKms gcp kms
//
// swagger:model gcp_kms
type GcpKms struct {

	// links
	Links *GcpKmsLinks `json:"_links,omitempty"`

	// Google Cloud application's service account credentials required to access the specified KMS. It is a JSON file containing an email address and the private key of the service account holder.
	// Example: { type: service_account, project_id: project-id, private_key_id: key-id, private_key: -----BEGIN PRIVATE KEY-----\nprivate-key\n-----END PRIVATE KEY-----\n, client_email: service-account-email, client_id: client-id, auth_uri: https://accounts.google.com/o/oauth2/auth, token_uri: https://accounts.google.com/o/oauth2/token, auth_provider_x509_cert_url: https://www.googleapis.com/oauth2/v1/certs, client_x509_cert_url: https://www.googleapis.com/robot/v1/metadata/x509/service-account-email }
	// Format: password
	ApplicationCredentials strfmt.Password `json:"application_credentials,omitempty"`

	// ekmip reachability
	EkmipReachability []*GcpKmsEkmipReachabilityItems0 `json:"ekmip_reachability,omitempty"`

	// google reachability
	GoogleReachability *GcpKmsGoogleReachability `json:"google_reachability,omitempty"`

	// Key Identifier of Google Cloud KMS key encryption key.
	// Example: cryptokey1
	KeyName string `json:"key_name,omitempty"`

	// Google Cloud KMS key ring location.
	// Example: global
	KeyRingLocation string `json:"key_ring_location,omitempty"`

	// Google Cloud KMS key ring name of the deployed Google Cloud application.
	// Example: gcpapp1-keyring
	KeyRingName string `json:"key_ring_name,omitempty"`

	// Google Cloud project (application) ID of the deployed Google Cloud application that has appropriate access to the Google Cloud KMS.
	// Example: gcpapp1
	ProjectID string `json:"project_id,omitempty"`

	// Proxy host name.
	// Example: proxy.eng.com
	ProxyHost string `json:"proxy_host,omitempty"`

	// Proxy password. Password is not audited.
	// Example: proxypassword
	ProxyPassword string `json:"proxy_password,omitempty"`

	// Proxy port number.
	// Example: 1234
	ProxyPort int64 `json:"proxy_port,omitempty"`

	// Type of proxy.
	// Example: http
	// Enum: [http https]
	ProxyType string `json:"proxy_type,omitempty"`

	// Proxy username.
	// Example: proxyuser
	ProxyUsername string `json:"proxy_username,omitempty"`

	// Set to "svm" for interfaces owned by an SVM. Otherwise, set to "cluster".
	// Enum: [svm cluster]
	Scope string `json:"scope,omitempty"`

	// state
	State *GcpKmsStateType `json:"state,omitempty"`

	// svm
	Svm *GcpKmsSvm `json:"svm,omitempty"`

	// A unique identifier for the Google Cloud KMS.
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	// Read Only: true
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this gcp kms
func (m *GcpKms) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateApplicationCredentials(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEkmipReachability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGoogleReachability(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProxyType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScope(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
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

func (m *GcpKms) validateLinks(formats strfmt.Registry) error {
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

func (m *GcpKms) validateApplicationCredentials(formats strfmt.Registry) error {
	if swag.IsZero(m.ApplicationCredentials) { // not required
		return nil
	}

	if err := validate.FormatOf("application_credentials", "body", "password", m.ApplicationCredentials.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GcpKms) validateEkmipReachability(formats strfmt.Registry) error {
	if swag.IsZero(m.EkmipReachability) { // not required
		return nil
	}

	for i := 0; i < len(m.EkmipReachability); i++ {
		if swag.IsZero(m.EkmipReachability[i]) { // not required
			continue
		}

		if m.EkmipReachability[i] != nil {
			if err := m.EkmipReachability[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ekmip_reachability" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GcpKms) validateGoogleReachability(formats strfmt.Registry) error {
	if swag.IsZero(m.GoogleReachability) { // not required
		return nil
	}

	if m.GoogleReachability != nil {
		if err := m.GoogleReachability.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("google_reachability")
			}
			return err
		}
	}

	return nil
}

var gcpKmsTypeProxyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["http","https"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gcpKmsTypeProxyTypePropEnum = append(gcpKmsTypeProxyTypePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// gcp_kms
	// GcpKms
	// proxy_type
	// ProxyType
	// http
	// END DEBUGGING
	// GcpKmsProxyTypeHTTP captures enum value "http"
	GcpKmsProxyTypeHTTP string = "http"

	// BEGIN DEBUGGING
	// gcp_kms
	// GcpKms
	// proxy_type
	// ProxyType
	// https
	// END DEBUGGING
	// GcpKmsProxyTypeHTTPS captures enum value "https"
	GcpKmsProxyTypeHTTPS string = "https"
)

// prop value enum
func (m *GcpKms) validateProxyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, gcpKmsTypeProxyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GcpKms) validateProxyType(formats strfmt.Registry) error {
	if swag.IsZero(m.ProxyType) { // not required
		return nil
	}

	// value enum
	if err := m.validateProxyTypeEnum("proxy_type", "body", m.ProxyType); err != nil {
		return err
	}

	return nil
}

var gcpKmsTypeScopePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["svm","cluster"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		gcpKmsTypeScopePropEnum = append(gcpKmsTypeScopePropEnum, v)
	}
}

const (

	// BEGIN DEBUGGING
	// gcp_kms
	// GcpKms
	// scope
	// Scope
	// svm
	// END DEBUGGING
	// GcpKmsScopeSvm captures enum value "svm"
	GcpKmsScopeSvm string = "svm"

	// BEGIN DEBUGGING
	// gcp_kms
	// GcpKms
	// scope
	// Scope
	// cluster
	// END DEBUGGING
	// GcpKmsScopeCluster captures enum value "cluster"
	GcpKmsScopeCluster string = "cluster"
)

// prop value enum
func (m *GcpKms) validateScopeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, gcpKmsTypeScopePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *GcpKms) validateScope(formats strfmt.Registry) error {
	if swag.IsZero(m.Scope) { // not required
		return nil
	}

	// value enum
	if err := m.validateScopeEnum("scope", "body", m.Scope); err != nil {
		return err
	}

	return nil
}

func (m *GcpKms) validateState(formats strfmt.Registry) error {
	if swag.IsZero(m.State) { // not required
		return nil
	}

	if m.State != nil {
		if err := m.State.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *GcpKms) validateSvm(formats strfmt.Registry) error {
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

// ContextValidate validate this gcp kms based on the context it is used
func (m *GcpKms) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateEkmipReachability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateGoogleReachability(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateState(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSvm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateUUID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpKms) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GcpKms) contextValidateEkmipReachability(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.EkmipReachability); i++ {

		if m.EkmipReachability[i] != nil {
			if err := m.EkmipReachability[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ekmip_reachability" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *GcpKms) contextValidateGoogleReachability(ctx context.Context, formats strfmt.Registry) error {

	if m.GoogleReachability != nil {
		if err := m.GoogleReachability.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("google_reachability")
			}
			return err
		}
	}

	return nil
}

func (m *GcpKms) contextValidateState(ctx context.Context, formats strfmt.Registry) error {

	if m.State != nil {
		if err := m.State.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("state")
			}
			return err
		}
	}

	return nil
}

func (m *GcpKms) contextValidateSvm(ctx context.Context, formats strfmt.Registry) error {

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

func (m *GcpKms) contextValidateUUID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "uuid", "body", string(m.UUID)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GcpKms) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpKms) UnmarshalBinary(b []byte) error {
	var res GcpKms
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GcpKmsEkmipReachabilityItems0 Provides the connectivity status for the given SVM on the given node to all EKMIP servers configured on all nodes of the cluster.
// This is an advanced property; there is an added cost to retrieving its value. The property is not populated for either a collection GET or an instance GET unless it is explicitly requested using the `fields` query parameter or GET for all advanced properties is enabled.
//
// swagger:model GcpKmsEkmipReachabilityItems0
type GcpKmsEkmipReachabilityItems0 struct {

	// Code corresponding to the error message. Returns a 0 if a given SVM is able to communicate to the EKMIP servers of all of the nodes in the cluster.
	// Example: 346758
	Code int64 `json:"code,omitempty"`

	// Error message set when cluster-wide EKMIP server availability from the given SVM and node is false.
	// Example: embedded KMIP server status unavailable on node.
	Message string `json:"message,omitempty"`

	// node
	Node *GcpKmsEkmipReachabilityItems0Node `json:"node,omitempty"`

	// Set to true if the given SVM on the given node is able to communicate to all EKMIP servers configured on all nodes in the cluster.
	Reachable bool `json:"reachable,omitempty"`
}

// Validate validates this gcp kms ekmip reachability items0
func (m *GcpKmsEkmipReachabilityItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpKmsEkmipReachabilityItems0) validateNode(formats strfmt.Registry) error {
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

// ContextValidate validate this gcp kms ekmip reachability items0 based on the context it is used
func (m *GcpKmsEkmipReachabilityItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateNode(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpKmsEkmipReachabilityItems0) contextValidateNode(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GcpKmsEkmipReachabilityItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpKmsEkmipReachabilityItems0) UnmarshalBinary(b []byte) error {
	var res GcpKmsEkmipReachabilityItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GcpKmsEkmipReachabilityItems0Node gcp kms ekmip reachability items0 node
//
// swagger:model GcpKmsEkmipReachabilityItems0Node
type GcpKmsEkmipReachabilityItems0Node struct {

	// links
	Links *GcpKmsEkmipReachabilityItems0NodeLinks `json:"_links,omitempty"`

	// name
	// Example: node1
	Name string `json:"name,omitempty"`

	// uuid
	// Example: 1cd8a442-86d1-11e0-ae1c-123478563412
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this gcp kms ekmip reachability items0 node
func (m *GcpKmsEkmipReachabilityItems0Node) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpKmsEkmipReachabilityItems0Node) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this gcp kms ekmip reachability items0 node based on the context it is used
func (m *GcpKmsEkmipReachabilityItems0Node) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpKmsEkmipReachabilityItems0Node) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GcpKmsEkmipReachabilityItems0Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpKmsEkmipReachabilityItems0Node) UnmarshalBinary(b []byte) error {
	var res GcpKmsEkmipReachabilityItems0Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GcpKmsEkmipReachabilityItems0NodeLinks gcp kms ekmip reachability items0 node links
//
// swagger:model GcpKmsEkmipReachabilityItems0NodeLinks
type GcpKmsEkmipReachabilityItems0NodeLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this gcp kms ekmip reachability items0 node links
func (m *GcpKmsEkmipReachabilityItems0NodeLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpKmsEkmipReachabilityItems0NodeLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this gcp kms ekmip reachability items0 node links based on the context it is used
func (m *GcpKmsEkmipReachabilityItems0NodeLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpKmsEkmipReachabilityItems0NodeLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GcpKmsEkmipReachabilityItems0NodeLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpKmsEkmipReachabilityItems0NodeLinks) UnmarshalBinary(b []byte) error {
	var res GcpKmsEkmipReachabilityItems0NodeLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GcpKmsGoogleReachability Indicates whether or not the Google Cloud KMS service is reachable from all nodes in the cluster.
// This is an advanced property; there is an added cost to retrieving its value. The property is not populated for either a collection GET or an instance GET unless it is explicitly requested using the `fields` query parameter or GET for all advanced properties is enabled.
//
// swagger:model GcpKmsGoogleReachability
type GcpKmsGoogleReachability struct {

	// Code corresponding to the error message. Returns a 0 if Google Cloud KMS service is reachable from all nodes in the cluster.
	// Example: 346758
	Code int64 `json:"code,omitempty"`

	// Set to the error message when 'reachable' is false.
	// Example: Google Cloud KMS service is not reachable from all nodes - \u003creason\u003e.
	Message string `json:"message,omitempty"`

	// Set to true if the Google Cloud KMS service is reachable from all nodes of the cluster.
	Reachable bool `json:"reachable,omitempty"`
}

// Validate validates this gcp kms google reachability
func (m *GcpKmsGoogleReachability) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gcp kms google reachability based on context it is used
func (m *GcpKmsGoogleReachability) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GcpKmsGoogleReachability) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpKmsGoogleReachability) UnmarshalBinary(b []byte) error {
	var res GcpKmsGoogleReachability
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GcpKmsLinks gcp kms links
//
// swagger:model GcpKmsLinks
type GcpKmsLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this gcp kms links
func (m *GcpKmsLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpKmsLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this gcp kms links based on the context it is used
func (m *GcpKmsLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpKmsLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GcpKmsLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpKmsLinks) UnmarshalBinary(b []byte) error {
	var res GcpKmsLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GcpKmsStateType Google Cloud Key Management Services is a cloud key management service (KMS) that provides a secure store for encryption keys. This object indicates whether or not the Google Cloud KMS key protection is available on all nodes in the cluster.
// This is an advanced property; there is an added cost to retrieving its value. The property is not populated for either a collection GET or an instance GET unless it is explicitly requested using the `fields` query parameter or GET for all advanced properties is enabled.
//
// swagger:model GcpKmsStateType
type GcpKmsStateType struct {

	// Set to true when Google Cloud KMS key protection is available on all nodes of the cluster.
	ClusterState bool `json:"cluster_state,omitempty"`

	// Error code corresponding to the status message. Returns 0 if Google Cloud KMS key protection is available in all nodes of the cluster.
	// Example: 346758
	Code int64 `json:"code,omitempty"`

	// Error message set when top-level internal key protection key (KEK) availability on cluster is false.
	// Example: Top-level internal key protection key (KEK) is unavailable on the following nodes with the associated reasons: Node: node1. Reason: No volumes created yet for the SVM. Wrapped KEK status will be available after creating encrypted volumes.
	Message string `json:"message,omitempty"`
}

// Validate validates this gcp kms state type
func (m *GcpKmsStateType) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this gcp kms state type based on context it is used
func (m *GcpKmsStateType) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GcpKmsStateType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpKmsStateType) UnmarshalBinary(b []byte) error {
	var res GcpKmsStateType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GcpKmsSvm gcp kms svm
//
// swagger:model GcpKmsSvm
type GcpKmsSvm struct {

	// links
	Links *GcpKmsSvmLinks `json:"_links,omitempty"`

	// The name of the SVM.
	//
	// Example: svm1
	Name string `json:"name,omitempty"`

	// The unique identifier of the SVM.
	//
	// Example: 02c9e252-41be-11e9-81d5-00a0986138f7
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this gcp kms svm
func (m *GcpKmsSvm) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateLinks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpKmsSvm) validateLinks(formats strfmt.Registry) error {
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

// ContextValidate validate this gcp kms svm based on the context it is used
func (m *GcpKmsSvm) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateLinks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpKmsSvm) contextValidateLinks(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GcpKmsSvm) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpKmsSvm) UnmarshalBinary(b []byte) error {
	var res GcpKmsSvm
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// GcpKmsSvmLinks gcp kms svm links
//
// swagger:model GcpKmsSvmLinks
type GcpKmsSvmLinks struct {

	// self
	Self *Href `json:"self,omitempty"`
}

// Validate validates this gcp kms svm links
func (m *GcpKmsSvmLinks) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateSelf(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpKmsSvmLinks) validateSelf(formats strfmt.Registry) error {
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

// ContextValidate validate this gcp kms svm links based on the context it is used
func (m *GcpKmsSvmLinks) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateSelf(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GcpKmsSvmLinks) contextValidateSelf(ctx context.Context, formats strfmt.Registry) error {

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
func (m *GcpKmsSvmLinks) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GcpKmsSvmLinks) UnmarshalBinary(b []byte) error {
	var res GcpKmsSvmLinks
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
