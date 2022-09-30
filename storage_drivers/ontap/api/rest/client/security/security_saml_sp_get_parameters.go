// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSecuritySamlSpGetParams creates a new SecuritySamlSpGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecuritySamlSpGetParams() *SecuritySamlSpGetParams {
	return &SecuritySamlSpGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecuritySamlSpGetParamsWithTimeout creates a new SecuritySamlSpGetParams object
// with the ability to set a timeout on a request.
func NewSecuritySamlSpGetParamsWithTimeout(timeout time.Duration) *SecuritySamlSpGetParams {
	return &SecuritySamlSpGetParams{
		timeout: timeout,
	}
}

// NewSecuritySamlSpGetParamsWithContext creates a new SecuritySamlSpGetParams object
// with the ability to set a context for a request.
func NewSecuritySamlSpGetParamsWithContext(ctx context.Context) *SecuritySamlSpGetParams {
	return &SecuritySamlSpGetParams{
		Context: ctx,
	}
}

// NewSecuritySamlSpGetParamsWithHTTPClient creates a new SecuritySamlSpGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecuritySamlSpGetParamsWithHTTPClient(client *http.Client) *SecuritySamlSpGetParams {
	return &SecuritySamlSpGetParams{
		HTTPClient: client,
	}
}

/*
SecuritySamlSpGetParams contains all the parameters to send to the API endpoint

	for the security saml sp get operation.

	Typically these are written to a http.Request.
*/
type SecuritySamlSpGetParams struct {

	/* CertificateCa.

	   Filter by certificate.ca
	*/
	CertificateCaQueryParameter *string

	/* CertificateCommonName.

	   Filter by certificate.common_name
	*/
	CertificateCommonNameQueryParameter *string

	/* CertificateSerialNumber.

	   Filter by certificate.serial_number
	*/
	CertificateSerialNumberQueryParameter *string

	/* Enabled.

	   Filter by enabled
	*/
	EnabledQueryParameter *bool

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Host.

	   Filter by host
	*/
	HostQueryParameter *string

	/* IdpURI.

	   Filter by idp_uri
	*/
	IdpURIQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security saml sp get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecuritySamlSpGetParams) WithDefaults() *SecuritySamlSpGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security saml sp get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecuritySamlSpGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security saml sp get params
func (o *SecuritySamlSpGetParams) WithTimeout(timeout time.Duration) *SecuritySamlSpGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security saml sp get params
func (o *SecuritySamlSpGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security saml sp get params
func (o *SecuritySamlSpGetParams) WithContext(ctx context.Context) *SecuritySamlSpGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security saml sp get params
func (o *SecuritySamlSpGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security saml sp get params
func (o *SecuritySamlSpGetParams) WithHTTPClient(client *http.Client) *SecuritySamlSpGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security saml sp get params
func (o *SecuritySamlSpGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCertificateCaQueryParameter adds the certificateCa to the security saml sp get params
func (o *SecuritySamlSpGetParams) WithCertificateCaQueryParameter(certificateCa *string) *SecuritySamlSpGetParams {
	o.SetCertificateCaQueryParameter(certificateCa)
	return o
}

// SetCertificateCaQueryParameter adds the certificateCa to the security saml sp get params
func (o *SecuritySamlSpGetParams) SetCertificateCaQueryParameter(certificateCa *string) {
	o.CertificateCaQueryParameter = certificateCa
}

// WithCertificateCommonNameQueryParameter adds the certificateCommonName to the security saml sp get params
func (o *SecuritySamlSpGetParams) WithCertificateCommonNameQueryParameter(certificateCommonName *string) *SecuritySamlSpGetParams {
	o.SetCertificateCommonNameQueryParameter(certificateCommonName)
	return o
}

// SetCertificateCommonNameQueryParameter adds the certificateCommonName to the security saml sp get params
func (o *SecuritySamlSpGetParams) SetCertificateCommonNameQueryParameter(certificateCommonName *string) {
	o.CertificateCommonNameQueryParameter = certificateCommonName
}

// WithCertificateSerialNumberQueryParameter adds the certificateSerialNumber to the security saml sp get params
func (o *SecuritySamlSpGetParams) WithCertificateSerialNumberQueryParameter(certificateSerialNumber *string) *SecuritySamlSpGetParams {
	o.SetCertificateSerialNumberQueryParameter(certificateSerialNumber)
	return o
}

// SetCertificateSerialNumberQueryParameter adds the certificateSerialNumber to the security saml sp get params
func (o *SecuritySamlSpGetParams) SetCertificateSerialNumberQueryParameter(certificateSerialNumber *string) {
	o.CertificateSerialNumberQueryParameter = certificateSerialNumber
}

// WithEnabledQueryParameter adds the enabled to the security saml sp get params
func (o *SecuritySamlSpGetParams) WithEnabledQueryParameter(enabled *bool) *SecuritySamlSpGetParams {
	o.SetEnabledQueryParameter(enabled)
	return o
}

// SetEnabledQueryParameter adds the enabled to the security saml sp get params
func (o *SecuritySamlSpGetParams) SetEnabledQueryParameter(enabled *bool) {
	o.EnabledQueryParameter = enabled
}

// WithFieldsQueryParameter adds the fields to the security saml sp get params
func (o *SecuritySamlSpGetParams) WithFieldsQueryParameter(fields []string) *SecuritySamlSpGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the security saml sp get params
func (o *SecuritySamlSpGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithHostQueryParameter adds the host to the security saml sp get params
func (o *SecuritySamlSpGetParams) WithHostQueryParameter(host *string) *SecuritySamlSpGetParams {
	o.SetHostQueryParameter(host)
	return o
}

// SetHostQueryParameter adds the host to the security saml sp get params
func (o *SecuritySamlSpGetParams) SetHostQueryParameter(host *string) {
	o.HostQueryParameter = host
}

// WithIdpURIQueryParameter adds the idpURI to the security saml sp get params
func (o *SecuritySamlSpGetParams) WithIdpURIQueryParameter(idpURI *string) *SecuritySamlSpGetParams {
	o.SetIdpURIQueryParameter(idpURI)
	return o
}

// SetIdpURIQueryParameter adds the idpUri to the security saml sp get params
func (o *SecuritySamlSpGetParams) SetIdpURIQueryParameter(idpURI *string) {
	o.IdpURIQueryParameter = idpURI
}

// WriteToRequest writes these params to a swagger request
func (o *SecuritySamlSpGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CertificateCaQueryParameter != nil {

		// query param certificate.ca
		var qrCertificateCa string

		if o.CertificateCaQueryParameter != nil {
			qrCertificateCa = *o.CertificateCaQueryParameter
		}
		qCertificateCa := qrCertificateCa
		if qCertificateCa != "" {

			if err := r.SetQueryParam("certificate.ca", qCertificateCa); err != nil {
				return err
			}
		}
	}

	if o.CertificateCommonNameQueryParameter != nil {

		// query param certificate.common_name
		var qrCertificateCommonName string

		if o.CertificateCommonNameQueryParameter != nil {
			qrCertificateCommonName = *o.CertificateCommonNameQueryParameter
		}
		qCertificateCommonName := qrCertificateCommonName
		if qCertificateCommonName != "" {

			if err := r.SetQueryParam("certificate.common_name", qCertificateCommonName); err != nil {
				return err
			}
		}
	}

	if o.CertificateSerialNumberQueryParameter != nil {

		// query param certificate.serial_number
		var qrCertificateSerialNumber string

		if o.CertificateSerialNumberQueryParameter != nil {
			qrCertificateSerialNumber = *o.CertificateSerialNumberQueryParameter
		}
		qCertificateSerialNumber := qrCertificateSerialNumber
		if qCertificateSerialNumber != "" {

			if err := r.SetQueryParam("certificate.serial_number", qCertificateSerialNumber); err != nil {
				return err
			}
		}
	}

	if o.EnabledQueryParameter != nil {

		// query param enabled
		var qrEnabled bool

		if o.EnabledQueryParameter != nil {
			qrEnabled = *o.EnabledQueryParameter
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.HostQueryParameter != nil {

		// query param host
		var qrHost string

		if o.HostQueryParameter != nil {
			qrHost = *o.HostQueryParameter
		}
		qHost := qrHost
		if qHost != "" {

			if err := r.SetQueryParam("host", qHost); err != nil {
				return err
			}
		}
	}

	if o.IdpURIQueryParameter != nil {

		// query param idp_uri
		var qrIdpURI string

		if o.IdpURIQueryParameter != nil {
			qrIdpURI = *o.IdpURIQueryParameter
		}
		qIdpURI := qrIdpURI
		if qIdpURI != "" {

			if err := r.SetQueryParam("idp_uri", qIdpURI); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSecuritySamlSpGet binds the parameter fields
func (o *SecuritySamlSpGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
