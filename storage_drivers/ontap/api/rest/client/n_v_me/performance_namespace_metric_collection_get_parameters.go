// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

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

// NewPerformanceNamespaceMetricCollectionGetParams creates a new PerformanceNamespaceMetricCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPerformanceNamespaceMetricCollectionGetParams() *PerformanceNamespaceMetricCollectionGetParams {
	return &PerformanceNamespaceMetricCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPerformanceNamespaceMetricCollectionGetParamsWithTimeout creates a new PerformanceNamespaceMetricCollectionGetParams object
// with the ability to set a timeout on a request.
func NewPerformanceNamespaceMetricCollectionGetParamsWithTimeout(timeout time.Duration) *PerformanceNamespaceMetricCollectionGetParams {
	return &PerformanceNamespaceMetricCollectionGetParams{
		timeout: timeout,
	}
}

// NewPerformanceNamespaceMetricCollectionGetParamsWithContext creates a new PerformanceNamespaceMetricCollectionGetParams object
// with the ability to set a context for a request.
func NewPerformanceNamespaceMetricCollectionGetParamsWithContext(ctx context.Context) *PerformanceNamespaceMetricCollectionGetParams {
	return &PerformanceNamespaceMetricCollectionGetParams{
		Context: ctx,
	}
}

// NewPerformanceNamespaceMetricCollectionGetParamsWithHTTPClient creates a new PerformanceNamespaceMetricCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPerformanceNamespaceMetricCollectionGetParamsWithHTTPClient(client *http.Client) *PerformanceNamespaceMetricCollectionGetParams {
	return &PerformanceNamespaceMetricCollectionGetParams{
		HTTPClient: client,
	}
}

/*
PerformanceNamespaceMetricCollectionGetParams contains all the parameters to send to the API endpoint

	for the performance namespace metric collection get operation.

	Typically these are written to a http.Request.
*/
type PerformanceNamespaceMetricCollectionGetParams struct {

	/* Duration.

	   Filter by duration
	*/
	DurationQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Interval.

	     The time range for the data. Examples can be 1h, 1d, 1m, 1w, 1y.
	The period for each time range is as follows:
	* 1h: Metrics over the most recent hour sampled over 15 seconds.
	* 1d: Metrics over the most recent day sampled over 5 minutes.
	* 1w: Metrics over the most recent week sampled over 30 minutes.
	* 1m: Metrics over the most recent month sampled over 2 hours.
	* 1y: Metrics over the most recent year sampled over a day.


	     Default: "1h"
	*/
	IntervalQueryParameter *string

	/* IopsOther.

	   Filter by iops.other
	*/
	IopsOtherQueryParameter *int64

	/* IopsRead.

	   Filter by iops.read
	*/
	IopsReadQueryParameter *int64

	/* IopsTotal.

	   Filter by iops.total
	*/
	IopsTotalQueryParameter *int64

	/* IopsWrite.

	   Filter by iops.write
	*/
	IopsWriteQueryParameter *int64

	/* LatencyOther.

	   Filter by latency.other
	*/
	LatencyOtherQueryParameter *int64

	/* LatencyRead.

	   Filter by latency.read
	*/
	LatencyReadQueryParameter *int64

	/* LatencyTotal.

	   Filter by latency.total
	*/
	LatencyTotalQueryParameter *int64

	/* LatencyWrite.

	   Filter by latency.write
	*/
	LatencyWriteQueryParameter *int64

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

	/* Status.

	   Filter by status
	*/
	StatusQueryParameter *string

	/* ThroughputRead.

	   Filter by throughput.read
	*/
	ThroughputReadQueryParameter *int64

	/* ThroughputTotal.

	   Filter by throughput.total
	*/
	ThroughputTotalQueryParameter *int64

	/* ThroughputWrite.

	   Filter by throughput.write
	*/
	ThroughputWriteQueryParameter *int64

	/* Timestamp.

	   Filter by timestamp
	*/
	TimestampQueryParameter *string

	/* UUID.

	   Unique identifier of the NVMe namespace.
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the performance namespace metric collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformanceNamespaceMetricCollectionGetParams) WithDefaults() *PerformanceNamespaceMetricCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the performance namespace metric collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PerformanceNamespaceMetricCollectionGetParams) SetDefaults() {
	var (
		intervalQueryParameterDefault = string("1h")

		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := PerformanceNamespaceMetricCollectionGetParams{
		IntervalQueryParameter:      &intervalQueryParameterDefault,
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithTimeout(timeout time.Duration) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithContext(ctx context.Context) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithHTTPClient(client *http.Client) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDurationQueryParameter adds the duration to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithDurationQueryParameter(duration *string) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetDurationQueryParameter(duration)
	return o
}

// SetDurationQueryParameter adds the duration to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetDurationQueryParameter(duration *string) {
	o.DurationQueryParameter = duration
}

// WithFieldsQueryParameter adds the fields to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithFieldsQueryParameter(fields []string) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIntervalQueryParameter adds the interval to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithIntervalQueryParameter(interval *string) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetIntervalQueryParameter(interval)
	return o
}

// SetIntervalQueryParameter adds the interval to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetIntervalQueryParameter(interval *string) {
	o.IntervalQueryParameter = interval
}

// WithIopsOtherQueryParameter adds the iopsOther to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithIopsOtherQueryParameter(iopsOther *int64) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetIopsOtherQueryParameter(iopsOther)
	return o
}

// SetIopsOtherQueryParameter adds the iopsOther to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetIopsOtherQueryParameter(iopsOther *int64) {
	o.IopsOtherQueryParameter = iopsOther
}

// WithIopsReadQueryParameter adds the iopsRead to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithIopsReadQueryParameter(iopsRead *int64) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetIopsReadQueryParameter(iopsRead)
	return o
}

// SetIopsReadQueryParameter adds the iopsRead to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetIopsReadQueryParameter(iopsRead *int64) {
	o.IopsReadQueryParameter = iopsRead
}

// WithIopsTotalQueryParameter adds the iopsTotal to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithIopsTotalQueryParameter(iopsTotal *int64) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetIopsTotalQueryParameter(iopsTotal)
	return o
}

// SetIopsTotalQueryParameter adds the iopsTotal to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetIopsTotalQueryParameter(iopsTotal *int64) {
	o.IopsTotalQueryParameter = iopsTotal
}

// WithIopsWriteQueryParameter adds the iopsWrite to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithIopsWriteQueryParameter(iopsWrite *int64) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetIopsWriteQueryParameter(iopsWrite)
	return o
}

// SetIopsWriteQueryParameter adds the iopsWrite to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetIopsWriteQueryParameter(iopsWrite *int64) {
	o.IopsWriteQueryParameter = iopsWrite
}

// WithLatencyOtherQueryParameter adds the latencyOther to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithLatencyOtherQueryParameter(latencyOther *int64) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetLatencyOtherQueryParameter(latencyOther)
	return o
}

// SetLatencyOtherQueryParameter adds the latencyOther to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetLatencyOtherQueryParameter(latencyOther *int64) {
	o.LatencyOtherQueryParameter = latencyOther
}

// WithLatencyReadQueryParameter adds the latencyRead to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithLatencyReadQueryParameter(latencyRead *int64) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetLatencyReadQueryParameter(latencyRead)
	return o
}

// SetLatencyReadQueryParameter adds the latencyRead to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetLatencyReadQueryParameter(latencyRead *int64) {
	o.LatencyReadQueryParameter = latencyRead
}

// WithLatencyTotalQueryParameter adds the latencyTotal to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithLatencyTotalQueryParameter(latencyTotal *int64) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetLatencyTotalQueryParameter(latencyTotal)
	return o
}

// SetLatencyTotalQueryParameter adds the latencyTotal to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetLatencyTotalQueryParameter(latencyTotal *int64) {
	o.LatencyTotalQueryParameter = latencyTotal
}

// WithLatencyWriteQueryParameter adds the latencyWrite to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithLatencyWriteQueryParameter(latencyWrite *int64) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetLatencyWriteQueryParameter(latencyWrite)
	return o
}

// SetLatencyWriteQueryParameter adds the latencyWrite to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetLatencyWriteQueryParameter(latencyWrite *int64) {
	o.LatencyWriteQueryParameter = latencyWrite
}

// WithMaxRecordsQueryParameter adds the maxRecords to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithStatusQueryParameter adds the status to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithStatusQueryParameter(status *string) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetStatusQueryParameter(status)
	return o
}

// SetStatusQueryParameter adds the status to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetStatusQueryParameter(status *string) {
	o.StatusQueryParameter = status
}

// WithThroughputReadQueryParameter adds the throughputRead to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithThroughputReadQueryParameter(throughputRead *int64) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetThroughputReadQueryParameter(throughputRead)
	return o
}

// SetThroughputReadQueryParameter adds the throughputRead to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetThroughputReadQueryParameter(throughputRead *int64) {
	o.ThroughputReadQueryParameter = throughputRead
}

// WithThroughputTotalQueryParameter adds the throughputTotal to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithThroughputTotalQueryParameter(throughputTotal *int64) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetThroughputTotalQueryParameter(throughputTotal)
	return o
}

// SetThroughputTotalQueryParameter adds the throughputTotal to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetThroughputTotalQueryParameter(throughputTotal *int64) {
	o.ThroughputTotalQueryParameter = throughputTotal
}

// WithThroughputWriteQueryParameter adds the throughputWrite to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithThroughputWriteQueryParameter(throughputWrite *int64) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetThroughputWriteQueryParameter(throughputWrite)
	return o
}

// SetThroughputWriteQueryParameter adds the throughputWrite to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetThroughputWriteQueryParameter(throughputWrite *int64) {
	o.ThroughputWriteQueryParameter = throughputWrite
}

// WithTimestampQueryParameter adds the timestamp to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithTimestampQueryParameter(timestamp *string) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetTimestampQueryParameter(timestamp)
	return o
}

// SetTimestampQueryParameter adds the timestamp to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetTimestampQueryParameter(timestamp *string) {
	o.TimestampQueryParameter = timestamp
}

// WithUUIDPathParameter adds the uuid to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) WithUUIDPathParameter(uuid string) *PerformanceNamespaceMetricCollectionGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the performance namespace metric collection get params
func (o *PerformanceNamespaceMetricCollectionGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PerformanceNamespaceMetricCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DurationQueryParameter != nil {

		// query param duration
		var qrDuration string

		if o.DurationQueryParameter != nil {
			qrDuration = *o.DurationQueryParameter
		}
		qDuration := qrDuration
		if qDuration != "" {

			if err := r.SetQueryParam("duration", qDuration); err != nil {
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

	if o.IntervalQueryParameter != nil {

		// query param interval
		var qrInterval string

		if o.IntervalQueryParameter != nil {
			qrInterval = *o.IntervalQueryParameter
		}
		qInterval := qrInterval
		if qInterval != "" {

			if err := r.SetQueryParam("interval", qInterval); err != nil {
				return err
			}
		}
	}

	if o.IopsOtherQueryParameter != nil {

		// query param iops.other
		var qrIopsOther int64

		if o.IopsOtherQueryParameter != nil {
			qrIopsOther = *o.IopsOtherQueryParameter
		}
		qIopsOther := swag.FormatInt64(qrIopsOther)
		if qIopsOther != "" {

			if err := r.SetQueryParam("iops.other", qIopsOther); err != nil {
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

	if o.IopsTotalQueryParameter != nil {

		// query param iops.total
		var qrIopsTotal int64

		if o.IopsTotalQueryParameter != nil {
			qrIopsTotal = *o.IopsTotalQueryParameter
		}
		qIopsTotal := swag.FormatInt64(qrIopsTotal)
		if qIopsTotal != "" {

			if err := r.SetQueryParam("iops.total", qIopsTotal); err != nil {
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

	if o.LatencyOtherQueryParameter != nil {

		// query param latency.other
		var qrLatencyOther int64

		if o.LatencyOtherQueryParameter != nil {
			qrLatencyOther = *o.LatencyOtherQueryParameter
		}
		qLatencyOther := swag.FormatInt64(qrLatencyOther)
		if qLatencyOther != "" {

			if err := r.SetQueryParam("latency.other", qLatencyOther); err != nil {
				return err
			}
		}
	}

	if o.LatencyReadQueryParameter != nil {

		// query param latency.read
		var qrLatencyRead int64

		if o.LatencyReadQueryParameter != nil {
			qrLatencyRead = *o.LatencyReadQueryParameter
		}
		qLatencyRead := swag.FormatInt64(qrLatencyRead)
		if qLatencyRead != "" {

			if err := r.SetQueryParam("latency.read", qLatencyRead); err != nil {
				return err
			}
		}
	}

	if o.LatencyTotalQueryParameter != nil {

		// query param latency.total
		var qrLatencyTotal int64

		if o.LatencyTotalQueryParameter != nil {
			qrLatencyTotal = *o.LatencyTotalQueryParameter
		}
		qLatencyTotal := swag.FormatInt64(qrLatencyTotal)
		if qLatencyTotal != "" {

			if err := r.SetQueryParam("latency.total", qLatencyTotal); err != nil {
				return err
			}
		}
	}

	if o.LatencyWriteQueryParameter != nil {

		// query param latency.write
		var qrLatencyWrite int64

		if o.LatencyWriteQueryParameter != nil {
			qrLatencyWrite = *o.LatencyWriteQueryParameter
		}
		qLatencyWrite := swag.FormatInt64(qrLatencyWrite)
		if qLatencyWrite != "" {

			if err := r.SetQueryParam("latency.write", qLatencyWrite); err != nil {
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

	if o.StatusQueryParameter != nil {

		// query param status
		var qrStatus string

		if o.StatusQueryParameter != nil {
			qrStatus = *o.StatusQueryParameter
		}
		qStatus := qrStatus
		if qStatus != "" {

			if err := r.SetQueryParam("status", qStatus); err != nil {
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

	if o.ThroughputTotalQueryParameter != nil {

		// query param throughput.total
		var qrThroughputTotal int64

		if o.ThroughputTotalQueryParameter != nil {
			qrThroughputTotal = *o.ThroughputTotalQueryParameter
		}
		qThroughputTotal := swag.FormatInt64(qrThroughputTotal)
		if qThroughputTotal != "" {

			if err := r.SetQueryParam("throughput.total", qThroughputTotal); err != nil {
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

	if o.TimestampQueryParameter != nil {

		// query param timestamp
		var qrTimestamp string

		if o.TimestampQueryParameter != nil {
			qrTimestamp = *o.TimestampQueryParameter
		}
		qTimestamp := qrTimestamp
		if qTimestamp != "" {

			if err := r.SetQueryParam("timestamp", qTimestamp); err != nil {
				return err
			}
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

// bindParamPerformanceNamespaceMetricCollectionGet binds the parameter fields
func (o *PerformanceNamespaceMetricCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamPerformanceNamespaceMetricCollectionGet binds the parameter order_by
func (o *PerformanceNamespaceMetricCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
