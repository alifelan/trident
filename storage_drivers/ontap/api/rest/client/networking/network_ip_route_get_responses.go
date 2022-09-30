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

// NetworkIPRouteGetReader is a Reader for the NetworkIPRouteGet structure.
type NetworkIPRouteGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkIPRouteGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkIPRouteGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkIPRouteGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkIPRouteGetOK creates a NetworkIPRouteGetOK with default headers values
func NewNetworkIPRouteGetOK() *NetworkIPRouteGetOK {
	return &NetworkIPRouteGetOK{}
}

/*
NetworkIPRouteGetOK describes a response with status code 200, with default header values.

OK
*/
type NetworkIPRouteGetOK struct {
	Payload *models.NetworkRoute
}

// IsSuccess returns true when this network Ip route get o k response has a 2xx status code
func (o *NetworkIPRouteGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network Ip route get o k response has a 3xx status code
func (o *NetworkIPRouteGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network Ip route get o k response has a 4xx status code
func (o *NetworkIPRouteGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network Ip route get o k response has a 5xx status code
func (o *NetworkIPRouteGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network Ip route get o k response a status code equal to that given
func (o *NetworkIPRouteGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *NetworkIPRouteGetOK) Error() string {
	return fmt.Sprintf("[GET /network/ip/routes/{uuid}][%d] networkIpRouteGetOK  %+v", 200, o.Payload)
}

func (o *NetworkIPRouteGetOK) String() string {
	return fmt.Sprintf("[GET /network/ip/routes/{uuid}][%d] networkIpRouteGetOK  %+v", 200, o.Payload)
}

func (o *NetworkIPRouteGetOK) GetPayload() *models.NetworkRoute {
	return o.Payload
}

func (o *NetworkIPRouteGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NetworkRoute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNetworkIPRouteGetDefault creates a NetworkIPRouteGetDefault with default headers values
func NewNetworkIPRouteGetDefault(code int) *NetworkIPRouteGetDefault {
	return &NetworkIPRouteGetDefault{
		_statusCode: code,
	}
}

/*
NetworkIPRouteGetDefault describes a response with status code -1, with default header values.

Error
*/
type NetworkIPRouteGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the network ip route get default response
func (o *NetworkIPRouteGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this network ip route get default response has a 2xx status code
func (o *NetworkIPRouteGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ip route get default response has a 3xx status code
func (o *NetworkIPRouteGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ip route get default response has a 4xx status code
func (o *NetworkIPRouteGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ip route get default response has a 5xx status code
func (o *NetworkIPRouteGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ip route get default response a status code equal to that given
func (o *NetworkIPRouteGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NetworkIPRouteGetDefault) Error() string {
	return fmt.Sprintf("[GET /network/ip/routes/{uuid}][%d] network_ip_route_get default  %+v", o._statusCode, o.Payload)
}

func (o *NetworkIPRouteGetDefault) String() string {
	return fmt.Sprintf("[GET /network/ip/routes/{uuid}][%d] network_ip_route_get default  %+v", o._statusCode, o.Payload)
}

func (o *NetworkIPRouteGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkIPRouteGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
