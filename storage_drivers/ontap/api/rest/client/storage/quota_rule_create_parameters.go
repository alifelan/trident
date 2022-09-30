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

// NewQuotaRuleCreateParams creates a new QuotaRuleCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewQuotaRuleCreateParams() *QuotaRuleCreateParams {
	return &QuotaRuleCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewQuotaRuleCreateParamsWithTimeout creates a new QuotaRuleCreateParams object
// with the ability to set a timeout on a request.
func NewQuotaRuleCreateParamsWithTimeout(timeout time.Duration) *QuotaRuleCreateParams {
	return &QuotaRuleCreateParams{
		timeout: timeout,
	}
}

// NewQuotaRuleCreateParamsWithContext creates a new QuotaRuleCreateParams object
// with the ability to set a context for a request.
func NewQuotaRuleCreateParamsWithContext(ctx context.Context) *QuotaRuleCreateParams {
	return &QuotaRuleCreateParams{
		Context: ctx,
	}
}

// NewQuotaRuleCreateParamsWithHTTPClient creates a new QuotaRuleCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewQuotaRuleCreateParamsWithHTTPClient(client *http.Client) *QuotaRuleCreateParams {
	return &QuotaRuleCreateParams{
		HTTPClient: client,
	}
}

/*
QuotaRuleCreateParams contains all the parameters to send to the API endpoint

	for the quota rule create operation.

	Typically these are written to a http.Request.
*/
type QuotaRuleCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.QuotaRule

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

// WithDefaults hydrates default values in the quota rule create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuotaRuleCreateParams) WithDefaults() *QuotaRuleCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the quota rule create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *QuotaRuleCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)

		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := QuotaRuleCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the quota rule create params
func (o *QuotaRuleCreateParams) WithTimeout(timeout time.Duration) *QuotaRuleCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the quota rule create params
func (o *QuotaRuleCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the quota rule create params
func (o *QuotaRuleCreateParams) WithContext(ctx context.Context) *QuotaRuleCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the quota rule create params
func (o *QuotaRuleCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the quota rule create params
func (o *QuotaRuleCreateParams) WithHTTPClient(client *http.Client) *QuotaRuleCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the quota rule create params
func (o *QuotaRuleCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the quota rule create params
func (o *QuotaRuleCreateParams) WithInfo(info *models.QuotaRule) *QuotaRuleCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the quota rule create params
func (o *QuotaRuleCreateParams) SetInfo(info *models.QuotaRule) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the quota rule create params
func (o *QuotaRuleCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *QuotaRuleCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the quota rule create params
func (o *QuotaRuleCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the quota rule create params
func (o *QuotaRuleCreateParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *QuotaRuleCreateParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the quota rule create params
func (o *QuotaRuleCreateParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *QuotaRuleCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
