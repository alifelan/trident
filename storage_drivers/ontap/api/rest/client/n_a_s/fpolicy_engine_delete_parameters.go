// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewFpolicyEngineDeleteParams creates a new FpolicyEngineDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFpolicyEngineDeleteParams() *FpolicyEngineDeleteParams {
	return &FpolicyEngineDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFpolicyEngineDeleteParamsWithTimeout creates a new FpolicyEngineDeleteParams object
// with the ability to set a timeout on a request.
func NewFpolicyEngineDeleteParamsWithTimeout(timeout time.Duration) *FpolicyEngineDeleteParams {
	return &FpolicyEngineDeleteParams{
		timeout: timeout,
	}
}

// NewFpolicyEngineDeleteParamsWithContext creates a new FpolicyEngineDeleteParams object
// with the ability to set a context for a request.
func NewFpolicyEngineDeleteParamsWithContext(ctx context.Context) *FpolicyEngineDeleteParams {
	return &FpolicyEngineDeleteParams{
		Context: ctx,
	}
}

// NewFpolicyEngineDeleteParamsWithHTTPClient creates a new FpolicyEngineDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewFpolicyEngineDeleteParamsWithHTTPClient(client *http.Client) *FpolicyEngineDeleteParams {
	return &FpolicyEngineDeleteParams{
		HTTPClient: client,
	}
}

/*
FpolicyEngineDeleteParams contains all the parameters to send to the API endpoint

	for the fpolicy engine delete operation.

	Typically these are written to a http.Request.
*/
type FpolicyEngineDeleteParams struct {

	// Name.
	NamePathParameter string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fpolicy engine delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyEngineDeleteParams) WithDefaults() *FpolicyEngineDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fpolicy engine delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyEngineDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fpolicy engine delete params
func (o *FpolicyEngineDeleteParams) WithTimeout(timeout time.Duration) *FpolicyEngineDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fpolicy engine delete params
func (o *FpolicyEngineDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fpolicy engine delete params
func (o *FpolicyEngineDeleteParams) WithContext(ctx context.Context) *FpolicyEngineDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fpolicy engine delete params
func (o *FpolicyEngineDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fpolicy engine delete params
func (o *FpolicyEngineDeleteParams) WithHTTPClient(client *http.Client) *FpolicyEngineDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fpolicy engine delete params
func (o *FpolicyEngineDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamePathParameter adds the name to the fpolicy engine delete params
func (o *FpolicyEngineDeleteParams) WithNamePathParameter(name string) *FpolicyEngineDeleteParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the fpolicy engine delete params
func (o *FpolicyEngineDeleteParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithSVMUUIDPathParameter adds the svmUUID to the fpolicy engine delete params
func (o *FpolicyEngineDeleteParams) WithSVMUUIDPathParameter(svmUUID string) *FpolicyEngineDeleteParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the fpolicy engine delete params
func (o *FpolicyEngineDeleteParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *FpolicyEngineDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
