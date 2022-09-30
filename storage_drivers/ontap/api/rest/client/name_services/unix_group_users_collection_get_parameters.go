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

// NewUnixGroupUsersCollectionGetParams creates a new UnixGroupUsersCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnixGroupUsersCollectionGetParams() *UnixGroupUsersCollectionGetParams {
	return &UnixGroupUsersCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnixGroupUsersCollectionGetParamsWithTimeout creates a new UnixGroupUsersCollectionGetParams object
// with the ability to set a timeout on a request.
func NewUnixGroupUsersCollectionGetParamsWithTimeout(timeout time.Duration) *UnixGroupUsersCollectionGetParams {
	return &UnixGroupUsersCollectionGetParams{
		timeout: timeout,
	}
}

// NewUnixGroupUsersCollectionGetParamsWithContext creates a new UnixGroupUsersCollectionGetParams object
// with the ability to set a context for a request.
func NewUnixGroupUsersCollectionGetParamsWithContext(ctx context.Context) *UnixGroupUsersCollectionGetParams {
	return &UnixGroupUsersCollectionGetParams{
		Context: ctx,
	}
}

// NewUnixGroupUsersCollectionGetParamsWithHTTPClient creates a new UnixGroupUsersCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnixGroupUsersCollectionGetParamsWithHTTPClient(client *http.Client) *UnixGroupUsersCollectionGetParams {
	return &UnixGroupUsersCollectionGetParams{
		HTTPClient: client,
	}
}

/*
UnixGroupUsersCollectionGetParams contains all the parameters to send to the API endpoint

	for the unix group users collection get operation.

	Typically these are written to a http.Request.
*/
type UnixGroupUsersCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

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

// WithDefaults hydrates default values in the unix group users collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupUsersCollectionGetParams) WithDefaults() *UnixGroupUsersCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unix group users collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixGroupUsersCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := UnixGroupUsersCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) WithTimeout(timeout time.Duration) *UnixGroupUsersCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) WithContext(ctx context.Context) *UnixGroupUsersCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) WithHTTPClient(client *http.Client) *UnixGroupUsersCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) WithFieldsQueryParameter(fields []string) *UnixGroupUsersCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *UnixGroupUsersCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) WithNameQueryParameter(name *string) *UnixGroupUsersCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *UnixGroupUsersCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *UnixGroupUsersCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *UnixGroupUsersCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMUUIDPathParameter adds the svmUUID to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) WithSVMUUIDPathParameter(svmUUID string) *UnixGroupUsersCollectionGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WithUnixGroupNamePathParameter adds the unixGroupName to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) WithUnixGroupNamePathParameter(unixGroupName string) *UnixGroupUsersCollectionGetParams {
	o.SetUnixGroupNamePathParameter(unixGroupName)
	return o
}

// SetUnixGroupNamePathParameter adds the unixGroupName to the unix group users collection get params
func (o *UnixGroupUsersCollectionGetParams) SetUnixGroupNamePathParameter(unixGroupName string) {
	o.UnixGroupNamePathParameter = unixGroupName
}

// WriteToRequest writes these params to a swagger request
func (o *UnixGroupUsersCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.NameQueryParameter != nil {

		// query param name
		var qrName string

		if o.NameQueryParameter != nil {
			qrName = *o.NameQueryParameter
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
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

	if o.ReturnTimeoutQueryParameter != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeoutQueryParameter != nil {
			qrReturnTimeout = *o.ReturnTimeoutQueryParameter
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
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

// bindParamUnixGroupUsersCollectionGet binds the parameter fields
func (o *UnixGroupUsersCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamUnixGroupUsersCollectionGet binds the parameter order_by
func (o *UnixGroupUsersCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
