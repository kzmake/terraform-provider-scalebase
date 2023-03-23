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

// BillingServiceListBillings2Reader is a Reader for the BillingServiceListBillings2 structure.
type BillingServiceListBillings2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingServiceListBillings2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingServiceListBillings2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBillingServiceListBillings2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBillingServiceListBillings2OK creates a BillingServiceListBillings2OK with default headers values
func NewBillingServiceListBillings2OK() *BillingServiceListBillings2OK {
	return &BillingServiceListBillings2OK{}
}

/*
BillingServiceListBillings2OK describes a response with status code 200, with default header values.

A successful response.
*/
type BillingServiceListBillings2OK struct {
	Payload *models.V1ListBillingsResponse
}

// IsSuccess returns true when this billing service list billings2 o k response has a 2xx status code
func (o *BillingServiceListBillings2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this billing service list billings2 o k response has a 3xx status code
func (o *BillingServiceListBillings2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing service list billings2 o k response has a 4xx status code
func (o *BillingServiceListBillings2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this billing service list billings2 o k response has a 5xx status code
func (o *BillingServiceListBillings2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this billing service list billings2 o k response a status code equal to that given
func (o *BillingServiceListBillings2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the billing service list billings2 o k response
func (o *BillingServiceListBillings2OK) Code() int {
	return 200
}

func (o *BillingServiceListBillings2OK) Error() string {
	return fmt.Sprintf("[GET /v1/billing/list][%d] billingServiceListBillings2OK  %+v", 200, o.Payload)
}

func (o *BillingServiceListBillings2OK) String() string {
	return fmt.Sprintf("[GET /v1/billing/list][%d] billingServiceListBillings2OK  %+v", 200, o.Payload)
}

func (o *BillingServiceListBillings2OK) GetPayload() *models.V1ListBillingsResponse {
	return o.Payload
}

func (o *BillingServiceListBillings2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ListBillingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingServiceListBillings2Default creates a BillingServiceListBillings2Default with default headers values
func NewBillingServiceListBillings2Default(code int) *BillingServiceListBillings2Default {
	return &BillingServiceListBillings2Default{
		_statusCode: code,
	}
}

/*
BillingServiceListBillings2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BillingServiceListBillings2Default struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this billing service list billings2 default response has a 2xx status code
func (o *BillingServiceListBillings2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this billing service list billings2 default response has a 3xx status code
func (o *BillingServiceListBillings2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this billing service list billings2 default response has a 4xx status code
func (o *BillingServiceListBillings2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this billing service list billings2 default response has a 5xx status code
func (o *BillingServiceListBillings2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this billing service list billings2 default response a status code equal to that given
func (o *BillingServiceListBillings2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the billing service list billings2 default response
func (o *BillingServiceListBillings2Default) Code() int {
	return o._statusCode
}

func (o *BillingServiceListBillings2Default) Error() string {
	return fmt.Sprintf("[GET /v1/billing/list][%d] BillingService_ListBillings2 default  %+v", o._statusCode, o.Payload)
}

func (o *BillingServiceListBillings2Default) String() string {
	return fmt.Sprintf("[GET /v1/billing/list][%d] BillingService_ListBillings2 default  %+v", o._statusCode, o.Payload)
}

func (o *BillingServiceListBillings2Default) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *BillingServiceListBillings2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
