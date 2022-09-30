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

// NvmeSubsystemMapCreateReader is a Reader for the NvmeSubsystemMapCreate structure.
type NvmeSubsystemMapCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *NvmeSubsystemMapCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewNvmeSubsystemMapCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewNvmeSubsystemMapCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewNvmeSubsystemMapCreateCreated creates a NvmeSubsystemMapCreateCreated with default headers values
func NewNvmeSubsystemMapCreateCreated() *NvmeSubsystemMapCreateCreated {
	return &NvmeSubsystemMapCreateCreated{}
}

/*
NvmeSubsystemMapCreateCreated describes a response with status code 201, with default header values.

Created
*/
type NvmeSubsystemMapCreateCreated struct {
	Payload *models.NvmeSubsystemMapResponse
}

// IsSuccess returns true when this nvme subsystem map create created response has a 2xx status code
func (o *NvmeSubsystemMapCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this nvme subsystem map create created response has a 3xx status code
func (o *NvmeSubsystemMapCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this nvme subsystem map create created response has a 4xx status code
func (o *NvmeSubsystemMapCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this nvme subsystem map create created response has a 5xx status code
func (o *NvmeSubsystemMapCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this nvme subsystem map create created response a status code equal to that given
func (o *NvmeSubsystemMapCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *NvmeSubsystemMapCreateCreated) Error() string {
	return fmt.Sprintf("[POST /protocols/nvme/subsystem-maps][%d] nvmeSubsystemMapCreateCreated  %+v", 201, o.Payload)
}

func (o *NvmeSubsystemMapCreateCreated) String() string {
	return fmt.Sprintf("[POST /protocols/nvme/subsystem-maps][%d] nvmeSubsystemMapCreateCreated  %+v", 201, o.Payload)
}

func (o *NvmeSubsystemMapCreateCreated) GetPayload() *models.NvmeSubsystemMapResponse {
	return o.Payload
}

func (o *NvmeSubsystemMapCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.NvmeSubsystemMapResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewNvmeSubsystemMapCreateDefault creates a NvmeSubsystemMapCreateDefault with default headers values
func NewNvmeSubsystemMapCreateDefault(code int) *NvmeSubsystemMapCreateDefault {
	return &NvmeSubsystemMapCreateDefault{
		_statusCode: code,
	}
}

/*
	NvmeSubsystemMapCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 72089790 | The supplied NVMe namespace is already mapped to the supplied NVMe subsystem. |
| 72089793 | An NVMe namespace in a Snapshot copy cannot be mapped. |
| 72089799 | The NVMe namespace is the destination of an ongoing restore operation and is inaccessible for I/O and management. |
| 72089902 | A node does not have an NVMe interface configured. |
| 72089903 | Multiple nodes do not have an NVMe interface configured. |
| 72089904 | The aggregate must be given back to its home node prior to mapping the NVMe namespace it contains. |
| 72090001 | The NVMe subsystem specified by `subsystem.uuid` was not found. |
| 72090005 | The specified `namespace.uuid` and `namespace.name` refer to different NVMe namespaces. |
| 72090006 | The NVMe namespace specified by `namespace.uuid` was not found. |
| 72090007 | The NVMe namespace specified by `namespace.name` was not found. |
| 72090020 | The specified `subsystem.uuid` and `subsystem.name` refer to different NVMe subsystems. |
| 72090021 | The NVMe subsystem specified by `subsystem.name` was not found. |
*/
type NvmeSubsystemMapCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the nvme subsystem map create default response
func (o *NvmeSubsystemMapCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this nvme subsystem map create default response has a 2xx status code
func (o *NvmeSubsystemMapCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this nvme subsystem map create default response has a 3xx status code
func (o *NvmeSubsystemMapCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this nvme subsystem map create default response has a 4xx status code
func (o *NvmeSubsystemMapCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this nvme subsystem map create default response has a 5xx status code
func (o *NvmeSubsystemMapCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this nvme subsystem map create default response a status code equal to that given
func (o *NvmeSubsystemMapCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *NvmeSubsystemMapCreateDefault) Error() string {
	return fmt.Sprintf("[POST /protocols/nvme/subsystem-maps][%d] nvme_subsystem_map_create default  %+v", o._statusCode, o.Payload)
}

func (o *NvmeSubsystemMapCreateDefault) String() string {
	return fmt.Sprintf("[POST /protocols/nvme/subsystem-maps][%d] nvme_subsystem_map_create default  %+v", o._statusCode, o.Payload)
}

func (o *NvmeSubsystemMapCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *NvmeSubsystemMapCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
