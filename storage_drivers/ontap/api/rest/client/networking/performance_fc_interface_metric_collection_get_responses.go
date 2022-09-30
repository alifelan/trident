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

// PerformanceFcInterfaceMetricCollectionGetReader is a Reader for the PerformanceFcInterfaceMetricCollectionGet structure.
type PerformanceFcInterfaceMetricCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PerformanceFcInterfaceMetricCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPerformanceFcInterfaceMetricCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewPerformanceFcInterfaceMetricCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPerformanceFcInterfaceMetricCollectionGetOK creates a PerformanceFcInterfaceMetricCollectionGetOK with default headers values
func NewPerformanceFcInterfaceMetricCollectionGetOK() *PerformanceFcInterfaceMetricCollectionGetOK {
	return &PerformanceFcInterfaceMetricCollectionGetOK{}
}

/*
PerformanceFcInterfaceMetricCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type PerformanceFcInterfaceMetricCollectionGetOK struct {
	Payload *models.PerformanceFcInterfaceMetricResponse
}

// IsSuccess returns true when this performance fc interface metric collection get o k response has a 2xx status code
func (o *PerformanceFcInterfaceMetricCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this performance fc interface metric collection get o k response has a 3xx status code
func (o *PerformanceFcInterfaceMetricCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this performance fc interface metric collection get o k response has a 4xx status code
func (o *PerformanceFcInterfaceMetricCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this performance fc interface metric collection get o k response has a 5xx status code
func (o *PerformanceFcInterfaceMetricCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this performance fc interface metric collection get o k response a status code equal to that given
func (o *PerformanceFcInterfaceMetricCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *PerformanceFcInterfaceMetricCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /network/fc/interfaces/{uuid}/metrics][%d] performanceFcInterfaceMetricCollectionGetOK  %+v", 200, o.Payload)
}

func (o *PerformanceFcInterfaceMetricCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /network/fc/interfaces/{uuid}/metrics][%d] performanceFcInterfaceMetricCollectionGetOK  %+v", 200, o.Payload)
}

func (o *PerformanceFcInterfaceMetricCollectionGetOK) GetPayload() *models.PerformanceFcInterfaceMetricResponse {
	return o.Payload
}

func (o *PerformanceFcInterfaceMetricCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceFcInterfaceMetricResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPerformanceFcInterfaceMetricCollectionGetDefault creates a PerformanceFcInterfaceMetricCollectionGetDefault with default headers values
func NewPerformanceFcInterfaceMetricCollectionGetDefault(code int) *PerformanceFcInterfaceMetricCollectionGetDefault {
	return &PerformanceFcInterfaceMetricCollectionGetDefault{
		_statusCode: code,
	}
}

/*
PerformanceFcInterfaceMetricCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type PerformanceFcInterfaceMetricCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the performance fc interface metric collection get default response
func (o *PerformanceFcInterfaceMetricCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this performance fc interface metric collection get default response has a 2xx status code
func (o *PerformanceFcInterfaceMetricCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this performance fc interface metric collection get default response has a 3xx status code
func (o *PerformanceFcInterfaceMetricCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this performance fc interface metric collection get default response has a 4xx status code
func (o *PerformanceFcInterfaceMetricCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this performance fc interface metric collection get default response has a 5xx status code
func (o *PerformanceFcInterfaceMetricCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this performance fc interface metric collection get default response a status code equal to that given
func (o *PerformanceFcInterfaceMetricCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *PerformanceFcInterfaceMetricCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /network/fc/interfaces/{uuid}/metrics][%d] performance_fc_interface_metric_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *PerformanceFcInterfaceMetricCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /network/fc/interfaces/{uuid}/metrics][%d] performance_fc_interface_metric_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *PerformanceFcInterfaceMetricCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *PerformanceFcInterfaceMetricCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
