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

// CifsSessionGetReader is a Reader for the CifsSessionGet structure.
type CifsSessionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CifsSessionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCifsSessionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCifsSessionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCifsSessionGetOK creates a CifsSessionGetOK with default headers values
func NewCifsSessionGetOK() *CifsSessionGetOK {
	return &CifsSessionGetOK{}
}

/*
CifsSessionGetOK describes a response with status code 200, with default header values.

OK
*/
type CifsSessionGetOK struct {
	Payload *models.CifsSession
}

// IsSuccess returns true when this cifs session get o k response has a 2xx status code
func (o *CifsSessionGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this cifs session get o k response has a 3xx status code
func (o *CifsSessionGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this cifs session get o k response has a 4xx status code
func (o *CifsSessionGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this cifs session get o k response has a 5xx status code
func (o *CifsSessionGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this cifs session get o k response a status code equal to that given
func (o *CifsSessionGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *CifsSessionGetOK) Error() string {
	return fmt.Sprintf("[GET /protocols/cifs/sessions/{node.uuid}/{svm.uuid}/{identifier}/{connection_id}][%d] cifsSessionGetOK  %+v", 200, o.Payload)
}

func (o *CifsSessionGetOK) String() string {
	return fmt.Sprintf("[GET /protocols/cifs/sessions/{node.uuid}/{svm.uuid}/{identifier}/{connection_id}][%d] cifsSessionGetOK  %+v", 200, o.Payload)
}

func (o *CifsSessionGetOK) GetPayload() *models.CifsSession {
	return o.Payload
}

func (o *CifsSessionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CifsSession)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCifsSessionGetDefault creates a CifsSessionGetDefault with default headers values
func NewCifsSessionGetDefault(code int) *CifsSessionGetDefault {
	return &CifsSessionGetDefault{
		_statusCode: code,
	}
}

/*
CifsSessionGetDefault describes a response with status code -1, with default header values.

Error
*/
type CifsSessionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the cifs session get default response
func (o *CifsSessionGetDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this cifs session get default response has a 2xx status code
func (o *CifsSessionGetDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this cifs session get default response has a 3xx status code
func (o *CifsSessionGetDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this cifs session get default response has a 4xx status code
func (o *CifsSessionGetDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this cifs session get default response has a 5xx status code
func (o *CifsSessionGetDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this cifs session get default response a status code equal to that given
func (o *CifsSessionGetDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *CifsSessionGetDefault) Error() string {
	return fmt.Sprintf("[GET /protocols/cifs/sessions/{node.uuid}/{svm.uuid}/{identifier}/{connection_id}][%d] cifs_session_get default  %+v", o._statusCode, o.Payload)
}

func (o *CifsSessionGetDefault) String() string {
	return fmt.Sprintf("[GET /protocols/cifs/sessions/{node.uuid}/{svm.uuid}/{identifier}/{connection_id}][%d] cifs_session_get default  %+v", o._statusCode, o.Payload)
}

func (o *CifsSessionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *CifsSessionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
