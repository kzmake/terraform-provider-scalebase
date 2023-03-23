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

// BillingServiceListBillingsReader is a Reader for the BillingServiceListBillings structure.
type BillingServiceListBillingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *BillingServiceListBillingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewBillingServiceListBillingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewBillingServiceListBillingsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewBillingServiceListBillingsOK creates a BillingServiceListBillingsOK with default headers values
func NewBillingServiceListBillingsOK() *BillingServiceListBillingsOK {
	return &BillingServiceListBillingsOK{}
}

/*
BillingServiceListBillingsOK describes a response with status code 200, with default header values.

A successful response.
*/
type BillingServiceListBillingsOK struct {
	Payload *models.V1ListBillingsResponse
}

// IsSuccess returns true when this billing service list billings o k response has a 2xx status code
func (o *BillingServiceListBillingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this billing service list billings o k response has a 3xx status code
func (o *BillingServiceListBillingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this billing service list billings o k response has a 4xx status code
func (o *BillingServiceListBillingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this billing service list billings o k response has a 5xx status code
func (o *BillingServiceListBillingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this billing service list billings o k response a status code equal to that given
func (o *BillingServiceListBillingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the billing service list billings o k response
func (o *BillingServiceListBillingsOK) Code() int {
	return 200
}

func (o *BillingServiceListBillingsOK) Error() string {
	return fmt.Sprintf("[POST /v1/billing/list][%d] billingServiceListBillingsOK  %+v", 200, o.Payload)
}

func (o *BillingServiceListBillingsOK) String() string {
	return fmt.Sprintf("[POST /v1/billing/list][%d] billingServiceListBillingsOK  %+v", 200, o.Payload)
}

func (o *BillingServiceListBillingsOK) GetPayload() *models.V1ListBillingsResponse {
	return o.Payload
}

func (o *BillingServiceListBillingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ListBillingsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewBillingServiceListBillingsDefault creates a BillingServiceListBillingsDefault with default headers values
func NewBillingServiceListBillingsDefault(code int) *BillingServiceListBillingsDefault {
	return &BillingServiceListBillingsDefault{
		_statusCode: code,
	}
}

/*
BillingServiceListBillingsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type BillingServiceListBillingsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this billing service list billings default response has a 2xx status code
func (o *BillingServiceListBillingsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this billing service list billings default response has a 3xx status code
func (o *BillingServiceListBillingsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this billing service list billings default response has a 4xx status code
func (o *BillingServiceListBillingsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this billing service list billings default response has a 5xx status code
func (o *BillingServiceListBillingsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this billing service list billings default response a status code equal to that given
func (o *BillingServiceListBillingsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the billing service list billings default response
func (o *BillingServiceListBillingsDefault) Code() int {
	return o._statusCode
}

func (o *BillingServiceListBillingsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/billing/list][%d] BillingService_ListBillings default  %+v", o._statusCode, o.Payload)
}

func (o *BillingServiceListBillingsDefault) String() string {
	return fmt.Sprintf("[POST /v1/billing/list][%d] BillingService_ListBillings default  %+v", o._statusCode, o.Payload)
}

func (o *BillingServiceListBillingsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *BillingServiceListBillingsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
