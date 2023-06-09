// Code generated by go-swagger; DO NOT EDIT.

package expected_billing_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

// ExpectedBillingServiceListExpectedBillingsByContract2Reader is a Reader for the ExpectedBillingServiceListExpectedBillingsByContract2 structure.
type ExpectedBillingServiceListExpectedBillingsByContract2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExpectedBillingServiceListExpectedBillingsByContract2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExpectedBillingServiceListExpectedBillingsByContract2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExpectedBillingServiceListExpectedBillingsByContract2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExpectedBillingServiceListExpectedBillingsByContract2OK creates a ExpectedBillingServiceListExpectedBillingsByContract2OK with default headers values
func NewExpectedBillingServiceListExpectedBillingsByContract2OK() *ExpectedBillingServiceListExpectedBillingsByContract2OK {
	return &ExpectedBillingServiceListExpectedBillingsByContract2OK{}
}

/*
ExpectedBillingServiceListExpectedBillingsByContract2OK describes a response with status code 200, with default header values.

A successful response.
*/
type ExpectedBillingServiceListExpectedBillingsByContract2OK struct {
	Payload *models.V1ListExpectedBillingsByContractResponse
}

// IsSuccess returns true when this expected billing service list expected billings by contract2 o k response has a 2xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContract2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this expected billing service list expected billings by contract2 o k response has a 3xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContract2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this expected billing service list expected billings by contract2 o k response has a 4xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContract2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this expected billing service list expected billings by contract2 o k response has a 5xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContract2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this expected billing service list expected billings by contract2 o k response a status code equal to that given
func (o *ExpectedBillingServiceListExpectedBillingsByContract2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the expected billing service list expected billings by contract2 o k response
func (o *ExpectedBillingServiceListExpectedBillingsByContract2OK) Code() int {
	return 200
}

func (o *ExpectedBillingServiceListExpectedBillingsByContract2OK) Error() string {
	return fmt.Sprintf("[GET /v1/expectedbilling/listbycontract][%d] expectedBillingServiceListExpectedBillingsByContract2OK  %+v", 200, o.Payload)
}

func (o *ExpectedBillingServiceListExpectedBillingsByContract2OK) String() string {
	return fmt.Sprintf("[GET /v1/expectedbilling/listbycontract][%d] expectedBillingServiceListExpectedBillingsByContract2OK  %+v", 200, o.Payload)
}

func (o *ExpectedBillingServiceListExpectedBillingsByContract2OK) GetPayload() *models.V1ListExpectedBillingsByContractResponse {
	return o.Payload
}

func (o *ExpectedBillingServiceListExpectedBillingsByContract2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ListExpectedBillingsByContractResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExpectedBillingServiceListExpectedBillingsByContract2Default creates a ExpectedBillingServiceListExpectedBillingsByContract2Default with default headers values
func NewExpectedBillingServiceListExpectedBillingsByContract2Default(code int) *ExpectedBillingServiceListExpectedBillingsByContract2Default {
	return &ExpectedBillingServiceListExpectedBillingsByContract2Default{
		_statusCode: code,
	}
}

/*
ExpectedBillingServiceListExpectedBillingsByContract2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ExpectedBillingServiceListExpectedBillingsByContract2Default struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this expected billing service list expected billings by contract2 default response has a 2xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContract2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this expected billing service list expected billings by contract2 default response has a 3xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContract2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this expected billing service list expected billings by contract2 default response has a 4xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContract2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this expected billing service list expected billings by contract2 default response has a 5xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContract2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this expected billing service list expected billings by contract2 default response a status code equal to that given
func (o *ExpectedBillingServiceListExpectedBillingsByContract2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the expected billing service list expected billings by contract2 default response
func (o *ExpectedBillingServiceListExpectedBillingsByContract2Default) Code() int {
	return o._statusCode
}

func (o *ExpectedBillingServiceListExpectedBillingsByContract2Default) Error() string {
	return fmt.Sprintf("[GET /v1/expectedbilling/listbycontract][%d] ExpectedBillingService_ListExpectedBillingsByContract2 default  %+v", o._statusCode, o.Payload)
}

func (o *ExpectedBillingServiceListExpectedBillingsByContract2Default) String() string {
	return fmt.Sprintf("[GET /v1/expectedbilling/listbycontract][%d] ExpectedBillingService_ListExpectedBillingsByContract2 default  %+v", o._statusCode, o.Payload)
}

func (o *ExpectedBillingServiceListExpectedBillingsByContract2Default) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ExpectedBillingServiceListExpectedBillingsByContract2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
