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

// S3UserCollectionGetReader is a Reader for the S3UserCollectionGet structure.
type S3UserCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *S3UserCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewS3UserCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewS3UserCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewS3UserCollectionGetOK creates a S3UserCollectionGetOK with default headers values
func NewS3UserCollectionGetOK() *S3UserCollectionGetOK {
	return &S3UserCollectionGetOK{}
}

/*
S3UserCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type S3UserCollectionGetOK struct {
	Payload *models.S3UserResponse
}

// IsSuccess returns true when this s3 user collection get o k response has a 2xx status code
func (o *S3UserCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this s3 user collection get o k response has a 3xx status code
func (o *S3UserCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this s3 user collection get o k response has a 4xx status code
func (o *S3UserCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this s3 user collection get o k response has a 5xx status code
func (o *S3UserCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this s3 user collection get o k response a status code equal to that given
func (o *S3UserCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *S3UserCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/users][%d] s3UserCollectionGetOK  %+v", 200, o.Payload)
}

func (o *S3UserCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/users][%d] s3UserCollectionGetOK  %+v", 200, o.Payload)
}

func (o *S3UserCollectionGetOK) GetPayload() *models.S3UserResponse {
	return o.Payload
}

func (o *S3UserCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.S3UserResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewS3UserCollectionGetDefault creates a S3UserCollectionGetDefault with default headers values
func NewS3UserCollectionGetDefault(code int) *S3UserCollectionGetDefault {
	return &S3UserCollectionGetDefault{
		_statusCode: code,
	}
}

/*
S3UserCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type S3UserCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the s3 user collection get default response
func (o *S3UserCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this s3 user collection get default response has a 2xx status code
func (o *S3UserCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this s3 user collection get default response has a 3xx status code
func (o *S3UserCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this s3 user collection get default response has a 4xx status code
func (o *S3UserCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this s3 user collection get default response has a 5xx status code
func (o *S3UserCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this s3 user collection get default response a status code equal to that given
func (o *S3UserCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *S3UserCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/users][%d] s3_user_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *S3UserCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/s3/services/{svm.uuid}/users][%d] s3_user_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *S3UserCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *S3UserCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
