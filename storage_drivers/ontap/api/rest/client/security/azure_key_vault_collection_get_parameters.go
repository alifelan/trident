// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewAzureKeyVaultCollectionGetParams creates a new AzureKeyVaultCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAzureKeyVaultCollectionGetParams() *AzureKeyVaultCollectionGetParams {
	return &AzureKeyVaultCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAzureKeyVaultCollectionGetParamsWithTimeout creates a new AzureKeyVaultCollectionGetParams object
// with the ability to set a timeout on a request.
func NewAzureKeyVaultCollectionGetParamsWithTimeout(timeout time.Duration) *AzureKeyVaultCollectionGetParams {
	return &AzureKeyVaultCollectionGetParams{
		timeout: timeout,
	}
}

// NewAzureKeyVaultCollectionGetParamsWithContext creates a new AzureKeyVaultCollectionGetParams object
// with the ability to set a context for a request.
func NewAzureKeyVaultCollectionGetParamsWithContext(ctx context.Context) *AzureKeyVaultCollectionGetParams {
	return &AzureKeyVaultCollectionGetParams{
		Context: ctx,
	}
}

// NewAzureKeyVaultCollectionGetParamsWithHTTPClient creates a new AzureKeyVaultCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAzureKeyVaultCollectionGetParamsWithHTTPClient(client *http.Client) *AzureKeyVaultCollectionGetParams {
	return &AzureKeyVaultCollectionGetParams{
		HTTPClient: client,
	}
}

/*
AzureKeyVaultCollectionGetParams contains all the parameters to send to the API endpoint

	for the azure key vault collection get operation.

	Typically these are written to a http.Request.
*/
type AzureKeyVaultCollectionGetParams struct {

	/* AuthenticationMethod.

	   Filter by authentication_method
	*/
	AuthenticationMethodQueryParameter *string

	/* AzureReachabilityCode.

	   Filter by azure_reachability.code
	*/
	AzureReachabilityCodeQueryParameter *int64

	/* AzureReachabilityMessage.

	   Filter by azure_reachability.message
	*/
	AzureReachabilityMessageQueryParameter *string

	/* AzureReachabilityReachable.

	   Filter by azure_reachability.reachable
	*/
	AzureReachabilityReachableQueryParameter *bool

	/* ClientID.

	   Filter by client_id
	*/
	ClientIDQueryParameter *string

	/* EkmipReachabilityCode.

	   Filter by ekmip_reachability.code
	*/
	EkmIPReachabilityCodeQueryParameter *int64

	/* EkmipReachabilityMessage.

	   Filter by ekmip_reachability.message
	*/
	EkmIPReachabilityMessageQueryParameter *string

	/* EkmipReachabilityNodeName.

	   Filter by ekmip_reachability.node.name
	*/
	EkmIPReachabilityNodeNameQueryParameter *string

	/* EkmipReachabilityNodeUUID.

	   Filter by ekmip_reachability.node.uuid
	*/
	EkmIPReachabilityNodeUUIDQueryParameter *string

	/* EkmipReachabilityReachable.

	   Filter by ekmip_reachability.reachable
	*/
	EkmIPReachabilityReachableQueryParameter *bool

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* KeyID.

	   Filter by key_id
	*/
	KeyIDQueryParameter *string

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

	/* ProxyHost.

	   Filter by proxy_host
	*/
	ProxyHostQueryParameter *string

	/* ProxyPort.

	   Filter by proxy_port
	*/
	ProxyPortQueryParameter *int64

	/* ProxyType.

	   Filter by proxy_type
	*/
	ProxyTypeQueryParameter *string

	/* ProxyUsername.

	   Filter by proxy_username
	*/
	ProxyUsernameQueryParameter *string

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

	/* StateAvailable.

	   Filter by state.available
	*/
	StateAvailableQueryParameter *bool

	/* StateCode.

	   Filter by state.code
	*/
	StateCodeQueryParameter *int64

	/* StateMessage.

	   Filter by state.message
	*/
	StateMessageQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* TenantID.

	   Filter by tenant_id
	*/
	TenantIDQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the azure key vault collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureKeyVaultCollectionGetParams) WithDefaults() *AzureKeyVaultCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the azure key vault collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AzureKeyVaultCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := AzureKeyVaultCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithTimeout(timeout time.Duration) *AzureKeyVaultCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithContext(ctx context.Context) *AzureKeyVaultCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithHTTPClient(client *http.Client) *AzureKeyVaultCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthenticationMethodQueryParameter adds the authenticationMethod to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithAuthenticationMethodQueryParameter(authenticationMethod *string) *AzureKeyVaultCollectionGetParams {
	o.SetAuthenticationMethodQueryParameter(authenticationMethod)
	return o
}

// SetAuthenticationMethodQueryParameter adds the authenticationMethod to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetAuthenticationMethodQueryParameter(authenticationMethod *string) {
	o.AuthenticationMethodQueryParameter = authenticationMethod
}

// WithAzureReachabilityCodeQueryParameter adds the azureReachabilityCode to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithAzureReachabilityCodeQueryParameter(azureReachabilityCode *int64) *AzureKeyVaultCollectionGetParams {
	o.SetAzureReachabilityCodeQueryParameter(azureReachabilityCode)
	return o
}

// SetAzureReachabilityCodeQueryParameter adds the azureReachabilityCode to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetAzureReachabilityCodeQueryParameter(azureReachabilityCode *int64) {
	o.AzureReachabilityCodeQueryParameter = azureReachabilityCode
}

// WithAzureReachabilityMessageQueryParameter adds the azureReachabilityMessage to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithAzureReachabilityMessageQueryParameter(azureReachabilityMessage *string) *AzureKeyVaultCollectionGetParams {
	o.SetAzureReachabilityMessageQueryParameter(azureReachabilityMessage)
	return o
}

// SetAzureReachabilityMessageQueryParameter adds the azureReachabilityMessage to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetAzureReachabilityMessageQueryParameter(azureReachabilityMessage *string) {
	o.AzureReachabilityMessageQueryParameter = azureReachabilityMessage
}

// WithAzureReachabilityReachableQueryParameter adds the azureReachabilityReachable to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithAzureReachabilityReachableQueryParameter(azureReachabilityReachable *bool) *AzureKeyVaultCollectionGetParams {
	o.SetAzureReachabilityReachableQueryParameter(azureReachabilityReachable)
	return o
}

// SetAzureReachabilityReachableQueryParameter adds the azureReachabilityReachable to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetAzureReachabilityReachableQueryParameter(azureReachabilityReachable *bool) {
	o.AzureReachabilityReachableQueryParameter = azureReachabilityReachable
}

// WithClientIDQueryParameter adds the clientID to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithClientIDQueryParameter(clientID *string) *AzureKeyVaultCollectionGetParams {
	o.SetClientIDQueryParameter(clientID)
	return o
}

// SetClientIDQueryParameter adds the clientId to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetClientIDQueryParameter(clientID *string) {
	o.ClientIDQueryParameter = clientID
}

// WithEkmIPReachabilityCodeQueryParameter adds the ekmipReachabilityCode to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithEkmIPReachabilityCodeQueryParameter(ekmipReachabilityCode *int64) *AzureKeyVaultCollectionGetParams {
	o.SetEkmIPReachabilityCodeQueryParameter(ekmipReachabilityCode)
	return o
}

// SetEkmIPReachabilityCodeQueryParameter adds the ekmipReachabilityCode to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetEkmIPReachabilityCodeQueryParameter(ekmipReachabilityCode *int64) {
	o.EkmIPReachabilityCodeQueryParameter = ekmipReachabilityCode
}

// WithEkmIPReachabilityMessageQueryParameter adds the ekmipReachabilityMessage to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithEkmIPReachabilityMessageQueryParameter(ekmipReachabilityMessage *string) *AzureKeyVaultCollectionGetParams {
	o.SetEkmIPReachabilityMessageQueryParameter(ekmipReachabilityMessage)
	return o
}

// SetEkmIPReachabilityMessageQueryParameter adds the ekmipReachabilityMessage to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetEkmIPReachabilityMessageQueryParameter(ekmipReachabilityMessage *string) {
	o.EkmIPReachabilityMessageQueryParameter = ekmipReachabilityMessage
}

// WithEkmIPReachabilityNodeNameQueryParameter adds the ekmipReachabilityNodeName to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithEkmIPReachabilityNodeNameQueryParameter(ekmipReachabilityNodeName *string) *AzureKeyVaultCollectionGetParams {
	o.SetEkmIPReachabilityNodeNameQueryParameter(ekmipReachabilityNodeName)
	return o
}

// SetEkmIPReachabilityNodeNameQueryParameter adds the ekmipReachabilityNodeName to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetEkmIPReachabilityNodeNameQueryParameter(ekmipReachabilityNodeName *string) {
	o.EkmIPReachabilityNodeNameQueryParameter = ekmipReachabilityNodeName
}

// WithEkmIPReachabilityNodeUUIDQueryParameter adds the ekmipReachabilityNodeUUID to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithEkmIPReachabilityNodeUUIDQueryParameter(ekmipReachabilityNodeUUID *string) *AzureKeyVaultCollectionGetParams {
	o.SetEkmIPReachabilityNodeUUIDQueryParameter(ekmipReachabilityNodeUUID)
	return o
}

// SetEkmIPReachabilityNodeUUIDQueryParameter adds the ekmipReachabilityNodeUuid to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetEkmIPReachabilityNodeUUIDQueryParameter(ekmipReachabilityNodeUUID *string) {
	o.EkmIPReachabilityNodeUUIDQueryParameter = ekmipReachabilityNodeUUID
}

// WithEkmIPReachabilityReachableQueryParameter adds the ekmipReachabilityReachable to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithEkmIPReachabilityReachableQueryParameter(ekmipReachabilityReachable *bool) *AzureKeyVaultCollectionGetParams {
	o.SetEkmIPReachabilityReachableQueryParameter(ekmipReachabilityReachable)
	return o
}

// SetEkmIPReachabilityReachableQueryParameter adds the ekmipReachabilityReachable to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetEkmIPReachabilityReachableQueryParameter(ekmipReachabilityReachable *bool) {
	o.EkmIPReachabilityReachableQueryParameter = ekmipReachabilityReachable
}

// WithFieldsQueryParameter adds the fields to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithFieldsQueryParameter(fields []string) *AzureKeyVaultCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithKeyIDQueryParameter adds the keyID to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithKeyIDQueryParameter(keyID *string) *AzureKeyVaultCollectionGetParams {
	o.SetKeyIDQueryParameter(keyID)
	return o
}

// SetKeyIDQueryParameter adds the keyId to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetKeyIDQueryParameter(keyID *string) {
	o.KeyIDQueryParameter = keyID
}

// WithMaxRecordsQueryParameter adds the maxRecords to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *AzureKeyVaultCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithNameQueryParameter(name *string) *AzureKeyVaultCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *AzureKeyVaultCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithProxyHostQueryParameter adds the proxyHost to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithProxyHostQueryParameter(proxyHost *string) *AzureKeyVaultCollectionGetParams {
	o.SetProxyHostQueryParameter(proxyHost)
	return o
}

// SetProxyHostQueryParameter adds the proxyHost to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetProxyHostQueryParameter(proxyHost *string) {
	o.ProxyHostQueryParameter = proxyHost
}

// WithProxyPortQueryParameter adds the proxyPort to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithProxyPortQueryParameter(proxyPort *int64) *AzureKeyVaultCollectionGetParams {
	o.SetProxyPortQueryParameter(proxyPort)
	return o
}

// SetProxyPortQueryParameter adds the proxyPort to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetProxyPortQueryParameter(proxyPort *int64) {
	o.ProxyPortQueryParameter = proxyPort
}

// WithProxyTypeQueryParameter adds the proxyType to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithProxyTypeQueryParameter(proxyType *string) *AzureKeyVaultCollectionGetParams {
	o.SetProxyTypeQueryParameter(proxyType)
	return o
}

// SetProxyTypeQueryParameter adds the proxyType to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetProxyTypeQueryParameter(proxyType *string) {
	o.ProxyTypeQueryParameter = proxyType
}

// WithProxyUsernameQueryParameter adds the proxyUsername to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithProxyUsernameQueryParameter(proxyUsername *string) *AzureKeyVaultCollectionGetParams {
	o.SetProxyUsernameQueryParameter(proxyUsername)
	return o
}

// SetProxyUsernameQueryParameter adds the proxyUsername to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetProxyUsernameQueryParameter(proxyUsername *string) {
	o.ProxyUsernameQueryParameter = proxyUsername
}

// WithReturnRecordsQueryParameter adds the returnRecords to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *AzureKeyVaultCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *AzureKeyVaultCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithScopeQueryParameter adds the scope to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithScopeQueryParameter(scope *string) *AzureKeyVaultCollectionGetParams {
	o.SetScopeQueryParameter(scope)
	return o
}

// SetScopeQueryParameter adds the scope to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetScopeQueryParameter(scope *string) {
	o.ScopeQueryParameter = scope
}

// WithStateAvailableQueryParameter adds the stateAvailable to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithStateAvailableQueryParameter(stateAvailable *bool) *AzureKeyVaultCollectionGetParams {
	o.SetStateAvailableQueryParameter(stateAvailable)
	return o
}

// SetStateAvailableQueryParameter adds the stateAvailable to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetStateAvailableQueryParameter(stateAvailable *bool) {
	o.StateAvailableQueryParameter = stateAvailable
}

// WithStateCodeQueryParameter adds the stateCode to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithStateCodeQueryParameter(stateCode *int64) *AzureKeyVaultCollectionGetParams {
	o.SetStateCodeQueryParameter(stateCode)
	return o
}

// SetStateCodeQueryParameter adds the stateCode to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetStateCodeQueryParameter(stateCode *int64) {
	o.StateCodeQueryParameter = stateCode
}

// WithStateMessageQueryParameter adds the stateMessage to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithStateMessageQueryParameter(stateMessage *string) *AzureKeyVaultCollectionGetParams {
	o.SetStateMessageQueryParameter(stateMessage)
	return o
}

// SetStateMessageQueryParameter adds the stateMessage to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetStateMessageQueryParameter(stateMessage *string) {
	o.StateMessageQueryParameter = stateMessage
}

// WithSVMNameQueryParameter adds the svmName to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *AzureKeyVaultCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *AzureKeyVaultCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithTenantIDQueryParameter adds the tenantID to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithTenantIDQueryParameter(tenantID *string) *AzureKeyVaultCollectionGetParams {
	o.SetTenantIDQueryParameter(tenantID)
	return o
}

// SetTenantIDQueryParameter adds the tenantId to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetTenantIDQueryParameter(tenantID *string) {
	o.TenantIDQueryParameter = tenantID
}

// WithUUIDQueryParameter adds the uuid to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) WithUUIDQueryParameter(uuid *string) *AzureKeyVaultCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the azure key vault collection get params
func (o *AzureKeyVaultCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *AzureKeyVaultCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AuthenticationMethodQueryParameter != nil {

		// query param authentication_method
		var qrAuthenticationMethod string

		if o.AuthenticationMethodQueryParameter != nil {
			qrAuthenticationMethod = *o.AuthenticationMethodQueryParameter
		}
		qAuthenticationMethod := qrAuthenticationMethod
		if qAuthenticationMethod != "" {

			if err := r.SetQueryParam("authentication_method", qAuthenticationMethod); err != nil {
				return err
			}
		}
	}

	if o.AzureReachabilityCodeQueryParameter != nil {

		// query param azure_reachability.code
		var qrAzureReachabilityCode int64

		if o.AzureReachabilityCodeQueryParameter != nil {
			qrAzureReachabilityCode = *o.AzureReachabilityCodeQueryParameter
		}
		qAzureReachabilityCode := swag.FormatInt64(qrAzureReachabilityCode)
		if qAzureReachabilityCode != "" {

			if err := r.SetQueryParam("azure_reachability.code", qAzureReachabilityCode); err != nil {
				return err
			}
		}
	}

	if o.AzureReachabilityMessageQueryParameter != nil {

		// query param azure_reachability.message
		var qrAzureReachabilityMessage string

		if o.AzureReachabilityMessageQueryParameter != nil {
			qrAzureReachabilityMessage = *o.AzureReachabilityMessageQueryParameter
		}
		qAzureReachabilityMessage := qrAzureReachabilityMessage
		if qAzureReachabilityMessage != "" {

			if err := r.SetQueryParam("azure_reachability.message", qAzureReachabilityMessage); err != nil {
				return err
			}
		}
	}

	if o.AzureReachabilityReachableQueryParameter != nil {

		// query param azure_reachability.reachable
		var qrAzureReachabilityReachable bool

		if o.AzureReachabilityReachableQueryParameter != nil {
			qrAzureReachabilityReachable = *o.AzureReachabilityReachableQueryParameter
		}
		qAzureReachabilityReachable := swag.FormatBool(qrAzureReachabilityReachable)
		if qAzureReachabilityReachable != "" {

			if err := r.SetQueryParam("azure_reachability.reachable", qAzureReachabilityReachable); err != nil {
				return err
			}
		}
	}

	if o.ClientIDQueryParameter != nil {

		// query param client_id
		var qrClientID string

		if o.ClientIDQueryParameter != nil {
			qrClientID = *o.ClientIDQueryParameter
		}
		qClientID := qrClientID
		if qClientID != "" {

			if err := r.SetQueryParam("client_id", qClientID); err != nil {
				return err
			}
		}
	}

	if o.EkmIPReachabilityCodeQueryParameter != nil {

		// query param ekmip_reachability.code
		var qrEkmipReachabilityCode int64

		if o.EkmIPReachabilityCodeQueryParameter != nil {
			qrEkmipReachabilityCode = *o.EkmIPReachabilityCodeQueryParameter
		}
		qEkmipReachabilityCode := swag.FormatInt64(qrEkmipReachabilityCode)
		if qEkmipReachabilityCode != "" {

			if err := r.SetQueryParam("ekmip_reachability.code", qEkmipReachabilityCode); err != nil {
				return err
			}
		}
	}

	if o.EkmIPReachabilityMessageQueryParameter != nil {

		// query param ekmip_reachability.message
		var qrEkmipReachabilityMessage string

		if o.EkmIPReachabilityMessageQueryParameter != nil {
			qrEkmipReachabilityMessage = *o.EkmIPReachabilityMessageQueryParameter
		}
		qEkmipReachabilityMessage := qrEkmipReachabilityMessage
		if qEkmipReachabilityMessage != "" {

			if err := r.SetQueryParam("ekmip_reachability.message", qEkmipReachabilityMessage); err != nil {
				return err
			}
		}
	}

	if o.EkmIPReachabilityNodeNameQueryParameter != nil {

		// query param ekmip_reachability.node.name
		var qrEkmipReachabilityNodeName string

		if o.EkmIPReachabilityNodeNameQueryParameter != nil {
			qrEkmipReachabilityNodeName = *o.EkmIPReachabilityNodeNameQueryParameter
		}
		qEkmipReachabilityNodeName := qrEkmipReachabilityNodeName
		if qEkmipReachabilityNodeName != "" {

			if err := r.SetQueryParam("ekmip_reachability.node.name", qEkmipReachabilityNodeName); err != nil {
				return err
			}
		}
	}

	if o.EkmIPReachabilityNodeUUIDQueryParameter != nil {

		// query param ekmip_reachability.node.uuid
		var qrEkmipReachabilityNodeUUID string

		if o.EkmIPReachabilityNodeUUIDQueryParameter != nil {
			qrEkmipReachabilityNodeUUID = *o.EkmIPReachabilityNodeUUIDQueryParameter
		}
		qEkmipReachabilityNodeUUID := qrEkmipReachabilityNodeUUID
		if qEkmipReachabilityNodeUUID != "" {

			if err := r.SetQueryParam("ekmip_reachability.node.uuid", qEkmipReachabilityNodeUUID); err != nil {
				return err
			}
		}
	}

	if o.EkmIPReachabilityReachableQueryParameter != nil {

		// query param ekmip_reachability.reachable
		var qrEkmipReachabilityReachable bool

		if o.EkmIPReachabilityReachableQueryParameter != nil {
			qrEkmipReachabilityReachable = *o.EkmIPReachabilityReachableQueryParameter
		}
		qEkmipReachabilityReachable := swag.FormatBool(qrEkmipReachabilityReachable)
		if qEkmipReachabilityReachable != "" {

			if err := r.SetQueryParam("ekmip_reachability.reachable", qEkmipReachabilityReachable); err != nil {
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

	if o.KeyIDQueryParameter != nil {

		// query param key_id
		var qrKeyID string

		if o.KeyIDQueryParameter != nil {
			qrKeyID = *o.KeyIDQueryParameter
		}
		qKeyID := qrKeyID
		if qKeyID != "" {

			if err := r.SetQueryParam("key_id", qKeyID); err != nil {
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

	if o.ProxyHostQueryParameter != nil {

		// query param proxy_host
		var qrProxyHost string

		if o.ProxyHostQueryParameter != nil {
			qrProxyHost = *o.ProxyHostQueryParameter
		}
		qProxyHost := qrProxyHost
		if qProxyHost != "" {

			if err := r.SetQueryParam("proxy_host", qProxyHost); err != nil {
				return err
			}
		}
	}

	if o.ProxyPortQueryParameter != nil {

		// query param proxy_port
		var qrProxyPort int64

		if o.ProxyPortQueryParameter != nil {
			qrProxyPort = *o.ProxyPortQueryParameter
		}
		qProxyPort := swag.FormatInt64(qrProxyPort)
		if qProxyPort != "" {

			if err := r.SetQueryParam("proxy_port", qProxyPort); err != nil {
				return err
			}
		}
	}

	if o.ProxyTypeQueryParameter != nil {

		// query param proxy_type
		var qrProxyType string

		if o.ProxyTypeQueryParameter != nil {
			qrProxyType = *o.ProxyTypeQueryParameter
		}
		qProxyType := qrProxyType
		if qProxyType != "" {

			if err := r.SetQueryParam("proxy_type", qProxyType); err != nil {
				return err
			}
		}
	}

	if o.ProxyUsernameQueryParameter != nil {

		// query param proxy_username
		var qrProxyUsername string

		if o.ProxyUsernameQueryParameter != nil {
			qrProxyUsername = *o.ProxyUsernameQueryParameter
		}
		qProxyUsername := qrProxyUsername
		if qProxyUsername != "" {

			if err := r.SetQueryParam("proxy_username", qProxyUsername); err != nil {
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

	if o.StateAvailableQueryParameter != nil {

		// query param state.available
		var qrStateAvailable bool

		if o.StateAvailableQueryParameter != nil {
			qrStateAvailable = *o.StateAvailableQueryParameter
		}
		qStateAvailable := swag.FormatBool(qrStateAvailable)
		if qStateAvailable != "" {

			if err := r.SetQueryParam("state.available", qStateAvailable); err != nil {
				return err
			}
		}
	}

	if o.StateCodeQueryParameter != nil {

		// query param state.code
		var qrStateCode int64

		if o.StateCodeQueryParameter != nil {
			qrStateCode = *o.StateCodeQueryParameter
		}
		qStateCode := swag.FormatInt64(qrStateCode)
		if qStateCode != "" {

			if err := r.SetQueryParam("state.code", qStateCode); err != nil {
				return err
			}
		}
	}

	if o.StateMessageQueryParameter != nil {

		// query param state.message
		var qrStateMessage string

		if o.StateMessageQueryParameter != nil {
			qrStateMessage = *o.StateMessageQueryParameter
		}
		qStateMessage := qrStateMessage
		if qStateMessage != "" {

			if err := r.SetQueryParam("state.message", qStateMessage); err != nil {
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

	if o.TenantIDQueryParameter != nil {

		// query param tenant_id
		var qrTenantID string

		if o.TenantIDQueryParameter != nil {
			qrTenantID = *o.TenantIDQueryParameter
		}
		qTenantID := qrTenantID
		if qTenantID != "" {

			if err := r.SetQueryParam("tenant_id", qTenantID); err != nil {
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

// bindParamAzureKeyVaultCollectionGet binds the parameter fields
func (o *AzureKeyVaultCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamAzureKeyVaultCollectionGet binds the parameter order_by
func (o *AzureKeyVaultCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
