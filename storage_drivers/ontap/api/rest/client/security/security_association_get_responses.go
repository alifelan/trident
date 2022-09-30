// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SecurityAssociationGetReader is a Reader for the SecurityAssociationGet structure.
type SecurityAssociationGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SecurityAssociationGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSecurityAssociationGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSecurityAssociationGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSecurityAssociationGetOK creates a SecurityAssociationGetOK with default headers values
func NewSecurityAssociationGetOK() *SecurityAssociationGetOK {
	return &SecurityAssociationGetOK{}
}

/*
SecurityAssociationGetOK describes a response with status code 200, with default header values.

OK
*/
type SecurityAssociationGetOK struct {
	Payload *models.SecurityAssociation
}

// IsSuccess returns true when this security association get o k response has a 2xx status code
func (o *SecurityAssociationGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this security association get o k response has a 3xx status code
func (o *SecurityAssociationGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this security association get o k response has a 4xx status code
func (o *SecurityAssociationGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this security association get o k response has a 5xx status code
func (o *SecurityAssociationGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this security association get o k response a status code equal to that given
func (o *SecurityAssociationGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *SecurityAssociationGetOK) Error() string {
	return fmt.Sprintf("[GET /security/ipsec/security-associations/{uuid}][%d] securityAssociationGetOK  %+v", 200, o.Payload)
}

func (o *SecurityAssociationGetOK) String() string {
	return fmt.Sprintf("[GET /security/ipsec/security-associations/{uuid}][%d] securityAssociationGetOK  %+v", 200, o.Payload)
}

func (o *SecurityAssociationGetOK) GetPayload() *models.SecurityAssociation {
	return o.Payload
}

func (o *SecurityAssociationGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SecurityAssociation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSecurityAssociationGetDefault creates a SecurityAssociationGetDefault with default headers values
func NewSecurityAssociationGetDefault(code int) *SecurityAssociationGetDefault {
	return &SecurityAssociationGetDefault{
		_statusCode: code,
	}
}

/*
	SecurityAssociationGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 66257118 | IPsec SA with the specified UUID was not found. |
| 66257119 | IPsec SA with the specified UUID was not found. |
*/
type SecurityAssociationGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the security association get default response
func (o *SecurityAssociationGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this security association get default response has a 2xx status code
func (o *SecurityAssociationGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this security association get default response has a 3xx status code
func (o *SecurityAssociationGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this security association get default response has a 4xx status code
func (o *SecurityAssociationGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this security association get default response has a 5xx status code
func (o *SecurityAssociationGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this security association get default response a status code equal to that given
func (o *SecurityAssociationGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SecurityAssociationGetDefault) Error() string {
	return fmt.Sprintf("[GET /security/ipsec/security-associations/{uuid}][%d] security_association_get default  %+v", o._statusCode, o.Payload)
}

func (o *SecurityAssociationGetDefault) String() string {
	return fmt.Sprintf("[GET /security/ipsec/security-associations/{uuid}][%d] security_association_get default  %+v", o._statusCode, o.Payload)
}

func (o *SecurityAssociationGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SecurityAssociationGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
