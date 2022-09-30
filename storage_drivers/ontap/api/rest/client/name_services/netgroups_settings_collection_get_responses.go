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

// NetgroupsSettingsCollectionGetReader is a Reader for the NetgroupsSettingsCollectionGet structure.
type NetgroupsSettingsCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetgroupsSettingsCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetgroupsSettingsCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetgroupsSettingsCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetgroupsSettingsCollectionGetOK creates a NetgroupsSettingsCollectionGetOK with default headers values
func NewNetgroupsSettingsCollectionGetOK() *NetgroupsSettingsCollectionGetOK {
	return &NetgroupsSettingsCollectionGetOK{}
}

/*
NetgroupsSettingsCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type NetgroupsSettingsCollectionGetOK struct {
	Payload *models.NetgroupsSettingsResponse
}

// IsSuccess returns true when this netgroups settings collection get o k response has a 2xx status code
func (o *NetgroupsSettingsCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this netgroups settings collection get o k response has a 3xx status code
func (o *NetgroupsSettingsCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this netgroups settings collection get o k response has a 4xx status code
func (o *NetgroupsSettingsCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this netgroups settings collection get o k response has a 5xx status code
func (o *NetgroupsSettingsCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this netgroups settings collection get o k response a status code equal to that given
func (o *NetgroupsSettingsCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *NetgroupsSettingsCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /name-services/cache/netgroup/settings][%d] netgroupsSettingsCollectionGetOK  %+v", 200, o.Payload)
}

func (o *NetgroupsSettingsCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /name-services/cache/netgroup/settings][%d] netgroupsSettingsCollectionGetOK  %+v", 200, o.Payload)
}

func (o *NetgroupsSettingsCollectionGetOK) GetPayload() *models.NetgroupsSettingsResponse {
	return o.Payload
}

func (o *NetgroupsSettingsCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetgroupsSettingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetgroupsSettingsCollectionGetDefault creates a NetgroupsSettingsCollectionGetDefault with default headers values
func NewNetgroupsSettingsCollectionGetDefault(code int) *NetgroupsSettingsCollectionGetDefault {
	return &NetgroupsSettingsCollectionGetDefault{
		_statusCode: code,
	}
}

/*
NetgroupsSettingsCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetgroupsSettingsCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the netgroups settings collection get default response
func (o *NetgroupsSettingsCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this netgroups settings collection get default response has a 2xx status code
func (o *NetgroupsSettingsCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this netgroups settings collection get default response has a 3xx status code
func (o *NetgroupsSettingsCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this netgroups settings collection get default response has a 4xx status code
func (o *NetgroupsSettingsCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this netgroups settings collection get default response has a 5xx status code
func (o *NetgroupsSettingsCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this netgroups settings collection get default response a status code equal to that given
func (o *NetgroupsSettingsCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NetgroupsSettingsCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /name-services/cache/netgroup/settings][%d] netgroups_settings_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *NetgroupsSettingsCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /name-services/cache/netgroup/settings][%d] netgroups_settings_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *NetgroupsSettingsCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetgroupsSettingsCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
