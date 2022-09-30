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

// NewCifsServiceCreateParams creates a new CifsServiceCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsServiceCreateParams() *CifsServiceCreateParams {
	return &CifsServiceCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsServiceCreateParamsWithTimeout creates a new CifsServiceCreateParams object
// with the ability to set a timeout on a request.
func NewCifsServiceCreateParamsWithTimeout(timeout time.Duration) *CifsServiceCreateParams {
	return &CifsServiceCreateParams{
		timeout: timeout,
	}
}

// NewCifsServiceCreateParamsWithContext creates a new CifsServiceCreateParams object
// with the ability to set a context for a request.
func NewCifsServiceCreateParamsWithContext(ctx context.Context) *CifsServiceCreateParams {
	return &CifsServiceCreateParams{
		Context: ctx,
	}
}

// NewCifsServiceCreateParamsWithHTTPClient creates a new CifsServiceCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsServiceCreateParamsWithHTTPClient(client *http.Client) *CifsServiceCreateParams {
	return &CifsServiceCreateParams{
		HTTPClient: client,
	}
}

/*
CifsServiceCreateParams contains all the parameters to send to the API endpoint

	for the cifs service create operation.

	Typically these are written to a http.Request.
*/
type CifsServiceCreateParams struct {

	/* Force.

	   If this is set and a machine account with the same name as specified in 'cifs-server name' exists in the Active Directory, existing  machine account will be overwritten and reused.
	*/
	ForceQueryParameter *bool

	/* Info.

	   Info specification
	*/
	Info *models.CifsService

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

// WithDefaults hydrates default values in the cifs service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsServiceCreateParams) WithDefaults() *CifsServiceCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs service create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsServiceCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)

		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := CifsServiceCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cifs service create params
func (o *CifsServiceCreateParams) WithTimeout(timeout time.Duration) *CifsServiceCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs service create params
func (o *CifsServiceCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs service create params
func (o *CifsServiceCreateParams) WithContext(ctx context.Context) *CifsServiceCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs service create params
func (o *CifsServiceCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs service create params
func (o *CifsServiceCreateParams) WithHTTPClient(client *http.Client) *CifsServiceCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs service create params
func (o *CifsServiceCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithForceQueryParameter adds the force to the cifs service create params
func (o *CifsServiceCreateParams) WithForceQueryParameter(force *bool) *CifsServiceCreateParams {
	o.SetForceQueryParameter(force)
	return o
}

// SetForceQueryParameter adds the force to the cifs service create params
func (o *CifsServiceCreateParams) SetForceQueryParameter(force *bool) {
	o.ForceQueryParameter = force
}

// WithInfo adds the info to the cifs service create params
func (o *CifsServiceCreateParams) WithInfo(info *models.CifsService) *CifsServiceCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cifs service create params
func (o *CifsServiceCreateParams) SetInfo(info *models.CifsService) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the cifs service create params
func (o *CifsServiceCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *CifsServiceCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the cifs service create params
func (o *CifsServiceCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the cifs service create params
func (o *CifsServiceCreateParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *CifsServiceCreateParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the cifs service create params
func (o *CifsServiceCreateParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *CifsServiceCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ForceQueryParameter != nil {

		// query param force
		var qrForce bool

		if o.ForceQueryParameter != nil {
			qrForce = *o.ForceQueryParameter
		}
		qForce := swag.FormatBool(qrForce)
		if qForce != "" {

			if err := r.SetQueryParam("force", qForce); err != nil {
				return err
			}
		}
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
