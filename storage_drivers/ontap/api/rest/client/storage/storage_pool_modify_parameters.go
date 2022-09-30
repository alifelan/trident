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

// NewStoragePoolModifyParams creates a new StoragePoolModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewStoragePoolModifyParams() *StoragePoolModifyParams {
	return &StoragePoolModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewStoragePoolModifyParamsWithTimeout creates a new StoragePoolModifyParams object
// with the ability to set a timeout on a request.
func NewStoragePoolModifyParamsWithTimeout(timeout time.Duration) *StoragePoolModifyParams {
	return &StoragePoolModifyParams{
		timeout: timeout,
	}
}

// NewStoragePoolModifyParamsWithContext creates a new StoragePoolModifyParams object
// with the ability to set a context for a request.
func NewStoragePoolModifyParamsWithContext(ctx context.Context) *StoragePoolModifyParams {
	return &StoragePoolModifyParams{
		Context: ctx,
	}
}

// NewStoragePoolModifyParamsWithHTTPClient creates a new StoragePoolModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewStoragePoolModifyParamsWithHTTPClient(client *http.Client) *StoragePoolModifyParams {
	return &StoragePoolModifyParams{
		HTTPClient: client,
	}
}

/*
StoragePoolModifyParams contains all the parameters to send to the API endpoint

	for the storage pool modify operation.

	Typically these are written to a http.Request.
*/
type StoragePoolModifyParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.StoragePool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	/* UUID.

	   Storage pool UUID.
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the storage pool modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StoragePoolModifyParams) WithDefaults() *StoragePoolModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the storage pool modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *StoragePoolModifyParams) SetDefaults() {
	var (
		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := StoragePoolModifyParams{
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the storage pool modify params
func (o *StoragePoolModifyParams) WithTimeout(timeout time.Duration) *StoragePoolModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the storage pool modify params
func (o *StoragePoolModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the storage pool modify params
func (o *StoragePoolModifyParams) WithContext(ctx context.Context) *StoragePoolModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the storage pool modify params
func (o *StoragePoolModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the storage pool modify params
func (o *StoragePoolModifyParams) WithHTTPClient(client *http.Client) *StoragePoolModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the storage pool modify params
func (o *StoragePoolModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the storage pool modify params
func (o *StoragePoolModifyParams) WithInfo(info *models.StoragePool) *StoragePoolModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the storage pool modify params
func (o *StoragePoolModifyParams) SetInfo(info *models.StoragePool) {
	o.Info = info
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the storage pool modify params
func (o *StoragePoolModifyParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *StoragePoolModifyParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the storage pool modify params
func (o *StoragePoolModifyParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithUUIDPathParameter adds the uuid to the storage pool modify params
func (o *StoragePoolModifyParams) WithUUIDPathParameter(uuid string) *StoragePoolModifyParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the storage pool modify params
func (o *StoragePoolModifyParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *StoragePoolModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
