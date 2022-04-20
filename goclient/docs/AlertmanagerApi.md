# \AlertmanagerApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**RouteCreateGrafanaSilence**](AlertmanagerApi.md#RouteCreateGrafanaSilence) | **Post** /alertmanager/grafana/api/v2/silences | 
[**RouteCreateSilence**](AlertmanagerApi.md#RouteCreateSilence) | **Post** /alertmanager/{Recipient}/api/v2/silences | 
[**RouteDeleteAlertingConfig**](AlertmanagerApi.md#RouteDeleteAlertingConfig) | **Delete** /alertmanager/{Recipient}/config/api/v1/alerts | 
[**RouteDeleteGrafanaAlertingConfig**](AlertmanagerApi.md#RouteDeleteGrafanaAlertingConfig) | **Delete** /alertmanager/grafana/config/api/v1/alerts | 
[**RouteDeleteGrafanaSilence**](AlertmanagerApi.md#RouteDeleteGrafanaSilence) | **Delete** /alertmanager/grafana/api/v2/silence/{SilenceId} | 
[**RouteDeleteSilence**](AlertmanagerApi.md#RouteDeleteSilence) | **Delete** /alertmanager/{Recipient}/api/v2/silence/{SilenceId} | 
[**RouteGetAMAlertGroups**](AlertmanagerApi.md#RouteGetAMAlertGroups) | **Get** /alertmanager/{Recipient}/api/v2/alerts/groups | 
[**RouteGetAMAlerts**](AlertmanagerApi.md#RouteGetAMAlerts) | **Get** /alertmanager/{Recipient}/api/v2/alerts | 
[**RouteGetAMStatus**](AlertmanagerApi.md#RouteGetAMStatus) | **Get** /alertmanager/{Recipient}/api/v2/status | 
[**RouteGetAlertingConfig**](AlertmanagerApi.md#RouteGetAlertingConfig) | **Get** /alertmanager/{Recipient}/config/api/v1/alerts | 
[**RouteGetGrafanaAMAlertGroups**](AlertmanagerApi.md#RouteGetGrafanaAMAlertGroups) | **Get** /alertmanager/grafana/api/v2/alerts/groups | 
[**RouteGetGrafanaAMAlerts**](AlertmanagerApi.md#RouteGetGrafanaAMAlerts) | **Get** /alertmanager/grafana/api/v2/alerts | 
[**RouteGetGrafanaAMStatus**](AlertmanagerApi.md#RouteGetGrafanaAMStatus) | **Get** /alertmanager/grafana/api/v2/status | 
[**RouteGetGrafanaAlertingConfig**](AlertmanagerApi.md#RouteGetGrafanaAlertingConfig) | **Get** /alertmanager/grafana/config/api/v1/alerts | 
[**RouteGetGrafanaSilence**](AlertmanagerApi.md#RouteGetGrafanaSilence) | **Get** /alertmanager/grafana/api/v2/silence/{SilenceId} | 
[**RouteGetGrafanaSilences**](AlertmanagerApi.md#RouteGetGrafanaSilences) | **Get** /alertmanager/grafana/api/v2/silences | 
[**RouteGetSilence**](AlertmanagerApi.md#RouteGetSilence) | **Get** /alertmanager/{Recipient}/api/v2/silence/{SilenceId} | 
[**RouteGetSilences**](AlertmanagerApi.md#RouteGetSilences) | **Get** /alertmanager/{Recipient}/api/v2/silences | 
[**RoutePostAMAlerts**](AlertmanagerApi.md#RoutePostAMAlerts) | **Post** /alertmanager/{Recipient}/api/v2/alerts | 
[**RoutePostAlertingConfig**](AlertmanagerApi.md#RoutePostAlertingConfig) | **Post** /alertmanager/{Recipient}/config/api/v1/alerts | 
[**RoutePostGrafanaAMAlerts**](AlertmanagerApi.md#RoutePostGrafanaAMAlerts) | **Post** /alertmanager/grafana/api/v2/alerts | 
[**RoutePostGrafanaAlertingConfig**](AlertmanagerApi.md#RoutePostGrafanaAlertingConfig) | **Post** /alertmanager/grafana/config/api/v1/alerts | 
[**RoutePostTestGrafanaReceivers**](AlertmanagerApi.md#RoutePostTestGrafanaReceivers) | **Post** /alertmanager/grafana/config/api/v1/receivers/test | Test Grafana managed receivers without saving them.
[**RoutePostTestReceivers**](AlertmanagerApi.md#RoutePostTestReceivers) | **Post** /alertmanager/{Recipient}/config/api/v1/receivers/test | Test Grafana managed receivers without saving them.


# **RouteCreateGrafanaSilence**
> GettableSilenceModel RouteCreateGrafanaSilence(ctx, optional)


create silence

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertmanagerApiRouteCreateGrafanaSilenceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertmanagerApiRouteCreateGrafanaSilenceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **silence** | [**optional.Interface of PostableSilenceModel**](PostableSilenceModel.md)|  | 

### Return type

[**GettableSilenceModel**](gettableSilence.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteCreateSilence**
> GettableSilenceModel RouteCreateSilence(ctx, recipient, optional)


create silence

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 
 **optional** | ***AlertmanagerApiRouteCreateSilenceOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertmanagerApiRouteCreateSilenceOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **silence** | [**optional.Interface of PostableSilenceModel**](PostableSilenceModel.md)|  | 

### Return type

[**GettableSilenceModel**](gettableSilence.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteDeleteAlertingConfig**
> AckModel RouteDeleteAlertingConfig(ctx, recipient)


deletes the Alerting config for a tenant

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteDeleteGrafanaAlertingConfig**
> AckModel RouteDeleteGrafanaAlertingConfig(ctx, )


deletes the Alerting config for a tenant

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteDeleteGrafanaSilence**
> AckModel RouteDeleteGrafanaSilence(ctx, silenceId)


delete silence

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **silenceId** | **string**|  | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteDeleteSilence**
> AckModel RouteDeleteSilence(ctx, silenceId, recipient)


delete silence

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **silenceId** | **string**|  | 
  **recipient** | **int64**| Recipient should be the numeric datasource id | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetAMAlertGroups**
> AlertGroupsModel RouteGetAMAlertGroups(ctx, recipient, optional)


get alertmanager alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 
 **optional** | ***AlertmanagerApiRouteGetAMAlertGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertmanagerApiRouteGetAMAlertGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **active** | **optional.Bool**| Show active alerts | [default to true]
 **silenced** | **optional.Bool**| Show silenced alerts | [default to true]
 **inhibited** | **optional.Bool**| Show inhibited alerts | [default to true]
 **filter** | [**optional.Interface of []string**](string.md)| A list of matchers to filter alerts by | 
 **receiver** | **optional.String**| A regex matching receivers to filter alerts by | 

### Return type

[**AlertGroupsModel**](alertGroups.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetAMAlerts**
> GettableAlertsModel RouteGetAMAlerts(ctx, recipient, optional)


get alertmanager alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 
 **optional** | ***AlertmanagerApiRouteGetAMAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertmanagerApiRouteGetAMAlertsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **active** | **optional.Bool**| Show active alerts | [default to true]
 **silenced** | **optional.Bool**| Show silenced alerts | [default to true]
 **inhibited** | **optional.Bool**| Show inhibited alerts | [default to true]
 **filter** | [**optional.Interface of []string**](string.md)| A list of matchers to filter alerts by | 
 **receiver** | **optional.String**| A regex matching receivers to filter alerts by | 

### Return type

[**GettableAlertsModel**](gettableAlerts.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetAMStatus**
> GettableStatusModel RouteGetAMStatus(ctx, recipient)


get alertmanager status and configuration

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 

### Return type

[**GettableStatusModel**](GettableStatus.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetAlertingConfig**
> GettableUserConfigModel RouteGetAlertingConfig(ctx, recipient)


gets an Alerting config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 

### Return type

[**GettableUserConfigModel**](GettableUserConfig.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetGrafanaAMAlertGroups**
> AlertGroupsModel RouteGetGrafanaAMAlertGroups(ctx, optional)


get alertmanager alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertmanagerApiRouteGetGrafanaAMAlertGroupsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertmanagerApiRouteGetGrafanaAMAlertGroupsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **optional.Bool**| Show active alerts | [default to true]
 **silenced** | **optional.Bool**| Show silenced alerts | [default to true]
 **inhibited** | **optional.Bool**| Show inhibited alerts | [default to true]
 **filter** | [**optional.Interface of []string**](string.md)| A list of matchers to filter alerts by | 
 **receiver** | **optional.String**| A regex matching receivers to filter alerts by | 

### Return type

[**AlertGroupsModel**](alertGroups.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetGrafanaAMAlerts**
> GettableAlertsModel RouteGetGrafanaAMAlerts(ctx, optional)


get alertmanager alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertmanagerApiRouteGetGrafanaAMAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertmanagerApiRouteGetGrafanaAMAlertsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **active** | **optional.Bool**| Show active alerts | [default to true]
 **silenced** | **optional.Bool**| Show silenced alerts | [default to true]
 **inhibited** | **optional.Bool**| Show inhibited alerts | [default to true]
 **filter** | [**optional.Interface of []string**](string.md)| A list of matchers to filter alerts by | 
 **receiver** | **optional.String**| A regex matching receivers to filter alerts by | 

### Return type

[**GettableAlertsModel**](gettableAlerts.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetGrafanaAMStatus**
> GettableStatusModel RouteGetGrafanaAMStatus(ctx, )


get alertmanager status and configuration

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GettableStatusModel**](GettableStatus.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetGrafanaAlertingConfig**
> GettableUserConfigModel RouteGetGrafanaAlertingConfig(ctx, )


gets an Alerting config

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**GettableUserConfigModel**](GettableUserConfig.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetGrafanaSilence**
> GettableSilenceModel RouteGetGrafanaSilence(ctx, silenceId)


get silence

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **silenceId** | **string**|  | 

### Return type

[**GettableSilenceModel**](gettableSilence.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetGrafanaSilences**
> GettableSilencesModel RouteGetGrafanaSilences(ctx, optional)


get silences

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertmanagerApiRouteGetGrafanaSilencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertmanagerApiRouteGetGrafanaSilencesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **filter** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**GettableSilencesModel**](gettableSilences.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetSilence**
> GettableSilenceModel RouteGetSilence(ctx, silenceId, recipient)


get silence

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **silenceId** | **string**|  | 
  **recipient** | **int64**| Recipient should be the numeric datasource id | 

### Return type

[**GettableSilenceModel**](gettableSilence.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RouteGetSilences**
> GettableSilencesModel RouteGetSilences(ctx, recipient, optional)


get silences

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 
 **optional** | ***AlertmanagerApiRouteGetSilencesOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertmanagerApiRouteGetSilencesOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **filter** | [**optional.Interface of []string**](string.md)|  | 

### Return type

[**GettableSilencesModel**](gettableSilences.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePostAMAlerts**
> AckModel RoutePostAMAlerts(ctx, recipient, optional)


create alertmanager alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 
 **optional** | ***AlertmanagerApiRoutePostAMAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertmanagerApiRoutePostAMAlertsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **postableAlerts** | [**optional.Interface of []PostableAlertModel**](postableAlert.md)|  | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePostAlertingConfig**
> AckModel RoutePostAlertingConfig(ctx, recipient, optional)


sets an Alerting config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 
 **optional** | ***AlertmanagerApiRoutePostAlertingConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertmanagerApiRoutePostAlertingConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of PostableUserConfigModel**](PostableUserConfigModel.md)|  | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePostGrafanaAMAlerts**
> AckModel RoutePostGrafanaAMAlerts(ctx, optional)


create alertmanager alerts

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertmanagerApiRoutePostGrafanaAMAlertsOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertmanagerApiRoutePostGrafanaAMAlertsOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **postableAlerts** | [**optional.Interface of []PostableAlertModel**](postableAlert.md)|  | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePostGrafanaAlertingConfig**
> AckModel RoutePostGrafanaAlertingConfig(ctx, optional)


sets an Alerting config

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertmanagerApiRoutePostGrafanaAlertingConfigOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertmanagerApiRoutePostGrafanaAlertingConfigOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of PostableUserConfigModel**](PostableUserConfigModel.md)|  | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePostTestGrafanaReceivers**
> AckModel RoutePostTestGrafanaReceivers(ctx, optional)
Test Grafana managed receivers without saving them.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AlertmanagerApiRoutePostTestGrafanaReceiversOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertmanagerApiRoutePostTestGrafanaReceiversOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**optional.Interface of TestReceiversConfigBodyParamsModel**](TestReceiversConfigBodyParamsModel.md)|  | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RoutePostTestReceivers**
> AckModel RoutePostTestReceivers(ctx, recipient, optional)
Test Grafana managed receivers without saving them.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **recipient** | **int64**| Recipient should be the numeric datasource id | 
 **optional** | ***AlertmanagerApiRoutePostTestReceiversOpts** | optional parameters | nil if no parameters

### Optional Parameters
Optional parameters are passed through a pointer to a AlertmanagerApiRoutePostTestReceiversOpts struct

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of TestReceiversConfigBodyParamsModel**](TestReceiversConfigBodyParamsModel.md)|  | 

### Return type

[**AckModel**](Ack.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

