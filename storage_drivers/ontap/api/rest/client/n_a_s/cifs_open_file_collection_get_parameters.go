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

// NewCifsOpenFileCollectionGetParams creates a new CifsOpenFileCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsOpenFileCollectionGetParams() *CifsOpenFileCollectionGetParams {
	return &CifsOpenFileCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsOpenFileCollectionGetParamsWithTimeout creates a new CifsOpenFileCollectionGetParams object
// with the ability to set a timeout on a request.
func NewCifsOpenFileCollectionGetParamsWithTimeout(timeout time.Duration) *CifsOpenFileCollectionGetParams {
	return &CifsOpenFileCollectionGetParams{
		timeout: timeout,
	}
}

// NewCifsOpenFileCollectionGetParamsWithContext creates a new CifsOpenFileCollectionGetParams object
// with the ability to set a context for a request.
func NewCifsOpenFileCollectionGetParamsWithContext(ctx context.Context) *CifsOpenFileCollectionGetParams {
	return &CifsOpenFileCollectionGetParams{
		Context: ctx,
	}
}

// NewCifsOpenFileCollectionGetParamsWithHTTPClient creates a new CifsOpenFileCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsOpenFileCollectionGetParamsWithHTTPClient(client *http.Client) *CifsOpenFileCollectionGetParams {
	return &CifsOpenFileCollectionGetParams{
		HTTPClient: client,
	}
}

/*
CifsOpenFileCollectionGetParams contains all the parameters to send to the API endpoint

	for the cifs open file collection get operation.

	Typically these are written to a http.Request.
*/
type CifsOpenFileCollectionGetParams struct {

	/* ConnectionCount.

	   Filter by connection.count
	*/
	ConnectionCountQueryParameter *int64

	/* ConnectionIdentifier.

	   Filter by connection.identifier
	*/
	ConnectionIDentifierQueryParameter *int64

	/* ContinuouslyAvailable.

	   Filter by continuously_available
	*/
	ContinuouslyAvailableQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Identifier.

	   Filter by identifier
	*/
	IdentifierQueryParameter *int64

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

	/* OpenMode.

	   Filter by open_mode
	*/
	OpenModeQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* Path.

	   Filter by path
	*/
	PathQueryParameter *string

	/* RangeLocksCount.

	   Filter by range_locks_count
	*/
	RangeLocksCountQueryParameter *int64

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

	/* SessionIdentifier.

	   Filter by session.identifier
	*/
	SessionIDentifierQueryParameter *int64

	/* ShareMode.

	   Filter by share.mode
	*/
	ShareModeQueryParameter *string

	/* ShareName.

	   Filter by share.name
	*/
	ShareNameQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* Type.

	   Filter by type
	*/
	TypeQueryParameter *string

	/* VolumeName.

	   Filter by volume.name
	*/
	VolumeNameQueryParameter *string

	/* VolumeUUID.

	   Filter by volume.uuid
	*/
	VolumeUUIDQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs open file collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsOpenFileCollectionGetParams) WithDefaults() *CifsOpenFileCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs open file collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsOpenFileCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := CifsOpenFileCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithTimeout(timeout time.Duration) *CifsOpenFileCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithContext(ctx context.Context) *CifsOpenFileCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithHTTPClient(client *http.Client) *CifsOpenFileCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithConnectionCountQueryParameter adds the connectionCount to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithConnectionCountQueryParameter(connectionCount *int64) *CifsOpenFileCollectionGetParams {
	o.SetConnectionCountQueryParameter(connectionCount)
	return o
}

// SetConnectionCountQueryParameter adds the connectionCount to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetConnectionCountQueryParameter(connectionCount *int64) {
	o.ConnectionCountQueryParameter = connectionCount
}

// WithConnectionIDentifierQueryParameter adds the connectionIdentifier to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithConnectionIDentifierQueryParameter(connectionIdentifier *int64) *CifsOpenFileCollectionGetParams {
	o.SetConnectionIDentifierQueryParameter(connectionIdentifier)
	return o
}

// SetConnectionIDentifierQueryParameter adds the connectionIdentifier to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetConnectionIDentifierQueryParameter(connectionIdentifier *int64) {
	o.ConnectionIDentifierQueryParameter = connectionIdentifier
}

// WithContinuouslyAvailableQueryParameter adds the continuouslyAvailable to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithContinuouslyAvailableQueryParameter(continuouslyAvailable *string) *CifsOpenFileCollectionGetParams {
	o.SetContinuouslyAvailableQueryParameter(continuouslyAvailable)
	return o
}

// SetContinuouslyAvailableQueryParameter adds the continuouslyAvailable to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetContinuouslyAvailableQueryParameter(continuouslyAvailable *string) {
	o.ContinuouslyAvailableQueryParameter = continuouslyAvailable
}

// WithFieldsQueryParameter adds the fields to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithFieldsQueryParameter(fields []string) *CifsOpenFileCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIdentifierQueryParameter adds the identifier to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithIdentifierQueryParameter(identifier *int64) *CifsOpenFileCollectionGetParams {
	o.SetIdentifierQueryParameter(identifier)
	return o
}

// SetIdentifierQueryParameter adds the identifier to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetIdentifierQueryParameter(identifier *int64) {
	o.IdentifierQueryParameter = identifier
}

// WithMaxRecordsQueryParameter adds the maxRecords to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *CifsOpenFileCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNodeNameQueryParameter adds the nodeName to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithNodeNameQueryParameter(nodeName *string) *CifsOpenFileCollectionGetParams {
	o.SetNodeNameQueryParameter(nodeName)
	return o
}

// SetNodeNameQueryParameter adds the nodeName to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetNodeNameQueryParameter(nodeName *string) {
	o.NodeNameQueryParameter = nodeName
}

// WithNodeUUIDQueryParameter adds the nodeUUID to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithNodeUUIDQueryParameter(nodeUUID *string) *CifsOpenFileCollectionGetParams {
	o.SetNodeUUIDQueryParameter(nodeUUID)
	return o
}

// SetNodeUUIDQueryParameter adds the nodeUuid to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetNodeUUIDQueryParameter(nodeUUID *string) {
	o.NodeUUIDQueryParameter = nodeUUID
}

// WithOpenModeQueryParameter adds the openMode to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithOpenModeQueryParameter(openMode *string) *CifsOpenFileCollectionGetParams {
	o.SetOpenModeQueryParameter(openMode)
	return o
}

// SetOpenModeQueryParameter adds the openMode to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetOpenModeQueryParameter(openMode *string) {
	o.OpenModeQueryParameter = openMode
}

// WithOrderByQueryParameter adds the orderBy to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *CifsOpenFileCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithPathQueryParameter adds the path to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithPathQueryParameter(path *string) *CifsOpenFileCollectionGetParams {
	o.SetPathQueryParameter(path)
	return o
}

// SetPathQueryParameter adds the path to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetPathQueryParameter(path *string) {
	o.PathQueryParameter = path
}

// WithRangeLocksCountQueryParameter adds the rangeLocksCount to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithRangeLocksCountQueryParameter(rangeLocksCount *int64) *CifsOpenFileCollectionGetParams {
	o.SetRangeLocksCountQueryParameter(rangeLocksCount)
	return o
}

// SetRangeLocksCountQueryParameter adds the rangeLocksCount to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetRangeLocksCountQueryParameter(rangeLocksCount *int64) {
	o.RangeLocksCountQueryParameter = rangeLocksCount
}

// WithReturnRecordsQueryParameter adds the returnRecords to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *CifsOpenFileCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *CifsOpenFileCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSessionIDentifierQueryParameter adds the sessionIdentifier to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithSessionIDentifierQueryParameter(sessionIdentifier *int64) *CifsOpenFileCollectionGetParams {
	o.SetSessionIDentifierQueryParameter(sessionIdentifier)
	return o
}

// SetSessionIDentifierQueryParameter adds the sessionIdentifier to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetSessionIDentifierQueryParameter(sessionIdentifier *int64) {
	o.SessionIDentifierQueryParameter = sessionIdentifier
}

// WithShareModeQueryParameter adds the shareMode to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithShareModeQueryParameter(shareMode *string) *CifsOpenFileCollectionGetParams {
	o.SetShareModeQueryParameter(shareMode)
	return o
}

// SetShareModeQueryParameter adds the shareMode to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetShareModeQueryParameter(shareMode *string) {
	o.ShareModeQueryParameter = shareMode
}

// WithShareNameQueryParameter adds the shareName to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithShareNameQueryParameter(shareName *string) *CifsOpenFileCollectionGetParams {
	o.SetShareNameQueryParameter(shareName)
	return o
}

// SetShareNameQueryParameter adds the shareName to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetShareNameQueryParameter(shareName *string) {
	o.ShareNameQueryParameter = shareName
}

// WithSVMNameQueryParameter adds the svmName to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *CifsOpenFileCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *CifsOpenFileCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithTypeQueryParameter adds the typeVar to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithTypeQueryParameter(typeVar *string) *CifsOpenFileCollectionGetParams {
	o.SetTypeQueryParameter(typeVar)
	return o
}

// SetTypeQueryParameter adds the type to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetTypeQueryParameter(typeVar *string) {
	o.TypeQueryParameter = typeVar
}

// WithVolumeNameQueryParameter adds the volumeName to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithVolumeNameQueryParameter(volumeName *string) *CifsOpenFileCollectionGetParams {
	o.SetVolumeNameQueryParameter(volumeName)
	return o
}

// SetVolumeNameQueryParameter adds the volumeName to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetVolumeNameQueryParameter(volumeName *string) {
	o.VolumeNameQueryParameter = volumeName
}

// WithVolumeUUIDQueryParameter adds the volumeUUID to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) WithVolumeUUIDQueryParameter(volumeUUID *string) *CifsOpenFileCollectionGetParams {
	o.SetVolumeUUIDQueryParameter(volumeUUID)
	return o
}

// SetVolumeUUIDQueryParameter adds the volumeUuid to the cifs open file collection get params
func (o *CifsOpenFileCollectionGetParams) SetVolumeUUIDQueryParameter(volumeUUID *string) {
	o.VolumeUUIDQueryParameter = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CifsOpenFileCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ConnectionCountQueryParameter != nil {

		// query param connection.count
		var qrConnectionCount int64

		if o.ConnectionCountQueryParameter != nil {
			qrConnectionCount = *o.ConnectionCountQueryParameter
		}
		qConnectionCount := swag.FormatInt64(qrConnectionCount)
		if qConnectionCount != "" {

			if err := r.SetQueryParam("connection.count", qConnectionCount); err != nil {
				return err
			}
		}
	}

	if o.ConnectionIDentifierQueryParameter != nil {

		// query param connection.identifier
		var qrConnectionIdentifier int64

		if o.ConnectionIDentifierQueryParameter != nil {
			qrConnectionIdentifier = *o.ConnectionIDentifierQueryParameter
		}
		qConnectionIdentifier := swag.FormatInt64(qrConnectionIdentifier)
		if qConnectionIdentifier != "" {

			if err := r.SetQueryParam("connection.identifier", qConnectionIdentifier); err != nil {
				return err
			}
		}
	}

	if o.ContinuouslyAvailableQueryParameter != nil {

		// query param continuously_available
		var qrContinuouslyAvailable string

		if o.ContinuouslyAvailableQueryParameter != nil {
			qrContinuouslyAvailable = *o.ContinuouslyAvailableQueryParameter
		}
		qContinuouslyAvailable := qrContinuouslyAvailable
		if qContinuouslyAvailable != "" {

			if err := r.SetQueryParam("continuously_available", qContinuouslyAvailable); err != nil {
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

	if o.IdentifierQueryParameter != nil {

		// query param identifier
		var qrIdentifier int64

		if o.IdentifierQueryParameter != nil {
			qrIdentifier = *o.IdentifierQueryParameter
		}
		qIdentifier := swag.FormatInt64(qrIdentifier)
		if qIdentifier != "" {

			if err := r.SetQueryParam("identifier", qIdentifier); err != nil {
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

	if o.OpenModeQueryParameter != nil {

		// query param open_mode
		var qrOpenMode string

		if o.OpenModeQueryParameter != nil {
			qrOpenMode = *o.OpenModeQueryParameter
		}
		qOpenMode := qrOpenMode
		if qOpenMode != "" {

			if err := r.SetQueryParam("open_mode", qOpenMode); err != nil {
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

	if o.PathQueryParameter != nil {

		// query param path
		var qrPath string

		if o.PathQueryParameter != nil {
			qrPath = *o.PathQueryParameter
		}
		qPath := qrPath
		if qPath != "" {

			if err := r.SetQueryParam("path", qPath); err != nil {
				return err
			}
		}
	}

	if o.RangeLocksCountQueryParameter != nil {

		// query param range_locks_count
		var qrRangeLocksCount int64

		if o.RangeLocksCountQueryParameter != nil {
			qrRangeLocksCount = *o.RangeLocksCountQueryParameter
		}
		qRangeLocksCount := swag.FormatInt64(qrRangeLocksCount)
		if qRangeLocksCount != "" {

			if err := r.SetQueryParam("range_locks_count", qRangeLocksCount); err != nil {
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

	if o.SessionIDentifierQueryParameter != nil {

		// query param session.identifier
		var qrSessionIdentifier int64

		if o.SessionIDentifierQueryParameter != nil {
			qrSessionIdentifier = *o.SessionIDentifierQueryParameter
		}
		qSessionIdentifier := swag.FormatInt64(qrSessionIdentifier)
		if qSessionIdentifier != "" {

			if err := r.SetQueryParam("session.identifier", qSessionIdentifier); err != nil {
				return err
			}
		}
	}

	if o.ShareModeQueryParameter != nil {

		// query param share.mode
		var qrShareMode string

		if o.ShareModeQueryParameter != nil {
			qrShareMode = *o.ShareModeQueryParameter
		}
		qShareMode := qrShareMode
		if qShareMode != "" {

			if err := r.SetQueryParam("share.mode", qShareMode); err != nil {
				return err
			}
		}
	}

	if o.ShareNameQueryParameter != nil {

		// query param share.name
		var qrShareName string

		if o.ShareNameQueryParameter != nil {
			qrShareName = *o.ShareNameQueryParameter
		}
		qShareName := qrShareName
		if qShareName != "" {

			if err := r.SetQueryParam("share.name", qShareName); err != nil {
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

	if o.VolumeNameQueryParameter != nil {

		// query param volume.name
		var qrVolumeName string

		if o.VolumeNameQueryParameter != nil {
			qrVolumeName = *o.VolumeNameQueryParameter
		}
		qVolumeName := qrVolumeName
		if qVolumeName != "" {

			if err := r.SetQueryParam("volume.name", qVolumeName); err != nil {
				return err
			}
		}
	}

	if o.VolumeUUIDQueryParameter != nil {

		// query param volume.uuid
		var qrVolumeUUID string

		if o.VolumeUUIDQueryParameter != nil {
			qrVolumeUUID = *o.VolumeUUIDQueryParameter
		}
		qVolumeUUID := qrVolumeUUID
		if qVolumeUUID != "" {

			if err := r.SetQueryParam("volume.uuid", qVolumeUUID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamCifsOpenFileCollectionGet binds the parameter fields
func (o *CifsOpenFileCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamCifsOpenFileCollectionGet binds the parameter order_by
func (o *CifsOpenFileCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
