// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// NewPortMetricsCollectionGetParams creates a new PortMetricsCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPortMetricsCollectionGetParams() *PortMetricsCollectionGetParams {
	return &PortMetricsCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPortMetricsCollectionGetParamsWithTimeout creates a new PortMetricsCollectionGetParams object
// with the ability to set a timeout on a request.
func NewPortMetricsCollectionGetParamsWithTimeout(timeout time.Duration) *PortMetricsCollectionGetParams {
	return &PortMetricsCollectionGetParams{
		timeout: timeout,
	}
}

// NewPortMetricsCollectionGetParamsWithContext creates a new PortMetricsCollectionGetParams object
// with the ability to set a context for a request.
func NewPortMetricsCollectionGetParamsWithContext(ctx context.Context) *PortMetricsCollectionGetParams {
	return &PortMetricsCollectionGetParams{
		Context: ctx,
	}
}

// NewPortMetricsCollectionGetParamsWithHTTPClient creates a new PortMetricsCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPortMetricsCollectionGetParamsWithHTTPClient(client *http.Client) *PortMetricsCollectionGetParams {
	return &PortMetricsCollectionGetParams{
		HTTPClient: client,
	}
}

/*
PortMetricsCollectionGetParams contains all the parameters to send to the API endpoint

	for the port metrics collection get operation.

	Typically these are written to a http.Request.
*/
type PortMetricsCollectionGetParams struct {

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

	   Unique identifier of the port.
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the port metrics collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortMetricsCollectionGetParams) WithDefaults() *PortMetricsCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the port metrics collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PortMetricsCollectionGetParams) SetDefaults() {
	var (
		intervalQueryParameterDefault = string("1h")

		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := PortMetricsCollectionGetParams{
		IntervalQueryParameter:      &intervalQueryParameterDefault,
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithTimeout(timeout time.Duration) *PortMetricsCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithContext(ctx context.Context) *PortMetricsCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithHTTPClient(client *http.Client) *PortMetricsCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDurationQueryParameter adds the duration to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithDurationQueryParameter(duration *string) *PortMetricsCollectionGetParams {
	o.SetDurationQueryParameter(duration)
	return o
}

// SetDurationQueryParameter adds the duration to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetDurationQueryParameter(duration *string) {
	o.DurationQueryParameter = duration
}

// WithFieldsQueryParameter adds the fields to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithFieldsQueryParameter(fields []string) *PortMetricsCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIntervalQueryParameter adds the interval to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithIntervalQueryParameter(interval *string) *PortMetricsCollectionGetParams {
	o.SetIntervalQueryParameter(interval)
	return o
}

// SetIntervalQueryParameter adds the interval to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetIntervalQueryParameter(interval *string) {
	o.IntervalQueryParameter = interval
}

// WithMaxRecordsQueryParameter adds the maxRecords to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *PortMetricsCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *PortMetricsCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *PortMetricsCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *PortMetricsCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithStatusQueryParameter adds the status to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithStatusQueryParameter(status *string) *PortMetricsCollectionGetParams {
	o.SetStatusQueryParameter(status)
	return o
}

// SetStatusQueryParameter adds the status to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetStatusQueryParameter(status *string) {
	o.StatusQueryParameter = status
}

// WithThroughputReadQueryParameter adds the throughputRead to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithThroughputReadQueryParameter(throughputRead *int64) *PortMetricsCollectionGetParams {
	o.SetThroughputReadQueryParameter(throughputRead)
	return o
}

// SetThroughputReadQueryParameter adds the throughputRead to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetThroughputReadQueryParameter(throughputRead *int64) {
	o.ThroughputReadQueryParameter = throughputRead
}

// WithThroughputTotalQueryParameter adds the throughputTotal to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithThroughputTotalQueryParameter(throughputTotal *int64) *PortMetricsCollectionGetParams {
	o.SetThroughputTotalQueryParameter(throughputTotal)
	return o
}

// SetThroughputTotalQueryParameter adds the throughputTotal to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetThroughputTotalQueryParameter(throughputTotal *int64) {
	o.ThroughputTotalQueryParameter = throughputTotal
}

// WithThroughputWriteQueryParameter adds the throughputWrite to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithThroughputWriteQueryParameter(throughputWrite *int64) *PortMetricsCollectionGetParams {
	o.SetThroughputWriteQueryParameter(throughputWrite)
	return o
}

// SetThroughputWriteQueryParameter adds the throughputWrite to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetThroughputWriteQueryParameter(throughputWrite *int64) {
	o.ThroughputWriteQueryParameter = throughputWrite
}

// WithTimestampQueryParameter adds the timestamp to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithTimestampQueryParameter(timestamp *string) *PortMetricsCollectionGetParams {
	o.SetTimestampQueryParameter(timestamp)
	return o
}

// SetTimestampQueryParameter adds the timestamp to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetTimestampQueryParameter(timestamp *string) {
	o.TimestampQueryParameter = timestamp
}

// WithUUIDPathParameter adds the uuid to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) WithUUIDPathParameter(uuid string) *PortMetricsCollectionGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the port metrics collection get params
func (o *PortMetricsCollectionGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *PortMetricsCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

// bindParamPortMetricsCollectionGet binds the parameter fields
func (o *PortMetricsCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamPortMetricsCollectionGet binds the parameter order_by
func (o *PortMetricsCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
