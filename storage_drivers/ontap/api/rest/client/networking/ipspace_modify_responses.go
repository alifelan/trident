// Code generated by go-swagger; DO NOT EDIT.

package networking

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// IpspaceModifyReader is a Reader for the IpspaceModify structure.
type IpspaceModifyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IpspaceModifyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewIpspaceModifyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewIpspaceModifyOK creates a IpspaceModifyOK with default headers values
func NewIpspaceModifyOK() *IpspaceModifyOK {
	return &IpspaceModifyOK{}
}

/*
IpspaceModifyOK describes a response with status code 200, with default header values.

OK
*/
type IpspaceModifyOK struct {
}

// IsSuccess returns true when this ipspace modify o k response has a 2xx status code
func (o *IpspaceModifyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this ipspace modify o k response has a 3xx status code
func (o *IpspaceModifyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this ipspace modify o k response has a 4xx status code
func (o *IpspaceModifyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this ipspace modify o k response has a 5xx status code
func (o *IpspaceModifyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this ipspace modify o k response a status code equal to that given
func (o *IpspaceModifyOK) IsCode(code int) bool {
	return code == 200
}

func (o *IpspaceModifyOK) Error() string {
	return fmt.Sprintf("[PATCH /network/ipspaces/{uuid}][%d] ipspaceModifyOK ", 200)
}

func (o *IpspaceModifyOK) String() string {
	return fmt.Sprintf("[PATCH /network/ipspaces/{uuid}][%d] ipspaceModifyOK ", 200)
}

func (o *IpspaceModifyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
