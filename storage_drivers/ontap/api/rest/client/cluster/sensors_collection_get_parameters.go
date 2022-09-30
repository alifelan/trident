// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewSensorsCollectionGetParams creates a new SensorsCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSensorsCollectionGetParams() *SensorsCollectionGetParams {
	return &SensorsCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSensorsCollectionGetParamsWithTimeout creates a new SensorsCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSensorsCollectionGetParamsWithTimeout(timeout time.Duration) *SensorsCollectionGetParams {
	return &SensorsCollectionGetParams{
		timeout: timeout,
	}
}

// NewSensorsCollectionGetParamsWithContext creates a new SensorsCollectionGetParams object
// with the ability to set a context for a request.
func NewSensorsCollectionGetParamsWithContext(ctx context.Context) *SensorsCollectionGetParams {
	return &SensorsCollectionGetParams{
		Context: ctx,
	}
}

// NewSensorsCollectionGetParamsWithHTTPClient creates a new SensorsCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSensorsCollectionGetParamsWithHTTPClient(client *http.Client) *SensorsCollectionGetParams {
	return &SensorsCollectionGetParams{
		HTTPClient: client,
	}
}

/*
SensorsCollectionGetParams contains all the parameters to send to the API endpoint

	for the sensors collection get operation.

	Typically these are written to a http.Request.
*/
type SensorsCollectionGetParams struct {

	/* CriticalHighThreshold.

	   Filter by critical_high_threshold
	*/
	CriticalHighThresholdQueryParameter *int64

	/* CriticalLowThreshold.

	   Filter by critical_low_threshold
	*/
	CriticalLowThresholdQueryParameter *int64

	/* DiscreteState.

	   Filter by discrete_state
	*/
	DiscreteStateQueryParameter *string

	/* DiscreteValue.

	   Filter by discrete_value
	*/
	DiscreteValueQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Index.

	   Filter by index
	*/
	IndexQueryParameter *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* NodeName.

	   Filter by node.name
	*/
	NodeNameQueryParameter *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUIDQueryParameter *string

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

	/* ThresholdState.

	   Filter by threshold_state
	*/
	ThresholdStateQueryParameter *string

	/* Type.

	   Filter by type
	*/
	TypeQueryParameter *string

	/* Value.

	   Filter by value
	*/
	ValueQueryParameter *int64

	/* ValueUnits.

	   Filter by value_units
	*/
	ValueUnitsQueryParameter *string

	/* WarningHighThreshold.

	   Filter by warning_high_threshold
	*/
	WarningHighThresholdQueryParameter *int64

	/* WarningLowThreshold.

	   Filter by warning_low_threshold
	*/
	WarningLowThresholdQueryParameter *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the sensors collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SensorsCollectionGetParams) WithDefaults() *SensorsCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the sensors collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SensorsCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := SensorsCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the sensors collection get params
func (o *SensorsCollectionGetParams) WithTimeout(timeout time.Duration) *SensorsCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the sensors collection get params
func (o *SensorsCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the sensors collection get params
func (o *SensorsCollectionGetParams) WithContext(ctx context.Context) *SensorsCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the sensors collection get params
func (o *SensorsCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the sensors collection get params
func (o *SensorsCollectionGetParams) WithHTTPClient(client *http.Client) *SensorsCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the sensors collection get params
func (o *SensorsCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCriticalHighThresholdQueryParameter adds the criticalHighThreshold to the sensors collection get params
func (o *SensorsCollectionGetParams) WithCriticalHighThresholdQueryParameter(criticalHighThreshold *int64) *SensorsCollectionGetParams {
	o.SetCriticalHighThresholdQueryParameter(criticalHighThreshold)
	return o
}

// SetCriticalHighThresholdQueryParameter adds the criticalHighThreshold to the sensors collection get params
func (o *SensorsCollectionGetParams) SetCriticalHighThresholdQueryParameter(criticalHighThreshold *int64) {
	o.CriticalHighThresholdQueryParameter = criticalHighThreshold
}

// WithCriticalLowThresholdQueryParameter adds the criticalLowThreshold to the sensors collection get params
func (o *SensorsCollectionGetParams) WithCriticalLowThresholdQueryParameter(criticalLowThreshold *int64) *SensorsCollectionGetParams {
	o.SetCriticalLowThresholdQueryParameter(criticalLowThreshold)
	return o
}

// SetCriticalLowThresholdQueryParameter adds the criticalLowThreshold to the sensors collection get params
func (o *SensorsCollectionGetParams) SetCriticalLowThresholdQueryParameter(criticalLowThreshold *int64) {
	o.CriticalLowThresholdQueryParameter = criticalLowThreshold
}

// WithDiscreteStateQueryParameter adds the discreteState to the sensors collection get params
func (o *SensorsCollectionGetParams) WithDiscreteStateQueryParameter(discreteState *string) *SensorsCollectionGetParams {
	o.SetDiscreteStateQueryParameter(discreteState)
	return o
}

// SetDiscreteStateQueryParameter adds the discreteState to the sensors collection get params
func (o *SensorsCollectionGetParams) SetDiscreteStateQueryParameter(discreteState *string) {
	o.DiscreteStateQueryParameter = discreteState
}

// WithDiscreteValueQueryParameter adds the discreteValue to the sensors collection get params
func (o *SensorsCollectionGetParams) WithDiscreteValueQueryParameter(discreteValue *string) *SensorsCollectionGetParams {
	o.SetDiscreteValueQueryParameter(discreteValue)
	return o
}

// SetDiscreteValueQueryParameter adds the discreteValue to the sensors collection get params
func (o *SensorsCollectionGetParams) SetDiscreteValueQueryParameter(discreteValue *string) {
	o.DiscreteValueQueryParameter = discreteValue
}

// WithFieldsQueryParameter adds the fields to the sensors collection get params
func (o *SensorsCollectionGetParams) WithFieldsQueryParameter(fields []string) *SensorsCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the sensors collection get params
func (o *SensorsCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIndexQueryParameter adds the index to the sensors collection get params
func (o *SensorsCollectionGetParams) WithIndexQueryParameter(index *int64) *SensorsCollectionGetParams {
	o.SetIndexQueryParameter(index)
	return o
}

// SetIndexQueryParameter adds the index to the sensors collection get params
func (o *SensorsCollectionGetParams) SetIndexQueryParameter(index *int64) {
	o.IndexQueryParameter = index
}

// WithMaxRecordsQueryParameter adds the maxRecords to the sensors collection get params
func (o *SensorsCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *SensorsCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the sensors collection get params
func (o *SensorsCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the sensors collection get params
func (o *SensorsCollectionGetParams) WithNameQueryParameter(name *string) *SensorsCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the sensors collection get params
func (o *SensorsCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithNodeNameQueryParameter adds the nodeName to the sensors collection get params
func (o *SensorsCollectionGetParams) WithNodeNameQueryParameter(nodeName *string) *SensorsCollectionGetParams {
	o.SetNodeNameQueryParameter(nodeName)
	return o
}

// SetNodeNameQueryParameter adds the nodeName to the sensors collection get params
func (o *SensorsCollectionGetParams) SetNodeNameQueryParameter(nodeName *string) {
	o.NodeNameQueryParameter = nodeName
}

// WithNodeUUIDQueryParameter adds the nodeUUID to the sensors collection get params
func (o *SensorsCollectionGetParams) WithNodeUUIDQueryParameter(nodeUUID *string) *SensorsCollectionGetParams {
	o.SetNodeUUIDQueryParameter(nodeUUID)
	return o
}

// SetNodeUUIDQueryParameter adds the nodeUuid to the sensors collection get params
func (o *SensorsCollectionGetParams) SetNodeUUIDQueryParameter(nodeUUID *string) {
	o.NodeUUIDQueryParameter = nodeUUID
}

// WithOrderByQueryParameter adds the orderBy to the sensors collection get params
func (o *SensorsCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *SensorsCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the sensors collection get params
func (o *SensorsCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the sensors collection get params
func (o *SensorsCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *SensorsCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the sensors collection get params
func (o *SensorsCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the sensors collection get params
func (o *SensorsCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SensorsCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the sensors collection get params
func (o *SensorsCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithThresholdStateQueryParameter adds the thresholdState to the sensors collection get params
func (o *SensorsCollectionGetParams) WithThresholdStateQueryParameter(thresholdState *string) *SensorsCollectionGetParams {
	o.SetThresholdStateQueryParameter(thresholdState)
	return o
}

// SetThresholdStateQueryParameter adds the thresholdState to the sensors collection get params
func (o *SensorsCollectionGetParams) SetThresholdStateQueryParameter(thresholdState *string) {
	o.ThresholdStateQueryParameter = thresholdState
}

// WithTypeQueryParameter adds the typeVar to the sensors collection get params
func (o *SensorsCollectionGetParams) WithTypeQueryParameter(typeVar *string) *SensorsCollectionGetParams {
	o.SetTypeQueryParameter(typeVar)
	return o
}

// SetTypeQueryParameter adds the type to the sensors collection get params
func (o *SensorsCollectionGetParams) SetTypeQueryParameter(typeVar *string) {
	o.TypeQueryParameter = typeVar
}

// WithValueQueryParameter adds the value to the sensors collection get params
func (o *SensorsCollectionGetParams) WithValueQueryParameter(value *int64) *SensorsCollectionGetParams {
	o.SetValueQueryParameter(value)
	return o
}

// SetValueQueryParameter adds the value to the sensors collection get params
func (o *SensorsCollectionGetParams) SetValueQueryParameter(value *int64) {
	o.ValueQueryParameter = value
}

// WithValueUnitsQueryParameter adds the valueUnits to the sensors collection get params
func (o *SensorsCollectionGetParams) WithValueUnitsQueryParameter(valueUnits *string) *SensorsCollectionGetParams {
	o.SetValueUnitsQueryParameter(valueUnits)
	return o
}

// SetValueUnitsQueryParameter adds the valueUnits to the sensors collection get params
func (o *SensorsCollectionGetParams) SetValueUnitsQueryParameter(valueUnits *string) {
	o.ValueUnitsQueryParameter = valueUnits
}

// WithWarningHighThresholdQueryParameter adds the warningHighThreshold to the sensors collection get params
func (o *SensorsCollectionGetParams) WithWarningHighThresholdQueryParameter(warningHighThreshold *int64) *SensorsCollectionGetParams {
	o.SetWarningHighThresholdQueryParameter(warningHighThreshold)
	return o
}

// SetWarningHighThresholdQueryParameter adds the warningHighThreshold to the sensors collection get params
func (o *SensorsCollectionGetParams) SetWarningHighThresholdQueryParameter(warningHighThreshold *int64) {
	o.WarningHighThresholdQueryParameter = warningHighThreshold
}

// WithWarningLowThresholdQueryParameter adds the warningLowThreshold to the sensors collection get params
func (o *SensorsCollectionGetParams) WithWarningLowThresholdQueryParameter(warningLowThreshold *int64) *SensorsCollectionGetParams {
	o.SetWarningLowThresholdQueryParameter(warningLowThreshold)
	return o
}

// SetWarningLowThresholdQueryParameter adds the warningLowThreshold to the sensors collection get params
func (o *SensorsCollectionGetParams) SetWarningLowThresholdQueryParameter(warningLowThreshold *int64) {
	o.WarningLowThresholdQueryParameter = warningLowThreshold
}

// WriteToRequest writes these params to a swagger request
func (o *SensorsCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CriticalHighThresholdQueryParameter != nil {

		// query param critical_high_threshold
		var qrCriticalHighThreshold int64

		if o.CriticalHighThresholdQueryParameter != nil {
			qrCriticalHighThreshold = *o.CriticalHighThresholdQueryParameter
		}
		qCriticalHighThreshold := swag.FormatInt64(qrCriticalHighThreshold)
		if qCriticalHighThreshold != "" {

			if err := r.SetQueryParam("critical_high_threshold", qCriticalHighThreshold); err != nil {
				return err
			}
		}
	}

	if o.CriticalLowThresholdQueryParameter != nil {

		// query param critical_low_threshold
		var qrCriticalLowThreshold int64

		if o.CriticalLowThresholdQueryParameter != nil {
			qrCriticalLowThreshold = *o.CriticalLowThresholdQueryParameter
		}
		qCriticalLowThreshold := swag.FormatInt64(qrCriticalLowThreshold)
		if qCriticalLowThreshold != "" {

			if err := r.SetQueryParam("critical_low_threshold", qCriticalLowThreshold); err != nil {
				return err
			}
		}
	}

	if o.DiscreteStateQueryParameter != nil {

		// query param discrete_state
		var qrDiscreteState string

		if o.DiscreteStateQueryParameter != nil {
			qrDiscreteState = *o.DiscreteStateQueryParameter
		}
		qDiscreteState := qrDiscreteState
		if qDiscreteState != "" {

			if err := r.SetQueryParam("discrete_state", qDiscreteState); err != nil {
				return err
			}
		}
	}

	if o.DiscreteValueQueryParameter != nil {

		// query param discrete_value
		var qrDiscreteValue string

		if o.DiscreteValueQueryParameter != nil {
			qrDiscreteValue = *o.DiscreteValueQueryParameter
		}
		qDiscreteValue := qrDiscreteValue
		if qDiscreteValue != "" {

			if err := r.SetQueryParam("discrete_value", qDiscreteValue); err != nil {
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

	if o.IndexQueryParameter != nil {

		// query param index
		var qrIndex int64

		if o.IndexQueryParameter != nil {
			qrIndex = *o.IndexQueryParameter
		}
		qIndex := swag.FormatInt64(qrIndex)
		if qIndex != "" {

			if err := r.SetQueryParam("index", qIndex); err != nil {
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

	if o.NodeNameQueryParameter != nil {

		// query param node.name
		var qrNodeName string

		if o.NodeNameQueryParameter != nil {
			qrNodeName = *o.NodeNameQueryParameter
		}
		qNodeName := qrNodeName
		if qNodeName != "" {

			if err := r.SetQueryParam("node.name", qNodeName); err != nil {
				return err
			}
		}
	}

	if o.NodeUUIDQueryParameter != nil {

		// query param node.uuid
		var qrNodeUUID string

		if o.NodeUUIDQueryParameter != nil {
			qrNodeUUID = *o.NodeUUIDQueryParameter
		}
		qNodeUUID := qrNodeUUID
		if qNodeUUID != "" {

			if err := r.SetQueryParam("node.uuid", qNodeUUID); err != nil {
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

	if o.ThresholdStateQueryParameter != nil {

		// query param threshold_state
		var qrThresholdState string

		if o.ThresholdStateQueryParameter != nil {
			qrThresholdState = *o.ThresholdStateQueryParameter
		}
		qThresholdState := qrThresholdState
		if qThresholdState != "" {

			if err := r.SetQueryParam("threshold_state", qThresholdState); err != nil {
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

	if o.ValueQueryParameter != nil {

		// query param value
		var qrValue int64

		if o.ValueQueryParameter != nil {
			qrValue = *o.ValueQueryParameter
		}
		qValue := swag.FormatInt64(qrValue)
		if qValue != "" {

			if err := r.SetQueryParam("value", qValue); err != nil {
				return err
			}
		}
	}

	if o.ValueUnitsQueryParameter != nil {

		// query param value_units
		var qrValueUnits string

		if o.ValueUnitsQueryParameter != nil {
			qrValueUnits = *o.ValueUnitsQueryParameter
		}
		qValueUnits := qrValueUnits
		if qValueUnits != "" {

			if err := r.SetQueryParam("value_units", qValueUnits); err != nil {
				return err
			}
		}
	}

	if o.WarningHighThresholdQueryParameter != nil {

		// query param warning_high_threshold
		var qrWarningHighThreshold int64

		if o.WarningHighThresholdQueryParameter != nil {
			qrWarningHighThreshold = *o.WarningHighThresholdQueryParameter
		}
		qWarningHighThreshold := swag.FormatInt64(qrWarningHighThreshold)
		if qWarningHighThreshold != "" {

			if err := r.SetQueryParam("warning_high_threshold", qWarningHighThreshold); err != nil {
				return err
			}
		}
	}

	if o.WarningLowThresholdQueryParameter != nil {

		// query param warning_low_threshold
		var qrWarningLowThreshold int64

		if o.WarningLowThresholdQueryParameter != nil {
			qrWarningLowThreshold = *o.WarningLowThresholdQueryParameter
		}
		qWarningLowThreshold := swag.FormatInt64(qrWarningLowThreshold)
		if qWarningLowThreshold != "" {

			if err := r.SetQueryParam("warning_low_threshold", qWarningLowThreshold); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSensorsCollectionGet binds the parameter fields
func (o *SensorsCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSensorsCollectionGet binds the parameter order_by
func (o *SensorsCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
