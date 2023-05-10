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

	"github.com/kzmake/terraform-provider-scalebase/api/models"
)

// NewCustomerServiceDeleteCustomerParams creates a new CustomerServiceDeleteCustomerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerServiceDeleteCustomerParams() *CustomerServiceDeleteCustomerParams {
	return &CustomerServiceDeleteCustomerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerServiceDeleteCustomerParamsWithTimeout creates a new CustomerServiceDeleteCustomerParams object
// with the ability to set a timeout on a request.
func NewCustomerServiceDeleteCustomerParamsWithTimeout(timeout time.Duration) *CustomerServiceDeleteCustomerParams {
	return &CustomerServiceDeleteCustomerParams{
		timeout: timeout,
	}
}

// NewCustomerServiceDeleteCustomerParamsWithContext creates a new CustomerServiceDeleteCustomerParams object
// with the ability to set a context for a request.
func NewCustomerServiceDeleteCustomerParamsWithContext(ctx context.Context) *CustomerServiceDeleteCustomerParams {
	return &CustomerServiceDeleteCustomerParams{
		Context: ctx,
	}
}

// NewCustomerServiceDeleteCustomerParamsWithHTTPClient creates a new CustomerServiceDeleteCustomerParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerServiceDeleteCustomerParamsWithHTTPClient(client *http.Client) *CustomerServiceDeleteCustomerParams {
	return &CustomerServiceDeleteCustomerParams{
		HTTPClient: client,
	}
}

/*
CustomerServiceDeleteCustomerParams contains all the parameters to send to the API endpoint

	for the customer service delete customer operation.

	Typically these are written to a http.Request.
*/
type CustomerServiceDeleteCustomerParams struct {

	// Body.
	Body *models.V1DeleteCustomerRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer service delete customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerServiceDeleteCustomerParams) WithDefaults() *CustomerServiceDeleteCustomerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer service delete customer params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerServiceDeleteCustomerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer service delete customer params
func (o *CustomerServiceDeleteCustomerParams) WithTimeout(timeout time.Duration) *CustomerServiceDeleteCustomerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer service delete customer params
func (o *CustomerServiceDeleteCustomerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer service delete customer params
func (o *CustomerServiceDeleteCustomerParams) WithContext(ctx context.Context) *CustomerServiceDeleteCustomerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer service delete customer params
func (o *CustomerServiceDeleteCustomerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer service delete customer params
func (o *CustomerServiceDeleteCustomerParams) WithHTTPClient(client *http.Client) *CustomerServiceDeleteCustomerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer service delete customer params
func (o *CustomerServiceDeleteCustomerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the customer service delete customer params
func (o *CustomerServiceDeleteCustomerParams) WithBody(body *models.V1DeleteCustomerRequest) *CustomerServiceDeleteCustomerParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the customer service delete customer params
func (o *CustomerServiceDeleteCustomerParams) SetBody(body *models.V1DeleteCustomerRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerServiceDeleteCustomerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
