// Code generated by go-swagger; DO NOT EDIT.

package networking

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

// NewFcZoneCollectionGetParams creates a new FcZoneCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewFcZoneCollectionGetParams() *FcZoneCollectionGetParams {
	return &FcZoneCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewFcZoneCollectionGetParamsWithTimeout creates a new FcZoneCollectionGetParams object
// with the ability to set a timeout on a request.
func NewFcZoneCollectionGetParamsWithTimeout(timeout time.Duration) *FcZoneCollectionGetParams {
	return &FcZoneCollectionGetParams{
		timeout: timeout,
	}
}

// NewFcZoneCollectionGetParamsWithContext creates a new FcZoneCollectionGetParams object
// with the ability to set a context for a request.
func NewFcZoneCollectionGetParamsWithContext(ctx context.Context) *FcZoneCollectionGetParams {
	return &FcZoneCollectionGetParams{
		Context: ctx,
	}
}

// NewFcZoneCollectionGetParamsWithHTTPClient creates a new FcZoneCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewFcZoneCollectionGetParamsWithHTTPClient(client *http.Client) *FcZoneCollectionGetParams {
	return &FcZoneCollectionGetParams{
		HTTPClient: client,
	}
}

/*
FcZoneCollectionGetParams contains all the parameters to send to the API endpoint

	for the fc zone collection get operation.

	Typically these are written to a http.Request.
*/
type FcZoneCollectionGetParams struct {

	/* CacheAge.

	   Filter by cache.age
	*/
	CacheAgeQueryParameter *string

	/* CacheIsCurrent.

	   Filter by cache.is_current
	*/
	CacheIsCurrentQueryParameter *bool

	/* CacheMaximumAge.

	   The maximum age of data in the Fibre Channel fabric cache before it should be refreshed from the fabric. The default is 15 minutes.

	   Format: iso8601
	   Default: "15 minutes"
	*/
	CacheMaximumAgeQueryParameter *string

	/* CacheUpdateTime.

	   Filter by cache.update_time
	*/
	CacheUpdateTimeQueryParameter *string

	/* FabricName.

	   The WWN of the primary switch of the Fibre Channel fabric.

	*/
	FabricNamePathParameter string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* MembersName.

	   Filter by members.name
	*/
	MembersNameQueryParameter *string

	/* MembersType.

	   Filter by members.type
	*/
	MembersTypeQueryParameter *string

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

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

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the fc zone collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcZoneCollectionGetParams) WithDefaults() *FcZoneCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the fc zone collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *FcZoneCollectionGetParams) SetDefaults() {
	var (
		cacheMaximumAgeQueryParameterDefault = string("15 minutes")

		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := FcZoneCollectionGetParams{
		CacheMaximumAgeQueryParameter: &cacheMaximumAgeQueryParameterDefault,
		ReturnRecordsQueryParameter:   &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter:   &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithTimeout(timeout time.Duration) *FcZoneCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithContext(ctx context.Context) *FcZoneCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithHTTPClient(client *http.Client) *FcZoneCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCacheAgeQueryParameter adds the cacheAge to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithCacheAgeQueryParameter(cacheAge *string) *FcZoneCollectionGetParams {
	o.SetCacheAgeQueryParameter(cacheAge)
	return o
}

// SetCacheAgeQueryParameter adds the cacheAge to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetCacheAgeQueryParameter(cacheAge *string) {
	o.CacheAgeQueryParameter = cacheAge
}

// WithCacheIsCurrentQueryParameter adds the cacheIsCurrent to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithCacheIsCurrentQueryParameter(cacheIsCurrent *bool) *FcZoneCollectionGetParams {
	o.SetCacheIsCurrentQueryParameter(cacheIsCurrent)
	return o
}

// SetCacheIsCurrentQueryParameter adds the cacheIsCurrent to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetCacheIsCurrentQueryParameter(cacheIsCurrent *bool) {
	o.CacheIsCurrentQueryParameter = cacheIsCurrent
}

// WithCacheMaximumAgeQueryParameter adds the cacheMaximumAge to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithCacheMaximumAgeQueryParameter(cacheMaximumAge *string) *FcZoneCollectionGetParams {
	o.SetCacheMaximumAgeQueryParameter(cacheMaximumAge)
	return o
}

// SetCacheMaximumAgeQueryParameter adds the cacheMaximumAge to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetCacheMaximumAgeQueryParameter(cacheMaximumAge *string) {
	o.CacheMaximumAgeQueryParameter = cacheMaximumAge
}

// WithCacheUpdateTimeQueryParameter adds the cacheUpdateTime to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithCacheUpdateTimeQueryParameter(cacheUpdateTime *string) *FcZoneCollectionGetParams {
	o.SetCacheUpdateTimeQueryParameter(cacheUpdateTime)
	return o
}

// SetCacheUpdateTimeQueryParameter adds the cacheUpdateTime to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetCacheUpdateTimeQueryParameter(cacheUpdateTime *string) {
	o.CacheUpdateTimeQueryParameter = cacheUpdateTime
}

// WithFabricNamePathParameter adds the fabricName to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithFabricNamePathParameter(fabricName string) *FcZoneCollectionGetParams {
	o.SetFabricNamePathParameter(fabricName)
	return o
}

// SetFabricNamePathParameter adds the fabricName to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetFabricNamePathParameter(fabricName string) {
	o.FabricNamePathParameter = fabricName
}

// WithFieldsQueryParameter adds the fields to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithFieldsQueryParameter(fields []string) *FcZoneCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithMaxRecordsQueryParameter adds the maxRecords to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *FcZoneCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithMembersNameQueryParameter adds the membersName to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithMembersNameQueryParameter(membersName *string) *FcZoneCollectionGetParams {
	o.SetMembersNameQueryParameter(membersName)
	return o
}

// SetMembersNameQueryParameter adds the membersName to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetMembersNameQueryParameter(membersName *string) {
	o.MembersNameQueryParameter = membersName
}

// WithMembersTypeQueryParameter adds the membersType to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithMembersTypeQueryParameter(membersType *string) *FcZoneCollectionGetParams {
	o.SetMembersTypeQueryParameter(membersType)
	return o
}

// SetMembersTypeQueryParameter adds the membersType to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetMembersTypeQueryParameter(membersType *string) {
	o.MembersTypeQueryParameter = membersType
}

// WithNameQueryParameter adds the name to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithNameQueryParameter(name *string) *FcZoneCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *FcZoneCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithReturnRecordsQueryParameter adds the returnRecords to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *FcZoneCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the fc zone collection get params
func (o *FcZoneCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *FcZoneCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the fc zone collection get params
func (o *FcZoneCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WriteToRequest writes these params to a swagger request
func (o *FcZoneCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CacheAgeQueryParameter != nil {

		// query param cache.age
		var qrCacheAge string

		if o.CacheAgeQueryParameter != nil {
			qrCacheAge = *o.CacheAgeQueryParameter
		}
		qCacheAge := qrCacheAge
		if qCacheAge != "" {

			if err := r.SetQueryParam("cache.age", qCacheAge); err != nil {
				return err
			}
		}
	}

	if o.CacheIsCurrentQueryParameter != nil {

		// query param cache.is_current
		var qrCacheIsCurrent bool

		if o.CacheIsCurrentQueryParameter != nil {
			qrCacheIsCurrent = *o.CacheIsCurrentQueryParameter
		}
		qCacheIsCurrent := swag.FormatBool(qrCacheIsCurrent)
		if qCacheIsCurrent != "" {

			if err := r.SetQueryParam("cache.is_current", qCacheIsCurrent); err != nil {
				return err
			}
		}
	}

	if o.CacheMaximumAgeQueryParameter != nil {

		// query param cache.maximum_age
		var qrCacheMaximumAge string

		if o.CacheMaximumAgeQueryParameter != nil {
			qrCacheMaximumAge = *o.CacheMaximumAgeQueryParameter
		}
		qCacheMaximumAge := qrCacheMaximumAge
		if qCacheMaximumAge != "" {

			if err := r.SetQueryParam("cache.maximum_age", qCacheMaximumAge); err != nil {
				return err
			}
		}
	}

	if o.CacheUpdateTimeQueryParameter != nil {

		// query param cache.update_time
		var qrCacheUpdateTime string

		if o.CacheUpdateTimeQueryParameter != nil {
			qrCacheUpdateTime = *o.CacheUpdateTimeQueryParameter
		}
		qCacheUpdateTime := qrCacheUpdateTime
		if qCacheUpdateTime != "" {

			if err := r.SetQueryParam("cache.update_time", qCacheUpdateTime); err != nil {
				return err
			}
		}
	}

	// path param fabric.name
	if err := r.SetPathParam("fabric.name", o.FabricNamePathParameter); err != nil {
		return err
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
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

	if o.MembersNameQueryParameter != nil {

		// query param members.name
		var qrMembersName string

		if o.MembersNameQueryParameter != nil {
			qrMembersName = *o.MembersNameQueryParameter
		}
		qMembersName := qrMembersName
		if qMembersName != "" {

			if err := r.SetQueryParam("members.name", qMembersName); err != nil {
				return err
			}
		}
	}

	if o.MembersTypeQueryParameter != nil {

		// query param members.type
		var qrMembersType string

		if o.MembersTypeQueryParameter != nil {
			qrMembersType = *o.MembersTypeQueryParameter
		}
		qMembersType := qrMembersType
		if qMembersType != "" {

			if err := r.SetQueryParam("members.type", qMembersType); err != nil {
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamFcZoneCollectionGet binds the parameter fields
func (o *FcZoneCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamFcZoneCollectionGet binds the parameter order_by
func (o *FcZoneCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
