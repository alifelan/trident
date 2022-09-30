// Code generated by go-swagger; DO NOT EDIT.

package snaplock

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

// NewSnaplockLegalHoldGetParams creates a new SnaplockLegalHoldGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnaplockLegalHoldGetParams() *SnaplockLegalHoldGetParams {
	return &SnaplockLegalHoldGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnaplockLegalHoldGetParamsWithTimeout creates a new SnaplockLegalHoldGetParams object
// with the ability to set a timeout on a request.
func NewSnaplockLegalHoldGetParamsWithTimeout(timeout time.Duration) *SnaplockLegalHoldGetParams {
	return &SnaplockLegalHoldGetParams{
		timeout: timeout,
	}
}

// NewSnaplockLegalHoldGetParamsWithContext creates a new SnaplockLegalHoldGetParams object
// with the ability to set a context for a request.
func NewSnaplockLegalHoldGetParamsWithContext(ctx context.Context) *SnaplockLegalHoldGetParams {
	return &SnaplockLegalHoldGetParams{
		Context: ctx,
	}
}

// NewSnaplockLegalHoldGetParamsWithHTTPClient creates a new SnaplockLegalHoldGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnaplockLegalHoldGetParamsWithHTTPClient(client *http.Client) *SnaplockLegalHoldGetParams {
	return &SnaplockLegalHoldGetParams{
		HTTPClient: client,
	}
}

/*
SnaplockLegalHoldGetParams contains all the parameters to send to the API endpoint

	for the snaplock legal hold get operation.

	Typically these are written to a http.Request.
*/
type SnaplockLegalHoldGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* ID.

	   Litigation ID
	*/
	IDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snaplock legal hold get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldGetParams) WithDefaults() *SnaplockLegalHoldGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snaplock legal hold get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnaplockLegalHoldGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the snaplock legal hold get params
func (o *SnaplockLegalHoldGetParams) WithTimeout(timeout time.Duration) *SnaplockLegalHoldGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snaplock legal hold get params
func (o *SnaplockLegalHoldGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snaplock legal hold get params
func (o *SnaplockLegalHoldGetParams) WithContext(ctx context.Context) *SnaplockLegalHoldGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snaplock legal hold get params
func (o *SnaplockLegalHoldGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snaplock legal hold get params
func (o *SnaplockLegalHoldGetParams) WithHTTPClient(client *http.Client) *SnaplockLegalHoldGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snaplock legal hold get params
func (o *SnaplockLegalHoldGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the snaplock legal hold get params
func (o *SnaplockLegalHoldGetParams) WithFieldsQueryParameter(fields []string) *SnaplockLegalHoldGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the snaplock legal hold get params
func (o *SnaplockLegalHoldGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIDPathParameter adds the id to the snaplock legal hold get params
func (o *SnaplockLegalHoldGetParams) WithIDPathParameter(id string) *SnaplockLegalHoldGetParams {
	o.SetIDPathParameter(id)
	return o
}

// SetIDPathParameter adds the id to the snaplock legal hold get params
func (o *SnaplockLegalHoldGetParams) SetIDPathParameter(id string) {
	o.IDPathParameter = id
}

// WriteToRequest writes these params to a swagger request
func (o *SnaplockLegalHoldGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.IDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSnaplockLegalHoldGet binds the parameter fields
func (o *SnaplockLegalHoldGetParams) bindParamFields(formats strfmt.Registry) []string {
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
