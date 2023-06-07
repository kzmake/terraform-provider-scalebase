// Code generated by go-swagger; DO NOT EDIT.

package custom_field_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

// CustomFieldServiceGetResource2Reader is a Reader for the CustomFieldServiceGetResource2 structure.
type CustomFieldServiceGetResource2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomFieldServiceGetResource2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomFieldServiceGetResource2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomFieldServiceGetResource2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomFieldServiceGetResource2OK creates a CustomFieldServiceGetResource2OK with default headers values
func NewCustomFieldServiceGetResource2OK() *CustomFieldServiceGetResource2OK {
	return &CustomFieldServiceGetResource2OK{}
}

/*
CustomFieldServiceGetResource2OK describes a response with status code 200, with default header values.

A successful response.
*/
type CustomFieldServiceGetResource2OK struct {
	Payload *models.V1GetResourceResponse
}

// IsSuccess returns true when this custom field service get resource2 o k response has a 2xx status code
func (o *CustomFieldServiceGetResource2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this custom field service get resource2 o k response has a 3xx status code
func (o *CustomFieldServiceGetResource2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom field service get resource2 o k response has a 4xx status code
func (o *CustomFieldServiceGetResource2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this custom field service get resource2 o k response has a 5xx status code
func (o *CustomFieldServiceGetResource2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this custom field service get resource2 o k response a status code equal to that given
func (o *CustomFieldServiceGetResource2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the custom field service get resource2 o k response
func (o *CustomFieldServiceGetResource2OK) Code() int {
	return 200
}

func (o *CustomFieldServiceGetResource2OK) Error() string {
	return fmt.Sprintf("[GET /v1/customfield/resource/get][%d] customFieldServiceGetResource2OK  %+v", 200, o.Payload)
}

func (o *CustomFieldServiceGetResource2OK) String() string {
	return fmt.Sprintf("[GET /v1/customfield/resource/get][%d] customFieldServiceGetResource2OK  %+v", 200, o.Payload)
}

func (o *CustomFieldServiceGetResource2OK) GetPayload() *models.V1GetResourceResponse {
	return o.Payload
}

func (o *CustomFieldServiceGetResource2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetResourceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomFieldServiceGetResource2Default creates a CustomFieldServiceGetResource2Default with default headers values
func NewCustomFieldServiceGetResource2Default(code int) *CustomFieldServiceGetResource2Default {
	return &CustomFieldServiceGetResource2Default{
		_statusCode: code,
	}
}

/*
CustomFieldServiceGetResource2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CustomFieldServiceGetResource2Default struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this custom field service get resource2 default response has a 2xx status code
func (o *CustomFieldServiceGetResource2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this custom field service get resource2 default response has a 3xx status code
func (o *CustomFieldServiceGetResource2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this custom field service get resource2 default response has a 4xx status code
func (o *CustomFieldServiceGetResource2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this custom field service get resource2 default response has a 5xx status code
func (o *CustomFieldServiceGetResource2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this custom field service get resource2 default response a status code equal to that given
func (o *CustomFieldServiceGetResource2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the custom field service get resource2 default response
func (o *CustomFieldServiceGetResource2Default) Code() int {
	return o._statusCode
}

func (o *CustomFieldServiceGetResource2Default) Error() string {
	return fmt.Sprintf("[GET /v1/customfield/resource/get][%d] CustomFieldService_GetResource2 default  %+v", o._statusCode, o.Payload)
}

func (o *CustomFieldServiceGetResource2Default) String() string {
	return fmt.Sprintf("[GET /v1/customfield/resource/get][%d] CustomFieldService_GetResource2 default  %+v", o._statusCode, o.Payload)
}

func (o *CustomFieldServiceGetResource2Default) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *CustomFieldServiceGetResource2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
