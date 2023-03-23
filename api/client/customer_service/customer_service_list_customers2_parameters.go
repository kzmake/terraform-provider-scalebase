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
	"github.com/go-openapi/swag"
)

// NewCustomerServiceListCustomers2Params creates a new CustomerServiceListCustomers2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCustomerServiceListCustomers2Params() *CustomerServiceListCustomers2Params {
	return &CustomerServiceListCustomers2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCustomerServiceListCustomers2ParamsWithTimeout creates a new CustomerServiceListCustomers2Params object
// with the ability to set a timeout on a request.
func NewCustomerServiceListCustomers2ParamsWithTimeout(timeout time.Duration) *CustomerServiceListCustomers2Params {
	return &CustomerServiceListCustomers2Params{
		timeout: timeout,
	}
}

// NewCustomerServiceListCustomers2ParamsWithContext creates a new CustomerServiceListCustomers2Params object
// with the ability to set a context for a request.
func NewCustomerServiceListCustomers2ParamsWithContext(ctx context.Context) *CustomerServiceListCustomers2Params {
	return &CustomerServiceListCustomers2Params{
		Context: ctx,
	}
}

// NewCustomerServiceListCustomers2ParamsWithHTTPClient creates a new CustomerServiceListCustomers2Params object
// with the ability to set a custom HTTPClient for a request.
func NewCustomerServiceListCustomers2ParamsWithHTTPClient(client *http.Client) *CustomerServiceListCustomers2Params {
	return &CustomerServiceListCustomers2Params{
		HTTPClient: client,
	}
}

/*
CustomerServiceListCustomers2Params contains all the parameters to send to the API endpoint

	for the customer service list customers2 operation.

	Typically these are written to a http.Request.
*/
type CustomerServiceListCustomers2Params struct {

	/* PageSize.

	   一覧取得する最大数

	   Format: int32
	*/
	PageSize int32

	/* PageToken.

	   一覧取得に使用するトークン
	*/
	PageToken *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the customer service list customers2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerServiceListCustomers2Params) WithDefaults() *CustomerServiceListCustomers2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the customer service list customers2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CustomerServiceListCustomers2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the customer service list customers2 params
func (o *CustomerServiceListCustomers2Params) WithTimeout(timeout time.Duration) *CustomerServiceListCustomers2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the customer service list customers2 params
func (o *CustomerServiceListCustomers2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the customer service list customers2 params
func (o *CustomerServiceListCustomers2Params) WithContext(ctx context.Context) *CustomerServiceListCustomers2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the customer service list customers2 params
func (o *CustomerServiceListCustomers2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the customer service list customers2 params
func (o *CustomerServiceListCustomers2Params) WithHTTPClient(client *http.Client) *CustomerServiceListCustomers2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the customer service list customers2 params
func (o *CustomerServiceListCustomers2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageSize adds the pageSize to the customer service list customers2 params
func (o *CustomerServiceListCustomers2Params) WithPageSize(pageSize int32) *CustomerServiceListCustomers2Params {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the customer service list customers2 params
func (o *CustomerServiceListCustomers2Params) SetPageSize(pageSize int32) {
	o.PageSize = pageSize
}

// WithPageToken adds the pageToken to the customer service list customers2 params
func (o *CustomerServiceListCustomers2Params) WithPageToken(pageToken *string) *CustomerServiceListCustomers2Params {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the customer service list customers2 params
func (o *CustomerServiceListCustomers2Params) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WriteToRequest writes these params to a swagger request
func (o *CustomerServiceListCustomers2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param pageSize
	qrPageSize := o.PageSize
	qPageSize := swag.FormatInt32(qrPageSize)
	if qPageSize != "" {

		if err := r.SetQueryParam("pageSize", qPageSize); err != nil {
			return err
		}
	}

	if o.PageToken != nil {

		// query param pageToken
		var qrPageToken string

		if o.PageToken != nil {
			qrPageToken = *o.PageToken
		}
		qPageToken := qrPageToken
		if qPageToken != "" {

			if err := r.SetQueryParam("pageToken", qPageToken); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}