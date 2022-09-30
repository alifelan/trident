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

// DNSGetReader is a Reader for the DNSGet structure.
type DNSGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DNSGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDNSGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewDNSGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDNSGetOK creates a DNSGetOK with default headers values
func NewDNSGetOK() *DNSGetOK {
	return &DNSGetOK{}
}

/*
DNSGetOK describes a response with status code 200, with default header values.

OK
*/
type DNSGetOK struct {
	Payload *models.DNS
}

// IsSuccess returns true when this dns get o k response has a 2xx status code
func (o *DNSGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this dns get o k response has a 3xx status code
func (o *DNSGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this dns get o k response has a 4xx status code
func (o *DNSGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this dns get o k response has a 5xx status code
func (o *DNSGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this dns get o k response a status code equal to that given
func (o *DNSGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *DNSGetOK) Error() string {
	return fmt.Sprintf("[GET /name-services/dns/{svm.uuid}][%d] dnsGetOK  %+v", 200, o.Payload)
}

func (o *DNSGetOK) String() string {
	return fmt.Sprintf("[GET /name-services/dns/{svm.uuid}][%d] dnsGetOK  %+v", 200, o.Payload)
}

func (o *DNSGetOK) GetPayload() *models.DNS {
	return o.Payload
}

func (o *DNSGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DNS)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDNSGetDefault creates a DNSGetDefault with default headers values
func NewDNSGetDefault(code int) *DNSGetDefault {
	return &DNSGetDefault{
		_statusCode: code,
	}
}

/*
DNSGetDefault describes a response with status code -1, with default header values.

Error
*/
type DNSGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the dns get default response
func (o *DNSGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this dns get default response has a 2xx status code
func (o *DNSGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this dns get default response has a 3xx status code
func (o *DNSGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this dns get default response has a 4xx status code
func (o *DNSGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this dns get default response has a 5xx status code
func (o *DNSGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this dns get default response a status code equal to that given
func (o *DNSGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *DNSGetDefault) Error() string {
	return fmt.Sprintf("[GET /name-services/dns/{svm.uuid}][%d] dns_get default  %+v", o._statusCode, o.Payload)
}

func (o *DNSGetDefault) String() string {
	return fmt.Sprintf("[GET /name-services/dns/{svm.uuid}][%d] dns_get default  %+v", o._statusCode, o.Payload)
}

func (o *DNSGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *DNSGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
