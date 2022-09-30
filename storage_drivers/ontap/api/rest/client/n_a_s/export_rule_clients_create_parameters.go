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

// NewExportRuleClientsCreateParams creates a new ExportRuleClientsCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewExportRuleClientsCreateParams() *ExportRuleClientsCreateParams {
	return &ExportRuleClientsCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewExportRuleClientsCreateParamsWithTimeout creates a new ExportRuleClientsCreateParams object
// with the ability to set a timeout on a request.
func NewExportRuleClientsCreateParamsWithTimeout(timeout time.Duration) *ExportRuleClientsCreateParams {
	return &ExportRuleClientsCreateParams{
		timeout: timeout,
	}
}

// NewExportRuleClientsCreateParamsWithContext creates a new ExportRuleClientsCreateParams object
// with the ability to set a context for a request.
func NewExportRuleClientsCreateParamsWithContext(ctx context.Context) *ExportRuleClientsCreateParams {
	return &ExportRuleClientsCreateParams{
		Context: ctx,
	}
}

// NewExportRuleClientsCreateParamsWithHTTPClient creates a new ExportRuleClientsCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewExportRuleClientsCreateParamsWithHTTPClient(client *http.Client) *ExportRuleClientsCreateParams {
	return &ExportRuleClientsCreateParams{
		HTTPClient: client,
	}
}

/*
ExportRuleClientsCreateParams contains all the parameters to send to the API endpoint

	for the export rule clients create operation.

	Typically these are written to a http.Request.
*/
type ExportRuleClientsCreateParams struct {

	/* Index.

	   Export Rule Index
	*/
	IndexPathParameter int64

	/* Info.

	   Info specification
	*/
	Info *models.ExportClient

	/* PolicyID.

	   Export Policy ID
	*/
	PolicyIDPathParameter int64

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the export rule clients create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleClientsCreateParams) WithDefaults() *ExportRuleClientsCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the export rule clients create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ExportRuleClientsCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)
	)

	val := ExportRuleClientsCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the export rule clients create params
func (o *ExportRuleClientsCreateParams) WithTimeout(timeout time.Duration) *ExportRuleClientsCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the export rule clients create params
func (o *ExportRuleClientsCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the export rule clients create params
func (o *ExportRuleClientsCreateParams) WithContext(ctx context.Context) *ExportRuleClientsCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the export rule clients create params
func (o *ExportRuleClientsCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the export rule clients create params
func (o *ExportRuleClientsCreateParams) WithHTTPClient(client *http.Client) *ExportRuleClientsCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the export rule clients create params
func (o *ExportRuleClientsCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIndexPathParameter adds the index to the export rule clients create params
func (o *ExportRuleClientsCreateParams) WithIndexPathParameter(index int64) *ExportRuleClientsCreateParams {
	o.SetIndexPathParameter(index)
	return o
}

// SetIndexPathParameter adds the index to the export rule clients create params
func (o *ExportRuleClientsCreateParams) SetIndexPathParameter(index int64) {
	o.IndexPathParameter = index
}

// WithInfo adds the info to the export rule clients create params
func (o *ExportRuleClientsCreateParams) WithInfo(info *models.ExportClient) *ExportRuleClientsCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the export rule clients create params
func (o *ExportRuleClientsCreateParams) SetInfo(info *models.ExportClient) {
	o.Info = info
}

// WithPolicyIDPathParameter adds the policyID to the export rule clients create params
func (o *ExportRuleClientsCreateParams) WithPolicyIDPathParameter(policyID int64) *ExportRuleClientsCreateParams {
	o.SetPolicyIDPathParameter(policyID)
	return o
}

// SetPolicyIDPathParameter adds the policyId to the export rule clients create params
func (o *ExportRuleClientsCreateParams) SetPolicyIDPathParameter(policyID int64) {
	o.PolicyIDPathParameter = policyID
}

// WithReturnRecordsQueryParameter adds the returnRecords to the export rule clients create params
func (o *ExportRuleClientsCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *ExportRuleClientsCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the export rule clients create params
func (o *ExportRuleClientsCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *ExportRuleClientsCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param index
	if err := r.SetPathParam("index", swag.FormatInt64(o.IndexPathParameter)); err != nil {
		return err
	}
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param policy.id
	if err := r.SetPathParam("policy.id", swag.FormatInt64(o.PolicyIDPathParameter)); err != nil {
		return err
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
