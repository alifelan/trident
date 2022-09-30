// Code generated by go-swagger; DO NOT EDIT.

package n_a_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// CifsSymlinkMappingCollectionGetReader is a Reader for the CifsSymlinkMappingCollectionGet structure.
type CifsSymlinkMappingCollectionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsSymlinkMappingCollectionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsSymlinkMappingCollectionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsSymlinkMappingCollectionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsSymlinkMappingCollectionGetOK creates a CifsSymlinkMappingCollectionGetOK with default headers values
func NewCifsSymlinkMappingCollectionGetOK() *CifsSymlinkMappingCollectionGetOK {
	return &CifsSymlinkMappingCollectionGetOK{}
}

/*
CifsSymlinkMappingCollectionGetOK describes a response with status code 200, with default header values.

OK
*/
type CifsSymlinkMappingCollectionGetOK struct {
	Payload *models.CifsSymlinkMappingResponse
}

// IsSuccess returns true when this cifs symlink mapping collection get o k response has a 2xx status code
func (o *CifsSymlinkMappingCollectionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs symlink mapping collection get o k response has a 3xx status code
func (o *CifsSymlinkMappingCollectionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs symlink mapping collection get o k response has a 4xx status code
func (o *CifsSymlinkMappingCollectionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs symlink mapping collection get o k response has a 5xx status code
func (o *CifsSymlinkMappingCollectionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs symlink mapping collection get o k response a status code equal to that given
func (o *CifsSymlinkMappingCollectionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *CifsSymlinkMappingCollectionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/cifs/unix-symlink-mapping][%d] cifsSymlinkMappingCollectionGetOK  %+v", 200, o.Payload)
}

func (o *CifsSymlinkMappingCollectionGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/cifs/unix-symlink-mapping][%d] cifsSymlinkMappingCollectionGetOK  %+v", 200, o.Payload)
}

func (o *CifsSymlinkMappingCollectionGetOK) GetPayload() *models.CifsSymlinkMappingResponse {
	return o.Payload
}

func (o *CifsSymlinkMappingCollectionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsSymlinkMappingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsSymlinkMappingCollectionGetDefault creates a CifsSymlinkMappingCollectionGetDefault with default headers values
func NewCifsSymlinkMappingCollectionGetDefault(code int) *CifsSymlinkMappingCollectionGetDefault {
	return &CifsSymlinkMappingCollectionGetDefault{
		_statusCode: code,
	}
}

/*
CifsSymlinkMappingCollectionGetDefault describes a response with status code -1, with default header values.

Error
*/
type CifsSymlinkMappingCollectionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cifs symlink mapping collection get default response
func (o *CifsSymlinkMappingCollectionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cifs symlink mapping collection get default response has a 2xx status code
func (o *CifsSymlinkMappingCollectionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs symlink mapping collection get default response has a 3xx status code
func (o *CifsSymlinkMappingCollectionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs symlink mapping collection get default response has a 4xx status code
func (o *CifsSymlinkMappingCollectionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs symlink mapping collection get default response has a 5xx status code
func (o *CifsSymlinkMappingCollectionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs symlink mapping collection get default response a status code equal to that given
func (o *CifsSymlinkMappingCollectionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CifsSymlinkMappingCollectionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/cifs/unix-symlink-mapping][%d] cifs_symlink_mapping_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *CifsSymlinkMappingCollectionGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/cifs/unix-symlink-mapping][%d] cifs_symlink_mapping_collection_get default  %+v", o._statusCode, o.Payload)
}

func (o *CifsSymlinkMappingCollectionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsSymlinkMappingCollectionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
