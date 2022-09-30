// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// QosPolicyDeleteReader is a Reader for the QosPolicyDelete structure.
type QosPolicyDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *QosPolicyDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewQosPolicyDeleteAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewQosPolicyDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewQosPolicyDeleteAccepted creates a QosPolicyDeleteAccepted with default headers values
func NewQosPolicyDeleteAccepted() *QosPolicyDeleteAccepted {
	return &QosPolicyDeleteAccepted{}
}

/*
QosPolicyDeleteAccepted describes a response with status code 202, with default header values.

Accepted
*/
type QosPolicyDeleteAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this qos policy delete accepted response has a 2xx status code
func (o *QosPolicyDeleteAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this qos policy delete accepted response has a 3xx status code
func (o *QosPolicyDeleteAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this qos policy delete accepted response has a 4xx status code
func (o *QosPolicyDeleteAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this qos policy delete accepted response has a 5xx status code
func (o *QosPolicyDeleteAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this qos policy delete accepted response a status code equal to that given
func (o *QosPolicyDeleteAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *QosPolicyDeleteAccepted) Error() string {
	return fmt.Sprintf("[DELETE /storage/qos/policies/{uuid}][%d] qosPolicyDeleteAccepted  %+v", 202, o.Payload)
}

func (o *QosPolicyDeleteAccepted) String() string {
	return fmt.Sprintf("[DELETE /storage/qos/policies/{uuid}][%d] qosPolicyDeleteAccepted  %+v", 202, o.Payload)
}

func (o *QosPolicyDeleteAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *QosPolicyDeleteAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewQosPolicyDeleteDefault creates a QosPolicyDeleteDefault with default headers values
func NewQosPolicyDeleteDefault(code int) *QosPolicyDeleteDefault {
	return &QosPolicyDeleteDefault{
		_statusCode: code,
	}
}

/*
QosPolicyDeleteDefault describes a response with status code -1, with default header values.

Error
*/
type QosPolicyDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the qos policy delete default response
func (o *QosPolicyDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this qos policy delete default response has a 2xx status code
func (o *QosPolicyDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this qos policy delete default response has a 3xx status code
func (o *QosPolicyDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this qos policy delete default response has a 4xx status code
func (o *QosPolicyDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this qos policy delete default response has a 5xx status code
func (o *QosPolicyDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this qos policy delete default response a status code equal to that given
func (o *QosPolicyDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *QosPolicyDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/qos/policies/{uuid}][%d] qos_policy_delete default  %+v", o._statusCode, o.Payload)
}

func (o *QosPolicyDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /storage/qos/policies/{uuid}][%d] qos_policy_delete default  %+v", o._statusCode, o.Payload)
}

func (o *QosPolicyDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *QosPolicyDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
