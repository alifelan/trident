// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewJobGetParams creates a new JobGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewJobGetParams() *JobGetParams {
	return &JobGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewJobGetParamsWithTimeout creates a new JobGetParams object
// with the ability to set a timeout on a request.
func NewJobGetParamsWithTimeout(timeout time.Duration) *JobGetParams {
	return &JobGetParams{
		timeout: timeout,
	}
}

// NewJobGetParamsWithContext creates a new JobGetParams object
// with the ability to set a context for a request.
func NewJobGetParamsWithContext(ctx context.Context) *JobGetParams {
	return &JobGetParams{
		Context: ctx,
	}
}

// NewJobGetParamsWithHTTPClient creates a new JobGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewJobGetParamsWithHTTPClient(client *http.Client) *JobGetParams {
	return &JobGetParams{
		HTTPClient: client,
	}
}

/*
JobGetParams contains all the parameters to send to the API endpoint

	for the job get operation.

	Typically these are written to a http.Request.
*/
type JobGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* UUID.

	   Job UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the job get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobGetParams) WithDefaults() *JobGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the job get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the job get params
func (o *JobGetParams) WithTimeout(timeout time.Duration) *JobGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job get params
func (o *JobGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job get params
func (o *JobGetParams) WithContext(ctx context.Context) *JobGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job get params
func (o *JobGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job get params
func (o *JobGetParams) WithHTTPClient(client *http.Client) *JobGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job get params
func (o *JobGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the job get params
func (o *JobGetParams) WithFieldsQueryParameter(fields []string) *JobGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the job get params
func (o *JobGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithUUIDPathParameter adds the uuid to the job get params
func (o *JobGetParams) WithUUIDPathParameter(uuid string) *JobGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the job get params
func (o *JobGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *JobGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamJobGet binds the parameter fields
func (o *JobGetParams) bindParamFields(formats strfmt.Registry) []string {
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
