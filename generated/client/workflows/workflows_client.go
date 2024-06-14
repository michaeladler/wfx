// Code generated by go-swagger; DO NOT EDIT.

// SPDX-FileCopyrightText: 2023 Siemens AG
//
// SPDX-License-Identifier: Apache-2.0
//

package workflows

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// New creates a new workflows API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

// New creates a new workflows API client with basic auth credentials.
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

// New creates a new workflows API client with a bearer token for authentication.
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
Client for workflows API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption may be used to customize the behavior of Client methods.
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteWorkflowsName(params *DeleteWorkflowsNameParams, opts ...ClientOption) (*DeleteWorkflowsNameNoContent, error)

	GetWorkflows(params *GetWorkflowsParams, opts ...ClientOption) (*GetWorkflowsOK, error)

	GetWorkflowsName(params *GetWorkflowsNameParams, opts ...ClientOption) (*GetWorkflowsNameOK, error)

	PostWorkflows(params *PostWorkflowsParams, opts ...ClientOption) (*PostWorkflowsCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
DeleteWorkflowsName deletes an existing workflow

Delete an existing workflow
*/
func (a *Client) DeleteWorkflowsName(params *DeleteWorkflowsNameParams, opts ...ClientOption) (*DeleteWorkflowsNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteWorkflowsNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "DeleteWorkflowsName",
		Method:             "DELETE",
		PathPattern:        "/workflows/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteWorkflowsNameReader{formats: a.formats},
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
	success, ok := result.(*DeleteWorkflowsNameNoContent)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*DeleteWorkflowsNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetWorkflows lists of available workflows

List of available workflows
*/
func (a *Client) GetWorkflows(params *GetWorkflowsParams, opts ...ClientOption) (*GetWorkflowsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWorkflows",
		Method:             "GET",
		PathPattern:        "/workflows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWorkflowsReader{formats: a.formats},
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
	success, ok := result.(*GetWorkflowsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetWorkflowsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetWorkflowsName workflows description for a given name

Workflow description for a given name
*/
func (a *Client) GetWorkflowsName(params *GetWorkflowsNameParams, opts ...ClientOption) (*GetWorkflowsNameOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetWorkflowsNameParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "GetWorkflowsName",
		Method:             "GET",
		PathPattern:        "/workflows/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetWorkflowsNameReader{formats: a.formats},
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
	success, ok := result.(*GetWorkflowsNameOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetWorkflowsNameDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
PostWorkflows adds a new workflow

Add a new workflow
*/
func (a *Client) PostWorkflows(params *PostWorkflowsParams, opts ...ClientOption) (*PostWorkflowsCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostWorkflowsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "PostWorkflows",
		Method:             "POST",
		PathPattern:        "/workflows",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PostWorkflowsReader{formats: a.formats},
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
	success, ok := result.(*PostWorkflowsCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*PostWorkflowsDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
