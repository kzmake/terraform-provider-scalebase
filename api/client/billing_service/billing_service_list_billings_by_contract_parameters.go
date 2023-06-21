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

// NewBillingServiceListBillingsByContractParams creates a new BillingServiceListBillingsByContractParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBillingServiceListBillingsByContractParams() *BillingServiceListBillingsByContractParams {
	return &BillingServiceListBillingsByContractParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewBillingServiceListBillingsByContractParamsWithTimeout creates a new BillingServiceListBillingsByContractParams object
// with the ability to set a timeout on a request.
func NewBillingServiceListBillingsByContractParamsWithTimeout(timeout time.Duration) *BillingServiceListBillingsByContractParams {
	return &BillingServiceListBillingsByContractParams{
		timeout: timeout,
	}
}

// NewBillingServiceListBillingsByContractParamsWithContext creates a new BillingServiceListBillingsByContractParams object
// with the ability to set a context for a request.
func NewBillingServiceListBillingsByContractParamsWithContext(ctx context.Context) *BillingServiceListBillingsByContractParams {
	return &BillingServiceListBillingsByContractParams{
		Context: ctx,
	}
}

// NewBillingServiceListBillingsByContractParamsWithHTTPClient creates a new BillingServiceListBillingsByContractParams object
// with the ability to set a custom HTTPClient for a request.
func NewBillingServiceListBillingsByContractParamsWithHTTPClient(client *http.Client) *BillingServiceListBillingsByContractParams {
	return &BillingServiceListBillingsByContractParams{
		HTTPClient: client,
	}
}

/*
BillingServiceListBillingsByContractParams contains all the parameters to send to the API endpoint

	for the billing service list billings by contract operation.

	Typically these are written to a http.Request.
*/
type BillingServiceListBillingsByContractParams struct {

	// Body.
	Body *models.V1ListBillingsByContractRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the billing service list billings by contract params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingServiceListBillingsByContractParams) WithDefaults() *BillingServiceListBillingsByContractParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the billing service list billings by contract params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingServiceListBillingsByContractParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the billing service list billings by contract params
func (o *BillingServiceListBillingsByContractParams) WithTimeout(timeout time.Duration) *BillingServiceListBillingsByContractParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing service list billings by contract params
func (o *BillingServiceListBillingsByContractParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing service list billings by contract params
func (o *BillingServiceListBillingsByContractParams) WithContext(ctx context.Context) *BillingServiceListBillingsByContractParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing service list billings by contract params
func (o *BillingServiceListBillingsByContractParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing service list billings by contract params
func (o *BillingServiceListBillingsByContractParams) WithHTTPClient(client *http.Client) *BillingServiceListBillingsByContractParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing service list billings by contract params
func (o *BillingServiceListBillingsByContractParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the billing service list billings by contract params
func (o *BillingServiceListBillingsByContractParams) WithBody(body *models.V1ListBillingsByContractRequest) *BillingServiceListBillingsByContractParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the billing service list billings by contract params
func (o *BillingServiceListBillingsByContractParams) SetBody(body *models.V1ListBillingsByContractRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *BillingServiceListBillingsByContractParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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