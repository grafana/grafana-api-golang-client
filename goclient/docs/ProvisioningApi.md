# \ProvisioningApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RouteDeleteContactpoints**](ProvisioningApi.md#RouteDeleteContactpoints) | **Delete** /provisioning/contact-points/{ID} | Delete a contact point.
[**RouteGetContactpoints**](ProvisioningApi.md#RouteGetContactpoints) | **Get** /provisioning/contact-points | Get all the contact points.
[**RouteGetPolicyTree**](ProvisioningApi.md#RouteGetPolicyTree) | **Get** /provisioning/policies | Get the notification policy tree.
[**RoutePostContactpoints**](ProvisioningApi.md#RoutePostContactpoints) | **Post** /provisioning/contact-points | Create a contact point.
[**RoutePostPolicyTree**](ProvisioningApi.md#RoutePostPolicyTree) | **Post** /provisioning/policies | Sets the notification policy tree.
[**RoutePutContactpoints**](ProvisioningApi.md#RoutePutContactpoints) | **Put** /provisioning/contact-points | Update an existing contact point.


# **RouteDeleteContactpoints**
> RouteDeleteContactpoints(ctx, )
Delete a contact point.

### Required Parameters
This endpoint does not need any parameter.

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetContactpoints**
> RouteModel RouteGetContactpoints(ctx, )
Get all the contact points.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RouteModel**](Route.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetPolicyTree**
> RouteModel RouteGetPolicyTree(ctx, )
Get the notification policy tree.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**RouteModel**](Route.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePostContactpoints**
> RoutePostContactpoints(ctx, optional)
Create a contact point.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProvisioningApiRoutePostContactpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRoutePostContactpointsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EmbeddedContactPointModel**](EmbeddedContactPointModel.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePostPolicyTree**
> RoutePostPolicyTree(ctx, optional)
Sets the notification policy tree.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProvisioningApiRoutePostPolicyTreeOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRoutePostPolicyTreeOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of RouteModel**](RouteModel.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePutContactpoints**
> RoutePutContactpoints(ctx, optional)
Update an existing contact point.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ProvisioningApiRoutePutContactpointsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a ProvisioningApiRoutePutContactpointsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of EmbeddedContactPointModel**](EmbeddedContactPointModel.md)|  | 

### Return type

 (empty response body)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

