// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// FileCopyCreateReader is a Reader for the FileCopyCreate structure.
type FileCopyCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FileCopyCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewFileCopyCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFileCopyCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFileCopyCreateAccepted creates a FileCopyCreateAccepted with default headers values
func NewFileCopyCreateAccepted() *FileCopyCreateAccepted {
	return &FileCopyCreateAccepted{}
}

/*
FileCopyCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type FileCopyCreateAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this file copy create accepted response has a 2xx status code
func (o *FileCopyCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this file copy create accepted response has a 3xx status code
func (o *FileCopyCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this file copy create accepted response has a 4xx status code
func (o *FileCopyCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this file copy create accepted response has a 5xx status code
func (o *FileCopyCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this file copy create accepted response a status code equal to that given
func (o *FileCopyCreateAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *FileCopyCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /storage/file/copy][%d] fileCopyCreateAccepted  %+v", 202, o.Payload)
}

func (o *FileCopyCreateAccepted) String() string {
	return fmt.Sprintf("[POST /storage/file/copy][%d] fileCopyCreateAccepted  %+v", 202, o.Payload)
}

func (o *FileCopyCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *FileCopyCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewFileCopyCreateDefault creates a FileCopyCreateDefault with default headers values
func NewFileCopyCreateDefault(code int) *FileCopyCreateDefault {
	return &FileCopyCreateDefault{
		_statusCode: code,
	}
}

/*
	FileCopyCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 7012352 | File locations are inconsistent. All files must be on the same volume. |
| 7012353 | Exceeded the file operations supported number of files. |
| 7012354 | Unable to pair the number of source files to destination files. |
| 7012357 | Cannot start a file operation until all cluster nodes support the file operations capability. |
| 7012358 | The specified source path is invalid. |
| 7012359 | The specified destination path is invalid. |
| 7012360 | The SVMs are not in an intracluster peering relationship. |
| 7012361 | The SVMs peering relationship does not include application \"file-copy\". |
| 7012362 | The SVMs are not yet in a peered state yet. |
| 7012363 | Cannot copy files. All file operations must be managed by the destination SVM's administrator. |
| 7012365 | Copying a file between clusters is not supported. |
| 7012367 | A reference path may only be specified if multiple source paths are specified. |
| 7012368 | The reference path must have a matching source path. |
| 7012371 | The reference cutover time exceeds the maximum allowable time. |
| 7012374 | Source volume and destination volume have different home clusters. |
| 7012376 | Operation not allowed on a volume that is part of a SnapMirror Synchronous relationship. |
| 7012377 | Cannot start a file copy operation on the volume because an active volume conversion is in progress. |
| 13107223 | Operation not supported for FlexGroup volumes or FlexGroup constituents. |
| 196608143 | Cannot start operation. The volume is undergoing a secure purge operation. |
*/
type FileCopyCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the file copy create default response
func (o *FileCopyCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this file copy create default response has a 2xx status code
func (o *FileCopyCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this file copy create default response has a 3xx status code
func (o *FileCopyCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this file copy create default response has a 4xx status code
func (o *FileCopyCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this file copy create default response has a 5xx status code
func (o *FileCopyCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this file copy create default response a status code equal to that given
func (o *FileCopyCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FileCopyCreateDefault) Error() string {
	return fmt.Sprintf("[POST /storage/file/copy][%d] file_copy_create default  %+v", o._statusCode, o.Payload)
}

func (o *FileCopyCreateDefault) String() string {
	return fmt.Sprintf("[POST /storage/file/copy][%d] file_copy_create default  %+v", o._statusCode, o.Payload)
}

func (o *FileCopyCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FileCopyCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
