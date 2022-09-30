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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewFpolicyEventModifyParams creates a new FpolicyEventModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFpolicyEventModifyParams() *FpolicyEventModifyParams {
	return &FpolicyEventModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFpolicyEventModifyParamsWithTimeout creates a new FpolicyEventModifyParams object
// with the ability to set a timeout on a request.
func NewFpolicyEventModifyParamsWithTimeout(timeout time.Duration) *FpolicyEventModifyParams {
	return &FpolicyEventModifyParams{
		timeout: timeout,
	}
}

// NewFpolicyEventModifyParamsWithContext creates a new FpolicyEventModifyParams object
// with the ability to set a context for a request.
func NewFpolicyEventModifyParamsWithContext(ctx context.Context) *FpolicyEventModifyParams {
	return &FpolicyEventModifyParams{
		Context: ctx,
	}
}

// NewFpolicyEventModifyParamsWithHTTPClient creates a new FpolicyEventModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewFpolicyEventModifyParamsWithHTTPClient(client *http.Client) *FpolicyEventModifyParams {
	return &FpolicyEventModifyParams{
		HTTPClient: client,
	}
}

/*
FpolicyEventModifyParams contains all the parameters to send to the API endpoint

	for the fpolicy event modify operation.

	Typically these are written to a http.Request.
*/
type FpolicyEventModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.FpolicyEvent

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

// WithDefaults hydrates default values in the fpolicy event modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyEventModifyParams) WithDefaults() *FpolicyEventModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fpolicy event modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FpolicyEventModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the fpolicy event modify params
func (o *FpolicyEventModifyParams) WithTimeout(timeout time.Duration) *FpolicyEventModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fpolicy event modify params
func (o *FpolicyEventModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fpolicy event modify params
func (o *FpolicyEventModifyParams) WithContext(ctx context.Context) *FpolicyEventModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fpolicy event modify params
func (o *FpolicyEventModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fpolicy event modify params
func (o *FpolicyEventModifyParams) WithHTTPClient(client *http.Client) *FpolicyEventModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fpolicy event modify params
func (o *FpolicyEventModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the fpolicy event modify params
func (o *FpolicyEventModifyParams) WithInfo(info *models.FpolicyEvent) *FpolicyEventModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the fpolicy event modify params
func (o *FpolicyEventModifyParams) SetInfo(info *models.FpolicyEvent) {
	o.Info = info
}

// WithNamePathParameter adds the name to the fpolicy event modify params
func (o *FpolicyEventModifyParams) WithNamePathParameter(name string) *FpolicyEventModifyParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the fpolicy event modify params
func (o *FpolicyEventModifyParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithSVMUUIDPathParameter adds the svmUUID to the fpolicy event modify params
func (o *FpolicyEventModifyParams) WithSVMUUIDPathParameter(svmUUID string) *FpolicyEventModifyParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the fpolicy event modify params
func (o *FpolicyEventModifyParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *FpolicyEventModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

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
