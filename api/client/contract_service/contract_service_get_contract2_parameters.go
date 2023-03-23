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
	"github.com/go-openapi/swag"
)

// NewContractServiceGetContract2Params creates a new ContractServiceGetContract2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewContractServiceGetContract2Params() *ContractServiceGetContract2Params {
	return &ContractServiceGetContract2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewContractServiceGetContract2ParamsWithTimeout creates a new ContractServiceGetContract2Params object
// with the ability to set a timeout on a request.
func NewContractServiceGetContract2ParamsWithTimeout(timeout time.Duration) *ContractServiceGetContract2Params {
	return &ContractServiceGetContract2Params{
		timeout: timeout,
	}
}

// NewContractServiceGetContract2ParamsWithContext creates a new ContractServiceGetContract2Params object
// with the ability to set a context for a request.
func NewContractServiceGetContract2ParamsWithContext(ctx context.Context) *ContractServiceGetContract2Params {
	return &ContractServiceGetContract2Params{
		Context: ctx,
	}
}

// NewContractServiceGetContract2ParamsWithHTTPClient creates a new ContractServiceGetContract2Params object
// with the ability to set a custom HTTPClient for a request.
func NewContractServiceGetContract2ParamsWithHTTPClient(client *http.Client) *ContractServiceGetContract2Params {
	return &ContractServiceGetContract2Params{
		HTTPClient: client,
	}
}

/*
ContractServiceGetContract2Params contains all the parameters to send to the API endpoint

	for the contract service get contract2 operation.

	Typically these are written to a http.Request.
*/
type ContractServiceGetContract2Params struct {

	/* ID.

	   契約ID
	*/
	ID string

	/* Version.

	   契約バージョン。指定がなければ最新バージョンの契約が返却される

	   Format: int32
	*/
	Version *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the contract service get contract2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContractServiceGetContract2Params) WithDefaults() *ContractServiceGetContract2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the contract service get contract2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ContractServiceGetContract2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the contract service get contract2 params
func (o *ContractServiceGetContract2Params) WithTimeout(timeout time.Duration) *ContractServiceGetContract2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the contract service get contract2 params
func (o *ContractServiceGetContract2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the contract service get contract2 params
func (o *ContractServiceGetContract2Params) WithContext(ctx context.Context) *ContractServiceGetContract2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the contract service get contract2 params
func (o *ContractServiceGetContract2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the contract service get contract2 params
func (o *ContractServiceGetContract2Params) WithHTTPClient(client *http.Client) *ContractServiceGetContract2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the contract service get contract2 params
func (o *ContractServiceGetContract2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the contract service get contract2 params
func (o *ContractServiceGetContract2Params) WithID(id string) *ContractServiceGetContract2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the contract service get contract2 params
func (o *ContractServiceGetContract2Params) SetID(id string) {
	o.ID = id
}

// WithVersion adds the version to the contract service get contract2 params
func (o *ContractServiceGetContract2Params) WithVersion(version *int32) *ContractServiceGetContract2Params {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the contract service get contract2 params
func (o *ContractServiceGetContract2Params) SetVersion(version *int32) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *ContractServiceGetContract2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param id
	qrID := o.ID
	qID := qrID
	if qID != "" {

		if err := r.SetQueryParam("id", qID); err != nil {
			return err
		}
	}

	if o.Version != nil {

		// query param version
		var qrVersion int32

		if o.Version != nil {
			qrVersion = *o.Version
		}
		qVersion := swag.FormatInt32(qrVersion)
		if qVersion != "" {

			if err := r.SetQueryParam("version", qVersion); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
