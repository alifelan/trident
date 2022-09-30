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

// NewMultiAdminVerifyRequestDeleteParams creates a new MultiAdminVerifyRequestDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMultiAdminVerifyRequestDeleteParams() *MultiAdminVerifyRequestDeleteParams {
	return &MultiAdminVerifyRequestDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMultiAdminVerifyRequestDeleteParamsWithTimeout creates a new MultiAdminVerifyRequestDeleteParams object
// with the ability to set a timeout on a request.
func NewMultiAdminVerifyRequestDeleteParamsWithTimeout(timeout time.Duration) *MultiAdminVerifyRequestDeleteParams {
	return &MultiAdminVerifyRequestDeleteParams{
		timeout: timeout,
	}
}

// NewMultiAdminVerifyRequestDeleteParamsWithContext creates a new MultiAdminVerifyRequestDeleteParams object
// with the ability to set a context for a request.
func NewMultiAdminVerifyRequestDeleteParamsWithContext(ctx context.Context) *MultiAdminVerifyRequestDeleteParams {
	return &MultiAdminVerifyRequestDeleteParams{
		Context: ctx,
	}
}

// NewMultiAdminVerifyRequestDeleteParamsWithHTTPClient creates a new MultiAdminVerifyRequestDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewMultiAdminVerifyRequestDeleteParamsWithHTTPClient(client *http.Client) *MultiAdminVerifyRequestDeleteParams {
	return &MultiAdminVerifyRequestDeleteParams{
		HTTPClient: client,
	}
}

/*
MultiAdminVerifyRequestDeleteParams contains all the parameters to send to the API endpoint

	for the multi admin verify request delete operation.

	Typically these are written to a http.Request.
*/
type MultiAdminVerifyRequestDeleteParams struct {

	// Index.
	IndexPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the multi admin verify request delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyRequestDeleteParams) WithDefaults() *MultiAdminVerifyRequestDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the multi admin verify request delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyRequestDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the multi admin verify request delete params
func (o *MultiAdminVerifyRequestDeleteParams) WithTimeout(timeout time.Duration) *MultiAdminVerifyRequestDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the multi admin verify request delete params
func (o *MultiAdminVerifyRequestDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the multi admin verify request delete params
func (o *MultiAdminVerifyRequestDeleteParams) WithContext(ctx context.Context) *MultiAdminVerifyRequestDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the multi admin verify request delete params
func (o *MultiAdminVerifyRequestDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the multi admin verify request delete params
func (o *MultiAdminVerifyRequestDeleteParams) WithHTTPClient(client *http.Client) *MultiAdminVerifyRequestDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the multi admin verify request delete params
func (o *MultiAdminVerifyRequestDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndexPathParameter adds the index to the multi admin verify request delete params
func (o *MultiAdminVerifyRequestDeleteParams) WithIndexPathParameter(index string) *MultiAdminVerifyRequestDeleteParams {
	o.SetIndexPathParameter(index)
	return o
}

// SetIndexPathParameter adds the index to the multi admin verify request delete params
func (o *MultiAdminVerifyRequestDeleteParams) SetIndexPathParameter(index string) {
	o.IndexPathParameter = index
}

// WriteToRequest writes these params to a swagger request
func (o *MultiAdminVerifyRequestDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", o.IndexPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
