// Code generated by go-swagger; DO NOT EDIT.

package n_v_me

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// NvmeCollectionPerformanceMetricsGetReader is a Reader for the NvmeCollectionPerformanceMetricsGet structure.
type NvmeCollectionPerformanceMetricsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeCollectionPerformanceMetricsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewNvmeCollectionPerformanceMetricsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeCollectionPerformanceMetricsGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeCollectionPerformanceMetricsGetOK creates a NvmeCollectionPerformanceMetricsGetOK with default headers values
func NewNvmeCollectionPerformanceMetricsGetOK() *NvmeCollectionPerformanceMetricsGetOK {
	return &NvmeCollectionPerformanceMetricsGetOK{}
}

/*
NvmeCollectionPerformanceMetricsGetOK describes a response with status code 200, with default header values.

OK
*/
type NvmeCollectionPerformanceMetricsGetOK struct {
	Payload *models.PerformanceNvmeMetricResponse
}

// IsSuccess returns true when this nvme collection performance metrics get o k response has a 2xx status code
func (o *NvmeCollectionPerformanceMetricsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme collection performance metrics get o k response has a 3xx status code
func (o *NvmeCollectionPerformanceMetricsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme collection performance metrics get o k response has a 4xx status code
func (o *NvmeCollectionPerformanceMetricsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme collection performance metrics get o k response has a 5xx status code
func (o *NvmeCollectionPerformanceMetricsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme collection performance metrics get o k response a status code equal to that given
func (o *NvmeCollectionPerformanceMetricsGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *NvmeCollectionPerformanceMetricsGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/nvme/services/{svm.uuid}/metrics][%d] nvmeCollectionPerformanceMetricsGetOK  %+v", 200, o.Payload)
}

func (o *NvmeCollectionPerformanceMetricsGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/nvme/services/{svm.uuid}/metrics][%d] nvmeCollectionPerformanceMetricsGetOK  %+v", 200, o.Payload)
}

func (o *NvmeCollectionPerformanceMetricsGetOK) GetPayload() *models.PerformanceNvmeMetricResponse {
	return o.Payload
}

func (o *NvmeCollectionPerformanceMetricsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PerformanceNvmeMetricResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeCollectionPerformanceMetricsGetDefault creates a NvmeCollectionPerformanceMetricsGetDefault with default headers values
func NewNvmeCollectionPerformanceMetricsGetDefault(code int) *NvmeCollectionPerformanceMetricsGetDefault {
	return &NvmeCollectionPerformanceMetricsGetDefault{
		_statusCode: code,
	}
}

/*
NvmeCollectionPerformanceMetricsGetDefault describes a response with status code -1, with default header values.

Error
*/
type NvmeCollectionPerformanceMetricsGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the nvme collection performance metrics get default response
func (o *NvmeCollectionPerformanceMetricsGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this nvme collection performance metrics get default response has a 2xx status code
func (o *NvmeCollectionPerformanceMetricsGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme collection performance metrics get default response has a 3xx status code
func (o *NvmeCollectionPerformanceMetricsGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme collection performance metrics get default response has a 4xx status code
func (o *NvmeCollectionPerformanceMetricsGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme collection performance metrics get default response has a 5xx status code
func (o *NvmeCollectionPerformanceMetricsGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme collection performance metrics get default response a status code equal to that given
func (o *NvmeCollectionPerformanceMetricsGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NvmeCollectionPerformanceMetricsGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/nvme/services/{svm.uuid}/metrics][%d] nvme_collection_performance_metrics_get default  %+v", o._statusCode, o.Payload)
}

func (o *NvmeCollectionPerformanceMetricsGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/nvme/services/{svm.uuid}/metrics][%d] nvme_collection_performance_metrics_get default  %+v", o._statusCode, o.Payload)
}

func (o *NvmeCollectionPerformanceMetricsGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeCollectionPerformanceMetricsGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
