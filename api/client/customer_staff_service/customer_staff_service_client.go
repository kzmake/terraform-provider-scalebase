// Code generated by go-swagger; DO NOT EDIT.

package customer_staff_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new customer staff service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customer staff service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CustomerStaffServiceCreateCustomerStaff(params *CustomerStaffServiceCreateCustomerStaffParams, opts ...ClientOption) (*CustomerStaffServiceCreateCustomerStaffOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CustomerStaffServiceCreateCustomerStaff 顧客担当者の作成s

新規に顧客担当者を作成します。
*/
func (a *Client) CustomerStaffServiceCreateCustomerStaff(params *CustomerStaffServiceCreateCustomerStaffParams, opts ...ClientOption) (*CustomerStaffServiceCreateCustomerStaffOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomerStaffServiceCreateCustomerStaffParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CustomerStaffService_CreateCustomerStaff",
		Method:             "POST",
		PathPattern:        "/v1/customerstaff/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomerStaffServiceCreateCustomerStaffReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CustomerStaffServiceCreateCustomerStaffOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomerStaffServiceCreateCustomerStaffDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
