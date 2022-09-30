// Code generated by go-swagger; DO NOT EDIT.

package object_store

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// S3PolicyCollectionGetReader is a Reader for the S3PolicyCollectionGet structure.
type S3PolicyCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3PolicyCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3PolicyCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3PolicyCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3PolicyCollectionGetOK creates a S3PolicyCollectionGetOK with default headers values
func NewS3PolicyCollectionGetOK() *S3PolicyCollectionGetOK {
	return &S3PolicyCollectionGetOK{}
}

/*
S3PolicyCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type S3PolicyCollectionGetOK struct {
	Payload *models.S3PolicyResponse
}

// IsSuccess returns true when this s3 policy collection get o k response has a 2xx status code
func (o *S3PolicyCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 policy collection get o k response has a 3xx status code
func (o *S3PolicyCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 policy collection get o k response has a 4xx status code
func (o *S3PolicyCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 policy collection get o k response has a 5xx status code
func (o *S3PolicyCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 policy collection get o k response a status code equal to that given
func (o *S3PolicyCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *S3PolicyCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/policies][%d] s3PolicyCollectionGetOK  %+v", 200, o.Payload)
}

func (o *S3PolicyCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/policies][%d] s3PolicyCollectionGetOK  %+v", 200, o.Payload)
}

func (o *S3PolicyCollectionGetOK) GetPayload() *models.S3PolicyResponse {
	return o.Payload
}

func (o *S3PolicyCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3PolicyResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3PolicyCollectionGetDefault creates a S3PolicyCollectionGetDefault with default headers values
func NewS3PolicyCollectionGetDefault(code int) *S3PolicyCollectionGetDefault {
	return &S3PolicyCollectionGetDefault{
		_statusCode: code,
	}
}

/*
S3PolicyCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type S3PolicyCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the s3 policy collection get default response
func (o *S3PolicyCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this s3 policy collection get default response has a 2xx status code
func (o *S3PolicyCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 policy collection get default response has a 3xx status code
func (o *S3PolicyCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 policy collection get default response has a 4xx status code
func (o *S3PolicyCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 policy collection get default response has a 5xx status code
func (o *S3PolicyCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 policy collection get default response a status code equal to that given
func (o *S3PolicyCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *S3PolicyCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/policies][%d] s3_policy_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *S3PolicyCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/policies][%d] s3_policy_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *S3PolicyCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3PolicyCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
