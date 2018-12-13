// Code generated by go-swagger; DO NOT EDIT.

package v1mpacks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v1mpacks API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v1mpacks API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
DeleteManagementPack deletes management pack by id

An Apache Ambari Management Pack (Mpack) can bundle multiple service definitions, stack definitions, stack add-on service definitions, view definitions services so that releasing these artifacts don’t enforce an Apache Ambari release.
*/
func (a *Client) DeleteManagementPack(params *DeleteManagementPackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteManagementPackParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteManagementPack",
		Method:             "DELETE",
		PathPattern:        "/v1/mpacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteManagementPackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeletePrivateManagementPack deletes private management pack by name

An Apache Ambari Management Pack (Mpack) can bundle multiple service definitions, stack definitions, stack add-on service definitions, view definitions services so that releasing these artifacts don’t enforce an Apache Ambari release.
*/
func (a *Client) DeletePrivateManagementPack(params *DeletePrivateManagementPackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePrivateManagementPackParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePrivateManagementPack",
		Method:             "DELETE",
		PathPattern:        "/v1/mpacks/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePrivateManagementPackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
DeletePublicManagementPack deletes public owned or private management pack by name

An Apache Ambari Management Pack (Mpack) can bundle multiple service definitions, stack definitions, stack add-on service definitions, view definitions services so that releasing these artifacts don’t enforce an Apache Ambari release.
*/
func (a *Client) DeletePublicManagementPack(params *DeletePublicManagementPackParams) error {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePublicManagementPackParams()
	}

	_, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePublicManagementPack",
		Method:             "DELETE",
		PathPattern:        "/v1/mpacks/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeletePublicManagementPackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return err
	}
	return nil

}

/*
GetManagementPack retrieves management pack by id

An Apache Ambari Management Pack (Mpack) can bundle multiple service definitions, stack definitions, stack add-on service definitions, view definitions services so that releasing these artifacts don’t enforce an Apache Ambari release.
*/
func (a *Client) GetManagementPack(params *GetManagementPackParams) (*GetManagementPackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetManagementPackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getManagementPack",
		Method:             "GET",
		PathPattern:        "/v1/mpacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetManagementPackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetManagementPackOK), nil

}

/*
GetPrivateManagementPack retrieves a private management pack by name

An Apache Ambari Management Pack (Mpack) can bundle multiple service definitions, stack definitions, stack add-on service definitions, view definitions services so that releasing these artifacts don’t enforce an Apache Ambari release.
*/
func (a *Client) GetPrivateManagementPack(params *GetPrivateManagementPackParams) (*GetPrivateManagementPackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateManagementPackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPrivateManagementPack",
		Method:             "GET",
		PathPattern:        "/v1/mpacks/user/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivateManagementPackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivateManagementPackOK), nil

}

/*
GetPrivateManagementPacks retrieves private management packs

An Apache Ambari Management Pack (Mpack) can bundle multiple service definitions, stack definitions, stack add-on service definitions, view definitions services so that releasing these artifacts don’t enforce an Apache Ambari release.
*/
func (a *Client) GetPrivateManagementPacks(params *GetPrivateManagementPacksParams) (*GetPrivateManagementPacksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPrivateManagementPacksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPrivateManagementPacks",
		Method:             "GET",
		PathPattern:        "/v1/mpacks/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPrivateManagementPacksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPrivateManagementPacksOK), nil

}

/*
GetPublicManagementPack retrieves a public or private owned management pack by name

An Apache Ambari Management Pack (Mpack) can bundle multiple service definitions, stack definitions, stack add-on service definitions, view definitions services so that releasing these artifacts don’t enforce an Apache Ambari release.
*/
func (a *Client) GetPublicManagementPack(params *GetPublicManagementPackParams) (*GetPublicManagementPackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicManagementPackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicManagementPack",
		Method:             "GET",
		PathPattern:        "/v1/mpacks/account/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicManagementPackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicManagementPackOK), nil

}

/*
GetPublicManagementPacks retrieves public and private owned management packs

An Apache Ambari Management Pack (Mpack) can bundle multiple service definitions, stack definitions, stack add-on service definitions, view definitions services so that releasing these artifacts don’t enforce an Apache Ambari release.
*/
func (a *Client) GetPublicManagementPacks(params *GetPublicManagementPacksParams) (*GetPublicManagementPacksOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPublicManagementPacksParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPublicManagementPacks",
		Method:             "GET",
		PathPattern:        "/v1/mpacks/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetPublicManagementPacksReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPublicManagementPacksOK), nil

}

/*
PostPrivateManagementPack creates management pack as private resource

An Apache Ambari Management Pack (Mpack) can bundle multiple service definitions, stack definitions, stack add-on service definitions, view definitions services so that releasing these artifacts don’t enforce an Apache Ambari release.
*/
func (a *Client) PostPrivateManagementPack(params *PostPrivateManagementPackParams) (*PostPrivateManagementPackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPrivateManagementPackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postPrivateManagementPack",
		Method:             "POST",
		PathPattern:        "/v1/mpacks/user",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPrivateManagementPackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPrivateManagementPackOK), nil

}

/*
PostPublicManagementPack creates management pack as public resource

An Apache Ambari Management Pack (Mpack) can bundle multiple service definitions, stack definitions, stack add-on service definitions, view definitions services so that releasing these artifacts don’t enforce an Apache Ambari release.
*/
func (a *Client) PostPublicManagementPack(params *PostPublicManagementPackParams) (*PostPublicManagementPackOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPostPublicManagementPackParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "postPublicManagementPack",
		Method:             "POST",
		PathPattern:        "/v1/mpacks/account",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &PostPublicManagementPackReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PostPublicManagementPackOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}