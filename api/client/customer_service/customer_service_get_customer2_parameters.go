// Code generated by go-swagger; DO NOT EDIT.

package customer_service

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
)

// NewCustomerServiceGetCustomer2Params creates a new CustomerServiceGetCustomer2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerServiceGetCustomer2Params() *CustomerServiceGetCustomer2Params {
	return &CustomerServiceGetCustomer2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerServiceGetCustomer2ParamsWithTimeout creates a new CustomerServiceGetCustomer2Params object
// with the ability to set a timeout on a request.
func NewCustomerServiceGetCustomer2ParamsWithTimeout(timeout time.Duration) *CustomerServiceGetCustomer2Params {
	return &CustomerServiceGetCustomer2Params{
		timeout: timeout,
	}
}

// NewCustomerServiceGetCustomer2ParamsWithContext creates a new CustomerServiceGetCustomer2Params object
// with the ability to set a context for a request.
func NewCustomerServiceGetCustomer2ParamsWithContext(ctx context.Context) *CustomerServiceGetCustomer2Params {
	return &CustomerServiceGetCustomer2Params{
		Context: ctx,
	}
}

// NewCustomerServiceGetCustomer2ParamsWithHTTPClient creates a new CustomerServiceGetCustomer2Params object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerServiceGetCustomer2ParamsWithHTTPClient(client *http.Client) *CustomerServiceGetCustomer2Params {
	return &CustomerServiceGetCustomer2Params{
		HTTPClient: client,
	}
}

/*
CustomerServiceGetCustomer2Params contains all the parameters to send to the API endpoint

	for the customer service get customer2 operation.

	Typically these are written to a http.Request.
*/
type CustomerServiceGetCustomer2Params struct {

	/* ID.

	   顧客ID
	*/
	ID *string

	/* OptionalID.

	   顧客管理ID(顧客IDの代わりに指定可)
	*/
	OptionalID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer service get customer2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerServiceGetCustomer2Params) WithDefaults() *CustomerServiceGetCustomer2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer service get customer2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerServiceGetCustomer2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer service get customer2 params
func (o *CustomerServiceGetCustomer2Params) WithTimeout(timeout time.Duration) *CustomerServiceGetCustomer2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer service get customer2 params
func (o *CustomerServiceGetCustomer2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer service get customer2 params
func (o *CustomerServiceGetCustomer2Params) WithContext(ctx context.Context) *CustomerServiceGetCustomer2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer service get customer2 params
func (o *CustomerServiceGetCustomer2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer service get customer2 params
func (o *CustomerServiceGetCustomer2Params) WithHTTPClient(client *http.Client) *CustomerServiceGetCustomer2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer service get customer2 params
func (o *CustomerServiceGetCustomer2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the customer service get customer2 params
func (o *CustomerServiceGetCustomer2Params) WithID(id *string) *CustomerServiceGetCustomer2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the customer service get customer2 params
func (o *CustomerServiceGetCustomer2Params) SetID(id *string) {
	o.ID = id
}

// WithOptionalID adds the optionalID to the customer service get customer2 params
func (o *CustomerServiceGetCustomer2Params) WithOptionalID(optionalID *string) *CustomerServiceGetCustomer2Params {
	o.SetOptionalID(optionalID)
	return o
}

// SetOptionalID adds the optionalId to the customer service get customer2 params
func (o *CustomerServiceGetCustomer2Params) SetOptionalID(optionalID *string) {
	o.OptionalID = optionalID
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerServiceGetCustomer2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.ID != nil {

		// query param id
		var qrID string

		if o.ID != nil {
			qrID = *o.ID
		}
		qID := qrID
		if qID != "" {

			if err := r.SetQueryParam("id", qID); err != nil {
				return err
			}
		}
	}

	if o.OptionalID != nil {

		// query param optionalId
		var qrOptionalID string

		if o.OptionalID != nil {
			qrOptionalID = *o.OptionalID
		}
		qOptionalID := qrOptionalID
		if qOptionalID != "" {

			if err := r.SetQueryParam("optionalId", qOptionalID); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}