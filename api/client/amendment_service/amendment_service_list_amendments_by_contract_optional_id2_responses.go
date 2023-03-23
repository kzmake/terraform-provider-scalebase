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

// AmendmentServiceListAmendmentsByContractOptionalId2Reader is a Reader for the AmendmentServiceListAmendmentsByContractOptionalId2 structure.
type AmendmentServiceListAmendmentsByContractOptionalId2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AmendmentServiceListAmendmentsByContractOptionalId2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewAmendmentServiceListAmendmentsByContractOptionalId2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewAmendmentServiceListAmendmentsByContractOptionalId2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAmendmentServiceListAmendmentsByContractOptionalId2OK creates a AmendmentServiceListAmendmentsByContractOptionalId2OK with default headers values
func NewAmendmentServiceListAmendmentsByContractOptionalId2OK() *AmendmentServiceListAmendmentsByContractOptionalId2OK {
	return &AmendmentServiceListAmendmentsByContractOptionalId2OK{}
}

/*
AmendmentServiceListAmendmentsByContractOptionalId2OK describes a response with status code 200, with default header values.

A successful response.
*/
type AmendmentServiceListAmendmentsByContractOptionalId2OK struct {
	Payload *models.V1ListAmendmentsByContractOptionalIDResponse
}

// IsSuccess returns true when this amendment service list amendments by contract optional id2 o k response has a 2xx status code
func (o *AmendmentServiceListAmendmentsByContractOptionalId2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this amendment service list amendments by contract optional id2 o k response has a 3xx status code
func (o *AmendmentServiceListAmendmentsByContractOptionalId2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this amendment service list amendments by contract optional id2 o k response has a 4xx status code
func (o *AmendmentServiceListAmendmentsByContractOptionalId2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this amendment service list amendments by contract optional id2 o k response has a 5xx status code
func (o *AmendmentServiceListAmendmentsByContractOptionalId2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this amendment service list amendments by contract optional id2 o k response a status code equal to that given
func (o *AmendmentServiceListAmendmentsByContractOptionalId2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the amendment service list amendments by contract optional id2 o k response
func (o *AmendmentServiceListAmendmentsByContractOptionalId2OK) Code() int {
	return 200
}

func (o *AmendmentServiceListAmendmentsByContractOptionalId2OK) Error() string {
	return fmt.Sprintf("[GET /v1/amendment/listbycontractoptionalid][%d] amendmentServiceListAmendmentsByContractOptionalId2OK  %+v", 200, o.Payload)
}

func (o *AmendmentServiceListAmendmentsByContractOptionalId2OK) String() string {
	return fmt.Sprintf("[GET /v1/amendment/listbycontractoptionalid][%d] amendmentServiceListAmendmentsByContractOptionalId2OK  %+v", 200, o.Payload)
}

func (o *AmendmentServiceListAmendmentsByContractOptionalId2OK) GetPayload() *models.V1ListAmendmentsByContractOptionalIDResponse {
	return o.Payload
}

func (o *AmendmentServiceListAmendmentsByContractOptionalId2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ListAmendmentsByContractOptionalIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAmendmentServiceListAmendmentsByContractOptionalId2Default creates a AmendmentServiceListAmendmentsByContractOptionalId2Default with default headers values
func NewAmendmentServiceListAmendmentsByContractOptionalId2Default(code int) *AmendmentServiceListAmendmentsByContractOptionalId2Default {
	return &AmendmentServiceListAmendmentsByContractOptionalId2Default{
		_statusCode: code,
	}
}

/*
AmendmentServiceListAmendmentsByContractOptionalId2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type AmendmentServiceListAmendmentsByContractOptionalId2Default struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this amendment service list amendments by contract optional id2 default response has a 2xx status code
func (o *AmendmentServiceListAmendmentsByContractOptionalId2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this amendment service list amendments by contract optional id2 default response has a 3xx status code
func (o *AmendmentServiceListAmendmentsByContractOptionalId2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this amendment service list amendments by contract optional id2 default response has a 4xx status code
func (o *AmendmentServiceListAmendmentsByContractOptionalId2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this amendment service list amendments by contract optional id2 default response has a 5xx status code
func (o *AmendmentServiceListAmendmentsByContractOptionalId2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this amendment service list amendments by contract optional id2 default response a status code equal to that given
func (o *AmendmentServiceListAmendmentsByContractOptionalId2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the amendment service list amendments by contract optional id2 default response
func (o *AmendmentServiceListAmendmentsByContractOptionalId2Default) Code() int {
	return o._statusCode
}

func (o *AmendmentServiceListAmendmentsByContractOptionalId2Default) Error() string {
	return fmt.Sprintf("[GET /v1/amendment/listbycontractoptionalid][%d] AmendmentService_ListAmendmentsByContractOptionalId2 default  %+v", o._statusCode, o.Payload)
}

func (o *AmendmentServiceListAmendmentsByContractOptionalId2Default) String() string {
	return fmt.Sprintf("[GET /v1/amendment/listbycontractoptionalid][%d] AmendmentService_ListAmendmentsByContractOptionalId2 default  %+v", o._statusCode, o.Payload)
}

func (o *AmendmentServiceListAmendmentsByContractOptionalId2Default) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *AmendmentServiceListAmendmentsByContractOptionalId2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
