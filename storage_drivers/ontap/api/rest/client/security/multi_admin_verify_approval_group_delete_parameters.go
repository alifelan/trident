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

// NewMultiAdminVerifyApprovalGroupDeleteParams creates a new MultiAdminVerifyApprovalGroupDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMultiAdminVerifyApprovalGroupDeleteParams() *MultiAdminVerifyApprovalGroupDeleteParams {
	return &MultiAdminVerifyApprovalGroupDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMultiAdminVerifyApprovalGroupDeleteParamsWithTimeout creates a new MultiAdminVerifyApprovalGroupDeleteParams object
// with the ability to set a timeout on a request.
func NewMultiAdminVerifyApprovalGroupDeleteParamsWithTimeout(timeout time.Duration) *MultiAdminVerifyApprovalGroupDeleteParams {
	return &MultiAdminVerifyApprovalGroupDeleteParams{
		timeout: timeout,
	}
}

// NewMultiAdminVerifyApprovalGroupDeleteParamsWithContext creates a new MultiAdminVerifyApprovalGroupDeleteParams object
// with the ability to set a context for a request.
func NewMultiAdminVerifyApprovalGroupDeleteParamsWithContext(ctx context.Context) *MultiAdminVerifyApprovalGroupDeleteParams {
	return &MultiAdminVerifyApprovalGroupDeleteParams{
		Context: ctx,
	}
}

// NewMultiAdminVerifyApprovalGroupDeleteParamsWithHTTPClient creates a new MultiAdminVerifyApprovalGroupDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewMultiAdminVerifyApprovalGroupDeleteParamsWithHTTPClient(client *http.Client) *MultiAdminVerifyApprovalGroupDeleteParams {
	return &MultiAdminVerifyApprovalGroupDeleteParams{
		HTTPClient: client,
	}
}

/*
MultiAdminVerifyApprovalGroupDeleteParams contains all the parameters to send to the API endpoint

	for the multi admin verify approval group delete operation.

	Typically these are written to a http.Request.
*/
type MultiAdminVerifyApprovalGroupDeleteParams struct {

	// Name.
	NamePathParameter string

	// OwnerUUID.
	OwnerUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the multi admin verify approval group delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyApprovalGroupDeleteParams) WithDefaults() *MultiAdminVerifyApprovalGroupDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the multi admin verify approval group delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyApprovalGroupDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the multi admin verify approval group delete params
func (o *MultiAdminVerifyApprovalGroupDeleteParams) WithTimeout(timeout time.Duration) *MultiAdminVerifyApprovalGroupDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the multi admin verify approval group delete params
func (o *MultiAdminVerifyApprovalGroupDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the multi admin verify approval group delete params
func (o *MultiAdminVerifyApprovalGroupDeleteParams) WithContext(ctx context.Context) *MultiAdminVerifyApprovalGroupDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the multi admin verify approval group delete params
func (o *MultiAdminVerifyApprovalGroupDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the multi admin verify approval group delete params
func (o *MultiAdminVerifyApprovalGroupDeleteParams) WithHTTPClient(client *http.Client) *MultiAdminVerifyApprovalGroupDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the multi admin verify approval group delete params
func (o *MultiAdminVerifyApprovalGroupDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamePathParameter adds the name to the multi admin verify approval group delete params
func (o *MultiAdminVerifyApprovalGroupDeleteParams) WithNamePathParameter(name string) *MultiAdminVerifyApprovalGroupDeleteParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the multi admin verify approval group delete params
func (o *MultiAdminVerifyApprovalGroupDeleteParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithOwnerUUIDPathParameter adds the ownerUUID to the multi admin verify approval group delete params
func (o *MultiAdminVerifyApprovalGroupDeleteParams) WithOwnerUUIDPathParameter(ownerUUID string) *MultiAdminVerifyApprovalGroupDeleteParams {
	o.SetOwnerUUIDPathParameter(ownerUUID)
	return o
}

// SetOwnerUUIDPathParameter adds the ownerUuid to the multi admin verify approval group delete params
func (o *MultiAdminVerifyApprovalGroupDeleteParams) SetOwnerUUIDPathParameter(ownerUUID string) {
	o.OwnerUUIDPathParameter = ownerUUID
}

// WriteToRequest writes these params to a swagger request
func (o *MultiAdminVerifyApprovalGroupDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	// path param owner.uuid
	if err := r.SetPathParam("owner.uuid", o.OwnerUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
