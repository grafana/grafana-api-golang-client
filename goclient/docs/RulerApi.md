# \RulerApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RouteDeleteGrafanaRuleGroupConfig**](RulerApi.md#RouteDeleteGrafanaRuleGroupConfig) | **Delete** /ruler/grafana/api/v1/rules/{Namespace}/{Groupname} | 
[**RouteDeleteNamespaceGrafanaRulesConfig**](RulerApi.md#RouteDeleteNamespaceGrafanaRulesConfig) | **Delete** /ruler/grafana/api/v1/rules/{Namespace} | 
[**RouteDeleteNamespaceRulesConfig**](RulerApi.md#RouteDeleteNamespaceRulesConfig) | **Delete** /ruler/{Recipient}/api/v1/rules/{Namespace} | 
[**RouteDeleteRuleGroupConfig**](RulerApi.md#RouteDeleteRuleGroupConfig) | **Delete** /ruler/{Recipient}/api/v1/rules/{Namespace}/{Groupname} | 
[**RouteGetGrafanaRuleGroupConfig**](RulerApi.md#RouteGetGrafanaRuleGroupConfig) | **Get** /ruler/grafana/api/v1/rules/{Namespace}/{Groupname} | 
[**RouteGetGrafanaRulesConfig**](RulerApi.md#RouteGetGrafanaRulesConfig) | **Get** /ruler/grafana/api/v1/rules | 
[**RouteGetNamespaceGrafanaRulesConfig**](RulerApi.md#RouteGetNamespaceGrafanaRulesConfig) | **Get** /ruler/grafana/api/v1/rules/{Namespace} | 
[**RouteGetNamespaceRulesConfig**](RulerApi.md#RouteGetNamespaceRulesConfig) | **Get** /ruler/{Recipient}/api/v1/rules/{Namespace} | 
[**RouteGetRulegGroupConfig**](RulerApi.md#RouteGetRulegGroupConfig) | **Get** /ruler/{Recipient}/api/v1/rules/{Namespace}/{Groupname} | 
[**RouteGetRulesConfig**](RulerApi.md#RouteGetRulesConfig) | **Get** /ruler/{Recipient}/api/v1/rules | 
[**RoutePostNameGrafanaRulesConfig**](RulerApi.md#RoutePostNameGrafanaRulesConfig) | **Post** /ruler/grafana/api/v1/rules/{Namespace} | 
[**RoutePostNameRulesConfig**](RulerApi.md#RoutePostNameRulesConfig) | **Post** /ruler/{Recipient}/api/v1/rules/{Namespace} | 


# **RouteDeleteGrafanaRuleGroupConfig**
> AckModel RouteDeleteGrafanaRuleGroupConfig(ctx, namespace, groupname)


Delete rule group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **groupname** | **string**|  | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteDeleteNamespaceGrafanaRulesConfig**
> AckModel RouteDeleteNamespaceGrafanaRulesConfig(ctx, namespace)


Delete namespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteDeleteNamespaceRulesConfig**
> AckModel RouteDeleteNamespaceRulesConfig(ctx, recipient, namespace)


Delete namespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 
  **namespace** | **string**|  | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteDeleteRuleGroupConfig**
> AckModel RouteDeleteRuleGroupConfig(ctx, recipient, namespace, groupname)


Delete rule group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 
  **namespace** | **string**|  | 
  **groupname** | **string**|  | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetGrafanaRuleGroupConfig**
> RuleGroupConfigResponseModel RouteGetGrafanaRuleGroupConfig(ctx, namespace, groupname)


Get rule group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
  **groupname** | **string**|  | 

### Return type

[**RuleGroupConfigResponseModel**](RuleGroupConfigResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetGrafanaRulesConfig**
> NamespaceConfigResponseModel RouteGetGrafanaRulesConfig(ctx, optional)


List rule groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***RulerApiRouteGetGrafanaRulesConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RulerApiRouteGetGrafanaRulesConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **dashboardUID** | **optional.String**|  | 
 **panelID** | **optional.Int64**|  | 

### Return type

[**NamespaceConfigResponseModel**](NamespaceConfigResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetNamespaceGrafanaRulesConfig**
> NamespaceConfigResponseModel RouteGetNamespaceGrafanaRulesConfig(ctx, namespace)


Get rule groups by namespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 

### Return type

[**NamespaceConfigResponseModel**](NamespaceConfigResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetNamespaceRulesConfig**
> NamespaceConfigResponseModel RouteGetNamespaceRulesConfig(ctx, recipient, namespace)


Get rule groups by namespace

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 
  **namespace** | **string**|  | 

### Return type

[**NamespaceConfigResponseModel**](NamespaceConfigResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetRulegGroupConfig**
> RuleGroupConfigResponseModel RouteGetRulegGroupConfig(ctx, recipient, namespace, groupname)


Get rule group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 
  **namespace** | **string**|  | 
  **groupname** | **string**|  | 

### Return type

[**RuleGroupConfigResponseModel**](RuleGroupConfigResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetRulesConfig**
> NamespaceConfigResponseModel RouteGetRulesConfig(ctx, recipient, optional)


List rule groups

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 
 **optional** | ***RulerApiRouteGetRulesConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RulerApiRouteGetRulesConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dashboardUID** | **optional.String**|  | 
 **panelID** | **optional.Int64**|  | 

### Return type

[**NamespaceConfigResponseModel**](NamespaceConfigResponse.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePostNameGrafanaRulesConfig**
> AckModel RoutePostNameGrafanaRulesConfig(ctx, namespace, optional)


Creates or updates a rule group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **namespace** | **string**|  | 
 **optional** | ***RulerApiRoutePostNameGrafanaRulesConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RulerApiRoutePostNameGrafanaRulesConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PostableRuleGroupConfigModel**](PostableRuleGroupConfigModel.md)|  | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePostNameRulesConfig**
> AckModel RoutePostNameRulesConfig(ctx, recipient, namespace, optional)


Creates or updates a rule group

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 
  **namespace** | **string**|  | 
 **optional** | ***RulerApiRoutePostNameRulesConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a RulerApiRoutePostNameRulesConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of PostableRuleGroupConfigModel**](PostableRuleGroupConfigModel.md)|  | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json, application/yaml
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

