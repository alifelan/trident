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

// NewCifsShareCreateParams creates a new CifsShareCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsShareCreateParams() *CifsShareCreateParams {
	return &CifsShareCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsShareCreateParamsWithTimeout creates a new CifsShareCreateParams object
// with the ability to set a timeout on a request.
func NewCifsShareCreateParamsWithTimeout(timeout time.Duration) *CifsShareCreateParams {
	return &CifsShareCreateParams{
		timeout: timeout,
	}
}

// NewCifsShareCreateParamsWithContext creates a new CifsShareCreateParams object
// with the ability to set a context for a request.
func NewCifsShareCreateParamsWithContext(ctx context.Context) *CifsShareCreateParams {
	return &CifsShareCreateParams{
		Context: ctx,
	}
}

// NewCifsShareCreateParamsWithHTTPClient creates a new CifsShareCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsShareCreateParamsWithHTTPClient(client *http.Client) *CifsShareCreateParams {
	return &CifsShareCreateParams{
		HTTPClient: client,
	}
}

/*
CifsShareCreateParams contains all the parameters to send to the API endpoint

	for the cifs share create operation.

	Typically these are written to a http.Request.
*/
type CifsShareCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.CifsShare

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs share create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsShareCreateParams) WithDefaults() *CifsShareCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs share create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsShareCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)
	)

	val := CifsShareCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cifs share create params
func (o *CifsShareCreateParams) WithTimeout(timeout time.Duration) *CifsShareCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs share create params
func (o *CifsShareCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs share create params
func (o *CifsShareCreateParams) WithContext(ctx context.Context) *CifsShareCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs share create params
func (o *CifsShareCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs share create params
func (o *CifsShareCreateParams) WithHTTPClient(client *http.Client) *CifsShareCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs share create params
func (o *CifsShareCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the cifs share create params
func (o *CifsShareCreateParams) WithInfo(info *models.CifsShare) *CifsShareCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cifs share create params
func (o *CifsShareCreateParams) SetInfo(info *models.CifsShare) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the cifs share create params
func (o *CifsShareCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *CifsShareCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the cifs share create params
func (o *CifsShareCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *CifsShareCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
