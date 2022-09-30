// Code generated by go-swagger; DO NOT EDIT.

package support

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

// NewSnmpTraphostsCollectionGetParams creates a new SnmpTraphostsCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnmpTraphostsCollectionGetParams() *SnmpTraphostsCollectionGetParams {
	return &SnmpTraphostsCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnmpTraphostsCollectionGetParamsWithTimeout creates a new SnmpTraphostsCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSnmpTraphostsCollectionGetParamsWithTimeout(timeout time.Duration) *SnmpTraphostsCollectionGetParams {
	return &SnmpTraphostsCollectionGetParams{
		timeout: timeout,
	}
}

// NewSnmpTraphostsCollectionGetParamsWithContext creates a new SnmpTraphostsCollectionGetParams object
// with the ability to set a context for a request.
func NewSnmpTraphostsCollectionGetParamsWithContext(ctx context.Context) *SnmpTraphostsCollectionGetParams {
	return &SnmpTraphostsCollectionGetParams{
		Context: ctx,
	}
}

// NewSnmpTraphostsCollectionGetParamsWithHTTPClient creates a new SnmpTraphostsCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnmpTraphostsCollectionGetParamsWithHTTPClient(client *http.Client) *SnmpTraphostsCollectionGetParams {
	return &SnmpTraphostsCollectionGetParams{
		HTTPClient: client,
	}
}

/*
SnmpTraphostsCollectionGetParams contains all the parameters to send to the API endpoint

	for the snmp traphosts collection get operation.

	Typically these are written to a http.Request.
*/
type SnmpTraphostsCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Host.

	   Filter by host
	*/
	HostQueryParameter *string

	/* IPAddress.

	   Filter by ip_address
	*/
	IPAddressQueryParameter *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecordsQueryParameter *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeoutQueryParameter *int64

	/* UserName.

	   Filter by user.name
	*/
	UserNameQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snmp traphosts collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnmpTraphostsCollectionGetParams) WithDefaults() *SnmpTraphostsCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snmp traphosts collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnmpTraphostsCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := SnmpTraphostsCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) WithTimeout(timeout time.Duration) *SnmpTraphostsCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) WithContext(ctx context.Context) *SnmpTraphostsCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) WithHTTPClient(client *http.Client) *SnmpTraphostsCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) WithFieldsQueryParameter(fields []string) *SnmpTraphostsCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithHostQueryParameter adds the host to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) WithHostQueryParameter(host *string) *SnmpTraphostsCollectionGetParams {
	o.SetHostQueryParameter(host)
	return o
}

// SetHostQueryParameter adds the host to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) SetHostQueryParameter(host *string) {
	o.HostQueryParameter = host
}

// WithIPAddressQueryParameter adds the iPAddress to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) WithIPAddressQueryParameter(iPAddress *string) *SnmpTraphostsCollectionGetParams {
	o.SetIPAddressQueryParameter(iPAddress)
	return o
}

// SetIPAddressQueryParameter adds the ipAddress to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) SetIPAddressQueryParameter(iPAddress *string) {
	o.IPAddressQueryParameter = iPAddress
}

// WithMaxRecordsQueryParameter adds the maxRecords to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *SnmpTraphostsCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithOrderByQueryParameter adds the orderBy to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *SnmpTraphostsCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *SnmpTraphostsCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *SnmpTraphostsCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithUserNameQueryParameter adds the userName to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) WithUserNameQueryParameter(userName *string) *SnmpTraphostsCollectionGetParams {
	o.SetUserNameQueryParameter(userName)
	return o
}

// SetUserNameQueryParameter adds the userName to the snmp traphosts collection get params
func (o *SnmpTraphostsCollectionGetParams) SetUserNameQueryParameter(userName *string) {
	o.UserNameQueryParameter = userName
}

// WriteToRequest writes these params to a swagger request
func (o *SnmpTraphostsCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.HostQueryParameter != nil {

		// query param host
		var qrHost string

		if o.HostQueryParameter != nil {
			qrHost = *o.HostQueryParameter
		}
		qHost := qrHost
		if qHost != "" {

			if err := r.SetQueryParam("host", qHost); err != nil {
				return err
			}
		}
	}

	if o.IPAddressQueryParameter != nil {

		// query param ip_address
		var qrIPAddress string

		if o.IPAddressQueryParameter != nil {
			qrIPAddress = *o.IPAddressQueryParameter
		}
		qIPAddress := qrIPAddress
		if qIPAddress != "" {

			if err := r.SetQueryParam("ip_address", qIPAddress); err != nil {
				return err
			}
		}
	}

	if o.MaxRecordsQueryParameter != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecordsQueryParameter != nil {
			qrMaxRecords = *o.MaxRecordsQueryParameter
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.OrderByQueryParameter != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
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

	if o.UserNameQueryParameter != nil {

		// query param user.name
		var qrUserName string

		if o.UserNameQueryParameter != nil {
			qrUserName = *o.UserNameQueryParameter
		}
		qUserName := qrUserName
		if qUserName != "" {

			if err := r.SetQueryParam("user.name", qUserName); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSnmpTraphostsCollectionGet binds the parameter fields
func (o *SnmpTraphostsCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamSnmpTraphostsCollectionGet binds the parameter order_by
func (o *SnmpTraphostsCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderByQueryParameter

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
