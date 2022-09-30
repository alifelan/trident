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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewS3AuditCreateParams creates a new S3AuditCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3AuditCreateParams() *S3AuditCreateParams {
	return &S3AuditCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3AuditCreateParamsWithTimeout creates a new S3AuditCreateParams object
// with the ability to set a timeout on a request.
func NewS3AuditCreateParamsWithTimeout(timeout time.Duration) *S3AuditCreateParams {
	return &S3AuditCreateParams{
		timeout: timeout,
	}
}

// NewS3AuditCreateParamsWithContext creates a new S3AuditCreateParams object
// with the ability to set a context for a request.
func NewS3AuditCreateParamsWithContext(ctx context.Context) *S3AuditCreateParams {
	return &S3AuditCreateParams{
		Context: ctx,
	}
}

// NewS3AuditCreateParamsWithHTTPClient creates a new S3AuditCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3AuditCreateParamsWithHTTPClient(client *http.Client) *S3AuditCreateParams {
	return &S3AuditCreateParams{
		HTTPClient: client,
	}
}

/*
S3AuditCreateParams contains all the parameters to send to the API endpoint

	for the s3 audit create operation.

	Typically these are written to a http.Request.
*/
type S3AuditCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.S3Audit

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 audit create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3AuditCreateParams) WithDefaults() *S3AuditCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 audit create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3AuditCreateParams) SetDefaults() {
	var (
		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := S3AuditCreateParams{
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the s3 audit create params
func (o *S3AuditCreateParams) WithTimeout(timeout time.Duration) *S3AuditCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 audit create params
func (o *S3AuditCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 audit create params
func (o *S3AuditCreateParams) WithContext(ctx context.Context) *S3AuditCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 audit create params
func (o *S3AuditCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 audit create params
func (o *S3AuditCreateParams) WithHTTPClient(client *http.Client) *S3AuditCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 audit create params
func (o *S3AuditCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the s3 audit create params
func (o *S3AuditCreateParams) WithInfo(info *models.S3Audit) *S3AuditCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the s3 audit create params
func (o *S3AuditCreateParams) SetInfo(info *models.S3Audit) {
	o.Info = info
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the s3 audit create params
func (o *S3AuditCreateParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *S3AuditCreateParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the s3 audit create params
func (o *S3AuditCreateParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithSVMUUIDPathParameter adds the svmUUID to the s3 audit create params
func (o *S3AuditCreateParams) WithSVMUUIDPathParameter(svmUUID string) *S3AuditCreateParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the s3 audit create params
func (o *S3AuditCreateParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *S3AuditCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SVMUUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
