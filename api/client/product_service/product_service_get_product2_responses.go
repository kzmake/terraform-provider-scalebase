// Code generated by go-swagger; DO NOT EDIT.

package product_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

// ProductServiceGetProduct2Reader is a Reader for the ProductServiceGetProduct2 structure.
type ProductServiceGetProduct2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProductServiceGetProduct2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProductServiceGetProduct2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProductServiceGetProduct2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProductServiceGetProduct2OK creates a ProductServiceGetProduct2OK with default headers values
func NewProductServiceGetProduct2OK() *ProductServiceGetProduct2OK {
	return &ProductServiceGetProduct2OK{}
}

/*
ProductServiceGetProduct2OK describes a response with status code 200, with default header values.

A successful response.
*/
type ProductServiceGetProduct2OK struct {
	Payload *models.V1GetProductResponse
}

// IsSuccess returns true when this product service get product2 o k response has a 2xx status code
func (o *ProductServiceGetProduct2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this product service get product2 o k response has a 3xx status code
func (o *ProductServiceGetProduct2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this product service get product2 o k response has a 4xx status code
func (o *ProductServiceGetProduct2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this product service get product2 o k response has a 5xx status code
func (o *ProductServiceGetProduct2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this product service get product2 o k response a status code equal to that given
func (o *ProductServiceGetProduct2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the product service get product2 o k response
func (o *ProductServiceGetProduct2OK) Code() int {
	return 200
}

func (o *ProductServiceGetProduct2OK) Error() string {
	return fmt.Sprintf("[GET /v1/product/get][%d] productServiceGetProduct2OK  %+v", 200, o.Payload)
}

func (o *ProductServiceGetProduct2OK) String() string {
	return fmt.Sprintf("[GET /v1/product/get][%d] productServiceGetProduct2OK  %+v", 200, o.Payload)
}

func (o *ProductServiceGetProduct2OK) GetPayload() *models.V1GetProductResponse {
	return o.Payload
}

func (o *ProductServiceGetProduct2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetProductResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProductServiceGetProduct2Default creates a ProductServiceGetProduct2Default with default headers values
func NewProductServiceGetProduct2Default(code int) *ProductServiceGetProduct2Default {
	return &ProductServiceGetProduct2Default{
		_statusCode: code,
	}
}

/*
ProductServiceGetProduct2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProductServiceGetProduct2Default struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this product service get product2 default response has a 2xx status code
func (o *ProductServiceGetProduct2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this product service get product2 default response has a 3xx status code
func (o *ProductServiceGetProduct2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this product service get product2 default response has a 4xx status code
func (o *ProductServiceGetProduct2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this product service get product2 default response has a 5xx status code
func (o *ProductServiceGetProduct2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this product service get product2 default response a status code equal to that given
func (o *ProductServiceGetProduct2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the product service get product2 default response
func (o *ProductServiceGetProduct2Default) Code() int {
	return o._statusCode
}

func (o *ProductServiceGetProduct2Default) Error() string {
	return fmt.Sprintf("[GET /v1/product/get][%d] ProductService_GetProduct2 default  %+v", o._statusCode, o.Payload)
}

func (o *ProductServiceGetProduct2Default) String() string {
	return fmt.Sprintf("[GET /v1/product/get][%d] ProductService_GetProduct2 default  %+v", o._statusCode, o.Payload)
}

func (o *ProductServiceGetProduct2Default) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ProductServiceGetProduct2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
