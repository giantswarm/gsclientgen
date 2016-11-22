# \DefaultApi

All URIs are relative to *https://api.giantswarm.io*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddKeyPair**](DefaultApi.md#AddKeyPair) | **Post** /v3/clusters/{cluster_id}/key-pairs/ | Add key-pair for cluster
[**GetKeyPairs**](DefaultApi.md#GetKeyPairs) | **Get** /v3/clusters/{cluster_id}/key-pairs/ | Get key-pairs for cluster
[**GetOrganizationClusters**](DefaultApi.md#GetOrganizationClusters) | **Get** /v3/orgs/{organization_name}/clusters/ | Get clusters for organization
[**GetUserOrganizations**](DefaultApi.md#GetUserOrganizations) | **Get** /v1/user/me/memberships | Get organizations for user
[**UserLogin**](DefaultApi.md#UserLogin) | **Post** /v1/user/{email}/login | Log in as a user
[**UserLogout**](DefaultApi.md#UserLogout) | **Post** /v1/token/logout | Expire the currently used auth token


# **AddKeyPair**
> AddKeyPairResponseModel AddKeyPair($authorization, $clusterId, $body)

Add key-pair for cluster


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string**| Header to pass an authorization token. The value has to be in the form &#x60;giantswarm &lt;token&gt;&#x60;. | 
 **clusterId** | **string**| ID of the cluster to create the key-pair for | 
 **body** | [**AddKeyPairBody**](AddKeyPairBody.md)| Description and expiry time for the new key-pair | 

### Return type

[**AddKeyPairResponseModel**](AddKeyPairResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetKeyPairs**
> KeyPairsResponseModel GetKeyPairs($authorization, $clusterId)

Get key-pairs for cluster


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string**| Header to pass an authorization token. The value has to be in the form &#x60;giantswarm &lt;token&gt;&#x60;. | 
 **clusterId** | **string**|  | 

### Return type

[**KeyPairsResponseModel**](KeyPairsResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetOrganizationClusters**
> OrganizationClustersResponseModel GetOrganizationClusters($authorization, $organizationName)

Get clusters for organization


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string**| Header to pass an authorization token. The value has to be in the form &#x60;giantswarm &lt;token&gt;&#x60;. | 
 **organizationName** | **string**|  | 

### Return type

[**OrganizationClustersResponseModel**](OrganizationClustersResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserOrganizations**
> UserOrganizationsResponseModel GetUserOrganizations($authorization)

Get organizations for user

Returns a list of organizations of which the current user is a member


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string**| Header to pass an authorization token. The value has to be in the form &#x60;giantswarm &lt;token&gt;&#x60;. | 

### Return type

[**UserOrganizationsResponseModel**](UserOrganizationsResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserLogin**
> LoginResponseModel UserLogin($email, $payload)

Log in as a user

This method takes email and password of a user and returns a new session token. The token can be found in the `data.Id` field of the response object.


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **email** | **string**| User email address | 
 **payload** | [**LoginBodyModel**](LoginBodyModel.md)| base64 encoded password | 

### Return type

[**LoginResponseModel**](LoginResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserLogout**
> GenericResponseModel UserLogout($authorization)

Expire the currently used auth token


### Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string**| Header to pass an authorization token. The value has to be in the form &#x60;giantswarm &lt;token&gt;&#x60;. | 

### Return type

[**GenericResponseModel**](GenericResponseModel.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/plain

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)
