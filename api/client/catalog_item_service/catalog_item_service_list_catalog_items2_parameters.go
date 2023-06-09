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
	"github.com/go-openapi/swag"
)

// NewCatalogItemServiceListCatalogItems2Params creates a new CatalogItemServiceListCatalogItems2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogItemServiceListCatalogItems2Params() *CatalogItemServiceListCatalogItems2Params {
	return &CatalogItemServiceListCatalogItems2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogItemServiceListCatalogItems2ParamsWithTimeout creates a new CatalogItemServiceListCatalogItems2Params object
// with the ability to set a timeout on a request.
func NewCatalogItemServiceListCatalogItems2ParamsWithTimeout(timeout time.Duration) *CatalogItemServiceListCatalogItems2Params {
	return &CatalogItemServiceListCatalogItems2Params{
		timeout: timeout,
	}
}

// NewCatalogItemServiceListCatalogItems2ParamsWithContext creates a new CatalogItemServiceListCatalogItems2Params object
// with the ability to set a context for a request.
func NewCatalogItemServiceListCatalogItems2ParamsWithContext(ctx context.Context) *CatalogItemServiceListCatalogItems2Params {
	return &CatalogItemServiceListCatalogItems2Params{
		Context: ctx,
	}
}

// NewCatalogItemServiceListCatalogItems2ParamsWithHTTPClient creates a new CatalogItemServiceListCatalogItems2Params object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogItemServiceListCatalogItems2ParamsWithHTTPClient(client *http.Client) *CatalogItemServiceListCatalogItems2Params {
	return &CatalogItemServiceListCatalogItems2Params{
		HTTPClient: client,
	}
}

/*
CatalogItemServiceListCatalogItems2Params contains all the parameters to send to the API endpoint

	for the catalog item service list catalog items2 operation.

	Typically these are written to a http.Request.
*/
type CatalogItemServiceListCatalogItems2Params struct {

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

// WithDefaults hydrates default values in the catalog item service list catalog items2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogItemServiceListCatalogItems2Params) WithDefaults() *CatalogItemServiceListCatalogItems2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog item service list catalog items2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogItemServiceListCatalogItems2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog item service list catalog items2 params
func (o *CatalogItemServiceListCatalogItems2Params) WithTimeout(timeout time.Duration) *CatalogItemServiceListCatalogItems2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog item service list catalog items2 params
func (o *CatalogItemServiceListCatalogItems2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog item service list catalog items2 params
func (o *CatalogItemServiceListCatalogItems2Params) WithContext(ctx context.Context) *CatalogItemServiceListCatalogItems2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog item service list catalog items2 params
func (o *CatalogItemServiceListCatalogItems2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog item service list catalog items2 params
func (o *CatalogItemServiceListCatalogItems2Params) WithHTTPClient(client *http.Client) *CatalogItemServiceListCatalogItems2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog item service list catalog items2 params
func (o *CatalogItemServiceListCatalogItems2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPageSize adds the pageSize to the catalog item service list catalog items2 params
func (o *CatalogItemServiceListCatalogItems2Params) WithPageSize(pageSize int32) *CatalogItemServiceListCatalogItems2Params {
	o.SetPageSize(pageSize)
	return o
}

// SetPageSize adds the pageSize to the catalog item service list catalog items2 params
func (o *CatalogItemServiceListCatalogItems2Params) SetPageSize(pageSize int32) {
	o.PageSize = pageSize
}

// WithPageToken adds the pageToken to the catalog item service list catalog items2 params
func (o *CatalogItemServiceListCatalogItems2Params) WithPageToken(pageToken *string) *CatalogItemServiceListCatalogItems2Params {
	o.SetPageToken(pageToken)
	return o
}

// SetPageToken adds the pageToken to the catalog item service list catalog items2 params
func (o *CatalogItemServiceListCatalogItems2Params) SetPageToken(pageToken *string) {
	o.PageToken = pageToken
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogItemServiceListCatalogItems2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
