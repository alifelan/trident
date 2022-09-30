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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewShelfModifyParams creates a new ShelfModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewShelfModifyParams() *ShelfModifyParams {
	return &ShelfModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewShelfModifyParamsWithTimeout creates a new ShelfModifyParams object
// with the ability to set a timeout on a request.
func NewShelfModifyParamsWithTimeout(timeout time.Duration) *ShelfModifyParams {
	return &ShelfModifyParams{
		timeout: timeout,
	}
}

// NewShelfModifyParamsWithContext creates a new ShelfModifyParams object
// with the ability to set a context for a request.
func NewShelfModifyParamsWithContext(ctx context.Context) *ShelfModifyParams {
	return &ShelfModifyParams{
		Context: ctx,
	}
}

// NewShelfModifyParamsWithHTTPClient creates a new ShelfModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewShelfModifyParamsWithHTTPClient(client *http.Client) *ShelfModifyParams {
	return &ShelfModifyParams{
		HTTPClient: client,
	}
}

/*
ShelfModifyParams contains all the parameters to send to the API endpoint

	for the shelf modify operation.

	Typically these are written to a http.Request.
*/
type ShelfModifyParams struct {

	/* Info.

	   The new property values for the shelf.
	*/
	Info *models.Shelf

	/* UID.

	   Shelf UID
	*/
	UIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the shelf modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShelfModifyParams) WithDefaults() *ShelfModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the shelf modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ShelfModifyParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the shelf modify params
func (o *ShelfModifyParams) WithTimeout(timeout time.Duration) *ShelfModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the shelf modify params
func (o *ShelfModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the shelf modify params
func (o *ShelfModifyParams) WithContext(ctx context.Context) *ShelfModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the shelf modify params
func (o *ShelfModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the shelf modify params
func (o *ShelfModifyParams) WithHTTPClient(client *http.Client) *ShelfModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the shelf modify params
func (o *ShelfModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the shelf modify params
func (o *ShelfModifyParams) WithInfo(info *models.Shelf) *ShelfModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the shelf modify params
func (o *ShelfModifyParams) SetInfo(info *models.Shelf) {
	o.Info = info
}

// WithUIDPathParameter adds the uid to the shelf modify params
func (o *ShelfModifyParams) WithUIDPathParameter(uid string) *ShelfModifyParams {
	o.SetUIDPathParameter(uid)
	return o
}

// SetUIDPathParameter adds the uid to the shelf modify params
func (o *ShelfModifyParams) SetUIDPathParameter(uid string) {
	o.UIDPathParameter = uid
}

// WriteToRequest writes these params to a swagger request
func (o *ShelfModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	// path param uid
	if err := r.SetPathParam("uid", o.UIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
