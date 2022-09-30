// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// FcpServiceGetReader is a Reader for the FcpServiceGet structure.
type FcpServiceGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FcpServiceGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFcpServiceGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFcpServiceGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFcpServiceGetOK creates a FcpServiceGetOK with default headers values
func NewFcpServiceGetOK() *FcpServiceGetOK {
	return &FcpServiceGetOK{}
}

/*
FcpServiceGetOK describes a response with status code 200, with default header values.

OK
*/
type FcpServiceGetOK struct {
	Payload *models.FcpService
}

// IsSuccess returns true when this fcp service get o k response has a 2xx status code
func (o *FcpServiceGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fcp service get o k response has a 3xx status code
func (o *FcpServiceGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fcp service get o k response has a 4xx status code
func (o *FcpServiceGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fcp service get o k response has a 5xx status code
func (o *FcpServiceGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fcp service get o k response a status code equal to that given
func (o *FcpServiceGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *FcpServiceGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/san/fcp/services/{svm.uuid}][%d] fcpServiceGetOK  %+v", 200, o.Payload)
}

func (o *FcpServiceGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/san/fcp/services/{svm.uuid}][%d] fcpServiceGetOK  %+v", 200, o.Payload)
}

func (o *FcpServiceGetOK) GetPayload() *models.FcpService {
	return o.Payload
}

func (o *FcpServiceGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FcpService)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFcpServiceGetDefault creates a FcpServiceGetDefault with default headers values
func NewFcpServiceGetDefault(code int) *FcpServiceGetDefault {
	return &FcpServiceGetDefault{
		_statusCode: code,
	}
}

/*
	FcpServiceGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2621462 | An SVM with the specified UUID does not exist. |
| 5374083 | There is no Fibre Channel Protocol service for the specified SVM. |
*/
type FcpServiceGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fcp service get default response
func (o *FcpServiceGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this fcp service get default response has a 2xx status code
func (o *FcpServiceGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fcp service get default response has a 3xx status code
func (o *FcpServiceGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fcp service get default response has a 4xx status code
func (o *FcpServiceGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fcp service get default response has a 5xx status code
func (o *FcpServiceGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fcp service get default response a status code equal to that given
func (o *FcpServiceGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FcpServiceGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/san/fcp/services/{svm.uuid}][%d] fcp_service_get default  %+v", o._statusCode, o.Payload)
}

func (o *FcpServiceGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/san/fcp/services/{svm.uuid}][%d] fcp_service_get default  %+v", o._statusCode, o.Payload)
}

func (o *FcpServiceGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FcpServiceGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
