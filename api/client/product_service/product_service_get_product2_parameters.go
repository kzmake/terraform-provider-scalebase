// Code generated by go-swagger; DO NOT EDIT.

package product_service

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

// NewProductServiceGetProduct2Params creates a new ProductServiceGetProduct2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProductServiceGetProduct2Params() *ProductServiceGetProduct2Params {
	return &ProductServiceGetProduct2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewProductServiceGetProduct2ParamsWithTimeout creates a new ProductServiceGetProduct2Params object
// with the ability to set a timeout on a request.
func NewProductServiceGetProduct2ParamsWithTimeout(timeout time.Duration) *ProductServiceGetProduct2Params {
	return &ProductServiceGetProduct2Params{
		timeout: timeout,
	}
}

// NewProductServiceGetProduct2ParamsWithContext creates a new ProductServiceGetProduct2Params object
// with the ability to set a context for a request.
func NewProductServiceGetProduct2ParamsWithContext(ctx context.Context) *ProductServiceGetProduct2Params {
	return &ProductServiceGetProduct2Params{
		Context: ctx,
	}
}

// NewProductServiceGetProduct2ParamsWithHTTPClient creates a new ProductServiceGetProduct2Params object
// with the ability to set a custom HTTPClient for a request.
func NewProductServiceGetProduct2ParamsWithHTTPClient(client *http.Client) *ProductServiceGetProduct2Params {
	return &ProductServiceGetProduct2Params{
		HTTPClient: client,
	}
}

/*
ProductServiceGetProduct2Params contains all the parameters to send to the API endpoint

	for the product service get product2 operation.

	Typically these are written to a http.Request.
*/
type ProductServiceGetProduct2Params struct {

	/* ID.

	   プロダクトID
	*/
	ID string

	/* OptionalID.

	   プロダクト管理ID(プロダクトIDの代わりに指定可)
	*/
	OptionalID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the product service get product2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductServiceGetProduct2Params) WithDefaults() *ProductServiceGetProduct2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the product service get product2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProductServiceGetProduct2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the product service get product2 params
func (o *ProductServiceGetProduct2Params) WithTimeout(timeout time.Duration) *ProductServiceGetProduct2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the product service get product2 params
func (o *ProductServiceGetProduct2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the product service get product2 params
func (o *ProductServiceGetProduct2Params) WithContext(ctx context.Context) *ProductServiceGetProduct2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the product service get product2 params
func (o *ProductServiceGetProduct2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the product service get product2 params
func (o *ProductServiceGetProduct2Params) WithHTTPClient(client *http.Client) *ProductServiceGetProduct2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the product service get product2 params
func (o *ProductServiceGetProduct2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the product service get product2 params
func (o *ProductServiceGetProduct2Params) WithID(id string) *ProductServiceGetProduct2Params {
	o.SetID(id)
	return o
}

// SetID adds the id to the product service get product2 params
func (o *ProductServiceGetProduct2Params) SetID(id string) {
	o.ID = id
}

// WithOptionalID adds the optionalID to the product service get product2 params
func (o *ProductServiceGetProduct2Params) WithOptionalID(optionalID *string) *ProductServiceGetProduct2Params {
	o.SetOptionalID(optionalID)
	return o
}

// SetOptionalID adds the optionalId to the product service get product2 params
func (o *ProductServiceGetProduct2Params) SetOptionalID(optionalID *string) {
	o.OptionalID = optionalID
}

// WriteToRequest writes these params to a swagger request
func (o *ProductServiceGetProduct2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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