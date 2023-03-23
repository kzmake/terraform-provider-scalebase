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

// ContractServiceGetContractReader is a Reader for the ContractServiceGetContract structure.
type ContractServiceGetContractReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ContractServiceGetContractReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewContractServiceGetContractOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewContractServiceGetContractDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewContractServiceGetContractOK creates a ContractServiceGetContractOK with default headers values
func NewContractServiceGetContractOK() *ContractServiceGetContractOK {
	return &ContractServiceGetContractOK{}
}

/*
ContractServiceGetContractOK describes a response with status code 200, with default header values.

A successful response.
*/
type ContractServiceGetContractOK struct {
	Payload *models.V1GetContractResponse
}

// IsSuccess returns true when this contract service get contract o k response has a 2xx status code
func (o *ContractServiceGetContractOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this contract service get contract o k response has a 3xx status code
func (o *ContractServiceGetContractOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this contract service get contract o k response has a 4xx status code
func (o *ContractServiceGetContractOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this contract service get contract o k response has a 5xx status code
func (o *ContractServiceGetContractOK) IsServerError() bool {
	return false
}

// IsCode returns true when this contract service get contract o k response a status code equal to that given
func (o *ContractServiceGetContractOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the contract service get contract o k response
func (o *ContractServiceGetContractOK) Code() int {
	return 200
}

func (o *ContractServiceGetContractOK) Error() string {
	return fmt.Sprintf("[POST /v1/contract/get][%d] contractServiceGetContractOK  %+v", 200, o.Payload)
}

func (o *ContractServiceGetContractOK) String() string {
	return fmt.Sprintf("[POST /v1/contract/get][%d] contractServiceGetContractOK  %+v", 200, o.Payload)
}

func (o *ContractServiceGetContractOK) GetPayload() *models.V1GetContractResponse {
	return o.Payload
}

func (o *ContractServiceGetContractOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetContractResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewContractServiceGetContractDefault creates a ContractServiceGetContractDefault with default headers values
func NewContractServiceGetContractDefault(code int) *ContractServiceGetContractDefault {
	return &ContractServiceGetContractDefault{
		_statusCode: code,
	}
}

/*
ContractServiceGetContractDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ContractServiceGetContractDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this contract service get contract default response has a 2xx status code
func (o *ContractServiceGetContractDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this contract service get contract default response has a 3xx status code
func (o *ContractServiceGetContractDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this contract service get contract default response has a 4xx status code
func (o *ContractServiceGetContractDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this contract service get contract default response has a 5xx status code
func (o *ContractServiceGetContractDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this contract service get contract default response a status code equal to that given
func (o *ContractServiceGetContractDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the contract service get contract default response
func (o *ContractServiceGetContractDefault) Code() int {
	return o._statusCode
}

func (o *ContractServiceGetContractDefault) Error() string {
	return fmt.Sprintf("[POST /v1/contract/get][%d] ContractService_GetContract default  %+v", o._statusCode, o.Payload)
}

func (o *ContractServiceGetContractDefault) String() string {
	return fmt.Sprintf("[POST /v1/contract/get][%d] ContractService_GetContract default  %+v", o._statusCode, o.Payload)
}

func (o *ContractServiceGetContractDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ContractServiceGetContractDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
