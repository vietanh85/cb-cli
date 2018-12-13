// Code generated by go-swagger; DO NOT EDIT.

package tenants

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new tenants API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for tenants API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DisableTenant disables tenant
*/
func (a *Client) DisableTenant(params *DisableTenantParams) (*DisableTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisableTenantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "disableTenant",
		Method:             "PUT",
		PathPattern:        "/caas/api/tenants/{tenantName}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DisableTenantReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DisableTenantOK), nil

}

/*
GetAllTenants lists tenants
*/
func (a *Client) GetAllTenants(params *GetAllTenantsParams) (*GetAllTenantsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetAllTenantsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getAllTenants",
		Method:             "GET",
		PathPattern:        "/caas/api/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetAllTenantsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetAllTenantsOK), nil

}

/*
RegisterTenant registers new tenant
*/
func (a *Client) RegisterTenant(params *RegisterTenantParams) (*RegisterTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRegisterTenantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "registerTenant",
		Method:             "POST",
		PathPattern:        "/caas/api/tenants",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RegisterTenantReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RegisterTenantOK), nil

}

/*
ResendMail resends mail
*/
func (a *Client) ResendMail(params *ResendMailParams) (*ResendMailOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewResendMailParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "resendMail",
		Method:             "POST",
		PathPattern:        "/caas/api/tenants/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ResendMailReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ResendMailOK), nil

}

/*
RetrieveTenant retrieves tenant
*/
func (a *Client) RetrieveTenant(params *RetrieveTenantParams) (*RetrieveTenantOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRetrieveTenantParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "retrieveTenant",
		Method:             "GET",
		PathPattern:        "/caas/api/tenants/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &RetrieveTenantReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*RetrieveTenantOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}