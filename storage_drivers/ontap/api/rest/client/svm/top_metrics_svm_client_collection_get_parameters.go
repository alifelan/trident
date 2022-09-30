// Code generated by go-swagger; DO NOT EDIT.

package svm

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

// NewTopMetricsSvmClientCollectionGetParams creates a new TopMetricsSvmClientCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTopMetricsSvmClientCollectionGetParams() *TopMetricsSvmClientCollectionGetParams {
	return &TopMetricsSvmClientCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTopMetricsSvmClientCollectionGetParamsWithTimeout creates a new TopMetricsSvmClientCollectionGetParams object
// with the ability to set a timeout on a request.
func NewTopMetricsSvmClientCollectionGetParamsWithTimeout(timeout time.Duration) *TopMetricsSvmClientCollectionGetParams {
	return &TopMetricsSvmClientCollectionGetParams{
		timeout: timeout,
	}
}

// NewTopMetricsSvmClientCollectionGetParamsWithContext creates a new TopMetricsSvmClientCollectionGetParams object
// with the ability to set a context for a request.
func NewTopMetricsSvmClientCollectionGetParamsWithContext(ctx context.Context) *TopMetricsSvmClientCollectionGetParams {
	return &TopMetricsSvmClientCollectionGetParams{
		Context: ctx,
	}
}

// NewTopMetricsSvmClientCollectionGetParamsWithHTTPClient creates a new TopMetricsSvmClientCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewTopMetricsSvmClientCollectionGetParamsWithHTTPClient(client *http.Client) *TopMetricsSvmClientCollectionGetParams {
	return &TopMetricsSvmClientCollectionGetParams{
		HTTPClient: client,
	}
}

/*
TopMetricsSvmClientCollectionGetParams contains all the parameters to send to the API endpoint

	for the top metrics svm client collection get operation.

	Typically these are written to a http.Request.
*/
type TopMetricsSvmClientCollectionGetParams struct {

	/* ClientIP.

	   Filter by client_ip
	*/
	ClientIPQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* IopsErrorLowerBound.

	   Filter by iops.error.lower_bound
	*/
	IopsErrorLowerBoundQueryParameter *int64

	/* IopsErrorUpperBound.

	   Filter by iops.error.upper_bound
	*/
	IopsErrorUpperBoundQueryParameter *int64

	/* IopsRead.

	   Filter by iops.read
	*/
	IopsReadQueryParameter *int64

	/* IopsWrite.

	   Filter by iops.write
	*/
	IopsWriteQueryParameter *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

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

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   SVM UUID
	*/
	SVMUUIDPathParameter string

	/* ThroughputErrorLowerBound.

	   Filter by throughput.error.lower_bound
	*/
	ThroughputErrorLowerBoundQueryParameter *int64

	/* ThroughputErrorUpperBound.

	   Filter by throughput.error.upper_bound
	*/
	ThroughputErrorUpperBoundQueryParameter *int64

	/* ThroughputRead.

	   Filter by throughput.read
	*/
	ThroughputReadQueryParameter *int64

	/* ThroughputWrite.

	   Filter by throughput.write
	*/
	ThroughputWriteQueryParameter *int64

	/* TopMetric.

	   IO activity type

	   Default: "iops.read"
	*/
	TopMetricQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the top metrics svm client collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TopMetricsSvmClientCollectionGetParams) WithDefaults() *TopMetricsSvmClientCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the top metrics svm client collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TopMetricsSvmClientCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)

		topMetricQueryParameterDefault = string("iops.read")
	)

	val := TopMetricsSvmClientCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
		TopMetricQueryParameter:     &topMetricQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithTimeout(timeout time.Duration) *TopMetricsSvmClientCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithContext(ctx context.Context) *TopMetricsSvmClientCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithHTTPClient(client *http.Client) *TopMetricsSvmClientCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClientIPQueryParameter adds the clientIP to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithClientIPQueryParameter(clientIP *string) *TopMetricsSvmClientCollectionGetParams {
	o.SetClientIPQueryParameter(clientIP)
	return o
}

// SetClientIPQueryParameter adds the clientIp to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetClientIPQueryParameter(clientIP *string) {
	o.ClientIPQueryParameter = clientIP
}

// WithFieldsQueryParameter adds the fields to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithFieldsQueryParameter(fields []string) *TopMetricsSvmClientCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIopsErrorLowerBoundQueryParameter adds the iopsErrorLowerBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithIopsErrorLowerBoundQueryParameter(iopsErrorLowerBound *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetIopsErrorLowerBoundQueryParameter(iopsErrorLowerBound)
	return o
}

// SetIopsErrorLowerBoundQueryParameter adds the iopsErrorLowerBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetIopsErrorLowerBoundQueryParameter(iopsErrorLowerBound *int64) {
	o.IopsErrorLowerBoundQueryParameter = iopsErrorLowerBound
}

// WithIopsErrorUpperBoundQueryParameter adds the iopsErrorUpperBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithIopsErrorUpperBoundQueryParameter(iopsErrorUpperBound *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetIopsErrorUpperBoundQueryParameter(iopsErrorUpperBound)
	return o
}

// SetIopsErrorUpperBoundQueryParameter adds the iopsErrorUpperBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetIopsErrorUpperBoundQueryParameter(iopsErrorUpperBound *int64) {
	o.IopsErrorUpperBoundQueryParameter = iopsErrorUpperBound
}

// WithIopsReadQueryParameter adds the iopsRead to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithIopsReadQueryParameter(iopsRead *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetIopsReadQueryParameter(iopsRead)
	return o
}

// SetIopsReadQueryParameter adds the iopsRead to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetIopsReadQueryParameter(iopsRead *int64) {
	o.IopsReadQueryParameter = iopsRead
}

// WithIopsWriteQueryParameter adds the iopsWrite to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithIopsWriteQueryParameter(iopsWrite *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetIopsWriteQueryParameter(iopsWrite)
	return o
}

// SetIopsWriteQueryParameter adds the iopsWrite to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetIopsWriteQueryParameter(iopsWrite *int64) {
	o.IopsWriteQueryParameter = iopsWrite
}

// WithMaxRecordsQueryParameter adds the maxRecords to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *TopMetricsSvmClientCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *TopMetricsSvmClientCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *TopMetricsSvmClientCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDPathParameter adds the svmUUID to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithSVMUUIDPathParameter(svmUUID string) *TopMetricsSvmClientCollectionGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WithThroughputErrorLowerBoundQueryParameter adds the throughputErrorLowerBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithThroughputErrorLowerBoundQueryParameter(throughputErrorLowerBound *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetThroughputErrorLowerBoundQueryParameter(throughputErrorLowerBound)
	return o
}

// SetThroughputErrorLowerBoundQueryParameter adds the throughputErrorLowerBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetThroughputErrorLowerBoundQueryParameter(throughputErrorLowerBound *int64) {
	o.ThroughputErrorLowerBoundQueryParameter = throughputErrorLowerBound
}

// WithThroughputErrorUpperBoundQueryParameter adds the throughputErrorUpperBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithThroughputErrorUpperBoundQueryParameter(throughputErrorUpperBound *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetThroughputErrorUpperBoundQueryParameter(throughputErrorUpperBound)
	return o
}

// SetThroughputErrorUpperBoundQueryParameter adds the throughputErrorUpperBound to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetThroughputErrorUpperBoundQueryParameter(throughputErrorUpperBound *int64) {
	o.ThroughputErrorUpperBoundQueryParameter = throughputErrorUpperBound
}

// WithThroughputReadQueryParameter adds the throughputRead to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithThroughputReadQueryParameter(throughputRead *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetThroughputReadQueryParameter(throughputRead)
	return o
}

// SetThroughputReadQueryParameter adds the throughputRead to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetThroughputReadQueryParameter(throughputRead *int64) {
	o.ThroughputReadQueryParameter = throughputRead
}

// WithThroughputWriteQueryParameter adds the throughputWrite to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithThroughputWriteQueryParameter(throughputWrite *int64) *TopMetricsSvmClientCollectionGetParams {
	o.SetThroughputWriteQueryParameter(throughputWrite)
	return o
}

// SetThroughputWriteQueryParameter adds the throughputWrite to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetThroughputWriteQueryParameter(throughputWrite *int64) {
	o.ThroughputWriteQueryParameter = throughputWrite
}

// WithTopMetricQueryParameter adds the topMetric to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) WithTopMetricQueryParameter(topMetric *string) *TopMetricsSvmClientCollectionGetParams {
	o.SetTopMetricQueryParameter(topMetric)
	return o
}

// SetTopMetricQueryParameter adds the topMetric to the top metrics svm client collection get params
func (o *TopMetricsSvmClientCollectionGetParams) SetTopMetricQueryParameter(topMetric *string) {
	o.TopMetricQueryParameter = topMetric
}

// WriteToRequest writes these params to a swagger request
func (o *TopMetricsSvmClientCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClientIPQueryParameter != nil {

		// query param client_ip
		var qrClientIP string

		if o.ClientIPQueryParameter != nil {
			qrClientIP = *o.ClientIPQueryParameter
		}
		qClientIP := qrClientIP
		if qClientIP != "" {

			if err := r.SetQueryParam("client_ip", qClientIP); err != nil {
				return err
			}
		}
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.IopsErrorLowerBoundQueryParameter != nil {

		// query param iops.error.lower_bound
		var qrIopsErrorLowerBound int64

		if o.IopsErrorLowerBoundQueryParameter != nil {
			qrIopsErrorLowerBound = *o.IopsErrorLowerBoundQueryParameter
		}
		qIopsErrorLowerBound := swag.FormatInt64(qrIopsErrorLowerBound)
		if qIopsErrorLowerBound != "" {

			if err := r.SetQueryParam("iops.error.lower_bound", qIopsErrorLowerBound); err != nil {
				return err
			}
		}
	}

	if o.IopsErrorUpperBoundQueryParameter != nil {

		// query param iops.error.upper_bound
		var qrIopsErrorUpperBound int64

		if o.IopsErrorUpperBoundQueryParameter != nil {
			qrIopsErrorUpperBound = *o.IopsErrorUpperBoundQueryParameter
		}
		qIopsErrorUpperBound := swag.FormatInt64(qrIopsErrorUpperBound)
		if qIopsErrorUpperBound != "" {

			if err := r.SetQueryParam("iops.error.upper_bound", qIopsErrorUpperBound); err != nil {
				return err
			}
		}
	}

	if o.IopsReadQueryParameter != nil {

		// query param iops.read
		var qrIopsRead int64

		if o.IopsReadQueryParameter != nil {
			qrIopsRead = *o.IopsReadQueryParameter
		}
		qIopsRead := swag.FormatInt64(qrIopsRead)
		if qIopsRead != "" {

			if err := r.SetQueryParam("iops.read", qIopsRead); err != nil {
				return err
			}
		}
	}

	if o.IopsWriteQueryParameter != nil {

		// query param iops.write
		var qrIopsWrite int64

		if o.IopsWriteQueryParameter != nil {
			qrIopsWrite = *o.IopsWriteQueryParameter
		}
		qIopsWrite := swag.FormatInt64(qrIopsWrite)
		if qIopsWrite != "" {

			if err := r.SetQueryParam("iops.write", qIopsWrite); err != nil {
				return err
			}
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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if o.ThroughputErrorLowerBoundQueryParameter != nil {

		// query param throughput.error.lower_bound
		var qrThroughputErrorLowerBound int64

		if o.ThroughputErrorLowerBoundQueryParameter != nil {
			qrThroughputErrorLowerBound = *o.ThroughputErrorLowerBoundQueryParameter
		}
		qThroughputErrorLowerBound := swag.FormatInt64(qrThroughputErrorLowerBound)
		if qThroughputErrorLowerBound != "" {

			if err := r.SetQueryParam("throughput.error.lower_bound", qThroughputErrorLowerBound); err != nil {
				return err
			}
		}
	}

	if o.ThroughputErrorUpperBoundQueryParameter != nil {

		// query param throughput.error.upper_bound
		var qrThroughputErrorUpperBound int64

		if o.ThroughputErrorUpperBoundQueryParameter != nil {
			qrThroughputErrorUpperBound = *o.ThroughputErrorUpperBoundQueryParameter
		}
		qThroughputErrorUpperBound := swag.FormatInt64(qrThroughputErrorUpperBound)
		if qThroughputErrorUpperBound != "" {

			if err := r.SetQueryParam("throughput.error.upper_bound", qThroughputErrorUpperBound); err != nil {
				return err
			}
		}
	}

	if o.ThroughputReadQueryParameter != nil {

		// query param throughput.read
		var qrThroughputRead int64

		if o.ThroughputReadQueryParameter != nil {
			qrThroughputRead = *o.ThroughputReadQueryParameter
		}
		qThroughputRead := swag.FormatInt64(qrThroughputRead)
		if qThroughputRead != "" {

			if err := r.SetQueryParam("throughput.read", qThroughputRead); err != nil {
				return err
			}
		}
	}

	if o.ThroughputWriteQueryParameter != nil {

		// query param throughput.write
		var qrThroughputWrite int64

		if o.ThroughputWriteQueryParameter != nil {
			qrThroughputWrite = *o.ThroughputWriteQueryParameter
		}
		qThroughputWrite := swag.FormatInt64(qrThroughputWrite)
		if qThroughputWrite != "" {

			if err := r.SetQueryParam("throughput.write", qThroughputWrite); err != nil {
				return err
			}
		}
	}

	if o.TopMetricQueryParameter != nil {

		// query param top_metric
		var qrTopMetric string

		if o.TopMetricQueryParameter != nil {
			qrTopMetric = *o.TopMetricQueryParameter
		}
		qTopMetric := qrTopMetric
		if qTopMetric != "" {

			if err := r.SetQueryParam("top_metric", qTopMetric); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamTopMetricsSvmClientCollectionGet binds the parameter fields
func (o *TopMetricsSvmClientCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamTopMetricsSvmClientCollectionGet binds the parameter order_by
func (o *TopMetricsSvmClientCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
