// Code generated by go-swagger; DO NOT EDIT.

package zz_health_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

// ZzHealthServiceCheck2Reader is a Reader for the ZzHealthServiceCheck2 structure.
type ZzHealthServiceCheck2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ZzHealthServiceCheck2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewZzHealthServiceCheck2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewZzHealthServiceCheck2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewZzHealthServiceCheck2OK creates a ZzHealthServiceCheck2OK with default headers values
func NewZzHealthServiceCheck2OK() *ZzHealthServiceCheck2OK {
	return &ZzHealthServiceCheck2OK{}
}

/*
ZzHealthServiceCheck2OK describes a response with status code 200, with default header values.

A successful response.
*/
type ZzHealthServiceCheck2OK struct {
	Payload *models.V1CheckResponse
}

// IsSuccess returns true when this zz health service check2 o k response has a 2xx status code
func (o *ZzHealthServiceCheck2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this zz health service check2 o k response has a 3xx status code
func (o *ZzHealthServiceCheck2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this zz health service check2 o k response has a 4xx status code
func (o *ZzHealthServiceCheck2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this zz health service check2 o k response has a 5xx status code
func (o *ZzHealthServiceCheck2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this zz health service check2 o k response a status code equal to that given
func (o *ZzHealthServiceCheck2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the zz health service check2 o k response
func (o *ZzHealthServiceCheck2OK) Code() int {
	return 200
}

func (o *ZzHealthServiceCheck2OK) Error() string {
	return fmt.Sprintf("[GET /hc][%d] zzHealthServiceCheck2OK  %+v", 200, o.Payload)
}

func (o *ZzHealthServiceCheck2OK) String() string {
	return fmt.Sprintf("[GET /hc][%d] zzHealthServiceCheck2OK  %+v", 200, o.Payload)
}

func (o *ZzHealthServiceCheck2OK) GetPayload() *models.V1CheckResponse {
	return o.Payload
}

func (o *ZzHealthServiceCheck2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1CheckResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewZzHealthServiceCheck2Default creates a ZzHealthServiceCheck2Default with default headers values
func NewZzHealthServiceCheck2Default(code int) *ZzHealthServiceCheck2Default {
	return &ZzHealthServiceCheck2Default{
		_statusCode: code,
	}
}

/*
ZzHealthServiceCheck2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ZzHealthServiceCheck2Default struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this zz health service check2 default response has a 2xx status code
func (o *ZzHealthServiceCheck2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this zz health service check2 default response has a 3xx status code
func (o *ZzHealthServiceCheck2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this zz health service check2 default response has a 4xx status code
func (o *ZzHealthServiceCheck2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this zz health service check2 default response has a 5xx status code
func (o *ZzHealthServiceCheck2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this zz health service check2 default response a status code equal to that given
func (o *ZzHealthServiceCheck2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the zz health service check2 default response
func (o *ZzHealthServiceCheck2Default) Code() int {
	return o._statusCode
}

func (o *ZzHealthServiceCheck2Default) Error() string {
	return fmt.Sprintf("[GET /hc][%d] ZzHealthService_Check2 default  %+v", o._statusCode, o.Payload)
}

func (o *ZzHealthServiceCheck2Default) String() string {
	return fmt.Sprintf("[GET /hc][%d] ZzHealthService_Check2 default  %+v", o._statusCode, o.Payload)
}

func (o *ZzHealthServiceCheck2Default) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *ZzHealthServiceCheck2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
