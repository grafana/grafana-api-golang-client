# \DashboardsApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CalcDashboardDiff**](DashboardsApi.md#CalcDashboardDiff) | **Post** /dashboards/calculate-diff | Perform diff on two dashboards.
[**DeleteDashboardByUID**](DashboardsApi.md#DeleteDashboardByUID) | **Delete** /dashboards/uid/{uid} | Delete dashboard by uid.
[**GetDashboardByUID**](DashboardsApi.md#GetDashboardByUID) | **Get** /dashboards/uid/{uid} | Get dashboard by uid.
[**GetDashboardTags**](DashboardsApi.md#GetDashboardTags) | **Get** /dashboards/tags | Get all dashboards tags of an organisation.
[**GetHomeDashboard**](DashboardsApi.md#GetHomeDashboard) | **Get** /dashboards/home | Get home dashboard.
[**ImportDashboard**](DashboardsApi.md#ImportDashboard) | **Post** /dashboards/import | Import dashboard.
[**PostDashboard**](DashboardsApi.md#PostDashboard) | **Post** /dashboards/db | Create / Update dashboard
[**TrimDashboard**](DashboardsApi.md#TrimDashboard) | **Post** /dashboards/trim | Trim defaults from dashboard.


# **CalcDashboardDiff**
> []int32 CalcDashboardDiff(ctx, body)
Perform diff on two dashboards.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**BodyModel**](BodyModel.md)|  | 

### Return type

**[]int32**

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json, text/html

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **DeleteDashboardByUID**
> InlineResponse2005Model DeleteDashboardByUID(ctx, uid)
Delete dashboard by uid.

Will delete the dashboard given the specified unique identifier (uid).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**|  | 

### Return type

[**InlineResponse2005Model**](inline_response_200_5.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardByUID**
> DashboardFullWithMetaModel GetDashboardByUID(ctx, uid)
Get dashboard by uid.

Will return the dashboard given the dashboard unique identifier (uid).

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **uid** | **string**|  | 

### Return type

[**DashboardFullWithMetaModel**](DashboardFullWithMeta.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetDashboardTags**
> []DashboardTagCloudItemModel GetDashboardTags(ctx, )
Get all dashboards tags of an organisation.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]DashboardTagCloudItemModel**](DashboardTagCloudItem.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetHomeDashboard**
> GetHomeDashboardResponseModel GetHomeDashboard(ctx, )
Get home dashboard.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GetHomeDashboardResponseModel**](GetHomeDashboardResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ImportDashboard**
> ImportDashboardResponseModel ImportDashboard(ctx, body)
Import dashboard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ImportDashboardRequestModel**](ImportDashboardRequestModel.md)|  | 

### Return type

[**ImportDashboardResponseModel**](ImportDashboardResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **PostDashboard**
> InlineResponse2004Model PostDashboard(ctx, body)
Create / Update dashboard

Creates a new dashboard or updates an existing dashboard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**SaveDashboardCommandModel**](SaveDashboardCommandModel.md)|  | 

### Return type

[**InlineResponse2004Model**](inline_response_200_4.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **TrimDashboard**
> TrimDashboardFullWithMetaModel TrimDashboard(ctx, body)
Trim defaults from dashboard.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**TrimDashboardCommandModel**](TrimDashboardCommandModel.md)|  | 

### Return type

[**TrimDashboardFullWithMetaModel**](TrimDashboardFullWithMeta.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

