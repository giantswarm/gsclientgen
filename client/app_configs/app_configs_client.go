// Code generated by go-swagger; DO NOT EDIT.

package app_configs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new app configs API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for app configs API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
CreateClusterAppConfig creates app config

This operation allows you to create a values configmap for a specific app. The app does
not have to exist before hand.

If the app does exist, this endpoint will ensure that the App CR gets it's
user_config field set correctly.

However, if the app exists and the user_config is already set to something,
then this request will fail. You will in that case most likely want to
update the config using the `PATCH /v4/clusters/{cluster_id}/apps/{app_name}/config/`
endpoint.


### Example POST request
```json
  {
    "agent": {
      "key": "secret-key-here",
      "endpointHost": "saas-eu-west-1.instana.io",
      "endpointPort": "443",
    },
    "zone": {
      "name": "giantswarm-cluster"
    }
  }
```

*/
func (a *Client) CreateClusterAppConfig(params *CreateClusterAppConfigParams, authInfo runtime.ClientAuthInfoWriter) (*CreateClusterAppConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateClusterAppConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createClusterAppConfig",
		Method:             "PUT",
		PathPattern:        "/v4/clusters/{cluster_id}/apps/{app_name}/config/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateClusterAppConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*CreateClusterAppConfigOK), nil

}

/*
DeleteClusterAppConfig deletes an app config

This operation allows a user to delete an app's user config if it has been named according to the convention of {app-name}-user-values and
stored in the cluster ID namespace.

Calling this endpoint will delete the ConfigMap, but it does not remove the reference to the ConfigMap in the (spec.user_config.configmap field) from the app.

Do make sure you also update the app and remove the reference.

The preferred order is to first remove the reference to the configmap by
updating the app, and only then delete the configmap using this endpoint.

*/
func (a *Client) DeleteClusterAppConfig(params *DeleteClusterAppConfigParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteClusterAppConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteClusterAppConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteClusterAppConfig",
		Method:             "DELETE",
		PathPattern:        "/v4/clusters/{cluster_id}/apps/{app_name}/config/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &DeleteClusterAppConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DeleteClusterAppConfigOK), nil

}

/*
GetClusterAppConfig gets app config

This operation allows you to fetch the user values configmap associated
with an app.

*/
func (a *Client) GetClusterAppConfig(params *GetClusterAppConfigParams, authInfo runtime.ClientAuthInfoWriter) (*GetClusterAppConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetClusterAppConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getClusterAppConfig",
		Method:             "GET",
		PathPattern:        "/v4/clusters/{cluster_id}/apps/{app_name}/config/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetClusterAppConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetClusterAppConfigOK), nil

}

/*
ModifyClusterAppConfig modifies app config

This operation allows you to modify the values configmap for a specific app.
It's only possible to modify app configs that have been named according to the convention of
{app-name}-user-values and stored in the cluster ID namespace.

The full values key of the configmap will be replaced by the JSON body
of your request.

### Example PATCH request
```json
  {
    "agent": {
      "key": "a-new-key-here",
    }
  }
```

If the configmap contained content like:

```json
  {
    "agent": {
      "key": "an-old-key-here",
      "admin": true,
    },
    "server": {
      "url": "giantswarm.io",
    }
  }
```

Then the "server" and "admin" keys will be removed.

*/
func (a *Client) ModifyClusterAppConfig(params *ModifyClusterAppConfigParams, authInfo runtime.ClientAuthInfoWriter) (*ModifyClusterAppConfigOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewModifyClusterAppConfigParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "modifyClusterAppConfig",
		Method:             "PATCH",
		PathPattern:        "/v4/clusters/{cluster_id}/apps/{app_name}/config/",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ModifyClusterAppConfigReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ModifyClusterAppConfigOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}