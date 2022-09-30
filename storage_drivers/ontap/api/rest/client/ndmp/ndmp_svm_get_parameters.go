// Code generated by go-swagger; DO NOT EDIT.

package ndmp

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

// NewNdmpSvmGetParams creates a new NdmpSvmGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNdmpSvmGetParams() *NdmpSvmGetParams {
	return &NdmpSvmGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNdmpSvmGetParamsWithTimeout creates a new NdmpSvmGetParams object
// with the ability to set a timeout on a request.
func NewNdmpSvmGetParamsWithTimeout(timeout time.Duration) *NdmpSvmGetParams {
	return &NdmpSvmGetParams{
		timeout: timeout,
	}
}

// NewNdmpSvmGetParamsWithContext creates a new NdmpSvmGetParams object
// with the ability to set a context for a request.
func NewNdmpSvmGetParamsWithContext(ctx context.Context) *NdmpSvmGetParams {
	return &NdmpSvmGetParams{
		Context: ctx,
	}
}

// NewNdmpSvmGetParamsWithHTTPClient creates a new NdmpSvmGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNdmpSvmGetParamsWithHTTPClient(client *http.Client) *NdmpSvmGetParams {
	return &NdmpSvmGetParams{
		HTTPClient: client,
	}
}

/*
NdmpSvmGetParams contains all the parameters to send to the API endpoint

	for the ndmp svm get operation.

	Typically these are written to a http.Request.
*/
type NdmpSvmGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* SvmUUID.

	   SVM UUID
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ndmp svm get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpSvmGetParams) WithDefaults() *NdmpSvmGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ndmp svm get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpSvmGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ndmp svm get params
func (o *NdmpSvmGetParams) WithTimeout(timeout time.Duration) *NdmpSvmGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ndmp svm get params
func (o *NdmpSvmGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ndmp svm get params
func (o *NdmpSvmGetParams) WithContext(ctx context.Context) *NdmpSvmGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ndmp svm get params
func (o *NdmpSvmGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ndmp svm get params
func (o *NdmpSvmGetParams) WithHTTPClient(client *http.Client) *NdmpSvmGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ndmp svm get params
func (o *NdmpSvmGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the ndmp svm get params
func (o *NdmpSvmGetParams) WithFieldsQueryParameter(fields []string) *NdmpSvmGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the ndmp svm get params
func (o *NdmpSvmGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithSVMUUIDPathParameter adds the svmUUID to the ndmp svm get params
func (o *NdmpSvmGetParams) WithSVMUUIDPathParameter(svmUUID string) *NdmpSvmGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the ndmp svm get params
func (o *NdmpSvmGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NdmpSvmGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
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

// bindParamNdmpSvmGet binds the parameter fields
func (o *NdmpSvmGetParams) bindParamFields(formats strfmt.Registry) []string {
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
