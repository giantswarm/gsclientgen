// Code generated by go-swagger; DO NOT EDIT.

package users

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new users API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for users API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateUser creates user

Creates a users in the system. Currently this endpoint is only available to users with admin permissions.

*/
func (a *Client) CreateUser(params *CreateUserParams, authInfo runtime.ClientAuthInfoWriter) (*CreateUserCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createUser",
		Method:             "PUT",
		PathPattern:        "/v4/users/{email}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateUserCreated), nil

}

/*
DeleteUser deletes user

Deletes a users in the system. Currently this endpoint is only available
to users with admin permissions.

*/
func (a *Client) DeleteUser(params *DeleteUserParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteUser",
		Method:             "DELETE",
		PathPattern:        "/v4/users/{email}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteUserOK), nil

}

/*
GetCurrentUser gets current user

Returns details about the currently authenticated user

*/
func (a *Client) GetCurrentUser(params *GetCurrentUserParams, authInfo runtime.ClientAuthInfoWriter) (*GetCurrentUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetCurrentUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getCurrentUser",
		Method:             "GET",
		PathPattern:        "/v4/user/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetCurrentUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetCurrentUserOK), nil

}

/*
GetUser gets user

Returns details about a specific user

*/
func (a *Client) GetUser(params *GetUserParams, authInfo runtime.ClientAuthInfoWriter) (*GetUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUser",
		Method:             "GET",
		PathPattern:        "/v4/users/{email}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUserOK), nil

}

/*
GetUsers gets users

Returns a list of all users in the system. Currently this endpoint is only available to users with admin permissions.

*/
func (a *Client) GetUsers(params *GetUsersParams, authInfo runtime.ClientAuthInfoWriter) (*GetUsersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetUsersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getUsers",
		Method:             "GET",
		PathPattern:        "/v4/users/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetUsersReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetUsersOK), nil

}

/*
ModifyPassword modifies password

This operation allows you to change your password. Admins are able to change passwords of other users.

*/
func (a *Client) ModifyPassword(params *ModifyPasswordParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyPasswordAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyPasswordParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "modifyPassword",
		Method:             "POST",
		PathPattern:        "/v4/users/{email}/password/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModifyPasswordReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyPasswordAccepted), nil

}

/*
ModifyUser modifies user

This operation allows you to change details of a given user. Only administrators can edit accounts of other users.

*/
func (a *Client) ModifyUser(params *ModifyUserParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyUserOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyUserParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "modifyUser",
		Method:             "PATCH",
		PathPattern:        "/v4/users/{email}/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModifyUserReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyUserOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
