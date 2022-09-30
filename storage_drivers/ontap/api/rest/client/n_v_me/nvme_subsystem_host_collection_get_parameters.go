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

// NewNvmeSubsystemHostCollectionGetParams creates a new NvmeSubsystemHostCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNvmeSubsystemHostCollectionGetParams() *NvmeSubsystemHostCollectionGetParams {
	return &NvmeSubsystemHostCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNvmeSubsystemHostCollectionGetParamsWithTimeout creates a new NvmeSubsystemHostCollectionGetParams object
// with the ability to set a timeout on a request.
func NewNvmeSubsystemHostCollectionGetParamsWithTimeout(timeout time.Duration) *NvmeSubsystemHostCollectionGetParams {
	return &NvmeSubsystemHostCollectionGetParams{
		timeout: timeout,
	}
}

// NewNvmeSubsystemHostCollectionGetParamsWithContext creates a new NvmeSubsystemHostCollectionGetParams object
// with the ability to set a context for a request.
func NewNvmeSubsystemHostCollectionGetParamsWithContext(ctx context.Context) *NvmeSubsystemHostCollectionGetParams {
	return &NvmeSubsystemHostCollectionGetParams{
		Context: ctx,
	}
}

// NewNvmeSubsystemHostCollectionGetParamsWithHTTPClient creates a new NvmeSubsystemHostCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNvmeSubsystemHostCollectionGetParamsWithHTTPClient(client *http.Client) *NvmeSubsystemHostCollectionGetParams {
	return &NvmeSubsystemHostCollectionGetParams{
		HTTPClient: client,
	}
}

/*
NvmeSubsystemHostCollectionGetParams contains all the parameters to send to the API endpoint

	for the nvme subsystem host collection get operation.

	Typically these are written to a http.Request.
*/
type NvmeSubsystemHostCollectionGetParams struct {

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

	/* SubsystemUUID.

	   The unique identifier of the NVMe subsystem.

	*/
	SubsystemUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nvme subsystem host collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemHostCollectionGetParams) WithDefaults() *NvmeSubsystemHostCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nvme subsystem host collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemHostCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := NvmeSubsystemHostCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithTimeout(timeout time.Duration) *NvmeSubsystemHostCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithContext(ctx context.Context) *NvmeSubsystemHostCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithHTTPClient(client *http.Client) *NvmeSubsystemHostCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithFieldsQueryParameter(fields []string) *NvmeSubsystemHostCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *NvmeSubsystemHostCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *NvmeSubsystemHostCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *NvmeSubsystemHostCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *NvmeSubsystemHostCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSubsystemUUIDPathParameter adds the subsystemUUID to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) WithSubsystemUUIDPathParameter(subsystemUUID string) *NvmeSubsystemHostCollectionGetParams {
	o.SetSubsystemUUIDPathParameter(subsystemUUID)
	return o
}

// SetSubsystemUUIDPathParameter adds the subsystemUuid to the nvme subsystem host collection get params
func (o *NvmeSubsystemHostCollectionGetParams) SetSubsystemUUIDPathParameter(subsystemUUID string) {
	o.SubsystemUUIDPathParameter = subsystemUUID
}

// WriteToRequest writes these params to a swagger request
func (o *NvmeSubsystemHostCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param subsystem.uuid
	if err := r.SetPathParam("subsystem.uuid", o.SubsystemUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNvmeSubsystemHostCollectionGet binds the parameter fields
func (o *NvmeSubsystemHostCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamNvmeSubsystemHostCollectionGet binds the parameter order_by
func (o *NvmeSubsystemHostCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
