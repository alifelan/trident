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

// NewIpsecGetParams creates a new IpsecGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIpsecGetParams() *IpsecGetParams {
	return &IpsecGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIpsecGetParamsWithTimeout creates a new IpsecGetParams object
// with the ability to set a timeout on a request.
func NewIpsecGetParamsWithTimeout(timeout time.Duration) *IpsecGetParams {
	return &IpsecGetParams{
		timeout: timeout,
	}
}

// NewIpsecGetParamsWithContext creates a new IpsecGetParams object
// with the ability to set a context for a request.
func NewIpsecGetParamsWithContext(ctx context.Context) *IpsecGetParams {
	return &IpsecGetParams{
		Context: ctx,
	}
}

// NewIpsecGetParamsWithHTTPClient creates a new IpsecGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewIpsecGetParamsWithHTTPClient(client *http.Client) *IpsecGetParams {
	return &IpsecGetParams{
		HTTPClient: client,
	}
}

/*
IpsecGetParams contains all the parameters to send to the API endpoint

	for the ipsec get operation.

	Typically these are written to a http.Request.
*/
type IpsecGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ipsec get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpsecGetParams) WithDefaults() *IpsecGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ipsec get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IpsecGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ipsec get params
func (o *IpsecGetParams) WithTimeout(timeout time.Duration) *IpsecGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipsec get params
func (o *IpsecGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipsec get params
func (o *IpsecGetParams) WithContext(ctx context.Context) *IpsecGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipsec get params
func (o *IpsecGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipsec get params
func (o *IpsecGetParams) WithHTTPClient(client *http.Client) *IpsecGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipsec get params
func (o *IpsecGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the ipsec get params
func (o *IpsecGetParams) WithFieldsQueryParameter(fields []string) *IpsecGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the ipsec get params
func (o *IpsecGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WriteToRequest writes these params to a swagger request
func (o *IpsecGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamIpsecGet binds the parameter fields
func (o *IpsecGetParams) bindParamFields(formats strfmt.Registry) []string {
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
