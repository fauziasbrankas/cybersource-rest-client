// Code generated by go-swagger; DO NOT EDIT.

package net_fundings

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new net fundings API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for net fundings API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
GetNetFundingDetails gets netfunding information for an account or a merchant

Get Netfunding information for an account or a merchant.
*/
func (a *Client) GetNetFundingDetails(params *GetNetFundingDetailsParams) (*GetNetFundingDetailsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetNetFundingDetailsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getNetFundingDetails",
		Method:             "GET",
		PathPattern:        "/reporting/v3/net-fundings",
		ProducesMediaTypes: []string{"application/hal+json", "application/xml"},
		ConsumesMediaTypes: []string{"application/json;charset=utf-8"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetNetFundingDetailsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetNetFundingDetailsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getNetFundingDetails: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}