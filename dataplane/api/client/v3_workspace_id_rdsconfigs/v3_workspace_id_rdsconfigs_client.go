// Code generated by go-swagger; DO NOT EDIT.

package v3_workspace_id_rdsconfigs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new v3 workspace id rdsconfigs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for v3 workspace id rdsconfigs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AttachRdsResourceToEnvironments attaches r d s resource to environemnts

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) AttachRdsResourceToEnvironments(params *AttachRdsResourceToEnvironmentsParams) (*AttachRdsResourceToEnvironmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAttachRdsResourceToEnvironmentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "attachRdsResourceToEnvironments",
		Method:             "PUT",
		PathPattern:        "/v3/{workspaceId}/rdsconfigs/{name}/attach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &AttachRdsResourceToEnvironmentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AttachRdsResourceToEnvironmentsOK), nil

}

/*
CreateRdsConfigInWorkspace creates r d s config in workspace

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) CreateRdsConfigInWorkspace(params *CreateRdsConfigInWorkspaceParams) (*CreateRdsConfigInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateRdsConfigInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createRdsConfigInWorkspace",
		Method:             "POST",
		PathPattern:        "/v3/{workspaceId}/rdsconfigs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &CreateRdsConfigInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateRdsConfigInWorkspaceOK), nil

}

/*
DeleteRdsConfigInWorkspace deletes r d s config by name in workspace

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) DeleteRdsConfigInWorkspace(params *DeleteRdsConfigInWorkspaceParams) (*DeleteRdsConfigInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRdsConfigInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRdsConfigInWorkspace",
		Method:             "DELETE",
		PathPattern:        "/v3/{workspaceId}/rdsconfigs/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DeleteRdsConfigInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteRdsConfigInWorkspaceOK), nil

}

/*
DetachRdsResourceFromEnvironments detaches r d s resource from environemnts

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) DetachRdsResourceFromEnvironments(params *DetachRdsResourceFromEnvironmentsParams) (*DetachRdsResourceFromEnvironmentsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDetachRdsResourceFromEnvironmentsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "detachRdsResourceFromEnvironments",
		Method:             "PUT",
		PathPattern:        "/v3/{workspaceId}/rdsconfigs/{name}/detach",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &DetachRdsResourceFromEnvironmentsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DetachRdsResourceFromEnvironmentsOK), nil

}

/*
GetRdsConfigInWorkspace gets r d s config by name in workspace

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) GetRdsConfigInWorkspace(params *GetRdsConfigInWorkspaceParams) (*GetRdsConfigInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRdsConfigInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRdsConfigInWorkspace",
		Method:             "GET",
		PathPattern:        "/v3/{workspaceId}/rdsconfigs/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRdsConfigInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRdsConfigInWorkspaceOK), nil

}

/*
GetRdsRequestFromNameInWorkspace gets request in workspace

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) GetRdsRequestFromNameInWorkspace(params *GetRdsRequestFromNameInWorkspaceParams) (*GetRdsRequestFromNameInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetRdsRequestFromNameInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getRdsRequestFromNameInWorkspace",
		Method:             "GET",
		PathPattern:        "/v3/{workspaceId}/rdsconfigs/{name}/request",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &GetRdsRequestFromNameInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetRdsRequestFromNameInWorkspaceOK), nil

}

/*
ListRdsConfigsByWorkspace lists r d s configs for the given workspace

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) ListRdsConfigsByWorkspace(params *ListRdsConfigsByWorkspaceParams) (*ListRdsConfigsByWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRdsConfigsByWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRdsConfigsByWorkspace",
		Method:             "GET",
		PathPattern:        "/v3/{workspaceId}/rdsconfigs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &ListRdsConfigsByWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ListRdsConfigsByWorkspaceOK), nil

}

/*
TestRdsConnectionInWorkspace tests r d s connectivity

An RDS Configuration describe a connection to an external Relational Database Service that can be used as the Hive Metastore.
*/
func (a *Client) TestRdsConnectionInWorkspace(params *TestRdsConnectionInWorkspaceParams) (*TestRdsConnectionInWorkspaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestRdsConnectionInWorkspaceParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "testRdsConnectionInWorkspace",
		Method:             "POST",
		PathPattern:        "/v3/{workspaceId}/rdsconfigs/testconnect",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http", "https"},
		Params:             params,
		Reader:             &TestRdsConnectionInWorkspaceReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*TestRdsConnectionInWorkspaceOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}