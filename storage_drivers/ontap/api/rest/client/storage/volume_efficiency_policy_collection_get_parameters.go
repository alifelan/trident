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

// NewVolumeEfficiencyPolicyCollectionGetParams creates a new VolumeEfficiencyPolicyCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeEfficiencyPolicyCollectionGetParams() *VolumeEfficiencyPolicyCollectionGetParams {
	return &VolumeEfficiencyPolicyCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeEfficiencyPolicyCollectionGetParamsWithTimeout creates a new VolumeEfficiencyPolicyCollectionGetParams object
// with the ability to set a timeout on a request.
func NewVolumeEfficiencyPolicyCollectionGetParamsWithTimeout(timeout time.Duration) *VolumeEfficiencyPolicyCollectionGetParams {
	return &VolumeEfficiencyPolicyCollectionGetParams{
		timeout: timeout,
	}
}

// NewVolumeEfficiencyPolicyCollectionGetParamsWithContext creates a new VolumeEfficiencyPolicyCollectionGetParams object
// with the ability to set a context for a request.
func NewVolumeEfficiencyPolicyCollectionGetParamsWithContext(ctx context.Context) *VolumeEfficiencyPolicyCollectionGetParams {
	return &VolumeEfficiencyPolicyCollectionGetParams{
		Context: ctx,
	}
}

// NewVolumeEfficiencyPolicyCollectionGetParamsWithHTTPClient creates a new VolumeEfficiencyPolicyCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeEfficiencyPolicyCollectionGetParamsWithHTTPClient(client *http.Client) *VolumeEfficiencyPolicyCollectionGetParams {
	return &VolumeEfficiencyPolicyCollectionGetParams{
		HTTPClient: client,
	}
}

/*
VolumeEfficiencyPolicyCollectionGetParams contains all the parameters to send to the API endpoint

	for the volume efficiency policy collection get operation.

	Typically these are written to a http.Request.
*/
type VolumeEfficiencyPolicyCollectionGetParams struct {

	/* Comment.

	   Filter by comment
	*/
	CommentQueryParameter *string

	/* Duration.

	   Filter by duration
	*/
	DurationQueryParameter *int64

	/* Enabled.

	   Filter by enabled
	*/
	EnabledQueryParameter *bool

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

	/* QosPolicy.

	   Filter by qos_policy
	*/
	QosPolicyQueryParameter *string

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

	/* ScheduleName.

	   Filter by schedule.name
	*/
	ScheduleNameQueryParameter *string

	/* StartThresholdPercent.

	   Filter by start_threshold_percent
	*/
	StartThresholdPercentQueryParameter *int64

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* Type.

	   Filter by type
	*/
	TypeQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume efficiency policy collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithDefaults() *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume efficiency policy collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := VolumeEfficiencyPolicyCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithTimeout(timeout time.Duration) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithContext(ctx context.Context) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithHTTPClient(client *http.Client) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommentQueryParameter adds the comment to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithCommentQueryParameter(comment *string) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetCommentQueryParameter(comment)
	return o
}

// SetCommentQueryParameter adds the comment to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetCommentQueryParameter(comment *string) {
	o.CommentQueryParameter = comment
}

// WithDurationQueryParameter adds the duration to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithDurationQueryParameter(duration *int64) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetDurationQueryParameter(duration)
	return o
}

// SetDurationQueryParameter adds the duration to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetDurationQueryParameter(duration *int64) {
	o.DurationQueryParameter = duration
}

// WithEnabledQueryParameter adds the enabled to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithEnabledQueryParameter(enabled *bool) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetEnabledQueryParameter(enabled)
	return o
}

// SetEnabledQueryParameter adds the enabled to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetEnabledQueryParameter(enabled *bool) {
	o.EnabledQueryParameter = enabled
}

// WithFieldsQueryParameter adds the fields to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithFieldsQueryParameter(fields []string) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithNameQueryParameter(name *string) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithQosPolicyQueryParameter adds the qosPolicy to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithQosPolicyQueryParameter(qosPolicy *string) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetQosPolicyQueryParameter(qosPolicy)
	return o
}

// SetQosPolicyQueryParameter adds the qosPolicy to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetQosPolicyQueryParameter(qosPolicy *string) {
	o.QosPolicyQueryParameter = qosPolicy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithScheduleNameQueryParameter adds the scheduleName to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithScheduleNameQueryParameter(scheduleName *string) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetScheduleNameQueryParameter(scheduleName)
	return o
}

// SetScheduleNameQueryParameter adds the scheduleName to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetScheduleNameQueryParameter(scheduleName *string) {
	o.ScheduleNameQueryParameter = scheduleName
}

// WithStartThresholdPercentQueryParameter adds the startThresholdPercent to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithStartThresholdPercentQueryParameter(startThresholdPercent *int64) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetStartThresholdPercentQueryParameter(startThresholdPercent)
	return o
}

// SetStartThresholdPercentQueryParameter adds the startThresholdPercent to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetStartThresholdPercentQueryParameter(startThresholdPercent *int64) {
	o.StartThresholdPercentQueryParameter = startThresholdPercent
}

// WithSVMNameQueryParameter adds the svmName to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithTypeQueryParameter adds the typeVar to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithTypeQueryParameter(typeVar *string) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetTypeQueryParameter(typeVar)
	return o
}

// SetTypeQueryParameter adds the type to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetTypeQueryParameter(typeVar *string) {
	o.TypeQueryParameter = typeVar
}

// WithUUIDQueryParameter adds the uuid to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) WithUUIDQueryParameter(uuid *string) *VolumeEfficiencyPolicyCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the volume efficiency policy collection get params
func (o *VolumeEfficiencyPolicyCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeEfficiencyPolicyCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CommentQueryParameter != nil {

		// query param comment
		var qrComment string

		if o.CommentQueryParameter != nil {
			qrComment = *o.CommentQueryParameter
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if o.DurationQueryParameter != nil {

		// query param duration
		var qrDuration int64

		if o.DurationQueryParameter != nil {
			qrDuration = *o.DurationQueryParameter
		}
		qDuration := swag.FormatInt64(qrDuration)
		if qDuration != "" {

			if err := r.SetQueryParam("duration", qDuration); err != nil {
				return err
			}
		}
	}

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

	if o.QosPolicyQueryParameter != nil {

		// query param qos_policy
		var qrQosPolicy string

		if o.QosPolicyQueryParameter != nil {
			qrQosPolicy = *o.QosPolicyQueryParameter
		}
		qQosPolicy := qrQosPolicy
		if qQosPolicy != "" {

			if err := r.SetQueryParam("qos_policy", qQosPolicy); err != nil {
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

	if o.ScheduleNameQueryParameter != nil {

		// query param schedule.name
		var qrScheduleName string

		if o.ScheduleNameQueryParameter != nil {
			qrScheduleName = *o.ScheduleNameQueryParameter
		}
		qScheduleName := qrScheduleName
		if qScheduleName != "" {

			if err := r.SetQueryParam("schedule.name", qScheduleName); err != nil {
				return err
			}
		}
	}

	if o.StartThresholdPercentQueryParameter != nil {

		// query param start_threshold_percent
		var qrStartThresholdPercent int64

		if o.StartThresholdPercentQueryParameter != nil {
			qrStartThresholdPercent = *o.StartThresholdPercentQueryParameter
		}
		qStartThresholdPercent := swag.FormatInt64(qrStartThresholdPercent)
		if qStartThresholdPercent != "" {

			if err := r.SetQueryParam("start_threshold_percent", qStartThresholdPercent); err != nil {
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

	if o.TypeQueryParameter != nil {

		// query param type
		var qrType string

		if o.TypeQueryParameter != nil {
			qrType = *o.TypeQueryParameter
		}
		qType := qrType
		if qType != "" {

			if err := r.SetQueryParam("type", qType); err != nil {
				return err
			}
		}
	}

	if o.UUIDQueryParameter != nil {

		// query param uuid
		var qrUUID string

		if o.UUIDQueryParameter != nil {
			qrUUID = *o.UUIDQueryParameter
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamVolumeEfficiencyPolicyCollectionGet binds the parameter fields
func (o *VolumeEfficiencyPolicyCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamVolumeEfficiencyPolicyCollectionGet binds the parameter order_by
func (o *VolumeEfficiencyPolicyCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
