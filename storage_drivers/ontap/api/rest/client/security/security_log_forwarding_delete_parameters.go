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

// NewSecurityLogForwardingDeleteParams creates a new SecurityLogForwardingDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityLogForwardingDeleteParams() *SecurityLogForwardingDeleteParams {
	return &SecurityLogForwardingDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityLogForwardingDeleteParamsWithTimeout creates a new SecurityLogForwardingDeleteParams object
// with the ability to set a timeout on a request.
func NewSecurityLogForwardingDeleteParamsWithTimeout(timeout time.Duration) *SecurityLogForwardingDeleteParams {
	return &SecurityLogForwardingDeleteParams{
		timeout: timeout,
	}
}

// NewSecurityLogForwardingDeleteParamsWithContext creates a new SecurityLogForwardingDeleteParams object
// with the ability to set a context for a request.
func NewSecurityLogForwardingDeleteParamsWithContext(ctx context.Context) *SecurityLogForwardingDeleteParams {
	return &SecurityLogForwardingDeleteParams{
		Context: ctx,
	}
}

// NewSecurityLogForwardingDeleteParamsWithHTTPClient creates a new SecurityLogForwardingDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityLogForwardingDeleteParamsWithHTTPClient(client *http.Client) *SecurityLogForwardingDeleteParams {
	return &SecurityLogForwardingDeleteParams{
		HTTPClient: client,
	}
}

/*
SecurityLogForwardingDeleteParams contains all the parameters to send to the API endpoint

	for the security log forwarding delete operation.

	Typically these are written to a http.Request.
*/
type SecurityLogForwardingDeleteParams struct {

	/* Address.

	   IP address of remote syslog/splunk server.
	*/
	AddressPathParameter string

	/* Port.

	   Port number of remote syslog/splunk server.
	*/
	PortPathParameter int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security log forwarding delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityLogForwardingDeleteParams) WithDefaults() *SecurityLogForwardingDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security log forwarding delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityLogForwardingDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the security log forwarding delete params
func (o *SecurityLogForwardingDeleteParams) WithTimeout(timeout time.Duration) *SecurityLogForwardingDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security log forwarding delete params
func (o *SecurityLogForwardingDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security log forwarding delete params
func (o *SecurityLogForwardingDeleteParams) WithContext(ctx context.Context) *SecurityLogForwardingDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security log forwarding delete params
func (o *SecurityLogForwardingDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security log forwarding delete params
func (o *SecurityLogForwardingDeleteParams) WithHTTPClient(client *http.Client) *SecurityLogForwardingDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security log forwarding delete params
func (o *SecurityLogForwardingDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAddressPathParameter adds the address to the security log forwarding delete params
func (o *SecurityLogForwardingDeleteParams) WithAddressPathParameter(address string) *SecurityLogForwardingDeleteParams {
	o.SetAddressPathParameter(address)
	return o
}

// SetAddressPathParameter adds the address to the security log forwarding delete params
func (o *SecurityLogForwardingDeleteParams) SetAddressPathParameter(address string) {
	o.AddressPathParameter = address
}

// WithPortPathParameter adds the port to the security log forwarding delete params
func (o *SecurityLogForwardingDeleteParams) WithPortPathParameter(port int64) *SecurityLogForwardingDeleteParams {
	o.SetPortPathParameter(port)
	return o
}

// SetPortPathParameter adds the port to the security log forwarding delete params
func (o *SecurityLogForwardingDeleteParams) SetPortPathParameter(port int64) {
	o.PortPathParameter = port
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityLogForwardingDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param address
	if err := r.SetPathParam("address", o.AddressPathParameter); err != nil {
		return err
	}

	// path param port
	if err := r.SetPathParam("port", swag.FormatInt64(o.PortPathParameter)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
