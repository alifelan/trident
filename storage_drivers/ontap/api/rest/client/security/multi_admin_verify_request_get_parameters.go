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

// NewMultiAdminVerifyRequestGetParams creates a new MultiAdminVerifyRequestGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMultiAdminVerifyRequestGetParams() *MultiAdminVerifyRequestGetParams {
	return &MultiAdminVerifyRequestGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMultiAdminVerifyRequestGetParamsWithTimeout creates a new MultiAdminVerifyRequestGetParams object
// with the ability to set a timeout on a request.
func NewMultiAdminVerifyRequestGetParamsWithTimeout(timeout time.Duration) *MultiAdminVerifyRequestGetParams {
	return &MultiAdminVerifyRequestGetParams{
		timeout: timeout,
	}
}

// NewMultiAdminVerifyRequestGetParamsWithContext creates a new MultiAdminVerifyRequestGetParams object
// with the ability to set a context for a request.
func NewMultiAdminVerifyRequestGetParamsWithContext(ctx context.Context) *MultiAdminVerifyRequestGetParams {
	return &MultiAdminVerifyRequestGetParams{
		Context: ctx,
	}
}

// NewMultiAdminVerifyRequestGetParamsWithHTTPClient creates a new MultiAdminVerifyRequestGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewMultiAdminVerifyRequestGetParamsWithHTTPClient(client *http.Client) *MultiAdminVerifyRequestGetParams {
	return &MultiAdminVerifyRequestGetParams{
		HTTPClient: client,
	}
}

/*
MultiAdminVerifyRequestGetParams contains all the parameters to send to the API endpoint

	for the multi admin verify request get operation.

	Typically these are written to a http.Request.
*/
type MultiAdminVerifyRequestGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	// Index.
	IndexPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the multi admin verify request get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyRequestGetParams) WithDefaults() *MultiAdminVerifyRequestGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the multi admin verify request get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MultiAdminVerifyRequestGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the multi admin verify request get params
func (o *MultiAdminVerifyRequestGetParams) WithTimeout(timeout time.Duration) *MultiAdminVerifyRequestGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the multi admin verify request get params
func (o *MultiAdminVerifyRequestGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the multi admin verify request get params
func (o *MultiAdminVerifyRequestGetParams) WithContext(ctx context.Context) *MultiAdminVerifyRequestGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the multi admin verify request get params
func (o *MultiAdminVerifyRequestGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the multi admin verify request get params
func (o *MultiAdminVerifyRequestGetParams) WithHTTPClient(client *http.Client) *MultiAdminVerifyRequestGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the multi admin verify request get params
func (o *MultiAdminVerifyRequestGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the multi admin verify request get params
func (o *MultiAdminVerifyRequestGetParams) WithFieldsQueryParameter(fields []string) *MultiAdminVerifyRequestGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the multi admin verify request get params
func (o *MultiAdminVerifyRequestGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIndexPathParameter adds the index to the multi admin verify request get params
func (o *MultiAdminVerifyRequestGetParams) WithIndexPathParameter(index string) *MultiAdminVerifyRequestGetParams {
	o.SetIndexPathParameter(index)
	return o
}

// SetIndexPathParameter adds the index to the multi admin verify request get params
func (o *MultiAdminVerifyRequestGetParams) SetIndexPathParameter(index string) {
	o.IndexPathParameter = index
}

// WriteToRequest writes these params to a swagger request
func (o *MultiAdminVerifyRequestGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param index
	if err := r.SetPathParam("index", o.IndexPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamMultiAdminVerifyRequestGet binds the parameter fields
func (o *MultiAdminVerifyRequestGetParams) bindParamFields(formats strfmt.Registry) []string {
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
