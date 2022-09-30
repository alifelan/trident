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

// NewNodeGetParams creates a new NodeGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNodeGetParams() *NodeGetParams {
	return &NodeGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNodeGetParamsWithTimeout creates a new NodeGetParams object
// with the ability to set a timeout on a request.
func NewNodeGetParamsWithTimeout(timeout time.Duration) *NodeGetParams {
	return &NodeGetParams{
		timeout: timeout,
	}
}

// NewNodeGetParamsWithContext creates a new NodeGetParams object
// with the ability to set a context for a request.
func NewNodeGetParamsWithContext(ctx context.Context) *NodeGetParams {
	return &NodeGetParams{
		Context: ctx,
	}
}

// NewNodeGetParamsWithHTTPClient creates a new NodeGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNodeGetParamsWithHTTPClient(client *http.Client) *NodeGetParams {
	return &NodeGetParams{
		HTTPClient: client,
	}
}

/*
NodeGetParams contains all the parameters to send to the API endpoint

	for the node get operation.

	Typically these are written to a http.Request.
*/
type NodeGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	// UUID.
	//
	// Format: uuid
	UUIDPathParameter strfmt.UUID

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the node get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodeGetParams) WithDefaults() *NodeGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the node get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NodeGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the node get params
func (o *NodeGetParams) WithTimeout(timeout time.Duration) *NodeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the node get params
func (o *NodeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the node get params
func (o *NodeGetParams) WithContext(ctx context.Context) *NodeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the node get params
func (o *NodeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the node get params
func (o *NodeGetParams) WithHTTPClient(client *http.Client) *NodeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the node get params
func (o *NodeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the node get params
func (o *NodeGetParams) WithFieldsQueryParameter(fields []string) *NodeGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the node get params
func (o *NodeGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithUUIDPathParameter adds the uuid to the node get params
func (o *NodeGetParams) WithUUIDPathParameter(uuid strfmt.UUID) *NodeGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the node get params
func (o *NodeGetParams) SetUUIDPathParameter(uuid strfmt.UUID) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *NodeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("uuid", o.UUIDPathParameter.String()); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNodeGet binds the parameter fields
func (o *NodeGetParams) bindParamFields(formats strfmt.Registry) []string {
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
