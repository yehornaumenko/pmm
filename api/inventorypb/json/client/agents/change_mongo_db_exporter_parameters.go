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

// NewChangeMongoDBExporterParams creates a new ChangeMongoDBExporterParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChangeMongoDBExporterParams() *ChangeMongoDBExporterParams {
	return &ChangeMongoDBExporterParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChangeMongoDBExporterParamsWithTimeout creates a new ChangeMongoDBExporterParams object
// with the ability to set a timeout on a request.
func NewChangeMongoDBExporterParamsWithTimeout(timeout time.Duration) *ChangeMongoDBExporterParams {
	return &ChangeMongoDBExporterParams{
		timeout: timeout,
	}
}

// NewChangeMongoDBExporterParamsWithContext creates a new ChangeMongoDBExporterParams object
// with the ability to set a context for a request.
func NewChangeMongoDBExporterParamsWithContext(ctx context.Context) *ChangeMongoDBExporterParams {
	return &ChangeMongoDBExporterParams{
		Context: ctx,
	}
}

// NewChangeMongoDBExporterParamsWithHTTPClient creates a new ChangeMongoDBExporterParams object
// with the ability to set a custom HTTPClient for a request.
func NewChangeMongoDBExporterParamsWithHTTPClient(client *http.Client) *ChangeMongoDBExporterParams {
	return &ChangeMongoDBExporterParams{
		HTTPClient: client,
	}
}

/*
ChangeMongoDBExporterParams contains all the parameters to send to the API endpoint

	for the change mongo DB exporter operation.

	Typically these are written to a http.Request.
*/
type ChangeMongoDBExporterParams struct {
	// Body.
	Body ChangeMongoDBExporterBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the change mongo DB exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeMongoDBExporterParams) WithDefaults() *ChangeMongoDBExporterParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the change mongo DB exporter params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChangeMongoDBExporterParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the change mongo DB exporter params
func (o *ChangeMongoDBExporterParams) WithTimeout(timeout time.Duration) *ChangeMongoDBExporterParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the change mongo DB exporter params
func (o *ChangeMongoDBExporterParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the change mongo DB exporter params
func (o *ChangeMongoDBExporterParams) WithContext(ctx context.Context) *ChangeMongoDBExporterParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the change mongo DB exporter params
func (o *ChangeMongoDBExporterParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the change mongo DB exporter params
func (o *ChangeMongoDBExporterParams) WithHTTPClient(client *http.Client) *ChangeMongoDBExporterParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the change mongo DB exporter params
func (o *ChangeMongoDBExporterParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the change mongo DB exporter params
func (o *ChangeMongoDBExporterParams) WithBody(body ChangeMongoDBExporterBody) *ChangeMongoDBExporterParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the change mongo DB exporter params
func (o *ChangeMongoDBExporterParams) SetBody(body ChangeMongoDBExporterBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *ChangeMongoDBExporterParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {
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