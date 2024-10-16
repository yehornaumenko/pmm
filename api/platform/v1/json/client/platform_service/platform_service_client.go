// Code generated by go-swagger; DO NOT EDIT.

package platform_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new platform service API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for platform service API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	Connect(params *ConnectParams, opts ...ClientOption) (*ConnectOK, error)

	Disconnect(params *DisconnectParams, opts ...ClientOption) (*DisconnectOK, error)

	GetContactInformation(params *GetContactInformationParams, opts ...ClientOption) (*GetContactInformationOK, error)

	SearchOrganizationEntitlements(params *SearchOrganizationEntitlementsParams, opts ...ClientOption) (*SearchOrganizationEntitlementsOK, error)

	SearchOrganizationTickets(params *SearchOrganizationTicketsParams, opts ...ClientOption) (*SearchOrganizationTicketsOK, error)

	ServerInfo(params *ServerInfoParams, opts ...ClientOption) (*ServerInfoOK, error)

	UserStatus(params *UserStatusParams, opts ...ClientOption) (*UserStatusOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
Connect connects PMM server

Connect a PMM server to the organization created on Percona Portal. That allows the user to sign in to the PMM server with their Percona Account.
*/
func (a *Client) Connect(params *ConnectParams, opts ...ClientOption) (*ConnectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConnectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Connect",
		Method:             "POST",
		PathPattern:        "/v1/platform:connect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ConnectReader{formats: a.formats},
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
	success, ok := result.(*ConnectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ConnectDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
Disconnect disconnects PMM server

Disconnect a PMM server from the organization created on Percona Portal.
*/
func (a *Client) Disconnect(params *DisconnectParams, opts ...ClientOption) (*DisconnectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisconnectParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "Disconnect",
		Method:             "POST",
		PathPattern:        "/v1/platform:disconnect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisconnectReader{formats: a.formats},
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
	success, ok := result.(*DisconnectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DisconnectDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetContactInformation gets contact information

Fetch the contact details of the customer success employee handling the Percona customer account.
*/
func (a *Client) GetContactInformation(params *GetContactInformationParams, opts ...ClientOption) (*GetContactInformationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetContactInformationParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetContactInformation",
		Method:             "GET",
		PathPattern:        "/v1/platform/contact",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetContactInformationReader{formats: a.formats},
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
	success, ok := result.(*GetContactInformationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetContactInformationDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchOrganizationEntitlements searches organization entitlements

Fetch entitlements available to the Portal organization that the PMM server is connected to.
*/
func (a *Client) SearchOrganizationEntitlements(params *SearchOrganizationEntitlementsParams, opts ...ClientOption) (*SearchOrganizationEntitlementsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchOrganizationEntitlementsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchOrganizationEntitlements",
		Method:             "GET",
		PathPattern:        "/v1/platform/organization/entitlements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchOrganizationEntitlementsReader{formats: a.formats},
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
	success, ok := result.(*SearchOrganizationEntitlementsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchOrganizationEntitlementsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
SearchOrganizationTickets searches organization tickets

Fetch support tickets belonging to the Percona Portal Organization that the PMM server is connected to.
*/
func (a *Client) SearchOrganizationTickets(params *SearchOrganizationTicketsParams, opts ...ClientOption) (*SearchOrganizationTicketsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSearchOrganizationTicketsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "SearchOrganizationTickets",
		Method:             "GET",
		PathPattern:        "/v1/platform/organization/tickets",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &SearchOrganizationTicketsReader{formats: a.formats},
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
	success, ok := result.(*SearchOrganizationTicketsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*SearchOrganizationTicketsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
ServerInfo gets server info

Return PMM server ID and name.
*/
func (a *Client) ServerInfo(params *ServerInfoParams, opts ...ClientOption) (*ServerInfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewServerInfoParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "ServerInfo",
		Method:             "GET",
		PathPattern:        "/v1/platform/server",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ServerInfoReader{formats: a.formats},
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
	success, ok := result.(*ServerInfoOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*ServerInfoDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
UserStatus gets user status

Check if the current user is logged in with their Percona Account.
*/
func (a *Client) UserStatus(params *UserStatusParams, opts ...ClientOption) (*UserStatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUserStatusParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "UserStatus",
		Method:             "GET",
		PathPattern:        "/v1/platform/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &UserStatusReader{formats: a.formats},
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
	success, ok := result.(*UserStatusOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*UserStatusDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}