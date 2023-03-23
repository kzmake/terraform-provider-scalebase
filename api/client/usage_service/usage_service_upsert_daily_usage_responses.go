// Code generated by go-swagger; DO NOT EDIT.

package usage_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

// UsageServiceUpsertDailyUsageReader is a Reader for the UsageServiceUpsertDailyUsage structure.
type UsageServiceUpsertDailyUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UsageServiceUpsertDailyUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUsageServiceUpsertDailyUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewUsageServiceUpsertDailyUsageDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUsageServiceUpsertDailyUsageOK creates a UsageServiceUpsertDailyUsageOK with default headers values
func NewUsageServiceUpsertDailyUsageOK() *UsageServiceUpsertDailyUsageOK {
	return &UsageServiceUpsertDailyUsageOK{}
}

/*
UsageServiceUpsertDailyUsageOK describes a response with status code 200, with default header values.

A successful response.
*/
type UsageServiceUpsertDailyUsageOK struct {
	Payload *models.V1UpsertDailyUsageResponse
}

// IsSuccess returns true when this usage service upsert daily usage o k response has a 2xx status code
func (o *UsageServiceUpsertDailyUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this usage service upsert daily usage o k response has a 3xx status code
func (o *UsageServiceUpsertDailyUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this usage service upsert daily usage o k response has a 4xx status code
func (o *UsageServiceUpsertDailyUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this usage service upsert daily usage o k response has a 5xx status code
func (o *UsageServiceUpsertDailyUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this usage service upsert daily usage o k response a status code equal to that given
func (o *UsageServiceUpsertDailyUsageOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the usage service upsert daily usage o k response
func (o *UsageServiceUpsertDailyUsageOK) Code() int {
	return 200
}

func (o *UsageServiceUpsertDailyUsageOK) Error() string {
	return fmt.Sprintf("[POST /v1/usage/daily/upsert][%d] usageServiceUpsertDailyUsageOK  %+v", 200, o.Payload)
}

func (o *UsageServiceUpsertDailyUsageOK) String() string {
	return fmt.Sprintf("[POST /v1/usage/daily/upsert][%d] usageServiceUpsertDailyUsageOK  %+v", 200, o.Payload)
}

func (o *UsageServiceUpsertDailyUsageOK) GetPayload() *models.V1UpsertDailyUsageResponse {
	return o.Payload
}

func (o *UsageServiceUpsertDailyUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1UpsertDailyUsageResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUsageServiceUpsertDailyUsageDefault creates a UsageServiceUpsertDailyUsageDefault with default headers values
func NewUsageServiceUpsertDailyUsageDefault(code int) *UsageServiceUpsertDailyUsageDefault {
	return &UsageServiceUpsertDailyUsageDefault{
		_statusCode: code,
	}
}

/*
UsageServiceUpsertDailyUsageDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type UsageServiceUpsertDailyUsageDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this usage service upsert daily usage default response has a 2xx status code
func (o *UsageServiceUpsertDailyUsageDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this usage service upsert daily usage default response has a 3xx status code
func (o *UsageServiceUpsertDailyUsageDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this usage service upsert daily usage default response has a 4xx status code
func (o *UsageServiceUpsertDailyUsageDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this usage service upsert daily usage default response has a 5xx status code
func (o *UsageServiceUpsertDailyUsageDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this usage service upsert daily usage default response a status code equal to that given
func (o *UsageServiceUpsertDailyUsageDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the usage service upsert daily usage default response
func (o *UsageServiceUpsertDailyUsageDefault) Code() int {
	return o._statusCode
}

func (o *UsageServiceUpsertDailyUsageDefault) Error() string {
	return fmt.Sprintf("[POST /v1/usage/daily/upsert][%d] UsageService_UpsertDailyUsage default  %+v", o._statusCode, o.Payload)
}

func (o *UsageServiceUpsertDailyUsageDefault) String() string {
	return fmt.Sprintf("[POST /v1/usage/daily/upsert][%d] UsageService_UpsertDailyUsage default  %+v", o._statusCode, o.Payload)
}

func (o *UsageServiceUpsertDailyUsageDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *UsageServiceUpsertDailyUsageDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}