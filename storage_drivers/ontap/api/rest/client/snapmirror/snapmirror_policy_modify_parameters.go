// Code generated by go-swagger; DO NOT EDIT.

package snapmirror

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

// NewSnapmirrorPolicyModifyParams creates a new SnapmirrorPolicyModifyParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapmirrorPolicyModifyParams() *SnapmirrorPolicyModifyParams {
	return &SnapmirrorPolicyModifyParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapmirrorPolicyModifyParamsWithTimeout creates a new SnapmirrorPolicyModifyParams object
// with the ability to set a timeout on a request.
func NewSnapmirrorPolicyModifyParamsWithTimeout(timeout time.Duration) *SnapmirrorPolicyModifyParams {
	return &SnapmirrorPolicyModifyParams{
		timeout: timeout,
	}
}

// NewSnapmirrorPolicyModifyParamsWithContext creates a new SnapmirrorPolicyModifyParams object
// with the ability to set a context for a request.
func NewSnapmirrorPolicyModifyParamsWithContext(ctx context.Context) *SnapmirrorPolicyModifyParams {
	return &SnapmirrorPolicyModifyParams{
		Context: ctx,
	}
}

// NewSnapmirrorPolicyModifyParamsWithHTTPClient creates a new SnapmirrorPolicyModifyParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapmirrorPolicyModifyParamsWithHTTPClient(client *http.Client) *SnapmirrorPolicyModifyParams {
	return &SnapmirrorPolicyModifyParams{
		HTTPClient: client,
	}
}

/*
SnapmirrorPolicyModifyParams contains all the parameters to send to the API endpoint

	for the snapmirror policy modify operation.

	Typically these are written to a http.Request.
*/
type SnapmirrorPolicyModifyParams struct {

	/* Info.

	   Information on the SnapMirror policy
	*/
	Info *models.SnapmirrorPolicy

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning. When doing a POST, PATCH, or DELETE operation on a single record, the default is 0 seconds.  This means that if an asynchronous operation is started, the server immediately returns HTTP code 202 (Accepted) along with a link to the job.  If a non-zero value is specified for POST, PATCH, or DELETE operations, ONTAP waits that length of time to see if the job completes so it can return something other than 202.
	*/
	ReturnTimeoutQueryParameter *int64

	/* UUID.

	   Policy UUID
	*/
	UUIDPathParameter string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapmirror policy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorPolicyModifyParams) WithDefaults() *SnapmirrorPolicyModifyParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapmirror policy modify params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapmirrorPolicyModifyParams) SetDefaults() {
	var (
		returnTimeoutQueryParameterDefault = int64(0)
	)

	val := SnapmirrorPolicyModifyParams{
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snapmirror policy modify params
func (o *SnapmirrorPolicyModifyParams) WithTimeout(timeout time.Duration) *SnapmirrorPolicyModifyParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapmirror policy modify params
func (o *SnapmirrorPolicyModifyParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapmirror policy modify params
func (o *SnapmirrorPolicyModifyParams) WithContext(ctx context.Context) *SnapmirrorPolicyModifyParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapmirror policy modify params
func (o *SnapmirrorPolicyModifyParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapmirror policy modify params
func (o *SnapmirrorPolicyModifyParams) WithHTTPClient(client *http.Client) *SnapmirrorPolicyModifyParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapmirror policy modify params
func (o *SnapmirrorPolicyModifyParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithInfo adds the info to the snapmirror policy modify params
func (o *SnapmirrorPolicyModifyParams) WithInfo(info *models.SnapmirrorPolicy) *SnapmirrorPolicyModifyParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the snapmirror policy modify params
func (o *SnapmirrorPolicyModifyParams) SetInfo(info *models.SnapmirrorPolicy) {
	o.Info = info
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the snapmirror policy modify params
func (o *SnapmirrorPolicyModifyParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SnapmirrorPolicyModifyParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the snapmirror policy modify params
func (o *SnapmirrorPolicyModifyParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithUUIDPathParameter adds the uuid to the snapmirror policy modify params
func (o *SnapmirrorPolicyModifyParams) WithUUIDPathParameter(uuid string) *SnapmirrorPolicyModifyParams {
	o.SetUUIDPathParameter(uuid)
	return o
}

// SetUUIDPathParameter adds the uuid to the snapmirror policy modify params
func (o *SnapmirrorPolicyModifyParams) SetUUIDPathParameter(uuid string) {
	o.UUIDPathParameter = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *SnapmirrorPolicyModifyParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Info != nil {
		if err := r.SetBodyParam(o.Info); err != nil {
			return err
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

	// path param uuid
	if err := r.SetPathParam("uuid", o.UUIDPathParameter); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
