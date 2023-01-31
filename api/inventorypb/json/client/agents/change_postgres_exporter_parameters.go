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

// NewChangePostgresExporterParams creates a new ChangePostgresExporterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangePostgresExporterParams() *ChangePostgresExporterParams {
	return &ChangePostgresExporterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangePostgresExporterParamsWithTimeout creates a new ChangePostgresExporterParams object
// with the ability to set a timeout on a request.
func NewChangePostgresExporterParamsWithTimeout(timeout time.Duration) *ChangePostgresExporterParams {
	return &ChangePostgresExporterParams{
		timeout: timeout,
	}
}

// NewChangePostgresExporterParamsWithContext creates a new ChangePostgresExporterParams object
// with the ability to set a context for a request.
func NewChangePostgresExporterParamsWithContext(ctx context.Context) *ChangePostgresExporterParams {
	return &ChangePostgresExporterParams{
		Context: ctx,
	}
}

// NewChangePostgresExporterParamsWithHTTPClient creates a new ChangePostgresExporterParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangePostgresExporterParamsWithHTTPClient(client *http.Client) *ChangePostgresExporterParams {
	return &ChangePostgresExporterParams{
		HTTPClient: client,
	}
}

/*
ChangePostgresExporterParams contains all the parameters to send to the API endpoint

	for the change postgres exporter operation.

	Typically these are written to a http.Request.
*/
type ChangePostgresExporterParams struct {
	// Body.
	Body ChangePostgresExporterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change postgres exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangePostgresExporterParams) WithDefaults() *ChangePostgresExporterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change postgres exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangePostgresExporterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change postgres exporter params
func (o *ChangePostgresExporterParams) WithTimeout(timeout time.Duration) *ChangePostgresExporterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change postgres exporter params
func (o *ChangePostgresExporterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change postgres exporter params
func (o *ChangePostgresExporterParams) WithContext(ctx context.Context) *ChangePostgresExporterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change postgres exporter params
func (o *ChangePostgresExporterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change postgres exporter params
func (o *ChangePostgresExporterParams) WithHTTPClient(client *http.Client) *ChangePostgresExporterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change postgres exporter params
func (o *ChangePostgresExporterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change postgres exporter params
func (o *ChangePostgresExporterParams) WithBody(body ChangePostgresExporterBody) *ChangePostgresExporterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change postgres exporter params
func (o *ChangePostgresExporterParams) SetBody(body ChangePostgresExporterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangePostgresExporterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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
