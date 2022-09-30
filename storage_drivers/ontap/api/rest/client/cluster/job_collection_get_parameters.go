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

// NewJobCollectionGetParams creates a new JobCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewJobCollectionGetParams() *JobCollectionGetParams {
	return &JobCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewJobCollectionGetParamsWithTimeout creates a new JobCollectionGetParams object
// with the ability to set a timeout on a request.
func NewJobCollectionGetParamsWithTimeout(timeout time.Duration) *JobCollectionGetParams {
	return &JobCollectionGetParams{
		timeout: timeout,
	}
}

// NewJobCollectionGetParamsWithContext creates a new JobCollectionGetParams object
// with the ability to set a context for a request.
func NewJobCollectionGetParamsWithContext(ctx context.Context) *JobCollectionGetParams {
	return &JobCollectionGetParams{
		Context: ctx,
	}
}

// NewJobCollectionGetParamsWithHTTPClient creates a new JobCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewJobCollectionGetParamsWithHTTPClient(client *http.Client) *JobCollectionGetParams {
	return &JobCollectionGetParams{
		HTTPClient: client,
	}
}

/*
JobCollectionGetParams contains all the parameters to send to the API endpoint

	for the job collection get operation.

	Typically these are written to a http.Request.
*/
type JobCollectionGetParams struct {

	/* Code.

	   Filter by code
	*/
	CodeQueryParameter *int64

	/* Description.

	   Filter by description
	*/
	DescriptionQueryParameter *string

	/* EndTime.

	   Filter by end_time
	*/
	EndTimeQueryParameter *string

	/* ErrorArgumentsCode.

	   Filter by error.arguments.code
	*/
	ErrorArgumentsCodeQueryParameter *string

	/* ErrorArgumentsMessage.

	   Filter by error.arguments.message
	*/
	ErrorArgumentsMessageQueryParameter *string

	/* ErrorCode.

	   Filter by error.code
	*/
	ErrorCodeQueryParameter *string

	/* ErrorMessage.

	   Filter by error.message
	*/
	ErrorMessageQueryParameter *string

	/* ErrorTarget.

	   Filter by error.target
	*/
	ErrorTargetQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* Message.

	   Filter by message
	*/
	MessageQueryParameter *string

	/* NodeName.

	   Filter by node.name
	*/
	NodeNameQueryParameter *string

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

	/* StartTime.

	   Filter by start_time
	*/
	StartTimeQueryParameter *string

	/* State.

	   Filter by state
	*/
	StateQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the job collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobCollectionGetParams) WithDefaults() *JobCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the job collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *JobCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := JobCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the job collection get params
func (o *JobCollectionGetParams) WithTimeout(timeout time.Duration) *JobCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the job collection get params
func (o *JobCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the job collection get params
func (o *JobCollectionGetParams) WithContext(ctx context.Context) *JobCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the job collection get params
func (o *JobCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the job collection get params
func (o *JobCollectionGetParams) WithHTTPClient(client *http.Client) *JobCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the job collection get params
func (o *JobCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCodeQueryParameter adds the code to the job collection get params
func (o *JobCollectionGetParams) WithCodeQueryParameter(code *int64) *JobCollectionGetParams {
	o.SetCodeQueryParameter(code)
	return o
}

// SetCodeQueryParameter adds the code to the job collection get params
func (o *JobCollectionGetParams) SetCodeQueryParameter(code *int64) {
	o.CodeQueryParameter = code
}

// WithDescriptionQueryParameter adds the description to the job collection get params
func (o *JobCollectionGetParams) WithDescriptionQueryParameter(description *string) *JobCollectionGetParams {
	o.SetDescriptionQueryParameter(description)
	return o
}

// SetDescriptionQueryParameter adds the description to the job collection get params
func (o *JobCollectionGetParams) SetDescriptionQueryParameter(description *string) {
	o.DescriptionQueryParameter = description
}

// WithEndTimeQueryParameter adds the endTime to the job collection get params
func (o *JobCollectionGetParams) WithEndTimeQueryParameter(endTime *string) *JobCollectionGetParams {
	o.SetEndTimeQueryParameter(endTime)
	return o
}

// SetEndTimeQueryParameter adds the endTime to the job collection get params
func (o *JobCollectionGetParams) SetEndTimeQueryParameter(endTime *string) {
	o.EndTimeQueryParameter = endTime
}

// WithErrorArgumentsCodeQueryParameter adds the errorArgumentsCode to the job collection get params
func (o *JobCollectionGetParams) WithErrorArgumentsCodeQueryParameter(errorArgumentsCode *string) *JobCollectionGetParams {
	o.SetErrorArgumentsCodeQueryParameter(errorArgumentsCode)
	return o
}

// SetErrorArgumentsCodeQueryParameter adds the errorArgumentsCode to the job collection get params
func (o *JobCollectionGetParams) SetErrorArgumentsCodeQueryParameter(errorArgumentsCode *string) {
	o.ErrorArgumentsCodeQueryParameter = errorArgumentsCode
}

// WithErrorArgumentsMessageQueryParameter adds the errorArgumentsMessage to the job collection get params
func (o *JobCollectionGetParams) WithErrorArgumentsMessageQueryParameter(errorArgumentsMessage *string) *JobCollectionGetParams {
	o.SetErrorArgumentsMessageQueryParameter(errorArgumentsMessage)
	return o
}

// SetErrorArgumentsMessageQueryParameter adds the errorArgumentsMessage to the job collection get params
func (o *JobCollectionGetParams) SetErrorArgumentsMessageQueryParameter(errorArgumentsMessage *string) {
	o.ErrorArgumentsMessageQueryParameter = errorArgumentsMessage
}

// WithErrorCodeQueryParameter adds the errorCode to the job collection get params
func (o *JobCollectionGetParams) WithErrorCodeQueryParameter(errorCode *string) *JobCollectionGetParams {
	o.SetErrorCodeQueryParameter(errorCode)
	return o
}

// SetErrorCodeQueryParameter adds the errorCode to the job collection get params
func (o *JobCollectionGetParams) SetErrorCodeQueryParameter(errorCode *string) {
	o.ErrorCodeQueryParameter = errorCode
}

// WithErrorMessageQueryParameter adds the errorMessage to the job collection get params
func (o *JobCollectionGetParams) WithErrorMessageQueryParameter(errorMessage *string) *JobCollectionGetParams {
	o.SetErrorMessageQueryParameter(errorMessage)
	return o
}

// SetErrorMessageQueryParameter adds the errorMessage to the job collection get params
func (o *JobCollectionGetParams) SetErrorMessageQueryParameter(errorMessage *string) {
	o.ErrorMessageQueryParameter = errorMessage
}

// WithErrorTargetQueryParameter adds the errorTarget to the job collection get params
func (o *JobCollectionGetParams) WithErrorTargetQueryParameter(errorTarget *string) *JobCollectionGetParams {
	o.SetErrorTargetQueryParameter(errorTarget)
	return o
}

// SetErrorTargetQueryParameter adds the errorTarget to the job collection get params
func (o *JobCollectionGetParams) SetErrorTargetQueryParameter(errorTarget *string) {
	o.ErrorTargetQueryParameter = errorTarget
}

// WithFieldsQueryParameter adds the fields to the job collection get params
func (o *JobCollectionGetParams) WithFieldsQueryParameter(fields []string) *JobCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the job collection get params
func (o *JobCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the job collection get params
func (o *JobCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *JobCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the job collection get params
func (o *JobCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithMessageQueryParameter adds the message to the job collection get params
func (o *JobCollectionGetParams) WithMessageQueryParameter(message *string) *JobCollectionGetParams {
	o.SetMessageQueryParameter(message)
	return o
}

// SetMessageQueryParameter adds the message to the job collection get params
func (o *JobCollectionGetParams) SetMessageQueryParameter(message *string) {
	o.MessageQueryParameter = message
}

// WithNodeNameQueryParameter adds the nodeName to the job collection get params
func (o *JobCollectionGetParams) WithNodeNameQueryParameter(nodeName *string) *JobCollectionGetParams {
	o.SetNodeNameQueryParameter(nodeName)
	return o
}

// SetNodeNameQueryParameter adds the nodeName to the job collection get params
func (o *JobCollectionGetParams) SetNodeNameQueryParameter(nodeName *string) {
	o.NodeNameQueryParameter = nodeName
}

// WithOrderByQueryParameter adds the orderBy to the job collection get params
func (o *JobCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *JobCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the job collection get params
func (o *JobCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the job collection get params
func (o *JobCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *JobCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the job collection get params
func (o *JobCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the job collection get params
func (o *JobCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *JobCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the job collection get params
func (o *JobCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithStartTimeQueryParameter adds the startTime to the job collection get params
func (o *JobCollectionGetParams) WithStartTimeQueryParameter(startTime *string) *JobCollectionGetParams {
	o.SetStartTimeQueryParameter(startTime)
	return o
}

// SetStartTimeQueryParameter adds the startTime to the job collection get params
func (o *JobCollectionGetParams) SetStartTimeQueryParameter(startTime *string) {
	o.StartTimeQueryParameter = startTime
}

// WithStateQueryParameter adds the state to the job collection get params
func (o *JobCollectionGetParams) WithStateQueryParameter(state *string) *JobCollectionGetParams {
	o.SetStateQueryParameter(state)
	return o
}

// SetStateQueryParameter adds the state to the job collection get params
func (o *JobCollectionGetParams) SetStateQueryParameter(state *string) {
	o.StateQueryParameter = state
}

// WithSVMNameQueryParameter adds the svmName to the job collection get params
func (o *JobCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *JobCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the job collection get params
func (o *JobCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the job collection get params
func (o *JobCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *JobCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the job collection get params
func (o *JobCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithUUIDQueryParameter adds the uuid to the job collection get params
func (o *JobCollectionGetParams) WithUUIDQueryParameter(uuid *string) *JobCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the job collection get params
func (o *JobCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *JobCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CodeQueryParameter != nil {

		// query param code
		var qrCode int64

		if o.CodeQueryParameter != nil {
			qrCode = *o.CodeQueryParameter
		}
		qCode := swag.FormatInt64(qrCode)
		if qCode != "" {

			if err := r.SetQueryParam("code", qCode); err != nil {
				return err
			}
		}
	}

	if o.DescriptionQueryParameter != nil {

		// query param description
		var qrDescription string

		if o.DescriptionQueryParameter != nil {
			qrDescription = *o.DescriptionQueryParameter
		}
		qDescription := qrDescription
		if qDescription != "" {

			if err := r.SetQueryParam("description", qDescription); err != nil {
				return err
			}
		}
	}

	if o.EndTimeQueryParameter != nil {

		// query param end_time
		var qrEndTime string

		if o.EndTimeQueryParameter != nil {
			qrEndTime = *o.EndTimeQueryParameter
		}
		qEndTime := qrEndTime
		if qEndTime != "" {

			if err := r.SetQueryParam("end_time", qEndTime); err != nil {
				return err
			}
		}
	}

	if o.ErrorArgumentsCodeQueryParameter != nil {

		// query param error.arguments.code
		var qrErrorArgumentsCode string

		if o.ErrorArgumentsCodeQueryParameter != nil {
			qrErrorArgumentsCode = *o.ErrorArgumentsCodeQueryParameter
		}
		qErrorArgumentsCode := qrErrorArgumentsCode
		if qErrorArgumentsCode != "" {

			if err := r.SetQueryParam("error.arguments.code", qErrorArgumentsCode); err != nil {
				return err
			}
		}
	}

	if o.ErrorArgumentsMessageQueryParameter != nil {

		// query param error.arguments.message
		var qrErrorArgumentsMessage string

		if o.ErrorArgumentsMessageQueryParameter != nil {
			qrErrorArgumentsMessage = *o.ErrorArgumentsMessageQueryParameter
		}
		qErrorArgumentsMessage := qrErrorArgumentsMessage
		if qErrorArgumentsMessage != "" {

			if err := r.SetQueryParam("error.arguments.message", qErrorArgumentsMessage); err != nil {
				return err
			}
		}
	}

	if o.ErrorCodeQueryParameter != nil {

		// query param error.code
		var qrErrorCode string

		if o.ErrorCodeQueryParameter != nil {
			qrErrorCode = *o.ErrorCodeQueryParameter
		}
		qErrorCode := qrErrorCode
		if qErrorCode != "" {

			if err := r.SetQueryParam("error.code", qErrorCode); err != nil {
				return err
			}
		}
	}

	if o.ErrorMessageQueryParameter != nil {

		// query param error.message
		var qrErrorMessage string

		if o.ErrorMessageQueryParameter != nil {
			qrErrorMessage = *o.ErrorMessageQueryParameter
		}
		qErrorMessage := qrErrorMessage
		if qErrorMessage != "" {

			if err := r.SetQueryParam("error.message", qErrorMessage); err != nil {
				return err
			}
		}
	}

	if o.ErrorTargetQueryParameter != nil {

		// query param error.target
		var qrErrorTarget string

		if o.ErrorTargetQueryParameter != nil {
			qrErrorTarget = *o.ErrorTargetQueryParameter
		}
		qErrorTarget := qrErrorTarget
		if qErrorTarget != "" {

			if err := r.SetQueryParam("error.target", qErrorTarget); err != nil {
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

	if o.MessageQueryParameter != nil {

		// query param message
		var qrMessage string

		if o.MessageQueryParameter != nil {
			qrMessage = *o.MessageQueryParameter
		}
		qMessage := qrMessage
		if qMessage != "" {

			if err := r.SetQueryParam("message", qMessage); err != nil {
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

	if o.StartTimeQueryParameter != nil {

		// query param start_time
		var qrStartTime string

		if o.StartTimeQueryParameter != nil {
			qrStartTime = *o.StartTimeQueryParameter
		}
		qStartTime := qrStartTime
		if qStartTime != "" {

			if err := r.SetQueryParam("start_time", qStartTime); err != nil {
				return err
			}
		}
	}

	if o.StateQueryParameter != nil {

		// query param state
		var qrState string

		if o.StateQueryParameter != nil {
			qrState = *o.StateQueryParameter
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
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

// bindParamJobCollectionGet binds the parameter fields
func (o *JobCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamJobCollectionGet binds the parameter order_by
func (o *JobCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
