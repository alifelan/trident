// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// FabricCollectionGetReader is a Reader for the FabricCollectionGet structure.
type FabricCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FabricCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFabricCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFabricCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFabricCollectionGetOK creates a FabricCollectionGetOK with default headers values
func NewFabricCollectionGetOK() *FabricCollectionGetOK {
	return &FabricCollectionGetOK{}
}

/*
FabricCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type FabricCollectionGetOK struct {
	Payload *models.FabricResponse
}

// IsSuccess returns true when this fabric collection get o k response has a 2xx status code
func (o *FabricCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fabric collection get o k response has a 3xx status code
func (o *FabricCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fabric collection get o k response has a 4xx status code
func (o *FabricCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fabric collection get o k response has a 5xx status code
func (o *FabricCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fabric collection get o k response a status code equal to that given
func (o *FabricCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *FabricCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /network/fc/fabrics][%d] fabricCollectionGetOK  %+v", 200, o.Payload)
}

func (o *FabricCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /network/fc/fabrics][%d] fabricCollectionGetOK  %+v", 200, o.Payload)
}

func (o *FabricCollectionGetOK) GetPayload() *models.FabricResponse {
	return o.Payload
}

func (o *FabricCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FabricResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFabricCollectionGetDefault creates a FabricCollectionGetDefault with default headers values
func NewFabricCollectionGetDefault(code int) *FabricCollectionGetDefault {
	return &FabricCollectionGetDefault{
		_statusCode: code,
	}
}

/*
FabricCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type FabricCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fabric collection get default response
func (o *FabricCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this fabric collection get default response has a 2xx status code
func (o *FabricCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fabric collection get default response has a 3xx status code
func (o *FabricCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fabric collection get default response has a 4xx status code
func (o *FabricCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fabric collection get default response has a 5xx status code
func (o *FabricCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fabric collection get default response a status code equal to that given
func (o *FabricCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FabricCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /network/fc/fabrics][%d] fabric_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *FabricCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /network/fc/fabrics][%d] fabric_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *FabricCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FabricCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
