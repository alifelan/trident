// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewSplitStatusCollectionGetParams creates a new SplitStatusCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSplitStatusCollectionGetParams() *SplitStatusCollectionGetParams {
	return &SplitStatusCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSplitStatusCollectionGetParamsWithTimeout creates a new SplitStatusCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSplitStatusCollectionGetParamsWithTimeout(timeout time.Duration) *SplitStatusCollectionGetParams {
	return &SplitStatusCollectionGetParams{
		timeout: timeout,
	}
}

// NewSplitStatusCollectionGetParamsWithContext creates a new SplitStatusCollectionGetParams object
// with the ability to set a context for a request.
func NewSplitStatusCollectionGetParamsWithContext(ctx context.Context) *SplitStatusCollectionGetParams {
	return &SplitStatusCollectionGetParams{
		Context: ctx,
	}
}

// NewSplitStatusCollectionGetParamsWithHTTPClient creates a new SplitStatusCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSplitStatusCollectionGetParamsWithHTTPClient(client *http.Client) *SplitStatusCollectionGetParams {
	return &SplitStatusCollectionGetParams{
		HTTPClient: client,
	}
}

/*
SplitStatusCollectionGetParams contains all the parameters to send to the API endpoint

	for the split status collection get operation.

	Typically these are written to a http.Request.
*/
type SplitStatusCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* PendingSplits.

	   Filter by pending_splits
	*/
	PendingSplitsQueryParameter *int64

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

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* UnsplitCloneSize.

	   Filter by unsplit_clone_size
	*/
	UnsplitCloneSizeQueryParameter *int64

	/* VolumeName.

	   Filter by volume.name
	*/
	VolumeNameQueryParameter *string

	/* VolumeUUID.

	   Filter by volume.uuid
	*/
	VolumeUUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the split status collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SplitStatusCollectionGetParams) WithDefaults() *SplitStatusCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the split status collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SplitStatusCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := SplitStatusCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the split status collection get params
func (o *SplitStatusCollectionGetParams) WithTimeout(timeout time.Duration) *SplitStatusCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the split status collection get params
func (o *SplitStatusCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the split status collection get params
func (o *SplitStatusCollectionGetParams) WithContext(ctx context.Context) *SplitStatusCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the split status collection get params
func (o *SplitStatusCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the split status collection get params
func (o *SplitStatusCollectionGetParams) WithHTTPClient(client *http.Client) *SplitStatusCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the split status collection get params
func (o *SplitStatusCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the split status collection get params
func (o *SplitStatusCollectionGetParams) WithFieldsQueryParameter(fields []string) *SplitStatusCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the split status collection get params
func (o *SplitStatusCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the split status collection get params
func (o *SplitStatusCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *SplitStatusCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the split status collection get params
func (o *SplitStatusCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the split status collection get params
func (o *SplitStatusCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *SplitStatusCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the split status collection get params
func (o *SplitStatusCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithPendingSplitsQueryParameter adds the pendingSplits to the split status collection get params
func (o *SplitStatusCollectionGetParams) WithPendingSplitsQueryParameter(pendingSplits *int64) *SplitStatusCollectionGetParams {
	o.SetPendingSplitsQueryParameter(pendingSplits)
	return o
}

// SetPendingSplitsQueryParameter adds the pendingSplits to the split status collection get params
func (o *SplitStatusCollectionGetParams) SetPendingSplitsQueryParameter(pendingSplits *int64) {
	o.PendingSplitsQueryParameter = pendingSplits
}

// WithReturnRecordsQueryParameter adds the returnRecords to the split status collection get params
func (o *SplitStatusCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *SplitStatusCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the split status collection get params
func (o *SplitStatusCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the split status collection get params
func (o *SplitStatusCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SplitStatusCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the split status collection get params
func (o *SplitStatusCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the split status collection get params
func (o *SplitStatusCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *SplitStatusCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the split status collection get params
func (o *SplitStatusCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the split status collection get params
func (o *SplitStatusCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *SplitStatusCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the split status collection get params
func (o *SplitStatusCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithUnsplitCloneSizeQueryParameter adds the unsplitCloneSize to the split status collection get params
func (o *SplitStatusCollectionGetParams) WithUnsplitCloneSizeQueryParameter(unsplitCloneSize *int64) *SplitStatusCollectionGetParams {
	o.SetUnsplitCloneSizeQueryParameter(unsplitCloneSize)
	return o
}

// SetUnsplitCloneSizeQueryParameter adds the unsplitCloneSize to the split status collection get params
func (o *SplitStatusCollectionGetParams) SetUnsplitCloneSizeQueryParameter(unsplitCloneSize *int64) {
	o.UnsplitCloneSizeQueryParameter = unsplitCloneSize
}

// WithVolumeNameQueryParameter adds the volumeName to the split status collection get params
func (o *SplitStatusCollectionGetParams) WithVolumeNameQueryParameter(volumeName *string) *SplitStatusCollectionGetParams {
	o.SetVolumeNameQueryParameter(volumeName)
	return o
}

// SetVolumeNameQueryParameter adds the volumeName to the split status collection get params
func (o *SplitStatusCollectionGetParams) SetVolumeNameQueryParameter(volumeName *string) {
	o.VolumeNameQueryParameter = volumeName
}

// WithVolumeUUIDQueryParameter adds the volumeUUID to the split status collection get params
func (o *SplitStatusCollectionGetParams) WithVolumeUUIDQueryParameter(volumeUUID *string) *SplitStatusCollectionGetParams {
	o.SetVolumeUUIDQueryParameter(volumeUUID)
	return o
}

// SetVolumeUUIDQueryParameter adds the volumeUuid to the split status collection get params
func (o *SplitStatusCollectionGetParams) SetVolumeUUIDQueryParameter(volumeUUID *string) {
	o.VolumeUUIDQueryParameter = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SplitStatusCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.PendingSplitsQueryParameter != nil {

		// query param pending_splits
		var qrPendingSplits int64

		if o.PendingSplitsQueryParameter != nil {
			qrPendingSplits = *o.PendingSplitsQueryParameter
		}
		qPendingSplits := swag.FormatInt64(qrPendingSplits)
		if qPendingSplits != "" {

			if err := r.SetQueryParam("pending_splits", qPendingSplits); err != nil {
				return err
			}
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

	if o.SVMNameQueryParameter != nil {

		// query param svm.name
		var qrSvmName string

		if o.SVMNameQueryParameter != nil {
			qrSvmName = *o.SVMNameQueryParameter
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SVMUUIDQueryParameter != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SVMUUIDQueryParameter != nil {
			qrSvmUUID = *o.SVMUUIDQueryParameter
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if o.UnsplitCloneSizeQueryParameter != nil {

		// query param unsplit_clone_size
		var qrUnsplitCloneSize int64

		if o.UnsplitCloneSizeQueryParameter != nil {
			qrUnsplitCloneSize = *o.UnsplitCloneSizeQueryParameter
		}
		qUnsplitCloneSize := swag.FormatInt64(qrUnsplitCloneSize)
		if qUnsplitCloneSize != "" {

			if err := r.SetQueryParam("unsplit_clone_size", qUnsplitCloneSize); err != nil {
				return err
			}
		}
	}

	if o.VolumeNameQueryParameter != nil {

		// query param volume.name
		var qrVolumeName string

		if o.VolumeNameQueryParameter != nil {
			qrVolumeName = *o.VolumeNameQueryParameter
		}
		qVolumeName := qrVolumeName
		if qVolumeName != "" {

			if err := r.SetQueryParam("volume.name", qVolumeName); err != nil {
				return err
			}
		}
	}

	if o.VolumeUUIDQueryParameter != nil {

		// query param volume.uuid
		var qrVolumeUUID string

		if o.VolumeUUIDQueryParameter != nil {
			qrVolumeUUID = *o.VolumeUUIDQueryParameter
		}
		qVolumeUUID := qrVolumeUUID
		if qVolumeUUID != "" {

			if err := r.SetQueryParam("volume.uuid", qVolumeUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSplitStatusCollectionGet binds the parameter fields
func (o *SplitStatusCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSplitStatusCollectionGet binds the parameter order_by
func (o *SplitStatusCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
