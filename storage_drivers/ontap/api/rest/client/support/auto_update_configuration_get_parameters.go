// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewAutoUpdateConfigurationGetParams creates a new AutoUpdateConfigurationGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAutoUpdateConfigurationGetParams() *AutoUpdateConfigurationGetParams {
	return &AutoUpdateConfigurationGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAutoUpdateConfigurationGetParamsWithTimeout creates a new AutoUpdateConfigurationGetParams object
// with the ability to set a timeout on a request.
func NewAutoUpdateConfigurationGetParamsWithTimeout(timeout time.Duration) *AutoUpdateConfigurationGetParams {
	return &AutoUpdateConfigurationGetParams{
		timeout: timeout,
	}
}

// NewAutoUpdateConfigurationGetParamsWithContext creates a new AutoUpdateConfigurationGetParams object
// with the ability to set a context for a request.
func NewAutoUpdateConfigurationGetParamsWithContext(ctx context.Context) *AutoUpdateConfigurationGetParams {
	return &AutoUpdateConfigurationGetParams{
		Context: ctx,
	}
}

// NewAutoUpdateConfigurationGetParamsWithHTTPClient creates a new AutoUpdateConfigurationGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAutoUpdateConfigurationGetParamsWithHTTPClient(client *http.Client) *AutoUpdateConfigurationGetParams {
	return &AutoUpdateConfigurationGetParams{
		HTTPClient: client,
	}
}

/*
AutoUpdateConfigurationGetParams contains all the parameters to send to the API endpoint

	for the auto update configuration get operation.

	Typically these are written to a http.Request.
*/
type AutoUpdateConfigurationGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* UUID.

	   Unique identifier for configuration record.
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the auto update configuration get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AutoUpdateConfigurationGetParams) WithDefaults() *AutoUpdateConfigurationGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the auto update configuration get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AutoUpdateConfigurationGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the auto update configuration get params
func (o *AutoUpdateConfigurationGetParams) WithTimeout(timeout time.Duration) *AutoUpdateConfigurationGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the auto update configuration get params
func (o *AutoUpdateConfigurationGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the auto update configuration get params
func (o *AutoUpdateConfigurationGetParams) WithContext(ctx context.Context) *AutoUpdateConfigurationGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the auto update configuration get params
func (o *AutoUpdateConfigurationGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the auto update configuration get params
func (o *AutoUpdateConfigurationGetParams) WithHTTPClient(client *http.Client) *AutoUpdateConfigurationGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the auto update configuration get params
func (o *AutoUpdateConfigurationGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the auto update configuration get params
func (o *AutoUpdateConfigurationGetParams) WithFieldsQueryParameter(fields []string) *AutoUpdateConfigurationGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the auto update configuration get params
func (o *AutoUpdateConfigurationGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithUUIDPathParameter adds the uuid to the auto update configuration get params
func (o *AutoUpdateConfigurationGetParams) WithUUIDPathParameter(uuid string) *AutoUpdateConfigurationGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the auto update configuration get params
func (o *AutoUpdateConfigurationGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *AutoUpdateConfigurationGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamAutoUpdateConfigurationGet binds the parameter fields
func (o *AutoUpdateConfigurationGetParams) bindParamFields(formats strfmt.Registry) []string {
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
