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

// SnaplockLegalHoldCreateReader is a Reader for the SnaplockLegalHoldCreate structure.
type SnaplockLegalHoldCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockLegalHoldCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSnaplockLegalHoldCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockLegalHoldCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockLegalHoldCreateCreated creates a SnaplockLegalHoldCreateCreated with default headers values
func NewSnaplockLegalHoldCreateCreated() *SnaplockLegalHoldCreateCreated {
	return &SnaplockLegalHoldCreateCreated{}
}

/*
SnaplockLegalHoldCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SnaplockLegalHoldCreateCreated struct {
	Payload *models.SnaplockLegalHoldOperation
}

// IsSuccess returns true when this snaplock legal hold create created response has a 2xx status code
func (o *SnaplockLegalHoldCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock legal hold create created response has a 3xx status code
func (o *SnaplockLegalHoldCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock legal hold create created response has a 4xx status code
func (o *SnaplockLegalHoldCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock legal hold create created response has a 5xx status code
func (o *SnaplockLegalHoldCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock legal hold create created response a status code equal to that given
func (o *SnaplockLegalHoldCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *SnaplockLegalHoldCreateCreated) Error() string {
	return fmt.Sprintf("[POST /storage/snaplock/litigations/{litigation.id}/operations][%d] snaplockLegalHoldCreateCreated  %+v", 201, o.Payload)
}

func (o *SnaplockLegalHoldCreateCreated) String() string {
	return fmt.Sprintf("[POST /storage/snaplock/litigations/{litigation.id}/operations][%d] snaplockLegalHoldCreateCreated  %+v", 201, o.Payload)
}

func (o *SnaplockLegalHoldCreateCreated) GetPayload() *models.SnaplockLegalHoldOperation {
	return o.Payload
}

func (o *SnaplockLegalHoldCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockLegalHoldOperation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockLegalHoldCreateDefault creates a SnaplockLegalHoldCreateDefault with default headers values
func NewSnaplockLegalHoldCreateDefault(code int) *SnaplockLegalHoldCreateDefault {
	return &SnaplockLegalHoldCreateDefault{
		_statusCode: code,
	}
}

/*
	SnaplockLegalHoldCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response codes

| Error code  |  Description |
|-------------|--------------|
| 14090346    | Internal Error. Wait a few minutes, then try the command again  |
| 14090343    | Invalid Field  |
*/
type SnaplockLegalHoldCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock legal hold create default response
func (o *SnaplockLegalHoldCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this snaplock legal hold create default response has a 2xx status code
func (o *SnaplockLegalHoldCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock legal hold create default response has a 3xx status code
func (o *SnaplockLegalHoldCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock legal hold create default response has a 4xx status code
func (o *SnaplockLegalHoldCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock legal hold create default response has a 5xx status code
func (o *SnaplockLegalHoldCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock legal hold create default response a status code equal to that given
func (o *SnaplockLegalHoldCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SnaplockLegalHoldCreateDefault) Error() string {
	return fmt.Sprintf("[POST /storage/snaplock/litigations/{litigation.id}/operations][%d] snaplock_legal_hold_create default  %+v", o._statusCode, o.Payload)
}

func (o *SnaplockLegalHoldCreateDefault) String() string {
	return fmt.Sprintf("[POST /storage/snaplock/litigations/{litigation.id}/operations][%d] snaplock_legal_hold_create default  %+v", o._statusCode, o.Payload)
}

func (o *SnaplockLegalHoldCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockLegalHoldCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
