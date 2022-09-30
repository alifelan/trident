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
)

// NewNetbiosCollectionGetParams creates a new NetbiosCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewNetbiosCollectionGetParams() *NetbiosCollectionGetParams {
	return &NetbiosCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewNetbiosCollectionGetParamsWithTimeout creates a new NetbiosCollectionGetParams object
// with the ability to set a timeout on a request.
func NewNetbiosCollectionGetParamsWithTimeout(timeout time.Duration) *NetbiosCollectionGetParams {
	return &NetbiosCollectionGetParams{
		timeout: timeout,
	}
}

// NewNetbiosCollectionGetParamsWithContext creates a new NetbiosCollectionGetParams object
// with the ability to set a context for a request.
func NewNetbiosCollectionGetParamsWithContext(ctx context.Context) *NetbiosCollectionGetParams {
	return &NetbiosCollectionGetParams{
		Context: ctx,
	}
}

// NewNetbiosCollectionGetParamsWithHTTPClient creates a new NetbiosCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewNetbiosCollectionGetParamsWithHTTPClient(client *http.Client) *NetbiosCollectionGetParams {
	return &NetbiosCollectionGetParams{
		HTTPClient: client,
	}
}

/*
NetbiosCollectionGetParams contains all the parameters to send to the API endpoint

	for the netbios collection get operation.

	Typically these are written to a http.Request.
*/
type NetbiosCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Interfaces.

	   Filter by interfaces
	*/
	InterfacesQueryParameter *string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* Mode.

	   Filter by mode
	*/
	ModeQueryParameter *string

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

	/* NameRegistrationType.

	   Filter by name_registration_type
	*/
	NameRegistrationTypeQueryParameter *string

	/* NodeName.

	   Filter by node.name
	*/
	NodeNameQueryParameter *string

	/* NodeUUID.

	   Filter by node.uuid
	*/
	NodeUUIDQueryParameter *string

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

	/* Scope.

	   Filter by scope
	*/
	ScopeQueryParameter *string

	/* State.

	   Filter by state
	*/
	StateQueryParameter *string

	/* Suffix.

	   Filter by suffix
	*/
	SuffixQueryParameter *string

	/* SvmName.

	   Filter by svm.name
	*/
	SVMNameQueryParameter *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SVMUUIDQueryParameter *string

	/* TimeLeft.

	   Filter by time_left
	*/
	TimeLeftQueryParameter *int64

	/* WinsServersIP.

	   Filter by wins_servers.ip
	*/
	WinsServersIPQueryParameter *string

	/* WinsServersState.

	   Filter by wins_servers.state
	*/
	WinsServersStateQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the netbios collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetbiosCollectionGetParams) WithDefaults() *NetbiosCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the netbios collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *NetbiosCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := NetbiosCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithTimeout(timeout time.Duration) *NetbiosCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithContext(ctx context.Context) *NetbiosCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithHTTPClient(client *http.Client) *NetbiosCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFieldsQueryParameter adds the fields to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithFieldsQueryParameter(fields []string) *NetbiosCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithInterfacesQueryParameter adds the interfaces to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithInterfacesQueryParameter(interfaces *string) *NetbiosCollectionGetParams {
	o.SetInterfacesQueryParameter(interfaces)
	return o
}

// SetInterfacesQueryParameter adds the interfaces to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetInterfacesQueryParameter(interfaces *string) {
	o.InterfacesQueryParameter = interfaces
}

// WithMaxRecordsQueryParameter adds the maxRecords to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *NetbiosCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithModeQueryParameter adds the mode to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithModeQueryParameter(mode *string) *NetbiosCollectionGetParams {
	o.SetModeQueryParameter(mode)
	return o
}

// SetModeQueryParameter adds the mode to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetModeQueryParameter(mode *string) {
	o.ModeQueryParameter = mode
}

// WithNameQueryParameter adds the name to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithNameQueryParameter(name *string) *NetbiosCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithNameRegistrationTypeQueryParameter adds the nameRegistrationType to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithNameRegistrationTypeQueryParameter(nameRegistrationType *string) *NetbiosCollectionGetParams {
	o.SetNameRegistrationTypeQueryParameter(nameRegistrationType)
	return o
}

// SetNameRegistrationTypeQueryParameter adds the nameRegistrationType to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetNameRegistrationTypeQueryParameter(nameRegistrationType *string) {
	o.NameRegistrationTypeQueryParameter = nameRegistrationType
}

// WithNodeNameQueryParameter adds the nodeName to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithNodeNameQueryParameter(nodeName *string) *NetbiosCollectionGetParams {
	o.SetNodeNameQueryParameter(nodeName)
	return o
}

// SetNodeNameQueryParameter adds the nodeName to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetNodeNameQueryParameter(nodeName *string) {
	o.NodeNameQueryParameter = nodeName
}

// WithNodeUUIDQueryParameter adds the nodeUUID to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithNodeUUIDQueryParameter(nodeUUID *string) *NetbiosCollectionGetParams {
	o.SetNodeUUIDQueryParameter(nodeUUID)
	return o
}

// SetNodeUUIDQueryParameter adds the nodeUuid to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetNodeUUIDQueryParameter(nodeUUID *string) {
	o.NodeUUIDQueryParameter = nodeUUID
}

// WithOrderByQueryParameter adds the orderBy to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *NetbiosCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *NetbiosCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *NetbiosCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithScopeQueryParameter adds the scope to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithScopeQueryParameter(scope *string) *NetbiosCollectionGetParams {
	o.SetScopeQueryParameter(scope)
	return o
}

// SetScopeQueryParameter adds the scope to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetScopeQueryParameter(scope *string) {
	o.ScopeQueryParameter = scope
}

// WithStateQueryParameter adds the state to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithStateQueryParameter(state *string) *NetbiosCollectionGetParams {
	o.SetStateQueryParameter(state)
	return o
}

// SetStateQueryParameter adds the state to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetStateQueryParameter(state *string) {
	o.StateQueryParameter = state
}

// WithSuffixQueryParameter adds the suffix to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithSuffixQueryParameter(suffix *string) *NetbiosCollectionGetParams {
	o.SetSuffixQueryParameter(suffix)
	return o
}

// SetSuffixQueryParameter adds the suffix to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetSuffixQueryParameter(suffix *string) {
	o.SuffixQueryParameter = suffix
}

// WithSVMNameQueryParameter adds the svmName to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithSVMNameQueryParameter(svmName *string) *NetbiosCollectionGetParams {
	o.SetSVMNameQueryParameter(svmName)
	return o
}

// SetSVMNameQueryParameter adds the svmName to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetSVMNameQueryParameter(svmName *string) {
	o.SVMNameQueryParameter = svmName
}

// WithSVMUUIDQueryParameter adds the svmUUID to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithSVMUUIDQueryParameter(svmUUID *string) *NetbiosCollectionGetParams {
	o.SetSVMUUIDQueryParameter(svmUUID)
	return o
}

// SetSVMUUIDQueryParameter adds the svmUuid to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetSVMUUIDQueryParameter(svmUUID *string) {
	o.SVMUUIDQueryParameter = svmUUID
}

// WithTimeLeftQueryParameter adds the timeLeft to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithTimeLeftQueryParameter(timeLeft *int64) *NetbiosCollectionGetParams {
	o.SetTimeLeftQueryParameter(timeLeft)
	return o
}

// SetTimeLeftQueryParameter adds the timeLeft to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetTimeLeftQueryParameter(timeLeft *int64) {
	o.TimeLeftQueryParameter = timeLeft
}

// WithWinsServersIPQueryParameter adds the winsServersIP to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithWinsServersIPQueryParameter(winsServersIP *string) *NetbiosCollectionGetParams {
	o.SetWinsServersIPQueryParameter(winsServersIP)
	return o
}

// SetWinsServersIPQueryParameter adds the winsServersIp to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetWinsServersIPQueryParameter(winsServersIP *string) {
	o.WinsServersIPQueryParameter = winsServersIP
}

// WithWinsServersStateQueryParameter adds the winsServersState to the netbios collection get params
func (o *NetbiosCollectionGetParams) WithWinsServersStateQueryParameter(winsServersState *string) *NetbiosCollectionGetParams {
	o.SetWinsServersStateQueryParameter(winsServersState)
	return o
}

// SetWinsServersStateQueryParameter adds the winsServersState to the netbios collection get params
func (o *NetbiosCollectionGetParams) SetWinsServersStateQueryParameter(winsServersState *string) {
	o.WinsServersStateQueryParameter = winsServersState
}

// WriteToRequest writes these params to a swagger request
func (o *NetbiosCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.InterfacesQueryParameter != nil {

		// query param interfaces
		var qrInterfaces string

		if o.InterfacesQueryParameter != nil {
			qrInterfaces = *o.InterfacesQueryParameter
		}
		qInterfaces := qrInterfaces
		if qInterfaces != "" {

			if err := r.SetQueryParam("interfaces", qInterfaces); err != nil {
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

	if o.ModeQueryParameter != nil {

		// query param mode
		var qrMode string

		if o.ModeQueryParameter != nil {
			qrMode = *o.ModeQueryParameter
		}
		qMode := qrMode
		if qMode != "" {

			if err := r.SetQueryParam("mode", qMode); err != nil {
				return err
			}
		}
	}

	if o.NameQueryParameter != nil {

		// query param name
		var qrName string

		if o.NameQueryParameter != nil {
			qrName = *o.NameQueryParameter
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.NameRegistrationTypeQueryParameter != nil {

		// query param name_registration_type
		var qrNameRegistrationType string

		if o.NameRegistrationTypeQueryParameter != nil {
			qrNameRegistrationType = *o.NameRegistrationTypeQueryParameter
		}
		qNameRegistrationType := qrNameRegistrationType
		if qNameRegistrationType != "" {

			if err := r.SetQueryParam("name_registration_type", qNameRegistrationType); err != nil {
				return err
			}
		}
	}

	if o.NodeNameQueryParameter != nil {

		// query param node.name
		var qrNodeName string

		if o.NodeNameQueryParameter != nil {
			qrNodeName = *o.NodeNameQueryParameter
		}
		qNodeName := qrNodeName
		if qNodeName != "" {

			if err := r.SetQueryParam("node.name", qNodeName); err != nil {
				return err
			}
		}
	}

	if o.NodeUUIDQueryParameter != nil {

		// query param node.uuid
		var qrNodeUUID string

		if o.NodeUUIDQueryParameter != nil {
			qrNodeUUID = *o.NodeUUIDQueryParameter
		}
		qNodeUUID := qrNodeUUID
		if qNodeUUID != "" {

			if err := r.SetQueryParam("node.uuid", qNodeUUID); err != nil {
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

	if o.ScopeQueryParameter != nil {

		// query param scope
		var qrScope string

		if o.ScopeQueryParameter != nil {
			qrScope = *o.ScopeQueryParameter
		}
		qScope := qrScope
		if qScope != "" {

			if err := r.SetQueryParam("scope", qScope); err != nil {
				return err
			}
		}
	}

	if o.StateQueryParameter != nil {

		// query param state
		var qrState string

		if o.StateQueryParameter != nil {
			qrState = *o.StateQueryParameter
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if o.SuffixQueryParameter != nil {

		// query param suffix
		var qrSuffix string

		if o.SuffixQueryParameter != nil {
			qrSuffix = *o.SuffixQueryParameter
		}
		qSuffix := qrSuffix
		if qSuffix != "" {

			if err := r.SetQueryParam("suffix", qSuffix); err != nil {
				return err
			}
		}
	}

	if o.SVMNameQueryParameter != nil {

		// query param svm.name
		var qrSvmName string

		if o.SVMNameQueryParameter != nil {
			qrSvmName = *o.SVMNameQueryParameter
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SVMUUIDQueryParameter != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SVMUUIDQueryParameter != nil {
			qrSvmUUID = *o.SVMUUIDQueryParameter
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if o.TimeLeftQueryParameter != nil {

		// query param time_left
		var qrTimeLeft int64

		if o.TimeLeftQueryParameter != nil {
			qrTimeLeft = *o.TimeLeftQueryParameter
		}
		qTimeLeft := swag.FormatInt64(qrTimeLeft)
		if qTimeLeft != "" {

			if err := r.SetQueryParam("time_left", qTimeLeft); err != nil {
				return err
			}
		}
	}

	if o.WinsServersIPQueryParameter != nil {

		// query param wins_servers.ip
		var qrWinsServersIP string

		if o.WinsServersIPQueryParameter != nil {
			qrWinsServersIP = *o.WinsServersIPQueryParameter
		}
		qWinsServersIP := qrWinsServersIP
		if qWinsServersIP != "" {

			if err := r.SetQueryParam("wins_servers.ip", qWinsServersIP); err != nil {
				return err
			}
		}
	}

	if o.WinsServersStateQueryParameter != nil {

		// query param wins_servers.state
		var qrWinsServersState string

		if o.WinsServersStateQueryParameter != nil {
			qrWinsServersState = *o.WinsServersStateQueryParameter
		}
		qWinsServersState := qrWinsServersState
		if qWinsServersState != "" {

			if err := r.SetQueryParam("wins_servers.state", qWinsServersState); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamNetbiosCollectionGet binds the parameter fields
func (o *NetbiosCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamNetbiosCollectionGet binds the parameter order_by
func (o *NetbiosCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
