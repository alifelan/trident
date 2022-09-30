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

// NewKerberosInterfaceCollectionGetParams creates a new KerberosInterfaceCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewKerberosInterfaceCollectionGetParams() *KerberosInterfaceCollectionGetParams {
	return &KerberosInterfaceCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewKerberosInterfaceCollectionGetParamsWithTimeout creates a new KerberosInterfaceCollectionGetParams object
// with the ability to set a timeout on a request.
func NewKerberosInterfaceCollectionGetParamsWithTimeout(timeout time.Duration) *KerberosInterfaceCollectionGetParams {
	return &KerberosInterfaceCollectionGetParams{
		timeout: timeout,
	}
}

// NewKerberosInterfaceCollectionGetParamsWithContext creates a new KerberosInterfaceCollectionGetParams object
// with the ability to set a context for a request.
func NewKerberosInterfaceCollectionGetParamsWithContext(ctx context.Context) *KerberosInterfaceCollectionGetParams {
	return &KerberosInterfaceCollectionGetParams{
		Context: ctx,
	}
}

// NewKerberosInterfaceCollectionGetParamsWithHTTPClient creates a new KerberosInterfaceCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewKerberosInterfaceCollectionGetParamsWithHTTPClient(client *http.Client) *KerberosInterfaceCollectionGetParams {
	return &KerberosInterfaceCollectionGetParams{
		HTTPClient: client,
	}
}

/*
KerberosInterfaceCollectionGetParams contains all the parameters to send to the API endpoint

	for the kerberos interface collection get operation.

	Typically these are written to a http.Request.
*/
type KerberosInterfaceCollectionGetParams struct {

	/* Enabled.

	   Filter by enabled
	*/
	EnabledQueryParameter *bool

	/* EncryptionTypes.

	   Filter by encryption_types
	*/
	EncryptionTypesQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* InterfaceIPAddress.

	   Filter by interface.ip.address
	*/
	InterfaceIPAddressQueryParameter *string

	/* InterfaceName.

	   Filter by interface.name
	*/
	InterfaceNameQueryParameter *string

	/* InterfaceUUID.

	   Filter by interface.uuid
	*/
	InterfaceUUIDQueryParameter *string

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

	/* Spn.

	   Filter by spn
	*/
	SpnQueryParameter *string

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

// WithDefaults hydrates default values in the kerberos interface collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KerberosInterfaceCollectionGetParams) WithDefaults() *KerberosInterfaceCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the kerberos interface collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *KerberosInterfaceCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := KerberosInterfaceCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithTimeout(timeout time.Duration) *KerberosInterfaceCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithContext(ctx context.Context) *KerberosInterfaceCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithHTTPClient(client *http.Client) *KerberosInterfaceCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEnabledQueryParameter adds the enabled to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithEnabledQueryParameter(enabled *bool) *KerberosInterfaceCollectionGetParams {
	o.SetEnabledQueryParameter(enabled)
	return o
}

// SetEnabledQueryParameter adds the enabled to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetEnabledQueryParameter(enabled *bool) {
	o.EnabledQueryParameter = enabled
}

// WithEncryptionTypesQueryParameter adds the encryptionTypes to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithEncryptionTypesQueryParameter(encryptionTypes *string) *KerberosInterfaceCollectionGetParams {
	o.SetEncryptionTypesQueryParameter(encryptionTypes)
	return o
}

// SetEncryptionTypesQueryParameter adds the encryptionTypes to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetEncryptionTypesQueryParameter(encryptionTypes *string) {
	o.EncryptionTypesQueryParameter = encryptionTypes
}

// WithFieldsQueryParameter adds the fields to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithFieldsQueryParameter(fields []string) *KerberosInterfaceCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithInterfaceIPAddressQueryParameter adds the interfaceIPAddress to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithInterfaceIPAddressQueryParameter(interfaceIPAddress *string) *KerberosInterfaceCollectionGetParams {
	o.SetInterfaceIPAddressQueryParameter(interfaceIPAddress)
	return o
}

// SetInterfaceIPAddressQueryParameter adds the interfaceIpAddress to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetInterfaceIPAddressQueryParameter(interfaceIPAddress *string) {
	o.InterfaceIPAddressQueryParameter = interfaceIPAddress
}

// WithInterfaceNameQueryParameter adds the interfaceName to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithInterfaceNameQueryParameter(interfaceName *string) *KerberosInterfaceCollectionGetParams {
	o.SetInterfaceNameQueryParameter(interfaceName)
	return o
}

// SetInterfaceNameQueryParameter adds the interfaceName to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetInterfaceNameQueryParameter(interfaceName *string) {
	o.InterfaceNameQueryParameter = interfaceName
}

// WithInterfaceUUIDQueryParameter adds the interfaceUUID to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithInterfaceUUIDQueryParameter(interfaceUUID *string) *KerberosInterfaceCollectionGetParams {
	o.SetInterfaceUUIDQueryParameter(interfaceUUID)
	return o
}

// SetInterfaceUUIDQueryParameter adds the interfaceUuid to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetInterfaceUUIDQueryParameter(interfaceUUID *string) {
	o.InterfaceUUIDQueryParameter = interfaceUUID
}

// WithMaxRecordsQueryParameter adds the maxRecords to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *KerberosInterfaceCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *KerberosInterfaceCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *KerberosInterfaceCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *KerberosInterfaceCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSpnQueryParameter adds the spn to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithSpnQueryParameter(spn *string) *KerberosInterfaceCollectionGetParams {
	o.SetSpnQueryParameter(spn)
	return o
}

// SetSpnQueryParameter adds the spn to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetSpnQueryParameter(spn *string) {
	o.SpnQueryParameter = spn
}

// WithSVMNameQueryParameter adds the svmName to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *KerberosInterfaceCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *KerberosInterfaceCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the kerberos interface collection get params
func (o *KerberosInterfaceCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *KerberosInterfaceCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.EncryptionTypesQueryParameter != nil {

		// query param encryption_types
		var qrEncryptionTypes string

		if o.EncryptionTypesQueryParameter != nil {
			qrEncryptionTypes = *o.EncryptionTypesQueryParameter
		}
		qEncryptionTypes := qrEncryptionTypes
		if qEncryptionTypes != "" {

			if err := r.SetQueryParam("encryption_types", qEncryptionTypes); err != nil {
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

	if o.InterfaceIPAddressQueryParameter != nil {

		// query param interface.ip.address
		var qrInterfaceIPAddress string

		if o.InterfaceIPAddressQueryParameter != nil {
			qrInterfaceIPAddress = *o.InterfaceIPAddressQueryParameter
		}
		qInterfaceIPAddress := qrInterfaceIPAddress
		if qInterfaceIPAddress != "" {

			if err := r.SetQueryParam("interface.ip.address", qInterfaceIPAddress); err != nil {
				return err
			}
		}
	}

	if o.InterfaceNameQueryParameter != nil {

		// query param interface.name
		var qrInterfaceName string

		if o.InterfaceNameQueryParameter != nil {
			qrInterfaceName = *o.InterfaceNameQueryParameter
		}
		qInterfaceName := qrInterfaceName
		if qInterfaceName != "" {

			if err := r.SetQueryParam("interface.name", qInterfaceName); err != nil {
				return err
			}
		}
	}

	if o.InterfaceUUIDQueryParameter != nil {

		// query param interface.uuid
		var qrInterfaceUUID string

		if o.InterfaceUUIDQueryParameter != nil {
			qrInterfaceUUID = *o.InterfaceUUIDQueryParameter
		}
		qInterfaceUUID := qrInterfaceUUID
		if qInterfaceUUID != "" {

			if err := r.SetQueryParam("interface.uuid", qInterfaceUUID); err != nil {
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

	if o.SpnQueryParameter != nil {

		// query param spn
		var qrSpn string

		if o.SpnQueryParameter != nil {
			qrSpn = *o.SpnQueryParameter
		}
		qSpn := qrSpn
		if qSpn != "" {

			if err := r.SetQueryParam("spn", qSpn); err != nil {
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

// bindParamKerberosInterfaceCollectionGet binds the parameter fields
func (o *KerberosInterfaceCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamKerberosInterfaceCollectionGet binds the parameter order_by
func (o *KerberosInterfaceCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
