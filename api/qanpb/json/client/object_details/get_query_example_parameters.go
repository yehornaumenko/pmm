// Code generated by go-swagger; DO NOT EDIT.

package object_details

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

// NewGetQueryExampleParams creates a new GetQueryExampleParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetQueryExampleParams() *GetQueryExampleParams {
	return &GetQueryExampleParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetQueryExampleParamsWithTimeout creates a new GetQueryExampleParams object
// with the ability to set a timeout on a request.
func NewGetQueryExampleParamsWithTimeout(timeout time.Duration) *GetQueryExampleParams {
	return &GetQueryExampleParams{
		timeout: timeout,
	}
}

// NewGetQueryExampleParamsWithContext creates a new GetQueryExampleParams object
// with the ability to set a context for a request.
func NewGetQueryExampleParamsWithContext(ctx context.Context) *GetQueryExampleParams {
	return &GetQueryExampleParams{
		Context: ctx,
	}
}

// NewGetQueryExampleParamsWithHTTPClient creates a new GetQueryExampleParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetQueryExampleParamsWithHTTPClient(client *http.Client) *GetQueryExampleParams {
	return &GetQueryExampleParams{
		HTTPClient: client,
	}
}

/*
GetQueryExampleParams contains all the parameters to send to the API endpoint

	for the get query example operation.

	Typically these are written to a http.Request.
*/
type GetQueryExampleParams struct {
	/* Body.

	     QueryExampleRequest defines filtering of query examples for specific value of
	dimension (ex.: host=hostname1 or queryid=1D410B4BE5060972.
	*/
	Body GetQueryExampleBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get query example params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueryExampleParams) WithDefaults() *GetQueryExampleParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get query example params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetQueryExampleParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get query example params
func (o *GetQueryExampleParams) WithTimeout(timeout time.Duration) *GetQueryExampleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get query example params
func (o *GetQueryExampleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get query example params
func (o *GetQueryExampleParams) WithContext(ctx context.Context) *GetQueryExampleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get query example params
func (o *GetQueryExampleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get query example params
func (o *GetQueryExampleParams) WithHTTPClient(client *http.Client) *GetQueryExampleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get query example params
func (o *GetQueryExampleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the get query example params
func (o *GetQueryExampleParams) WithBody(body GetQueryExampleBody) *GetQueryExampleParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the get query example params
func (o *GetQueryExampleParams) SetBody(body GetQueryExampleBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *GetQueryExampleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
