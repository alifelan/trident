// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

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

// NewCifsSymlinkMappingGetParams creates a new CifsSymlinkMappingGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsSymlinkMappingGetParams() *CifsSymlinkMappingGetParams {
	return &CifsSymlinkMappingGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsSymlinkMappingGetParamsWithTimeout creates a new CifsSymlinkMappingGetParams object
// with the ability to set a timeout on a request.
func NewCifsSymlinkMappingGetParamsWithTimeout(timeout time.Duration) *CifsSymlinkMappingGetParams {
	return &CifsSymlinkMappingGetParams{
		timeout: timeout,
	}
}

// NewCifsSymlinkMappingGetParamsWithContext creates a new CifsSymlinkMappingGetParams object
// with the ability to set a context for a request.
func NewCifsSymlinkMappingGetParamsWithContext(ctx context.Context) *CifsSymlinkMappingGetParams {
	return &CifsSymlinkMappingGetParams{
		Context: ctx,
	}
}

// NewCifsSymlinkMappingGetParamsWithHTTPClient creates a new CifsSymlinkMappingGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsSymlinkMappingGetParamsWithHTTPClient(client *http.Client) *CifsSymlinkMappingGetParams {
	return &CifsSymlinkMappingGetParams{
		HTTPClient: client,
	}
}

/*
CifsSymlinkMappingGetParams contains all the parameters to send to the API endpoint

	for the cifs symlink mapping get operation.

	Typically these are written to a http.Request.
*/
type CifsSymlinkMappingGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	/* UnixPath.

	   UNIX symbolic link path
	*/
	UnixPathPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs symlink mapping get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSymlinkMappingGetParams) WithDefaults() *CifsSymlinkMappingGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs symlink mapping get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsSymlinkMappingGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cifs symlink mapping get params
func (o *CifsSymlinkMappingGetParams) WithTimeout(timeout time.Duration) *CifsSymlinkMappingGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs symlink mapping get params
func (o *CifsSymlinkMappingGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs symlink mapping get params
func (o *CifsSymlinkMappingGetParams) WithContext(ctx context.Context) *CifsSymlinkMappingGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs symlink mapping get params
func (o *CifsSymlinkMappingGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs symlink mapping get params
func (o *CifsSymlinkMappingGetParams) WithHTTPClient(client *http.Client) *CifsSymlinkMappingGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs symlink mapping get params
func (o *CifsSymlinkMappingGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the cifs symlink mapping get params
func (o *CifsSymlinkMappingGetParams) WithFieldsQueryParameter(fields []string) *CifsSymlinkMappingGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the cifs symlink mapping get params
func (o *CifsSymlinkMappingGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithSVMUUIDPathParameter adds the svmUUID to the cifs symlink mapping get params
func (o *CifsSymlinkMappingGetParams) WithSVMUUIDPathParameter(svmUUID string) *CifsSymlinkMappingGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the cifs symlink mapping get params
func (o *CifsSymlinkMappingGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WithUnixPathPathParameter adds the unixPath to the cifs symlink mapping get params
func (o *CifsSymlinkMappingGetParams) WithUnixPathPathParameter(unixPath string) *CifsSymlinkMappingGetParams {
	o.SetUnixPathPathParameter(unixPath)
	return o
}

// SetUnixPathPathParameter adds the unixPath to the cifs symlink mapping get params
func (o *CifsSymlinkMappingGetParams) SetUnixPathPathParameter(unixPath string) {
	o.UnixPathPathParameter = unixPath
}

// WriteToRequest writes these params to a swagger request
func (o *CifsSymlinkMappingGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param unix_path
	if err := r.SetPathParam("unix_path", o.UnixPathPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCifsSymlinkMappingGet binds the parameter fields
func (o *CifsSymlinkMappingGetParams) bindParamFields(formats strfmt.Registry) []string {
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
