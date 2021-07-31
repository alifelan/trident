// Code generated by go-swagger; DO NOT EDIT.

package support

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnmpGetReader is a Reader for the SnmpGet structure.
type SnmpGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnmpGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnmpGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnmpGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnmpGetOK creates a SnmpGetOK with default headers values
func NewSnmpGetOK() *SnmpGetOK {
	return &SnmpGetOK{}
}

/* SnmpGetOK describes a response with status code 200, with default header values.

OK
*/
type SnmpGetOK struct {
	Payload *models.Snmp
}

func (o *SnmpGetOK) Error() string {
	return fmt.Sprintf("[GET /support/snmp][%d] snmpGetOK  %+v", 200, o.Payload)
}
func (o *SnmpGetOK) GetPayload() *models.Snmp {
	return o.Payload
}

func (o *SnmpGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Snmp)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnmpGetDefault creates a SnmpGetDefault with default headers values
func NewSnmpGetDefault(code int) *SnmpGetDefault {
	return &SnmpGetDefault{
		_statusCode: code,
	}
}

/* SnmpGetDefault describes a response with status code -1, with default header values.

Error
*/
type SnmpGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snmp get default response
func (o *SnmpGetDefault) Code() int {
	return o._statusCode
}

func (o *SnmpGetDefault) Error() string {
	return fmt.Sprintf("[GET /support/snmp][%d] snmp_get default  %+v", o._statusCode, o.Payload)
}
func (o *SnmpGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnmpGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}