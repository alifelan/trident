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

// NewAccountCollectionGetParams creates a new AccountCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewAccountCollectionGetParams() *AccountCollectionGetParams {
	return &AccountCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewAccountCollectionGetParamsWithTimeout creates a new AccountCollectionGetParams object
// with the ability to set a timeout on a request.
func NewAccountCollectionGetParamsWithTimeout(timeout time.Duration) *AccountCollectionGetParams {
	return &AccountCollectionGetParams{
		timeout: timeout,
	}
}

// NewAccountCollectionGetParamsWithContext creates a new AccountCollectionGetParams object
// with the ability to set a context for a request.
func NewAccountCollectionGetParamsWithContext(ctx context.Context) *AccountCollectionGetParams {
	return &AccountCollectionGetParams{
		Context: ctx,
	}
}

// NewAccountCollectionGetParamsWithHTTPClient creates a new AccountCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewAccountCollectionGetParamsWithHTTPClient(client *http.Client) *AccountCollectionGetParams {
	return &AccountCollectionGetParams{
		HTTPClient: client,
	}
}

/*
AccountCollectionGetParams contains all the parameters to send to the API endpoint

	for the account collection get operation.

	Typically these are written to a http.Request.
*/
type AccountCollectionGetParams struct {

	/* ApplicationsApplication.

	   Filter by applications.application
	*/
	ApplicationsApplicationQueryParameter *string

	/* ApplicationsAuthenticationMethods.

	   Filter by applications.authentication_methods
	*/
	ApplicationsAuthenticationMethodsQueryParameter *string

	/* ApplicationsSecondAuthenticationMethod.

	   Filter by applications.second_authentication_method
	*/
	ApplicationsSecondAuthenticationMethodQueryParameter *string

	/* AuthenticationMethods.

	   Filter by authentication_methods
	*/
	AuthenticationMethodsQueryParameter *string

	/* Comment.

	   Filter by comment
	*/
	CommentQueryParameter *string

	/* Fields.

	   Specify the fields to return.
	*/
	FieldsQueryParameter []string

	/* LdapFastbind.

	   Filter by ldap_fastbind
	*/
	LdapFastbindQueryParameter *bool

	/* Locked.

	   Filter by locked
	*/
	LockedQueryParameter *bool

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecordsQueryParameter *int64

	/* Name.

	   Filter by name
	*/
	NameQueryParameter *string

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

	/* PasswordHashAlgorithm.

	   Filter by password_hash_algorithm
	*/
	PasswordHashAlgorithmQueryParameter *string

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

	/* RoleName.

	   Filter by role.name
	*/
	RoleNameQueryParameter *string

	/* Scope.

	   Filter by scope
	*/
	ScopeQueryParameter *string

	/* SslCaCertificate.

	   Filter by ssl_ca_certificate
	*/
	SslCaCertificateQueryParameter *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the account collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountCollectionGetParams) WithDefaults() *AccountCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the account collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *AccountCollectionGetParams) SetDefaults() {
	var (
		returnRecordsQueryParameterDefault = bool(true)

		returnTimeoutQueryParameterDefault = int64(15)
	)

	val := AccountCollectionGetParams{
		ReturnRecordsQueryParameter: &returnRecordsQueryParameterDefault,
		ReturnTimeoutQueryParameter: &returnTimeoutQueryParameterDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the account collection get params
func (o *AccountCollectionGetParams) WithTimeout(timeout time.Duration) *AccountCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the account collection get params
func (o *AccountCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the account collection get params
func (o *AccountCollectionGetParams) WithContext(ctx context.Context) *AccountCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the account collection get params
func (o *AccountCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the account collection get params
func (o *AccountCollectionGetParams) WithHTTPClient(client *http.Client) *AccountCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the account collection get params
func (o *AccountCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithApplicationsApplicationQueryParameter adds the applicationsApplication to the account collection get params
func (o *AccountCollectionGetParams) WithApplicationsApplicationQueryParameter(applicationsApplication *string) *AccountCollectionGetParams {
	o.SetApplicationsApplicationQueryParameter(applicationsApplication)
	return o
}

// SetApplicationsApplicationQueryParameter adds the applicationsApplication to the account collection get params
func (o *AccountCollectionGetParams) SetApplicationsApplicationQueryParameter(applicationsApplication *string) {
	o.ApplicationsApplicationQueryParameter = applicationsApplication
}

// WithApplicationsAuthenticationMethodsQueryParameter adds the applicationsAuthenticationMethods to the account collection get params
func (o *AccountCollectionGetParams) WithApplicationsAuthenticationMethodsQueryParameter(applicationsAuthenticationMethods *string) *AccountCollectionGetParams {
	o.SetApplicationsAuthenticationMethodsQueryParameter(applicationsAuthenticationMethods)
	return o
}

// SetApplicationsAuthenticationMethodsQueryParameter adds the applicationsAuthenticationMethods to the account collection get params
func (o *AccountCollectionGetParams) SetApplicationsAuthenticationMethodsQueryParameter(applicationsAuthenticationMethods *string) {
	o.ApplicationsAuthenticationMethodsQueryParameter = applicationsAuthenticationMethods
}

// WithApplicationsSecondAuthenticationMethodQueryParameter adds the applicationsSecondAuthenticationMethod to the account collection get params
func (o *AccountCollectionGetParams) WithApplicationsSecondAuthenticationMethodQueryParameter(applicationsSecondAuthenticationMethod *string) *AccountCollectionGetParams {
	o.SetApplicationsSecondAuthenticationMethodQueryParameter(applicationsSecondAuthenticationMethod)
	return o
}

// SetApplicationsSecondAuthenticationMethodQueryParameter adds the applicationsSecondAuthenticationMethod to the account collection get params
func (o *AccountCollectionGetParams) SetApplicationsSecondAuthenticationMethodQueryParameter(applicationsSecondAuthenticationMethod *string) {
	o.ApplicationsSecondAuthenticationMethodQueryParameter = applicationsSecondAuthenticationMethod
}

// WithAuthenticationMethodsQueryParameter adds the authenticationMethods to the account collection get params
func (o *AccountCollectionGetParams) WithAuthenticationMethodsQueryParameter(authenticationMethods *string) *AccountCollectionGetParams {
	o.SetAuthenticationMethodsQueryParameter(authenticationMethods)
	return o
}

// SetAuthenticationMethodsQueryParameter adds the authenticationMethods to the account collection get params
func (o *AccountCollectionGetParams) SetAuthenticationMethodsQueryParameter(authenticationMethods *string) {
	o.AuthenticationMethodsQueryParameter = authenticationMethods
}

// WithCommentQueryParameter adds the comment to the account collection get params
func (o *AccountCollectionGetParams) WithCommentQueryParameter(comment *string) *AccountCollectionGetParams {
	o.SetCommentQueryParameter(comment)
	return o
}

// SetCommentQueryParameter adds the comment to the account collection get params
func (o *AccountCollectionGetParams) SetCommentQueryParameter(comment *string) {
	o.CommentQueryParameter = comment
}

// WithFieldsQueryParameter adds the fields to the account collection get params
func (o *AccountCollectionGetParams) WithFieldsQueryParameter(fields []string) *AccountCollectionGetParams {
	o.SetFieldsQueryParameter(fields)
	return o
}

// SetFieldsQueryParameter adds the fields to the account collection get params
func (o *AccountCollectionGetParams) SetFieldsQueryParameter(fields []string) {
	o.FieldsQueryParameter = fields
}

// WithLdapFastbindQueryParameter adds the ldapFastbind to the account collection get params
func (o *AccountCollectionGetParams) WithLdapFastbindQueryParameter(ldapFastbind *bool) *AccountCollectionGetParams {
	o.SetLdapFastbindQueryParameter(ldapFastbind)
	return o
}

// SetLdapFastbindQueryParameter adds the ldapFastbind to the account collection get params
func (o *AccountCollectionGetParams) SetLdapFastbindQueryParameter(ldapFastbind *bool) {
	o.LdapFastbindQueryParameter = ldapFastbind
}

// WithLockedQueryParameter adds the locked to the account collection get params
func (o *AccountCollectionGetParams) WithLockedQueryParameter(locked *bool) *AccountCollectionGetParams {
	o.SetLockedQueryParameter(locked)
	return o
}

// SetLockedQueryParameter adds the locked to the account collection get params
func (o *AccountCollectionGetParams) SetLockedQueryParameter(locked *bool) {
	o.LockedQueryParameter = locked
}

// WithMaxRecordsQueryParameter adds the maxRecords to the account collection get params
func (o *AccountCollectionGetParams) WithMaxRecordsQueryParameter(maxRecords *int64) *AccountCollectionGetParams {
	o.SetMaxRecordsQueryParameter(maxRecords)
	return o
}

// SetMaxRecordsQueryParameter adds the maxRecords to the account collection get params
func (o *AccountCollectionGetParams) SetMaxRecordsQueryParameter(maxRecords *int64) {
	o.MaxRecordsQueryParameter = maxRecords
}

// WithNameQueryParameter adds the name to the account collection get params
func (o *AccountCollectionGetParams) WithNameQueryParameter(name *string) *AccountCollectionGetParams {
	o.SetNameQueryParameter(name)
	return o
}

// SetNameQueryParameter adds the name to the account collection get params
func (o *AccountCollectionGetParams) SetNameQueryParameter(name *string) {
	o.NameQueryParameter = name
}

// WithOrderByQueryParameter adds the orderBy to the account collection get params
func (o *AccountCollectionGetParams) WithOrderByQueryParameter(orderBy []string) *AccountCollectionGetParams {
	o.SetOrderByQueryParameter(orderBy)
	return o
}

// SetOrderByQueryParameter adds the orderBy to the account collection get params
func (o *AccountCollectionGetParams) SetOrderByQueryParameter(orderBy []string) {
	o.OrderByQueryParameter = orderBy
}

// WithOwnerNameQueryParameter adds the ownerName to the account collection get params
func (o *AccountCollectionGetParams) WithOwnerNameQueryParameter(ownerName *string) *AccountCollectionGetParams {
	o.SetOwnerNameQueryParameter(ownerName)
	return o
}

// SetOwnerNameQueryParameter adds the ownerName to the account collection get params
func (o *AccountCollectionGetParams) SetOwnerNameQueryParameter(ownerName *string) {
	o.OwnerNameQueryParameter = ownerName
}

// WithOwnerUUIDQueryParameter adds the ownerUUID to the account collection get params
func (o *AccountCollectionGetParams) WithOwnerUUIDQueryParameter(ownerUUID *string) *AccountCollectionGetParams {
	o.SetOwnerUUIDQueryParameter(ownerUUID)
	return o
}

// SetOwnerUUIDQueryParameter adds the ownerUuid to the account collection get params
func (o *AccountCollectionGetParams) SetOwnerUUIDQueryParameter(ownerUUID *string) {
	o.OwnerUUIDQueryParameter = ownerUUID
}

// WithPasswordHashAlgorithmQueryParameter adds the passwordHashAlgorithm to the account collection get params
func (o *AccountCollectionGetParams) WithPasswordHashAlgorithmQueryParameter(passwordHashAlgorithm *string) *AccountCollectionGetParams {
	o.SetPasswordHashAlgorithmQueryParameter(passwordHashAlgorithm)
	return o
}

// SetPasswordHashAlgorithmQueryParameter adds the passwordHashAlgorithm to the account collection get params
func (o *AccountCollectionGetParams) SetPasswordHashAlgorithmQueryParameter(passwordHashAlgorithm *string) {
	o.PasswordHashAlgorithmQueryParameter = passwordHashAlgorithm
}

// WithPublicKeyQueryParameter adds the publicKey to the account collection get params
func (o *AccountCollectionGetParams) WithPublicKeyQueryParameter(publicKey *string) *AccountCollectionGetParams {
	o.SetPublicKeyQueryParameter(publicKey)
	return o
}

// SetPublicKeyQueryParameter adds the publicKey to the account collection get params
func (o *AccountCollectionGetParams) SetPublicKeyQueryParameter(publicKey *string) {
	o.PublicKeyQueryParameter = publicKey
}

// WithReturnRecordsQueryParameter adds the returnRecords to the account collection get params
func (o *AccountCollectionGetParams) WithReturnRecordsQueryParameter(returnRecords *bool) *AccountCollectionGetParams {
	o.SetReturnRecordsQueryParameter(returnRecords)
	return o
}

// SetReturnRecordsQueryParameter adds the returnRecords to the account collection get params
func (o *AccountCollectionGetParams) SetReturnRecordsQueryParameter(returnRecords *bool) {
	o.ReturnRecordsQueryParameter = returnRecords
}

// WithReturnTimeoutQueryParameter adds the returnTimeout to the account collection get params
func (o *AccountCollectionGetParams) WithReturnTimeoutQueryParameter(returnTimeout *int64) *AccountCollectionGetParams {
	o.SetReturnTimeoutQueryParameter(returnTimeout)
	return o
}

// SetReturnTimeoutQueryParameter adds the returnTimeout to the account collection get params
func (o *AccountCollectionGetParams) SetReturnTimeoutQueryParameter(returnTimeout *int64) {
	o.ReturnTimeoutQueryParameter = returnTimeout
}

// WithRoleNameQueryParameter adds the roleName to the account collection get params
func (o *AccountCollectionGetParams) WithRoleNameQueryParameter(roleName *string) *AccountCollectionGetParams {
	o.SetRoleNameQueryParameter(roleName)
	return o
}

// SetRoleNameQueryParameter adds the roleName to the account collection get params
func (o *AccountCollectionGetParams) SetRoleNameQueryParameter(roleName *string) {
	o.RoleNameQueryParameter = roleName
}

// WithScopeQueryParameter adds the scope to the account collection get params
func (o *AccountCollectionGetParams) WithScopeQueryParameter(scope *string) *AccountCollectionGetParams {
	o.SetScopeQueryParameter(scope)
	return o
}

// SetScopeQueryParameter adds the scope to the account collection get params
func (o *AccountCollectionGetParams) SetScopeQueryParameter(scope *string) {
	o.ScopeQueryParameter = scope
}

// WithSslCaCertificateQueryParameter adds the sslCaCertificate to the account collection get params
func (o *AccountCollectionGetParams) WithSslCaCertificateQueryParameter(sslCaCertificate *string) *AccountCollectionGetParams {
	o.SetSslCaCertificateQueryParameter(sslCaCertificate)
	return o
}

// SetSslCaCertificateQueryParameter adds the sslCaCertificate to the account collection get params
func (o *AccountCollectionGetParams) SetSslCaCertificateQueryParameter(sslCaCertificate *string) {
	o.SslCaCertificateQueryParameter = sslCaCertificate
}

// WriteToRequest writes these params to a swagger request
func (o *AccountCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ApplicationsApplicationQueryParameter != nil {

		// query param applications.application
		var qrApplicationsApplication string

		if o.ApplicationsApplicationQueryParameter != nil {
			qrApplicationsApplication = *o.ApplicationsApplicationQueryParameter
		}
		qApplicationsApplication := qrApplicationsApplication
		if qApplicationsApplication != "" {

			if err := r.SetQueryParam("applications.application", qApplicationsApplication); err != nil {
				return err
			}
		}
	}

	if o.ApplicationsAuthenticationMethodsQueryParameter != nil {

		// query param applications.authentication_methods
		var qrApplicationsAuthenticationMethods string

		if o.ApplicationsAuthenticationMethodsQueryParameter != nil {
			qrApplicationsAuthenticationMethods = *o.ApplicationsAuthenticationMethodsQueryParameter
		}
		qApplicationsAuthenticationMethods := qrApplicationsAuthenticationMethods
		if qApplicationsAuthenticationMethods != "" {

			if err := r.SetQueryParam("applications.authentication_methods", qApplicationsAuthenticationMethods); err != nil {
				return err
			}
		}
	}

	if o.ApplicationsSecondAuthenticationMethodQueryParameter != nil {

		// query param applications.second_authentication_method
		var qrApplicationsSecondAuthenticationMethod string

		if o.ApplicationsSecondAuthenticationMethodQueryParameter != nil {
			qrApplicationsSecondAuthenticationMethod = *o.ApplicationsSecondAuthenticationMethodQueryParameter
		}
		qApplicationsSecondAuthenticationMethod := qrApplicationsSecondAuthenticationMethod
		if qApplicationsSecondAuthenticationMethod != "" {

			if err := r.SetQueryParam("applications.second_authentication_method", qApplicationsSecondAuthenticationMethod); err != nil {
				return err
			}
		}
	}

	if o.AuthenticationMethodsQueryParameter != nil {

		// query param authentication_methods
		var qrAuthenticationMethods string

		if o.AuthenticationMethodsQueryParameter != nil {
			qrAuthenticationMethods = *o.AuthenticationMethodsQueryParameter
		}
		qAuthenticationMethods := qrAuthenticationMethods
		if qAuthenticationMethods != "" {

			if err := r.SetQueryParam("authentication_methods", qAuthenticationMethods); err != nil {
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

	if o.LdapFastbindQueryParameter != nil {

		// query param ldap_fastbind
		var qrLdapFastbind bool

		if o.LdapFastbindQueryParameter != nil {
			qrLdapFastbind = *o.LdapFastbindQueryParameter
		}
		qLdapFastbind := swag.FormatBool(qrLdapFastbind)
		if qLdapFastbind != "" {

			if err := r.SetQueryParam("ldap_fastbind", qLdapFastbind); err != nil {
				return err
			}
		}
	}

	if o.LockedQueryParameter != nil {

		// query param locked
		var qrLocked bool

		if o.LockedQueryParameter != nil {
			qrLocked = *o.LockedQueryParameter
		}
		qLocked := swag.FormatBool(qrLocked)
		if qLocked != "" {

			if err := r.SetQueryParam("locked", qLocked); err != nil {
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

	if o.PasswordHashAlgorithmQueryParameter != nil {

		// query param password_hash_algorithm
		var qrPasswordHashAlgorithm string

		if o.PasswordHashAlgorithmQueryParameter != nil {
			qrPasswordHashAlgorithm = *o.PasswordHashAlgorithmQueryParameter
		}
		qPasswordHashAlgorithm := qrPasswordHashAlgorithm
		if qPasswordHashAlgorithm != "" {

			if err := r.SetQueryParam("password_hash_algorithm", qPasswordHashAlgorithm); err != nil {
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

	if o.RoleNameQueryParameter != nil {

		// query param role.name
		var qrRoleName string

		if o.RoleNameQueryParameter != nil {
			qrRoleName = *o.RoleNameQueryParameter
		}
		qRoleName := qrRoleName
		if qRoleName != "" {

			if err := r.SetQueryParam("role.name", qRoleName); err != nil {
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

	if o.SslCaCertificateQueryParameter != nil {

		// query param ssl_ca_certificate
		var qrSslCaCertificate string

		if o.SslCaCertificateQueryParameter != nil {
			qrSslCaCertificate = *o.SslCaCertificateQueryParameter
		}
		qSslCaCertificate := qrSslCaCertificate
		if qSslCaCertificate != "" {

			if err := r.SetQueryParam("ssl_ca_certificate", qSslCaCertificate); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamAccountCollectionGet binds the parameter fields
func (o *AccountCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamAccountCollectionGet binds the parameter order_by
func (o *AccountCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
