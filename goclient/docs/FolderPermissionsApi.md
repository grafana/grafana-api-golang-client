# \FolderPermissionsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFolderPermissions**](FolderPermissionsApi.md#GetFolderPermissions) | **Get** /folders/{folder_uid}/permissions | Gets all existing permissions for the folder with the given &#x60;uid&#x60;.
[**UpdateFolderPermissions**](FolderPermissionsApi.md#UpdateFolderPermissions) | **Post** /folders/{folder_uid}/permissions | Updates permissions for a folder. This operation will remove existing permissions if they’re not included in the request.


# **GetFolderPermissions**
> []DashboardAclInfoDtoModel GetFolderPermissions(ctx, folderUid)
Gets all existing permissions for the folder with the given `uid`.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **folderUid** | **string**|  | 

### Return type

[**[]DashboardAclInfoDtoModel**](DashboardAclInfoDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateFolderPermissions**
> SuccessResponseBodyModel UpdateFolderPermissions(ctx, body, folderUid)
Updates permissions for a folder. This operation will remove existing permissions if they’re not included in the request.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateDashboardAclCommandModel**](UpdateDashboardAclCommandModel.md)|  | 
  **folderUid** | **string**|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

