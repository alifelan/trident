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

// NewHTTPProxyCollectionGetParams creates a new HTTPProxyCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewHTTPProxyCollectionGetParams() *HTTPProxyCollectionGetParams {
	return &HTTPProxyCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewHTTPProxyCollectionGetParamsWithTimeout creates a new HTTPProxyCollectionGetParams object
// with the ability to set a timeout on a request.
func NewHTTPProxyCollectionGetParamsWithTimeout(timeout time.Duration) *HTTPProxyCollectionGetParams {
	return &HTTPProxyCollectionGetParams{
		timeout: timeout,
	}
}

// NewHTTPProxyCollectionGetParamsWithContext creates a new HTTPProxyCollectionGetParams object
// with the ability to set a context for a request.
func NewHTTPProxyCollectionGetParamsWithContext(ctx context.Context) *HTTPProxyCollectionGetParams {
	return &HTTPProxyCollectionGetParams{
		Context: ctx,
	}
}

// NewHTTPProxyCollectionGetParamsWithHTTPClient creates a new HTTPProxyCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewHTTPProxyCollectionGetParamsWithHTTPClient(client *http.Client) *HTTPProxyCollectionGetParams {
	return &HTTPProxyCollectionGetParams{
		HTTPClient: client,
	}
}

/*
HTTPProxyCollectionGetParams contains all the parameters to send to the API endpoint

	for the http proxy collection get operation.

	Typically these are written to a http.Request.
*/
type HTTPProxyCollectionGetParams struct {

	/* AuthenticationEnabled.

	   Filter by authentication_enabled
	*/
	AuthenticationEnabledQueryParameter *bool

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* IpspaceName.

	   Filter by ipspace.name
	*/
	IpspaceNameQueryParameter *string

	/* IpspaceUUID.

	   Filter by ipspace.uuid
	*/
	IpspaceUUIDQueryParameter *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* Port.

	   Filter by port
	*/
	PortQueryParameter *int64

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

	/* Scope.

	   Filter by scope
	*/
	ScopeQueryParameter *string

	/* Server.

	   Filter by server
	*/
	ServerQueryParameter *string

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

// WithDefaults hydrates default values in the http proxy collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HTTPProxyCollectionGetParams) WithDefaults() *HTTPProxyCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the http proxy collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *HTTPProxyCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := HTTPProxyCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithTimeout(timeout time.Duration) *HTTPProxyCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithContext(ctx context.Context) *HTTPProxyCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithHTTPClient(client *http.Client) *HTTPProxyCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthenticationEnabledQueryParameter adds the authenticationEnabled to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithAuthenticationEnabledQueryParameter(authenticationEnabled *bool) *HTTPProxyCollectionGetParams {
	o.SetAuthenticationEnabledQueryParameter(authenticationEnabled)
	return o
}

// SetAuthenticationEnabledQueryParameter adds the authenticationEnabled to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetAuthenticationEnabledQueryParameter(authenticationEnabled *bool) {
	o.AuthenticationEnabledQueryParameter = authenticationEnabled
}

// WithFieldsQueryParameter adds the fields to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithFieldsQueryParameter(fields []string) *HTTPProxyCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIpspaceNameQueryParameter adds the ipspaceName to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithIpspaceNameQueryParameter(ipspaceName *string) *HTTPProxyCollectionGetParams {
	o.SetIpspaceNameQueryParameter(ipspaceName)
	return o
}

// SetIpspaceNameQueryParameter adds the ipspaceName to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetIpspaceNameQueryParameter(ipspaceName *string) {
	o.IpspaceNameQueryParameter = ipspaceName
}

// WithIpspaceUUIDQueryParameter adds the ipspaceUUID to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithIpspaceUUIDQueryParameter(ipspaceUUID *string) *HTTPProxyCollectionGetParams {
	o.SetIpspaceUUIDQueryParameter(ipspaceUUID)
	return o
}

// SetIpspaceUUIDQueryParameter adds the ipspaceUuid to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetIpspaceUUIDQueryParameter(ipspaceUUID *string) {
	o.IpspaceUUIDQueryParameter = ipspaceUUID
}

// WithMaxRecordsQueryParameter adds the maxRecords to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *HTTPProxyCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *HTTPProxyCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithPortQueryParameter adds the port to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithPortQueryParameter(port *int64) *HTTPProxyCollectionGetParams {
	o.SetPortQueryParameter(port)
	return o
}

// SetPortQueryParameter adds the port to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetPortQueryParameter(port *int64) {
	o.PortQueryParameter = port
}

// WithReturnRecordsQueryParameter adds the returnRecords to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *HTTPProxyCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *HTTPProxyCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithScopeQueryParameter adds the scope to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithScopeQueryParameter(scope *string) *HTTPProxyCollectionGetParams {
	o.SetScopeQueryParameter(scope)
	return o
}

// SetScopeQueryParameter adds the scope to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetScopeQueryParameter(scope *string) {
	o.ScopeQueryParameter = scope
}

// WithServerQueryParameter adds the server to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithServerQueryParameter(server *string) *HTTPProxyCollectionGetParams {
	o.SetServerQueryParameter(server)
	return o
}

// SetServerQueryParameter adds the server to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetServerQueryParameter(server *string) {
	o.ServerQueryParameter = server
}

// WithSVMNameQueryParameter adds the svmName to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *HTTPProxyCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *HTTPProxyCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithUUIDQueryParameter adds the uuid to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) WithUUIDQueryParameter(uuid *string) *HTTPProxyCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the http proxy collection get params
func (o *HTTPProxyCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *HTTPProxyCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AuthenticationEnabledQueryParameter != nil {

		// query param authentication_enabled
		var qrAuthenticationEnabled bool

		if o.AuthenticationEnabledQueryParameter != nil {
			qrAuthenticationEnabled = *o.AuthenticationEnabledQueryParameter
		}
		qAuthenticationEnabled := swag.FormatBool(qrAuthenticationEnabled)
		if qAuthenticationEnabled != "" {

			if err := r.SetQueryParam("authentication_enabled", qAuthenticationEnabled); err != nil {
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

	if o.IpspaceNameQueryParameter != nil {

		// query param ipspace.name
		var qrIpspaceName string

		if o.IpspaceNameQueryParameter != nil {
			qrIpspaceName = *o.IpspaceNameQueryParameter
		}
		qIpspaceName := qrIpspaceName
		if qIpspaceName != "" {

			if err := r.SetQueryParam("ipspace.name", qIpspaceName); err != nil {
				return err
			}
		}
	}

	if o.IpspaceUUIDQueryParameter != nil {

		// query param ipspace.uuid
		var qrIpspaceUUID string

		if o.IpspaceUUIDQueryParameter != nil {
			qrIpspaceUUID = *o.IpspaceUUIDQueryParameter
		}
		qIpspaceUUID := qrIpspaceUUID
		if qIpspaceUUID != "" {

			if err := r.SetQueryParam("ipspace.uuid", qIpspaceUUID); err != nil {
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

	if o.PortQueryParameter != nil {

		// query param port
		var qrPort int64

		if o.PortQueryParameter != nil {
			qrPort = *o.PortQueryParameter
		}
		qPort := swag.FormatInt64(qrPort)
		if qPort != "" {

			if err := r.SetQueryParam("port", qPort); err != nil {
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

	if o.ScopeQueryParameter != nil {

		// query param scope
		var qrScope string

		if o.ScopeQueryParameter != nil {
			qrScope = *o.ScopeQueryParameter
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}
	}

	if o.ServerQueryParameter != nil {

		// query param server
		var qrServer string

		if o.ServerQueryParameter != nil {
			qrServer = *o.ServerQueryParameter
		}
		qServer := qrServer
		if qServer != "" {

			if err := r.SetQueryParam("server", qServer); err != nil {
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

// bindParamHTTPProxyCollectionGet binds the parameter fields
func (o *HTTPProxyCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamHTTPProxyCollectionGet binds the parameter order_by
func (o *HTTPProxyCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
