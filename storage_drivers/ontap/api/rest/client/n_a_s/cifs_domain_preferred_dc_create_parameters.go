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

// NewCifsDomainPreferredDcCreateParams creates a new CifsDomainPreferredDcCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCifsDomainPreferredDcCreateParams() *CifsDomainPreferredDcCreateParams {
	return &CifsDomainPreferredDcCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCifsDomainPreferredDcCreateParamsWithTimeout creates a new CifsDomainPreferredDcCreateParams object
// with the ability to set a timeout on a request.
func NewCifsDomainPreferredDcCreateParamsWithTimeout(timeout time.Duration) *CifsDomainPreferredDcCreateParams {
	return &CifsDomainPreferredDcCreateParams{
		timeout: timeout,
	}
}

// NewCifsDomainPreferredDcCreateParamsWithContext creates a new CifsDomainPreferredDcCreateParams object
// with the ability to set a context for a request.
func NewCifsDomainPreferredDcCreateParamsWithContext(ctx context.Context) *CifsDomainPreferredDcCreateParams {
	return &CifsDomainPreferredDcCreateParams{
		Context: ctx,
	}
}

// NewCifsDomainPreferredDcCreateParamsWithHTTPClient creates a new CifsDomainPreferredDcCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewCifsDomainPreferredDcCreateParamsWithHTTPClient(client *http.Client) *CifsDomainPreferredDcCreateParams {
	return &CifsDomainPreferredDcCreateParams{
		HTTPClient: client,
	}
}

/*
CifsDomainPreferredDcCreateParams contains all the parameters to send to the API endpoint

	for the cifs domain preferred dc create operation.

	Typically these are written to a http.Request.
*/
type CifsDomainPreferredDcCreateParams struct {

	/* Info.

	   Info specification
	*/
	Info *models.CifsDomainPreferredDc

	/* ReturnRecords.

	   The default is false.  If set to true, the records are returned.
	*/
	ReturnRecordsQueryParameter *bool

	/* SkipConfigValidation.

	   Skip the validation of the specified preferred DC configuration.
	*/
	SkIPConfigValidationQueryParameter *bool

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cifs domain preferred dc create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsDomainPreferredDcCreateParams) WithDefaults() *CifsDomainPreferredDcCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cifs domain preferred dc create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CifsDomainPreferredDcCreateParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(false)

		skIPConfigValidationQueryParameterDefault = bool(false)
	)

	val := CifsDomainPreferredDcCreateParams{
		ReturnRecordsQueryParameter:        &returnRecordsQueryParameterDefault,
		SkIPConfigValidationQueryParameter: &skIPConfigValidationQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the cifs domain preferred dc create params
func (o *CifsDomainPreferredDcCreateParams) WithTimeout(timeout time.Duration) *CifsDomainPreferredDcCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cifs domain preferred dc create params
func (o *CifsDomainPreferredDcCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cifs domain preferred dc create params
func (o *CifsDomainPreferredDcCreateParams) WithContext(ctx context.Context) *CifsDomainPreferredDcCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cifs domain preferred dc create params
func (o *CifsDomainPreferredDcCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cifs domain preferred dc create params
func (o *CifsDomainPreferredDcCreateParams) WithHTTPClient(client *http.Client) *CifsDomainPreferredDcCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cifs domain preferred dc create params
func (o *CifsDomainPreferredDcCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the cifs domain preferred dc create params
func (o *CifsDomainPreferredDcCreateParams) WithInfo(info *models.CifsDomainPreferredDc) *CifsDomainPreferredDcCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cifs domain preferred dc create params
func (o *CifsDomainPreferredDcCreateParams) SetInfo(info *models.CifsDomainPreferredDc) {
	o.Info = info
}

// WithReturnRecordsQueryParameter adds the returnRecords to the cifs domain preferred dc create params
func (o *CifsDomainPreferredDcCreateParams) WithReturnRecordsQueryParameter(returnRecords *bool) *CifsDomainPreferredDcCreateParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the cifs domain preferred dc create params
func (o *CifsDomainPreferredDcCreateParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithSkIPConfigValidationQueryParameter adds the skipConfigValidation to the cifs domain preferred dc create params
func (o *CifsDomainPreferredDcCreateParams) WithSkIPConfigValidationQueryParameter(skipConfigValidation *bool) *CifsDomainPreferredDcCreateParams {
	o.SetSkIPConfigValidationQueryParameter(skipConfigValidation)
	return o
}

// SetSkIPConfigValidationQueryParameter adds the skipConfigValidation to the cifs domain preferred dc create params
func (o *CifsDomainPreferredDcCreateParams) SetSkIPConfigValidationQueryParameter(skipConfigValidation *bool) {
	o.SkIPConfigValidationQueryParameter = skipConfigValidation
}

// WithSVMUUIDPathParameter adds the svmUUID to the cifs domain preferred dc create params
func (o *CifsDomainPreferredDcCreateParams) WithSVMUUIDPathParameter(svmUUID string) *CifsDomainPreferredDcCreateParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the cifs domain preferred dc create params
func (o *CifsDomainPreferredDcCreateParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *CifsDomainPreferredDcCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.SkIPConfigValidationQueryParameter != nil {

		// query param skip_config_validation
		var qrSkipConfigValidation bool

		if o.SkIPConfigValidationQueryParameter != nil {
			qrSkipConfigValidation = *o.SkIPConfigValidationQueryParameter
		}
		qSkipConfigValidation := swag.FormatBool(qrSkipConfigValidation)
		if qSkipConfigValidation != "" {

			if err := r.SetQueryParam("skip_config_validation", qSkipConfigValidation); err != nil {
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
