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

// SnaplockRetentionOperationCreateReader is a Reader for the SnaplockRetentionOperationCreate structure.
type SnaplockRetentionOperationCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockRetentionOperationCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewSnaplockRetentionOperationCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockRetentionOperationCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockRetentionOperationCreateCreated creates a SnaplockRetentionOperationCreateCreated with default headers values
func NewSnaplockRetentionOperationCreateCreated() *SnaplockRetentionOperationCreateCreated {
	return &SnaplockRetentionOperationCreateCreated{}
}

/*
SnaplockRetentionOperationCreateCreated describes a response with status code 201, with default header values.

Created
*/
type SnaplockRetentionOperationCreateCreated struct {
	Payload *models.EbrOperation
}

// IsSuccess returns true when this snaplock retention operation create created response has a 2xx status code
func (o *SnaplockRetentionOperationCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snaplock retention operation create created response has a 3xx status code
func (o *SnaplockRetentionOperationCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snaplock retention operation create created response has a 4xx status code
func (o *SnaplockRetentionOperationCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this snaplock retention operation create created response has a 5xx status code
func (o *SnaplockRetentionOperationCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this snaplock retention operation create created response a status code equal to that given
func (o *SnaplockRetentionOperationCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *SnaplockRetentionOperationCreateCreated) Error() string {
	return fmt.Sprintf("[POST /storage/snaplock/event-retention/operations][%d] snaplockRetentionOperationCreateCreated  %+v", 201, o.Payload)
}

func (o *SnaplockRetentionOperationCreateCreated) String() string {
	return fmt.Sprintf("[POST /storage/snaplock/event-retention/operations][%d] snaplockRetentionOperationCreateCreated  %+v", 201, o.Payload)
}

func (o *SnaplockRetentionOperationCreateCreated) GetPayload() *models.EbrOperation {
	return o.Payload
}

func (o *SnaplockRetentionOperationCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EbrOperation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockRetentionOperationCreateDefault creates a SnaplockRetentionOperationCreateDefault with default headers values
func NewSnaplockRetentionOperationCreateDefault(code int) *SnaplockRetentionOperationCreateDefault {
	return &SnaplockRetentionOperationCreateDefault{
		_statusCode: code,
	}
}

/*
SnaplockRetentionOperationCreateDefault describes a response with status code -1, with default header values.

Error
*/
type SnaplockRetentionOperationCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock retention operation create default response
func (o *SnaplockRetentionOperationCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this snaplock retention operation create default response has a 2xx status code
func (o *SnaplockRetentionOperationCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snaplock retention operation create default response has a 3xx status code
func (o *SnaplockRetentionOperationCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snaplock retention operation create default response has a 4xx status code
func (o *SnaplockRetentionOperationCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snaplock retention operation create default response has a 5xx status code
func (o *SnaplockRetentionOperationCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snaplock retention operation create default response a status code equal to that given
func (o *SnaplockRetentionOperationCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SnaplockRetentionOperationCreateDefault) Error() string {
	return fmt.Sprintf("[POST /storage/snaplock/event-retention/operations][%d] snaplock_retention_operation_create default  %+v", o._statusCode, o.Payload)
}

func (o *SnaplockRetentionOperationCreateDefault) String() string {
	return fmt.Sprintf("[POST /storage/snaplock/event-retention/operations][%d] snaplock_retention_operation_create default  %+v", o._statusCode, o.Payload)
}

func (o *SnaplockRetentionOperationCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockRetentionOperationCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
