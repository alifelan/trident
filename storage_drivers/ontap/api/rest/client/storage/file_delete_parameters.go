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

// NewFileDeleteParams creates a new FileDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFileDeleteParams() *FileDeleteParams {
	return &FileDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFileDeleteParamsWithTimeout creates a new FileDeleteParams object
// with the ability to set a timeout on a request.
func NewFileDeleteParamsWithTimeout(timeout time.Duration) *FileDeleteParams {
	return &FileDeleteParams{
		timeout: timeout,
	}
}

// NewFileDeleteParamsWithContext creates a new FileDeleteParams object
// with the ability to set a context for a request.
func NewFileDeleteParamsWithContext(ctx context.Context) *FileDeleteParams {
	return &FileDeleteParams{
		Context: ctx,
	}
}

// NewFileDeleteParamsWithHTTPClient creates a new FileDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewFileDeleteParamsWithHTTPClient(client *http.Client) *FileDeleteParams {
	return &FileDeleteParams{
		HTTPClient: client,
	}
}

/*
FileDeleteParams contains all the parameters to send to the API endpoint

	for the file delete operation.

	Typically these are written to a http.Request.
*/
type FileDeleteParams struct {

	/* Path.

	   The relative path of a directory in the volume. The path field requires using "%2E" to represent "." and "%2F" to represent "/" for the path provided.
	*/
	PathPathParameter string

	/* Recurse.

	   Delete an entire directory. The behaviour of this call is equivalent to rm -rf.
	*/
	RecurseQueryParameter *bool

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	/* ThrottleDeletion.

	   The maximum number of directory delete operations per second. A valid throttle_deletion number is an integer from 10 to 100000.
	*/
	ThrottleDeletionQueryParameter *int64

	/* VolumeUUID.

	   Volume UUID
	*/
	VolumeUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the file delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileDeleteParams) WithDefaults() *FileDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the file delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FileDeleteParams) SetDefaults() {
	var (
		recurseQueryParameterDefault = bool(false)

		returnRecordsQueryParameterDefault = bool(false)

		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := FileDeleteParams{
		RecurseQueryParameter:       &recurseQueryParameterDefault,
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the file delete params
func (o *FileDeleteParams) WithTimeout(timeout time.Duration) *FileDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the file delete params
func (o *FileDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the file delete params
func (o *FileDeleteParams) WithContext(ctx context.Context) *FileDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the file delete params
func (o *FileDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the file delete params
func (o *FileDeleteParams) WithHTTPClient(client *http.Client) *FileDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the file delete params
func (o *FileDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPathPathParameter adds the path to the file delete params
func (o *FileDeleteParams) WithPathPathParameter(path string) *FileDeleteParams {
	o.SetPathPathParameter(path)
	return o
}

// SetPathPathParameter adds the path to the file delete params
func (o *FileDeleteParams) SetPathPathParameter(path string) {
	o.PathPathParameter = path
}

// WithRecurseQueryParameter adds the recurse to the file delete params
func (o *FileDeleteParams) WithRecurseQueryParameter(recurse *bool) *FileDeleteParams {
	o.SetRecurseQueryParameter(recurse)
	return o
}

// SetRecurseQueryParameter adds the recurse to the file delete params
func (o *FileDeleteParams) SetRecurseQueryParameter(recurse *bool) {
	o.RecurseQueryParameter = recurse
}

// WithReturnRecordsQueryParameter adds the returnRecords to the file delete params
func (o *FileDeleteParams) WithReturnRecordsQueryParameter(returnRecords *bool) *FileDeleteParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the file delete params
func (o *FileDeleteParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the file delete params
func (o *FileDeleteParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *FileDeleteParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the file delete params
func (o *FileDeleteParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithThrottleDeletionQueryParameter adds the throttleDeletion to the file delete params
func (o *FileDeleteParams) WithThrottleDeletionQueryParameter(throttleDeletion *int64) *FileDeleteParams {
	o.SetThrottleDeletionQueryParameter(throttleDeletion)
	return o
}

// SetThrottleDeletionQueryParameter adds the throttleDeletion to the file delete params
func (o *FileDeleteParams) SetThrottleDeletionQueryParameter(throttleDeletion *int64) {
	o.ThrottleDeletionQueryParameter = throttleDeletion
}

// WithVolumeUUIDPathParameter adds the volumeUUID to the file delete params
func (o *FileDeleteParams) WithVolumeUUIDPathParameter(volumeUUID string) *FileDeleteParams {
	o.SetVolumeUUIDPathParameter(volumeUUID)
	return o
}

// SetVolumeUUIDPathParameter adds the volumeUuid to the file delete params
func (o *FileDeleteParams) SetVolumeUUIDPathParameter(volumeUUID string) {
	o.VolumeUUIDPathParameter = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *FileDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param path
	if err := r.SetPathParam("path", o.PathPathParameter); err != nil {
		return err
	}

	if o.RecurseQueryParameter != nil {

		// query param recurse
		var qrRecurse bool

		if o.RecurseQueryParameter != nil {
			qrRecurse = *o.RecurseQueryParameter
		}
		qRecurse := swag.FormatBool(qrRecurse)
		if qRecurse != "" {

			if err := r.SetQueryParam("recurse", qRecurse); err != nil {
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

	if o.ThrottleDeletionQueryParameter != nil {

		// query param throttle_deletion
		var qrThrottleDeletion int64

		if o.ThrottleDeletionQueryParameter != nil {
			qrThrottleDeletion = *o.ThrottleDeletionQueryParameter
		}
		qThrottleDeletion := swag.FormatInt64(qrThrottleDeletion)
		if qThrottleDeletion != "" {

			if err := r.SetQueryParam("throttle_deletion", qThrottleDeletion); err != nil {
				return err
			}
		}
	}

	// path param volume.uuid
	if err := r.SetPathParam("volume.uuid", o.VolumeUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
