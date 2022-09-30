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

// LunMapReportingNodeGetReader is a Reader for the LunMapReportingNodeGet structure.
type LunMapReportingNodeGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *LunMapReportingNodeGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewLunMapReportingNodeGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewLunMapReportingNodeGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewLunMapReportingNodeGetOK creates a LunMapReportingNodeGetOK with default headers values
func NewLunMapReportingNodeGetOK() *LunMapReportingNodeGetOK {
	return &LunMapReportingNodeGetOK{}
}

/*
LunMapReportingNodeGetOK describes a response with status code 200, with default header values.

OK
*/
type LunMapReportingNodeGetOK struct {
	Payload *models.LunMapReportingNode
}

// IsSuccess returns true when this lun map reporting node get o k response has a 2xx status code
func (o *LunMapReportingNodeGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this lun map reporting node get o k response has a 3xx status code
func (o *LunMapReportingNodeGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this lun map reporting node get o k response has a 4xx status code
func (o *LunMapReportingNodeGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this lun map reporting node get o k response has a 5xx status code
func (o *LunMapReportingNodeGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this lun map reporting node get o k response a status code equal to that given
func (o *LunMapReportingNodeGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *LunMapReportingNodeGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}/reporting-nodes/{uuid}][%d] lunMapReportingNodeGetOK  %+v", 200, o.Payload)
}

func (o *LunMapReportingNodeGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}/reporting-nodes/{uuid}][%d] lunMapReportingNodeGetOK  %+v", 200, o.Payload)
}

func (o *LunMapReportingNodeGetOK) GetPayload() *models.LunMapReportingNode {
	return o.Payload
}

func (o *LunMapReportingNodeGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.LunMapReportingNode)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewLunMapReportingNodeGetDefault creates a LunMapReportingNodeGetDefault with default headers values
func NewLunMapReportingNodeGetDefault(code int) *LunMapReportingNodeGetDefault {
	return &LunMapReportingNodeGetDefault{
		_statusCode: code,
	}
}

/*
	LunMapReportingNodeGetDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 5374875 | The specified LUN does not exist or is not accessible to the caller. |
| 5374878 | The specified initiator group does not exist, is not accessible to the caller, or is not in the same SVM as the specified LUN. |
| 5374922 | The specified LUN map does not exist. |
*/
type LunMapReportingNodeGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the lun map reporting node get default response
func (o *LunMapReportingNodeGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this lun map reporting node get default response has a 2xx status code
func (o *LunMapReportingNodeGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this lun map reporting node get default response has a 3xx status code
func (o *LunMapReportingNodeGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this lun map reporting node get default response has a 4xx status code
func (o *LunMapReportingNodeGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this lun map reporting node get default response has a 5xx status code
func (o *LunMapReportingNodeGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this lun map reporting node get default response a status code equal to that given
func (o *LunMapReportingNodeGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *LunMapReportingNodeGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}/reporting-nodes/{uuid}][%d] lun_map_reporting_node_get default  %+v", o._statusCode, o.Payload)
}

func (o *LunMapReportingNodeGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/san/lun-maps/{lun.uuid}/{igroup.uuid}/reporting-nodes/{uuid}][%d] lun_map_reporting_node_get default  %+v", o._statusCode, o.Payload)
}

func (o *LunMapReportingNodeGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *LunMapReportingNodeGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
