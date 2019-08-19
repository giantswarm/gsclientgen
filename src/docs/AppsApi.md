# \AppsApi

All URIs are relative to *https://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateClusterApp**](AppsApi.md#CreateClusterApp) | **Put** /v4/clusters/{cluster_id}/apps/{app_name}/ | Install an app
[**DeleteClusterApp**](AppsApi.md#DeleteClusterApp) | **Delete** /v4/clusters/{cluster_id}/apps/{app_name}/ | Delete an app
[**GetAppCatalogs**](AppsApi.md#GetAppCatalogs) | **Get** /v4/appcatalogs/ | Get a list of app catalogs configured on your installation.
[**GetClusterApps**](AppsApi.md#GetClusterApps) | **Get** /v4/clusters/{cluster_id}/apps/ | Get cluster apps
[**ModifyClusterApp**](AppsApi.md#ModifyClusterApp) | **Patch** /v4/clusters/{cluster_id}/apps/{app_name}/ | Modify an app



## CreateClusterApp

> V4App CreateClusterApp(ctx, authorization, clusterId, appName, optional)
Install an app

Install an app on a tenant cluster by posting to this endpoint.  The spec field represents the app we'll be installing, and so spec.name refers to the name of the chart that installs this app in the catalog.  The response you get on a succesful create includes the status of the app. However since the App is still initialising and this is an asynchronous operation, it is likely that the fields in this status object will be all empty values.  To check on the status of your app, perform a GET to /v4/clusters/{cluster_id}/apps/, and check the status field of the app.  ### Example PUT request ```json   {     \"spec\": {       \"catalog\": \"sample-catalog\",       \"name\": \"prometheus-chart\",       \"namespace\": \"prometheus\",       \"version\": \"0.2.0\",     }   } ```  ### About the user_config field in the response This field is not editable by you, but is set automatically by the API if a configmap named `{app_name}-user-values` exists in the tenant cluster namespace on the control plane.  The `/v4/clusters/{cluster_id}/apps/{app_name}/config/` endpoints allows you to create such a configmap using this API.  It is recommended to create your config before creating your app. This will result in a faster deploy.  However, you can create your config after creating the app if you wish, this API will take care of setting the `user_config` field of the app correctly for you.  ### Why can't I just set the `user_config` value myself? It simplifies usage while also being a security measure.  Furthermore it is also a security measure and ensures that users of this API can't access arbitrary configmaps of the control plane.  This API will only allow you to edit or access configmaps that adhere to a strict naming convention. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
**appName** | **string**| App Name | 
 **optional** | ***CreateClusterAppOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateClusterAppOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 
 **v4CreateAppRequest** | [**optional.Interface of V4CreateAppRequest**](V4CreateAppRequest.md)|  | 

### Return type

[**V4App**](V4App.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteClusterApp

> V4GenericResponse DeleteClusterApp(ctx, authorization, clusterId, appName, optional)
Delete an app

This operation allows a user to delete an app. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
**appName** | **string**| App Name | 
 **optional** | ***DeleteClusterAppOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteClusterAppOpts struct


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


## GetAppCatalogs

> []map[string]interface{} GetAppCatalogs(ctx, authorization, optional)
Get a list of app catalogs configured on your installation.

Returns an array of app catalog objects, which contains further metadata, including a URL to fetch the full index of each catalog.  #### About the Labels  - `application.giantswarm.io/catalog-type`   Describes the type of catalog.    - `managed` - Apps in this catalog are managed by Giant Swarm.   - `incubator` - Apps in this catalog are a work in progress. They're                   made with your Giant Swarm cluster in mind though, so                   they should work. Feedback is appreciated on these apps.   - `test` - Apps in this catalog will soon graduate to incubator status,               you most likely will not see any `test` catalogs on your               installations.   - `community` - Apps in this catalog are provided by the kubernetes                   community. They will most likely not work without making                   some changes to the security settings of your cluster.    App Catalogs can also be labeled as `internal`, however these catalogs   contain apps that run on our control planes. These `internal` app catalogs   will be filtered out and never shown when calling this endpoint.    For more details on app catalogs visit: https://docs.giantswarm.io/basics/app-catalog/   ### Example ```json   [     {       \"metadata\": {         \"name\": \"sample-catalog\",         \"labels\": {           \"application.giantswarm.io/catalog-type\": \"test\",           \"app-operator.giantswarm.io/version\": \"1.0.0\",         },       },        \"spec\": {         \"description\": \"Giant Swarm's Sample Catalog with a few apps to test things out.\",         \"logoURL\": \"https://s.giantswarm.io/app-catalog/1/images/sample-catalog.png\",          \"storage\": {           \"URL\": \"https://giantswarm.github.com/sample-catalog/\",           \"type\": \"helm\"         },         \"title\": \"Sample Catalog\"       }     }   ] ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
 **optional** | ***GetAppCatalogsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAppCatalogsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**[]map[string]interface{}**](map[string]interface{}.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterApps

> []V4App GetClusterApps(ctx, authorization, clusterId, optional)
Get cluster apps

Returns an array of apps installed on a given cluster.  ### Example ```json   [     {       \"metadata\": {         \"name\": \"my-awesome-prometheus\",       },        \"spec\": {         \"catalog\": \"sample-catalog\"         \"name\": \"prometheus-chart\",         \"namespace\": \"giantswarm\",         \"version\": \"0.2.0\",         \"user_config\": {           \"configmap\": {             \"name\": \"prometheus-user-values\",             \"namespace\": \"123ab\"           }         }       },        \"status\": {         \"app_version\": \"1.0.0\",         \"release\": {           \"last_deployed\": \"2019-04-08T12:34:00Z\",           \"status\": \"DEPLOYED\"         },         \"version\": \"0.2.0\",       }     }   ] ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
 **optional** | ***GetClusterAppsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetClusterAppsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**[]V4App**](V4App.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyClusterApp

> V4App ModifyClusterApp(ctx, authorization, clusterId, appName, optional)
Modify an app

This operation allows you to modify an existing app.  The following attributes can be modified:  - `version`: Changing this field lets you upgrade or downgrade an app.  `catalog`, `name`, `namespace`, and `user_config` are not editable. If you need to move or rename an app, you should instead delete the app and make it again.  The request body must conform with the [JSON Patch Merge (RFC 7386)](https://tools.ietf.org/html/rfc7386) standard. Requests have to be sent with the `Content-Type: application/merge-patch+json` header. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
**appName** | **string**| App Name | 
 **optional** | ***ModifyClusterAppOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ModifyClusterAppOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 
 **v4ModifyAppRequest** | [**optional.Interface of V4ModifyAppRequest**](V4ModifyAppRequest.md)|  | 

### Return type

[**V4App**](V4App.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

