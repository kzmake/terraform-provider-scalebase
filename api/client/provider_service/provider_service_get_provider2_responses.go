// Code generated by go-swagger; DO NOT EDIT.

package provider_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

// ProviderServiceGetProvider2Reader is a Reader for the ProviderServiceGetProvider2 structure.
type ProviderServiceGetProvider2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProviderServiceGetProvider2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProviderServiceGetProvider2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProviderServiceGetProvider2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProviderServiceGetProvider2OK creates a ProviderServiceGetProvider2OK with default headers values
func NewProviderServiceGetProvider2OK() *ProviderServiceGetProvider2OK {
	return &ProviderServiceGetProvider2OK{}
}

/*
ProviderServiceGetProvider2OK describes a response with status code 200, with default header values.

A successful response.
*/
type ProviderServiceGetProvider2OK struct {
	Payload *models.V1GetProviderResponse
}

// IsSuccess returns true when this provider service get provider2 o k response has a 2xx status code
func (o *ProviderServiceGetProvider2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this provider service get provider2 o k response has a 3xx status code
func (o *ProviderServiceGetProvider2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this provider service get provider2 o k response has a 4xx status code
func (o *ProviderServiceGetProvider2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this provider service get provider2 o k response has a 5xx status code
func (o *ProviderServiceGetProvider2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this provider service get provider2 o k response a status code equal to that given
func (o *ProviderServiceGetProvider2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the provider service get provider2 o k response
func (o *ProviderServiceGetProvider2OK) Code() int {
	return 200
}

func (o *ProviderServiceGetProvider2OK) Error() string {
	return fmt.Sprintf("[GET /v1/provider/get][%d] providerServiceGetProvider2OK  %+v", 200, o.Payload)
}

func (o *ProviderServiceGetProvider2OK) String() string {
	return fmt.Sprintf("[GET /v1/provider/get][%d] providerServiceGetProvider2OK  %+v", 200, o.Payload)
}

func (o *ProviderServiceGetProvider2OK) GetPayload() *models.V1GetProviderResponse {
	return o.Payload
}

func (o *ProviderServiceGetProvider2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProviderServiceGetProvider2Default creates a ProviderServiceGetProvider2Default with default headers values
func NewProviderServiceGetProvider2Default(code int) *ProviderServiceGetProvider2Default {
	return &ProviderServiceGetProvider2Default{
		_statusCode: code,
	}
}

/*
ProviderServiceGetProvider2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProviderServiceGetProvider2Default struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this provider service get provider2 default response has a 2xx status code
func (o *ProviderServiceGetProvider2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this provider service get provider2 default response has a 3xx status code
func (o *ProviderServiceGetProvider2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this provider service get provider2 default response has a 4xx status code
func (o *ProviderServiceGetProvider2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this provider service get provider2 default response has a 5xx status code
func (o *ProviderServiceGetProvider2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this provider service get provider2 default response a status code equal to that given
func (o *ProviderServiceGetProvider2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the provider service get provider2 default response
func (o *ProviderServiceGetProvider2Default) Code() int {
	return o._statusCode
}

func (o *ProviderServiceGetProvider2Default) Error() string {
	return fmt.Sprintf("[GET /v1/provider/get][%d] ProviderService_GetProvider2 default  %+v", o._statusCode, o.Payload)
}

func (o *ProviderServiceGetProvider2Default) String() string {
	return fmt.Sprintf("[GET /v1/provider/get][%d] ProviderService_GetProvider2 default  %+v", o._statusCode, o.Payload)
}

func (o *ProviderServiceGetProvider2Default) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ProviderServiceGetProvider2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
