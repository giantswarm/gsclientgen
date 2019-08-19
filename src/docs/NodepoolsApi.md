# \NodepoolsApi

All URIs are relative to *https://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddNodePool**](NodepoolsApi.md#AddNodePool) | **Post** /v5/clusters/{cluster_id}/nodepools/ | Create node pool
[**DeleteNodePool**](NodepoolsApi.md#DeleteNodePool) | **Delete** /v5/clusters/{cluster_id}/nodepools/{nodepool_id}/ | Delete node pool
[**GetNodePool**](NodepoolsApi.md#GetNodePool) | **Get** /v5/clusters/{cluster_id}/nodepools/{nodepool_id}/ | Get node pool details
[**GetNodePools**](NodepoolsApi.md#GetNodePools) | **Get** /v5/clusters/{cluster_id}/nodepools/ | Get node pools
[**ModifyNodePool**](NodepoolsApi.md#ModifyNodePool) | **Patch** /v5/clusters/{cluster_id}/nodepools/{nodepool_id}/ | Modify node pool



## AddNodePool

> V5GetNodePoolResponse AddNodePool(ctx, authorization, clusterId, optional)
Create node pool

This allows to add a node pool to a cluster.  Some, but not all, node pool configuration can be changed after creation. The following settings will have a permanent effect:  - `availability_zones` - `instance_type` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
 **optional** | ***AddNodePoolOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddNodePoolOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 
 **v5AddNodePoolRequest** | [**optional.Interface of V5AddNodePoolRequest**](V5AddNodePoolRequest.md)|  | 

### Return type

[**V5GetNodePoolResponse**](V5GetNodePoolResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteNodePool

> V4GenericResponse DeleteNodePool(ctx, authorization, clusterId, nodepoolId, optional)
Delete node pool

Triggers the deletion of a node pool.  Nodes in the pool will first be cordoned and drained. Note that it is your responsibililty to make sure that workloads using the node pool can be scheduled elsewhere. We recommend to double-check the available capacity of remaining node pools, as well as any node selectors and taints. Also you can do the draining yourself before issuing the delete request, to observe the outcome. Use  ``` kubectl drain nodes -l giantswarm.nodepool_id= ... ```  TODO: adapt the command for correct label syntax 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
**nodepoolId** | **string**| Node Pool ID | 
 **optional** | ***DeleteNodePoolOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteNodePoolOpts struct


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


## GetNodePool

> V5GetNodePoolResponse GetNodePool(ctx, authorization, clusterId, nodepoolId, optional)
Get node pool details

Returns all available details on a specific node pool. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
**nodepoolId** | **string**| Node Pool ID | 
 **optional** | ***GetNodePoolOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetNodePoolOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V5GetNodePoolResponse**](V5GetNodePoolResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodePools

> []V5GetNodePoolResponse GetNodePools(ctx, authorization, clusterId, optional)
Get node pools

Returns a list of node pools from a given cluster. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
 **optional** | ***GetNodePoolsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetNodePoolsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**[]V5GetNodePoolResponse**](V5GetNodePoolResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyNodePool

> V5GetNodePoolResponse ModifyNodePool(ctx, authorization, clusterId, nodepoolId, v5ModifyNodePoolRequest, optional)
Modify node pool

Allows to rename a node pool or change its scaling settings. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
**nodepoolId** | **string**| Node Pool ID | 
**v5ModifyNodePoolRequest** | [**V5ModifyNodePoolRequest**](V5ModifyNodePoolRequest.md)| Merge-patch body | 
 **optional** | ***ModifyNodePoolOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ModifyNodePoolOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V5GetNodePoolResponse**](V5GetNodePoolResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

