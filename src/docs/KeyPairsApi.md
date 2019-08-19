# \KeyPairsApi

All URIs are relative to *https://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKeyPair**](KeyPairsApi.md#AddKeyPair) | **Post** /v4/clusters/{cluster_id}/key-pairs/ | Create key pair
[**GetKeyPairs**](KeyPairsApi.md#GetKeyPairs) | **Get** /v4/clusters/{cluster_id}/key-pairs/ | Get key pairs



## AddKeyPair

> V4AddKeyPairResponse AddKeyPair(ctx, authorization, clusterId, v4AddKeyPairRequest, optional)
Create key pair

This operation allows to create a new key pair for accessing a specific cluster.  A key pair consists of an unencrypted private RSA key and an X.509 certificate. In addition, when obtaining a key pair for a cluster, the cluster's certificate authority file (CA certificate) is delivered, which is required by TLS clients to establish trust to the cluster.  In addition to the credentials itself, a key pair has some metadata like a unique ID, a creation timestamp and a free text `description` that you can use at will, for example to note for whom a key pair has been issued.  ### Customizing the certificate's subject for K8s RBAC  It is possible to set the Common Name and Organization fields of the generated certificate's subject.  - `cn_prefix`: The certificate's common name uses this format: `.user.`.    `clusterdomain` is specific to your cluster and is not editable.    The `cn_prefix` however is editable. When left blank it will default   to the email address of the Giant Swarm user that is performing the   create key pair request.    The common name is used as the username for requests to the Kubernetes API. This allows you   to set up role-based access controls.   - `certificate_organizations`: This will set the certificate's `organization` fields. Use a comma separated list of values.   The Kubernetes API will use these values as group memberships.  __Note:__ The actual credentials coming with the key pair (key, certificate) can only be accessed once, as the result of the `POST` request that triggers their creation. This restriction exists to minimize the risk of credentials being leaked. If you fail to capture the credentials upon creation, you'll have to repeat the creation request. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
**v4AddKeyPairRequest** | [**V4AddKeyPairRequest**](V4AddKeyPairRequest.md)| While the &#x60;ttl_hours&#x60; attribute is optional and will be set to a default value when omitted, the &#x60;description&#x60; is mandatory.  | 
 **optional** | ***AddKeyPairOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddKeyPairOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V4AddKeyPairResponse**](V4AddKeyPairResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetKeyPairs

> []map[string]interface{} GetKeyPairs(ctx, authorization, clusterId, optional)
Get key pairs

Returns a list of information on all key pairs of a cluster as an array.  The individual array items contain metadata on the key pairs, but neither the key nor the certificate. These can only be obtained upon creation, using the [addKeypair](#operation/addKeyPair) operation. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**clusterId** | **string**| Cluster ID | 
 **optional** | ***GetKeyPairsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetKeyPairsOpts struct


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

