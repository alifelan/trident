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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewClusterNtpKeysCreateParams creates a new ClusterNtpKeysCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterNtpKeysCreateParams() *ClusterNtpKeysCreateParams {
	return &ClusterNtpKeysCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterNtpKeysCreateParamsWithTimeout creates a new ClusterNtpKeysCreateParams object
// with the ability to set a timeout on a request.
func NewClusterNtpKeysCreateParamsWithTimeout(timeout time.Duration) *ClusterNtpKeysCreateParams {
	return &ClusterNtpKeysCreateParams{
		timeout: timeout,
	}
}

// NewClusterNtpKeysCreateParamsWithContext creates a new ClusterNtpKeysCreateParams object
// with the ability to set a context for a request.
func NewClusterNtpKeysCreateParamsWithContext(ctx context.Context) *ClusterNtpKeysCreateParams {
	return &ClusterNtpKeysCreateParams{
		Context: ctx,
	}
}

// NewClusterNtpKeysCreateParamsWithHTTPClient creates a new ClusterNtpKeysCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterNtpKeysCreateParamsWithHTTPClient(client *http.Client) *ClusterNtpKeysCreateParams {
	return &ClusterNtpKeysCreateParams{
		HTTPClient: client,
	}
}

/*
ClusterNtpKeysCreateParams contains all the parameters to send to the API endpoint

	for the cluster ntp keys create operation.

	Typically these are written to a http.Request.
*/
type ClusterNtpKeysCreateParams struct {

	/* Info.

	   Information specification
	*/
	Info *models.NtpKey

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster ntp keys create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterNtpKeysCreateParams) WithDefaults() *ClusterNtpKeysCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster ntp keys create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterNtpKeysCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)
	)

	val := ClusterNtpKeysCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cluster ntp keys create params
func (o *ClusterNtpKeysCreateParams) WithTimeout(timeout time.Duration) *ClusterNtpKeysCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster ntp keys create params
func (o *ClusterNtpKeysCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster ntp keys create params
func (o *ClusterNtpKeysCreateParams) WithContext(ctx context.Context) *ClusterNtpKeysCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster ntp keys create params
func (o *ClusterNtpKeysCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster ntp keys create params
func (o *ClusterNtpKeysCreateParams) WithHTTPClient(client *http.Client) *ClusterNtpKeysCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster ntp keys create params
func (o *ClusterNtpKeysCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the cluster ntp keys create params
func (o *ClusterNtpKeysCreateParams) WithInfo(info *models.NtpKey) *ClusterNtpKeysCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cluster ntp keys create params
func (o *ClusterNtpKeysCreateParams) SetInfo(info *models.NtpKey) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the cluster ntp keys create params
func (o *ClusterNtpKeysCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *ClusterNtpKeysCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the cluster ntp keys create params
func (o *ClusterNtpKeysCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterNtpKeysCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
