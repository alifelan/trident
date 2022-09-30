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
)

// NewPublickeyCollectionGetParams creates a new PublickeyCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPublickeyCollectionGetParams() *PublickeyCollectionGetParams {
	return &PublickeyCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPublickeyCollectionGetParamsWithTimeout creates a new PublickeyCollectionGetParams object
// with the ability to set a timeout on a request.
func NewPublickeyCollectionGetParamsWithTimeout(timeout time.Duration) *PublickeyCollectionGetParams {
	return &PublickeyCollectionGetParams{
		timeout: timeout,
	}
}

// NewPublickeyCollectionGetParamsWithContext creates a new PublickeyCollectionGetParams object
// with the ability to set a context for a request.
func NewPublickeyCollectionGetParamsWithContext(ctx context.Context) *PublickeyCollectionGetParams {
	return &PublickeyCollectionGetParams{
		Context: ctx,
	}
}

// NewPublickeyCollectionGetParamsWithHTTPClient creates a new PublickeyCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewPublickeyCollectionGetParamsWithHTTPClient(client *http.Client) *PublickeyCollectionGetParams {
	return &PublickeyCollectionGetParams{
		HTTPClient: client,
	}
}

/*
PublickeyCollectionGetParams contains all the parameters to send to the API endpoint

	for the publickey collection get operation.

	Typically these are written to a http.Request.
*/
type PublickeyCollectionGetParams struct {

	/* AccountName.

	   Filter by account.name
	*/
	AccountNameQueryParameter *string

	/* Comment.

	   Filter by comment
	*/
	CommentQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* Index.

	   Filter by index
	*/
	IndexQueryParameter *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* ObfuscatedFingerprint.

	   Filter by obfuscated_fingerprint
	*/
	ObfuscatedFingerprintQueryParameter *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderByQueryParameter []string

	/* OwnerName.

	   Filter by owner.name
	*/
	OwnerNameQueryParameter *string

	/* OwnerUUID.

	   Filter by owner.uuid
	*/
	OwnerUUIDQueryParameter *string

	/* PublicKey.

	   Filter by public_key
	*/
	PublicKeyQueryParameter *string

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

	/* ShaFingerprint.

	   Filter by sha_fingerprint
	*/
	ShaFingerprintQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the publickey collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublickeyCollectionGetParams) WithDefaults() *PublickeyCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the publickey collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PublickeyCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := PublickeyCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithTimeout(timeout time.Duration) *PublickeyCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithContext(ctx context.Context) *PublickeyCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithHTTPClient(client *http.Client) *PublickeyCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAccountNameQueryParameter adds the accountName to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithAccountNameQueryParameter(accountName *string) *PublickeyCollectionGetParams {
	o.SetAccountNameQueryParameter(accountName)
	return o
}

// SetAccountNameQueryParameter adds the accountName to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetAccountNameQueryParameter(accountName *string) {
	o.AccountNameQueryParameter = accountName
}

// WithCommentQueryParameter adds the comment to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithCommentQueryParameter(comment *string) *PublickeyCollectionGetParams {
	o.SetCommentQueryParameter(comment)
	return o
}

// SetCommentQueryParameter adds the comment to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetCommentQueryParameter(comment *string) {
	o.CommentQueryParameter = comment
}

// WithFieldsQueryParameter adds the fields to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithFieldsQueryParameter(fields []string) *PublickeyCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithIndexQueryParameter adds the index to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithIndexQueryParameter(index *int64) *PublickeyCollectionGetParams {
	o.SetIndexQueryParameter(index)
	return o
}

// SetIndexQueryParameter adds the index to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetIndexQueryParameter(index *int64) {
	o.IndexQueryParameter = index
}

// WithMaxRecordsQueryParameter adds the maxRecords to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *PublickeyCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithObfuscatedFingerprintQueryParameter adds the obfuscatedFingerprint to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithObfuscatedFingerprintQueryParameter(obfuscatedFingerprint *string) *PublickeyCollectionGetParams {
	o.SetObfuscatedFingerprintQueryParameter(obfuscatedFingerprint)
	return o
}

// SetObfuscatedFingerprintQueryParameter adds the obfuscatedFingerprint to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetObfuscatedFingerprintQueryParameter(obfuscatedFingerprint *string) {
	o.ObfuscatedFingerprintQueryParameter = obfuscatedFingerprint
}

// WithOrderByQueryParameter adds the orderBy to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *PublickeyCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithOwnerNameQueryParameter adds the ownerName to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithOwnerNameQueryParameter(ownerName *string) *PublickeyCollectionGetParams {
	o.SetOwnerNameQueryParameter(ownerName)
	return o
}

// SetOwnerNameQueryParameter adds the ownerName to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetOwnerNameQueryParameter(ownerName *string) {
	o.OwnerNameQueryParameter = ownerName
}

// WithOwnerUUIDQueryParameter adds the ownerUUID to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithOwnerUUIDQueryParameter(ownerUUID *string) *PublickeyCollectionGetParams {
	o.SetOwnerUUIDQueryParameter(ownerUUID)
	return o
}

// SetOwnerUUIDQueryParameter adds the ownerUuid to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetOwnerUUIDQueryParameter(ownerUUID *string) {
	o.OwnerUUIDQueryParameter = ownerUUID
}

// WithPublicKeyQueryParameter adds the publicKey to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithPublicKeyQueryParameter(publicKey *string) *PublickeyCollectionGetParams {
	o.SetPublicKeyQueryParameter(publicKey)
	return o
}

// SetPublicKeyQueryParameter adds the publicKey to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetPublicKeyQueryParameter(publicKey *string) {
	o.PublicKeyQueryParameter = publicKey
}

// WithReturnRecordsQueryParameter adds the returnRecords to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *PublickeyCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *PublickeyCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithScopeQueryParameter adds the scope to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithScopeQueryParameter(scope *string) *PublickeyCollectionGetParams {
	o.SetScopeQueryParameter(scope)
	return o
}

// SetScopeQueryParameter adds the scope to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetScopeQueryParameter(scope *string) {
	o.ScopeQueryParameter = scope
}

// WithShaFingerprintQueryParameter adds the shaFingerprint to the publickey collection get params
func (o *PublickeyCollectionGetParams) WithShaFingerprintQueryParameter(shaFingerprint *string) *PublickeyCollectionGetParams {
	o.SetShaFingerprintQueryParameter(shaFingerprint)
	return o
}

// SetShaFingerprintQueryParameter adds the shaFingerprint to the publickey collection get params
func (o *PublickeyCollectionGetParams) SetShaFingerprintQueryParameter(shaFingerprint *string) {
	o.ShaFingerprintQueryParameter = shaFingerprint
}

// WriteToRequest writes these params to a swagger request
func (o *PublickeyCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AccountNameQueryParameter != nil {

		// query param account.name
		var qrAccountName string

		if o.AccountNameQueryParameter != nil {
			qrAccountName = *o.AccountNameQueryParameter
		}
		qAccountName := qrAccountName
		if qAccountName != "" {

			if err := r.SetQueryParam("account.name", qAccountName); err != nil {
				return err
			}
		}
	}

	if o.CommentQueryParameter != nil {

		// query param comment
		var qrComment string

		if o.CommentQueryParameter != nil {
			qrComment = *o.CommentQueryParameter
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if o.FieldsQueryParameter != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.IndexQueryParameter != nil {

		// query param index
		var qrIndex int64

		if o.IndexQueryParameter != nil {
			qrIndex = *o.IndexQueryParameter
		}
		qIndex := swag.FormatInt64(qrIndex)
		if qIndex != "" {

			if err := r.SetQueryParam("index", qIndex); err != nil {
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

	if o.ObfuscatedFingerprintQueryParameter != nil {

		// query param obfuscated_fingerprint
		var qrObfuscatedFingerprint string

		if o.ObfuscatedFingerprintQueryParameter != nil {
			qrObfuscatedFingerprint = *o.ObfuscatedFingerprintQueryParameter
		}
		qObfuscatedFingerprint := qrObfuscatedFingerprint
		if qObfuscatedFingerprint != "" {

			if err := r.SetQueryParam("obfuscated_fingerprint", qObfuscatedFingerprint); err != nil {
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

	if o.OwnerNameQueryParameter != nil {

		// query param owner.name
		var qrOwnerName string

		if o.OwnerNameQueryParameter != nil {
			qrOwnerName = *o.OwnerNameQueryParameter
		}
		qOwnerName := qrOwnerName
		if qOwnerName != "" {

			if err := r.SetQueryParam("owner.name", qOwnerName); err != nil {
				return err
			}
		}
	}

	if o.OwnerUUIDQueryParameter != nil {

		// query param owner.uuid
		var qrOwnerUUID string

		if o.OwnerUUIDQueryParameter != nil {
			qrOwnerUUID = *o.OwnerUUIDQueryParameter
		}
		qOwnerUUID := qrOwnerUUID
		if qOwnerUUID != "" {

			if err := r.SetQueryParam("owner.uuid", qOwnerUUID); err != nil {
				return err
			}
		}
	}

	if o.PublicKeyQueryParameter != nil {

		// query param public_key
		var qrPublicKey string

		if o.PublicKeyQueryParameter != nil {
			qrPublicKey = *o.PublicKeyQueryParameter
		}
		qPublicKey := qrPublicKey
		if qPublicKey != "" {

			if err := r.SetQueryParam("public_key", qPublicKey); err != nil {
				return err
			}
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

	if o.ShaFingerprintQueryParameter != nil {

		// query param sha_fingerprint
		var qrShaFingerprint string

		if o.ShaFingerprintQueryParameter != nil {
			qrShaFingerprint = *o.ShaFingerprintQueryParameter
		}
		qShaFingerprint := qrShaFingerprint
		if qShaFingerprint != "" {

			if err := r.SetQueryParam("sha_fingerprint", qShaFingerprint); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPublickeyCollectionGet binds the parameter fields
func (o *PublickeyCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamPublickeyCollectionGet binds the parameter order_by
func (o *PublickeyCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
