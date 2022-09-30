// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewIgroupGetParams creates a new IgroupGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewIgroupGetParams() *IgroupGetParams {
	return &IgroupGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewIgroupGetParamsWithTimeout creates a new IgroupGetParams object
// with the ability to set a timeout on a request.
func NewIgroupGetParamsWithTimeout(timeout time.Duration) *IgroupGetParams {
	return &IgroupGetParams{
		timeout: timeout,
	}
}

// NewIgroupGetParamsWithContext creates a new IgroupGetParams object
// with the ability to set a context for a request.
func NewIgroupGetParamsWithContext(ctx context.Context) *IgroupGetParams {
	return &IgroupGetParams{
		Context: ctx,
	}
}

// NewIgroupGetParamsWithHTTPClient creates a new IgroupGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewIgroupGetParamsWithHTTPClient(client *http.Client) *IgroupGetParams {
	return &IgroupGetParams{
		HTTPClient: client,
	}
}

/*
IgroupGetParams contains all the parameters to send to the API endpoint

	for the igroup get operation.

	Typically these are written to a http.Request.
*/
type IgroupGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* UUID.

	   The unique identifier of the initiator group.

	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the igroup get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IgroupGetParams) WithDefaults() *IgroupGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the igroup get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *IgroupGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the igroup get params
func (o *IgroupGetParams) WithTimeout(timeout time.Duration) *IgroupGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the igroup get params
func (o *IgroupGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the igroup get params
func (o *IgroupGetParams) WithContext(ctx context.Context) *IgroupGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the igroup get params
func (o *IgroupGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the igroup get params
func (o *IgroupGetParams) WithHTTPClient(client *http.Client) *IgroupGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the igroup get params
func (o *IgroupGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the igroup get params
func (o *IgroupGetParams) WithFieldsQueryParameter(fields []string) *IgroupGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the igroup get params
func (o *IgroupGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithUUIDPathParameter adds the uuid to the igroup get params
func (o *IgroupGetParams) WithUUIDPathParameter(uuid string) *IgroupGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the igroup get params
func (o *IgroupGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *IgroupGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamIgroupGet binds the parameter fields
func (o *IgroupGetParams) bindParamFields(formats strfmt.Registry) []string {
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
