// Code generated by go-swagger; DO NOT EDIT.

package svm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SvmMigrateDeleteReader is a Reader for the SvmMigrateDelete structure.
type SvmMigrateDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmMigrateDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewSvmMigrateDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmMigrateDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmMigrateDeleteAccepted creates a SvmMigrateDeleteAccepted with default headers values
func NewSvmMigrateDeleteAccepted() *SvmMigrateDeleteAccepted {
	return &SvmMigrateDeleteAccepted{}
}

/*
SvmMigrateDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type SvmMigrateDeleteAccepted struct {
}

// IsSuccess returns true when this svm migrate delete accepted response has a 2xx status code
func (o *SvmMigrateDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm migrate delete accepted response has a 3xx status code
func (o *SvmMigrateDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm migrate delete accepted response has a 4xx status code
func (o *SvmMigrateDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm migrate delete accepted response has a 5xx status code
func (o *SvmMigrateDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this svm migrate delete accepted response a status code equal to that given
func (o *SvmMigrateDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *SvmMigrateDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /svm/migrations/{uuid}][%d] svmMigrateDeleteAccepted ", 202)
}

func (o *SvmMigrateDeleteAccepted) String() string {
	return fmt.Sprintf("[DELETE /svm/migrations/{uuid}][%d] svmMigrateDeleteAccepted ", 202)
}

func (o *SvmMigrateDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSvmMigrateDeleteDefault creates a SvmMigrateDeleteDefault with default headers values
func NewSvmMigrateDeleteDefault(code int) *SvmMigrateDeleteDefault {
	return &SvmMigrateDeleteDefault{
		_statusCode: code,
	}
}

/*
	SvmMigrateDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 13172783 | Migrate RDB lookup failed |
*/
type SvmMigrateDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the svm migrate delete default response
func (o *SvmMigrateDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this svm migrate delete default response has a 2xx status code
func (o *SvmMigrateDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm migrate delete default response has a 3xx status code
func (o *SvmMigrateDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm migrate delete default response has a 4xx status code
func (o *SvmMigrateDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm migrate delete default response has a 5xx status code
func (o *SvmMigrateDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm migrate delete default response a status code equal to that given
func (o *SvmMigrateDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SvmMigrateDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /svm/migrations/{uuid}][%d] svm_migrate_delete default  %+v", o._statusCode, o.Payload)
}

func (o *SvmMigrateDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /svm/migrations/{uuid}][%d] svm_migrate_delete default  %+v", o._statusCode, o.Payload)
}

func (o *SvmMigrateDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmMigrateDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
