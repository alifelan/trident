// Code generated by go-swagger; DO NOT EDIT.

package cluster

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

// NewSoftwarePackageDeleteParams creates a new SoftwarePackageDeleteParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSoftwarePackageDeleteParams() *SoftwarePackageDeleteParams {
	return &SoftwarePackageDeleteParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSoftwarePackageDeleteParamsWithTimeout creates a new SoftwarePackageDeleteParams object
// with the ability to set a timeout on a request.
func NewSoftwarePackageDeleteParamsWithTimeout(timeout time.Duration) *SoftwarePackageDeleteParams {
	return &SoftwarePackageDeleteParams{
		timeout: timeout,
	}
}

// NewSoftwarePackageDeleteParamsWithContext creates a new SoftwarePackageDeleteParams object
// with the ability to set a context for a request.
func NewSoftwarePackageDeleteParamsWithContext(ctx context.Context) *SoftwarePackageDeleteParams {
	return &SoftwarePackageDeleteParams{
		Context: ctx,
	}
}

// NewSoftwarePackageDeleteParamsWithHTTPClient creates a new SoftwarePackageDeleteParams object
// with the ability to set a custom HTTPClient for a request.
func NewSoftwarePackageDeleteParamsWithHTTPClient(client *http.Client) *SoftwarePackageDeleteParams {
	return &SoftwarePackageDeleteParams{
		HTTPClient: client,
	}
}

/*
SoftwarePackageDeleteParams contains all the parameters to send to the API endpoint

	for the software package delete operation.

	Typically these are written to a http.Request.
*/
type SoftwarePackageDeleteParams struct {

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	// Version.
	VersionPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the software package delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SoftwarePackageDeleteParams) WithDefaults() *SoftwarePackageDeleteParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the software package delete params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SoftwarePackageDeleteParams) SetDefaults() {
	var (
		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := SoftwarePackageDeleteParams{
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the software package delete params
func (o *SoftwarePackageDeleteParams) WithTimeout(timeout time.Duration) *SoftwarePackageDeleteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the software package delete params
func (o *SoftwarePackageDeleteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the software package delete params
func (o *SoftwarePackageDeleteParams) WithContext(ctx context.Context) *SoftwarePackageDeleteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the software package delete params
func (o *SoftwarePackageDeleteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the software package delete params
func (o *SoftwarePackageDeleteParams) WithHTTPClient(client *http.Client) *SoftwarePackageDeleteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the software package delete params
func (o *SoftwarePackageDeleteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the software package delete params
func (o *SoftwarePackageDeleteParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SoftwarePackageDeleteParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the software package delete params
func (o *SoftwarePackageDeleteParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithVersionPathParameter adds the version to the software package delete params
func (o *SoftwarePackageDeleteParams) WithVersionPathParameter(version string) *SoftwarePackageDeleteParams {
	o.SetVersionPathParameter(version)
	return o
}

// SetVersionPathParameter adds the version to the software package delete params
func (o *SoftwarePackageDeleteParams) SetVersionPathParameter(version string) {
	o.VersionPathParameter = version
}

// WriteToRequest writes these params to a swagger request
func (o *SoftwarePackageDeleteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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

	// path param version
	if err := r.SetPathParam("version", o.VersionPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
