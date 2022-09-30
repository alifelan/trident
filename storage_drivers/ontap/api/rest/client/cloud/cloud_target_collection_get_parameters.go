// Code generated by go-swagger; DO NOT EDIT.

package cloud

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

// NewCloudTargetCollectionGetParams creates a new CloudTargetCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloudTargetCollectionGetParams() *CloudTargetCollectionGetParams {
	return &CloudTargetCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloudTargetCollectionGetParamsWithTimeout creates a new CloudTargetCollectionGetParams object
// with the ability to set a timeout on a request.
func NewCloudTargetCollectionGetParamsWithTimeout(timeout time.Duration) *CloudTargetCollectionGetParams {
	return &CloudTargetCollectionGetParams{
		timeout: timeout,
	}
}

// NewCloudTargetCollectionGetParamsWithContext creates a new CloudTargetCollectionGetParams object
// with the ability to set a context for a request.
func NewCloudTargetCollectionGetParamsWithContext(ctx context.Context) *CloudTargetCollectionGetParams {
	return &CloudTargetCollectionGetParams{
		Context: ctx,
	}
}

// NewCloudTargetCollectionGetParamsWithHTTPClient creates a new CloudTargetCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloudTargetCollectionGetParamsWithHTTPClient(client *http.Client) *CloudTargetCollectionGetParams {
	return &CloudTargetCollectionGetParams{
		HTTPClient: client,
	}
}

/*
CloudTargetCollectionGetParams contains all the parameters to send to the API endpoint

	for the cloud target collection get operation.

	Typically these are written to a http.Request.
*/
type CloudTargetCollectionGetParams struct {

	/* AccessKey.

	   Filter by access_key
	*/
	AccessKeyQueryParameter *string

	/* AuthenticationType.

	   Filter by authentication_type
	*/
	AuthenticationTypeQueryParameter *string

	/* AzureAccount.

	   Filter by azure_account
	*/
	AzureAccountQueryParameter *string

	/* CapURL.

	   Filter by cap_url
	*/
	CapURLQueryParameter *string

	/* CertificateValidationEnabled.

	   Filter by certificate_validation_enabled
	*/
	CertificateValidationEnabledQueryParameter *bool

	/* ClusterName.

	   Filter by cluster.name
	*/
	ClusterNameQueryParameter *string

	/* ClusterUUID.

	   Filter by cluster.uuid
	*/
	ClusterUUIDQueryParameter *string

	/* Container.

	   Filter by container
	*/
	ContainerQueryParameter *string

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

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* Owner.

	   Filter by owner
	*/
	OwnerQueryParameter *string

	/* Port.

	   Filter by port
	*/
	PortQueryParameter *int64

	/* ProviderType.

	   Filter by provider_type
	*/
	ProviderTypeQueryParameter *string

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

	/* Server.

	   Filter by server
	*/
	ServerQueryParameter *string

	/* ServerSideEncryption.

	   Filter by server_side_encryption
	*/
	ServerSideEncryptionQueryParameter *string

	/* SnapmirrorUse.

	   Filter by snapmirror_use
	*/
	SnapmirrorUseQueryParameter *string

	/* SslEnabled.

	   Filter by ssl_enabled
	*/
	SslEnabledQueryParameter *bool

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* URLStyle.

	   Filter by url_style
	*/
	URLStyleQueryParameter *string

	/* UseHTTPProxy.

	   Filter by use_http_proxy
	*/
	UseHTTPProxyQueryParameter *bool

	/* Used.

	   Filter by used
	*/
	UsedQueryParameter *int64

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cloud target collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudTargetCollectionGetParams) WithDefaults() *CloudTargetCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cloud target collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudTargetCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := CloudTargetCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithTimeout(timeout time.Duration) *CloudTargetCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithContext(ctx context.Context) *CloudTargetCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithHTTPClient(client *http.Client) *CloudTargetCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccessKeyQueryParameter adds the accessKey to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithAccessKeyQueryParameter(accessKey *string) *CloudTargetCollectionGetParams {
	o.SetAccessKeyQueryParameter(accessKey)
	return o
}

// SetAccessKeyQueryParameter adds the accessKey to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetAccessKeyQueryParameter(accessKey *string) {
	o.AccessKeyQueryParameter = accessKey
}

// WithAuthenticationTypeQueryParameter adds the authenticationType to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithAuthenticationTypeQueryParameter(authenticationType *string) *CloudTargetCollectionGetParams {
	o.SetAuthenticationTypeQueryParameter(authenticationType)
	return o
}

// SetAuthenticationTypeQueryParameter adds the authenticationType to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetAuthenticationTypeQueryParameter(authenticationType *string) {
	o.AuthenticationTypeQueryParameter = authenticationType
}

// WithAzureAccountQueryParameter adds the azureAccount to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithAzureAccountQueryParameter(azureAccount *string) *CloudTargetCollectionGetParams {
	o.SetAzureAccountQueryParameter(azureAccount)
	return o
}

// SetAzureAccountQueryParameter adds the azureAccount to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetAzureAccountQueryParameter(azureAccount *string) {
	o.AzureAccountQueryParameter = azureAccount
}

// WithCapURLQueryParameter adds the capURL to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithCapURLQueryParameter(capURL *string) *CloudTargetCollectionGetParams {
	o.SetCapURLQueryParameter(capURL)
	return o
}

// SetCapURLQueryParameter adds the capUrl to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetCapURLQueryParameter(capURL *string) {
	o.CapURLQueryParameter = capURL
}

// WithCertificateValidationEnabledQueryParameter adds the certificateValidationEnabled to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithCertificateValidationEnabledQueryParameter(certificateValidationEnabled *bool) *CloudTargetCollectionGetParams {
	o.SetCertificateValidationEnabledQueryParameter(certificateValidationEnabled)
	return o
}

// SetCertificateValidationEnabledQueryParameter adds the certificateValidationEnabled to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetCertificateValidationEnabledQueryParameter(certificateValidationEnabled *bool) {
	o.CertificateValidationEnabledQueryParameter = certificateValidationEnabled
}

// WithClusterNameQueryParameter adds the clusterName to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithClusterNameQueryParameter(clusterName *string) *CloudTargetCollectionGetParams {
	o.SetClusterNameQueryParameter(clusterName)
	return o
}

// SetClusterNameQueryParameter adds the clusterName to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetClusterNameQueryParameter(clusterName *string) {
	o.ClusterNameQueryParameter = clusterName
}

// WithClusterUUIDQueryParameter adds the clusterUUID to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithClusterUUIDQueryParameter(clusterUUID *string) *CloudTargetCollectionGetParams {
	o.SetClusterUUIDQueryParameter(clusterUUID)
	return o
}

// SetClusterUUIDQueryParameter adds the clusterUuid to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetClusterUUIDQueryParameter(clusterUUID *string) {
	o.ClusterUUIDQueryParameter = clusterUUID
}

// WithContainerQueryParameter adds the container to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithContainerQueryParameter(container *string) *CloudTargetCollectionGetParams {
	o.SetContainerQueryParameter(container)
	return o
}

// SetContainerQueryParameter adds the container to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetContainerQueryParameter(container *string) {
	o.ContainerQueryParameter = container
}

// WithFieldsQueryParameter adds the fields to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithFieldsQueryParameter(fields []string) *CloudTargetCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIpspaceNameQueryParameter adds the ipspaceName to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithIpspaceNameQueryParameter(ipspaceName *string) *CloudTargetCollectionGetParams {
	o.SetIpspaceNameQueryParameter(ipspaceName)
	return o
}

// SetIpspaceNameQueryParameter adds the ipspaceName to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetIpspaceNameQueryParameter(ipspaceName *string) {
	o.IpspaceNameQueryParameter = ipspaceName
}

// WithIpspaceUUIDQueryParameter adds the ipspaceUUID to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithIpspaceUUIDQueryParameter(ipspaceUUID *string) *CloudTargetCollectionGetParams {
	o.SetIpspaceUUIDQueryParameter(ipspaceUUID)
	return o
}

// SetIpspaceUUIDQueryParameter adds the ipspaceUuid to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetIpspaceUUIDQueryParameter(ipspaceUUID *string) {
	o.IpspaceUUIDQueryParameter = ipspaceUUID
}

// WithMaxRecordsQueryParameter adds the maxRecords to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *CloudTargetCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithNameQueryParameter(name *string) *CloudTargetCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *CloudTargetCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithOwnerQueryParameter adds the owner to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithOwnerQueryParameter(owner *string) *CloudTargetCollectionGetParams {
	o.SetOwnerQueryParameter(owner)
	return o
}

// SetOwnerQueryParameter adds the owner to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetOwnerQueryParameter(owner *string) {
	o.OwnerQueryParameter = owner
}

// WithPortQueryParameter adds the port to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithPortQueryParameter(port *int64) *CloudTargetCollectionGetParams {
	o.SetPortQueryParameter(port)
	return o
}

// SetPortQueryParameter adds the port to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetPortQueryParameter(port *int64) {
	o.PortQueryParameter = port
}

// WithProviderTypeQueryParameter adds the providerType to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithProviderTypeQueryParameter(providerType *string) *CloudTargetCollectionGetParams {
	o.SetProviderTypeQueryParameter(providerType)
	return o
}

// SetProviderTypeQueryParameter adds the providerType to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetProviderTypeQueryParameter(providerType *string) {
	o.ProviderTypeQueryParameter = providerType
}

// WithReturnRecordsQueryParameter adds the returnRecords to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *CloudTargetCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *CloudTargetCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithServerQueryParameter adds the server to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithServerQueryParameter(server *string) *CloudTargetCollectionGetParams {
	o.SetServerQueryParameter(server)
	return o
}

// SetServerQueryParameter adds the server to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetServerQueryParameter(server *string) {
	o.ServerQueryParameter = server
}

// WithServerSideEncryptionQueryParameter adds the serverSideEncryption to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithServerSideEncryptionQueryParameter(serverSideEncryption *string) *CloudTargetCollectionGetParams {
	o.SetServerSideEncryptionQueryParameter(serverSideEncryption)
	return o
}

// SetServerSideEncryptionQueryParameter adds the serverSideEncryption to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetServerSideEncryptionQueryParameter(serverSideEncryption *string) {
	o.ServerSideEncryptionQueryParameter = serverSideEncryption
}

// WithSnapmirrorUseQueryParameter adds the snapmirrorUse to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithSnapmirrorUseQueryParameter(snapmirrorUse *string) *CloudTargetCollectionGetParams {
	o.SetSnapmirrorUseQueryParameter(snapmirrorUse)
	return o
}

// SetSnapmirrorUseQueryParameter adds the snapmirrorUse to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetSnapmirrorUseQueryParameter(snapmirrorUse *string) {
	o.SnapmirrorUseQueryParameter = snapmirrorUse
}

// WithSslEnabledQueryParameter adds the sslEnabled to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithSslEnabledQueryParameter(sslEnabled *bool) *CloudTargetCollectionGetParams {
	o.SetSslEnabledQueryParameter(sslEnabled)
	return o
}

// SetSslEnabledQueryParameter adds the sslEnabled to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetSslEnabledQueryParameter(sslEnabled *bool) {
	o.SslEnabledQueryParameter = sslEnabled
}

// WithSVMNameQueryParameter adds the svmName to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *CloudTargetCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *CloudTargetCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithURLStyleQueryParameter adds the uRLStyle to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithURLStyleQueryParameter(uRLStyle *string) *CloudTargetCollectionGetParams {
	o.SetURLStyleQueryParameter(uRLStyle)
	return o
}

// SetURLStyleQueryParameter adds the urlStyle to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetURLStyleQueryParameter(uRLStyle *string) {
	o.URLStyleQueryParameter = uRLStyle
}

// WithUseHTTPProxyQueryParameter adds the useHTTPProxy to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithUseHTTPProxyQueryParameter(useHTTPProxy *bool) *CloudTargetCollectionGetParams {
	o.SetUseHTTPProxyQueryParameter(useHTTPProxy)
	return o
}

// SetUseHTTPProxyQueryParameter adds the useHttpProxy to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetUseHTTPProxyQueryParameter(useHTTPProxy *bool) {
	o.UseHTTPProxyQueryParameter = useHTTPProxy
}

// WithUsedQueryParameter adds the used to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithUsedQueryParameter(used *int64) *CloudTargetCollectionGetParams {
	o.SetUsedQueryParameter(used)
	return o
}

// SetUsedQueryParameter adds the used to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetUsedQueryParameter(used *int64) {
	o.UsedQueryParameter = used
}

// WithUUIDQueryParameter adds the uuid to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) WithUUIDQueryParameter(uuid *string) *CloudTargetCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the cloud target collection get params
func (o *CloudTargetCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *CloudTargetCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccessKeyQueryParameter != nil {

		// query param access_key
		var qrAccessKey string

		if o.AccessKeyQueryParameter != nil {
			qrAccessKey = *o.AccessKeyQueryParameter
		}
		qAccessKey := qrAccessKey
		if qAccessKey != "" {

			if err := r.SetQueryParam("access_key", qAccessKey); err != nil {
				return err
			}
		}
	}

	if o.AuthenticationTypeQueryParameter != nil {

		// query param authentication_type
		var qrAuthenticationType string

		if o.AuthenticationTypeQueryParameter != nil {
			qrAuthenticationType = *o.AuthenticationTypeQueryParameter
		}
		qAuthenticationType := qrAuthenticationType
		if qAuthenticationType != "" {

			if err := r.SetQueryParam("authentication_type", qAuthenticationType); err != nil {
				return err
			}
		}
	}

	if o.AzureAccountQueryParameter != nil {

		// query param azure_account
		var qrAzureAccount string

		if o.AzureAccountQueryParameter != nil {
			qrAzureAccount = *o.AzureAccountQueryParameter
		}
		qAzureAccount := qrAzureAccount
		if qAzureAccount != "" {

			if err := r.SetQueryParam("azure_account", qAzureAccount); err != nil {
				return err
			}
		}
	}

	if o.CapURLQueryParameter != nil {

		// query param cap_url
		var qrCapURL string

		if o.CapURLQueryParameter != nil {
			qrCapURL = *o.CapURLQueryParameter
		}
		qCapURL := qrCapURL
		if qCapURL != "" {

			if err := r.SetQueryParam("cap_url", qCapURL); err != nil {
				return err
			}
		}
	}

	if o.CertificateValidationEnabledQueryParameter != nil {

		// query param certificate_validation_enabled
		var qrCertificateValidationEnabled bool

		if o.CertificateValidationEnabledQueryParameter != nil {
			qrCertificateValidationEnabled = *o.CertificateValidationEnabledQueryParameter
		}
		qCertificateValidationEnabled := swag.FormatBool(qrCertificateValidationEnabled)
		if qCertificateValidationEnabled != "" {

			if err := r.SetQueryParam("certificate_validation_enabled", qCertificateValidationEnabled); err != nil {
				return err
			}
		}
	}

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

	if o.ContainerQueryParameter != nil {

		// query param container
		var qrContainer string

		if o.ContainerQueryParameter != nil {
			qrContainer = *o.ContainerQueryParameter
		}
		qContainer := qrContainer
		if qContainer != "" {

			if err := r.SetQueryParam("container", qContainer); err != nil {
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

	if o.OwnerQueryParameter != nil {

		// query param owner
		var qrOwner string

		if o.OwnerQueryParameter != nil {
			qrOwner = *o.OwnerQueryParameter
		}
		qOwner := qrOwner
		if qOwner != "" {

			if err := r.SetQueryParam("owner", qOwner); err != nil {
				return err
			}
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

	if o.ProviderTypeQueryParameter != nil {

		// query param provider_type
		var qrProviderType string

		if o.ProviderTypeQueryParameter != nil {
			qrProviderType = *o.ProviderTypeQueryParameter
		}
		qProviderType := qrProviderType
		if qProviderType != "" {

			if err := r.SetQueryParam("provider_type", qProviderType); err != nil {
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

	if o.ServerSideEncryptionQueryParameter != nil {

		// query param server_side_encryption
		var qrServerSideEncryption string

		if o.ServerSideEncryptionQueryParameter != nil {
			qrServerSideEncryption = *o.ServerSideEncryptionQueryParameter
		}
		qServerSideEncryption := qrServerSideEncryption
		if qServerSideEncryption != "" {

			if err := r.SetQueryParam("server_side_encryption", qServerSideEncryption); err != nil {
				return err
			}
		}
	}

	if o.SnapmirrorUseQueryParameter != nil {

		// query param snapmirror_use
		var qrSnapmirrorUse string

		if o.SnapmirrorUseQueryParameter != nil {
			qrSnapmirrorUse = *o.SnapmirrorUseQueryParameter
		}
		qSnapmirrorUse := qrSnapmirrorUse
		if qSnapmirrorUse != "" {

			if err := r.SetQueryParam("snapmirror_use", qSnapmirrorUse); err != nil {
				return err
			}
		}
	}

	if o.SslEnabledQueryParameter != nil {

		// query param ssl_enabled
		var qrSslEnabled bool

		if o.SslEnabledQueryParameter != nil {
			qrSslEnabled = *o.SslEnabledQueryParameter
		}
		qSslEnabled := swag.FormatBool(qrSslEnabled)
		if qSslEnabled != "" {

			if err := r.SetQueryParam("ssl_enabled", qSslEnabled); err != nil {
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

	if o.URLStyleQueryParameter != nil {

		// query param url_style
		var qrURLStyle string

		if o.URLStyleQueryParameter != nil {
			qrURLStyle = *o.URLStyleQueryParameter
		}
		qURLStyle := qrURLStyle
		if qURLStyle != "" {

			if err := r.SetQueryParam("url_style", qURLStyle); err != nil {
				return err
			}
		}
	}

	if o.UseHTTPProxyQueryParameter != nil {

		// query param use_http_proxy
		var qrUseHTTPProxy bool

		if o.UseHTTPProxyQueryParameter != nil {
			qrUseHTTPProxy = *o.UseHTTPProxyQueryParameter
		}
		qUseHTTPProxy := swag.FormatBool(qrUseHTTPProxy)
		if qUseHTTPProxy != "" {

			if err := r.SetQueryParam("use_http_proxy", qUseHTTPProxy); err != nil {
				return err
			}
		}
	}

	if o.UsedQueryParameter != nil {

		// query param used
		var qrUsed int64

		if o.UsedQueryParameter != nil {
			qrUsed = *o.UsedQueryParameter
		}
		qUsed := swag.FormatInt64(qrUsed)
		if qUsed != "" {

			if err := r.SetQueryParam("used", qUsed); err != nil {
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

// bindParamCloudTargetCollectionGet binds the parameter fields
func (o *CloudTargetCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamCloudTargetCollectionGet binds the parameter order_by
func (o *CloudTargetCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
