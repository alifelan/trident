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

// NewNvmeSubsystemCollectionGetParams creates a new NvmeSubsystemCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNvmeSubsystemCollectionGetParams() *NvmeSubsystemCollectionGetParams {
	return &NvmeSubsystemCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNvmeSubsystemCollectionGetParamsWithTimeout creates a new NvmeSubsystemCollectionGetParams object
// with the ability to set a timeout on a request.
func NewNvmeSubsystemCollectionGetParamsWithTimeout(timeout time.Duration) *NvmeSubsystemCollectionGetParams {
	return &NvmeSubsystemCollectionGetParams{
		timeout: timeout,
	}
}

// NewNvmeSubsystemCollectionGetParamsWithContext creates a new NvmeSubsystemCollectionGetParams object
// with the ability to set a context for a request.
func NewNvmeSubsystemCollectionGetParamsWithContext(ctx context.Context) *NvmeSubsystemCollectionGetParams {
	return &NvmeSubsystemCollectionGetParams{
		Context: ctx,
	}
}

// NewNvmeSubsystemCollectionGetParamsWithHTTPClient creates a new NvmeSubsystemCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNvmeSubsystemCollectionGetParamsWithHTTPClient(client *http.Client) *NvmeSubsystemCollectionGetParams {
	return &NvmeSubsystemCollectionGetParams{
		HTTPClient: client,
	}
}

/*
NvmeSubsystemCollectionGetParams contains all the parameters to send to the API endpoint

	for the nvme subsystem collection get operation.

	Typically these are written to a http.Request.
*/
type NvmeSubsystemCollectionGetParams struct {

	/* Comment.

	   Filter by comment
	*/
	CommentQueryParameter *string

	/* DeleteOnUnmap.

	   Filter by delete_on_unmap
	*/
	DeleteOnUnmapQueryParameter *bool

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* HostsNqn.

	   Filter by hosts.nqn
	*/
	HostsNqnQueryParameter *string

	/* IoQueueDefaultCount.

	   Filter by io_queue.default.count
	*/
	IoQueueDefaultCountQueryParameter *int64

	/* IoQueueDefaultDepth.

	   Filter by io_queue.default.depth
	*/
	IoQueueDefaultDepthQueryParameter *int64

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

	/* OsType.

	   Filter by os_type
	*/
	OsTypeQueryParameter *string

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

	/* SerialNumber.

	   Filter by serial_number
	*/
	SerialNumberQueryParameter *string

	/* SubsystemMapsAnagrpid.

	   Filter by subsystem_maps.anagrpid
	*/
	SubsystemMapsAnagrpIDQueryParameter *string

	/* SubsystemMapsNamespaceName.

	   Filter by subsystem_maps.namespace.name
	*/
	SubsystemMapsNamespaceNameQueryParameter *string

	/* SubsystemMapsNamespaceUUID.

	   Filter by subsystem_maps.namespace.uuid
	*/
	SubsystemMapsNamespaceUUIDQueryParameter *string

	/* SubsystemMapsNsid.

	   Filter by subsystem_maps.nsid
	*/
	SubsystemMapsNsIDQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* TargetNqn.

	   Filter by target_nqn
	*/
	TargetNqnQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

	/* VendorUuids.

	   Filter by vendor_uuids
	*/
	VendorUUIDsQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the nvme subsystem collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemCollectionGetParams) WithDefaults() *NvmeSubsystemCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the nvme subsystem collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NvmeSubsystemCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := NvmeSubsystemCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithTimeout(timeout time.Duration) *NvmeSubsystemCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithContext(ctx context.Context) *NvmeSubsystemCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithHTTPClient(client *http.Client) *NvmeSubsystemCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCommentQueryParameter adds the comment to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithCommentQueryParameter(comment *string) *NvmeSubsystemCollectionGetParams {
	o.SetCommentQueryParameter(comment)
	return o
}

// SetCommentQueryParameter adds the comment to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetCommentQueryParameter(comment *string) {
	o.CommentQueryParameter = comment
}

// WithDeleteOnUnmapQueryParameter adds the deleteOnUnmap to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithDeleteOnUnmapQueryParameter(deleteOnUnmap *bool) *NvmeSubsystemCollectionGetParams {
	o.SetDeleteOnUnmapQueryParameter(deleteOnUnmap)
	return o
}

// SetDeleteOnUnmapQueryParameter adds the deleteOnUnmap to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetDeleteOnUnmapQueryParameter(deleteOnUnmap *bool) {
	o.DeleteOnUnmapQueryParameter = deleteOnUnmap
}

// WithFieldsQueryParameter adds the fields to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithFieldsQueryParameter(fields []string) *NvmeSubsystemCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithHostsNqnQueryParameter adds the hostsNqn to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithHostsNqnQueryParameter(hostsNqn *string) *NvmeSubsystemCollectionGetParams {
	o.SetHostsNqnQueryParameter(hostsNqn)
	return o
}

// SetHostsNqnQueryParameter adds the hostsNqn to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetHostsNqnQueryParameter(hostsNqn *string) {
	o.HostsNqnQueryParameter = hostsNqn
}

// WithIoQueueDefaultCountQueryParameter adds the ioQueueDefaultCount to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithIoQueueDefaultCountQueryParameter(ioQueueDefaultCount *int64) *NvmeSubsystemCollectionGetParams {
	o.SetIoQueueDefaultCountQueryParameter(ioQueueDefaultCount)
	return o
}

// SetIoQueueDefaultCountQueryParameter adds the ioQueueDefaultCount to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetIoQueueDefaultCountQueryParameter(ioQueueDefaultCount *int64) {
	o.IoQueueDefaultCountQueryParameter = ioQueueDefaultCount
}

// WithIoQueueDefaultDepthQueryParameter adds the ioQueueDefaultDepth to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithIoQueueDefaultDepthQueryParameter(ioQueueDefaultDepth *int64) *NvmeSubsystemCollectionGetParams {
	o.SetIoQueueDefaultDepthQueryParameter(ioQueueDefaultDepth)
	return o
}

// SetIoQueueDefaultDepthQueryParameter adds the ioQueueDefaultDepth to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetIoQueueDefaultDepthQueryParameter(ioQueueDefaultDepth *int64) {
	o.IoQueueDefaultDepthQueryParameter = ioQueueDefaultDepth
}

// WithMaxRecordsQueryParameter adds the maxRecords to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *NvmeSubsystemCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithNameQueryParameter(name *string) *NvmeSubsystemCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *NvmeSubsystemCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithOsTypeQueryParameter adds the osType to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithOsTypeQueryParameter(osType *string) *NvmeSubsystemCollectionGetParams {
	o.SetOsTypeQueryParameter(osType)
	return o
}

// SetOsTypeQueryParameter adds the osType to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetOsTypeQueryParameter(osType *string) {
	o.OsTypeQueryParameter = osType
}

// WithReturnRecordsQueryParameter adds the returnRecords to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *NvmeSubsystemCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *NvmeSubsystemCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSerialNumberQueryParameter adds the serialNumber to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithSerialNumberQueryParameter(serialNumber *string) *NvmeSubsystemCollectionGetParams {
	o.SetSerialNumberQueryParameter(serialNumber)
	return o
}

// SetSerialNumberQueryParameter adds the serialNumber to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetSerialNumberQueryParameter(serialNumber *string) {
	o.SerialNumberQueryParameter = serialNumber
}

// WithSubsystemMapsAnagrpIDQueryParameter adds the subsystemMapsAnagrpid to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithSubsystemMapsAnagrpIDQueryParameter(subsystemMapsAnagrpid *string) *NvmeSubsystemCollectionGetParams {
	o.SetSubsystemMapsAnagrpIDQueryParameter(subsystemMapsAnagrpid)
	return o
}

// SetSubsystemMapsAnagrpIDQueryParameter adds the subsystemMapsAnagrpid to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetSubsystemMapsAnagrpIDQueryParameter(subsystemMapsAnagrpid *string) {
	o.SubsystemMapsAnagrpIDQueryParameter = subsystemMapsAnagrpid
}

// WithSubsystemMapsNamespaceNameQueryParameter adds the subsystemMapsNamespaceName to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithSubsystemMapsNamespaceNameQueryParameter(subsystemMapsNamespaceName *string) *NvmeSubsystemCollectionGetParams {
	o.SetSubsystemMapsNamespaceNameQueryParameter(subsystemMapsNamespaceName)
	return o
}

// SetSubsystemMapsNamespaceNameQueryParameter adds the subsystemMapsNamespaceName to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetSubsystemMapsNamespaceNameQueryParameter(subsystemMapsNamespaceName *string) {
	o.SubsystemMapsNamespaceNameQueryParameter = subsystemMapsNamespaceName
}

// WithSubsystemMapsNamespaceUUIDQueryParameter adds the subsystemMapsNamespaceUUID to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithSubsystemMapsNamespaceUUIDQueryParameter(subsystemMapsNamespaceUUID *string) *NvmeSubsystemCollectionGetParams {
	o.SetSubsystemMapsNamespaceUUIDQueryParameter(subsystemMapsNamespaceUUID)
	return o
}

// SetSubsystemMapsNamespaceUUIDQueryParameter adds the subsystemMapsNamespaceUuid to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetSubsystemMapsNamespaceUUIDQueryParameter(subsystemMapsNamespaceUUID *string) {
	o.SubsystemMapsNamespaceUUIDQueryParameter = subsystemMapsNamespaceUUID
}

// WithSubsystemMapsNsIDQueryParameter adds the subsystemMapsNsid to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithSubsystemMapsNsIDQueryParameter(subsystemMapsNsid *string) *NvmeSubsystemCollectionGetParams {
	o.SetSubsystemMapsNsIDQueryParameter(subsystemMapsNsid)
	return o
}

// SetSubsystemMapsNsIDQueryParameter adds the subsystemMapsNsid to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetSubsystemMapsNsIDQueryParameter(subsystemMapsNsid *string) {
	o.SubsystemMapsNsIDQueryParameter = subsystemMapsNsid
}

// WithSVMNameQueryParameter adds the svmName to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *NvmeSubsystemCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *NvmeSubsystemCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithTargetNqnQueryParameter adds the targetNqn to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithTargetNqnQueryParameter(targetNqn *string) *NvmeSubsystemCollectionGetParams {
	o.SetTargetNqnQueryParameter(targetNqn)
	return o
}

// SetTargetNqnQueryParameter adds the targetNqn to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetTargetNqnQueryParameter(targetNqn *string) {
	o.TargetNqnQueryParameter = targetNqn
}

// WithUUIDQueryParameter adds the uuid to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithUUIDQueryParameter(uuid *string) *NvmeSubsystemCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WithVendorUUIDsQueryParameter adds the vendorUuids to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) WithVendorUUIDsQueryParameter(vendorUuids *string) *NvmeSubsystemCollectionGetParams {
	o.SetVendorUUIDsQueryParameter(vendorUuids)
	return o
}

// SetVendorUUIDsQueryParameter adds the vendorUuids to the nvme subsystem collection get params
func (o *NvmeSubsystemCollectionGetParams) SetVendorUUIDsQueryParameter(vendorUuids *string) {
	o.VendorUUIDsQueryParameter = vendorUuids
}

// WriteToRequest writes these params to a swagger request
func (o *NvmeSubsystemCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CommentQueryParameter != nil {

		// query param comment
		var qrComment string

		if o.CommentQueryParameter != nil {
			qrComment = *o.CommentQueryParameter
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if o.DeleteOnUnmapQueryParameter != nil {

		// query param delete_on_unmap
		var qrDeleteOnUnmap bool

		if o.DeleteOnUnmapQueryParameter != nil {
			qrDeleteOnUnmap = *o.DeleteOnUnmapQueryParameter
		}
		qDeleteOnUnmap := swag.FormatBool(qrDeleteOnUnmap)
		if qDeleteOnUnmap != "" {

			if err := r.SetQueryParam("delete_on_unmap", qDeleteOnUnmap); err != nil {
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

	if o.HostsNqnQueryParameter != nil {

		// query param hosts.nqn
		var qrHostsNqn string

		if o.HostsNqnQueryParameter != nil {
			qrHostsNqn = *o.HostsNqnQueryParameter
		}
		qHostsNqn := qrHostsNqn
		if qHostsNqn != "" {

			if err := r.SetQueryParam("hosts.nqn", qHostsNqn); err != nil {
				return err
			}
		}
	}

	if o.IoQueueDefaultCountQueryParameter != nil {

		// query param io_queue.default.count
		var qrIoQueueDefaultCount int64

		if o.IoQueueDefaultCountQueryParameter != nil {
			qrIoQueueDefaultCount = *o.IoQueueDefaultCountQueryParameter
		}
		qIoQueueDefaultCount := swag.FormatInt64(qrIoQueueDefaultCount)
		if qIoQueueDefaultCount != "" {

			if err := r.SetQueryParam("io_queue.default.count", qIoQueueDefaultCount); err != nil {
				return err
			}
		}
	}

	if o.IoQueueDefaultDepthQueryParameter != nil {

		// query param io_queue.default.depth
		var qrIoQueueDefaultDepth int64

		if o.IoQueueDefaultDepthQueryParameter != nil {
			qrIoQueueDefaultDepth = *o.IoQueueDefaultDepthQueryParameter
		}
		qIoQueueDefaultDepth := swag.FormatInt64(qrIoQueueDefaultDepth)
		if qIoQueueDefaultDepth != "" {

			if err := r.SetQueryParam("io_queue.default.depth", qIoQueueDefaultDepth); err != nil {
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

	if o.OsTypeQueryParameter != nil {

		// query param os_type
		var qrOsType string

		if o.OsTypeQueryParameter != nil {
			qrOsType = *o.OsTypeQueryParameter
		}
		qOsType := qrOsType
		if qOsType != "" {

			if err := r.SetQueryParam("os_type", qOsType); err != nil {
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

	if o.SerialNumberQueryParameter != nil {

		// query param serial_number
		var qrSerialNumber string

		if o.SerialNumberQueryParameter != nil {
			qrSerialNumber = *o.SerialNumberQueryParameter
		}
		qSerialNumber := qrSerialNumber
		if qSerialNumber != "" {

			if err := r.SetQueryParam("serial_number", qSerialNumber); err != nil {
				return err
			}
		}
	}

	if o.SubsystemMapsAnagrpIDQueryParameter != nil {

		// query param subsystem_maps.anagrpid
		var qrSubsystemMapsAnagrpid string

		if o.SubsystemMapsAnagrpIDQueryParameter != nil {
			qrSubsystemMapsAnagrpid = *o.SubsystemMapsAnagrpIDQueryParameter
		}
		qSubsystemMapsAnagrpid := qrSubsystemMapsAnagrpid
		if qSubsystemMapsAnagrpid != "" {

			if err := r.SetQueryParam("subsystem_maps.anagrpid", qSubsystemMapsAnagrpid); err != nil {
				return err
			}
		}
	}

	if o.SubsystemMapsNamespaceNameQueryParameter != nil {

		// query param subsystem_maps.namespace.name
		var qrSubsystemMapsNamespaceName string

		if o.SubsystemMapsNamespaceNameQueryParameter != nil {
			qrSubsystemMapsNamespaceName = *o.SubsystemMapsNamespaceNameQueryParameter
		}
		qSubsystemMapsNamespaceName := qrSubsystemMapsNamespaceName
		if qSubsystemMapsNamespaceName != "" {

			if err := r.SetQueryParam("subsystem_maps.namespace.name", qSubsystemMapsNamespaceName); err != nil {
				return err
			}
		}
	}

	if o.SubsystemMapsNamespaceUUIDQueryParameter != nil {

		// query param subsystem_maps.namespace.uuid
		var qrSubsystemMapsNamespaceUUID string

		if o.SubsystemMapsNamespaceUUIDQueryParameter != nil {
			qrSubsystemMapsNamespaceUUID = *o.SubsystemMapsNamespaceUUIDQueryParameter
		}
		qSubsystemMapsNamespaceUUID := qrSubsystemMapsNamespaceUUID
		if qSubsystemMapsNamespaceUUID != "" {

			if err := r.SetQueryParam("subsystem_maps.namespace.uuid", qSubsystemMapsNamespaceUUID); err != nil {
				return err
			}
		}
	}

	if o.SubsystemMapsNsIDQueryParameter != nil {

		// query param subsystem_maps.nsid
		var qrSubsystemMapsNsid string

		if o.SubsystemMapsNsIDQueryParameter != nil {
			qrSubsystemMapsNsid = *o.SubsystemMapsNsIDQueryParameter
		}
		qSubsystemMapsNsid := qrSubsystemMapsNsid
		if qSubsystemMapsNsid != "" {

			if err := r.SetQueryParam("subsystem_maps.nsid", qSubsystemMapsNsid); err != nil {
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

	if o.TargetNqnQueryParameter != nil {

		// query param target_nqn
		var qrTargetNqn string

		if o.TargetNqnQueryParameter != nil {
			qrTargetNqn = *o.TargetNqnQueryParameter
		}
		qTargetNqn := qrTargetNqn
		if qTargetNqn != "" {

			if err := r.SetQueryParam("target_nqn", qTargetNqn); err != nil {
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

	if o.VendorUUIDsQueryParameter != nil {

		// query param vendor_uuids
		var qrVendorUuids string

		if o.VendorUUIDsQueryParameter != nil {
			qrVendorUuids = *o.VendorUUIDsQueryParameter
		}
		qVendorUuids := qrVendorUuids
		if qVendorUuids != "" {

			if err := r.SetQueryParam("vendor_uuids", qVendorUuids); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNvmeSubsystemCollectionGet binds the parameter fields
func (o *NvmeSubsystemCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamNvmeSubsystemCollectionGet binds the parameter order_by
func (o *NvmeSubsystemCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
