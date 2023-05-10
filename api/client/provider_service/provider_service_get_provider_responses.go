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

// ProviderServiceGetProviderReader is a Reader for the ProviderServiceGetProvider structure.
type ProviderServiceGetProviderReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ProviderServiceGetProviderReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewProviderServiceGetProviderOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewProviderServiceGetProviderDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewProviderServiceGetProviderOK creates a ProviderServiceGetProviderOK with default headers values
func NewProviderServiceGetProviderOK() *ProviderServiceGetProviderOK {
	return &ProviderServiceGetProviderOK{}
}

/*
ProviderServiceGetProviderOK describes a response with status code 200, with default header values.

A successful response.
*/
type ProviderServiceGetProviderOK struct {
	Payload *models.V1GetProviderResponse
}

// IsSuccess returns true when this provider service get provider o k response has a 2xx status code
func (o *ProviderServiceGetProviderOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this provider service get provider o k response has a 3xx status code
func (o *ProviderServiceGetProviderOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this provider service get provider o k response has a 4xx status code
func (o *ProviderServiceGetProviderOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this provider service get provider o k response has a 5xx status code
func (o *ProviderServiceGetProviderOK) IsServerError() bool {
	return false
}

// IsCode returns true when this provider service get provider o k response a status code equal to that given
func (o *ProviderServiceGetProviderOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the provider service get provider o k response
func (o *ProviderServiceGetProviderOK) Code() int {
	return 200
}

func (o *ProviderServiceGetProviderOK) Error() string {
	return fmt.Sprintf("[POST /v1/provider/get][%d] providerServiceGetProviderOK  %+v", 200, o.Payload)
}

func (o *ProviderServiceGetProviderOK) String() string {
	return fmt.Sprintf("[POST /v1/provider/get][%d] providerServiceGetProviderOK  %+v", 200, o.Payload)
}

func (o *ProviderServiceGetProviderOK) GetPayload() *models.V1GetProviderResponse {
	return o.Payload
}

func (o *ProviderServiceGetProviderOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetProviderResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewProviderServiceGetProviderDefault creates a ProviderServiceGetProviderDefault with default headers values
func NewProviderServiceGetProviderDefault(code int) *ProviderServiceGetProviderDefault {
	return &ProviderServiceGetProviderDefault{
		_statusCode: code,
	}
}

/*
ProviderServiceGetProviderDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ProviderServiceGetProviderDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this provider service get provider default response has a 2xx status code
func (o *ProviderServiceGetProviderDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this provider service get provider default response has a 3xx status code
func (o *ProviderServiceGetProviderDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this provider service get provider default response has a 4xx status code
func (o *ProviderServiceGetProviderDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this provider service get provider default response has a 5xx status code
func (o *ProviderServiceGetProviderDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this provider service get provider default response a status code equal to that given
func (o *ProviderServiceGetProviderDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the provider service get provider default response
func (o *ProviderServiceGetProviderDefault) Code() int {
	return o._statusCode
}

func (o *ProviderServiceGetProviderDefault) Error() string {
	return fmt.Sprintf("[POST /v1/provider/get][%d] ProviderService_GetProvider default  %+v", o._statusCode, o.Payload)
}

func (o *ProviderServiceGetProviderDefault) String() string {
	return fmt.Sprintf("[POST /v1/provider/get][%d] ProviderService_GetProvider default  %+v", o._statusCode, o.Payload)
}

func (o *ProviderServiceGetProviderDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ProviderServiceGetProviderDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
