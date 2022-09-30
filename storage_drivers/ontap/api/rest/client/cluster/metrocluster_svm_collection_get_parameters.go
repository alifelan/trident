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

// NewMetroclusterSvmCollectionGetParams creates a new MetroclusterSvmCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewMetroclusterSvmCollectionGetParams() *MetroclusterSvmCollectionGetParams {
	return &MetroclusterSvmCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewMetroclusterSvmCollectionGetParamsWithTimeout creates a new MetroclusterSvmCollectionGetParams object
// with the ability to set a timeout on a request.
func NewMetroclusterSvmCollectionGetParamsWithTimeout(timeout time.Duration) *MetroclusterSvmCollectionGetParams {
	return &MetroclusterSvmCollectionGetParams{
		timeout: timeout,
	}
}

// NewMetroclusterSvmCollectionGetParamsWithContext creates a new MetroclusterSvmCollectionGetParams object
// with the ability to set a context for a request.
func NewMetroclusterSvmCollectionGetParamsWithContext(ctx context.Context) *MetroclusterSvmCollectionGetParams {
	return &MetroclusterSvmCollectionGetParams{
		Context: ctx,
	}
}

// NewMetroclusterSvmCollectionGetParamsWithHTTPClient creates a new MetroclusterSvmCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewMetroclusterSvmCollectionGetParamsWithHTTPClient(client *http.Client) *MetroclusterSvmCollectionGetParams {
	return &MetroclusterSvmCollectionGetParams{
		HTTPClient: client,
	}
}

/*
MetroclusterSvmCollectionGetParams contains all the parameters to send to the API endpoint

	for the metrocluster svm collection get operation.

	Typically these are written to a http.Request.
*/
type MetroclusterSvmCollectionGetParams struct {

	/* ClusterName.

	   Filter by cluster.name
	*/
	ClusterNameQueryParameter *string

	/* ClusterUUID.

	   Filter by cluster.uuid
	*/
	ClusterUUIDQueryParameter *string

	/* ConfigurationState.

	   Filter by configuration_state
	*/
	ConfigurationStateQueryParameter *string

	/* FailedReason.

	   Filter by failed_reason
	*/
	FailedReasonQueryParameter *string

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

	/* PartnerSvmName.

	   Filter by partner_svm.name
	*/
	PartnerSVMNameQueryParameter *string

	/* PartnerSvmUUID.

	   Filter by partner_svm.uuid
	*/
	PartnerSVMUUIDQueryParameter *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the metrocluster svm collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroclusterSvmCollectionGetParams) WithDefaults() *MetroclusterSvmCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the metrocluster svm collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *MetroclusterSvmCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := MetroclusterSvmCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithTimeout(timeout time.Duration) *MetroclusterSvmCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithContext(ctx context.Context) *MetroclusterSvmCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithHTTPClient(client *http.Client) *MetroclusterSvmCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterNameQueryParameter adds the clusterName to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithClusterNameQueryParameter(clusterName *string) *MetroclusterSvmCollectionGetParams {
	o.SetClusterNameQueryParameter(clusterName)
	return o
}

// SetClusterNameQueryParameter adds the clusterName to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetClusterNameQueryParameter(clusterName *string) {
	o.ClusterNameQueryParameter = clusterName
}

// WithClusterUUIDQueryParameter adds the clusterUUID to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithClusterUUIDQueryParameter(clusterUUID *string) *MetroclusterSvmCollectionGetParams {
	o.SetClusterUUIDQueryParameter(clusterUUID)
	return o
}

// SetClusterUUIDQueryParameter adds the clusterUuid to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetClusterUUIDQueryParameter(clusterUUID *string) {
	o.ClusterUUIDQueryParameter = clusterUUID
}

// WithConfigurationStateQueryParameter adds the configurationState to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithConfigurationStateQueryParameter(configurationState *string) *MetroclusterSvmCollectionGetParams {
	o.SetConfigurationStateQueryParameter(configurationState)
	return o
}

// SetConfigurationStateQueryParameter adds the configurationState to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetConfigurationStateQueryParameter(configurationState *string) {
	o.ConfigurationStateQueryParameter = configurationState
}

// WithFailedReasonQueryParameter adds the failedReason to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithFailedReasonQueryParameter(failedReason *string) *MetroclusterSvmCollectionGetParams {
	o.SetFailedReasonQueryParameter(failedReason)
	return o
}

// SetFailedReasonQueryParameter adds the failedReason to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetFailedReasonQueryParameter(failedReason *string) {
	o.FailedReasonQueryParameter = failedReason
}

// WithFieldsQueryParameter adds the fields to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithFieldsQueryParameter(fields []string) *MetroclusterSvmCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *MetroclusterSvmCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *MetroclusterSvmCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithPartnerSVMNameQueryParameter adds the partnerSvmName to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithPartnerSVMNameQueryParameter(partnerSvmName *string) *MetroclusterSvmCollectionGetParams {
	o.SetPartnerSVMNameQueryParameter(partnerSvmName)
	return o
}

// SetPartnerSVMNameQueryParameter adds the partnerSvmName to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetPartnerSVMNameQueryParameter(partnerSvmName *string) {
	o.PartnerSVMNameQueryParameter = partnerSvmName
}

// WithPartnerSVMUUIDQueryParameter adds the partnerSvmUUID to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithPartnerSVMUUIDQueryParameter(partnerSvmUUID *string) *MetroclusterSvmCollectionGetParams {
	o.SetPartnerSVMUUIDQueryParameter(partnerSvmUUID)
	return o
}

// SetPartnerSVMUUIDQueryParameter adds the partnerSvmUuid to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetPartnerSVMUUIDQueryParameter(partnerSvmUUID *string) {
	o.PartnerSVMUUIDQueryParameter = partnerSvmUUID
}

// WithReturnRecordsQueryParameter adds the returnRecords to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *MetroclusterSvmCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *MetroclusterSvmCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMNameQueryParameter adds the svmName to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *MetroclusterSvmCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *MetroclusterSvmCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the metrocluster svm collection get params
func (o *MetroclusterSvmCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *MetroclusterSvmCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ClusterNameQueryParameter != nil {

		// query param cluster.name
		var qrClusterName string

		if o.ClusterNameQueryParameter != nil {
			qrClusterName = *o.ClusterNameQueryParameter
		}
		qClusterName := qrClusterName
		if qClusterName != "" {

			if err := r.SetQueryParam("cluster.name", qClusterName); err != nil {
				return err
			}
		}
	}

	if o.ClusterUUIDQueryParameter != nil {

		// query param cluster.uuid
		var qrClusterUUID string

		if o.ClusterUUIDQueryParameter != nil {
			qrClusterUUID = *o.ClusterUUIDQueryParameter
		}
		qClusterUUID := qrClusterUUID
		if qClusterUUID != "" {

			if err := r.SetQueryParam("cluster.uuid", qClusterUUID); err != nil {
				return err
			}
		}
	}

	if o.ConfigurationStateQueryParameter != nil {

		// query param configuration_state
		var qrConfigurationState string

		if o.ConfigurationStateQueryParameter != nil {
			qrConfigurationState = *o.ConfigurationStateQueryParameter
		}
		qConfigurationState := qrConfigurationState
		if qConfigurationState != "" {

			if err := r.SetQueryParam("configuration_state", qConfigurationState); err != nil {
				return err
			}
		}
	}

	if o.FailedReasonQueryParameter != nil {

		// query param failed_reason
		var qrFailedReason string

		if o.FailedReasonQueryParameter != nil {
			qrFailedReason = *o.FailedReasonQueryParameter
		}
		qFailedReason := qrFailedReason
		if qFailedReason != "" {

			if err := r.SetQueryParam("failed_reason", qFailedReason); err != nil {
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

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.PartnerSVMNameQueryParameter != nil {

		// query param partner_svm.name
		var qrPartnerSvmName string

		if o.PartnerSVMNameQueryParameter != nil {
			qrPartnerSvmName = *o.PartnerSVMNameQueryParameter
		}
		qPartnerSvmName := qrPartnerSvmName
		if qPartnerSvmName != "" {

			if err := r.SetQueryParam("partner_svm.name", qPartnerSvmName); err != nil {
				return err
			}
		}
	}

	if o.PartnerSVMUUIDQueryParameter != nil {

		// query param partner_svm.uuid
		var qrPartnerSvmUUID string

		if o.PartnerSVMUUIDQueryParameter != nil {
			qrPartnerSvmUUID = *o.PartnerSVMUUIDQueryParameter
		}
		qPartnerSvmUUID := qrPartnerSvmUUID
		if qPartnerSvmUUID != "" {

			if err := r.SetQueryParam("partner_svm.uuid", qPartnerSvmUUID); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamMetroclusterSvmCollectionGet binds the parameter fields
func (o *MetroclusterSvmCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamMetroclusterSvmCollectionGet binds the parameter order_by
func (o *MetroclusterSvmCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
