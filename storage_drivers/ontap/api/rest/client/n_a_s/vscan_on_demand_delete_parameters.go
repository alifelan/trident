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

// NewVscanOnDemandDeleteParams creates a new VscanOnDemandDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVscanOnDemandDeleteParams() *VscanOnDemandDeleteParams {
	return &VscanOnDemandDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVscanOnDemandDeleteParamsWithTimeout creates a new VscanOnDemandDeleteParams object
// with the ability to set a timeout on a request.
func NewVscanOnDemandDeleteParamsWithTimeout(timeout time.Duration) *VscanOnDemandDeleteParams {
	return &VscanOnDemandDeleteParams{
		timeout: timeout,
	}
}

// NewVscanOnDemandDeleteParamsWithContext creates a new VscanOnDemandDeleteParams object
// with the ability to set a context for a request.
func NewVscanOnDemandDeleteParamsWithContext(ctx context.Context) *VscanOnDemandDeleteParams {
	return &VscanOnDemandDeleteParams{
		Context: ctx,
	}
}

// NewVscanOnDemandDeleteParamsWithHTTPClient creates a new VscanOnDemandDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewVscanOnDemandDeleteParamsWithHTTPClient(client *http.Client) *VscanOnDemandDeleteParams {
	return &VscanOnDemandDeleteParams{
		HTTPClient: client,
	}
}

/*
VscanOnDemandDeleteParams contains all the parameters to send to the API endpoint

	for the vscan on demand delete operation.

	Typically these are written to a http.Request.
*/
type VscanOnDemandDeleteParams struct {

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

// WithDefaults hydrates default values in the vscan on demand delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanOnDemandDeleteParams) WithDefaults() *VscanOnDemandDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vscan on demand delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanOnDemandDeleteParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the vscan on demand delete params
func (o *VscanOnDemandDeleteParams) WithTimeout(timeout time.Duration) *VscanOnDemandDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vscan on demand delete params
func (o *VscanOnDemandDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vscan on demand delete params
func (o *VscanOnDemandDeleteParams) WithContext(ctx context.Context) *VscanOnDemandDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vscan on demand delete params
func (o *VscanOnDemandDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vscan on demand delete params
func (o *VscanOnDemandDeleteParams) WithHTTPClient(client *http.Client) *VscanOnDemandDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vscan on demand delete params
func (o *VscanOnDemandDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithNamePathParameter adds the name to the vscan on demand delete params
func (o *VscanOnDemandDeleteParams) WithNamePathParameter(name string) *VscanOnDemandDeleteParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the vscan on demand delete params
func (o *VscanOnDemandDeleteParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithSVMUUIDPathParameter adds the svmUUID to the vscan on demand delete params
func (o *VscanOnDemandDeleteParams) WithSVMUUIDPathParameter(svmUUID string) *VscanOnDemandDeleteParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the vscan on demand delete params
func (o *VscanOnDemandDeleteParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *VscanOnDemandDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
