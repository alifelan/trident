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

// VolumeCreateReader is a Reader for the VolumeCreate structure.
type VolumeCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VolumeCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewVolumeCreateAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewVolumeCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewVolumeCreateAccepted creates a VolumeCreateAccepted with default headers values
func NewVolumeCreateAccepted() *VolumeCreateAccepted {
	return &VolumeCreateAccepted{}
}

/*
VolumeCreateAccepted describes a response with status code 202, with default header values.

Accepted
*/
type VolumeCreateAccepted struct {
	Payload *models.JobLinkResponse
}

// IsSuccess returns true when this volume create accepted response has a 2xx status code
func (o *VolumeCreateAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this volume create accepted response has a 3xx status code
func (o *VolumeCreateAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this volume create accepted response has a 4xx status code
func (o *VolumeCreateAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this volume create accepted response has a 5xx status code
func (o *VolumeCreateAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this volume create accepted response a status code equal to that given
func (o *VolumeCreateAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *VolumeCreateAccepted) Error() string {
	return fmt.Sprintf("[POST /storage/volumes][%d] volumeCreateAccepted  %+v", 202, o.Payload)
}

func (o *VolumeCreateAccepted) String() string {
	return fmt.Sprintf("[POST /storage/volumes][%d] volumeCreateAccepted  %+v", 202, o.Payload)
}

func (o *VolumeCreateAccepted) GetPayload() *models.JobLinkResponse {
	return o.Payload
}

func (o *VolumeCreateAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.JobLinkResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewVolumeCreateDefault creates a VolumeCreateDefault with default headers values
func NewVolumeCreateDefault(code int) *VolumeCreateDefault {
	return &VolumeCreateDefault{
		_statusCode: code,
	}
}

/*
	VolumeCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 787140 | One of \"aggregates.uuid\", \"aggregates.name\", or \"style\" must be provided. |
| 787141 | The specified \"aggregates.name\" and \"aggregates.uuid\" refer to different aggregates. |
| 917526 | The volume name specified is a duplicate. |
| 917829 | Volume autosize grow threshold must be larger than autosize shrink threshold. |
| 917831 | Volume minimum autosize must be smaller than the maximum autosize. |
| 917835 | Maximum allowed snapshot.reserve_percent value during a volume creation is 90. Use PATCH to set it to a higher value after the volume has been created. |
| 918191 | Flexvol tiering min cooling days requires an effective cluster version of ONTAP 9.4 or later. |
| 918194 | Tiering min cooling days not supported for SVMDR. |
| 918195 | Tiering min cooling days not supported for non data volumes. |
| 918196 | Tiering min cooling days not allowed for the provided tiering policy. |
| 918215 | FlexGroup tiering min cooling days requires an effective cluster version of ONTAP 9.5 or later. |
| 918233 | The target field cannot be specified for this operation. |
| 918236 | The specified \"parent_volume.uuid\" and \"parent_volume.name\" do not refer to the same volume. |
| 918240 | The target style is an invalid volume style. |
| 918241 | The target style is an unsupported volume style for volume creation. |
| 918242 | When creating a flexible volume, exactly one aggregate must be specified via either \"aggregates.name\" or \"aggregates.uuid\". |
| 918243 | The specified Snapshot copy UUID is not correct for the specified Snapshot copy name. |
| 918244 | Invalid \"volume.type\" for clone volume. |
| 918246 | \"volume.clone.parent_volume.name\" or \"volume.clone.parent_volume.uuid\" must be provided. |
| 918247 | Specifying a value is not valid for a volume FlexClone creation. |
| 918252 | \"nas.path\" is invalid. |
| 918290 | cloud retrieval policy requires an effective cluster version of 9.8 or later. |
| 918291 | Invalid volume cloud retrieval policy for the provided tiering policy. |
| 918292 | cloud retrieval policy not supported for non data volume. |
| 918521 | The volume maximum autosize must be smaller than or equal to the maximum volume size. |
| 918524 | Volume minimum autosize must be less than or equal to the current volume size. |
| 2621706 | The specified \"svm.uuid\" and \"svm.name\" do not refer to the same SVM. |
| 2621707 | No SVM was specified. Either \"svm.name\" or \"svm.uuid\" must be supplied. |
| 111411205 | File system analytics requires an effective cluster version of 9.8 or later. |
| 111411206 | The specified \"analytics.state\" is invalid. |
| 111411207 | File system analytics cannot be enabled on volumes that contain LUNs. |
*/
type VolumeCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the volume create default response
func (o *VolumeCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this volume create default response has a 2xx status code
func (o *VolumeCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this volume create default response has a 3xx status code
func (o *VolumeCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this volume create default response has a 4xx status code
func (o *VolumeCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this volume create default response has a 5xx status code
func (o *VolumeCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this volume create default response a status code equal to that given
func (o *VolumeCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *VolumeCreateDefault) Error() string {
	return fmt.Sprintf("[POST /storage/volumes][%d] volume_create default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeCreateDefault) String() string {
	return fmt.Sprintf("[POST /storage/volumes][%d] volume_create default  %+v", o._statusCode, o.Payload)
}

func (o *VolumeCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *VolumeCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
