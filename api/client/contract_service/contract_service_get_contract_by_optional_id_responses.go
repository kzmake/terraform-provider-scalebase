// Code generated by go-swagger; DO NOT EDIT.

package contract_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

// ContractServiceGetContractByOptionalIDReader is a Reader for the ContractServiceGetContractByOptionalID structure.
type ContractServiceGetContractByOptionalIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContractServiceGetContractByOptionalIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContractServiceGetContractByOptionalIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewContractServiceGetContractByOptionalIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewContractServiceGetContractByOptionalIDOK creates a ContractServiceGetContractByOptionalIDOK with default headers values
func NewContractServiceGetContractByOptionalIDOK() *ContractServiceGetContractByOptionalIDOK {
	return &ContractServiceGetContractByOptionalIDOK{}
}

/*
ContractServiceGetContractByOptionalIDOK describes a response with status code 200, with default header values.

A successful response.
*/
type ContractServiceGetContractByOptionalIDOK struct {
	Payload *models.V1GetContractByOptionalIDResponse
}

// IsSuccess returns true when this contract service get contract by optional Id o k response has a 2xx status code
func (o *ContractServiceGetContractByOptionalIDOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contract service get contract by optional Id o k response has a 3xx status code
func (o *ContractServiceGetContractByOptionalIDOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contract service get contract by optional Id o k response has a 4xx status code
func (o *ContractServiceGetContractByOptionalIDOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contract service get contract by optional Id o k response has a 5xx status code
func (o *ContractServiceGetContractByOptionalIDOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contract service get contract by optional Id o k response a status code equal to that given
func (o *ContractServiceGetContractByOptionalIDOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contract service get contract by optional Id o k response
func (o *ContractServiceGetContractByOptionalIDOK) Code() int {
	return 200
}

func (o *ContractServiceGetContractByOptionalIDOK) Error() string {
	return fmt.Sprintf("[POST /v1/contract/getbyoptionalid][%d] contractServiceGetContractByOptionalIdOK  %+v", 200, o.Payload)
}

func (o *ContractServiceGetContractByOptionalIDOK) String() string {
	return fmt.Sprintf("[POST /v1/contract/getbyoptionalid][%d] contractServiceGetContractByOptionalIdOK  %+v", 200, o.Payload)
}

func (o *ContractServiceGetContractByOptionalIDOK) GetPayload() *models.V1GetContractByOptionalIDResponse {
	return o.Payload
}

func (o *ContractServiceGetContractByOptionalIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetContractByOptionalIDResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContractServiceGetContractByOptionalIDDefault creates a ContractServiceGetContractByOptionalIDDefault with default headers values
func NewContractServiceGetContractByOptionalIDDefault(code int) *ContractServiceGetContractByOptionalIDDefault {
	return &ContractServiceGetContractByOptionalIDDefault{
		_statusCode: code,
	}
}

/*
ContractServiceGetContractByOptionalIDDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ContractServiceGetContractByOptionalIDDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this contract service get contract by optional Id default response has a 2xx status code
func (o *ContractServiceGetContractByOptionalIDDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this contract service get contract by optional Id default response has a 3xx status code
func (o *ContractServiceGetContractByOptionalIDDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this contract service get contract by optional Id default response has a 4xx status code
func (o *ContractServiceGetContractByOptionalIDDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this contract service get contract by optional Id default response has a 5xx status code
func (o *ContractServiceGetContractByOptionalIDDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this contract service get contract by optional Id default response a status code equal to that given
func (o *ContractServiceGetContractByOptionalIDDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the contract service get contract by optional Id default response
func (o *ContractServiceGetContractByOptionalIDDefault) Code() int {
	return o._statusCode
}

func (o *ContractServiceGetContractByOptionalIDDefault) Error() string {
	return fmt.Sprintf("[POST /v1/contract/getbyoptionalid][%d] ContractService_GetContractByOptionalId default  %+v", o._statusCode, o.Payload)
}

func (o *ContractServiceGetContractByOptionalIDDefault) String() string {
	return fmt.Sprintf("[POST /v1/contract/getbyoptionalid][%d] ContractService_GetContractByOptionalId default  %+v", o._statusCode, o.Payload)
}

func (o *ContractServiceGetContractByOptionalIDDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ContractServiceGetContractByOptionalIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
