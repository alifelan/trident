// Code generated by go-swagger; DO NOT EDIT.

package name_services

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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewUnixUserSettingsModifyParams creates a new UnixUserSettingsModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUnixUserSettingsModifyParams() *UnixUserSettingsModifyParams {
	return &UnixUserSettingsModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUnixUserSettingsModifyParamsWithTimeout creates a new UnixUserSettingsModifyParams object
// with the ability to set a timeout on a request.
func NewUnixUserSettingsModifyParamsWithTimeout(timeout time.Duration) *UnixUserSettingsModifyParams {
	return &UnixUserSettingsModifyParams{
		timeout: timeout,
	}
}

// NewUnixUserSettingsModifyParamsWithContext creates a new UnixUserSettingsModifyParams object
// with the ability to set a context for a request.
func NewUnixUserSettingsModifyParamsWithContext(ctx context.Context) *UnixUserSettingsModifyParams {
	return &UnixUserSettingsModifyParams{
		Context: ctx,
	}
}

// NewUnixUserSettingsModifyParamsWithHTTPClient creates a new UnixUserSettingsModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewUnixUserSettingsModifyParamsWithHTTPClient(client *http.Client) *UnixUserSettingsModifyParams {
	return &UnixUserSettingsModifyParams{
		HTTPClient: client,
	}
}

/*
UnixUserSettingsModifyParams contains all the parameters to send to the API endpoint

	for the unix user settings modify operation.

	Typically these are written to a http.Request.
*/
type UnixUserSettingsModifyParams struct {

	/* Info.

	   Info specification.
	*/
	Info *models.UnixUserSettings

	/* SvmUUID.

	   SVM UUID.
	*/
	SVMUUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the unix user settings modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixUserSettingsModifyParams) WithDefaults() *UnixUserSettingsModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the unix user settings modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UnixUserSettingsModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the unix user settings modify params
func (o *UnixUserSettingsModifyParams) WithTimeout(timeout time.Duration) *UnixUserSettingsModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the unix user settings modify params
func (o *UnixUserSettingsModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the unix user settings modify params
func (o *UnixUserSettingsModifyParams) WithContext(ctx context.Context) *UnixUserSettingsModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the unix user settings modify params
func (o *UnixUserSettingsModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the unix user settings modify params
func (o *UnixUserSettingsModifyParams) WithHTTPClient(client *http.Client) *UnixUserSettingsModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the unix user settings modify params
func (o *UnixUserSettingsModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the unix user settings modify params
func (o *UnixUserSettingsModifyParams) WithInfo(info *models.UnixUserSettings) *UnixUserSettingsModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the unix user settings modify params
func (o *UnixUserSettingsModifyParams) SetInfo(info *models.UnixUserSettings) {
	o.Info = info
}

// WithSVMUUIDPathParameter adds the svmUUID to the unix user settings modify params
func (o *UnixUserSettingsModifyParams) WithSVMUUIDPathParameter(svmUUID string) *UnixUserSettingsModifyParams {
	o.SetSVMUUIDPathParameter(svmUUID)
	return o
}

// SetSVMUUIDPathParameter adds the svmUuid to the unix user settings modify params
func (o *UnixUserSettingsModifyParams) SetSVMUUIDPathParameter(svmUUID string) {
	o.SVMUUIDPathParameter = svmUUID
}

// WriteToRequest writes these params to a swagger request
func (o *UnixUserSettingsModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
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
