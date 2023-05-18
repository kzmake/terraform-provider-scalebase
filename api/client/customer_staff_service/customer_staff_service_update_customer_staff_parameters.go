// Code generated by go-swagger; DO NOT EDIT.

package customer_staff_service

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

// NewCustomerStaffServiceUpdateCustomerStaffParams creates a new CustomerStaffServiceUpdateCustomerStaffParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerStaffServiceUpdateCustomerStaffParams() *CustomerStaffServiceUpdateCustomerStaffParams {
	return &CustomerStaffServiceUpdateCustomerStaffParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerStaffServiceUpdateCustomerStaffParamsWithTimeout creates a new CustomerStaffServiceUpdateCustomerStaffParams object
// with the ability to set a timeout on a request.
func NewCustomerStaffServiceUpdateCustomerStaffParamsWithTimeout(timeout time.Duration) *CustomerStaffServiceUpdateCustomerStaffParams {
	return &CustomerStaffServiceUpdateCustomerStaffParams{
		timeout: timeout,
	}
}

// NewCustomerStaffServiceUpdateCustomerStaffParamsWithContext creates a new CustomerStaffServiceUpdateCustomerStaffParams object
// with the ability to set a context for a request.
func NewCustomerStaffServiceUpdateCustomerStaffParamsWithContext(ctx context.Context) *CustomerStaffServiceUpdateCustomerStaffParams {
	return &CustomerStaffServiceUpdateCustomerStaffParams{
		Context: ctx,
	}
}

// NewCustomerStaffServiceUpdateCustomerStaffParamsWithHTTPClient creates a new CustomerStaffServiceUpdateCustomerStaffParams object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerStaffServiceUpdateCustomerStaffParamsWithHTTPClient(client *http.Client) *CustomerStaffServiceUpdateCustomerStaffParams {
	return &CustomerStaffServiceUpdateCustomerStaffParams{
		HTTPClient: client,
	}
}

/*
CustomerStaffServiceUpdateCustomerStaffParams contains all the parameters to send to the API endpoint

	for the customer staff service update customer staff operation.

	Typically these are written to a http.Request.
*/
type CustomerStaffServiceUpdateCustomerStaffParams struct {

	/* CustomerStaff.

	   顧客担当者
	*/
	CustomerStaff *models.Publicv1CustomerStaff

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer staff service update customer staff params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerStaffServiceUpdateCustomerStaffParams) WithDefaults() *CustomerStaffServiceUpdateCustomerStaffParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer staff service update customer staff params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerStaffServiceUpdateCustomerStaffParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer staff service update customer staff params
func (o *CustomerStaffServiceUpdateCustomerStaffParams) WithTimeout(timeout time.Duration) *CustomerStaffServiceUpdateCustomerStaffParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer staff service update customer staff params
func (o *CustomerStaffServiceUpdateCustomerStaffParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer staff service update customer staff params
func (o *CustomerStaffServiceUpdateCustomerStaffParams) WithContext(ctx context.Context) *CustomerStaffServiceUpdateCustomerStaffParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer staff service update customer staff params
func (o *CustomerStaffServiceUpdateCustomerStaffParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer staff service update customer staff params
func (o *CustomerStaffServiceUpdateCustomerStaffParams) WithHTTPClient(client *http.Client) *CustomerStaffServiceUpdateCustomerStaffParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer staff service update customer staff params
func (o *CustomerStaffServiceUpdateCustomerStaffParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCustomerStaff adds the customerStaff to the customer staff service update customer staff params
func (o *CustomerStaffServiceUpdateCustomerStaffParams) WithCustomerStaff(customerStaff *models.Publicv1CustomerStaff) *CustomerStaffServiceUpdateCustomerStaffParams {
	o.SetCustomerStaff(customerStaff)
	return o
}

// SetCustomerStaff adds the customerStaff to the customer staff service update customer staff params
func (o *CustomerStaffServiceUpdateCustomerStaffParams) SetCustomerStaff(customerStaff *models.Publicv1CustomerStaff) {
	o.CustomerStaff = customerStaff
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerStaffServiceUpdateCustomerStaffParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if o.CustomerStaff != nil {
		if err := r.SetBodyParam(o.CustomerStaff); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}