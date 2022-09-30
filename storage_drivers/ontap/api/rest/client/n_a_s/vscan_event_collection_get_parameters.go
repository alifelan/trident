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

// NewVscanEventCollectionGetParams creates a new VscanEventCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVscanEventCollectionGetParams() *VscanEventCollectionGetParams {
	return &VscanEventCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVscanEventCollectionGetParamsWithTimeout creates a new VscanEventCollectionGetParams object
// with the ability to set a timeout on a request.
func NewVscanEventCollectionGetParamsWithTimeout(timeout time.Duration) *VscanEventCollectionGetParams {
	return &VscanEventCollectionGetParams{
		timeout: timeout,
	}
}

// NewVscanEventCollectionGetParamsWithContext creates a new VscanEventCollectionGetParams object
// with the ability to set a context for a request.
func NewVscanEventCollectionGetParamsWithContext(ctx context.Context) *VscanEventCollectionGetParams {
	return &VscanEventCollectionGetParams{
		Context: ctx,
	}
}

// NewVscanEventCollectionGetParamsWithHTTPClient creates a new VscanEventCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewVscanEventCollectionGetParamsWithHTTPClient(client *http.Client) *VscanEventCollectionGetParams {
	return &VscanEventCollectionGetParams{
		HTTPClient: client,
	}
}

/*
VscanEventCollectionGetParams contains all the parameters to send to the API endpoint

	for the vscan event collection get operation.

	Typically these are written to a http.Request.
*/
type VscanEventCollectionGetParams struct {

	/* DisconnectReason.

	   Filter by disconnect_reason
	*/
	DisconnectReasonQueryParameter *string

	/* EventTime.

	   Filter by event_time
	*/
	EventTimeQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* FilePath.

	   Filter by file_path
	*/
	FilePathQueryParameter *string

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

	/* Server.

	   Filter by server
	*/
	ServerQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	/* Type.

	   Filter by type
	*/
	TypeQueryParameter *string

	/* Vendor.

	   Filter by vendor
	*/
	VendorQueryParameter *string

	/* Version.

	   Filter by version
	*/
	VersionQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the vscan event collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanEventCollectionGetParams) WithDefaults() *VscanEventCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the vscan event collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VscanEventCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := VscanEventCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithTimeout(timeout time.Duration) *VscanEventCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithContext(ctx context.Context) *VscanEventCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithHTTPClient(client *http.Client) *VscanEventCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDisconnectReasonQueryParameter adds the disconnectReason to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithDisconnectReasonQueryParameter(disconnectReason *string) *VscanEventCollectionGetParams {
	o.SetDisconnectReasonQueryParameter(disconnectReason)
	return o
}

// SetDisconnectReasonQueryParameter adds the disconnectReason to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetDisconnectReasonQueryParameter(disconnectReason *string) {
	o.DisconnectReasonQueryParameter = disconnectReason
}

// WithEventTimeQueryParameter adds the eventTime to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithEventTimeQueryParameter(eventTime *string) *VscanEventCollectionGetParams {
	o.SetEventTimeQueryParameter(eventTime)
	return o
}

// SetEventTimeQueryParameter adds the eventTime to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetEventTimeQueryParameter(eventTime *string) {
	o.EventTimeQueryParameter = eventTime
}

// WithFieldsQueryParameter adds the fields to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithFieldsQueryParameter(fields []string) *VscanEventCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithFilePathQueryParameter adds the filePath to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithFilePathQueryParameter(filePath *string) *VscanEventCollectionGetParams {
	o.SetFilePathQueryParameter(filePath)
	return o
}

// SetFilePathQueryParameter adds the filePath to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetFilePathQueryParameter(filePath *string) {
	o.FilePathQueryParameter = filePath
}

// WithInterfaceIPAddressQueryParameter adds the interfaceIPAddress to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithInterfaceIPAddressQueryParameter(interfaceIPAddress *string) *VscanEventCollectionGetParams {
	o.SetInterfaceIPAddressQueryParameter(interfaceIPAddress)
	return o
}

// SetInterfaceIPAddressQueryParameter adds the interfaceIpAddress to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetInterfaceIPAddressQueryParameter(interfaceIPAddress *string) {
	o.InterfaceIPAddressQueryParameter = interfaceIPAddress
}

// WithInterfaceNameQueryParameter adds the interfaceName to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithInterfaceNameQueryParameter(interfaceName *string) *VscanEventCollectionGetParams {
	o.SetInterfaceNameQueryParameter(interfaceName)
	return o
}

// SetInterfaceNameQueryParameter adds the interfaceName to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetInterfaceNameQueryParameter(interfaceName *string) {
	o.InterfaceNameQueryParameter = interfaceName
}

// WithInterfaceUUIDQueryParameter adds the interfaceUUID to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithInterfaceUUIDQueryParameter(interfaceUUID *string) *VscanEventCollectionGetParams {
	o.SetInterfaceUUIDQueryParameter(interfaceUUID)
	return o
}

// SetInterfaceUUIDQueryParameter adds the interfaceUuid to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetInterfaceUUIDQueryParameter(interfaceUUID *string) {
	o.InterfaceUUIDQueryParameter = interfaceUUID
}

// WithMaxRecordsQueryParameter adds the maxRecords to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *VscanEventCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNodeNameQueryParameter adds the nodeName to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithNodeNameQueryParameter(nodeName *string) *VscanEventCollectionGetParams {
	o.SetNodeNameQueryParameter(nodeName)
	return o
}

// SetNodeNameQueryParameter adds the nodeName to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetNodeNameQueryParameter(nodeName *string) {
	o.NodeNameQueryParameter = nodeName
}

// WithNodeUUIDQueryParameter adds the nodeUUID to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithNodeUUIDQueryParameter(nodeUUID *string) *VscanEventCollectionGetParams {
	o.SetNodeUUIDQueryParameter(nodeUUID)
	return o
}

// SetNodeUUIDQueryParameter adds the nodeUuid to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetNodeUUIDQueryParameter(nodeUUID *string) {
	o.NodeUUIDQueryParameter = nodeUUID
}

// WithOrderByQueryParameter adds the orderBy to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *VscanEventCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *VscanEventCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *VscanEventCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithServerQueryParameter adds the server to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithServerQueryParameter(server *string) *VscanEventCollectionGetParams {
	o.SetServerQueryParameter(server)
	return o
}

// SetServerQueryParameter adds the server to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetServerQueryParameter(server *string) {
	o.ServerQueryParameter = server
}

// WithSVMNameQueryParameter adds the svmName to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *VscanEventCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDPathParameter adds the svmUUID to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithSVMUUIDPathParameter(svmUUID string) *VscanEventCollectionGetParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WithTypeQueryParameter adds the typeVar to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithTypeQueryParameter(typeVar *string) *VscanEventCollectionGetParams {
	o.SetTypeQueryParameter(typeVar)
	return o
}

// SetTypeQueryParameter adds the type to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetTypeQueryParameter(typeVar *string) {
	o.TypeQueryParameter = typeVar
}

// WithVendorQueryParameter adds the vendor to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithVendorQueryParameter(vendor *string) *VscanEventCollectionGetParams {
	o.SetVendorQueryParameter(vendor)
	return o
}

// SetVendorQueryParameter adds the vendor to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetVendorQueryParameter(vendor *string) {
	o.VendorQueryParameter = vendor
}

// WithVersionQueryParameter adds the version to the vscan event collection get params
func (o *VscanEventCollectionGetParams) WithVersionQueryParameter(version *string) *VscanEventCollectionGetParams {
	o.SetVersionQueryParameter(version)
	return o
}

// SetVersionQueryParameter adds the version to the vscan event collection get params
func (o *VscanEventCollectionGetParams) SetVersionQueryParameter(version *string) {
	o.VersionQueryParameter = version
}

// WriteToRequest writes these params to a swagger request
func (o *VscanEventCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.DisconnectReasonQueryParameter != nil {

		// query param disconnect_reason
		var qrDisconnectReason string

		if o.DisconnectReasonQueryParameter != nil {
			qrDisconnectReason = *o.DisconnectReasonQueryParameter
		}
		qDisconnectReason := qrDisconnectReason
		if qDisconnectReason != "" {

			if err := r.SetQueryParam("disconnect_reason", qDisconnectReason); err != nil {
				return err
			}
		}
	}

	if o.EventTimeQueryParameter != nil {

		// query param event_time
		var qrEventTime string

		if o.EventTimeQueryParameter != nil {
			qrEventTime = *o.EventTimeQueryParameter
		}
		qEventTime := qrEventTime
		if qEventTime != "" {

			if err := r.SetQueryParam("event_time", qEventTime); err != nil {
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

	if o.FilePathQueryParameter != nil {

		// query param file_path
		var qrFilePath string

		if o.FilePathQueryParameter != nil {
			qrFilePath = *o.FilePathQueryParameter
		}
		qFilePath := qrFilePath
		if qFilePath != "" {

			if err := r.SetQueryParam("file_path", qFilePath); err != nil {
				return err
			}
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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
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

	if o.VendorQueryParameter != nil {

		// query param vendor
		var qrVendor string

		if o.VendorQueryParameter != nil {
			qrVendor = *o.VendorQueryParameter
		}
		qVendor := qrVendor
		if qVendor != "" {

			if err := r.SetQueryParam("vendor", qVendor); err != nil {
				return err
			}
		}
	}

	if o.VersionQueryParameter != nil {

		// query param version
		var qrVersion string

		if o.VersionQueryParameter != nil {
			qrVersion = *o.VersionQueryParameter
		}
		qVersion := qrVersion
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamVscanEventCollectionGet binds the parameter fields
func (o *VscanEventCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamVscanEventCollectionGet binds the parameter order_by
func (o *VscanEventCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
