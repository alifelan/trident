// Code generated by go-swagger; DO NOT EDIT.

package security

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// ClusterNisCreateReader is a Reader for the ClusterNisCreate structure.
type ClusterNisCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClusterNisCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewClusterNisCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewClusterNisCreateDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewClusterNisCreateCreated creates a ClusterNisCreateCreated with default headers values
func NewClusterNisCreateCreated() *ClusterNisCreateCreated {
	return &ClusterNisCreateCreated{}
}

/*
ClusterNisCreateCreated describes a response with status code 201, with default header values.

Created
*/
type ClusterNisCreateCreated struct {
	Payload *models.ClusterNisServiceResponse
}

// IsSuccess returns true when this cluster nis create created response has a 2xx status code
func (o *ClusterNisCreateCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cluster nis create created response has a 3xx status code
func (o *ClusterNisCreateCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cluster nis create created response has a 4xx status code
func (o *ClusterNisCreateCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this cluster nis create created response has a 5xx status code
func (o *ClusterNisCreateCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this cluster nis create created response a status code equal to that given
func (o *ClusterNisCreateCreated) IsCode(code int) bool {
	return code == 201
}

func (o *ClusterNisCreateCreated) Error() string {
	return fmt.Sprintf("[POST /security/authentication/cluster/nis][%d] clusterNisCreateCreated  %+v", 201, o.Payload)
}

func (o *ClusterNisCreateCreated) String() string {
	return fmt.Sprintf("[POST /security/authentication/cluster/nis][%d] clusterNisCreateCreated  %+v", 201, o.Payload)
}

func (o *ClusterNisCreateCreated) GetPayload() *models.ClusterNisServiceResponse {
	return o.Payload
}

func (o *ClusterNisCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ClusterNisServiceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewClusterNisCreateDefault creates a ClusterNisCreateDefault with default headers values
func NewClusterNisCreateDefault(code int) *ClusterNisCreateDefault {
	return &ClusterNisCreateDefault{
		_statusCode: code,
	}
}

/*
	ClusterNisCreateDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 1966253    | IPv6 is not enabled in the cluster. |
| 3276964    | The NIS domain name or NIS server domain is too long. The maximum supported for domain name is 64 characters and the maximum supported for NIS server domain is 255 characters. |
| 3276933    | A maximum of 10 NIS servers can be configured per SVM. |
| 13434916   | The SVM is in the process of being created. Wait a few minutes, and then try the command again. |
| 23724109   | DNS resolution failed for one or more specified servers.  |
| 23724112   | DNS resolution failed due to an internal error. Contact technical support if this issue persists.  |
| 23724132   | DNS resolution failed for all the specified servers.  |
| 23724130   | Cannot use an IPv6 name server address because there are no IPv6 interfaces. |
*/
type ClusterNisCreateDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cluster nis create default response
func (o *ClusterNisCreateDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cluster nis create default response has a 2xx status code
func (o *ClusterNisCreateDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cluster nis create default response has a 3xx status code
func (o *ClusterNisCreateDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cluster nis create default response has a 4xx status code
func (o *ClusterNisCreateDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cluster nis create default response has a 5xx status code
func (o *ClusterNisCreateDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cluster nis create default response a status code equal to that given
func (o *ClusterNisCreateDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *ClusterNisCreateDefault) Error() string {
	return fmt.Sprintf("[POST /security/authentication/cluster/nis][%d] cluster_nis_create default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterNisCreateDefault) String() string {
	return fmt.Sprintf("[POST /security/authentication/cluster/nis][%d] cluster_nis_create default  %+v", o._statusCode, o.Payload)
}

func (o *ClusterNisCreateDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ClusterNisCreateDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
