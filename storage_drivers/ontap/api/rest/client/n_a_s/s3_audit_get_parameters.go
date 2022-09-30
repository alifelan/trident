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

// NewS3AuditGetParams creates a new S3AuditGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3AuditGetParams() *S3AuditGetParams {
	return &S3AuditGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3AuditGetParamsWithTimeout creates a new S3AuditGetParams object
// with the ability to set a timeout on a request.
func NewS3AuditGetParamsWithTimeout(timeout time.Duration) *S3AuditGetParams {
	return &S3AuditGetParams{
		timeout: timeout,
	}
}

// NewS3AuditGetParamsWithContext creates a new S3AuditGetParams object
// with the ability to set a context for a request.
func NewS3AuditGetParamsWithContext(ctx context.Context) *S3AuditGetParams {
	return &S3AuditGetParams{
		Context: ctx,
	}
}

// NewS3AuditGetParamsWithHTTPClient creates a new S3AuditGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3AuditGetParamsWithHTTPClient(client *http.Client) *S3AuditGetParams {
	return &S3AuditGetParams{
		HTTPClient: client,
	}
}

/*
S3AuditGetParams contains all the parameters to send to the API endpoint

	for the s3 audit get operation.

	Typically these are written to a http.Request.
*/
type S3AuditGetParams struct {

	/* Enabled.

	   Filter by enabled
	*/
	EnabledQueryParameter *bool

	/* EventsData.

	   Filter by events.data
	*/
	EventsDataQueryParameter *bool

	/* EventsManagement.

	   Filter by events.management
	*/
	EventsManagementQueryParameter *bool

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* LogFormat.

	   Filter by log.format
	*/
	LogFormatQueryParameter *string

	/* LogRetentionCount.

	   Filter by log.retention.count
	*/
	LogRetentionCountQueryParameter *int64

	/* LogRetentionDuration.

	   Filter by log.retention.duration
	*/
	LogRetentionDurationQueryParameter *string

	/* LogRotationScheduleDays.

	   Filter by log.rotation.schedule.days
	*/
	LogRotationScheduleDaysQueryParameter *int64

	/* LogRotationScheduleHours.

	   Filter by log.rotation.schedule.hours
	*/
	LogRotationScheduleHoursQueryParameter *int64

	/* LogRotationScheduleMinutes.

	   Filter by log.rotation.schedule.minutes
	*/
	LogRotationScheduleMinutesQueryParameter *int64

	/* LogRotationScheduleMonths.

	   Filter by log.rotation.schedule.months
	*/
	LogRotationScheduleMonthsQueryParameter *int64

	/* LogRotationScheduleWeekdays.

	   Filter by log.rotation.schedule.weekdays
	*/
	LogRotationScheduleWeekdaysQueryParameter *int64

	/* LogRotationSize.

	   Filter by log.rotation.size
	*/
	LogRotationSizeQueryParameter *int64

	/* LogPath.

	   Filter by log_path
	*/
	LogPathQueryParameter *string

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

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 audit get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3AuditGetParams) WithDefaults() *S3AuditGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 audit get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3AuditGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := S3AuditGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the s3 audit get params
func (o *S3AuditGetParams) WithTimeout(timeout time.Duration) *S3AuditGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 audit get params
func (o *S3AuditGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 audit get params
func (o *S3AuditGetParams) WithContext(ctx context.Context) *S3AuditGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 audit get params
func (o *S3AuditGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 audit get params
func (o *S3AuditGetParams) WithHTTPClient(client *http.Client) *S3AuditGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 audit get params
func (o *S3AuditGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabledQueryParameter adds the enabled to the s3 audit get params
func (o *S3AuditGetParams) WithEnabledQueryParameter(enabled *bool) *S3AuditGetParams {
	o.SetEnabledQueryParameter(enabled)
	return o
}

// SetEnabledQueryParameter adds the enabled to the s3 audit get params
func (o *S3AuditGetParams) SetEnabledQueryParameter(enabled *bool) {
	o.EnabledQueryParameter = enabled
}

// WithEventsDataQueryParameter adds the eventsData to the s3 audit get params
func (o *S3AuditGetParams) WithEventsDataQueryParameter(eventsData *bool) *S3AuditGetParams {
	o.SetEventsDataQueryParameter(eventsData)
	return o
}

// SetEventsDataQueryParameter adds the eventsData to the s3 audit get params
func (o *S3AuditGetParams) SetEventsDataQueryParameter(eventsData *bool) {
	o.EventsDataQueryParameter = eventsData
}

// WithEventsManagementQueryParameter adds the eventsManagement to the s3 audit get params
func (o *S3AuditGetParams) WithEventsManagementQueryParameter(eventsManagement *bool) *S3AuditGetParams {
	o.SetEventsManagementQueryParameter(eventsManagement)
	return o
}

// SetEventsManagementQueryParameter adds the eventsManagement to the s3 audit get params
func (o *S3AuditGetParams) SetEventsManagementQueryParameter(eventsManagement *bool) {
	o.EventsManagementQueryParameter = eventsManagement
}

// WithFieldsQueryParameter adds the fields to the s3 audit get params
func (o *S3AuditGetParams) WithFieldsQueryParameter(fields []string) *S3AuditGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the s3 audit get params
func (o *S3AuditGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithLogFormatQueryParameter adds the logFormat to the s3 audit get params
func (o *S3AuditGetParams) WithLogFormatQueryParameter(logFormat *string) *S3AuditGetParams {
	o.SetLogFormatQueryParameter(logFormat)
	return o
}

// SetLogFormatQueryParameter adds the logFormat to the s3 audit get params
func (o *S3AuditGetParams) SetLogFormatQueryParameter(logFormat *string) {
	o.LogFormatQueryParameter = logFormat
}

// WithLogRetentionCountQueryParameter adds the logRetentionCount to the s3 audit get params
func (o *S3AuditGetParams) WithLogRetentionCountQueryParameter(logRetentionCount *int64) *S3AuditGetParams {
	o.SetLogRetentionCountQueryParameter(logRetentionCount)
	return o
}

// SetLogRetentionCountQueryParameter adds the logRetentionCount to the s3 audit get params
func (o *S3AuditGetParams) SetLogRetentionCountQueryParameter(logRetentionCount *int64) {
	o.LogRetentionCountQueryParameter = logRetentionCount
}

// WithLogRetentionDurationQueryParameter adds the logRetentionDuration to the s3 audit get params
func (o *S3AuditGetParams) WithLogRetentionDurationQueryParameter(logRetentionDuration *string) *S3AuditGetParams {
	o.SetLogRetentionDurationQueryParameter(logRetentionDuration)
	return o
}

// SetLogRetentionDurationQueryParameter adds the logRetentionDuration to the s3 audit get params
func (o *S3AuditGetParams) SetLogRetentionDurationQueryParameter(logRetentionDuration *string) {
	o.LogRetentionDurationQueryParameter = logRetentionDuration
}

// WithLogRotationScheduleDaysQueryParameter adds the logRotationScheduleDays to the s3 audit get params
func (o *S3AuditGetParams) WithLogRotationScheduleDaysQueryParameter(logRotationScheduleDays *int64) *S3AuditGetParams {
	o.SetLogRotationScheduleDaysQueryParameter(logRotationScheduleDays)
	return o
}

// SetLogRotationScheduleDaysQueryParameter adds the logRotationScheduleDays to the s3 audit get params
func (o *S3AuditGetParams) SetLogRotationScheduleDaysQueryParameter(logRotationScheduleDays *int64) {
	o.LogRotationScheduleDaysQueryParameter = logRotationScheduleDays
}

// WithLogRotationScheduleHoursQueryParameter adds the logRotationScheduleHours to the s3 audit get params
func (o *S3AuditGetParams) WithLogRotationScheduleHoursQueryParameter(logRotationScheduleHours *int64) *S3AuditGetParams {
	o.SetLogRotationScheduleHoursQueryParameter(logRotationScheduleHours)
	return o
}

// SetLogRotationScheduleHoursQueryParameter adds the logRotationScheduleHours to the s3 audit get params
func (o *S3AuditGetParams) SetLogRotationScheduleHoursQueryParameter(logRotationScheduleHours *int64) {
	o.LogRotationScheduleHoursQueryParameter = logRotationScheduleHours
}

// WithLogRotationScheduleMinutesQueryParameter adds the logRotationScheduleMinutes to the s3 audit get params
func (o *S3AuditGetParams) WithLogRotationScheduleMinutesQueryParameter(logRotationScheduleMinutes *int64) *S3AuditGetParams {
	o.SetLogRotationScheduleMinutesQueryParameter(logRotationScheduleMinutes)
	return o
}

// SetLogRotationScheduleMinutesQueryParameter adds the logRotationScheduleMinutes to the s3 audit get params
func (o *S3AuditGetParams) SetLogRotationScheduleMinutesQueryParameter(logRotationScheduleMinutes *int64) {
	o.LogRotationScheduleMinutesQueryParameter = logRotationScheduleMinutes
}

// WithLogRotationScheduleMonthsQueryParameter adds the logRotationScheduleMonths to the s3 audit get params
func (o *S3AuditGetParams) WithLogRotationScheduleMonthsQueryParameter(logRotationScheduleMonths *int64) *S3AuditGetParams {
	o.SetLogRotationScheduleMonthsQueryParameter(logRotationScheduleMonths)
	return o
}

// SetLogRotationScheduleMonthsQueryParameter adds the logRotationScheduleMonths to the s3 audit get params
func (o *S3AuditGetParams) SetLogRotationScheduleMonthsQueryParameter(logRotationScheduleMonths *int64) {
	o.LogRotationScheduleMonthsQueryParameter = logRotationScheduleMonths
}

// WithLogRotationScheduleWeekdaysQueryParameter adds the logRotationScheduleWeekdays to the s3 audit get params
func (o *S3AuditGetParams) WithLogRotationScheduleWeekdaysQueryParameter(logRotationScheduleWeekdays *int64) *S3AuditGetParams {
	o.SetLogRotationScheduleWeekdaysQueryParameter(logRotationScheduleWeekdays)
	return o
}

// SetLogRotationScheduleWeekdaysQueryParameter adds the logRotationScheduleWeekdays to the s3 audit get params
func (o *S3AuditGetParams) SetLogRotationScheduleWeekdaysQueryParameter(logRotationScheduleWeekdays *int64) {
	o.LogRotationScheduleWeekdaysQueryParameter = logRotationScheduleWeekdays
}

// WithLogRotationSizeQueryParameter adds the logRotationSize to the s3 audit get params
func (o *S3AuditGetParams) WithLogRotationSizeQueryParameter(logRotationSize *int64) *S3AuditGetParams {
	o.SetLogRotationSizeQueryParameter(logRotationSize)
	return o
}

// SetLogRotationSizeQueryParameter adds the logRotationSize to the s3 audit get params
func (o *S3AuditGetParams) SetLogRotationSizeQueryParameter(logRotationSize *int64) {
	o.LogRotationSizeQueryParameter = logRotationSize
}

// WithLogPathQueryParameter adds the logPath to the s3 audit get params
func (o *S3AuditGetParams) WithLogPathQueryParameter(logPath *string) *S3AuditGetParams {
	o.SetLogPathQueryParameter(logPath)
	return o
}

// SetLogPathQueryParameter adds the logPath to the s3 audit get params
func (o *S3AuditGetParams) SetLogPathQueryParameter(logPath *string) {
	o.LogPathQueryParameter = logPath
}

// WithMaxRecordsQueryParameter adds the maxRecords to the s3 audit get params
func (o *S3AuditGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *S3AuditGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the s3 audit get params
func (o *S3AuditGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the s3 audit get params
func (o *S3AuditGetParams) WithOrderByQueryParameter(orderBy []string) *S3AuditGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the s3 audit get params
func (o *S3AuditGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the s3 audit get params
func (o *S3AuditGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *S3AuditGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the s3 audit get params
func (o *S3AuditGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the s3 audit get params
func (o *S3AuditGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *S3AuditGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the s3 audit get params
func (o *S3AuditGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the s3 audit get params
func (o *S3AuditGetParams) WithSVMNameQueryParameter(svmName *string) *S3AuditGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the s3 audit get params
func (o *S3AuditGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDPathParameter adds the svmUUID to the s3 audit get params
func (o *S3AuditGetParams) WithSVMUUIDPathParameter(svmUUID string) *S3AuditGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the s3 audit get params
func (o *S3AuditGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *S3AuditGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.EnabledQueryParameter != nil {

		// query param enabled
		var qrEnabled bool

		if o.EnabledQueryParameter != nil {
			qrEnabled = *o.EnabledQueryParameter
		}
		qEnabled := swag.FormatBool(qrEnabled)
		if qEnabled != "" {

			if err := r.SetQueryParam("enabled", qEnabled); err != nil {
				return err
			}
		}
	}

	if o.EventsDataQueryParameter != nil {

		// query param events.data
		var qrEventsData bool

		if o.EventsDataQueryParameter != nil {
			qrEventsData = *o.EventsDataQueryParameter
		}
		qEventsData := swag.FormatBool(qrEventsData)
		if qEventsData != "" {

			if err := r.SetQueryParam("events.data", qEventsData); err != nil {
				return err
			}
		}
	}

	if o.EventsManagementQueryParameter != nil {

		// query param events.management
		var qrEventsManagement bool

		if o.EventsManagementQueryParameter != nil {
			qrEventsManagement = *o.EventsManagementQueryParameter
		}
		qEventsManagement := swag.FormatBool(qrEventsManagement)
		if qEventsManagement != "" {

			if err := r.SetQueryParam("events.management", qEventsManagement); err != nil {
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

	if o.LogFormatQueryParameter != nil {

		// query param log.format
		var qrLogFormat string

		if o.LogFormatQueryParameter != nil {
			qrLogFormat = *o.LogFormatQueryParameter
		}
		qLogFormat := qrLogFormat
		if qLogFormat != "" {

			if err := r.SetQueryParam("log.format", qLogFormat); err != nil {
				return err
			}
		}
	}

	if o.LogRetentionCountQueryParameter != nil {

		// query param log.retention.count
		var qrLogRetentionCount int64

		if o.LogRetentionCountQueryParameter != nil {
			qrLogRetentionCount = *o.LogRetentionCountQueryParameter
		}
		qLogRetentionCount := swag.FormatInt64(qrLogRetentionCount)
		if qLogRetentionCount != "" {

			if err := r.SetQueryParam("log.retention.count", qLogRetentionCount); err != nil {
				return err
			}
		}
	}

	if o.LogRetentionDurationQueryParameter != nil {

		// query param log.retention.duration
		var qrLogRetentionDuration string

		if o.LogRetentionDurationQueryParameter != nil {
			qrLogRetentionDuration = *o.LogRetentionDurationQueryParameter
		}
		qLogRetentionDuration := qrLogRetentionDuration
		if qLogRetentionDuration != "" {

			if err := r.SetQueryParam("log.retention.duration", qLogRetentionDuration); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleDaysQueryParameter != nil {

		// query param log.rotation.schedule.days
		var qrLogRotationScheduleDays int64

		if o.LogRotationScheduleDaysQueryParameter != nil {
			qrLogRotationScheduleDays = *o.LogRotationScheduleDaysQueryParameter
		}
		qLogRotationScheduleDays := swag.FormatInt64(qrLogRotationScheduleDays)
		if qLogRotationScheduleDays != "" {

			if err := r.SetQueryParam("log.rotation.schedule.days", qLogRotationScheduleDays); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleHoursQueryParameter != nil {

		// query param log.rotation.schedule.hours
		var qrLogRotationScheduleHours int64

		if o.LogRotationScheduleHoursQueryParameter != nil {
			qrLogRotationScheduleHours = *o.LogRotationScheduleHoursQueryParameter
		}
		qLogRotationScheduleHours := swag.FormatInt64(qrLogRotationScheduleHours)
		if qLogRotationScheduleHours != "" {

			if err := r.SetQueryParam("log.rotation.schedule.hours", qLogRotationScheduleHours); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleMinutesQueryParameter != nil {

		// query param log.rotation.schedule.minutes
		var qrLogRotationScheduleMinutes int64

		if o.LogRotationScheduleMinutesQueryParameter != nil {
			qrLogRotationScheduleMinutes = *o.LogRotationScheduleMinutesQueryParameter
		}
		qLogRotationScheduleMinutes := swag.FormatInt64(qrLogRotationScheduleMinutes)
		if qLogRotationScheduleMinutes != "" {

			if err := r.SetQueryParam("log.rotation.schedule.minutes", qLogRotationScheduleMinutes); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleMonthsQueryParameter != nil {

		// query param log.rotation.schedule.months
		var qrLogRotationScheduleMonths int64

		if o.LogRotationScheduleMonthsQueryParameter != nil {
			qrLogRotationScheduleMonths = *o.LogRotationScheduleMonthsQueryParameter
		}
		qLogRotationScheduleMonths := swag.FormatInt64(qrLogRotationScheduleMonths)
		if qLogRotationScheduleMonths != "" {

			if err := r.SetQueryParam("log.rotation.schedule.months", qLogRotationScheduleMonths); err != nil {
				return err
			}
		}
	}

	if o.LogRotationScheduleWeekdaysQueryParameter != nil {

		// query param log.rotation.schedule.weekdays
		var qrLogRotationScheduleWeekdays int64

		if o.LogRotationScheduleWeekdaysQueryParameter != nil {
			qrLogRotationScheduleWeekdays = *o.LogRotationScheduleWeekdaysQueryParameter
		}
		qLogRotationScheduleWeekdays := swag.FormatInt64(qrLogRotationScheduleWeekdays)
		if qLogRotationScheduleWeekdays != "" {

			if err := r.SetQueryParam("log.rotation.schedule.weekdays", qLogRotationScheduleWeekdays); err != nil {
				return err
			}
		}
	}

	if o.LogRotationSizeQueryParameter != nil {

		// query param log.rotation.size
		var qrLogRotationSize int64

		if o.LogRotationSizeQueryParameter != nil {
			qrLogRotationSize = *o.LogRotationSizeQueryParameter
		}
		qLogRotationSize := swag.FormatInt64(qrLogRotationSize)
		if qLogRotationSize != "" {

			if err := r.SetQueryParam("log.rotation.size", qLogRotationSize); err != nil {
				return err
			}
		}
	}

	if o.LogPathQueryParameter != nil {

		// query param log_path
		var qrLogPath string

		if o.LogPathQueryParameter != nil {
			qrLogPath = *o.LogPathQueryParameter
		}
		qLogPath := qrLogPath
		if qLogPath != "" {

			if err := r.SetQueryParam("log_path", qLogPath); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamS3AuditGet binds the parameter fields
func (o *S3AuditGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamS3AuditGet binds the parameter order_by
func (o *S3AuditGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
