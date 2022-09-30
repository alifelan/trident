// Code generated by go-swagger; DO NOT EDIT.

package object_store

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

// NewS3BucketCreateParams creates a new S3BucketCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3BucketCreateParams() *S3BucketCreateParams {
	return &S3BucketCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3BucketCreateParamsWithTimeout creates a new S3BucketCreateParams object
// with the ability to set a timeout on a request.
func NewS3BucketCreateParamsWithTimeout(timeout time.Duration) *S3BucketCreateParams {
	return &S3BucketCreateParams{
		timeout: timeout,
	}
}

// NewS3BucketCreateParamsWithContext creates a new S3BucketCreateParams object
// with the ability to set a context for a request.
func NewS3BucketCreateParamsWithContext(ctx context.Context) *S3BucketCreateParams {
	return &S3BucketCreateParams{
		Context: ctx,
	}
}

// NewS3BucketCreateParamsWithHTTPClient creates a new S3BucketCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3BucketCreateParamsWithHTTPClient(client *http.Client) *S3BucketCreateParams {
	return &S3BucketCreateParams{
		HTTPClient: client,
	}
}

/*
S3BucketCreateParams contains all the parameters to send to the API endpoint

	for the s3 bucket create operation.

	Typically these are written to a http.Request.
*/
type S3BucketCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.S3Bucket

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

// WithDefaults hydrates default values in the s3 bucket create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3BucketCreateParams) WithDefaults() *S3BucketCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 bucket create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3BucketCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)

		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := S3BucketCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the s3 bucket create params
func (o *S3BucketCreateParams) WithTimeout(timeout time.Duration) *S3BucketCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 bucket create params
func (o *S3BucketCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 bucket create params
func (o *S3BucketCreateParams) WithContext(ctx context.Context) *S3BucketCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 bucket create params
func (o *S3BucketCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 bucket create params
func (o *S3BucketCreateParams) WithHTTPClient(client *http.Client) *S3BucketCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 bucket create params
func (o *S3BucketCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the s3 bucket create params
func (o *S3BucketCreateParams) WithInfo(info *models.S3Bucket) *S3BucketCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the s3 bucket create params
func (o *S3BucketCreateParams) SetInfo(info *models.S3Bucket) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the s3 bucket create params
func (o *S3BucketCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *S3BucketCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the s3 bucket create params
func (o *S3BucketCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the s3 bucket create params
func (o *S3BucketCreateParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *S3BucketCreateParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the s3 bucket create params
func (o *S3BucketCreateParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *S3BucketCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
