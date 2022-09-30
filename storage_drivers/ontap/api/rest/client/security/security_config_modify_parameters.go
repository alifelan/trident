// Code generated by go-swagger; DO NOT EDIT.

package security

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

// NewSecurityConfigModifyParams creates a new SecurityConfigModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSecurityConfigModifyParams() *SecurityConfigModifyParams {
	return &SecurityConfigModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSecurityConfigModifyParamsWithTimeout creates a new SecurityConfigModifyParams object
// with the ability to set a timeout on a request.
func NewSecurityConfigModifyParamsWithTimeout(timeout time.Duration) *SecurityConfigModifyParams {
	return &SecurityConfigModifyParams{
		timeout: timeout,
	}
}

// NewSecurityConfigModifyParamsWithContext creates a new SecurityConfigModifyParams object
// with the ability to set a context for a request.
func NewSecurityConfigModifyParamsWithContext(ctx context.Context) *SecurityConfigModifyParams {
	return &SecurityConfigModifyParams{
		Context: ctx,
	}
}

// NewSecurityConfigModifyParamsWithHTTPClient creates a new SecurityConfigModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSecurityConfigModifyParamsWithHTTPClient(client *http.Client) *SecurityConfigModifyParams {
	return &SecurityConfigModifyParams{
		HTTPClient: client,
	}
}

/*
SecurityConfigModifyParams contains all the parameters to send to the API endpoint

	for the security config modify operation.

	Typically these are written to a http.Request.
*/
type SecurityConfigModifyParams struct {

	/* Info.

	   security info specification
	*/
	Info *models.SecurityConfig

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the security config modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityConfigModifyParams) WithDefaults() *SecurityConfigModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the security config modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SecurityConfigModifyParams) SetDefaults() {
	var (
		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := SecurityConfigModifyParams{
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the security config modify params
func (o *SecurityConfigModifyParams) WithTimeout(timeout time.Duration) *SecurityConfigModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the security config modify params
func (o *SecurityConfigModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the security config modify params
func (o *SecurityConfigModifyParams) WithContext(ctx context.Context) *SecurityConfigModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the security config modify params
func (o *SecurityConfigModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the security config modify params
func (o *SecurityConfigModifyParams) WithHTTPClient(client *http.Client) *SecurityConfigModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the security config modify params
func (o *SecurityConfigModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the security config modify params
func (o *SecurityConfigModifyParams) WithInfo(info *models.SecurityConfig) *SecurityConfigModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the security config modify params
func (o *SecurityConfigModifyParams) SetInfo(info *models.SecurityConfig) {
	o.Info = info
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the security config modify params
func (o *SecurityConfigModifyParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SecurityConfigModifyParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the security config modify params
func (o *SecurityConfigModifyParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *SecurityConfigModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
