// Code generated by go-swagger; DO NOT EDIT.

package contract_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

// NewContractServiceListContractsParams creates a new ContractServiceListContractsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContractServiceListContractsParams() *ContractServiceListContractsParams {
	return &ContractServiceListContractsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewContractServiceListContractsParamsWithTimeout creates a new ContractServiceListContractsParams object
// with the ability to set a timeout on a request.
func NewContractServiceListContractsParamsWithTimeout(timeout time.Duration) *ContractServiceListContractsParams {
	return &ContractServiceListContractsParams{
		timeout: timeout,
	}
}

// NewContractServiceListContractsParamsWithContext creates a new ContractServiceListContractsParams object
// with the ability to set a context for a request.
func NewContractServiceListContractsParamsWithContext(ctx context.Context) *ContractServiceListContractsParams {
	return &ContractServiceListContractsParams{
		Context: ctx,
	}
}

// NewContractServiceListContractsParamsWithHTTPClient creates a new ContractServiceListContractsParams object
// with the ability to set a custom HTTPClient for a request.
func NewContractServiceListContractsParamsWithHTTPClient(client *http.Client) *ContractServiceListContractsParams {
	return &ContractServiceListContractsParams{
		HTTPClient: client,
	}
}

/*
ContractServiceListContractsParams contains all the parameters to send to the API endpoint

	for the contract service list contracts operation.

	Typically these are written to a http.Request.
*/
type ContractServiceListContractsParams struct {

	// Body.
	Body *models.V1ListContractsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the contract service list contracts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContractServiceListContractsParams) WithDefaults() *ContractServiceListContractsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contract service list contracts params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContractServiceListContractsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contract service list contracts params
func (o *ContractServiceListContractsParams) WithTimeout(timeout time.Duration) *ContractServiceListContractsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contract service list contracts params
func (o *ContractServiceListContractsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contract service list contracts params
func (o *ContractServiceListContractsParams) WithContext(ctx context.Context) *ContractServiceListContractsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contract service list contracts params
func (o *ContractServiceListContractsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contract service list contracts params
func (o *ContractServiceListContractsParams) WithHTTPClient(client *http.Client) *ContractServiceListContractsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contract service list contracts params
func (o *ContractServiceListContractsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the contract service list contracts params
func (o *ContractServiceListContractsParams) WithBody(body *models.V1ListContractsRequest) *ContractServiceListContractsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the contract service list contracts params
func (o *ContractServiceListContractsParams) SetBody(body *models.V1ListContractsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ContractServiceListContractsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
