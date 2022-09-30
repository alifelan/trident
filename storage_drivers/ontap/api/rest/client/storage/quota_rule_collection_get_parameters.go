// Code generated by go-swagger; DO NOT EDIT.

package storage

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

// NewQuotaRuleCollectionGetParams creates a new QuotaRuleCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQuotaRuleCollectionGetParams() *QuotaRuleCollectionGetParams {
	return &QuotaRuleCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQuotaRuleCollectionGetParamsWithTimeout creates a new QuotaRuleCollectionGetParams object
// with the ability to set a timeout on a request.
func NewQuotaRuleCollectionGetParamsWithTimeout(timeout time.Duration) *QuotaRuleCollectionGetParams {
	return &QuotaRuleCollectionGetParams{
		timeout: timeout,
	}
}

// NewQuotaRuleCollectionGetParamsWithContext creates a new QuotaRuleCollectionGetParams object
// with the ability to set a context for a request.
func NewQuotaRuleCollectionGetParamsWithContext(ctx context.Context) *QuotaRuleCollectionGetParams {
	return &QuotaRuleCollectionGetParams{
		Context: ctx,
	}
}

// NewQuotaRuleCollectionGetParamsWithHTTPClient creates a new QuotaRuleCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewQuotaRuleCollectionGetParamsWithHTTPClient(client *http.Client) *QuotaRuleCollectionGetParams {
	return &QuotaRuleCollectionGetParams{
		HTTPClient: client,
	}
}

/*
QuotaRuleCollectionGetParams contains all the parameters to send to the API endpoint

	for the quota rule collection get operation.

	Typically these are written to a http.Request.
*/
type QuotaRuleCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* FilesHardLimit.

	   Filter by files.hard_limit
	*/
	FilesHardLimitQueryParameter *int64

	/* FilesSoftLimit.

	   Filter by files.soft_limit
	*/
	FilesSoftLimitQueryParameter *int64

	/* GroupID.

	   Filter by group.id
	*/
	GroupIDQueryParameter *string

	/* GroupName.

	   Filter by group.name
	*/
	GroupNameQueryParameter *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* QtreeID.

	   Filter by qtree.id
	*/
	QtreeIDQueryParameter *int64

	/* QtreeName.

	   Filter by qtree.name
	*/
	QtreeNameQueryParameter *string

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

	/* SpaceHardLimit.

	   Filter by space.hard_limit
	*/
	SpaceHardLimitQueryParameter *int64

	/* SpaceSoftLimit.

	   Filter by space.soft_limit
	*/
	SpaceSoftLimitQueryParameter *int64

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

	/* UserMapping.

	   Filter by user_mapping
	*/
	UserMappingQueryParameter *bool

	/* UsersID.

	   Filter by users.id
	*/
	UsersIDQueryParameter *string

	/* UsersName.

	   Filter by users.name
	*/
	UsersNameQueryParameter *string

	/* UUID.

	   Filter by uuid
	*/
	UUIDQueryParameter *string

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

// WithDefaults hydrates default values in the quota rule collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuotaRuleCollectionGetParams) WithDefaults() *QuotaRuleCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the quota rule collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuotaRuleCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := QuotaRuleCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithTimeout(timeout time.Duration) *QuotaRuleCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithContext(ctx context.Context) *QuotaRuleCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithHTTPClient(client *http.Client) *QuotaRuleCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithFieldsQueryParameter(fields []string) *QuotaRuleCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithFilesHardLimitQueryParameter adds the filesHardLimit to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithFilesHardLimitQueryParameter(filesHardLimit *int64) *QuotaRuleCollectionGetParams {
	o.SetFilesHardLimitQueryParameter(filesHardLimit)
	return o
}

// SetFilesHardLimitQueryParameter adds the filesHardLimit to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetFilesHardLimitQueryParameter(filesHardLimit *int64) {
	o.FilesHardLimitQueryParameter = filesHardLimit
}

// WithFilesSoftLimitQueryParameter adds the filesSoftLimit to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithFilesSoftLimitQueryParameter(filesSoftLimit *int64) *QuotaRuleCollectionGetParams {
	o.SetFilesSoftLimitQueryParameter(filesSoftLimit)
	return o
}

// SetFilesSoftLimitQueryParameter adds the filesSoftLimit to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetFilesSoftLimitQueryParameter(filesSoftLimit *int64) {
	o.FilesSoftLimitQueryParameter = filesSoftLimit
}

// WithGroupIDQueryParameter adds the groupID to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithGroupIDQueryParameter(groupID *string) *QuotaRuleCollectionGetParams {
	o.SetGroupIDQueryParameter(groupID)
	return o
}

// SetGroupIDQueryParameter adds the groupId to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetGroupIDQueryParameter(groupID *string) {
	o.GroupIDQueryParameter = groupID
}

// WithGroupNameQueryParameter adds the groupName to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithGroupNameQueryParameter(groupName *string) *QuotaRuleCollectionGetParams {
	o.SetGroupNameQueryParameter(groupName)
	return o
}

// SetGroupNameQueryParameter adds the groupName to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetGroupNameQueryParameter(groupName *string) {
	o.GroupNameQueryParameter = groupName
}

// WithMaxRecordsQueryParameter adds the maxRecords to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *QuotaRuleCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *QuotaRuleCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithQtreeIDQueryParameter adds the qtreeID to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithQtreeIDQueryParameter(qtreeID *int64) *QuotaRuleCollectionGetParams {
	o.SetQtreeIDQueryParameter(qtreeID)
	return o
}

// SetQtreeIDQueryParameter adds the qtreeId to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetQtreeIDQueryParameter(qtreeID *int64) {
	o.QtreeIDQueryParameter = qtreeID
}

// WithQtreeNameQueryParameter adds the qtreeName to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithQtreeNameQueryParameter(qtreeName *string) *QuotaRuleCollectionGetParams {
	o.SetQtreeNameQueryParameter(qtreeName)
	return o
}

// SetQtreeNameQueryParameter adds the qtreeName to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetQtreeNameQueryParameter(qtreeName *string) {
	o.QtreeNameQueryParameter = qtreeName
}

// WithReturnRecordsQueryParameter adds the returnRecords to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *QuotaRuleCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *QuotaRuleCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSpaceHardLimitQueryParameter adds the spaceHardLimit to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithSpaceHardLimitQueryParameter(spaceHardLimit *int64) *QuotaRuleCollectionGetParams {
	o.SetSpaceHardLimitQueryParameter(spaceHardLimit)
	return o
}

// SetSpaceHardLimitQueryParameter adds the spaceHardLimit to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetSpaceHardLimitQueryParameter(spaceHardLimit *int64) {
	o.SpaceHardLimitQueryParameter = spaceHardLimit
}

// WithSpaceSoftLimitQueryParameter adds the spaceSoftLimit to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithSpaceSoftLimitQueryParameter(spaceSoftLimit *int64) *QuotaRuleCollectionGetParams {
	o.SetSpaceSoftLimitQueryParameter(spaceSoftLimit)
	return o
}

// SetSpaceSoftLimitQueryParameter adds the spaceSoftLimit to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetSpaceSoftLimitQueryParameter(spaceSoftLimit *int64) {
	o.SpaceSoftLimitQueryParameter = spaceSoftLimit
}

// WithSVMNameQueryParameter adds the svmName to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *QuotaRuleCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *QuotaRuleCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithTypeQueryParameter adds the typeVar to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithTypeQueryParameter(typeVar *string) *QuotaRuleCollectionGetParams {
	o.SetTypeQueryParameter(typeVar)
	return o
}

// SetTypeQueryParameter adds the type to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetTypeQueryParameter(typeVar *string) {
	o.TypeQueryParameter = typeVar
}

// WithUserMappingQueryParameter adds the userMapping to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithUserMappingQueryParameter(userMapping *bool) *QuotaRuleCollectionGetParams {
	o.SetUserMappingQueryParameter(userMapping)
	return o
}

// SetUserMappingQueryParameter adds the userMapping to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetUserMappingQueryParameter(userMapping *bool) {
	o.UserMappingQueryParameter = userMapping
}

// WithUsersIDQueryParameter adds the usersID to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithUsersIDQueryParameter(usersID *string) *QuotaRuleCollectionGetParams {
	o.SetUsersIDQueryParameter(usersID)
	return o
}

// SetUsersIDQueryParameter adds the usersId to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetUsersIDQueryParameter(usersID *string) {
	o.UsersIDQueryParameter = usersID
}

// WithUsersNameQueryParameter adds the usersName to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithUsersNameQueryParameter(usersName *string) *QuotaRuleCollectionGetParams {
	o.SetUsersNameQueryParameter(usersName)
	return o
}

// SetUsersNameQueryParameter adds the usersName to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetUsersNameQueryParameter(usersName *string) {
	o.UsersNameQueryParameter = usersName
}

// WithUUIDQueryParameter adds the uuid to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithUUIDQueryParameter(uuid *string) *QuotaRuleCollectionGetParams {
	o.SetUUIDQueryParameter(uuid)
	return o
}

// SetUUIDQueryParameter adds the uuid to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetUUIDQueryParameter(uuid *string) {
	o.UUIDQueryParameter = uuid
}

// WithVolumeNameQueryParameter adds the volumeName to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithVolumeNameQueryParameter(volumeName *string) *QuotaRuleCollectionGetParams {
	o.SetVolumeNameQueryParameter(volumeName)
	return o
}

// SetVolumeNameQueryParameter adds the volumeName to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetVolumeNameQueryParameter(volumeName *string) {
	o.VolumeNameQueryParameter = volumeName
}

// WithVolumeUUIDQueryParameter adds the volumeUUID to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) WithVolumeUUIDQueryParameter(volumeUUID *string) *QuotaRuleCollectionGetParams {
	o.SetVolumeUUIDQueryParameter(volumeUUID)
	return o
}

// SetVolumeUUIDQueryParameter adds the volumeUuid to the quota rule collection get params
func (o *QuotaRuleCollectionGetParams) SetVolumeUUIDQueryParameter(volumeUUID *string) {
	o.VolumeUUIDQueryParameter = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *QuotaRuleCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.FilesHardLimitQueryParameter != nil {

		// query param files.hard_limit
		var qrFilesHardLimit int64

		if o.FilesHardLimitQueryParameter != nil {
			qrFilesHardLimit = *o.FilesHardLimitQueryParameter
		}
		qFilesHardLimit := swag.FormatInt64(qrFilesHardLimit)
		if qFilesHardLimit != "" {

			if err := r.SetQueryParam("files.hard_limit", qFilesHardLimit); err != nil {
				return err
			}
		}
	}

	if o.FilesSoftLimitQueryParameter != nil {

		// query param files.soft_limit
		var qrFilesSoftLimit int64

		if o.FilesSoftLimitQueryParameter != nil {
			qrFilesSoftLimit = *o.FilesSoftLimitQueryParameter
		}
		qFilesSoftLimit := swag.FormatInt64(qrFilesSoftLimit)
		if qFilesSoftLimit != "" {

			if err := r.SetQueryParam("files.soft_limit", qFilesSoftLimit); err != nil {
				return err
			}
		}
	}

	if o.GroupIDQueryParameter != nil {

		// query param group.id
		var qrGroupID string

		if o.GroupIDQueryParameter != nil {
			qrGroupID = *o.GroupIDQueryParameter
		}
		qGroupID := qrGroupID
		if qGroupID != "" {

			if err := r.SetQueryParam("group.id", qGroupID); err != nil {
				return err
			}
		}
	}

	if o.GroupNameQueryParameter != nil {

		// query param group.name
		var qrGroupName string

		if o.GroupNameQueryParameter != nil {
			qrGroupName = *o.GroupNameQueryParameter
		}
		qGroupName := qrGroupName
		if qGroupName != "" {

			if err := r.SetQueryParam("group.name", qGroupName); err != nil {
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

	if o.QtreeIDQueryParameter != nil {

		// query param qtree.id
		var qrQtreeID int64

		if o.QtreeIDQueryParameter != nil {
			qrQtreeID = *o.QtreeIDQueryParameter
		}
		qQtreeID := swag.FormatInt64(qrQtreeID)
		if qQtreeID != "" {

			if err := r.SetQueryParam("qtree.id", qQtreeID); err != nil {
				return err
			}
		}
	}

	if o.QtreeNameQueryParameter != nil {

		// query param qtree.name
		var qrQtreeName string

		if o.QtreeNameQueryParameter != nil {
			qrQtreeName = *o.QtreeNameQueryParameter
		}
		qQtreeName := qrQtreeName
		if qQtreeName != "" {

			if err := r.SetQueryParam("qtree.name", qQtreeName); err != nil {
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

	if o.SpaceHardLimitQueryParameter != nil {

		// query param space.hard_limit
		var qrSpaceHardLimit int64

		if o.SpaceHardLimitQueryParameter != nil {
			qrSpaceHardLimit = *o.SpaceHardLimitQueryParameter
		}
		qSpaceHardLimit := swag.FormatInt64(qrSpaceHardLimit)
		if qSpaceHardLimit != "" {

			if err := r.SetQueryParam("space.hard_limit", qSpaceHardLimit); err != nil {
				return err
			}
		}
	}

	if o.SpaceSoftLimitQueryParameter != nil {

		// query param space.soft_limit
		var qrSpaceSoftLimit int64

		if o.SpaceSoftLimitQueryParameter != nil {
			qrSpaceSoftLimit = *o.SpaceSoftLimitQueryParameter
		}
		qSpaceSoftLimit := swag.FormatInt64(qrSpaceSoftLimit)
		if qSpaceSoftLimit != "" {

			if err := r.SetQueryParam("space.soft_limit", qSpaceSoftLimit); err != nil {
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

	if o.UserMappingQueryParameter != nil {

		// query param user_mapping
		var qrUserMapping bool

		if o.UserMappingQueryParameter != nil {
			qrUserMapping = *o.UserMappingQueryParameter
		}
		qUserMapping := swag.FormatBool(qrUserMapping)
		if qUserMapping != "" {

			if err := r.SetQueryParam("user_mapping", qUserMapping); err != nil {
				return err
			}
		}
	}

	if o.UsersIDQueryParameter != nil {

		// query param users.id
		var qrUsersID string

		if o.UsersIDQueryParameter != nil {
			qrUsersID = *o.UsersIDQueryParameter
		}
		qUsersID := qrUsersID
		if qUsersID != "" {

			if err := r.SetQueryParam("users.id", qUsersID); err != nil {
				return err
			}
		}
	}

	if o.UsersNameQueryParameter != nil {

		// query param users.name
		var qrUsersName string

		if o.UsersNameQueryParameter != nil {
			qrUsersName = *o.UsersNameQueryParameter
		}
		qUsersName := qrUsersName
		if qUsersName != "" {

			if err := r.SetQueryParam("users.name", qUsersName); err != nil {
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

// bindParamQuotaRuleCollectionGet binds the parameter fields
func (o *QuotaRuleCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamQuotaRuleCollectionGet binds the parameter order_by
func (o *QuotaRuleCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
