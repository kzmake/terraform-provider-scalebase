// Code generated by go-swagger; DO NOT EDIT.

package contract_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new contract service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for contract service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	ContractServiceCreateContract(params *ContractServiceCreateContractParams, opts ...ClientOption) (*ContractServiceCreateContractOK, error)

	ContractServiceGetContract(params *ContractServiceGetContractParams, opts ...ClientOption) (*ContractServiceGetContractOK, error)

	ContractServiceGetContract2(params *ContractServiceGetContract2Params, opts ...ClientOption) (*ContractServiceGetContract2OK, error)

	ContractServiceGetContractByOptionalID(params *ContractServiceGetContractByOptionalIDParams, opts ...ClientOption) (*ContractServiceGetContractByOptionalIDOK, error)

	ContractServiceGetContractByOptionalId2(params *ContractServiceGetContractByOptionalId2Params, opts ...ClientOption) (*ContractServiceGetContractByOptionalId2OK, error)

	ContractServiceListContracts(params *ContractServiceListContractsParams, opts ...ClientOption) (*ContractServiceListContractsOK, error)

	ContractServiceListContracts2(params *ContractServiceListContracts2Params, opts ...ClientOption) (*ContractServiceListContracts2OK, error)

	ContractServiceListContractsByCustomer(params *ContractServiceListContractsByCustomerParams, opts ...ClientOption) (*ContractServiceListContractsByCustomerOK, error)

	ContractServiceListContractsByCustomer2(params *ContractServiceListContractsByCustomer2Params, opts ...ClientOption) (*ContractServiceListContractsByCustomer2OK, error)

	ContractServiceListContractsByCustomer3(params *ContractServiceListContractsByCustomer3Params, opts ...ClientOption) (*ContractServiceListContractsByCustomer3OK, error)

	ContractServiceListContractsByCustomer4(params *ContractServiceListContractsByCustomer4Params, opts ...ClientOption) (*ContractServiceListContractsByCustomer4OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
ContractServiceCreateContract 契約の作成s

新規に契約を作成します。
*/
func (a *Client) ContractServiceCreateContract(params *ContractServiceCreateContractParams, opts ...ClientOption) (*ContractServiceCreateContractOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractServiceCreateContractParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContractService_CreateContract",
		Method:             "POST",
		PathPattern:        "/v1/contract/create",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContractServiceCreateContractReader{formats: a.formats},
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
	success, ok := result.(*ContractServiceCreateContractOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ContractServiceCreateContractDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ContractServiceGetContract 契約の取得s

指定されたIDを持つ契約を取得します。
*/
func (a *Client) ContractServiceGetContract(params *ContractServiceGetContractParams, opts ...ClientOption) (*ContractServiceGetContractOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractServiceGetContractParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContractService_GetContract",
		Method:             "POST",
		PathPattern:        "/v1/contract/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContractServiceGetContractReader{formats: a.formats},
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
	success, ok := result.(*ContractServiceGetContractOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ContractServiceGetContractDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ContractServiceGetContract2 契約の取得s

指定されたIDを持つ契約を取得します。
*/
func (a *Client) ContractServiceGetContract2(params *ContractServiceGetContract2Params, opts ...ClientOption) (*ContractServiceGetContract2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractServiceGetContract2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContractService_GetContract2",
		Method:             "GET",
		PathPattern:        "/v1/contract/get",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContractServiceGetContract2Reader{formats: a.formats},
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
	success, ok := result.(*ContractServiceGetContract2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ContractServiceGetContract2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ContractServiceGetContractByOptionalID 契約管理s ID を指定した契約の取得

契約管理IDを指定して契約を取得します。
*/
func (a *Client) ContractServiceGetContractByOptionalID(params *ContractServiceGetContractByOptionalIDParams, opts ...ClientOption) (*ContractServiceGetContractByOptionalIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractServiceGetContractByOptionalIDParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContractService_GetContractByOptionalId",
		Method:             "POST",
		PathPattern:        "/v1/contract/getbyoptionalid",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContractServiceGetContractByOptionalIDReader{formats: a.formats},
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
	success, ok := result.(*ContractServiceGetContractByOptionalIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ContractServiceGetContractByOptionalIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ContractServiceGetContractByOptionalId2 契約管理s ID を指定した契約の取得

契約管理IDを指定して契約を取得します。
*/
func (a *Client) ContractServiceGetContractByOptionalId2(params *ContractServiceGetContractByOptionalId2Params, opts ...ClientOption) (*ContractServiceGetContractByOptionalId2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractServiceGetContractByOptionalId2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContractService_GetContractByOptionalId2",
		Method:             "GET",
		PathPattern:        "/v1/contract/getbyoptionalid",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContractServiceGetContractByOptionalId2Reader{formats: a.formats},
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
	success, ok := result.(*ContractServiceGetContractByOptionalId2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ContractServiceGetContractByOptionalId2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ContractServiceListContracts 契約の一覧取得s

契約の一覧を取得します。
*/
func (a *Client) ContractServiceListContracts(params *ContractServiceListContractsParams, opts ...ClientOption) (*ContractServiceListContractsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractServiceListContractsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContractService_ListContracts",
		Method:             "POST",
		PathPattern:        "/v1/contract/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContractServiceListContractsReader{formats: a.formats},
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
	success, ok := result.(*ContractServiceListContractsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ContractServiceListContractsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ContractServiceListContracts2 契約の一覧取得s

契約の一覧を取得します。
*/
func (a *Client) ContractServiceListContracts2(params *ContractServiceListContracts2Params, opts ...ClientOption) (*ContractServiceListContracts2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractServiceListContracts2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContractService_ListContracts2",
		Method:             "GET",
		PathPattern:        "/v1/contract/list",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContractServiceListContracts2Reader{formats: a.formats},
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
	success, ok := result.(*ContractServiceListContracts2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ContractServiceListContracts2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ContractServiceListContractsByCustomer 顧客に紐づく契約の一覧取得s

顧客に紐づく契約を一覧取得します。
*/
func (a *Client) ContractServiceListContractsByCustomer(params *ContractServiceListContractsByCustomerParams, opts ...ClientOption) (*ContractServiceListContractsByCustomerOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractServiceListContractsByCustomerParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContractService_ListContractsByCustomer",
		Method:             "POST",
		PathPattern:        "/v1/contract/listbycustomer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContractServiceListContractsByCustomerReader{formats: a.formats},
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
	success, ok := result.(*ContractServiceListContractsByCustomerOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ContractServiceListContractsByCustomerDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ContractServiceListContractsByCustomer2 顧客に紐づく契約の一覧取得s

顧客に紐づく契約を一覧取得します。
*/
func (a *Client) ContractServiceListContractsByCustomer2(params *ContractServiceListContractsByCustomer2Params, opts ...ClientOption) (*ContractServiceListContractsByCustomer2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractServiceListContractsByCustomer2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContractService_ListContractsByCustomer2",
		Method:             "GET",
		PathPattern:        "/v1/contract/listbycustomer",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContractServiceListContractsByCustomer2Reader{formats: a.formats},
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
	success, ok := result.(*ContractServiceListContractsByCustomer2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ContractServiceListContractsByCustomer2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ContractServiceListContractsByCustomer3 顧客に紐づく契約の一覧取得s

顧客に紐づく契約を一覧取得します。
*/
func (a *Client) ContractServiceListContractsByCustomer3(params *ContractServiceListContractsByCustomer3Params, opts ...ClientOption) (*ContractServiceListContractsByCustomer3OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractServiceListContractsByCustomer3Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContractService_ListContractsByCustomer3",
		Method:             "POST",
		PathPattern:        "/v1/contract/listbycustomerid",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContractServiceListContractsByCustomer3Reader{formats: a.formats},
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
	success, ok := result.(*ContractServiceListContractsByCustomer3OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ContractServiceListContractsByCustomer3Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ContractServiceListContractsByCustomer4 顧客に紐づく契約の一覧取得s

顧客に紐づく契約を一覧取得します。
*/
func (a *Client) ContractServiceListContractsByCustomer4(params *ContractServiceListContractsByCustomer4Params, opts ...ClientOption) (*ContractServiceListContractsByCustomer4OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewContractServiceListContractsByCustomer4Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "ContractService_ListContractsByCustomer4",
		Method:             "GET",
		PathPattern:        "/v1/contract/listbycustomerid",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ContractServiceListContractsByCustomer4Reader{formats: a.formats},
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
	success, ok := result.(*ContractServiceListContractsByCustomer4OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ContractServiceListContractsByCustomer4Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}