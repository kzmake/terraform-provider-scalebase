// Code generated by go-swagger; DO NOT EDIT.

package custom_field_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new custom field service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for custom field service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CustomFieldServiceGetResource(params *CustomFieldServiceGetResourceParams, opts ...ClientOption) (*CustomFieldServiceGetResourceOK, error)

	CustomFieldServiceGetResource2(params *CustomFieldServiceGetResource2Params, opts ...ClientOption) (*CustomFieldServiceGetResource2OK, error)

	CustomFieldServiceListCustomFieldMasters(params *CustomFieldServiceListCustomFieldMastersParams, opts ...ClientOption) (*CustomFieldServiceListCustomFieldMastersOK, error)

	CustomFieldServiceListCustomFieldMasters2(params *CustomFieldServiceListCustomFieldMasters2Params, opts ...ClientOption) (*CustomFieldServiceListCustomFieldMasters2OK, error)

	CustomFieldServiceUpdateResource(params *CustomFieldServiceUpdateResourceParams, opts ...ClientOption) (*CustomFieldServiceUpdateResourceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CustomFieldServiceGetResource 対象リソースのカスタムフィールド取得s

対象リソースのカスタムフィールド取得します。
*/
func (a *Client) CustomFieldServiceGetResource(params *CustomFieldServiceGetResourceParams, opts ...ClientOption) (*CustomFieldServiceGetResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomFieldServiceGetResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CustomFieldService_GetResource",
		Method:             "POST",
		PathPattern:        "/v1/customfield/resource/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomFieldServiceGetResourceReader{formats: a.formats},
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
	success, ok := result.(*CustomFieldServiceGetResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomFieldServiceGetResourceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CustomFieldServiceGetResource2 対象リソースのカスタムフィールド取得s

対象リソースのカスタムフィールド取得します。
*/
func (a *Client) CustomFieldServiceGetResource2(params *CustomFieldServiceGetResource2Params, opts ...ClientOption) (*CustomFieldServiceGetResource2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomFieldServiceGetResource2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "CustomFieldService_GetResource2",
		Method:             "GET",
		PathPattern:        "/v1/customfield/resource/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomFieldServiceGetResource2Reader{formats: a.formats},
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
	success, ok := result.(*CustomFieldServiceGetResource2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomFieldServiceGetResource2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CustomFieldServiceListCustomFieldMasters カスタムフィールドマスターの一覧取得s

カスタムフィールドマスターを一覧取得します。
*/
func (a *Client) CustomFieldServiceListCustomFieldMasters(params *CustomFieldServiceListCustomFieldMastersParams, opts ...ClientOption) (*CustomFieldServiceListCustomFieldMastersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomFieldServiceListCustomFieldMastersParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CustomFieldService_ListCustomFieldMasters",
		Method:             "POST",
		PathPattern:        "/v1/customfield/master/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomFieldServiceListCustomFieldMastersReader{formats: a.formats},
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
	success, ok := result.(*CustomFieldServiceListCustomFieldMastersOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomFieldServiceListCustomFieldMastersDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CustomFieldServiceListCustomFieldMasters2 カスタムフィールドマスターの一覧取得s

カスタムフィールドマスターを一覧取得します。
*/
func (a *Client) CustomFieldServiceListCustomFieldMasters2(params *CustomFieldServiceListCustomFieldMasters2Params, opts ...ClientOption) (*CustomFieldServiceListCustomFieldMasters2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomFieldServiceListCustomFieldMasters2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "CustomFieldService_ListCustomFieldMasters2",
		Method:             "GET",
		PathPattern:        "/v1/customfield/master/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomFieldServiceListCustomFieldMasters2Reader{formats: a.formats},
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
	success, ok := result.(*CustomFieldServiceListCustomFieldMasters2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomFieldServiceListCustomFieldMasters2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
CustomFieldServiceUpdateResource カスタムフィールド更新s

カスタムフィールドの更新を行います。
*/
func (a *Client) CustomFieldServiceUpdateResource(params *CustomFieldServiceUpdateResourceParams, opts ...ClientOption) (*CustomFieldServiceUpdateResourceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCustomFieldServiceUpdateResourceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CustomFieldService_UpdateResource",
		Method:             "PATCH",
		PathPattern:        "/v1/customfield/resource/update",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CustomFieldServiceUpdateResourceReader{formats: a.formats},
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
	success, ok := result.(*CustomFieldServiceUpdateResourceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CustomFieldServiceUpdateResourceDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
