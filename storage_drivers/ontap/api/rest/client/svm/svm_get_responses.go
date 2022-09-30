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

// SvmGetReader is a Reader for the SvmGet structure.
type SvmGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SvmGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSvmGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSvmGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSvmGetOK creates a SvmGetOK with default headers values
func NewSvmGetOK() *SvmGetOK {
	return &SvmGetOK{}
}

/*
SvmGetOK describes a response with status code 200, with default header values.

OK
*/
type SvmGetOK struct {
	Payload *models.Svm
}

// IsSuccess returns true when this svm get o k response has a 2xx status code
func (o *SvmGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this svm get o k response has a 3xx status code
func (o *SvmGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this svm get o k response has a 4xx status code
func (o *SvmGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this svm get o k response has a 5xx status code
func (o *SvmGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this svm get o k response a status code equal to that given
func (o *SvmGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *SvmGetOK) Error() string {
	return fmt.Sprintf("[GET /svm/svms/{uuid}][%d] svmGetOK  %+v", 200, o.Payload)
}

func (o *SvmGetOK) String() string {
	return fmt.Sprintf("[GET /svm/svms/{uuid}][%d] svmGetOK  %+v", 200, o.Payload)
}

func (o *SvmGetOK) GetPayload() *models.Svm {
	return o.Payload
}

func (o *SvmGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Svm)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSvmGetDefault creates a SvmGetDefault with default headers values
func NewSvmGetDefault(code int) *SvmGetDefault {
	return &SvmGetDefault{
		_statusCode: code,
	}
}

/*
SvmGetDefault describes a response with status code -1, with default header values.

Error
*/
type SvmGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the svm get default response
func (o *SvmGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this svm get default response has a 2xx status code
func (o *SvmGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this svm get default response has a 3xx status code
func (o *SvmGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this svm get default response has a 4xx status code
func (o *SvmGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this svm get default response has a 5xx status code
func (o *SvmGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this svm get default response a status code equal to that given
func (o *SvmGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SvmGetDefault) Error() string {
	return fmt.Sprintf("[GET /svm/svms/{uuid}][%d] svm_get default  %+v", o._statusCode, o.Payload)
}

func (o *SvmGetDefault) String() string {
	return fmt.Sprintf("[GET /svm/svms/{uuid}][%d] svm_get default  %+v", o._statusCode, o.Payload)
}

func (o *SvmGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SvmGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
