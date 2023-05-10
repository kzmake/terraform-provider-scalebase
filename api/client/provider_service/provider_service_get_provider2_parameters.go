// Code generated by go-swagger; DO NOT EDIT.

package provider_service

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

// NewProviderServiceGetProvider2Params creates a new ProviderServiceGetProvider2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewProviderServiceGetProvider2Params() *ProviderServiceGetProvider2Params {
	return &ProviderServiceGetProvider2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewProviderServiceGetProvider2ParamsWithTimeout creates a new ProviderServiceGetProvider2Params object
// with the ability to set a timeout on a request.
func NewProviderServiceGetProvider2ParamsWithTimeout(timeout time.Duration) *ProviderServiceGetProvider2Params {
	return &ProviderServiceGetProvider2Params{
		timeout: timeout,
	}
}

// NewProviderServiceGetProvider2ParamsWithContext creates a new ProviderServiceGetProvider2Params object
// with the ability to set a context for a request.
func NewProviderServiceGetProvider2ParamsWithContext(ctx context.Context) *ProviderServiceGetProvider2Params {
	return &ProviderServiceGetProvider2Params{
		Context: ctx,
	}
}

// NewProviderServiceGetProvider2ParamsWithHTTPClient creates a new ProviderServiceGetProvider2Params object
// with the ability to set a custom HTTPClient for a request.
func NewProviderServiceGetProvider2ParamsWithHTTPClient(client *http.Client) *ProviderServiceGetProvider2Params {
	return &ProviderServiceGetProvider2Params{
		HTTPClient: client,
	}
}

/*
ProviderServiceGetProvider2Params contains all the parameters to send to the API endpoint

	for the provider service get provider2 operation.

	Typically these are written to a http.Request.
*/
type ProviderServiceGetProvider2Params struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the provider service get provider2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProviderServiceGetProvider2Params) WithDefaults() *ProviderServiceGetProvider2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the provider service get provider2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ProviderServiceGetProvider2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the provider service get provider2 params
func (o *ProviderServiceGetProvider2Params) WithTimeout(timeout time.Duration) *ProviderServiceGetProvider2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the provider service get provider2 params
func (o *ProviderServiceGetProvider2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the provider service get provider2 params
func (o *ProviderServiceGetProvider2Params) WithContext(ctx context.Context) *ProviderServiceGetProvider2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the provider service get provider2 params
func (o *ProviderServiceGetProvider2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the provider service get provider2 params
func (o *ProviderServiceGetProvider2Params) WithHTTPClient(client *http.Client) *ProviderServiceGetProvider2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the provider service get provider2 params
func (o *ProviderServiceGetProvider2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ProviderServiceGetProvider2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}