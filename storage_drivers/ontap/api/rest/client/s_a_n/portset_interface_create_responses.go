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

// PortsetInterfaceCreateReader is a Reader for the PortsetInterfaceCreate structure.
type PortsetInterfaceCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PortsetInterfaceCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewPortsetInterfaceCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPortsetInterfaceCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPortsetInterfaceCreateCreated creates a PortsetInterfaceCreateCreated with default headers values
func NewPortsetInterfaceCreateCreated() *PortsetInterfaceCreateCreated {
	return &PortsetInterfaceCreateCreated{}
}

/*
PortsetInterfaceCreateCreated describes a response with status code 201, with default header values.

Created
*/
type PortsetInterfaceCreateCreated struct {
	Payload *models.PortsetInterfaceResponse
}

// IsSuccess returns true when this portset interface create created response has a 2xx status code
func (o *PortsetInterfaceCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this portset interface create created response has a 3xx status code
func (o *PortsetInterfaceCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this portset interface create created response has a 4xx status code
func (o *PortsetInterfaceCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this portset interface create created response has a 5xx status code
func (o *PortsetInterfaceCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this portset interface create created response a status code equal to that given
func (o *PortsetInterfaceCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *PortsetInterfaceCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/san/portsets/{portset.uuid}/interfaces][%d] portsetInterfaceCreateCreated  %+v", 201, o.Payload)
}

func (o *PortsetInterfaceCreateCreated) String() string {
	return fmt.Sprintf("[POST /protocols/san/portsets/{portset.uuid}/interfaces][%d] portsetInterfaceCreateCreated  %+v", 201, o.Payload)
}

func (o *PortsetInterfaceCreateCreated) GetPayload() *models.PortsetInterfaceResponse {
	return o.Payload
}

func (o *PortsetInterfaceCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortsetInterfaceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPortsetInterfaceCreateDefault creates a PortsetInterfaceCreateDefault with default headers values
func NewPortsetInterfaceCreateDefault(code int) *PortsetInterfaceCreateDefault {
	return &PortsetInterfaceCreateDefault{
		_statusCode: code,
	}
}

/*
	PortsetInterfaceCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5374906 | A specified network interface was not found. |
| 5374907 | The specified network interface UUID and name don't identify the same network interface. |
| 5374909 | An invalid combination of network interface properties was specified. |
| 5374910 | An incomplete set of network interface properties was specified. |
| 5374914 | An attempt was made to add a network interface of an incompatible protocol to a portset. |
| 5374915 | An attempt was made to add a duplicate network interface to a portset. |
*/
type PortsetInterfaceCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the portset interface create default response
func (o *PortsetInterfaceCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this portset interface create default response has a 2xx status code
func (o *PortsetInterfaceCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this portset interface create default response has a 3xx status code
func (o *PortsetInterfaceCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this portset interface create default response has a 4xx status code
func (o *PortsetInterfaceCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this portset interface create default response has a 5xx status code
func (o *PortsetInterfaceCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this portset interface create default response a status code equal to that given
func (o *PortsetInterfaceCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PortsetInterfaceCreateDefault) Error() string {
	return fmt.Sprintf("[POST /protocols/san/portsets/{portset.uuid}/interfaces][%d] portset_interface_create default  %+v", o._statusCode, o.Payload)
}

func (o *PortsetInterfaceCreateDefault) String() string {
	return fmt.Sprintf("[POST /protocols/san/portsets/{portset.uuid}/interfaces][%d] portset_interface_create default  %+v", o._statusCode, o.Payload)
}

func (o *PortsetInterfaceCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PortsetInterfaceCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
