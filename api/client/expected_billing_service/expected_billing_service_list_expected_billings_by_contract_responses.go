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

// ExpectedBillingServiceListExpectedBillingsByContractReader is a Reader for the ExpectedBillingServiceListExpectedBillingsByContract structure.
type ExpectedBillingServiceListExpectedBillingsByContractReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExpectedBillingServiceListExpectedBillingsByContractReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewExpectedBillingServiceListExpectedBillingsByContractOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewExpectedBillingServiceListExpectedBillingsByContractDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExpectedBillingServiceListExpectedBillingsByContractOK creates a ExpectedBillingServiceListExpectedBillingsByContractOK with default headers values
func NewExpectedBillingServiceListExpectedBillingsByContractOK() *ExpectedBillingServiceListExpectedBillingsByContractOK {
	return &ExpectedBillingServiceListExpectedBillingsByContractOK{}
}

/*
ExpectedBillingServiceListExpectedBillingsByContractOK describes a response with status code 200, with default header values.

A successful response.
*/
type ExpectedBillingServiceListExpectedBillingsByContractOK struct {
	Payload *models.V1ListExpectedBillingsByContractResponse
}

// IsSuccess returns true when this expected billing service list expected billings by contract o k response has a 2xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContractOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this expected billing service list expected billings by contract o k response has a 3xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContractOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this expected billing service list expected billings by contract o k response has a 4xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContractOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this expected billing service list expected billings by contract o k response has a 5xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContractOK) IsServerError() bool {
	return false
}

// IsCode returns true when this expected billing service list expected billings by contract o k response a status code equal to that given
func (o *ExpectedBillingServiceListExpectedBillingsByContractOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the expected billing service list expected billings by contract o k response
func (o *ExpectedBillingServiceListExpectedBillingsByContractOK) Code() int {
	return 200
}

func (o *ExpectedBillingServiceListExpectedBillingsByContractOK) Error() string {
	return fmt.Sprintf("[POST /v1/expectedbilling/listbycontract][%d] expectedBillingServiceListExpectedBillingsByContractOK  %+v", 200, o.Payload)
}

func (o *ExpectedBillingServiceListExpectedBillingsByContractOK) String() string {
	return fmt.Sprintf("[POST /v1/expectedbilling/listbycontract][%d] expectedBillingServiceListExpectedBillingsByContractOK  %+v", 200, o.Payload)
}

func (o *ExpectedBillingServiceListExpectedBillingsByContractOK) GetPayload() *models.V1ListExpectedBillingsByContractResponse {
	return o.Payload
}

func (o *ExpectedBillingServiceListExpectedBillingsByContractOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ListExpectedBillingsByContractResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExpectedBillingServiceListExpectedBillingsByContractDefault creates a ExpectedBillingServiceListExpectedBillingsByContractDefault with default headers values
func NewExpectedBillingServiceListExpectedBillingsByContractDefault(code int) *ExpectedBillingServiceListExpectedBillingsByContractDefault {
	return &ExpectedBillingServiceListExpectedBillingsByContractDefault{
		_statusCode: code,
	}
}

/*
ExpectedBillingServiceListExpectedBillingsByContractDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ExpectedBillingServiceListExpectedBillingsByContractDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this expected billing service list expected billings by contract default response has a 2xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContractDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this expected billing service list expected billings by contract default response has a 3xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContractDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this expected billing service list expected billings by contract default response has a 4xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContractDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this expected billing service list expected billings by contract default response has a 5xx status code
func (o *ExpectedBillingServiceListExpectedBillingsByContractDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this expected billing service list expected billings by contract default response a status code equal to that given
func (o *ExpectedBillingServiceListExpectedBillingsByContractDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the expected billing service list expected billings by contract default response
func (o *ExpectedBillingServiceListExpectedBillingsByContractDefault) Code() int {
	return o._statusCode
}

func (o *ExpectedBillingServiceListExpectedBillingsByContractDefault) Error() string {
	return fmt.Sprintf("[POST /v1/expectedbilling/listbycontract][%d] ExpectedBillingService_ListExpectedBillingsByContract default  %+v", o._statusCode, o.Payload)
}

func (o *ExpectedBillingServiceListExpectedBillingsByContractDefault) String() string {
	return fmt.Sprintf("[POST /v1/expectedbilling/listbycontract][%d] ExpectedBillingService_ListExpectedBillingsByContract default  %+v", o._statusCode, o.Payload)
}

func (o *ExpectedBillingServiceListExpectedBillingsByContractDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ExpectedBillingServiceListExpectedBillingsByContractDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
