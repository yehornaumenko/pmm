// Code generated by go-swagger; DO NOT EDIT.

package services

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

// NewListServicesParams creates a new ListServicesParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewListServicesParams() *ListServicesParams {
	return &ListServicesParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewListServicesParamsWithTimeout creates a new ListServicesParams object
// with the ability to set a timeout on a request.
func NewListServicesParamsWithTimeout(timeout time.Duration) *ListServicesParams {
	return &ListServicesParams{
		timeout: timeout,
	}
}

// NewListServicesParamsWithContext creates a new ListServicesParams object
// with the ability to set a context for a request.
func NewListServicesParamsWithContext(ctx context.Context) *ListServicesParams {
	return &ListServicesParams{
		Context: ctx,
	}
}

// NewListServicesParamsWithHTTPClient creates a new ListServicesParams object
// with the ability to set a custom HTTPClient for a request.
func NewListServicesParamsWithHTTPClient(client *http.Client) *ListServicesParams {
	return &ListServicesParams{
		HTTPClient: client,
	}
}

/*
ListServicesParams contains all the parameters to send to the API endpoint

	for the list services operation.

	Typically these are written to a http.Request.
*/
type ListServicesParams struct {
	// Body.
	Body ListServicesBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the list services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServicesParams) WithDefaults() *ListServicesParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the list services params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ListServicesParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the list services params
func (o *ListServicesParams) WithTimeout(timeout time.Duration) *ListServicesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list services params
func (o *ListServicesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list services params
func (o *ListServicesParams) WithContext(ctx context.Context) *ListServicesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list services params
func (o *ListServicesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list services params
func (o *ListServicesParams) WithHTTPClient(client *http.Client) *ListServicesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list services params
func (o *ListServicesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the list services params
func (o *ListServicesParams) WithBody(body ListServicesBody) *ListServicesParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the list services params
func (o *ListServicesParams) SetBody(body ListServicesBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ListServicesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
