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

// BillingServiceGetBillingReader is a Reader for the BillingServiceGetBilling structure.
type BillingServiceGetBillingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingServiceGetBillingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingServiceGetBillingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBillingServiceGetBillingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBillingServiceGetBillingOK creates a BillingServiceGetBillingOK with default headers values
func NewBillingServiceGetBillingOK() *BillingServiceGetBillingOK {
	return &BillingServiceGetBillingOK{}
}

/*
BillingServiceGetBillingOK describes a response with status code 200, with default header values.

A successful response.
*/
type BillingServiceGetBillingOK struct {
	Payload *models.V1GetBillingResponse
}

// IsSuccess returns true when this billing service get billing o k response has a 2xx status code
func (o *BillingServiceGetBillingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this billing service get billing o k response has a 3xx status code
func (o *BillingServiceGetBillingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing service get billing o k response has a 4xx status code
func (o *BillingServiceGetBillingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this billing service get billing o k response has a 5xx status code
func (o *BillingServiceGetBillingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this billing service get billing o k response a status code equal to that given
func (o *BillingServiceGetBillingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the billing service get billing o k response
func (o *BillingServiceGetBillingOK) Code() int {
	return 200
}

func (o *BillingServiceGetBillingOK) Error() string {
	return fmt.Sprintf("[POST /v1/billing/get][%d] billingServiceGetBillingOK  %+v", 200, o.Payload)
}

func (o *BillingServiceGetBillingOK) String() string {
	return fmt.Sprintf("[POST /v1/billing/get][%d] billingServiceGetBillingOK  %+v", 200, o.Payload)
}

func (o *BillingServiceGetBillingOK) GetPayload() *models.V1GetBillingResponse {
	return o.Payload
}

func (o *BillingServiceGetBillingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetBillingResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingServiceGetBillingDefault creates a BillingServiceGetBillingDefault with default headers values
func NewBillingServiceGetBillingDefault(code int) *BillingServiceGetBillingDefault {
	return &BillingServiceGetBillingDefault{
		_statusCode: code,
	}
}

/*
BillingServiceGetBillingDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BillingServiceGetBillingDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this billing service get billing default response has a 2xx status code
func (o *BillingServiceGetBillingDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this billing service get billing default response has a 3xx status code
func (o *BillingServiceGetBillingDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this billing service get billing default response has a 4xx status code
func (o *BillingServiceGetBillingDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this billing service get billing default response has a 5xx status code
func (o *BillingServiceGetBillingDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this billing service get billing default response a status code equal to that given
func (o *BillingServiceGetBillingDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the billing service get billing default response
func (o *BillingServiceGetBillingDefault) Code() int {
	return o._statusCode
}

func (o *BillingServiceGetBillingDefault) Error() string {
	return fmt.Sprintf("[POST /v1/billing/get][%d] BillingService_GetBilling default  %+v", o._statusCode, o.Payload)
}

func (o *BillingServiceGetBillingDefault) String() string {
	return fmt.Sprintf("[POST /v1/billing/get][%d] BillingService_GetBilling default  %+v", o._statusCode, o.Payload)
}

func (o *BillingServiceGetBillingDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *BillingServiceGetBillingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
