// Code generated by go-swagger; DO NOT EDIT.

package customer_staff_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

// CustomerStaffServiceUpdateCustomerStaffReader is a Reader for the CustomerStaffServiceUpdateCustomerStaff structure.
type CustomerStaffServiceUpdateCustomerStaffReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerStaffServiceUpdateCustomerStaffReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerStaffServiceUpdateCustomerStaffOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerStaffServiceUpdateCustomerStaffDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerStaffServiceUpdateCustomerStaffOK creates a CustomerStaffServiceUpdateCustomerStaffOK with default headers values
func NewCustomerStaffServiceUpdateCustomerStaffOK() *CustomerStaffServiceUpdateCustomerStaffOK {
	return &CustomerStaffServiceUpdateCustomerStaffOK{}
}

/*
CustomerStaffServiceUpdateCustomerStaffOK describes a response with status code 200, with default header values.

A successful response.
*/
type CustomerStaffServiceUpdateCustomerStaffOK struct {
	Payload *models.V1UpdateCustomerStaffResponse
}

// IsSuccess returns true when this customer staff service update customer staff o k response has a 2xx status code
func (o *CustomerStaffServiceUpdateCustomerStaffOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer staff service update customer staff o k response has a 3xx status code
func (o *CustomerStaffServiceUpdateCustomerStaffOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer staff service update customer staff o k response has a 4xx status code
func (o *CustomerStaffServiceUpdateCustomerStaffOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer staff service update customer staff o k response has a 5xx status code
func (o *CustomerStaffServiceUpdateCustomerStaffOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer staff service update customer staff o k response a status code equal to that given
func (o *CustomerStaffServiceUpdateCustomerStaffOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer staff service update customer staff o k response
func (o *CustomerStaffServiceUpdateCustomerStaffOK) Code() int {
	return 200
}

func (o *CustomerStaffServiceUpdateCustomerStaffOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/customerstaff/update][%d] customerStaffServiceUpdateCustomerStaffOK  %+v", 200, o.Payload)
}

func (o *CustomerStaffServiceUpdateCustomerStaffOK) String() string {
	return fmt.Sprintf("[PATCH /v1/customerstaff/update][%d] customerStaffServiceUpdateCustomerStaffOK  %+v", 200, o.Payload)
}

func (o *CustomerStaffServiceUpdateCustomerStaffOK) GetPayload() *models.V1UpdateCustomerStaffResponse {
	return o.Payload
}

func (o *CustomerStaffServiceUpdateCustomerStaffOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1UpdateCustomerStaffResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerStaffServiceUpdateCustomerStaffDefault creates a CustomerStaffServiceUpdateCustomerStaffDefault with default headers values
func NewCustomerStaffServiceUpdateCustomerStaffDefault(code int) *CustomerStaffServiceUpdateCustomerStaffDefault {
	return &CustomerStaffServiceUpdateCustomerStaffDefault{
		_statusCode: code,
	}
}

/*
CustomerStaffServiceUpdateCustomerStaffDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CustomerStaffServiceUpdateCustomerStaffDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this customer staff service update customer staff default response has a 2xx status code
func (o *CustomerStaffServiceUpdateCustomerStaffDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer staff service update customer staff default response has a 3xx status code
func (o *CustomerStaffServiceUpdateCustomerStaffDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer staff service update customer staff default response has a 4xx status code
func (o *CustomerStaffServiceUpdateCustomerStaffDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer staff service update customer staff default response has a 5xx status code
func (o *CustomerStaffServiceUpdateCustomerStaffDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer staff service update customer staff default response a status code equal to that given
func (o *CustomerStaffServiceUpdateCustomerStaffDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer staff service update customer staff default response
func (o *CustomerStaffServiceUpdateCustomerStaffDefault) Code() int {
	return o._statusCode
}

func (o *CustomerStaffServiceUpdateCustomerStaffDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/customerstaff/update][%d] CustomerStaffService_UpdateCustomerStaff default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerStaffServiceUpdateCustomerStaffDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/customerstaff/update][%d] CustomerStaffService_UpdateCustomerStaff default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerStaffServiceUpdateCustomerStaffDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *CustomerStaffServiceUpdateCustomerStaffDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
