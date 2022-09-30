// Code generated by go-swagger; DO NOT EDIT.

package application

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

// NewApplicationSnapshotCreateParams creates a new ApplicationSnapshotCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewApplicationSnapshotCreateParams() *ApplicationSnapshotCreateParams {
	return &ApplicationSnapshotCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewApplicationSnapshotCreateParamsWithTimeout creates a new ApplicationSnapshotCreateParams object
// with the ability to set a timeout on a request.
func NewApplicationSnapshotCreateParamsWithTimeout(timeout time.Duration) *ApplicationSnapshotCreateParams {
	return &ApplicationSnapshotCreateParams{
		timeout: timeout,
	}
}

// NewApplicationSnapshotCreateParamsWithContext creates a new ApplicationSnapshotCreateParams object
// with the ability to set a context for a request.
func NewApplicationSnapshotCreateParamsWithContext(ctx context.Context) *ApplicationSnapshotCreateParams {
	return &ApplicationSnapshotCreateParams{
		Context: ctx,
	}
}

// NewApplicationSnapshotCreateParamsWithHTTPClient creates a new ApplicationSnapshotCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewApplicationSnapshotCreateParamsWithHTTPClient(client *http.Client) *ApplicationSnapshotCreateParams {
	return &ApplicationSnapshotCreateParams{
		HTTPClient: client,
	}
}

/*
ApplicationSnapshotCreateParams contains all the parameters to send to the API endpoint

	for the application snapshot create operation.

	Typically these are written to a http.Request.
*/
type ApplicationSnapshotCreateParams struct {

	/* ApplicationUUID.

	   Application UUID
	*/
	ApplicationUUIDPathParameter string

	/* Info.

	   Info specification
	*/
	Info *models.ApplicationSnapshot

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the application snapshot create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationSnapshotCreateParams) WithDefaults() *ApplicationSnapshotCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the application snapshot create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ApplicationSnapshotCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)

		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := ApplicationSnapshotCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the application snapshot create params
func (o *ApplicationSnapshotCreateParams) WithTimeout(timeout time.Duration) *ApplicationSnapshotCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the application snapshot create params
func (o *ApplicationSnapshotCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the application snapshot create params
func (o *ApplicationSnapshotCreateParams) WithContext(ctx context.Context) *ApplicationSnapshotCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the application snapshot create params
func (o *ApplicationSnapshotCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the application snapshot create params
func (o *ApplicationSnapshotCreateParams) WithHTTPClient(client *http.Client) *ApplicationSnapshotCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the application snapshot create params
func (o *ApplicationSnapshotCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationUUIDPathParameter adds the applicationUUID to the application snapshot create params
func (o *ApplicationSnapshotCreateParams) WithApplicationUUIDPathParameter(applicationUUID string) *ApplicationSnapshotCreateParams {
	o.SetApplicationUUIDPathParameter(applicationUUID)
	return o
}

// SetApplicationUUIDPathParameter adds the applicationUuid to the application snapshot create params
func (o *ApplicationSnapshotCreateParams) SetApplicationUUIDPathParameter(applicationUUID string) {
	o.ApplicationUUIDPathParameter = applicationUUID
}

// WithInfo adds the info to the application snapshot create params
func (o *ApplicationSnapshotCreateParams) WithInfo(info *models.ApplicationSnapshot) *ApplicationSnapshotCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the application snapshot create params
func (o *ApplicationSnapshotCreateParams) SetInfo(info *models.ApplicationSnapshot) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the application snapshot create params
func (o *ApplicationSnapshotCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *ApplicationSnapshotCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the application snapshot create params
func (o *ApplicationSnapshotCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the application snapshot create params
func (o *ApplicationSnapshotCreateParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *ApplicationSnapshotCreateParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the application snapshot create params
func (o *ApplicationSnapshotCreateParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *ApplicationSnapshotCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param application.uuid
	if err := r.SetPathParam("application.uuid", o.ApplicationUUIDPathParameter); err != nil {
		return err
	}
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
