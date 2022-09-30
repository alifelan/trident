// Code generated by go-swagger; DO NOT EDIT.

package name_services

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// UnixGroupCreateReader is a Reader for the UnixGroupCreate structure.
type UnixGroupCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UnixGroupCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewUnixGroupCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUnixGroupCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUnixGroupCreateCreated creates a UnixGroupCreateCreated with default headers values
func NewUnixGroupCreateCreated() *UnixGroupCreateCreated {
	return &UnixGroupCreateCreated{}
}

/*
UnixGroupCreateCreated describes a response with status code 201, with default header values.

Created
*/
type UnixGroupCreateCreated struct {
}

// IsSuccess returns true when this unix group create created response has a 2xx status code
func (o *UnixGroupCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this unix group create created response has a 3xx status code
func (o *UnixGroupCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this unix group create created response has a 4xx status code
func (o *UnixGroupCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this unix group create created response has a 5xx status code
func (o *UnixGroupCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this unix group create created response a status code equal to that given
func (o *UnixGroupCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *UnixGroupCreateCreated) Error() string {
	return fmt.Sprintf("[POST /name-services/unix-groups][%d] unixGroupCreateCreated ", 201)
}

func (o *UnixGroupCreateCreated) String() string {
	return fmt.Sprintf("[POST /name-services/unix-groups][%d] unixGroupCreateCreated ", 201)
}

func (o *UnixGroupCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUnixGroupCreateDefault creates a UnixGroupCreateDefault with default headers values
func NewUnixGroupCreateDefault(code int) *UnixGroupCreateDefault {
	return &UnixGroupCreateDefault{
		_statusCode: code,
	}
}

/*
	UnixGroupCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621706    | The specified SVM UUID is incorrect for the specified SVM name. |
| 3277025    | Maximum supported limit of UNIX group count reached. |
| 3277051    | Invalid characters in group name. Valid characters are 0-9, A-Z, a-z, ".", "_" and "-". Names cannot start with "-". |
| 23724067   | Group name too long. Maximum supported length is 64 characters. |
| 23724141   | Duplicate group ID. Group ID must be unique.|
*/
type UnixGroupCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the unix group create default response
func (o *UnixGroupCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this unix group create default response has a 2xx status code
func (o *UnixGroupCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this unix group create default response has a 3xx status code
func (o *UnixGroupCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this unix group create default response has a 4xx status code
func (o *UnixGroupCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this unix group create default response has a 5xx status code
func (o *UnixGroupCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this unix group create default response a status code equal to that given
func (o *UnixGroupCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *UnixGroupCreateDefault) Error() string {
	return fmt.Sprintf("[POST /name-services/unix-groups][%d] unix_group_create default  %+v", o._statusCode, o.Payload)
}

func (o *UnixGroupCreateDefault) String() string {
	return fmt.Sprintf("[POST /name-services/unix-groups][%d] unix_group_create default  %+v", o._statusCode, o.Payload)
}

func (o *UnixGroupCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *UnixGroupCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
