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

// NewAccountPasswordCreateParams creates a new AccountPasswordCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountPasswordCreateParams() *AccountPasswordCreateParams {
	return &AccountPasswordCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountPasswordCreateParamsWithTimeout creates a new AccountPasswordCreateParams object
// with the ability to set a timeout on a request.
func NewAccountPasswordCreateParamsWithTimeout(timeout time.Duration) *AccountPasswordCreateParams {
	return &AccountPasswordCreateParams{
		timeout: timeout,
	}
}

// NewAccountPasswordCreateParamsWithContext creates a new AccountPasswordCreateParams object
// with the ability to set a context for a request.
func NewAccountPasswordCreateParamsWithContext(ctx context.Context) *AccountPasswordCreateParams {
	return &AccountPasswordCreateParams{
		Context: ctx,
	}
}

// NewAccountPasswordCreateParamsWithHTTPClient creates a new AccountPasswordCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountPasswordCreateParamsWithHTTPClient(client *http.Client) *AccountPasswordCreateParams {
	return &AccountPasswordCreateParams{
		HTTPClient: client,
	}
}

/*
AccountPasswordCreateParams contains all the parameters to send to the API endpoint

	for the account password create operation.

	Typically these are written to a http.Request.
*/
type AccountPasswordCreateParams struct {

	/* Info.

	   New password for the user account.
	*/
	Info *models.AccountPassword

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account password create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountPasswordCreateParams) WithDefaults() *AccountPasswordCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account password create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountPasswordCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)
	)

	val := AccountPasswordCreateParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the account password create params
func (o *AccountPasswordCreateParams) WithTimeout(timeout time.Duration) *AccountPasswordCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account password create params
func (o *AccountPasswordCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account password create params
func (o *AccountPasswordCreateParams) WithContext(ctx context.Context) *AccountPasswordCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account password create params
func (o *AccountPasswordCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account password create params
func (o *AccountPasswordCreateParams) WithHTTPClient(client *http.Client) *AccountPasswordCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account password create params
func (o *AccountPasswordCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the account password create params
func (o *AccountPasswordCreateParams) WithInfo(info *models.AccountPassword) *AccountPasswordCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the account password create params
func (o *AccountPasswordCreateParams) SetInfo(info *models.AccountPassword) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the account password create params
func (o *AccountPasswordCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *AccountPasswordCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the account password create params
func (o *AccountPasswordCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WriteToRequest writes these params to a swagger request
func (o *AccountPasswordCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
