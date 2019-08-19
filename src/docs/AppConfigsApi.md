# \AppConfigsApi

All URIs are relative to *https://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterAppConfig**](AppConfigsApi.md#CreateClusterAppConfig) | **Put** /v4/clusters/{cluster_id}/apps/{app_name}/config/ | Create app config
[**DeleteClusterAppConfig**](AppConfigsApi.md#DeleteClusterAppConfig) | **Delete** /v4/clusters/{cluster_id}/apps/{app_name}/config/ | Delete an app config
[**GetClusterAppConfig**](AppConfigsApi.md#GetClusterAppConfig) | **Get** /v4/clusters/{cluster_id}/apps/{app_name}/config/ | Get app config
[**ModifyClusterAppConfig**](AppConfigsApi.md#ModifyClusterAppConfig) | **Patch** /v4/clusters/{cluster_id}/apps/{app_name}/config/ | Modify app config



## CreateClusterAppConfig

> V4GenericResponse CreateClusterAppConfig(ctx, authorization, clusterId, appName, optional)
Create app config

This operation allows you to create a values configmap for a specific app. The app does not have to exist before hand.  If the app does exist, this endpoint will ensure that the App CR gets it's user_config field set correctly.  However, if the app exists and the user_config is already set to something, then this request will fail. You will in that case most likely want to update the config using the `PATCH /v4/clusters/{cluster_id}/apps/{app_name}/config/` endpoint.  ### Example POST request ```json   {     \"agent\": {       \"key\": \"secret-key-here\",       \"endpointHost\": \"saas-eu-west-1.instana.io\",       \"endpointPort\": \"443\",     },     \"zone\": {       \"name\": \"giantswarm-cluster\"     }   } ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
**appName** | **string**| App Name | 
 **optional** | ***CreateClusterAppConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateClusterAppConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 
 **requestBody** | [**optional.Interface of map[string]map[string]interface{}**](map[string]interface{}.md)|  | 

### Return type

[**V4GenericResponse**](V4GenericResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, appication/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClusterAppConfig

> V4GenericResponse DeleteClusterAppConfig(ctx, authorization, clusterId, appName, optional)
Delete an app config

This operation allows a user to delete an app's user config if it has been named according to the convention of {app-name}-user-values and stored in the cluster ID namespace.  Calling this endpoint will delete the ConfigMap, but it does not remove the reference to the ConfigMap in the (spec.user_config.configmap field) from the app.  Do make sure you also update the app and remove the reference.  The preferred order is to first remove the reference to the configmap by updating the app, and only then delete the configmap using this endpoint. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
**appName** | **string**| App Name | 
 **optional** | ***DeleteClusterAppConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteClusterAppConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V4GenericResponse**](V4GenericResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterAppConfig

> map[string]map[string]interface{} GetClusterAppConfig(ctx, authorization, clusterId, appName, optional)
Get app config

This operation allows you to fetch the user values configmap associated with an app. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
**appName** | **string**| App Name | 
 **optional** | ***GetClusterAppConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetClusterAppConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**map[string]map[string]interface{}**](map[string]interface{}.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, appication/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyClusterAppConfig

> V4GenericResponse ModifyClusterAppConfig(ctx, authorization, clusterId, appName, optional)
Modify app config

This operation allows you to modify the values configmap for a specific app. It's only possible to modify app configs that have been named according to the convention of {app-name}-user-values and stored in the cluster ID namespace.  The full values key of the configmap will be replaced by the JSON body of your request.  ### Example PATCH request ```json   {     \"agent\": {       \"key\": \"a-new-key-here\",     }   } ```  If the configmap contained content like:  ```json   {     \"agent\": {       \"key\": \"an-old-key-here\",       \"admin\": true,     },     \"server\": {       \"url\": \"giantswarm.io\",     }   } ```  Then the \"server\" and \"admin\" keys will be removed. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
**appName** | **string**| App Name | 
 **optional** | ***ModifyClusterAppConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ModifyClusterAppConfigOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 
 **requestBody** | [**optional.Interface of map[string]map[string]interface{}**](map[string]interface{}.md)|  | 

### Return type

[**V4GenericResponse**](V4GenericResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json, appication/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

