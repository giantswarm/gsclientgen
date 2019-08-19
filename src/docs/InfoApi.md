# \InfoApi

All URIs are relative to *https://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetInfo**](InfoApi.md#GetInfo) | **Get** /v4/info/ | Get information on the installation



## GetInfo

> V4InfoResponse GetInfo(ctx, authorization, optional)
Get information on the installation

Returns a set of details on the installation. The output varies based on the provider used in the installation.  This information is useful for example when creating new cluster, to prevent creating clusters with more worker nodes than possible.  ### Example for an AWS-based installation  ```json {   \"general\": {     \"installation_name\": \"shire\",     \"provider\": \"aws\",     \"datacenter\": \"eu-central-1\",     \"availability_zones\": {       \"max\": 3,       \"default\": 1,     }   },   \"stats\": {     \"cluster_creation_duration\": {       \"median\": 750,       \"p25\": 700,       \"p75\": 800     }   },   \"workers\": {     \"count_per_cluster\": {       \"max\": null,       \"default\": 3     },     \"instance_type\": {       \"options\": [         \"m3.medium\", \"m3.large\", \"m3.xlarge\"       ],       \"default\": \"m3.large\"     }   } } ```  ### Example for a KVM-based installation  ```json {   \"general\": {     \"installation_name\": \"isengard\",     \"provider\": \"kvm\",     \"datacenter\": \"string\",     \"availability_zones\": {       \"max\": 1,       \"default\": 1,     }   },   \"stats\": {     \"cluster_creation_duration\": {       \"median\": 750,       \"p25\": 700,       \"p75\": 800     }   },   \"workers\": {     \"count_per_cluster\": {       \"max\": 8,       \"default\": 3     },   } } ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
 **optional** | ***GetInfoOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetInfoOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V4InfoResponse**](V4InfoResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

