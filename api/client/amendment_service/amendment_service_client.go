// Code generated by go-swagger; DO NOT EDIT.

package amendment_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new amendment service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for amendment service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	AmendmentServiceAddOptionItem(params *AmendmentServiceAddOptionItemParams, opts ...ClientOption) (*AmendmentServiceAddOptionItemOK, error)

	AmendmentServiceListAmendments(params *AmendmentServiceListAmendmentsParams, opts ...ClientOption) (*AmendmentServiceListAmendmentsOK, error)

	AmendmentServiceListAmendments2(params *AmendmentServiceListAmendments2Params, opts ...ClientOption) (*AmendmentServiceListAmendments2OK, error)

	AmendmentServiceListAmendmentsByContractOptionalID(params *AmendmentServiceListAmendmentsByContractOptionalIDParams, opts ...ClientOption) (*AmendmentServiceListAmendmentsByContractOptionalIDOK, error)

	AmendmentServiceListAmendmentsByContractOptionalId2(params *AmendmentServiceListAmendmentsByContractOptionalId2Params, opts ...ClientOption) (*AmendmentServiceListAmendmentsByContractOptionalId2OK, error)

	AmendmentServiceRenewContractItem(params *AmendmentServiceRenewContractItemParams, opts ...ClientOption) (*AmendmentServiceRenewContractItemOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
AmendmentServiceAddOptionItem オプションアイテムの追加s

指定したカタログアイテムをオプションアイテムとして契約に追加します。
*/
func (a *Client) AmendmentServiceAddOptionItem(params *AmendmentServiceAddOptionItemParams, opts ...ClientOption) (*AmendmentServiceAddOptionItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAmendmentServiceAddOptionItemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AmendmentService_AddOptionItem",
		Method:             "POST",
		PathPattern:        "/v1/amendment/addoptionitem",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AmendmentServiceAddOptionItemReader{formats: a.formats},
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
	success, ok := result.(*AmendmentServiceAddOptionItemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AmendmentServiceAddOptionItemDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AmendmentServiceListAmendments 改定の一覧取得s

プロバイダーに紐づく改定の一覧を取得します
*/
func (a *Client) AmendmentServiceListAmendments(params *AmendmentServiceListAmendmentsParams, opts ...ClientOption) (*AmendmentServiceListAmendmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAmendmentServiceListAmendmentsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AmendmentService_ListAmendments",
		Method:             "POST",
		PathPattern:        "/v1/amendment/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AmendmentServiceListAmendmentsReader{formats: a.formats},
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
	success, ok := result.(*AmendmentServiceListAmendmentsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AmendmentServiceListAmendmentsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AmendmentServiceListAmendments2 改定の一覧取得s

プロバイダーに紐づく改定の一覧を取得します
*/
func (a *Client) AmendmentServiceListAmendments2(params *AmendmentServiceListAmendments2Params, opts ...ClientOption) (*AmendmentServiceListAmendments2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAmendmentServiceListAmendments2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "AmendmentService_ListAmendments2",
		Method:             "GET",
		PathPattern:        "/v1/amendment/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AmendmentServiceListAmendments2Reader{formats: a.formats},
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
	success, ok := result.(*AmendmentServiceListAmendments2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AmendmentServiceListAmendments2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AmendmentServiceListAmendmentsByContractOptionalID 契約管理s ID に紐づく改定の一覧取得

契約管理IDに紐づく改定を一覧取得します。
*/
func (a *Client) AmendmentServiceListAmendmentsByContractOptionalID(params *AmendmentServiceListAmendmentsByContractOptionalIDParams, opts ...ClientOption) (*AmendmentServiceListAmendmentsByContractOptionalIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAmendmentServiceListAmendmentsByContractOptionalIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AmendmentService_ListAmendmentsByContractOptionalId",
		Method:             "POST",
		PathPattern:        "/v1/amendment/listbycontractoptionalid",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AmendmentServiceListAmendmentsByContractOptionalIDReader{formats: a.formats},
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
	success, ok := result.(*AmendmentServiceListAmendmentsByContractOptionalIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AmendmentServiceListAmendmentsByContractOptionalIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AmendmentServiceListAmendmentsByContractOptionalId2 契約管理s ID に紐づく改定の一覧取得

契約管理IDに紐づく改定を一覧取得します。
*/
func (a *Client) AmendmentServiceListAmendmentsByContractOptionalId2(params *AmendmentServiceListAmendmentsByContractOptionalId2Params, opts ...ClientOption) (*AmendmentServiceListAmendmentsByContractOptionalId2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAmendmentServiceListAmendmentsByContractOptionalId2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "AmendmentService_ListAmendmentsByContractOptionalId2",
		Method:             "GET",
		PathPattern:        "/v1/amendment/listbycontractoptionalid",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AmendmentServiceListAmendmentsByContractOptionalId2Reader{formats: a.formats},
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
	success, ok := result.(*AmendmentServiceListAmendmentsByContractOptionalId2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AmendmentServiceListAmendmentsByContractOptionalId2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
AmendmentServiceRenewContractItem 契約アイテムの更新s

指定した契約アイテムを更新します。
*/
func (a *Client) AmendmentServiceRenewContractItem(params *AmendmentServiceRenewContractItemParams, opts ...ClientOption) (*AmendmentServiceRenewContractItemOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAmendmentServiceRenewContractItemParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "AmendmentService_RenewContractItem",
		Method:             "POST",
		PathPattern:        "/v1/amendment/renewcontractitem",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &AmendmentServiceRenewContractItemReader{formats: a.formats},
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
	success, ok := result.(*AmendmentServiceRenewContractItemOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*AmendmentServiceRenewContractItemDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}