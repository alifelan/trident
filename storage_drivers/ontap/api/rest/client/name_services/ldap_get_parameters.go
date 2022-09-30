// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

// NewLdapGetParams creates a new LdapGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLdapGetParams() *LdapGetParams {
	return &LdapGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLdapGetParamsWithTimeout creates a new LdapGetParams object
// with the ability to set a timeout on a request.
func NewLdapGetParamsWithTimeout(timeout time.Duration) *LdapGetParams {
	return &LdapGetParams{
		timeout: timeout,
	}
}

// NewLdapGetParamsWithContext creates a new LdapGetParams object
// with the ability to set a context for a request.
func NewLdapGetParamsWithContext(ctx context.Context) *LdapGetParams {
	return &LdapGetParams{
		Context: ctx,
	}
}

// NewLdapGetParamsWithHTTPClient creates a new LdapGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewLdapGetParamsWithHTTPClient(client *http.Client) *LdapGetParams {
	return &LdapGetParams{
		HTTPClient: client,
	}
}

/*
LdapGetParams contains all the parameters to send to the API endpoint

	for the ldap get operation.

	Typically these are written to a http.Request.
*/
type LdapGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ldap get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LdapGetParams) WithDefaults() *LdapGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ldap get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LdapGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ldap get params
func (o *LdapGetParams) WithTimeout(timeout time.Duration) *LdapGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ldap get params
func (o *LdapGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ldap get params
func (o *LdapGetParams) WithContext(ctx context.Context) *LdapGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ldap get params
func (o *LdapGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ldap get params
func (o *LdapGetParams) WithHTTPClient(client *http.Client) *LdapGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ldap get params
func (o *LdapGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the ldap get params
func (o *LdapGetParams) WithFieldsQueryParameter(fields []string) *LdapGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the ldap get params
func (o *LdapGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithSVMUUIDPathParameter adds the svmUUID to the ldap get params
func (o *LdapGetParams) WithSVMUUIDPathParameter(svmUUID string) *LdapGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the ldap get params
func (o *LdapGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *LdapGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamLdapGet binds the parameter fields
func (o *LdapGetParams) bindParamFields(formats strfmt.Registry) []string {
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
