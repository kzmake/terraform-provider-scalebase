// Code generated by go-swagger; DO NOT EDIT.

package billing_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

// BillingServiceListBillingsByContractReader is a Reader for the BillingServiceListBillingsByContract structure.
type BillingServiceListBillingsByContractReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingServiceListBillingsByContractReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingServiceListBillingsByContractOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBillingServiceListBillingsByContractDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBillingServiceListBillingsByContractOK creates a BillingServiceListBillingsByContractOK with default headers values
func NewBillingServiceListBillingsByContractOK() *BillingServiceListBillingsByContractOK {
	return &BillingServiceListBillingsByContractOK{}
}

/*
BillingServiceListBillingsByContractOK describes a response with status code 200, with default header values.

A successful response.
*/
type BillingServiceListBillingsByContractOK struct {
	Payload *models.V1ListBillingsByContractResponse
}

// IsSuccess returns true when this billing service list billings by contract o k response has a 2xx status code
func (o *BillingServiceListBillingsByContractOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this billing service list billings by contract o k response has a 3xx status code
func (o *BillingServiceListBillingsByContractOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing service list billings by contract o k response has a 4xx status code
func (o *BillingServiceListBillingsByContractOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this billing service list billings by contract o k response has a 5xx status code
func (o *BillingServiceListBillingsByContractOK) IsServerError() bool {
	return false
}

// IsCode returns true when this billing service list billings by contract o k response a status code equal to that given
func (o *BillingServiceListBillingsByContractOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the billing service list billings by contract o k response
func (o *BillingServiceListBillingsByContractOK) Code() int {
	return 200
}

func (o *BillingServiceListBillingsByContractOK) Error() string {
	return fmt.Sprintf("[POST /v1/billing/listbycontract][%d] billingServiceListBillingsByContractOK  %+v", 200, o.Payload)
}

func (o *BillingServiceListBillingsByContractOK) String() string {
	return fmt.Sprintf("[POST /v1/billing/listbycontract][%d] billingServiceListBillingsByContractOK  %+v", 200, o.Payload)
}

func (o *BillingServiceListBillingsByContractOK) GetPayload() *models.V1ListBillingsByContractResponse {
	return o.Payload
}

func (o *BillingServiceListBillingsByContractOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ListBillingsByContractResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingServiceListBillingsByContractDefault creates a BillingServiceListBillingsByContractDefault with default headers values
func NewBillingServiceListBillingsByContractDefault(code int) *BillingServiceListBillingsByContractDefault {
	return &BillingServiceListBillingsByContractDefault{
		_statusCode: code,
	}
}

/*
BillingServiceListBillingsByContractDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BillingServiceListBillingsByContractDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this billing service list billings by contract default response has a 2xx status code
func (o *BillingServiceListBillingsByContractDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this billing service list billings by contract default response has a 3xx status code
func (o *BillingServiceListBillingsByContractDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this billing service list billings by contract default response has a 4xx status code
func (o *BillingServiceListBillingsByContractDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this billing service list billings by contract default response has a 5xx status code
func (o *BillingServiceListBillingsByContractDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this billing service list billings by contract default response a status code equal to that given
func (o *BillingServiceListBillingsByContractDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the billing service list billings by contract default response
func (o *BillingServiceListBillingsByContractDefault) Code() int {
	return o._statusCode
}

func (o *BillingServiceListBillingsByContractDefault) Error() string {
	return fmt.Sprintf("[POST /v1/billing/listbycontract][%d] BillingService_ListBillingsByContract default  %+v", o._statusCode, o.Payload)
}

func (o *BillingServiceListBillingsByContractDefault) String() string {
	return fmt.Sprintf("[POST /v1/billing/listbycontract][%d] BillingService_ListBillingsByContract default  %+v", o._statusCode, o.Payload)
}

func (o *BillingServiceListBillingsByContractDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *BillingServiceListBillingsByContractDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
