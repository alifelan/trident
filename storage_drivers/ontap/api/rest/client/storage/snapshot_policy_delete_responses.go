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

// SnapshotPolicyDeleteReader is a Reader for the SnapshotPolicyDelete structure.
type SnapshotPolicyDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnapshotPolicyDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnapshotPolicyDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnapshotPolicyDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnapshotPolicyDeleteOK creates a SnapshotPolicyDeleteOK with default headers values
func NewSnapshotPolicyDeleteOK() *SnapshotPolicyDeleteOK {
	return &SnapshotPolicyDeleteOK{}
}

/*
SnapshotPolicyDeleteOK describes a response with status code 200, with default header values.

OK
*/
type SnapshotPolicyDeleteOK struct {
}

// IsSuccess returns true when this snapshot policy delete o k response has a 2xx status code
func (o *SnapshotPolicyDeleteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this snapshot policy delete o k response has a 3xx status code
func (o *SnapshotPolicyDeleteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this snapshot policy delete o k response has a 4xx status code
func (o *SnapshotPolicyDeleteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this snapshot policy delete o k response has a 5xx status code
func (o *SnapshotPolicyDeleteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this snapshot policy delete o k response a status code equal to that given
func (o *SnapshotPolicyDeleteOK) IsCode(code int) bool {
	return code == 200
}

func (o *SnapshotPolicyDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /storage/snapshot-policies/{uuid}][%d] snapshotPolicyDeleteOK ", 200)
}

func (o *SnapshotPolicyDeleteOK) String() string {
	return fmt.Sprintf("[DELETE /storage/snapshot-policies/{uuid}][%d] snapshotPolicyDeleteOK ", 200)
}

func (o *SnapshotPolicyDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewSnapshotPolicyDeleteDefault creates a SnapshotPolicyDeleteDefault with default headers values
func NewSnapshotPolicyDeleteDefault(code int) *SnapshotPolicyDeleteDefault {
	return &SnapshotPolicyDeleteDefault{
		_statusCode: code,
	}
}

/*
	SnapshotPolicyDeleteDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Code

| Error Code | Description |
| ---------- | ----------- |
| 1638415    | Cannot delete policy. Reason: Policy is in use by at least one volume. |
| 1638416    | Cannot delete policy. Reason: Cannot verify whether policy is in use. |
| 1638430    | Cannot delete policy. Reason: Policy is in use by at least one Vserver. |
| 1638430    | Cannot delete built-in policy. |
*/
type SnapshotPolicyDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snapshot policy delete default response
func (o *SnapshotPolicyDeleteDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this snapshot policy delete default response has a 2xx status code
func (o *SnapshotPolicyDeleteDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this snapshot policy delete default response has a 3xx status code
func (o *SnapshotPolicyDeleteDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this snapshot policy delete default response has a 4xx status code
func (o *SnapshotPolicyDeleteDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this snapshot policy delete default response has a 5xx status code
func (o *SnapshotPolicyDeleteDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this snapshot policy delete default response a status code equal to that given
func (o *SnapshotPolicyDeleteDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *SnapshotPolicyDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /storage/snapshot-policies/{uuid}][%d] snapshot_policy_delete default  %+v", o._statusCode, o.Payload)
}

func (o *SnapshotPolicyDeleteDefault) String() string {
	return fmt.Sprintf("[DELETE /storage/snapshot-policies/{uuid}][%d] snapshot_policy_delete default  %+v", o._statusCode, o.Payload)
}

func (o *SnapshotPolicyDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnapshotPolicyDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
