// Code generated by go-swagger; DO NOT EDIT.

package actions_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new actions service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new actions service API client with basic auth credentials.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - user: user for basic authentication header.
// - password: password for basic authentication header.
func NewClientWithBasicAuth(host, basePath, scheme, user, password string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BasicAuth(user, password)
	return &Client{transport: transport, formats: strfmt.Default}
}

// New creates a new actions service API client with a bearer token for authentication.
// It takes the following parameters:
// - host: http host (github.com).
// - basePath: any base path for the API client ("/v1", "/v3").
// - scheme: http scheme ("http", "https").
// - bearerToken: bearer token for Bearer authentication header.
func NewClientWithBearerToken(host, basePath, scheme, bearerToken string) ClientService {
	transport := httptransport.New(host, basePath, []string{scheme})
	transport.DefaultAuthentication = httptransport.BearerToken(bearerToken)
	return &Client{transport: transport, formats: strfmt.Default}
}

/*
Client for actions service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	CancelAction(params *CancelActionParams, opts ...ClientOption) (*CancelActionOK, error)

	GetAction(params *GetActionParams, opts ...ClientOption) (*GetActionOK, error)

	StartPTSummaryAction(params *StartPTSummaryActionParams, opts ...ClientOption) (*StartPTSummaryActionOK, error)

	StartServiceAction(params *StartServiceActionParams, opts ...ClientOption) (*StartServiceActionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
CancelAction cancels an action

Stops an Action.
*/
func (a *Client) CancelAction(params *CancelActionParams, opts ...ClientOption) (*CancelActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCancelActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "CancelAction",
		Method:             "POST",
		PathPattern:        "/v1/actions:cancelAction",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CancelActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CancelActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CancelActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetAction gets action

Gets the result of a given Action.
*/
func (a *Client) GetAction(params *GetActionParams, opts ...ClientOption) (*GetActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetAction",
		Method:             "GET",
		PathPattern:        "/v1/actions/{action_id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
StartPTSummaryAction starts PT summary action

Starts 'Percona Toolkit Summary' Action.
*/
func (a *Client) StartPTSummaryAction(params *StartPTSummaryActionParams, opts ...ClientOption) (*StartPTSummaryActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartPTSummaryActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StartPTSummaryAction",
		Method:             "POST",
		PathPattern:        "/v1/actions:startNodeAction",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartPTSummaryActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartPTSummaryActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartPTSummaryActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
StartServiceAction starts a service action

Starts a Service Action.
*/
func (a *Client) StartServiceAction(params *StartServiceActionParams, opts ...ClientOption) (*StartServiceActionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStartServiceActionParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "StartServiceAction",
		Method:             "POST",
		PathPattern:        "/v1/actions:startServiceAction",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &StartServiceActionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*StartServiceActionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*StartServiceActionDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}