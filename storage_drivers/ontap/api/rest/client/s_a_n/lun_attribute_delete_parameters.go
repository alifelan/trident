// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewLunAttributeDeleteParams creates a new LunAttributeDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLunAttributeDeleteParams() *LunAttributeDeleteParams {
	return &LunAttributeDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLunAttributeDeleteParamsWithTimeout creates a new LunAttributeDeleteParams object
// with the ability to set a timeout on a request.
func NewLunAttributeDeleteParamsWithTimeout(timeout time.Duration) *LunAttributeDeleteParams {
	return &LunAttributeDeleteParams{
		timeout: timeout,
	}
}

// NewLunAttributeDeleteParamsWithContext creates a new LunAttributeDeleteParams object
// with the ability to set a context for a request.
func NewLunAttributeDeleteParamsWithContext(ctx context.Context) *LunAttributeDeleteParams {
	return &LunAttributeDeleteParams{
		Context: ctx,
	}
}

// NewLunAttributeDeleteParamsWithHTTPClient creates a new LunAttributeDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewLunAttributeDeleteParamsWithHTTPClient(client *http.Client) *LunAttributeDeleteParams {
	return &LunAttributeDeleteParams{
		HTTPClient: client,
	}
}

/*
LunAttributeDeleteParams contains all the parameters to send to the API endpoint

	for the lun attribute delete operation.

	Typically these are written to a http.Request.
*/
type LunAttributeDeleteParams struct {

	/* LunUUID.

	   The unique identifier of the LUN.

	*/
	LunUUIDPathParameter string

	/* Name.

	   The name of the attribute.

	*/
	NamePathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lun attribute delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunAttributeDeleteParams) WithDefaults() *LunAttributeDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lun attribute delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunAttributeDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the lun attribute delete params
func (o *LunAttributeDeleteParams) WithTimeout(timeout time.Duration) *LunAttributeDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lun attribute delete params
func (o *LunAttributeDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lun attribute delete params
func (o *LunAttributeDeleteParams) WithContext(ctx context.Context) *LunAttributeDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lun attribute delete params
func (o *LunAttributeDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lun attribute delete params
func (o *LunAttributeDeleteParams) WithHTTPClient(client *http.Client) *LunAttributeDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lun attribute delete params
func (o *LunAttributeDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLunUUIDPathParameter adds the lunUUID to the lun attribute delete params
func (o *LunAttributeDeleteParams) WithLunUUIDPathParameter(lunUUID string) *LunAttributeDeleteParams {
	o.SetLunUUIDPathParameter(lunUUID)
	return o
}

// SetLunUUIDPathParameter adds the lunUuid to the lun attribute delete params
func (o *LunAttributeDeleteParams) SetLunUUIDPathParameter(lunUUID string) {
	o.LunUUIDPathParameter = lunUUID
}

// WithNamePathParameter adds the name to the lun attribute delete params
func (o *LunAttributeDeleteParams) WithNamePathParameter(name string) *LunAttributeDeleteParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the lun attribute delete params
func (o *LunAttributeDeleteParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WriteToRequest writes these params to a swagger request
func (o *LunAttributeDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param lun.uuid
	if err := r.SetPathParam("lun.uuid", o.LunUUIDPathParameter); err != nil {
		return err
	}

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
