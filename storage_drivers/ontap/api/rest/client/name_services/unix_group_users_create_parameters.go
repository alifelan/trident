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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewUnixGroupUsersCreateParams creates a new UnixGroupUsersCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnixGroupUsersCreateParams() *UnixGroupUsersCreateParams {
	return &UnixGroupUsersCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnixGroupUsersCreateParamsWithTimeout creates a new UnixGroupUsersCreateParams object
// with the ability to set a timeout on a request.
func NewUnixGroupUsersCreateParamsWithTimeout(timeout time.Duration) *UnixGroupUsersCreateParams {
	return &UnixGroupUsersCreateParams{
		timeout: timeout,
	}
}

// NewUnixGroupUsersCreateParamsWithContext creates a new UnixGroupUsersCreateParams object
// with the ability to set a context for a request.
func NewUnixGroupUsersCreateParamsWithContext(ctx context.Context) *UnixGroupUsersCreateParams {
	return &UnixGroupUsersCreateParams{
		Context: ctx,
	}
}

// NewUnixGroupUsersCreateParamsWithHTTPClient creates a new UnixGroupUsersCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnixGroupUsersCreateParamsWithHTTPClient(client *http.Client) *UnixGroupUsersCreateParams {
	return &UnixGroupUsersCreateParams{
		HTTPClient: client,
	}
}

/*
UnixGroupUsersCreateParams contains all the parameters to send to the API endpoint

	for the unix group users create operation.

	Typically these are written to a http.Request.
*/
type UnixGroupUsersCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.UnixGroupUsers

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	/* UnixGroupName.

	   UNIX group name.
	*/
	UnixGroupNamePathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unix group users create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupUsersCreateParams) WithDefaults() *UnixGroupUsersCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unix group users create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupUsersCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)
	)

	val := UnixGroupUsersCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the unix group users create params
func (o *UnixGroupUsersCreateParams) WithTimeout(timeout time.Duration) *UnixGroupUsersCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unix group users create params
func (o *UnixGroupUsersCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unix group users create params
func (o *UnixGroupUsersCreateParams) WithContext(ctx context.Context) *UnixGroupUsersCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unix group users create params
func (o *UnixGroupUsersCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unix group users create params
func (o *UnixGroupUsersCreateParams) WithHTTPClient(client *http.Client) *UnixGroupUsersCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unix group users create params
func (o *UnixGroupUsersCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the unix group users create params
func (o *UnixGroupUsersCreateParams) WithInfo(info *models.UnixGroupUsers) *UnixGroupUsersCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the unix group users create params
func (o *UnixGroupUsersCreateParams) SetInfo(info *models.UnixGroupUsers) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the unix group users create params
func (o *UnixGroupUsersCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *UnixGroupUsersCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the unix group users create params
func (o *UnixGroupUsersCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithSVMUUIDPathParameter adds the svmUUID to the unix group users create params
func (o *UnixGroupUsersCreateParams) WithSVMUUIDPathParameter(svmUUID string) *UnixGroupUsersCreateParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the unix group users create params
func (o *UnixGroupUsersCreateParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WithUnixGroupNamePathParameter adds the unixGroupName to the unix group users create params
func (o *UnixGroupUsersCreateParams) WithUnixGroupNamePathParameter(unixGroupName string) *UnixGroupUsersCreateParams {
	o.SetUnixGroupNamePathParameter(unixGroupName)
	return o
}

// SetUnixGroupNamePathParameter adds the unixGroupName to the unix group users create params
func (o *UnixGroupUsersCreateParams) SetUnixGroupNamePathParameter(unixGroupName string) {
	o.UnixGroupNamePathParameter = unixGroupName
}

// WriteToRequest writes these params to a swagger request
func (o *UnixGroupUsersCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if o.ReturnRecordsQueryParameter != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecordsQueryParameter != nil {
			qrReturnRecords = *o.ReturnRecordsQueryParameter
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	// path param unix_group.name
	if err := r.SetPathParam("unix_group.name", o.UnixGroupNamePathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
