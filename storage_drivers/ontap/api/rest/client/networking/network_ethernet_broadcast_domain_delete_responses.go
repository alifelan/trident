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

// NetworkEthernetBroadcastDomainDeleteReader is a Reader for the NetworkEthernetBroadcastDomainDelete structure.
type NetworkEthernetBroadcastDomainDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NetworkEthernetBroadcastDomainDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNetworkEthernetBroadcastDomainDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNetworkEthernetBroadcastDomainDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNetworkEthernetBroadcastDomainDeleteOK creates a NetworkEthernetBroadcastDomainDeleteOK with default headers values
func NewNetworkEthernetBroadcastDomainDeleteOK() *NetworkEthernetBroadcastDomainDeleteOK {
	return &NetworkEthernetBroadcastDomainDeleteOK{}
}

/*
NetworkEthernetBroadcastDomainDeleteOK describes a response with status code 200, with default header values.

OK
*/
type NetworkEthernetBroadcastDomainDeleteOK struct {
}

// IsSuccess returns true when this network ethernet broadcast domain delete o k response has a 2xx status code
func (o *NetworkEthernetBroadcastDomainDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this network ethernet broadcast domain delete o k response has a 3xx status code
func (o *NetworkEthernetBroadcastDomainDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this network ethernet broadcast domain delete o k response has a 4xx status code
func (o *NetworkEthernetBroadcastDomainDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this network ethernet broadcast domain delete o k response has a 5xx status code
func (o *NetworkEthernetBroadcastDomainDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this network ethernet broadcast domain delete o k response a status code equal to that given
func (o *NetworkEthernetBroadcastDomainDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *NetworkEthernetBroadcastDomainDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /network/ethernet/broadcast-domains/{uuid}][%d] networkEthernetBroadcastDomainDeleteOK ", 200)
}

func (o *NetworkEthernetBroadcastDomainDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /network/ethernet/broadcast-domains/{uuid}][%d] networkEthernetBroadcastDomainDeleteOK ", 200)
}

func (o *NetworkEthernetBroadcastDomainDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewNetworkEthernetBroadcastDomainDeleteDefault creates a NetworkEthernetBroadcastDomainDeleteDefault with default headers values
func NewNetworkEthernetBroadcastDomainDeleteDefault(code int) *NetworkEthernetBroadcastDomainDeleteDefault {
	return &NetworkEthernetBroadcastDomainDeleteDefault{
		_statusCode: code,
	}
}

/*
	NetworkEthernetBroadcastDomainDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1967103 | A broadcast domain with ports cannot be deleted. |
*/
type NetworkEthernetBroadcastDomainDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the network ethernet broadcast domain delete default response
func (o *NetworkEthernetBroadcastDomainDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this network ethernet broadcast domain delete default response has a 2xx status code
func (o *NetworkEthernetBroadcastDomainDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this network ethernet broadcast domain delete default response has a 3xx status code
func (o *NetworkEthernetBroadcastDomainDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this network ethernet broadcast domain delete default response has a 4xx status code
func (o *NetworkEthernetBroadcastDomainDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this network ethernet broadcast domain delete default response has a 5xx status code
func (o *NetworkEthernetBroadcastDomainDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this network ethernet broadcast domain delete default response a status code equal to that given
func (o *NetworkEthernetBroadcastDomainDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NetworkEthernetBroadcastDomainDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /network/ethernet/broadcast-domains/{uuid}][%d] network_ethernet_broadcast_domain_delete default  %+v", o._statusCode, o.Payload)
}

func (o *NetworkEthernetBroadcastDomainDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /network/ethernet/broadcast-domains/{uuid}][%d] network_ethernet_broadcast_domain_delete default  %+v", o._statusCode, o.Payload)
}

func (o *NetworkEthernetBroadcastDomainDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NetworkEthernetBroadcastDomainDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
