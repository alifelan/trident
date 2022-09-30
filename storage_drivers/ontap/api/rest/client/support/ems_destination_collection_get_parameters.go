// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewEmsDestinationCollectionGetParams creates a new EmsDestinationCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewEmsDestinationCollectionGetParams() *EmsDestinationCollectionGetParams {
	return &EmsDestinationCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEmsDestinationCollectionGetParamsWithTimeout creates a new EmsDestinationCollectionGetParams object
// with the ability to set a timeout on a request.
func NewEmsDestinationCollectionGetParamsWithTimeout(timeout time.Duration) *EmsDestinationCollectionGetParams {
	return &EmsDestinationCollectionGetParams{
		timeout: timeout,
	}
}

// NewEmsDestinationCollectionGetParamsWithContext creates a new EmsDestinationCollectionGetParams object
// with the ability to set a context for a request.
func NewEmsDestinationCollectionGetParamsWithContext(ctx context.Context) *EmsDestinationCollectionGetParams {
	return &EmsDestinationCollectionGetParams{
		Context: ctx,
	}
}

// NewEmsDestinationCollectionGetParamsWithHTTPClient creates a new EmsDestinationCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewEmsDestinationCollectionGetParamsWithHTTPClient(client *http.Client) *EmsDestinationCollectionGetParams {
	return &EmsDestinationCollectionGetParams{
		HTTPClient: client,
	}
}

/*
EmsDestinationCollectionGetParams contains all the parameters to send to the API endpoint

	for the ems destination collection get operation.

	Typically these are written to a http.Request.
*/
type EmsDestinationCollectionGetParams struct {

	/* CertificateCa.

	   Filter by certificate.ca
	*/
	CertificateCaQueryParameter *string

	/* CertificateName.

	   Filter by certificate.name
	*/
	CertificateNameQueryParameter *string

	/* CertificateSerialNumber.

	   Filter by certificate.serial_number
	*/
	CertificateSerialNumberQueryParameter *string

	/* ConnectivityErrorsMessageArgumentsCode.

	   Filter by connectivity.errors.message.arguments.code
	*/
	ConnectivityErrorsMessageArgumentsCodeQueryParameter *string

	/* ConnectivityErrorsMessageArgumentsMessage.

	   Filter by connectivity.errors.message.arguments.message
	*/
	ConnectivityErrorsMessageArgumentsMessageQueryParameter *string

	/* ConnectivityErrorsMessageCode.

	   Filter by connectivity.errors.message.code
	*/
	ConnectivityErrorsMessageCodeQueryParameter *string

	/* ConnectivityErrorsMessageMessage.

	   Filter by connectivity.errors.message.message
	*/
	ConnectivityErrorsMessageMessageQueryParameter *string

	/* ConnectivityErrorsNodeName.

	   Filter by connectivity.errors.node.name
	*/
	ConnectivityErrorsNodeNameQueryParameter *string

	/* ConnectivityErrorsNodeUUID.

	   Filter by connectivity.errors.node.uuid
	*/
	ConnectivityErrorsNodeUUIDQueryParameter *string

	/* ConnectivityState.

	   Filter by connectivity.state
	*/
	ConnectivityStateQueryParameter *string

	/* Destination.

	   Filter by destination
	*/
	DestinationQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* FiltersName.

	   Filter by filters.name
	*/
	FiltersNameQueryParameter *string

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

	/* SystemDefined.

	   Filter by system_defined
	*/
	SystemDefinedQueryParameter *bool

	/* Type.

	   Filter by type
	*/
	TypeQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the ems destination collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsDestinationCollectionGetParams) WithDefaults() *EmsDestinationCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the ems destination collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EmsDestinationCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := EmsDestinationCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithTimeout(timeout time.Duration) *EmsDestinationCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithContext(ctx context.Context) *EmsDestinationCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithHTTPClient(client *http.Client) *EmsDestinationCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCertificateCaQueryParameter adds the certificateCa to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithCertificateCaQueryParameter(certificateCa *string) *EmsDestinationCollectionGetParams {
	o.SetCertificateCaQueryParameter(certificateCa)
	return o
}

// SetCertificateCaQueryParameter adds the certificateCa to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetCertificateCaQueryParameter(certificateCa *string) {
	o.CertificateCaQueryParameter = certificateCa
}

// WithCertificateNameQueryParameter adds the certificateName to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithCertificateNameQueryParameter(certificateName *string) *EmsDestinationCollectionGetParams {
	o.SetCertificateNameQueryParameter(certificateName)
	return o
}

// SetCertificateNameQueryParameter adds the certificateName to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetCertificateNameQueryParameter(certificateName *string) {
	o.CertificateNameQueryParameter = certificateName
}

// WithCertificateSerialNumberQueryParameter adds the certificateSerialNumber to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithCertificateSerialNumberQueryParameter(certificateSerialNumber *string) *EmsDestinationCollectionGetParams {
	o.SetCertificateSerialNumberQueryParameter(certificateSerialNumber)
	return o
}

// SetCertificateSerialNumberQueryParameter adds the certificateSerialNumber to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetCertificateSerialNumberQueryParameter(certificateSerialNumber *string) {
	o.CertificateSerialNumberQueryParameter = certificateSerialNumber
}

// WithConnectivityErrorsMessageArgumentsCodeQueryParameter adds the connectivityErrorsMessageArgumentsCode to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithConnectivityErrorsMessageArgumentsCodeQueryParameter(connectivityErrorsMessageArgumentsCode *string) *EmsDestinationCollectionGetParams {
	o.SetConnectivityErrorsMessageArgumentsCodeQueryParameter(connectivityErrorsMessageArgumentsCode)
	return o
}

// SetConnectivityErrorsMessageArgumentsCodeQueryParameter adds the connectivityErrorsMessageArgumentsCode to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetConnectivityErrorsMessageArgumentsCodeQueryParameter(connectivityErrorsMessageArgumentsCode *string) {
	o.ConnectivityErrorsMessageArgumentsCodeQueryParameter = connectivityErrorsMessageArgumentsCode
}

// WithConnectivityErrorsMessageArgumentsMessageQueryParameter adds the connectivityErrorsMessageArgumentsMessage to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithConnectivityErrorsMessageArgumentsMessageQueryParameter(connectivityErrorsMessageArgumentsMessage *string) *EmsDestinationCollectionGetParams {
	o.SetConnectivityErrorsMessageArgumentsMessageQueryParameter(connectivityErrorsMessageArgumentsMessage)
	return o
}

// SetConnectivityErrorsMessageArgumentsMessageQueryParameter adds the connectivityErrorsMessageArgumentsMessage to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetConnectivityErrorsMessageArgumentsMessageQueryParameter(connectivityErrorsMessageArgumentsMessage *string) {
	o.ConnectivityErrorsMessageArgumentsMessageQueryParameter = connectivityErrorsMessageArgumentsMessage
}

// WithConnectivityErrorsMessageCodeQueryParameter adds the connectivityErrorsMessageCode to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithConnectivityErrorsMessageCodeQueryParameter(connectivityErrorsMessageCode *string) *EmsDestinationCollectionGetParams {
	o.SetConnectivityErrorsMessageCodeQueryParameter(connectivityErrorsMessageCode)
	return o
}

// SetConnectivityErrorsMessageCodeQueryParameter adds the connectivityErrorsMessageCode to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetConnectivityErrorsMessageCodeQueryParameter(connectivityErrorsMessageCode *string) {
	o.ConnectivityErrorsMessageCodeQueryParameter = connectivityErrorsMessageCode
}

// WithConnectivityErrorsMessageMessageQueryParameter adds the connectivityErrorsMessageMessage to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithConnectivityErrorsMessageMessageQueryParameter(connectivityErrorsMessageMessage *string) *EmsDestinationCollectionGetParams {
	o.SetConnectivityErrorsMessageMessageQueryParameter(connectivityErrorsMessageMessage)
	return o
}

// SetConnectivityErrorsMessageMessageQueryParameter adds the connectivityErrorsMessageMessage to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetConnectivityErrorsMessageMessageQueryParameter(connectivityErrorsMessageMessage *string) {
	o.ConnectivityErrorsMessageMessageQueryParameter = connectivityErrorsMessageMessage
}

// WithConnectivityErrorsNodeNameQueryParameter adds the connectivityErrorsNodeName to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithConnectivityErrorsNodeNameQueryParameter(connectivityErrorsNodeName *string) *EmsDestinationCollectionGetParams {
	o.SetConnectivityErrorsNodeNameQueryParameter(connectivityErrorsNodeName)
	return o
}

// SetConnectivityErrorsNodeNameQueryParameter adds the connectivityErrorsNodeName to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetConnectivityErrorsNodeNameQueryParameter(connectivityErrorsNodeName *string) {
	o.ConnectivityErrorsNodeNameQueryParameter = connectivityErrorsNodeName
}

// WithConnectivityErrorsNodeUUIDQueryParameter adds the connectivityErrorsNodeUUID to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithConnectivityErrorsNodeUUIDQueryParameter(connectivityErrorsNodeUUID *string) *EmsDestinationCollectionGetParams {
	o.SetConnectivityErrorsNodeUUIDQueryParameter(connectivityErrorsNodeUUID)
	return o
}

// SetConnectivityErrorsNodeUUIDQueryParameter adds the connectivityErrorsNodeUuid to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetConnectivityErrorsNodeUUIDQueryParameter(connectivityErrorsNodeUUID *string) {
	o.ConnectivityErrorsNodeUUIDQueryParameter = connectivityErrorsNodeUUID
}

// WithConnectivityStateQueryParameter adds the connectivityState to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithConnectivityStateQueryParameter(connectivityState *string) *EmsDestinationCollectionGetParams {
	o.SetConnectivityStateQueryParameter(connectivityState)
	return o
}

// SetConnectivityStateQueryParameter adds the connectivityState to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetConnectivityStateQueryParameter(connectivityState *string) {
	o.ConnectivityStateQueryParameter = connectivityState
}

// WithDestinationQueryParameter adds the destination to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithDestinationQueryParameter(destination *string) *EmsDestinationCollectionGetParams {
	o.SetDestinationQueryParameter(destination)
	return o
}

// SetDestinationQueryParameter adds the destination to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetDestinationQueryParameter(destination *string) {
	o.DestinationQueryParameter = destination
}

// WithFieldsQueryParameter adds the fields to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithFieldsQueryParameter(fields []string) *EmsDestinationCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithFiltersNameQueryParameter adds the filtersName to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithFiltersNameQueryParameter(filtersName *string) *EmsDestinationCollectionGetParams {
	o.SetFiltersNameQueryParameter(filtersName)
	return o
}

// SetFiltersNameQueryParameter adds the filtersName to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetFiltersNameQueryParameter(filtersName *string) {
	o.FiltersNameQueryParameter = filtersName
}

// WithMaxRecordsQueryParameter adds the maxRecords to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *EmsDestinationCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithNameQueryParameter(name *string) *EmsDestinationCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *EmsDestinationCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *EmsDestinationCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *EmsDestinationCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSystemDefinedQueryParameter adds the systemDefined to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithSystemDefinedQueryParameter(systemDefined *bool) *EmsDestinationCollectionGetParams {
	o.SetSystemDefinedQueryParameter(systemDefined)
	return o
}

// SetSystemDefinedQueryParameter adds the systemDefined to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetSystemDefinedQueryParameter(systemDefined *bool) {
	o.SystemDefinedQueryParameter = systemDefined
}

// WithTypeQueryParameter adds the typeVar to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) WithTypeQueryParameter(typeVar *string) *EmsDestinationCollectionGetParams {
	o.SetTypeQueryParameter(typeVar)
	return o
}

// SetTypeQueryParameter adds the type to the ems destination collection get params
func (o *EmsDestinationCollectionGetParams) SetTypeQueryParameter(typeVar *string) {
	o.TypeQueryParameter = typeVar
}

// WriteToRequest writes these params to a swagger request
func (o *EmsDestinationCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CertificateCaQueryParameter != nil {

		// query param certificate.ca
		var qrCertificateCa string

		if o.CertificateCaQueryParameter != nil {
			qrCertificateCa = *o.CertificateCaQueryParameter
		}
		qCertificateCa := qrCertificateCa
		if qCertificateCa != "" {

			if err := r.SetQueryParam("certificate.ca", qCertificateCa); err != nil {
				return err
			}
		}
	}

	if o.CertificateNameQueryParameter != nil {

		// query param certificate.name
		var qrCertificateName string

		if o.CertificateNameQueryParameter != nil {
			qrCertificateName = *o.CertificateNameQueryParameter
		}
		qCertificateName := qrCertificateName
		if qCertificateName != "" {

			if err := r.SetQueryParam("certificate.name", qCertificateName); err != nil {
				return err
			}
		}
	}

	if o.CertificateSerialNumberQueryParameter != nil {

		// query param certificate.serial_number
		var qrCertificateSerialNumber string

		if o.CertificateSerialNumberQueryParameter != nil {
			qrCertificateSerialNumber = *o.CertificateSerialNumberQueryParameter
		}
		qCertificateSerialNumber := qrCertificateSerialNumber
		if qCertificateSerialNumber != "" {

			if err := r.SetQueryParam("certificate.serial_number", qCertificateSerialNumber); err != nil {
				return err
			}
		}
	}

	if o.ConnectivityErrorsMessageArgumentsCodeQueryParameter != nil {

		// query param connectivity.errors.message.arguments.code
		var qrConnectivityErrorsMessageArgumentsCode string

		if o.ConnectivityErrorsMessageArgumentsCodeQueryParameter != nil {
			qrConnectivityErrorsMessageArgumentsCode = *o.ConnectivityErrorsMessageArgumentsCodeQueryParameter
		}
		qConnectivityErrorsMessageArgumentsCode := qrConnectivityErrorsMessageArgumentsCode
		if qConnectivityErrorsMessageArgumentsCode != "" {

			if err := r.SetQueryParam("connectivity.errors.message.arguments.code", qConnectivityErrorsMessageArgumentsCode); err != nil {
				return err
			}
		}
	}

	if o.ConnectivityErrorsMessageArgumentsMessageQueryParameter != nil {

		// query param connectivity.errors.message.arguments.message
		var qrConnectivityErrorsMessageArgumentsMessage string

		if o.ConnectivityErrorsMessageArgumentsMessageQueryParameter != nil {
			qrConnectivityErrorsMessageArgumentsMessage = *o.ConnectivityErrorsMessageArgumentsMessageQueryParameter
		}
		qConnectivityErrorsMessageArgumentsMessage := qrConnectivityErrorsMessageArgumentsMessage
		if qConnectivityErrorsMessageArgumentsMessage != "" {

			if err := r.SetQueryParam("connectivity.errors.message.arguments.message", qConnectivityErrorsMessageArgumentsMessage); err != nil {
				return err
			}
		}
	}

	if o.ConnectivityErrorsMessageCodeQueryParameter != nil {

		// query param connectivity.errors.message.code
		var qrConnectivityErrorsMessageCode string

		if o.ConnectivityErrorsMessageCodeQueryParameter != nil {
			qrConnectivityErrorsMessageCode = *o.ConnectivityErrorsMessageCodeQueryParameter
		}
		qConnectivityErrorsMessageCode := qrConnectivityErrorsMessageCode
		if qConnectivityErrorsMessageCode != "" {

			if err := r.SetQueryParam("connectivity.errors.message.code", qConnectivityErrorsMessageCode); err != nil {
				return err
			}
		}
	}

	if o.ConnectivityErrorsMessageMessageQueryParameter != nil {

		// query param connectivity.errors.message.message
		var qrConnectivityErrorsMessageMessage string

		if o.ConnectivityErrorsMessageMessageQueryParameter != nil {
			qrConnectivityErrorsMessageMessage = *o.ConnectivityErrorsMessageMessageQueryParameter
		}
		qConnectivityErrorsMessageMessage := qrConnectivityErrorsMessageMessage
		if qConnectivityErrorsMessageMessage != "" {

			if err := r.SetQueryParam("connectivity.errors.message.message", qConnectivityErrorsMessageMessage); err != nil {
				return err
			}
		}
	}

	if o.ConnectivityErrorsNodeNameQueryParameter != nil {

		// query param connectivity.errors.node.name
		var qrConnectivityErrorsNodeName string

		if o.ConnectivityErrorsNodeNameQueryParameter != nil {
			qrConnectivityErrorsNodeName = *o.ConnectivityErrorsNodeNameQueryParameter
		}
		qConnectivityErrorsNodeName := qrConnectivityErrorsNodeName
		if qConnectivityErrorsNodeName != "" {

			if err := r.SetQueryParam("connectivity.errors.node.name", qConnectivityErrorsNodeName); err != nil {
				return err
			}
		}
	}

	if o.ConnectivityErrorsNodeUUIDQueryParameter != nil {

		// query param connectivity.errors.node.uuid
		var qrConnectivityErrorsNodeUUID string

		if o.ConnectivityErrorsNodeUUIDQueryParameter != nil {
			qrConnectivityErrorsNodeUUID = *o.ConnectivityErrorsNodeUUIDQueryParameter
		}
		qConnectivityErrorsNodeUUID := qrConnectivityErrorsNodeUUID
		if qConnectivityErrorsNodeUUID != "" {

			if err := r.SetQueryParam("connectivity.errors.node.uuid", qConnectivityErrorsNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.ConnectivityStateQueryParameter != nil {

		// query param connectivity.state
		var qrConnectivityState string

		if o.ConnectivityStateQueryParameter != nil {
			qrConnectivityState = *o.ConnectivityStateQueryParameter
		}
		qConnectivityState := qrConnectivityState
		if qConnectivityState != "" {

			if err := r.SetQueryParam("connectivity.state", qConnectivityState); err != nil {
				return err
			}
		}
	}

	if o.DestinationQueryParameter != nil {

		// query param destination
		var qrDestination string

		if o.DestinationQueryParameter != nil {
			qrDestination = *o.DestinationQueryParameter
		}
		qDestination := qrDestination
		if qDestination != "" {

			if err := r.SetQueryParam("destination", qDestination); err != nil {
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

	if o.FiltersNameQueryParameter != nil {

		// query param filters.name
		var qrFiltersName string

		if o.FiltersNameQueryParameter != nil {
			qrFiltersName = *o.FiltersNameQueryParameter
		}
		qFiltersName := qrFiltersName
		if qFiltersName != "" {

			if err := r.SetQueryParam("filters.name", qFiltersName); err != nil {
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

	if o.SystemDefinedQueryParameter != nil {

		// query param system_defined
		var qrSystemDefined bool

		if o.SystemDefinedQueryParameter != nil {
			qrSystemDefined = *o.SystemDefinedQueryParameter
		}
		qSystemDefined := swag.FormatBool(qrSystemDefined)
		if qSystemDefined != "" {

			if err := r.SetQueryParam("system_defined", qSystemDefined); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamEmsDestinationCollectionGet binds the parameter fields
func (o *EmsDestinationCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamEmsDestinationCollectionGet binds the parameter order_by
func (o *EmsDestinationCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
