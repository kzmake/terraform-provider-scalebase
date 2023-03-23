// Code generated by go-swagger; DO NOT EDIT.

package catalog_item_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

// CatalogItemServiceListCatalogItemsReader is a Reader for the CatalogItemServiceListCatalogItems structure.
type CatalogItemServiceListCatalogItemsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CatalogItemServiceListCatalogItemsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCatalogItemServiceListCatalogItemsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewCatalogItemServiceListCatalogItemsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewCatalogItemServiceListCatalogItemsOK creates a CatalogItemServiceListCatalogItemsOK with default headers values
func NewCatalogItemServiceListCatalogItemsOK() *CatalogItemServiceListCatalogItemsOK {
	return &CatalogItemServiceListCatalogItemsOK{}
}

/*
CatalogItemServiceListCatalogItemsOK describes a response with status code 200, with default header values.

A successful response.
*/
type CatalogItemServiceListCatalogItemsOK struct {
	Payload *models.V1ListCatalogItemsResponse
}

// IsSuccess returns true when this catalog item service list catalog items o k response has a 2xx status code
func (o *CatalogItemServiceListCatalogItemsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this catalog item service list catalog items o k response has a 3xx status code
func (o *CatalogItemServiceListCatalogItemsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this catalog item service list catalog items o k response has a 4xx status code
func (o *CatalogItemServiceListCatalogItemsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this catalog item service list catalog items o k response has a 5xx status code
func (o *CatalogItemServiceListCatalogItemsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this catalog item service list catalog items o k response a status code equal to that given
func (o *CatalogItemServiceListCatalogItemsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the catalog item service list catalog items o k response
func (o *CatalogItemServiceListCatalogItemsOK) Code() int {
	return 200
}

func (o *CatalogItemServiceListCatalogItemsOK) Error() string {
	return fmt.Sprintf("[POST /v1/catalogitem/list][%d] catalogItemServiceListCatalogItemsOK  %+v", 200, o.Payload)
}

func (o *CatalogItemServiceListCatalogItemsOK) String() string {
	return fmt.Sprintf("[POST /v1/catalogitem/list][%d] catalogItemServiceListCatalogItemsOK  %+v", 200, o.Payload)
}

func (o *CatalogItemServiceListCatalogItemsOK) GetPayload() *models.V1ListCatalogItemsResponse {
	return o.Payload
}

func (o *CatalogItemServiceListCatalogItemsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1ListCatalogItemsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCatalogItemServiceListCatalogItemsDefault creates a CatalogItemServiceListCatalogItemsDefault with default headers values
func NewCatalogItemServiceListCatalogItemsDefault(code int) *CatalogItemServiceListCatalogItemsDefault {
	return &CatalogItemServiceListCatalogItemsDefault{
		_statusCode: code,
	}
}

/*
CatalogItemServiceListCatalogItemsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type CatalogItemServiceListCatalogItemsDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this catalog item service list catalog items default response has a 2xx status code
func (o *CatalogItemServiceListCatalogItemsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this catalog item service list catalog items default response has a 3xx status code
func (o *CatalogItemServiceListCatalogItemsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this catalog item service list catalog items default response has a 4xx status code
func (o *CatalogItemServiceListCatalogItemsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this catalog item service list catalog items default response has a 5xx status code
func (o *CatalogItemServiceListCatalogItemsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this catalog item service list catalog items default response a status code equal to that given
func (o *CatalogItemServiceListCatalogItemsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the catalog item service list catalog items default response
func (o *CatalogItemServiceListCatalogItemsDefault) Code() int {
	return o._statusCode
}

func (o *CatalogItemServiceListCatalogItemsDefault) Error() string {
	return fmt.Sprintf("[POST /v1/catalogitem/list][%d] CatalogItemService_ListCatalogItems default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogItemServiceListCatalogItemsDefault) String() string {
	return fmt.Sprintf("[POST /v1/catalogitem/list][%d] CatalogItemService_ListCatalogItems default  %+v", o._statusCode, o.Payload)
}

func (o *CatalogItemServiceListCatalogItemsDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *CatalogItemServiceListCatalogItemsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}