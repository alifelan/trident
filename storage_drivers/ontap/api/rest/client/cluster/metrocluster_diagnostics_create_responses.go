// Code generated by go-swagger; DO NOT EDIT.

package cluster

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// MetroclusterDiagnosticsCreateReader is a Reader for the MetroclusterDiagnosticsCreate structure.
type MetroclusterDiagnosticsCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MetroclusterDiagnosticsCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewMetroclusterDiagnosticsCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewMetroclusterDiagnosticsCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewMetroclusterDiagnosticsCreateAccepted creates a MetroclusterDiagnosticsCreateAccepted with default headers values
func NewMetroclusterDiagnosticsCreateAccepted() *MetroclusterDiagnosticsCreateAccepted {
	return &MetroclusterDiagnosticsCreateAccepted{}
}

/*
MetroclusterDiagnosticsCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type MetroclusterDiagnosticsCreateAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this metrocluster diagnostics create accepted response has a 2xx status code
func (o *MetroclusterDiagnosticsCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this metrocluster diagnostics create accepted response has a 3xx status code
func (o *MetroclusterDiagnosticsCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this metrocluster diagnostics create accepted response has a 4xx status code
func (o *MetroclusterDiagnosticsCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this metrocluster diagnostics create accepted response has a 5xx status code
func (o *MetroclusterDiagnosticsCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this metrocluster diagnostics create accepted response a status code equal to that given
func (o *MetroclusterDiagnosticsCreateAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *MetroclusterDiagnosticsCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /cluster/metrocluster/diagnostics][%d] metroclusterDiagnosticsCreateAccepted  %+v", 202, o.Payload)
}

func (o *MetroclusterDiagnosticsCreateAccepted) String() string {
	return fmt.Sprintf("[POST /cluster/metrocluster/diagnostics][%d] metroclusterDiagnosticsCreateAccepted  %+v", 202, o.Payload)
}

func (o *MetroclusterDiagnosticsCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *MetroclusterDiagnosticsCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMetroclusterDiagnosticsCreateDefault creates a MetroclusterDiagnosticsCreateDefault with default headers values
func NewMetroclusterDiagnosticsCreateDefault(code int) *MetroclusterDiagnosticsCreateDefault {
	return &MetroclusterDiagnosticsCreateDefault{
		_statusCode: code,
	}
}

/*
	MetroclusterDiagnosticsCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 2425734 | An internal error occurred. Wait a few minutes, and try the operation again. For further assistance, contact technical support. |
| 2427132 | MetroCluster is not configured on this cluster. |
| 2432833 | Operation is already running. |
| 2432852 | MetroCluster diagnostics start |
| 2432853 | MetroCluster diagnostics job scheduled |
| 2432854 | MetroCluster diagnostics complete |
| 2432855 | MetroCluster diagnostics operation failed. Use the REST API GET method on "/api/cluster/metrocluster/operations?type=check&fields=*" for more information. |
*/
type MetroclusterDiagnosticsCreateDefault struct {
	_statusCode int
}

// Code gets the status code for the metrocluster diagnostics create default response
func (o *MetroclusterDiagnosticsCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this metrocluster diagnostics create default response has a 2xx status code
func (o *MetroclusterDiagnosticsCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this metrocluster diagnostics create default response has a 3xx status code
func (o *MetroclusterDiagnosticsCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this metrocluster diagnostics create default response has a 4xx status code
func (o *MetroclusterDiagnosticsCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this metrocluster diagnostics create default response has a 5xx status code
func (o *MetroclusterDiagnosticsCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this metrocluster diagnostics create default response a status code equal to that given
func (o *MetroclusterDiagnosticsCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *MetroclusterDiagnosticsCreateDefault) Error() string {
	return fmt.Sprintf("[POST /cluster/metrocluster/diagnostics][%d] metrocluster_diagnostics_create default ", o._statusCode)
}

func (o *MetroclusterDiagnosticsCreateDefault) String() string {
	return fmt.Sprintf("[POST /cluster/metrocluster/diagnostics][%d] metrocluster_diagnostics_create default ", o._statusCode)
}

func (o *MetroclusterDiagnosticsCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
