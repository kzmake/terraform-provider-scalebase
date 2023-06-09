// Code generated by go-swagger; DO NOT EDIT.

package customer_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

// CustomerServiceUpdateCustomerReader is a Reader for the CustomerServiceUpdateCustomer structure.
type CustomerServiceUpdateCustomerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomerServiceUpdateCustomerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomerServiceUpdateCustomerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomerServiceUpdateCustomerDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomerServiceUpdateCustomerOK creates a CustomerServiceUpdateCustomerOK with default headers values
func NewCustomerServiceUpdateCustomerOK() *CustomerServiceUpdateCustomerOK {
	return &CustomerServiceUpdateCustomerOK{}
}

/*
CustomerServiceUpdateCustomerOK describes a response with status code 200, with default header values.

A successful response.
*/
type CustomerServiceUpdateCustomerOK struct {
	Payload *models.V1UpdateCustomerResponse
}

// IsSuccess returns true when this customer service update customer o k response has a 2xx status code
func (o *CustomerServiceUpdateCustomerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this customer service update customer o k response has a 3xx status code
func (o *CustomerServiceUpdateCustomerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this customer service update customer o k response has a 4xx status code
func (o *CustomerServiceUpdateCustomerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this customer service update customer o k response has a 5xx status code
func (o *CustomerServiceUpdateCustomerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this customer service update customer o k response a status code equal to that given
func (o *CustomerServiceUpdateCustomerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the customer service update customer o k response
func (o *CustomerServiceUpdateCustomerOK) Code() int {
	return 200
}

func (o *CustomerServiceUpdateCustomerOK) Error() string {
	return fmt.Sprintf("[PATCH /v1/customer/update][%d] customerServiceUpdateCustomerOK  %+v", 200, o.Payload)
}

func (o *CustomerServiceUpdateCustomerOK) String() string {
	return fmt.Sprintf("[PATCH /v1/customer/update][%d] customerServiceUpdateCustomerOK  %+v", 200, o.Payload)
}

func (o *CustomerServiceUpdateCustomerOK) GetPayload() *models.V1UpdateCustomerResponse {
	return o.Payload
}

func (o *CustomerServiceUpdateCustomerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1UpdateCustomerResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomerServiceUpdateCustomerDefault creates a CustomerServiceUpdateCustomerDefault with default headers values
func NewCustomerServiceUpdateCustomerDefault(code int) *CustomerServiceUpdateCustomerDefault {
	return &CustomerServiceUpdateCustomerDefault{
		_statusCode: code,
	}
}

/*
CustomerServiceUpdateCustomerDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CustomerServiceUpdateCustomerDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this customer service update customer default response has a 2xx status code
func (o *CustomerServiceUpdateCustomerDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this customer service update customer default response has a 3xx status code
func (o *CustomerServiceUpdateCustomerDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this customer service update customer default response has a 4xx status code
func (o *CustomerServiceUpdateCustomerDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this customer service update customer default response has a 5xx status code
func (o *CustomerServiceUpdateCustomerDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this customer service update customer default response a status code equal to that given
func (o *CustomerServiceUpdateCustomerDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the customer service update customer default response
func (o *CustomerServiceUpdateCustomerDefault) Code() int {
	return o._statusCode
}

func (o *CustomerServiceUpdateCustomerDefault) Error() string {
	return fmt.Sprintf("[PATCH /v1/customer/update][%d] CustomerService_UpdateCustomer default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerServiceUpdateCustomerDefault) String() string {
	return fmt.Sprintf("[PATCH /v1/customer/update][%d] CustomerService_UpdateCustomer default  %+v", o._statusCode, o.Payload)
}

func (o *CustomerServiceUpdateCustomerDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *CustomerServiceUpdateCustomerDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
