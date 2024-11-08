// Code generated by go-swagger; DO NOT EDIT.

package agents

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

// NewChangeNodeExporterParams creates a new ChangeNodeExporterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangeNodeExporterParams() *ChangeNodeExporterParams {
	return &ChangeNodeExporterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangeNodeExporterParamsWithTimeout creates a new ChangeNodeExporterParams object
// with the ability to set a timeout on a request.
func NewChangeNodeExporterParamsWithTimeout(timeout time.Duration) *ChangeNodeExporterParams {
	return &ChangeNodeExporterParams{
		timeout: timeout,
	}
}

// NewChangeNodeExporterParamsWithContext creates a new ChangeNodeExporterParams object
// with the ability to set a context for a request.
func NewChangeNodeExporterParamsWithContext(ctx context.Context) *ChangeNodeExporterParams {
	return &ChangeNodeExporterParams{
		Context: ctx,
	}
}

// NewChangeNodeExporterParamsWithHTTPClient creates a new ChangeNodeExporterParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangeNodeExporterParamsWithHTTPClient(client *http.Client) *ChangeNodeExporterParams {
	return &ChangeNodeExporterParams{
		HTTPClient: client,
	}
}

/*
ChangeNodeExporterParams contains all the parameters to send to the API endpoint

	for the change node exporter operation.

	Typically these are written to a http.Request.
*/
type ChangeNodeExporterParams struct {
	// Body.
	Body ChangeNodeExporterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change node exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeNodeExporterParams) WithDefaults() *ChangeNodeExporterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change node exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeNodeExporterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change node exporter params
func (o *ChangeNodeExporterParams) WithTimeout(timeout time.Duration) *ChangeNodeExporterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change node exporter params
func (o *ChangeNodeExporterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change node exporter params
func (o *ChangeNodeExporterParams) WithContext(ctx context.Context) *ChangeNodeExporterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change node exporter params
func (o *ChangeNodeExporterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change node exporter params
func (o *ChangeNodeExporterParams) WithHTTPClient(client *http.Client) *ChangeNodeExporterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change node exporter params
func (o *ChangeNodeExporterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change node exporter params
func (o *ChangeNodeExporterParams) WithBody(body ChangeNodeExporterBody) *ChangeNodeExporterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change node exporter params
func (o *ChangeNodeExporterParams) SetBody(body ChangeNodeExporterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeNodeExporterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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