/*
 * The Giant Swarm API
 *
 * This is the documentation for the Giant Swarm API.  For an introduction to Giant Swarm, refer to the [documentation site](https://docs.giantswarm.io/).  The Giant Swarm API attempts to behave in a __restful__ way. As a developer, you access resources using the `GET` method and, for example, delete them using the same path and the `DELETE` method.  Accessing resources via GET usually returns all information available about a resource, while collections, like for example the list of all clusters you have access to, only contain a selected few attributes of each member item.  Some requests, like for example the request to create a new cluster, don't return the resource itself. Instead, the response delivers a standard message body, showing a `code` and a `message` part. The `message` contains information for you or a client's end user. The `code` attribute contains some string (example: `RESOURCE_CREATED`) that is supposed to give you details on the state of the operation, in addition to standard HTTP status codes. This message format is also used in the case of errors. We provide a [list of all response codes](https://github.com/giantswarm/api-spec/blob/master/details/RESPONSE_CODES.md) outside this documentation.  Feedback on the API as well as this documentation is welcome via `support@giantswarm.io` or on IRC channel [#giantswarm](irc://irc.freenode.org:6667/#giantswarm) on freenode.  ## Source  The source of this documentation is available on [GitHub](https://github.com/giantswarm/api-spec). 
 *
 * API version: 4.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"fmt"
	"strings"
	"github.com/antihax/optional"
)

// Linger please
var (
	_ context.Context
)

type AppsApiService service

/*
AppsApiService Install an app
Install an app on a tenant cluster by posting to this endpoint.  The spec field represents the app we&#39;ll be installing, and so spec.name refers to the name of the chart that installs this app in the catalog.  The response you get on a succesful create includes the status of the app. However since the App is still initialising and this is an asynchronous operation, it is likely that the fields in this status object will be all empty values.  To check on the status of your app, perform a GET to /v4/clusters/{cluster_id}/apps/, and check the status field of the app.  ### Example PUT request &#x60;&#x60;&#x60;json   {     \&quot;spec\&quot;: {       \&quot;catalog\&quot;: \&quot;sample-catalog\&quot;,       \&quot;name\&quot;: \&quot;prometheus-chart\&quot;,       \&quot;namespace\&quot;: \&quot;prometheus\&quot;,       \&quot;version\&quot;: \&quot;0.2.0\&quot;,     }   } &#x60;&#x60;&#x60;  ### About the user_config field in the response This field is not editable by you, but is set automatically by the API if a configmap named &#x60;{app_name}-user-values&#x60; exists in the tenant cluster namespace on the control plane.  The &#x60;/v4/clusters/{cluster_id}/apps/{app_name}/config/&#x60; endpoints allows you to create such a configmap using this API.  It is recommended to create your config before creating your app. This will result in a faster deploy.  However, you can create your config after creating the app if you wish, this API will take care of setting the &#x60;user_config&#x60; field of the app correctly for you.  ### Why can&#39;t I just set the &#x60;user_config&#x60; value myself? It simplifies usage while also being a security measure.  Furthermore it is also a security measure and ensures that users of this API can&#39;t access arbitrary configmaps of the control plane.  This API will only allow you to edit or access configmaps that adhere to a strict naming convention. 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization As described in the [authentication](#section/Authentication) section 
 * @param clusterId Cluster ID
 * @param appName App Name
 * @param optional nil or *CreateClusterAppOpts - Optional Parameters:
 * @param "XRequestID" (optional.String) -  A randomly generated key that can be used to track a request throughout services of Giant Swarm. 
 * @param "XGiantSwarmActivity" (optional.String) -  Name of an activity to track, like \"list-clusters\". This allows to analyze several API requests sent in context and gives an idea on the purpose. 
 * @param "XGiantSwarmCmdLine" (optional.String) -  If activity has been issued by a CLI, this header can contain the command line 
 * @param "V4CreateAppRequest" (optional.Interface of V4CreateAppRequest) - 
@return V4App
*/

type CreateClusterAppOpts struct {
	XRequestID optional.String
	XGiantSwarmActivity optional.String
	XGiantSwarmCmdLine optional.String
	V4CreateAppRequest optional.Interface
}

func (a *AppsApiService) CreateClusterApp(ctx context.Context, authorization string, clusterId string, appName string, localVarOptionals *CreateClusterAppOpts) (V4App, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V4App
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v4/clusters/{cluster_id}/apps/{app_name}/"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", fmt.Sprintf("%v", clusterId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_name"+"}", fmt.Sprintf("%v", appName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(authorization, "")
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XGiantSwarmActivity.IsSet() {
		localVarHeaderParams["X-Giant-Swarm-Activity"] = parameterToString(localVarOptionals.XGiantSwarmActivity.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XGiantSwarmCmdLine.IsSet() {
		localVarHeaderParams["X-Giant-Swarm-CmdLine"] = parameterToString(localVarOptionals.XGiantSwarmCmdLine.Value(), "")
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.V4CreateAppRequest.IsSet() {
		localVarOptionalV4CreateAppRequest, localVarOptionalV4CreateAppRequestok := localVarOptionals.V4CreateAppRequest.Value().(V4CreateAppRequest)
		if !localVarOptionalV4CreateAppRequestok {
			return localVarReturnValue, nil, reportError("v4CreateAppRequest should be V4CreateAppRequest")
		}
		localVarPostBody = &localVarOptionalV4CreateAppRequest
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v V4App
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 409 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
AppsApiService Delete an app
This operation allows a user to delete an app. 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization As described in the [authentication](#section/Authentication) section 
 * @param clusterId Cluster ID
 * @param appName App Name
 * @param optional nil or *DeleteClusterAppOpts - Optional Parameters:
 * @param "XRequestID" (optional.String) -  A randomly generated key that can be used to track a request throughout services of Giant Swarm. 
 * @param "XGiantSwarmActivity" (optional.String) -  Name of an activity to track, like \"list-clusters\". This allows to analyze several API requests sent in context and gives an idea on the purpose. 
 * @param "XGiantSwarmCmdLine" (optional.String) -  If activity has been issued by a CLI, this header can contain the command line 
@return V4GenericResponse
*/

type DeleteClusterAppOpts struct {
	XRequestID optional.String
	XGiantSwarmActivity optional.String
	XGiantSwarmCmdLine optional.String
}

func (a *AppsApiService) DeleteClusterApp(ctx context.Context, authorization string, clusterId string, appName string, localVarOptionals *DeleteClusterAppOpts) (V4GenericResponse, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V4GenericResponse
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v4/clusters/{cluster_id}/apps/{app_name}/"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", fmt.Sprintf("%v", clusterId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_name"+"}", fmt.Sprintf("%v", appName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(authorization, "")
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XGiantSwarmActivity.IsSet() {
		localVarHeaderParams["X-Giant-Swarm-Activity"] = parameterToString(localVarOptionals.XGiantSwarmActivity.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XGiantSwarmCmdLine.IsSet() {
		localVarHeaderParams["X-Giant-Swarm-CmdLine"] = parameterToString(localVarOptionals.XGiantSwarmCmdLine.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
AppsApiService Get a list of app catalogs configured on your installation.
Returns an array of app catalog objects, which contains further metadata, including a URL to fetch the full index of each catalog.  #### About the Labels  - &#x60;application.giantswarm.io/catalog-type&#x60;   Describes the type of catalog.    - &#x60;managed&#x60; - Apps in this catalog are managed by Giant Swarm.   - &#x60;incubator&#x60; - Apps in this catalog are a work in progress. They&#39;re                   made with your Giant Swarm cluster in mind though, so                   they should work. Feedback is appreciated on these apps.   - &#x60;test&#x60; - Apps in this catalog will soon graduate to incubator status,               you most likely will not see any &#x60;test&#x60; catalogs on your               installations.   - &#x60;community&#x60; - Apps in this catalog are provided by the kubernetes                   community. They will most likely not work without making                   some changes to the security settings of your cluster.    App Catalogs can also be labeled as &#x60;internal&#x60;, however these catalogs   contain apps that run on our control planes. These &#x60;internal&#x60; app catalogs   will be filtered out and never shown when calling this endpoint.    For more details on app catalogs visit: https://docs.giantswarm.io/basics/app-catalog/   ### Example &#x60;&#x60;&#x60;json   [     {       \&quot;metadata\&quot;: {         \&quot;name\&quot;: \&quot;sample-catalog\&quot;,         \&quot;labels\&quot;: {           \&quot;application.giantswarm.io/catalog-type\&quot;: \&quot;test\&quot;,           \&quot;app-operator.giantswarm.io/version\&quot;: \&quot;1.0.0\&quot;,         },       },        \&quot;spec\&quot;: {         \&quot;description\&quot;: \&quot;Giant Swarm&#39;s Sample Catalog with a few apps to test things out.\&quot;,         \&quot;logoURL\&quot;: \&quot;https://s.giantswarm.io/app-catalog/1/images/sample-catalog.png\&quot;,          \&quot;storage\&quot;: {           \&quot;URL\&quot;: \&quot;https://giantswarm.github.com/sample-catalog/\&quot;,           \&quot;type\&quot;: \&quot;helm\&quot;         },         \&quot;title\&quot;: \&quot;Sample Catalog\&quot;       }     }   ] &#x60;&#x60;&#x60; 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization As described in the [authentication](#section/Authentication) section 
 * @param optional nil or *GetAppCatalogsOpts - Optional Parameters:
 * @param "XRequestID" (optional.String) -  A randomly generated key that can be used to track a request throughout services of Giant Swarm. 
 * @param "XGiantSwarmActivity" (optional.String) -  Name of an activity to track, like \"list-clusters\". This allows to analyze several API requests sent in context and gives an idea on the purpose. 
 * @param "XGiantSwarmCmdLine" (optional.String) -  If activity has been issued by a CLI, this header can contain the command line 
@return []map[string]interface{}
*/

type GetAppCatalogsOpts struct {
	XRequestID optional.String
	XGiantSwarmActivity optional.String
	XGiantSwarmCmdLine optional.String
}

func (a *AppsApiService) GetAppCatalogs(ctx context.Context, authorization string, localVarOptionals *GetAppCatalogsOpts) ([]map[string]interface{}, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []map[string]interface{}
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v4/appcatalogs/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(authorization, "")
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XGiantSwarmActivity.IsSet() {
		localVarHeaderParams["X-Giant-Swarm-Activity"] = parameterToString(localVarOptionals.XGiantSwarmActivity.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XGiantSwarmCmdLine.IsSet() {
		localVarHeaderParams["X-Giant-Swarm-CmdLine"] = parameterToString(localVarOptionals.XGiantSwarmCmdLine.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []map[string]interface{}
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
AppsApiService Get cluster apps
Returns an array of apps installed on a given cluster.  ### Example &#x60;&#x60;&#x60;json   [     {       \&quot;metadata\&quot;: {         \&quot;name\&quot;: \&quot;my-awesome-prometheus\&quot;,       },        \&quot;spec\&quot;: {         \&quot;catalog\&quot;: \&quot;sample-catalog\&quot;         \&quot;name\&quot;: \&quot;prometheus-chart\&quot;,         \&quot;namespace\&quot;: \&quot;giantswarm\&quot;,         \&quot;version\&quot;: \&quot;0.2.0\&quot;,         \&quot;user_config\&quot;: {           \&quot;configmap\&quot;: {             \&quot;name\&quot;: \&quot;prometheus-user-values\&quot;,             \&quot;namespace\&quot;: \&quot;123ab\&quot;           }         }       },        \&quot;status\&quot;: {         \&quot;app_version\&quot;: \&quot;1.0.0\&quot;,         \&quot;release\&quot;: {           \&quot;last_deployed\&quot;: \&quot;2019-04-08T12:34:00Z\&quot;,           \&quot;status\&quot;: \&quot;DEPLOYED\&quot;         },         \&quot;version\&quot;: \&quot;0.2.0\&quot;,       }     }   ] &#x60;&#x60;&#x60; 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization As described in the [authentication](#section/Authentication) section 
 * @param clusterId Cluster ID
 * @param optional nil or *GetClusterAppsOpts - Optional Parameters:
 * @param "XRequestID" (optional.String) -  A randomly generated key that can be used to track a request throughout services of Giant Swarm. 
 * @param "XGiantSwarmActivity" (optional.String) -  Name of an activity to track, like \"list-clusters\". This allows to analyze several API requests sent in context and gives an idea on the purpose. 
 * @param "XGiantSwarmCmdLine" (optional.String) -  If activity has been issued by a CLI, this header can contain the command line 
@return []V4App
*/

type GetClusterAppsOpts struct {
	XRequestID optional.String
	XGiantSwarmActivity optional.String
	XGiantSwarmCmdLine optional.String
}

func (a *AppsApiService) GetClusterApps(ctx context.Context, authorization string, clusterId string, localVarOptionals *GetClusterAppsOpts) ([]V4App, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  []V4App
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v4/clusters/{cluster_id}/apps/"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", fmt.Sprintf("%v", clusterId), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(authorization, "")
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XGiantSwarmActivity.IsSet() {
		localVarHeaderParams["X-Giant-Swarm-Activity"] = parameterToString(localVarOptionals.XGiantSwarmActivity.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XGiantSwarmCmdLine.IsSet() {
		localVarHeaderParams["X-Giant-Swarm-CmdLine"] = parameterToString(localVarOptionals.XGiantSwarmCmdLine.Value(), "")
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v []V4App
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}

/*
AppsApiService Modify an app
This operation allows you to modify an existing app.  The following attributes can be modified:  - &#x60;version&#x60;: Changing this field lets you upgrade or downgrade an app.  &#x60;catalog&#x60;, &#x60;name&#x60;, &#x60;namespace&#x60;, and &#x60;user_config&#x60; are not editable. If you need to move or rename an app, you should instead delete the app and make it again.  The request body must conform with the [JSON Patch Merge (RFC 7386)](https://tools.ietf.org/html/rfc7386) standard. Requests have to be sent with the &#x60;Content-Type: application/merge-patch+json&#x60; header. 
 * @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param authorization As described in the [authentication](#section/Authentication) section 
 * @param clusterId Cluster ID
 * @param appName App Name
 * @param optional nil or *ModifyClusterAppOpts - Optional Parameters:
 * @param "XRequestID" (optional.String) -  A randomly generated key that can be used to track a request throughout services of Giant Swarm. 
 * @param "XGiantSwarmActivity" (optional.String) -  Name of an activity to track, like \"list-clusters\". This allows to analyze several API requests sent in context and gives an idea on the purpose. 
 * @param "XGiantSwarmCmdLine" (optional.String) -  If activity has been issued by a CLI, this header can contain the command line 
 * @param "V4ModifyAppRequest" (optional.Interface of V4ModifyAppRequest) - 
@return V4App
*/

type ModifyClusterAppOpts struct {
	XRequestID optional.String
	XGiantSwarmActivity optional.String
	XGiantSwarmCmdLine optional.String
	V4ModifyAppRequest optional.Interface
}

func (a *AppsApiService) ModifyClusterApp(ctx context.Context, authorization string, clusterId string, appName string, localVarOptionals *ModifyClusterAppOpts) (V4App, *http.Response, error) {
	var (
		localVarHttpMethod   = http.MethodPatch
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  V4App
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/v4/clusters/{cluster_id}/apps/{app_name}/"
	localVarPath = strings.Replace(localVarPath, "{"+"cluster_id"+"}", fmt.Sprintf("%v", clusterId), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"app_name"+"}", fmt.Sprintf("%v", appName), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	localVarHeaderParams["Authorization"] = parameterToString(authorization, "")
	if localVarOptionals != nil && localVarOptionals.XRequestID.IsSet() {
		localVarHeaderParams["X-Request-ID"] = parameterToString(localVarOptionals.XRequestID.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XGiantSwarmActivity.IsSet() {
		localVarHeaderParams["X-Giant-Swarm-Activity"] = parameterToString(localVarOptionals.XGiantSwarmActivity.Value(), "")
	}
	if localVarOptionals != nil && localVarOptionals.XGiantSwarmCmdLine.IsSet() {
		localVarHeaderParams["X-Giant-Swarm-CmdLine"] = parameterToString(localVarOptionals.XGiantSwarmCmdLine.Value(), "")
	}
	// body params
	if localVarOptionals != nil && localVarOptionals.V4ModifyAppRequest.IsSet() {
		localVarOptionalV4ModifyAppRequest, localVarOptionalV4ModifyAppRequestok := localVarOptionals.V4ModifyAppRequest.Value().(V4ModifyAppRequest)
		if !localVarOptionalV4ModifyAppRequestok {
			return localVarReturnValue, nil, reportError("v4ModifyAppRequest should be V4ModifyAppRequest")
		}
		localVarPostBody = &localVarOptionalV4ModifyAppRequest
	}

	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["Authorization"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHttpResponse.Body)
	localVarHttpResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHttpResponse, err
	}

	if localVarHttpResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHttpResponse.Status,
		}
		if localVarHttpResponse.StatusCode == 200 {
			var v V4App
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 400 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 401 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 404 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		if localVarHttpResponse.StatusCode == 0 {
			var v V4GenericResponse
			err = a.client.decode(&v, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHttpResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHttpResponse, newErr
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHttpResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHttpResponse, newErr
	}

	return localVarReturnValue, localVarHttpResponse, nil
}
