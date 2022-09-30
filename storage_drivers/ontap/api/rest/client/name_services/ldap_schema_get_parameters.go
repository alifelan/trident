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

// NewLdapSchemaGetParams creates a new LdapSchemaGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLdapSchemaGetParams() *LdapSchemaGetParams {
	return &LdapSchemaGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLdapSchemaGetParamsWithTimeout creates a new LdapSchemaGetParams object
// with the ability to set a timeout on a request.
func NewLdapSchemaGetParamsWithTimeout(timeout time.Duration) *LdapSchemaGetParams {
	return &LdapSchemaGetParams{
		timeout: timeout,
	}
}

// NewLdapSchemaGetParamsWithContext creates a new LdapSchemaGetParams object
// with the ability to set a context for a request.
func NewLdapSchemaGetParamsWithContext(ctx context.Context) *LdapSchemaGetParams {
	return &LdapSchemaGetParams{
		Context: ctx,
	}
}

// NewLdapSchemaGetParamsWithHTTPClient creates a new LdapSchemaGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewLdapSchemaGetParamsWithHTTPClient(client *http.Client) *LdapSchemaGetParams {
	return &LdapSchemaGetParams{
		HTTPClient: client,
	}
}

/*
LdapSchemaGetParams contains all the parameters to send to the API endpoint

	for the ldap schema get operation.

	Typically these are written to a http.Request.
*/
type LdapSchemaGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Name.

	   LDAP schema name.
	*/
	NamePathParameter string

	/* OwnerUUID.

	   UUID of the owner to which this object belongs.
	*/
	OwnerUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ldap schema get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LdapSchemaGetParams) WithDefaults() *LdapSchemaGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ldap schema get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LdapSchemaGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the ldap schema get params
func (o *LdapSchemaGetParams) WithTimeout(timeout time.Duration) *LdapSchemaGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ldap schema get params
func (o *LdapSchemaGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ldap schema get params
func (o *LdapSchemaGetParams) WithContext(ctx context.Context) *LdapSchemaGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ldap schema get params
func (o *LdapSchemaGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ldap schema get params
func (o *LdapSchemaGetParams) WithHTTPClient(client *http.Client) *LdapSchemaGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ldap schema get params
func (o *LdapSchemaGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the ldap schema get params
func (o *LdapSchemaGetParams) WithFieldsQueryParameter(fields []string) *LdapSchemaGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the ldap schema get params
func (o *LdapSchemaGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithNamePathParameter adds the name to the ldap schema get params
func (o *LdapSchemaGetParams) WithNamePathParameter(name string) *LdapSchemaGetParams {
	o.SetNamePathParameter(name)
	return o
}

// SetNamePathParameter adds the name to the ldap schema get params
func (o *LdapSchemaGetParams) SetNamePathParameter(name string) {
	o.NamePathParameter = name
}

// WithOwnerUUIDPathParameter adds the ownerUUID to the ldap schema get params
func (o *LdapSchemaGetParams) WithOwnerUUIDPathParameter(ownerUUID string) *LdapSchemaGetParams {
	o.SetOwnerUUIDPathParameter(ownerUUID)
	return o
}

// SetOwnerUUIDPathParameter adds the ownerUuid to the ldap schema get params
func (o *LdapSchemaGetParams) SetOwnerUUIDPathParameter(ownerUUID string) {
	o.OwnerUUIDPathParameter = ownerUUID
}

// WriteToRequest writes these params to a swagger request
func (o *LdapSchemaGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param name
	if err := r.SetPathParam("name", o.NamePathParameter); err != nil {
		return err
	}

	// path param owner.uuid
	if err := r.SetPathParam("owner.uuid", o.OwnerUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamLdapSchemaGet binds the parameter fields
func (o *LdapSchemaGetParams) bindParamFields(formats strfmt.Registry) []string {
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
