// Code generated by go-swagger; DO NOT EDIT.

package amendment_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

// AmendmentServiceChangeQuantityReader is a Reader for the AmendmentServiceChangeQuantity structure.
type AmendmentServiceChangeQuantityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AmendmentServiceChangeQuantityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAmendmentServiceChangeQuantityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAmendmentServiceChangeQuantityDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAmendmentServiceChangeQuantityOK creates a AmendmentServiceChangeQuantityOK with default headers values
func NewAmendmentServiceChangeQuantityOK() *AmendmentServiceChangeQuantityOK {
	return &AmendmentServiceChangeQuantityOK{}
}

/*
AmendmentServiceChangeQuantityOK describes a response with status code 200, with default header values.

A successful response.
*/
type AmendmentServiceChangeQuantityOK struct {
	Payload *models.V1ChangeQuantityResponse
}

// IsSuccess returns true when this amendment service change quantity o k response has a 2xx status code
func (o *AmendmentServiceChangeQuantityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this amendment service change quantity o k response has a 3xx status code
func (o *AmendmentServiceChangeQuantityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this amendment service change quantity o k response has a 4xx status code
func (o *AmendmentServiceChangeQuantityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this amendment service change quantity o k response has a 5xx status code
func (o *AmendmentServiceChangeQuantityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this amendment service change quantity o k response a status code equal to that given
func (o *AmendmentServiceChangeQuantityOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the amendment service change quantity o k response
func (o *AmendmentServiceChangeQuantityOK) Code() int {
	return 200
}

func (o *AmendmentServiceChangeQuantityOK) Error() string {
	return fmt.Sprintf("[POST /v1/amendment/changequantity][%d] amendmentServiceChangeQuantityOK  %+v", 200, o.Payload)
}

func (o *AmendmentServiceChangeQuantityOK) String() string {
	return fmt.Sprintf("[POST /v1/amendment/changequantity][%d] amendmentServiceChangeQuantityOK  %+v", 200, o.Payload)
}

func (o *AmendmentServiceChangeQuantityOK) GetPayload() *models.V1ChangeQuantityResponse {
	return o.Payload
}

func (o *AmendmentServiceChangeQuantityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ChangeQuantityResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAmendmentServiceChangeQuantityDefault creates a AmendmentServiceChangeQuantityDefault with default headers values
func NewAmendmentServiceChangeQuantityDefault(code int) *AmendmentServiceChangeQuantityDefault {
	return &AmendmentServiceChangeQuantityDefault{
		_statusCode: code,
	}
}

/*
AmendmentServiceChangeQuantityDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AmendmentServiceChangeQuantityDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this amendment service change quantity default response has a 2xx status code
func (o *AmendmentServiceChangeQuantityDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this amendment service change quantity default response has a 3xx status code
func (o *AmendmentServiceChangeQuantityDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this amendment service change quantity default response has a 4xx status code
func (o *AmendmentServiceChangeQuantityDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this amendment service change quantity default response has a 5xx status code
func (o *AmendmentServiceChangeQuantityDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this amendment service change quantity default response a status code equal to that given
func (o *AmendmentServiceChangeQuantityDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the amendment service change quantity default response
func (o *AmendmentServiceChangeQuantityDefault) Code() int {
	return o._statusCode
}

func (o *AmendmentServiceChangeQuantityDefault) Error() string {
	return fmt.Sprintf("[POST /v1/amendment/changequantity][%d] AmendmentService_ChangeQuantity default  %+v", o._statusCode, o.Payload)
}

func (o *AmendmentServiceChangeQuantityDefault) String() string {
	return fmt.Sprintf("[POST /v1/amendment/changequantity][%d] AmendmentService_ChangeQuantity default  %+v", o._statusCode, o.Payload)
}

func (o *AmendmentServiceChangeQuantityDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *AmendmentServiceChangeQuantityDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
