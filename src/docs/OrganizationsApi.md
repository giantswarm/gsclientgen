# \OrganizationsApi

All URIs are relative to *https://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCredentials**](OrganizationsApi.md#AddCredentials) | **Post** /v4/organizations/{organization_id}/credentials/ | Set credentials
[**AddOrganization**](OrganizationsApi.md#AddOrganization) | **Put** /v4/organizations/{organization_id}/ | Create an organization
[**DeleteOrganization**](OrganizationsApi.md#DeleteOrganization) | **Delete** /v4/organizations/{organization_id}/ | Delete an organization
[**GetCredential**](OrganizationsApi.md#GetCredential) | **Get** /v4/organizations/{organization_id}/credentials/{credential_id}/ | Get credential details
[**GetCredentials**](OrganizationsApi.md#GetCredentials) | **Get** /v4/organizations/{organization_id}/credentials/ | Get credentials
[**GetOrganization**](OrganizationsApi.md#GetOrganization) | **Get** /v4/organizations/{organization_id}/ | Get organization details
[**GetOrganizations**](OrganizationsApi.md#GetOrganizations) | **Get** /v4/organizations/ | Get organizations
[**ModifyOrganization**](OrganizationsApi.md#ModifyOrganization) | **Patch** /v4/organizations/{organization_id}/ | Modify organization



## AddCredentials

> V4GenericResponse AddCredentials(ctx, authorization, organizationId, v4AddCredentialsRequest, optional)
Set credentials

Add a set of credentials to the organization allowing the creation and operation of clusters within a cloud provider account/subscription.  The actual type of these credentials depends on the cloud provider the installation is running on. AWS and Azure are currently supported.  Credentials in an organization are immutable. Each organization can only have one set of credentials.  Once credentials have been set for an organization, they are used for every new cluster that will be created for the organization.  ### Example request body for AWS  ```json {   \"provider\": \"aws\",   \"aws\": {     \"roles\": {       \"admin\": \"arn:aws:iam::123456789012:role/GiantSwarmAdmin\",       \"awsoperator\": \"arn:aws:iam::123456789012:role/GiantSwarmAWSOperator\"     }   } } ```  ### Example request body for Azure  ```json {   \"provider\": \"azure\",   \"azure\": {     \"credential\": {       \"client_id\": \"c93bf55e-5bf7-4966-ad2b-e6f6e7721d50\",       \"secret_key\": \"720e38f7-3af4-463c-9313-abcdf2ead612\",       \"subscription_id\": \"b388b7c7-4479-4040-9ac5-1e13edd6b1cd\",       \"tenant_id\": \"3dd2e94a-92ba-434c-99be-32bb65864a99\"     }   } } ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**organizationId** | **string**| An ID for the organization. This ID must be unique and match this regular expression: ^[a-z0-9_]{4,30}$  | 
**v4AddCredentialsRequest** | [**V4AddCredentialsRequest**](V4AddCredentialsRequest.md)|  | 
 **optional** | ***AddCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddCredentialsOpts struct


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


## AddOrganization

> V4Organization AddOrganization(ctx, authorization, organizationId, v4Organization, optional)
Create an organization

This operation allows a user to create an organization. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**organizationId** | **string**| An ID for the organization. This ID must be unique and match this regular expression: ^[a-z0-9_]{4,30}$  | 
**v4Organization** | [**V4Organization**](V4Organization.md)|  | 
 **optional** | ***AddOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddOrganizationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V4Organization**](V4Organization.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteOrganization

> V4GenericResponse DeleteOrganization(ctx, authorization, organizationId, optional)
Delete an organization

This operation allows a user to delete an organization that they are a member of. Admin users can delete any organization. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**organizationId** | **string**| An ID for the organization. This ID must be unique and match this regular expression: ^[a-z0-9_]{4,30}$  | 
 **optional** | ***DeleteOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteOrganizationOpts struct


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


## GetCredential

> V4GetCredentialResponse GetCredential(ctx, authorization, organizationId, credentialId, optional)
Get credential details

Returns details for a particular set of credentials, identified by its ID. The returned data does not contain any secrets (i. e. passphrase, secret key). For more information on credentials, see [Set credentials](#operation/addCredentials).  ### Example response body for AWS  ```json {   \"id\": \"a1b2c3\",   \"provider\": \"aws\",   \"aws\": {     \"roles\": {       \"admin\": \"arn:aws:iam::123456789012:role/GiantSwarmAdmin\",       \"awsoperator\": \"arn:aws:iam::123456789012:role/GiantSwarmAWSOperator\"     }   } } ```  ### Example response body for Azure  ```json {   \"id\": \"a1b2c3\",   \"provider\": \"azure\",   \"azure\": {     \"credential\": {       \"client_id\": \"c93bf55e-5bf7-4966-ad2b-e6f6e7721d50\",       \"subscription_id\": \"b388b7c7-4479-4040-9ac5-1e13edd6b1cd\",       \"tenant_id\": \"3dd2e94a-92ba-434c-99be-32bb65864a99\"     }   } } ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**organizationId** | **string**| An ID for the organization. This ID must be unique and match this regular expression: ^[a-z0-9_]{4,30}$  | 
**credentialId** | **string**| Unique ID of a credential set | 
 **optional** | ***GetCredentialOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCredentialOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V4GetCredentialResponse**](V4GetCredentialResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCredentials

> []V4GetCredentialResponse GetCredentials(ctx, authorization, organizationId, optional)
Get credentials

Returns credentials for an organization, if available. For more information on credentials, see [Set credentials](#operation/addCredentials).  Here is another paragraph.  ### Example response body for AWS  ```json [   {     \"id\": \"a1b2c3\",     \"provider\": \"aws\",     \"aws\": {       \"roles\": {         \"admin\": \"arn:aws:iam::123456789012:role/GiantSwarmAdmin\",         \"awsoperator\": \"arn:aws:iam::123456789012:role/GiantSwarmAWSOperator\"       }     }   } ] ```  ### Example response body for Azure  ```json [   {     \"id\": \"a1b2c3\",     \"provider\": \"azure\",     \"azure\": {       \"credential\": {         \"client_id\": \"c93bf55e-5bf7-4966-ad2b-e6f6e7721d50\",         \"subscription_id\": \"b388b7c7-4479-4040-9ac5-1e13edd6b1cd\",         \"tenant_id\": \"3dd2e94a-92ba-434c-99be-32bb65864a99\"       }     }   } ] ``` 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**organizationId** | **string**| An ID for the organization. This ID must be unique and match this regular expression: ^[a-z0-9_]{4,30}$  | 
 **optional** | ***GetCredentialsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCredentialsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**[]V4GetCredentialResponse**](V4GetCredentialResponse.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganization

> V4Organization GetOrganization(ctx, authorization, organizationId, optional)
Get organization details

This operation fetches organization details. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**organizationId** | **string**| An ID for the organization. This ID must be unique and match this regular expression: ^[a-z0-9_]{4,30}$  | 
 **optional** | ***GetOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOrganizationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V4Organization**](V4Organization.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetOrganizations

> []V4OrganizationListItem GetOrganizations(ctx, authorization, optional)
Get organizations

This operation allows to fetch a list of organizations the user is a member of. In the case of an admin user, the result includes all existing organizations. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
 **optional** | ***GetOrganizationsOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetOrganizationsOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**[]V4OrganizationListItem**](V4OrganizationListItem.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyOrganization

> V4Organization ModifyOrganization(ctx, authorization, organizationId, inlineObject, optional)
Modify organization

This operation allows you to modify an existing organization. You must be a member of the organization or an admin in order to use this endpoint.  The following attributes can be modified:  - `members`: By modifying the array of members, members can be added to or removed from the organization  The request body must conform with the [JSON Patch Merge (RFC 7386)](https://tools.ietf.org/html/rfc7386) standard. Requests have to be sent with the `Content-Type: application/merge-patch+json` header.  The full request must be valid before it will be executed, currently this means every member you attempt to add to the organization must actually exist in the system. If any member you attempt to add is invalid, the entire patch operation will fail, no members will be added or removed, and an error message will explain which members in your request are invalid. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**organizationId** | **string**| An ID for the organization. This ID must be unique and match this regular expression: ^[a-z0-9_]{4,30}$  | 
**inlineObject** | [**InlineObject**](InlineObject.md)|  | 
 **optional** | ***ModifyOrganizationOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ModifyOrganizationOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V4Organization**](V4Organization.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

