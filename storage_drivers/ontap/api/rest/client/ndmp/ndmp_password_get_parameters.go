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

// NewNdmpPasswordGetParams creates a new NdmpPasswordGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNdmpPasswordGetParams() *NdmpPasswordGetParams {
	return &NdmpPasswordGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNdmpPasswordGetParamsWithTimeout creates a new NdmpPasswordGetParams object
// with the ability to set a timeout on a request.
func NewNdmpPasswordGetParamsWithTimeout(timeout time.Duration) *NdmpPasswordGetParams {
	return &NdmpPasswordGetParams{
		timeout: timeout,
	}
}

// NewNdmpPasswordGetParamsWithContext creates a new NdmpPasswordGetParams object
// with the ability to set a context for a request.
func NewNdmpPasswordGetParamsWithContext(ctx context.Context) *NdmpPasswordGetParams {
	return &NdmpPasswordGetParams{
		Context: ctx,
	}
}

// NewNdmpPasswordGetParamsWithHTTPClient creates a new NdmpPasswordGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNdmpPasswordGetParamsWithHTTPClient(client *http.Client) *NdmpPasswordGetParams {
	return &NdmpPasswordGetParams{
		HTTPClient: client,
	}
}

/*
NdmpPasswordGetParams contains all the parameters to send to the API endpoint

	for the ndmp password get operation.

	Typically these are written to a http.Request.
*/
type NdmpPasswordGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* SvmUUID.

	   SVM UUID
	*/
	SVMUUIDPathParameter string

	/* User.

	   User name
	*/
	UserPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ndmp password get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpPasswordGetParams) WithDefaults() *NdmpPasswordGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ndmp password get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NdmpPasswordGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ndmp password get params
func (o *NdmpPasswordGetParams) WithTimeout(timeout time.Duration) *NdmpPasswordGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ndmp password get params
func (o *NdmpPasswordGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ndmp password get params
func (o *NdmpPasswordGetParams) WithContext(ctx context.Context) *NdmpPasswordGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ndmp password get params
func (o *NdmpPasswordGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ndmp password get params
func (o *NdmpPasswordGetParams) WithHTTPClient(client *http.Client) *NdmpPasswordGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ndmp password get params
func (o *NdmpPasswordGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the ndmp password get params
func (o *NdmpPasswordGetParams) WithFieldsQueryParameter(fields []string) *NdmpPasswordGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the ndmp password get params
func (o *NdmpPasswordGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithSVMUUIDPathParameter adds the svmUUID to the ndmp password get params
func (o *NdmpPasswordGetParams) WithSVMUUIDPathParameter(svmUUID string) *NdmpPasswordGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the ndmp password get params
func (o *NdmpPasswordGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WithUserPathParameter adds the user to the ndmp password get params
func (o *NdmpPasswordGetParams) WithUserPathParameter(user string) *NdmpPasswordGetParams {
	o.SetUserPathParameter(user)
	return o
}

// SetUserPathParameter adds the user to the ndmp password get params
func (o *NdmpPasswordGetParams) SetUserPathParameter(user string) {
	o.UserPathParameter = user
}

// WriteToRequest writes these params to a swagger request
func (o *NdmpPasswordGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param user
	if err := r.SetPathParam("user", o.UserPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNdmpPasswordGet binds the parameter fields
func (o *NdmpPasswordGetParams) bindParamFields(formats strfmt.Registry) []string {
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
