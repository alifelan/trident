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
)

// NewVolumeGetParams creates a new VolumeGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewVolumeGetParams() *VolumeGetParams {
	return &VolumeGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewVolumeGetParamsWithTimeout creates a new VolumeGetParams object
// with the ability to set a timeout on a request.
func NewVolumeGetParamsWithTimeout(timeout time.Duration) *VolumeGetParams {
	return &VolumeGetParams{
		timeout: timeout,
	}
}

// NewVolumeGetParamsWithContext creates a new VolumeGetParams object
// with the ability to set a context for a request.
func NewVolumeGetParamsWithContext(ctx context.Context) *VolumeGetParams {
	return &VolumeGetParams{
		Context: ctx,
	}
}

// NewVolumeGetParamsWithHTTPClient creates a new VolumeGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewVolumeGetParamsWithHTTPClient(client *http.Client) *VolumeGetParams {
	return &VolumeGetParams{
		HTTPClient: client,
	}
}

/*
VolumeGetParams contains all the parameters to send to the API endpoint

	for the volume get operation.

	Typically these are written to a http.Request.
*/
type VolumeGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* IsConstituent.

	   When set to false, only FlexVol and FlexGroup volumes are returned.  When set to true, only FlexGroup constituent volumes are returned. Default for GET calls is false.
	*/
	IsConstituentQueryParameter *bool

	/* UUID.

	   Unique identifier of the volume.
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the volume get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeGetParams) WithDefaults() *VolumeGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the volume get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *VolumeGetParams) SetDefaults() {
	var (
		isConstituentQueryParameterDefault = bool(false)
	)

	val := VolumeGetParams{
		IsConstituentQueryParameter: &isConstituentQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the volume get params
func (o *VolumeGetParams) WithTimeout(timeout time.Duration) *VolumeGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the volume get params
func (o *VolumeGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the volume get params
func (o *VolumeGetParams) WithContext(ctx context.Context) *VolumeGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the volume get params
func (o *VolumeGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the volume get params
func (o *VolumeGetParams) WithHTTPClient(client *http.Client) *VolumeGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the volume get params
func (o *VolumeGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the volume get params
func (o *VolumeGetParams) WithFieldsQueryParameter(fields []string) *VolumeGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the volume get params
func (o *VolumeGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIsConstituentQueryParameter adds the isConstituent to the volume get params
func (o *VolumeGetParams) WithIsConstituentQueryParameter(isConstituent *bool) *VolumeGetParams {
	o.SetIsConstituentQueryParameter(isConstituent)
	return o
}

// SetIsConstituentQueryParameter adds the isConstituent to the volume get params
func (o *VolumeGetParams) SetIsConstituentQueryParameter(isConstituent *bool) {
	o.IsConstituentQueryParameter = isConstituent
}

// WithUUIDPathParameter adds the uuid to the volume get params
func (o *VolumeGetParams) WithUUIDPathParameter(uuid string) *VolumeGetParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the volume get params
func (o *VolumeGetParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *VolumeGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.IsConstituentQueryParameter != nil {

		// query param is_constituent
		var qrIsConstituent bool

		if o.IsConstituentQueryParameter != nil {
			qrIsConstituent = *o.IsConstituentQueryParameter
		}
		qIsConstituent := swag.FormatBool(qrIsConstituent)
		if qIsConstituent != "" {

			if err := r.SetQueryParam("is_constituent", qIsConstituent); err != nil {
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

// bindParamVolumeGet binds the parameter fields
func (o *VolumeGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.FieldsQueryParameter

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}
