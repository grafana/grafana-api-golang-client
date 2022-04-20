# \UserPreferencesApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserPreferences**](UserPreferencesApi.md#GetUserPreferences) | **Get** /user/preferences | Get user preferences.
[**PatchUserPreferences**](UserPreferencesApi.md#PatchUserPreferences) | **Patch** /user/preferences | Patch user preferences.
[**UpdateUserPreferences**](UserPreferencesApi.md#UpdateUserPreferences) | **Put** /user/preferences | Update user preferences.


# **GetUserPreferences**
> PrefsModel GetUserPreferences(ctx, )
Get user preferences.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**PrefsModel**](Prefs.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PatchUserPreferences**
> SuccessResponseBodyModel PatchUserPreferences(ctx, body)
Patch user preferences.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**PatchPrefsCmdModel**](PatchPrefsCmdModel.md)|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateUserPreferences**
> SuccessResponseBodyModel UpdateUserPreferences(ctx, body)
Update user preferences.

Omitting a key (`theme`, `homeDashboardId`, `timezone`) will cause the current value to be replaced with the system default value.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdatePrefsCmdModel**](UpdatePrefsCmdModel.md)|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

