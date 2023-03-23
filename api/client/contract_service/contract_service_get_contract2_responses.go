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

// ContractServiceGetContract2Reader is a Reader for the ContractServiceGetContract2 structure.
type ContractServiceGetContract2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContractServiceGetContract2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContractServiceGetContract2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewContractServiceGetContract2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewContractServiceGetContract2OK creates a ContractServiceGetContract2OK with default headers values
func NewContractServiceGetContract2OK() *ContractServiceGetContract2OK {
	return &ContractServiceGetContract2OK{}
}

/*
ContractServiceGetContract2OK describes a response with status code 200, with default header values.

A successful response.
*/
type ContractServiceGetContract2OK struct {
	Payload *models.V1GetContractResponse
}

// IsSuccess returns true when this contract service get contract2 o k response has a 2xx status code
func (o *ContractServiceGetContract2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contract service get contract2 o k response has a 3xx status code
func (o *ContractServiceGetContract2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contract service get contract2 o k response has a 4xx status code
func (o *ContractServiceGetContract2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contract service get contract2 o k response has a 5xx status code
func (o *ContractServiceGetContract2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this contract service get contract2 o k response a status code equal to that given
func (o *ContractServiceGetContract2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contract service get contract2 o k response
func (o *ContractServiceGetContract2OK) Code() int {
	return 200
}

func (o *ContractServiceGetContract2OK) Error() string {
	return fmt.Sprintf("[GET /v1/contract/get][%d] contractServiceGetContract2OK  %+v", 200, o.Payload)
}

func (o *ContractServiceGetContract2OK) String() string {
	return fmt.Sprintf("[GET /v1/contract/get][%d] contractServiceGetContract2OK  %+v", 200, o.Payload)
}

func (o *ContractServiceGetContract2OK) GetPayload() *models.V1GetContractResponse {
	return o.Payload
}

func (o *ContractServiceGetContract2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetContractResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContractServiceGetContract2Default creates a ContractServiceGetContract2Default with default headers values
func NewContractServiceGetContract2Default(code int) *ContractServiceGetContract2Default {
	return &ContractServiceGetContract2Default{
		_statusCode: code,
	}
}

/*
ContractServiceGetContract2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ContractServiceGetContract2Default struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this contract service get contract2 default response has a 2xx status code
func (o *ContractServiceGetContract2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this contract service get contract2 default response has a 3xx status code
func (o *ContractServiceGetContract2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this contract service get contract2 default response has a 4xx status code
func (o *ContractServiceGetContract2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this contract service get contract2 default response has a 5xx status code
func (o *ContractServiceGetContract2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this contract service get contract2 default response a status code equal to that given
func (o *ContractServiceGetContract2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the contract service get contract2 default response
func (o *ContractServiceGetContract2Default) Code() int {
	return o._statusCode
}

func (o *ContractServiceGetContract2Default) Error() string {
	return fmt.Sprintf("[GET /v1/contract/get][%d] ContractService_GetContract2 default  %+v", o._statusCode, o.Payload)
}

func (o *ContractServiceGetContract2Default) String() string {
	return fmt.Sprintf("[GET /v1/contract/get][%d] ContractService_GetContract2 default  %+v", o._statusCode, o.Payload)
}

func (o *ContractServiceGetContract2Default) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ContractServiceGetContract2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
