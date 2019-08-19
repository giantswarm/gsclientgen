# \ClustersApi

All URIs are relative to *https://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCluster**](ClustersApi.md#AddCluster) | **Post** /v4/clusters/ | Create cluster (v4)
[**AddClusterV5**](ClustersApi.md#AddClusterV5) | **Post** /v5/clusters/ | Create cluster (v5)
[**DeleteCluster**](ClustersApi.md#DeleteCluster) | **Delete** /v4/clusters/{cluster_id}/ | Delete cluster
[**GetCluster**](ClustersApi.md#GetCluster) | **Get** /v4/clusters/{cluster_id}/ | Get cluster details (v4)
[**GetClusterStatus**](ClustersApi.md#GetClusterStatus) | **Get** /v4/clusters/{cluster_id}/status/ | Get cluster status
[**GetClusterV5**](ClustersApi.md#GetClusterV5) | **Get** /v5/clusters/{cluster_id}/ | Get cluster details (v5)
[**GetClusters**](ClustersApi.md#GetClusters) | **Get** /v4/clusters/ | Get clusters
[**ModifyCluster**](ClustersApi.md#ModifyCluster) | **Patch** /v4/clusters/{cluster_id}/ | Modify cluster (v4)
[**ModifyClusterV5**](ClustersApi.md#ModifyClusterV5) | **Patch** /v5/clusters/{cluster_id}/ | Modify cluster (v5)



## AddCluster

> V4GenericResponse AddCluster(ctx, authorization, v4AddClusterRequest, optional)
Create cluster (v4)

This operation is used to create a new Kubernetes cluster or \"tenant cluster\".  ### Cluster definition  The cluster definition format allows to set a number of optional configuration details, like worker node configuration, with node specification depending on the provider (e. g. on Azure the VM size, or on KVM the memory size and number of CPU cores).  One attribute is __mandatory__ upon creation: The `owner` attribute must carry the name of the organization the cluster will belong to. Note that the acting user must be a member of that organization in order to create a cluster.  For all other attributes, defaults will be applied if the attribute is not set. Check out the [getInfo](#operation/getInfo) operation for more info about defaults. If no `release_version` is set, the latest release version available for the provider will be used. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**v4AddClusterRequest** | [**V4AddClusterRequest**](V4AddClusterRequest.md)| New cluster definition | 
 **optional** | ***AddClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddClusterOpts struct


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

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddClusterV5

> V5ClusterDetailsResponse AddClusterV5(ctx, authorization, v5AddClusterRequest, optional)
Create cluster (v5)

Allows to create most recent clusters on AWS installations.  ### Node pools  In the Giant Swarm API v5, worker nodes are grouped into pools of worker nodes where all nodes share the same configuration.  When creating a cluster without submitting the `nodepools` attribute, or with its value being an empty array, one node pool with default configuration will be created.  Node pools can be created, deleted and modified during the entire lifetime of a cluster.  See [node pools](#tag/nodepools) and [Create node pool](#operation/addNodePool) for details. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**v5AddClusterRequest** | [**V5AddClusterRequest**](V5AddClusterRequest.md)| New cluster definition | 
 **optional** | ***AddClusterV5Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddClusterV5Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V5ClusterDetailsResponse**](V5ClusterDetailsResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCluster

> V4GenericResponse DeleteCluster(ctx, authorization, clusterId, optional)
Delete cluster

This operation triggers deleting a cluster with all resources attached to it.  Deleting a cluster causes the termination of all workloads running on the cluster. Data stored on the worker nodes will be lost. There is no way to undo this operation.  The response is sent as soon as the request is validated. At that point, workloads might still be running on the cluster and may be accessible for a little wile, until the cluster is actually deleted. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
 **optional** | ***DeleteClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteClusterOpts struct


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


## GetCluster

> V4ClusterDetailsResponse GetCluster(ctx, authorization, clusterId, optional)
Get cluster details (v4)

This operation allows to obtain basic details on a particular cluster. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
 **optional** | ***GetClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetClusterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V4ClusterDetailsResponse**](V4ClusterDetailsResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterStatus

> map[string]map[string]interface{} GetClusterStatus(ctx, authorization, clusterId, optional)
Get cluster status

Returns an object about a cluster's current state and past status transitions.  This endpoint exposes the status content of the Kubernetes resources representing a cluster in the corresponding custom resource. That is, depending on the provider:  - [`awsconfig.provider.giantswarm.io`](https://godoc.org/github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1#AWSConfig) - [`azureconfig.provider.giantswarm.io`](https://godoc.org/github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1#AzureConfig) - [`kvmconfig.provider.giantswarm.io`](https://godoc.org/github.com/giantswarm/apiextensions/pkg/apis/provider/v1alpha1#KVMConfig)  Note that structure and style differ from the rest of the v4 API. Also note that the structure depends on the release version and changes can be expected frequently. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
 **optional** | ***GetClusterStatusOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetClusterStatusOpts struct


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
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusterV5

> V5ClusterDetailsResponse GetClusterV5(ctx, authorization, clusterId, optional)
Get cluster details (v5)

Allows to retrieve cluster details on AWS installations. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
 **optional** | ***GetClusterV5Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetClusterV5Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V5ClusterDetailsResponse**](V5ClusterDetailsResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetClusters

> []V4ClusterListItem GetClusters(ctx, authorization, optional)
Get clusters

This operation fetches a list of clusters.  The result depends on the permissions of the user. A normal user will get all the clusters the user has access to, via organization membership. A user with admin permission will receive a list of all existing clusters.  The result array items are sparse representations of the cluster objects. To fetch more details on a cluster, use the [getClusterStatus](#operation/getClusterStatus) operation. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
 **optional** | ***GetClustersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetClustersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**[]V4ClusterListItem**](V4ClusterListItem.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyCluster

> V4ClusterDetailsResponse ModifyCluster(ctx, authorization, clusterId, v4ModifyClusterRequest, optional)
Modify cluster (v4)

This operation allows to modify an existing cluster.  A cluster modification is performed by submitting a `PATCH` request to the cluster resource (as described in the [addCluster](#operation/addCluster) and [getCluster](#operation/getCluster)) in form of a [JSON Patch Merge (RFC 7386)](https://tools.ietf.org/html/rfc7386). This means, only the attributes to be modified have to be contained in the request body.  The following attributes can be modified:  - `name`: Rename the cluster to something more fitting.  - `owner`: Changing the owner organization name means to change cluster ownership from one organization to another. The user performing the request has to be a member of both organizations.  - `release_version`: By changing this attribute you can upgrade a cluster to a newer [release](https://docs.giantswarm.io/api/#tag/releases).  - `scaling`: Adjust the cluster node limits to make use of auto scaling or to have full control over the node count. The latter can be achieved by setting `min` and `max` to the same values. Note that setting `min` and `max` to different values (effectively enabling autoscaling) is only available on AWS with releases from 6.2.0.    - `workers` (deprecated): For backward compatibility reasons, it is possible to provide this attribute as an array, where the number of items contained in the array determines the intended number of worker nodes in the cluster. The item count will be applied as both `min` and `max` value of the scaling limits, effectively disabling autoscaling. This requires the `scaling` attribute must not be present in the same request.  ### Limitations  - As of now, existing worker nodes cannot be modified. - The number of availability zones cannot be modified. - When removing nodes (scaling down), it is not possible to determine which nodes will be removed. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
**v4ModifyClusterRequest** | [**V4ModifyClusterRequest**](V4ModifyClusterRequest.md)| Merge-patch body | 
 **optional** | ***ModifyClusterOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ModifyClusterOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V4ClusterDetailsResponse**](V4ClusterDetailsResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyClusterV5

> V5ClusterDetailsResponse ModifyClusterV5(ctx, authorization, clusterId, v5ModifyClusterRequest, optional)
Modify cluster (v5)

Allows to change cluster properties on AWS installations. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
**v5ModifyClusterRequest** | [**V5ModifyClusterRequest**](V5ModifyClusterRequest.md)| Merge-patch body | 
 **optional** | ***ModifyClusterV5Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ModifyClusterV5Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V5ClusterDetailsResponse**](V5ClusterDetailsResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

