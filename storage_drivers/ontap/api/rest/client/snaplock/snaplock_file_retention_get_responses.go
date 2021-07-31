// Code generated by go-swagger; DO NOT EDIT.

package snaplock

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netapp/trident/storage_drivers/ontap/api/rest/models"
)

// SnaplockFileRetentionGetReader is a Reader for the SnaplockFileRetentionGet structure.
type SnaplockFileRetentionGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SnaplockFileRetentionGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSnaplockFileRetentionGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewSnaplockFileRetentionGetDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewSnaplockFileRetentionGetOK creates a SnaplockFileRetentionGetOK with default headers values
func NewSnaplockFileRetentionGetOK() *SnaplockFileRetentionGetOK {
	return &SnaplockFileRetentionGetOK{}
}

/* SnaplockFileRetentionGetOK describes a response with status code 200, with default header values.

OK
*/
type SnaplockFileRetentionGetOK struct {
	Payload *models.SnaplockFileRetention
}

func (o *SnaplockFileRetentionGetOK) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/file/{volume.uuid}/{path}][%d] snaplockFileRetentionGetOK  %+v", 200, o.Payload)
}
func (o *SnaplockFileRetentionGetOK) GetPayload() *models.SnaplockFileRetention {
	return o.Payload
}

func (o *SnaplockFileRetentionGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.SnaplockFileRetention)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewSnaplockFileRetentionGetDefault creates a SnaplockFileRetentionGetDefault with default headers values
func NewSnaplockFileRetentionGetDefault(code int) *SnaplockFileRetentionGetDefault {
	return &SnaplockFileRetentionGetDefault{
		_statusCode: code,
	}
}

/* SnaplockFileRetentionGetDefault describes a response with status code -1, with default header values.

 ONTAP Error Response codes
| Error code  |  Description |
|-------------|--------------|
| 14090347    | File path must be in the format \"/<dir>/<file path>\"  |

*/
type SnaplockFileRetentionGetDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the snaplock file retention get default response
func (o *SnaplockFileRetentionGetDefault) Code() int {
	return o._statusCode
}

func (o *SnaplockFileRetentionGetDefault) Error() string {
	return fmt.Sprintf("[GET /storage/snaplock/file/{volume.uuid}/{path}][%d] snaplock_file_retention_get default  %+v", o._statusCode, o.Payload)
}
func (o *SnaplockFileRetentionGetDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *SnaplockFileRetentionGetDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}