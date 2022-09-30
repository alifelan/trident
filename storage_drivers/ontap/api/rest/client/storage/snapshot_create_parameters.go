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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewSnapshotCreateParams creates a new SnapshotCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapshotCreateParams() *SnapshotCreateParams {
	return &SnapshotCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotCreateParamsWithTimeout creates a new SnapshotCreateParams object
// with the ability to set a timeout on a request.
func NewSnapshotCreateParamsWithTimeout(timeout time.Duration) *SnapshotCreateParams {
	return &SnapshotCreateParams{
		timeout: timeout,
	}
}

// NewSnapshotCreateParamsWithContext creates a new SnapshotCreateParams object
// with the ability to set a context for a request.
func NewSnapshotCreateParamsWithContext(ctx context.Context) *SnapshotCreateParams {
	return &SnapshotCreateParams{
		Context: ctx,
	}
}

// NewSnapshotCreateParamsWithHTTPClient creates a new SnapshotCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapshotCreateParamsWithHTTPClient(client *http.Client) *SnapshotCreateParams {
	return &SnapshotCreateParams{
		HTTPClient: client,
	}
}

/*
SnapshotCreateParams contains all the parameters to send to the API endpoint

	for the snapshot create operation.

	Typically these are written to a http.Request.
*/
type SnapshotCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.Snapshot

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	/* VolumeUUID.

	   Volume UUID
	*/
	VolumeUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapshot create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotCreateParams) WithDefaults() *SnapshotCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapshot create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)

		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := SnapshotCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snapshot create params
func (o *SnapshotCreateParams) WithTimeout(timeout time.Duration) *SnapshotCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot create params
func (o *SnapshotCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot create params
func (o *SnapshotCreateParams) WithContext(ctx context.Context) *SnapshotCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot create params
func (o *SnapshotCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot create params
func (o *SnapshotCreateParams) WithHTTPClient(client *http.Client) *SnapshotCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot create params
func (o *SnapshotCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the snapshot create params
func (o *SnapshotCreateParams) WithInfo(info *models.Snapshot) *SnapshotCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the snapshot create params
func (o *SnapshotCreateParams) SetInfo(info *models.Snapshot) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the snapshot create params
func (o *SnapshotCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *SnapshotCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the snapshot create params
func (o *SnapshotCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the snapshot create params
func (o *SnapshotCreateParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SnapshotCreateParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the snapshot create params
func (o *SnapshotCreateParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithVolumeUUIDPathParameter adds the volumeUUID to the snapshot create params
func (o *SnapshotCreateParams) WithVolumeUUIDPathParameter(volumeUUID string) *SnapshotCreateParams {
	o.SetVolumeUUIDPathParameter(volumeUUID)
	return o
}

// SetVolumeUUIDPathParameter adds the volumeUuid to the snapshot create params
func (o *SnapshotCreateParams) SetVolumeUUIDPathParameter(volumeUUID string) {
	o.VolumeUUIDPathParameter = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
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

	// path param volume.uuid
	if err := r.SetPathParam("volume.uuid", o.VolumeUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
