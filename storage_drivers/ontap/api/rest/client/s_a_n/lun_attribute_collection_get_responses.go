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

// LunAttributeCollectionGetReader is a Reader for the LunAttributeCollectionGet structure.
type LunAttributeCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunAttributeCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLunAttributeCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunAttributeCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunAttributeCollectionGetOK creates a LunAttributeCollectionGetOK with default headers values
func NewLunAttributeCollectionGetOK() *LunAttributeCollectionGetOK {
	return &LunAttributeCollectionGetOK{}
}

/*
LunAttributeCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type LunAttributeCollectionGetOK struct {
	Payload *models.LunAttributeResponse
}

// IsSuccess returns true when this lun attribute collection get o k response has a 2xx status code
func (o *LunAttributeCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lun attribute collection get o k response has a 3xx status code
func (o *LunAttributeCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lun attribute collection get o k response has a 4xx status code
func (o *LunAttributeCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lun attribute collection get o k response has a 5xx status code
func (o *LunAttributeCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lun attribute collection get o k response a status code equal to that given
func (o *LunAttributeCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *LunAttributeCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/luns/{lun.uuid}/attributes][%d] lunAttributeCollectionGetOK  %+v", 200, o.Payload)
}

func (o *LunAttributeCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /storage/luns/{lun.uuid}/attributes][%d] lunAttributeCollectionGetOK  %+v", 200, o.Payload)
}

func (o *LunAttributeCollectionGetOK) GetPayload() *models.LunAttributeResponse {
	return o.Payload
}

func (o *LunAttributeCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LunAttributeResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLunAttributeCollectionGetDefault creates a LunAttributeCollectionGetDefault with default headers values
func NewLunAttributeCollectionGetDefault(code int) *LunAttributeCollectionGetDefault {
	return &LunAttributeCollectionGetDefault{
		_statusCode: code,
	}
}

/*
LunAttributeCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type LunAttributeCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the lun attribute collection get default response
func (o *LunAttributeCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this lun attribute collection get default response has a 2xx status code
func (o *LunAttributeCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lun attribute collection get default response has a 3xx status code
func (o *LunAttributeCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lun attribute collection get default response has a 4xx status code
func (o *LunAttributeCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lun attribute collection get default response has a 5xx status code
func (o *LunAttributeCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lun attribute collection get default response a status code equal to that given
func (o *LunAttributeCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *LunAttributeCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/luns/{lun.uuid}/attributes][%d] lun_attribute_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *LunAttributeCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /storage/luns/{lun.uuid}/attributes][%d] lun_attribute_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *LunAttributeCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunAttributeCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
