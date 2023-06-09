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

// AmendmentServiceChangeMainItemReader is a Reader for the AmendmentServiceChangeMainItem structure.
type AmendmentServiceChangeMainItemReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AmendmentServiceChangeMainItemReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAmendmentServiceChangeMainItemOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAmendmentServiceChangeMainItemDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAmendmentServiceChangeMainItemOK creates a AmendmentServiceChangeMainItemOK with default headers values
func NewAmendmentServiceChangeMainItemOK() *AmendmentServiceChangeMainItemOK {
	return &AmendmentServiceChangeMainItemOK{}
}

/*
AmendmentServiceChangeMainItemOK describes a response with status code 200, with default header values.

A successful response.
*/
type AmendmentServiceChangeMainItemOK struct {
	Payload *models.V1ChangeMainItemResponse
}

// IsSuccess returns true when this amendment service change main item o k response has a 2xx status code
func (o *AmendmentServiceChangeMainItemOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this amendment service change main item o k response has a 3xx status code
func (o *AmendmentServiceChangeMainItemOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this amendment service change main item o k response has a 4xx status code
func (o *AmendmentServiceChangeMainItemOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this amendment service change main item o k response has a 5xx status code
func (o *AmendmentServiceChangeMainItemOK) IsServerError() bool {
	return false
}

// IsCode returns true when this amendment service change main item o k response a status code equal to that given
func (o *AmendmentServiceChangeMainItemOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the amendment service change main item o k response
func (o *AmendmentServiceChangeMainItemOK) Code() int {
	return 200
}

func (o *AmendmentServiceChangeMainItemOK) Error() string {
	return fmt.Sprintf("[POST /v1/amendment/changemainitem][%d] amendmentServiceChangeMainItemOK  %+v", 200, o.Payload)
}

func (o *AmendmentServiceChangeMainItemOK) String() string {
	return fmt.Sprintf("[POST /v1/amendment/changemainitem][%d] amendmentServiceChangeMainItemOK  %+v", 200, o.Payload)
}

func (o *AmendmentServiceChangeMainItemOK) GetPayload() *models.V1ChangeMainItemResponse {
	return o.Payload
}

func (o *AmendmentServiceChangeMainItemOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ChangeMainItemResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAmendmentServiceChangeMainItemDefault creates a AmendmentServiceChangeMainItemDefault with default headers values
func NewAmendmentServiceChangeMainItemDefault(code int) *AmendmentServiceChangeMainItemDefault {
	return &AmendmentServiceChangeMainItemDefault{
		_statusCode: code,
	}
}

/*
AmendmentServiceChangeMainItemDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AmendmentServiceChangeMainItemDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this amendment service change main item default response has a 2xx status code
func (o *AmendmentServiceChangeMainItemDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this amendment service change main item default response has a 3xx status code
func (o *AmendmentServiceChangeMainItemDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this amendment service change main item default response has a 4xx status code
func (o *AmendmentServiceChangeMainItemDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this amendment service change main item default response has a 5xx status code
func (o *AmendmentServiceChangeMainItemDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this amendment service change main item default response a status code equal to that given
func (o *AmendmentServiceChangeMainItemDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the amendment service change main item default response
func (o *AmendmentServiceChangeMainItemDefault) Code() int {
	return o._statusCode
}

func (o *AmendmentServiceChangeMainItemDefault) Error() string {
	return fmt.Sprintf("[POST /v1/amendment/changemainitem][%d] AmendmentService_ChangeMainItem default  %+v", o._statusCode, o.Payload)
}

func (o *AmendmentServiceChangeMainItemDefault) String() string {
	return fmt.Sprintf("[POST /v1/amendment/changemainitem][%d] AmendmentService_ChangeMainItem default  %+v", o._statusCode, o.Payload)
}

func (o *AmendmentServiceChangeMainItemDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *AmendmentServiceChangeMainItemDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
