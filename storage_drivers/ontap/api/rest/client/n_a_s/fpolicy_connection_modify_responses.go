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

// FpolicyConnectionModifyReader is a Reader for the FpolicyConnectionModify structure.
type FpolicyConnectionModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyConnectionModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyConnectionModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyConnectionModifyDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyConnectionModifyOK creates a FpolicyConnectionModifyOK with default headers values
func NewFpolicyConnectionModifyOK() *FpolicyConnectionModifyOK {
	return &FpolicyConnectionModifyOK{}
}

/*
FpolicyConnectionModifyOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyConnectionModifyOK struct {
}

// IsSuccess returns true when this fpolicy connection modify o k response has a 2xx status code
func (o *FpolicyConnectionModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this fpolicy connection modify o k response has a 3xx status code
func (o *FpolicyConnectionModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this fpolicy connection modify o k response has a 4xx status code
func (o *FpolicyConnectionModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this fpolicy connection modify o k response has a 5xx status code
func (o *FpolicyConnectionModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this fpolicy connection modify o k response a status code equal to that given
func (o *FpolicyConnectionModifyOK) IsCode(code int) bool {
	return code == 200
}

func (o *FpolicyConnectionModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/connections/{node.uuid}/{policy.name}/{server}][%d] fpolicyConnectionModifyOK ", 200)
}

func (o *FpolicyConnectionModifyOK) String() string {
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/connections/{node.uuid}/{policy.name}/{server}][%d] fpolicyConnectionModifyOK ", 200)
}

func (o *FpolicyConnectionModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFpolicyConnectionModifyDefault creates a FpolicyConnectionModifyDefault with default headers values
func NewFpolicyConnectionModifyDefault(code int) *FpolicyConnectionModifyDefault {
	return &FpolicyConnectionModifyDefault{
		_statusCode: code,
	}
}

/*
	FpolicyConnectionModifyDefault describes a response with status code -1, with default header values.

	ONTAP Error Response Codes

| Error Code | Description |
| ---------- | ----------- |
| 9764954    | The specified policy does not exist |
| 9764911    | Failed to connect to the FPolicy server. Reason: The specified entry does not exist |
| 9764948    | The specified policy is disabled. Using a disabled policy with this command is not supported. Use the 'fpolicy enable' command to enable the policy |
| 9764912    | Failed to disconnect the FPolicy server. Reason: The specified entry does not exist |
*/
type FpolicyConnectionModifyDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fpolicy connection modify default response
func (o *FpolicyConnectionModifyDefault) Code() int {
	return o._statusCode
}

// IsSuccess returns true when this fpolicy connection modify default response has a 2xx status code
func (o *FpolicyConnectionModifyDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this fpolicy connection modify default response has a 3xx status code
func (o *FpolicyConnectionModifyDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this fpolicy connection modify default response has a 4xx status code
func (o *FpolicyConnectionModifyDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this fpolicy connection modify default response has a 5xx status code
func (o *FpolicyConnectionModifyDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this fpolicy connection modify default response a status code equal to that given
func (o *FpolicyConnectionModifyDefault) IsCode(code int) bool {
	return o._statusCode == code
}

func (o *FpolicyConnectionModifyDefault) Error() string {
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/connections/{node.uuid}/{policy.name}/{server}][%d] fpolicy_connection_modify default  %+v", o._statusCode, o.Payload)
}

func (o *FpolicyConnectionModifyDefault) String() string {
	return fmt.Sprintf("[PATCH /protocols/fpolicy/{svm.uuid}/connections/{node.uuid}/{policy.name}/{server}][%d] fpolicy_connection_modify default  %+v", o._statusCode, o.Payload)
}

func (o *FpolicyConnectionModifyDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyConnectionModifyDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
