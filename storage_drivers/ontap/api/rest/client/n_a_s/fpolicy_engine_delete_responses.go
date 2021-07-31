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

// FpolicyEngineDeleteReader is a Reader for the FpolicyEngineDelete structure.
type FpolicyEngineDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *FpolicyEngineDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewFpolicyEngineDeleteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewFpolicyEngineDeleteDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewFpolicyEngineDeleteOK creates a FpolicyEngineDeleteOK with default headers values
func NewFpolicyEngineDeleteOK() *FpolicyEngineDeleteOK {
	return &FpolicyEngineDeleteOK{}
}

/* FpolicyEngineDeleteOK describes a response with status code 200, with default header values.

OK
*/
type FpolicyEngineDeleteOK struct {
}

func (o *FpolicyEngineDeleteOK) Error() string {
	return fmt.Sprintf("[DELETE /protocols/fpolicy/{svm.uuid}/engines/{name}][%d] fpolicyEngineDeleteOK ", 200)
}

func (o *FpolicyEngineDeleteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewFpolicyEngineDeleteDefault creates a FpolicyEngineDeleteDefault with default headers values
func NewFpolicyEngineDeleteDefault(code int) *FpolicyEngineDeleteDefault {
	return &FpolicyEngineDeleteDefault{
		_statusCode: code,
	}
}

/* FpolicyEngineDeleteDefault describes a response with status code -1, with default header values.

 ONTAP Error Response Codes
| Error Code | Description |
| ---------- | ----------- |
| 9764940    | At least one FPolicy policy is using the FPolicy engine |
| 9764887    | The FPolicy engine is a cluster level FPolicy engine |

*/
type FpolicyEngineDeleteDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the fpolicy engine delete default response
func (o *FpolicyEngineDeleteDefault) Code() int {
	return o._statusCode
}

func (o *FpolicyEngineDeleteDefault) Error() string {
	return fmt.Sprintf("[DELETE /protocols/fpolicy/{svm.uuid}/engines/{name}][%d] fpolicy_engine_delete default  %+v", o._statusCode, o.Payload)
}
func (o *FpolicyEngineDeleteDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *FpolicyEngineDeleteDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}