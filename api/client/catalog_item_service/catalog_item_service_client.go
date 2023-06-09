// Code generated by go-swagger; DO NOT EDIT.

package catalog_item_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new catalog item service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for catalog item service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CatalogItemServiceGetCatalogItem(params *CatalogItemServiceGetCatalogItemParams, opts ...ClientOption) (*CatalogItemServiceGetCatalogItemOK, error)

	CatalogItemServiceGetCatalogItem2(params *CatalogItemServiceGetCatalogItem2Params, opts ...ClientOption) (*CatalogItemServiceGetCatalogItem2OK, error)

	CatalogItemServiceListCatalogItems(params *CatalogItemServiceListCatalogItemsParams, opts ...ClientOption) (*CatalogItemServiceListCatalogItemsOK, error)

	CatalogItemServiceListCatalogItems2(params *CatalogItemServiceListCatalogItems2Params, opts ...ClientOption) (*CatalogItemServiceListCatalogItems2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CatalogItemServiceGetCatalogItem カタログアイテムの取得s

指定したカタログアイテムを取得します。
*/
func (a *Client) CatalogItemServiceGetCatalogItem(params *CatalogItemServiceGetCatalogItemParams, opts ...ClientOption) (*CatalogItemServiceGetCatalogItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogItemServiceGetCatalogItemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CatalogItemService_GetCatalogItem",
		Method:             "POST",
		PathPattern:        "/v1/catalogitem/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogItemServiceGetCatalogItemReader{formats: a.formats},
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
	success, ok := result.(*CatalogItemServiceGetCatalogItemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CatalogItemServiceGetCatalogItemDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CatalogItemServiceGetCatalogItem2 カタログアイテムの取得s

指定したカタログアイテムを取得します。
*/
func (a *Client) CatalogItemServiceGetCatalogItem2(params *CatalogItemServiceGetCatalogItem2Params, opts ...ClientOption) (*CatalogItemServiceGetCatalogItem2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogItemServiceGetCatalogItem2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "CatalogItemService_GetCatalogItem2",
		Method:             "GET",
		PathPattern:        "/v1/catalogitem/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogItemServiceGetCatalogItem2Reader{formats: a.formats},
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
	success, ok := result.(*CatalogItemServiceGetCatalogItem2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CatalogItemServiceGetCatalogItem2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CatalogItemServiceListCatalogItems カタログアイテムの一覧取得s

カタログアイテムの一覧を取得します。
*/
func (a *Client) CatalogItemServiceListCatalogItems(params *CatalogItemServiceListCatalogItemsParams, opts ...ClientOption) (*CatalogItemServiceListCatalogItemsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogItemServiceListCatalogItemsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CatalogItemService_ListCatalogItems",
		Method:             "POST",
		PathPattern:        "/v1/catalogitem/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogItemServiceListCatalogItemsReader{formats: a.formats},
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
	success, ok := result.(*CatalogItemServiceListCatalogItemsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CatalogItemServiceListCatalogItemsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CatalogItemServiceListCatalogItems2 カタログアイテムの一覧取得s

カタログアイテムの一覧を取得します。
*/
func (a *Client) CatalogItemServiceListCatalogItems2(params *CatalogItemServiceListCatalogItems2Params, opts ...ClientOption) (*CatalogItemServiceListCatalogItems2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCatalogItemServiceListCatalogItems2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "CatalogItemService_ListCatalogItems2",
		Method:             "GET",
		PathPattern:        "/v1/catalogitem/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CatalogItemServiceListCatalogItems2Reader{formats: a.formats},
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
	success, ok := result.(*CatalogItemServiceListCatalogItems2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CatalogItemServiceListCatalogItems2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
