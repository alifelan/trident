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
)

// NewSecurityCertificateDeleteParams creates a new SecurityCertificateDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityCertificateDeleteParams() *SecurityCertificateDeleteParams {
	return &SecurityCertificateDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityCertificateDeleteParamsWithTimeout creates a new SecurityCertificateDeleteParams object
// with the ability to set a timeout on a request.
func NewSecurityCertificateDeleteParamsWithTimeout(timeout time.Duration) *SecurityCertificateDeleteParams {
	return &SecurityCertificateDeleteParams{
		timeout: timeout,
	}
}

// NewSecurityCertificateDeleteParamsWithContext creates a new SecurityCertificateDeleteParams object
// with the ability to set a context for a request.
func NewSecurityCertificateDeleteParamsWithContext(ctx context.Context) *SecurityCertificateDeleteParams {
	return &SecurityCertificateDeleteParams{
		Context: ctx,
	}
}

// NewSecurityCertificateDeleteParamsWithHTTPClient creates a new SecurityCertificateDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityCertificateDeleteParamsWithHTTPClient(client *http.Client) *SecurityCertificateDeleteParams {
	return &SecurityCertificateDeleteParams{
		HTTPClient: client,
	}
}

/*
SecurityCertificateDeleteParams contains all the parameters to send to the API endpoint

	for the security certificate delete operation.

	Typically these are written to a http.Request.
*/
type SecurityCertificateDeleteParams struct {

	/* UUID.

	   Certificate UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security certificate delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityCertificateDeleteParams) WithDefaults() *SecurityCertificateDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security certificate delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityCertificateDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security certificate delete params
func (o *SecurityCertificateDeleteParams) WithTimeout(timeout time.Duration) *SecurityCertificateDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security certificate delete params
func (o *SecurityCertificateDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security certificate delete params
func (o *SecurityCertificateDeleteParams) WithContext(ctx context.Context) *SecurityCertificateDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security certificate delete params
func (o *SecurityCertificateDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security certificate delete params
func (o *SecurityCertificateDeleteParams) WithHTTPClient(client *http.Client) *SecurityCertificateDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security certificate delete params
func (o *SecurityCertificateDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithUUIDPathParameter adds the uuid to the security certificate delete params
func (o *SecurityCertificateDeleteParams) WithUUIDPathParameter(uuid string) *SecurityCertificateDeleteParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the security certificate delete params
func (o *SecurityCertificateDeleteParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityCertificateDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
