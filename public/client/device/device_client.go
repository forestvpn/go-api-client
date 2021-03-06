// Code generated by go-swagger; DO NOT EDIT.

package device

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new device API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for device API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BindingList(params *BindingListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BindingListOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  BindingList devices bindings

  Each device has bindings to specific servers. This method returns hostnames where this device
bonded at a particular time.

*/
func (a *Client) BindingList(params *BindingListParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BindingListOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBindingListParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "BindingList",
		Method:             "GET",
		PathPattern:        "/devices/{deviceID}/bindings/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &BindingListReader{formats: a.formats},
		AuthInfo:           authInfo,
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
	success, ok := result.(*BindingListOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*BindingListDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
