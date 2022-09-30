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

// NewMultiAdminVerifyRuleDeleteParams creates a new MultiAdminVerifyRuleDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMultiAdminVerifyRuleDeleteParams() *MultiAdminVerifyRuleDeleteParams {
	return &MultiAdminVerifyRuleDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMultiAdminVerifyRuleDeleteParamsWithTimeout creates a new MultiAdminVerifyRuleDeleteParams object
// with the ability to set a timeout on a request.
func NewMultiAdminVerifyRuleDeleteParamsWithTimeout(timeout time.Duration) *MultiAdminVerifyRuleDeleteParams {
	return &MultiAdminVerifyRuleDeleteParams{
		timeout: timeout,
	}
}

// NewMultiAdminVerifyRuleDeleteParamsWithContext creates a new MultiAdminVerifyRuleDeleteParams object
// with the ability to set a context for a request.
func NewMultiAdminVerifyRuleDeleteParamsWithContext(ctx context.Context) *MultiAdminVerifyRuleDeleteParams {
	return &MultiAdminVerifyRuleDeleteParams{
		Context: ctx,
	}
}

// NewMultiAdminVerifyRuleDeleteParamsWithHTTPClient creates a new MultiAdminVerifyRuleDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewMultiAdminVerifyRuleDeleteParamsWithHTTPClient(client *http.Client) *MultiAdminVerifyRuleDeleteParams {
	return &MultiAdminVerifyRuleDeleteParams{
		HTTPClient: client,
	}
}

/*
MultiAdminVerifyRuleDeleteParams contains all the parameters to send to the API endpoint

	for the multi admin verify rule delete operation.

	Typically these are written to a http.Request.
*/
type MultiAdminVerifyRuleDeleteParams struct {

	// Operation.
	OperationPathParameter string

	// OwnerUUID.
	OwnerUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the multi admin verify rule delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyRuleDeleteParams) WithDefaults() *MultiAdminVerifyRuleDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the multi admin verify rule delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyRuleDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the multi admin verify rule delete params
func (o *MultiAdminVerifyRuleDeleteParams) WithTimeout(timeout time.Duration) *MultiAdminVerifyRuleDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the multi admin verify rule delete params
func (o *MultiAdminVerifyRuleDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the multi admin verify rule delete params
func (o *MultiAdminVerifyRuleDeleteParams) WithContext(ctx context.Context) *MultiAdminVerifyRuleDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the multi admin verify rule delete params
func (o *MultiAdminVerifyRuleDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the multi admin verify rule delete params
func (o *MultiAdminVerifyRuleDeleteParams) WithHTTPClient(client *http.Client) *MultiAdminVerifyRuleDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the multi admin verify rule delete params
func (o *MultiAdminVerifyRuleDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithOperationPathParameter adds the operation to the multi admin verify rule delete params
func (o *MultiAdminVerifyRuleDeleteParams) WithOperationPathParameter(operation string) *MultiAdminVerifyRuleDeleteParams {
	o.SetOperationPathParameter(operation)
	return o
}

// SetOperationPathParameter adds the operation to the multi admin verify rule delete params
func (o *MultiAdminVerifyRuleDeleteParams) SetOperationPathParameter(operation string) {
	o.OperationPathParameter = operation
}

// WithOwnerUUIDPathParameter adds the ownerUUID to the multi admin verify rule delete params
func (o *MultiAdminVerifyRuleDeleteParams) WithOwnerUUIDPathParameter(ownerUUID string) *MultiAdminVerifyRuleDeleteParams {
	o.SetOwnerUUIDPathParameter(ownerUUID)
	return o
}

// SetOwnerUUIDPathParameter adds the ownerUuid to the multi admin verify rule delete params
func (o *MultiAdminVerifyRuleDeleteParams) SetOwnerUUIDPathParameter(ownerUUID string) {
	o.OwnerUUIDPathParameter = ownerUUID
}

// WriteToRequest writes these params to a swagger request
func (o *MultiAdminVerifyRuleDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param operation
	if err := r.SetPathParam("operation", o.OperationPathParameter); err != nil {
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
