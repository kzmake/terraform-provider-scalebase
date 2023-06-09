// Code generated by go-swagger; DO NOT EDIT.

package catalog_item_service

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

// NewCatalogItemServiceListCatalogItemsParams creates a new CatalogItemServiceListCatalogItemsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogItemServiceListCatalogItemsParams() *CatalogItemServiceListCatalogItemsParams {
	return &CatalogItemServiceListCatalogItemsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogItemServiceListCatalogItemsParamsWithTimeout creates a new CatalogItemServiceListCatalogItemsParams object
// with the ability to set a timeout on a request.
func NewCatalogItemServiceListCatalogItemsParamsWithTimeout(timeout time.Duration) *CatalogItemServiceListCatalogItemsParams {
	return &CatalogItemServiceListCatalogItemsParams{
		timeout: timeout,
	}
}

// NewCatalogItemServiceListCatalogItemsParamsWithContext creates a new CatalogItemServiceListCatalogItemsParams object
// with the ability to set a context for a request.
func NewCatalogItemServiceListCatalogItemsParamsWithContext(ctx context.Context) *CatalogItemServiceListCatalogItemsParams {
	return &CatalogItemServiceListCatalogItemsParams{
		Context: ctx,
	}
}

// NewCatalogItemServiceListCatalogItemsParamsWithHTTPClient creates a new CatalogItemServiceListCatalogItemsParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogItemServiceListCatalogItemsParamsWithHTTPClient(client *http.Client) *CatalogItemServiceListCatalogItemsParams {
	return &CatalogItemServiceListCatalogItemsParams{
		HTTPClient: client,
	}
}

/*
CatalogItemServiceListCatalogItemsParams contains all the parameters to send to the API endpoint

	for the catalog item service list catalog items operation.

	Typically these are written to a http.Request.
*/
type CatalogItemServiceListCatalogItemsParams struct {

	// Body.
	Body *models.V1ListCatalogItemsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the catalog item service list catalog items params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogItemServiceListCatalogItemsParams) WithDefaults() *CatalogItemServiceListCatalogItemsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog item service list catalog items params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogItemServiceListCatalogItemsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog item service list catalog items params
func (o *CatalogItemServiceListCatalogItemsParams) WithTimeout(timeout time.Duration) *CatalogItemServiceListCatalogItemsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog item service list catalog items params
func (o *CatalogItemServiceListCatalogItemsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog item service list catalog items params
func (o *CatalogItemServiceListCatalogItemsParams) WithContext(ctx context.Context) *CatalogItemServiceListCatalogItemsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog item service list catalog items params
func (o *CatalogItemServiceListCatalogItemsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog item service list catalog items params
func (o *CatalogItemServiceListCatalogItemsParams) WithHTTPClient(client *http.Client) *CatalogItemServiceListCatalogItemsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog item service list catalog items params
func (o *CatalogItemServiceListCatalogItemsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the catalog item service list catalog items params
func (o *CatalogItemServiceListCatalogItemsParams) WithBody(body *models.V1ListCatalogItemsRequest) *CatalogItemServiceListCatalogItemsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the catalog item service list catalog items params
func (o *CatalogItemServiceListCatalogItemsParams) SetBody(body *models.V1ListCatalogItemsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogItemServiceListCatalogItemsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
