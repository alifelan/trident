// Code generated by go-swagger; DO NOT EDIT.

package cloud

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

// NewCloudTargetModifyParams creates a new CloudTargetModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCloudTargetModifyParams() *CloudTargetModifyParams {
	return &CloudTargetModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCloudTargetModifyParamsWithTimeout creates a new CloudTargetModifyParams object
// with the ability to set a timeout on a request.
func NewCloudTargetModifyParamsWithTimeout(timeout time.Duration) *CloudTargetModifyParams {
	return &CloudTargetModifyParams{
		timeout: timeout,
	}
}

// NewCloudTargetModifyParamsWithContext creates a new CloudTargetModifyParams object
// with the ability to set a context for a request.
func NewCloudTargetModifyParamsWithContext(ctx context.Context) *CloudTargetModifyParams {
	return &CloudTargetModifyParams{
		Context: ctx,
	}
}

// NewCloudTargetModifyParamsWithHTTPClient creates a new CloudTargetModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewCloudTargetModifyParamsWithHTTPClient(client *http.Client) *CloudTargetModifyParams {
	return &CloudTargetModifyParams{
		HTTPClient: client,
	}
}

/*
CloudTargetModifyParams contains all the parameters to send to the API endpoint

	for the cloud target modify operation.

	Typically these are written to a http.Request.
*/
type CloudTargetModifyParams struct {

	/* CheckOnly.

	   Do not modify the configuration, only check that the PATCH request succeeds.
	*/
	CheckOnlyQueryParameter *bool

	/* IgnoreWarnings.

	   Specifies whether or not warnings should be ignored.
	*/
	IgnoreWarningsQueryParameter *bool

	/* Info.

	   Info specification
	*/
	Info *models.CloudTarget

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	/* UUID.

	   Cloud target UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cloud target modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudTargetModifyParams) WithDefaults() *CloudTargetModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cloud target modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CloudTargetModifyParams) SetDefaults() {
	var (
		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := CloudTargetModifyParams{
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cloud target modify params
func (o *CloudTargetModifyParams) WithTimeout(timeout time.Duration) *CloudTargetModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cloud target modify params
func (o *CloudTargetModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cloud target modify params
func (o *CloudTargetModifyParams) WithContext(ctx context.Context) *CloudTargetModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cloud target modify params
func (o *CloudTargetModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cloud target modify params
func (o *CloudTargetModifyParams) WithHTTPClient(client *http.Client) *CloudTargetModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cloud target modify params
func (o *CloudTargetModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCheckOnlyQueryParameter adds the checkOnly to the cloud target modify params
func (o *CloudTargetModifyParams) WithCheckOnlyQueryParameter(checkOnly *bool) *CloudTargetModifyParams {
	o.SetCheckOnlyQueryParameter(checkOnly)
	return o
}

// SetCheckOnlyQueryParameter adds the checkOnly to the cloud target modify params
func (o *CloudTargetModifyParams) SetCheckOnlyQueryParameter(checkOnly *bool) {
	o.CheckOnlyQueryParameter = checkOnly
}

// WithIgnoreWarningsQueryParameter adds the ignoreWarnings to the cloud target modify params
func (o *CloudTargetModifyParams) WithIgnoreWarningsQueryParameter(ignoreWarnings *bool) *CloudTargetModifyParams {
	o.SetIgnoreWarningsQueryParameter(ignoreWarnings)
	return o
}

// SetIgnoreWarningsQueryParameter adds the ignoreWarnings to the cloud target modify params
func (o *CloudTargetModifyParams) SetIgnoreWarningsQueryParameter(ignoreWarnings *bool) {
	o.IgnoreWarningsQueryParameter = ignoreWarnings
}

// WithInfo adds the info to the cloud target modify params
func (o *CloudTargetModifyParams) WithInfo(info *models.CloudTarget) *CloudTargetModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cloud target modify params
func (o *CloudTargetModifyParams) SetInfo(info *models.CloudTarget) {
	o.Info = info
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the cloud target modify params
func (o *CloudTargetModifyParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *CloudTargetModifyParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the cloud target modify params
func (o *CloudTargetModifyParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithUUIDPathParameter adds the uuid to the cloud target modify params
func (o *CloudTargetModifyParams) WithUUIDPathParameter(uuid string) *CloudTargetModifyParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the cloud target modify params
func (o *CloudTargetModifyParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *CloudTargetModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CheckOnlyQueryParameter != nil {

		// query param check_only
		var qrCheckOnly bool

		if o.CheckOnlyQueryParameter != nil {
			qrCheckOnly = *o.CheckOnlyQueryParameter
		}
		qCheckOnly := swag.FormatBool(qrCheckOnly)
		if qCheckOnly != "" {

			if err := r.SetQueryParam("check_only", qCheckOnly); err != nil {
				return err
			}
		}
	}

	if o.IgnoreWarningsQueryParameter != nil {

		// query param ignore_warnings
		var qrIgnoreWarnings bool

		if o.IgnoreWarningsQueryParameter != nil {
			qrIgnoreWarnings = *o.IgnoreWarningsQueryParameter
		}
		qIgnoreWarnings := swag.FormatBool(qrIgnoreWarnings)
		if qIgnoreWarnings != "" {

			if err := r.SetQueryParam("ignore_warnings", qIgnoreWarnings); err != nil {
				return err
			}
		}
	}
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
