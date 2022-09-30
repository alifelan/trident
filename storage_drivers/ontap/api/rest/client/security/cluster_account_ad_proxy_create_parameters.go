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

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NewClusterAccountAdProxyCreateParams creates a new ClusterAccountAdProxyCreateParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewClusterAccountAdProxyCreateParams() *ClusterAccountAdProxyCreateParams {
	return &ClusterAccountAdProxyCreateParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewClusterAccountAdProxyCreateParamsWithTimeout creates a new ClusterAccountAdProxyCreateParams object
// with the ability to set a timeout on a request.
func NewClusterAccountAdProxyCreateParamsWithTimeout(timeout time.Duration) *ClusterAccountAdProxyCreateParams {
	return &ClusterAccountAdProxyCreateParams{
		timeout: timeout,
	}
}

// NewClusterAccountAdProxyCreateParamsWithContext creates a new ClusterAccountAdProxyCreateParams object
// with the ability to set a context for a request.
func NewClusterAccountAdProxyCreateParamsWithContext(ctx context.Context) *ClusterAccountAdProxyCreateParams {
	return &ClusterAccountAdProxyCreateParams{
		Context: ctx,
	}
}

// NewClusterAccountAdProxyCreateParamsWithHTTPClient creates a new ClusterAccountAdProxyCreateParams object
// with the ability to set a custom HTTPClient for a request.
func NewClusterAccountAdProxyCreateParamsWithHTTPClient(client *http.Client) *ClusterAccountAdProxyCreateParams {
	return &ClusterAccountAdProxyCreateParams{
		HTTPClient: client,
	}
}

/*
ClusterAccountAdProxyCreateParams contains all the parameters to send to the API endpoint

	for the cluster account ad proxy create operation.

	Typically these are written to a http.Request.
*/
type ClusterAccountAdProxyCreateParams struct {

	/* Info.

	   The data SVM that tunnels the Active Directory authentication requests.
	*/
	Info *models.ClusterAdProxy

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the cluster account ad proxy create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterAccountAdProxyCreateParams) WithDefaults() *ClusterAccountAdProxyCreateParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the cluster account ad proxy create params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ClusterAccountAdProxyCreateParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the cluster account ad proxy create params
func (o *ClusterAccountAdProxyCreateParams) WithTimeout(timeout time.Duration) *ClusterAccountAdProxyCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the cluster account ad proxy create params
func (o *ClusterAccountAdProxyCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the cluster account ad proxy create params
func (o *ClusterAccountAdProxyCreateParams) WithContext(ctx context.Context) *ClusterAccountAdProxyCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the cluster account ad proxy create params
func (o *ClusterAccountAdProxyCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the cluster account ad proxy create params
func (o *ClusterAccountAdProxyCreateParams) WithHTTPClient(client *http.Client) *ClusterAccountAdProxyCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the cluster account ad proxy create params
func (o *ClusterAccountAdProxyCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the cluster account ad proxy create params
func (o *ClusterAccountAdProxyCreateParams) WithInfo(info *models.ClusterAdProxy) *ClusterAccountAdProxyCreateParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the cluster account ad proxy create params
func (o *ClusterAccountAdProxyCreateParams) SetInfo(info *models.ClusterAdProxy) {
	o.Info = info
}

// WriteToRequest writes these params to a swagger request
func (o *ClusterAccountAdProxyCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
