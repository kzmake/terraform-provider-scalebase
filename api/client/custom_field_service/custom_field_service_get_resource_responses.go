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

// CustomFieldServiceGetResourceReader is a Reader for the CustomFieldServiceGetResource structure.
type CustomFieldServiceGetResourceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CustomFieldServiceGetResourceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCustomFieldServiceGetResourceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCustomFieldServiceGetResourceDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCustomFieldServiceGetResourceOK creates a CustomFieldServiceGetResourceOK with default headers values
func NewCustomFieldServiceGetResourceOK() *CustomFieldServiceGetResourceOK {
	return &CustomFieldServiceGetResourceOK{}
}

/*
CustomFieldServiceGetResourceOK describes a response with status code 200, with default header values.

A successful response.
*/
type CustomFieldServiceGetResourceOK struct {
	Payload *models.V1GetResourceResponse
}

// IsSuccess returns true when this custom field service get resource o k response has a 2xx status code
func (o *CustomFieldServiceGetResourceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this custom field service get resource o k response has a 3xx status code
func (o *CustomFieldServiceGetResourceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this custom field service get resource o k response has a 4xx status code
func (o *CustomFieldServiceGetResourceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this custom field service get resource o k response has a 5xx status code
func (o *CustomFieldServiceGetResourceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this custom field service get resource o k response a status code equal to that given
func (o *CustomFieldServiceGetResourceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the custom field service get resource o k response
func (o *CustomFieldServiceGetResourceOK) Code() int {
	return 200
}

func (o *CustomFieldServiceGetResourceOK) Error() string {
	return fmt.Sprintf("[POST /v1/customfield/resource/get][%d] customFieldServiceGetResourceOK  %+v", 200, o.Payload)
}

func (o *CustomFieldServiceGetResourceOK) String() string {
	return fmt.Sprintf("[POST /v1/customfield/resource/get][%d] customFieldServiceGetResourceOK  %+v", 200, o.Payload)
}

func (o *CustomFieldServiceGetResourceOK) GetPayload() *models.V1GetResourceResponse {
	return o.Payload
}

func (o *CustomFieldServiceGetResourceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1GetResourceResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCustomFieldServiceGetResourceDefault creates a CustomFieldServiceGetResourceDefault with default headers values
func NewCustomFieldServiceGetResourceDefault(code int) *CustomFieldServiceGetResourceDefault {
	return &CustomFieldServiceGetResourceDefault{
		_statusCode: code,
	}
}

/*
CustomFieldServiceGetResourceDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CustomFieldServiceGetResourceDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this custom field service get resource default response has a 2xx status code
func (o *CustomFieldServiceGetResourceDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this custom field service get resource default response has a 3xx status code
func (o *CustomFieldServiceGetResourceDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this custom field service get resource default response has a 4xx status code
func (o *CustomFieldServiceGetResourceDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this custom field service get resource default response has a 5xx status code
func (o *CustomFieldServiceGetResourceDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this custom field service get resource default response a status code equal to that given
func (o *CustomFieldServiceGetResourceDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the custom field service get resource default response
func (o *CustomFieldServiceGetResourceDefault) Code() int {
	return o._statusCode
}

func (o *CustomFieldServiceGetResourceDefault) Error() string {
	return fmt.Sprintf("[POST /v1/customfield/resource/get][%d] CustomFieldService_GetResource default  %+v", o._statusCode, o.Payload)
}

func (o *CustomFieldServiceGetResourceDefault) String() string {
	return fmt.Sprintf("[POST /v1/customfield/resource/get][%d] CustomFieldService_GetResource default  %+v", o._statusCode, o.Payload)
}

func (o *CustomFieldServiceGetResourceDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *CustomFieldServiceGetResourceDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
