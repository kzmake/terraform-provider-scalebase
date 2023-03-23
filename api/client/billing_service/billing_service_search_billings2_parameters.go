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
	"github.com/go-openapi/swag"
)

// NewBillingServiceSearchBillings2Params creates a new BillingServiceSearchBillings2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewBillingServiceSearchBillings2Params() *BillingServiceSearchBillings2Params {
	return &BillingServiceSearchBillings2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewBillingServiceSearchBillings2ParamsWithTimeout creates a new BillingServiceSearchBillings2Params object
// with the ability to set a timeout on a request.
func NewBillingServiceSearchBillings2ParamsWithTimeout(timeout time.Duration) *BillingServiceSearchBillings2Params {
	return &BillingServiceSearchBillings2Params{
		timeout: timeout,
	}
}

// NewBillingServiceSearchBillings2ParamsWithContext creates a new BillingServiceSearchBillings2Params object
// with the ability to set a context for a request.
func NewBillingServiceSearchBillings2ParamsWithContext(ctx context.Context) *BillingServiceSearchBillings2Params {
	return &BillingServiceSearchBillings2Params{
		Context: ctx,
	}
}

// NewBillingServiceSearchBillings2ParamsWithHTTPClient creates a new BillingServiceSearchBillings2Params object
// with the ability to set a custom HTTPClient for a request.
func NewBillingServiceSearchBillings2ParamsWithHTTPClient(client *http.Client) *BillingServiceSearchBillings2Params {
	return &BillingServiceSearchBillings2Params{
		HTTPClient: client,
	}
}

/*
BillingServiceSearchBillings2Params contains all the parameters to send to the API endpoint

	for the billing service search billings2 operation.

	Typically these are written to a http.Request.
*/
type BillingServiceSearchBillings2Params struct {

	/* PageSize.

	   一覧取得する最大数

	   Format: int32
	*/
	PageSize int32

	/* PageToken.

	   一覧取得に使用するトークン
	*/
	PageToken *string

	/* Query.

	     検索クエリ
	example:
	```
	status=STATUS_UNBILLED
	```

	status:
	[`any`|`*`|下記のパターン] が指定可能
	```
	STATUS_UNSPECIFIED
	STATUS_HOLD_INVOICE_CREATION
	STATUS_UNBILLED
	STATUS_POSTED
	STATUS_PARTIALLY_PAID
	STATUS_PAID
	STATUS_PAYMENT_EXCLUDED
	STATUS_PAYMENT_DUE
	STATUS_PAYMENT_UNNECESSARY
	STATUS_PARTIALLY_UNPAID
	STATUS_PAYMENT_REFUNDED
	STATUS_PAID_WITH_CARD
	STATUS_INVOICE_CREATED
	STATUS_INVOICE_DELETED
	STATUS_INVALIDATED
	STATUS_ANY
	```
	*/
	Query *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the billing service search billings2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingServiceSearchBillings2Params) WithDefaults() *BillingServiceSearchBillings2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the billing service search billings2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *BillingServiceSearchBillings2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the billing service search billings2 params
func (o *BillingServiceSearchBillings2Params) WithTimeout(timeout time.Duration) *BillingServiceSearchBillings2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the billing service search billings2 params
func (o *BillingServiceSearchBillings2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the billing service search billings2 params
func (o *BillingServiceSearchBillings2Params) WithContext(ctx context.Context) *BillingServiceSearchBillings2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the billing service search billings2 params
func (o *BillingServiceSearchBillings2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the billing service search billings2 params
func (o *BillingServiceSearchBillings2Params) WithHTTPClient(client *http.Client) *BillingServiceSearchBillings2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the billing service search billings2 params
func (o *BillingServiceSearchBillings2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageSize adds the pageSize to the billing service search billings2 params
func (o *BillingServiceSearchBillings2Params) WithPageSize(pageSize int32) *BillingServiceSearchBillings2Params {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the billing service search billings2 params
func (o *BillingServiceSearchBillings2Params) SetPageSize(pageSize int32) {
	o.PageSize = pageSize
}

// WithPageToken adds the pageToken to the billing service search billings2 params
func (o *BillingServiceSearchBillings2Params) WithPageToken(pageToken *string) *BillingServiceSearchBillings2Params {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the billing service search billings2 params
func (o *BillingServiceSearchBillings2Params) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WithQuery adds the query to the billing service search billings2 params
func (o *BillingServiceSearchBillings2Params) WithQuery(query *string) *BillingServiceSearchBillings2Params {
	o.SetQuery(query)
	return o
}

// SetQuery adds the query to the billing service search billings2 params
func (o *BillingServiceSearchBillings2Params) SetQuery(query *string) {
	o.Query = query
}

// WriteToRequest writes these params to a swagger request
func (o *BillingServiceSearchBillings2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Query != nil {

		// query param query
		var qrQuery string

		if o.Query != nil {
			qrQuery = *o.Query
		}
		qQuery := qrQuery
		if qQuery != "" {

			if err := r.SetQueryParam("query", qQuery); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
