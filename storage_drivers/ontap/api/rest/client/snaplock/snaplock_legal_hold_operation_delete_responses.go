// Code generated by go-swagger; DO NOT EDIT.

package snaplock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnaplockLegalHoldOperationDeleteReader is a Reader for the SnaplockLegalHoldOperationDelete structure.
type SnaplockLegalHoldOperationDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockLegalHoldOperationDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockLegalHoldOperationDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockLegalHoldOperationDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockLegalHoldOperationDeleteOK creates a SnaplockLegalHoldOperationDeleteOK with default headers values
func NewSnaplockLegalHoldOperationDeleteOK() *SnaplockLegalHoldOperationDeleteOK {
	return &SnaplockLegalHoldOperationDeleteOK{}
}

/*
SnaplockLegalHoldOperationDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockLegalHoldOperationDeleteOK struct {
}

// IsSuccess returns true when this snaplock legal hold operation delete o k response has a 2xx status code
func (o *SnaplockLegalHoldOperationDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock legal hold operation delete o k response has a 3xx status code
func (o *SnaplockLegalHoldOperationDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock legal hold operation delete o k response has a 4xx status code
func (o *SnaplockLegalHoldOperationDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock legal hold operation delete o k response has a 5xx status code
func (o *SnaplockLegalHoldOperationDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock legal hold operation delete o k response a status code equal to that given
func (o *SnaplockLegalHoldOperationDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *SnaplockLegalHoldOperationDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /storage/snaplock/litigations/{litigation.id}/operations/{id}][%d] snaplockLegalHoldOperationDeleteOK ", 200)
}

func (o *SnaplockLegalHoldOperationDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /storage/snaplock/litigations/{litigation.id}/operations/{id}][%d] snaplockLegalHoldOperationDeleteOK ", 200)
}

func (o *SnaplockLegalHoldOperationDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSnaplockLegalHoldOperationDeleteDefault creates a SnaplockLegalHoldOperationDeleteDefault with default headers values
func NewSnaplockLegalHoldOperationDeleteDefault(code int) *SnaplockLegalHoldOperationDeleteDefault {
	return &SnaplockLegalHoldOperationDeleteDefault{
		_statusCode: code,
	}
}

/*
	SnaplockLegalHoldOperationDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 14090346    | Internal Error. Wait a few minutes, then try the command again  |
| 14090541    | A completed or failed operation cannot be aborted  |
*/
type SnaplockLegalHoldOperationDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock legal hold operation delete default response
func (o *SnaplockLegalHoldOperationDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this snaplock legal hold operation delete default response has a 2xx status code
func (o *SnaplockLegalHoldOperationDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock legal hold operation delete default response has a 3xx status code
func (o *SnaplockLegalHoldOperationDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock legal hold operation delete default response has a 4xx status code
func (o *SnaplockLegalHoldOperationDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock legal hold operation delete default response has a 5xx status code
func (o *SnaplockLegalHoldOperationDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock legal hold operation delete default response a status code equal to that given
func (o *SnaplockLegalHoldOperationDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SnaplockLegalHoldOperationDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/snaplock/litigations/{litigation.id}/operations/{id}][%d] snaplock_legal_hold_operation_delete default  %+v", o._statusCode, o.Payload)
}

func (o *SnaplockLegalHoldOperationDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /storage/snaplock/litigations/{litigation.id}/operations/{id}][%d] snaplock_legal_hold_operation_delete default  %+v", o._statusCode, o.Payload)
}

func (o *SnaplockLegalHoldOperationDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockLegalHoldOperationDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
