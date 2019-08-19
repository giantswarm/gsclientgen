# \UsersApi

All URIs are relative to *https://localhost:8000*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUser**](UsersApi.md#CreateUser) | **Put** /v4/users/{email}/ | Create user
[**DeleteUser**](UsersApi.md#DeleteUser) | **Delete** /v4/users/{email}/ | Delete user
[**GetCurrentUser**](UsersApi.md#GetCurrentUser) | **Get** /v4/user/ | Get current user
[**GetUser**](UsersApi.md#GetUser) | **Get** /v4/users/{email}/ | Get user
[**GetUsers**](UsersApi.md#GetUsers) | **Get** /v4/users/ | Get users
[**ModifyPassword**](UsersApi.md#ModifyPassword) | **Post** /v4/users/{email}/password/ | Modify password
[**ModifyUser**](UsersApi.md#ModifyUser) | **Patch** /v4/users/{email}/ | Modify user



## CreateUser

> V4GenericResponse CreateUser(ctx, authorization, email, v4CreateUserRequest, optional)
Create user

Creates a users in the system. Currently this endpoint is only available to users with admin permissions. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**email** | **string**| The user&#39;s email address | 
**v4CreateUserRequest** | [**V4CreateUserRequest**](V4CreateUserRequest.md)| User account details | 
 **optional** | ***CreateUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUserOpts struct


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


## DeleteUser

> V4GenericResponse DeleteUser(ctx, authorization, email, optional)
Delete user

Deletes a users in the system. Currently this endpoint is only available to users with admin permissions. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**email** | **string**| The user&#39;s email address | 
 **optional** | ***DeleteUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteUserOpts struct


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


## GetCurrentUser

> V4UserListItem GetCurrentUser(ctx, authorization, optional)
Get current user

Returns details about the currently authenticated user 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
 **optional** | ***GetCurrentUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCurrentUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V4UserListItem**](V4UserListItem.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUser

> V4UserListItem GetUser(ctx, authorization, email, optional)
Get user

Returns details about a specific user 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**email** | **string**| The user&#39;s email address | 
 **optional** | ***GetUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V4UserListItem**](V4UserListItem.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsers

> []V4UserListItem GetUsers(ctx, authorization, optional)
Get users

Returns a list of all users in the system. Currently this endpoint is only available to users with admin permissions. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
 **optional** | ***GetUsersOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUsersOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**[]V4UserListItem**](V4UserListItem.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ModifyPassword

> V4GenericResponse ModifyPassword(ctx, authorization, email, v4ModifyUserPasswordRequest, optional)
Modify password

This operation allows you to change your password. Admins are able to change passwords of other users. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**email** | **string**| The user&#39;s email address | 
**v4ModifyUserPasswordRequest** | [**V4ModifyUserPasswordRequest**](V4ModifyUserPasswordRequest.md)| Modify password request | 
 **optional** | ***ModifyPasswordOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ModifyPasswordOpts struct


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


## ModifyUser

> V4UserListItem ModifyUser(ctx, authorization, email, v4ModifyUserRequest, optional)
Modify user

This operation allows you to change details of a given user. Only administrators can edit accounts of other users. 

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| As described in the [authentication](#section/Authentication) section  | 
**email** | **string**| The user&#39;s email address | 
**v4ModifyUserRequest** | [**V4ModifyUserRequest**](V4ModifyUserRequest.md)| User account details | 
 **optional** | ***ModifyUserOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ModifyUserOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xRequestID** | **optional.String**| A randomly generated key that can be used to track a request throughout services of Giant Swarm.  | 
 **xGiantSwarmActivity** | **optional.String**| Name of an activity to track, like \&quot;list-clusters\&quot;. This allows to analyze several API requests sent in context and gives an idea on the purpose.  | 
 **xGiantSwarmCmdLine** | **optional.String**| If activity has been issued by a CLI, this header can contain the command line  | 

### Return type

[**V4UserListItem**](V4UserListItem.md)

### Authorization

[AuthorizationHeaderToken](../README.md#AuthorizationHeaderToken)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

