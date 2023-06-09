// Code generated by go-swagger; DO NOT EDIT.

package billing_service

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

// NewBillingServiceGetBillingParams creates a new BillingServiceGetBillingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBillingServiceGetBillingParams() *BillingServiceGetBillingParams {
	return &BillingServiceGetBillingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBillingServiceGetBillingParamsWithTimeout creates a new BillingServiceGetBillingParams object
// with the ability to set a timeout on a request.
func NewBillingServiceGetBillingParamsWithTimeout(timeout time.Duration) *BillingServiceGetBillingParams {
	return &BillingServiceGetBillingParams{
		timeout: timeout,
	}
}

// NewBillingServiceGetBillingParamsWithContext creates a new BillingServiceGetBillingParams object
// with the ability to set a context for a request.
func NewBillingServiceGetBillingParamsWithContext(ctx context.Context) *BillingServiceGetBillingParams {
	return &BillingServiceGetBillingParams{
		Context: ctx,
	}
}

// NewBillingServiceGetBillingParamsWithHTTPClient creates a new BillingServiceGetBillingParams object
// with the ability to set a custom HTTPClient for a request.
func NewBillingServiceGetBillingParamsWithHTTPClient(client *http.Client) *BillingServiceGetBillingParams {
	return &BillingServiceGetBillingParams{
		HTTPClient: client,
	}
}

/*
BillingServiceGetBillingParams contains all the parameters to send to the API endpoint

	for the billing service get billing operation.

	Typically these are written to a http.Request.
*/
type BillingServiceGetBillingParams struct {

	// Body.
	Body *models.V1GetBillingRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the billing service get billing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingServiceGetBillingParams) WithDefaults() *BillingServiceGetBillingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the billing service get billing params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingServiceGetBillingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the billing service get billing params
func (o *BillingServiceGetBillingParams) WithTimeout(timeout time.Duration) *BillingServiceGetBillingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing service get billing params
func (o *BillingServiceGetBillingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing service get billing params
func (o *BillingServiceGetBillingParams) WithContext(ctx context.Context) *BillingServiceGetBillingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing service get billing params
func (o *BillingServiceGetBillingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing service get billing params
func (o *BillingServiceGetBillingParams) WithHTTPClient(client *http.Client) *BillingServiceGetBillingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing service get billing params
func (o *BillingServiceGetBillingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the billing service get billing params
func (o *BillingServiceGetBillingParams) WithBody(body *models.V1GetBillingRequest) *BillingServiceGetBillingParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the billing service get billing params
func (o *BillingServiceGetBillingParams) SetBody(body *models.V1GetBillingRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BillingServiceGetBillingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
