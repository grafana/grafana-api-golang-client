# \SignedInUserApi

All URIs are relative to *http://localhost/api*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangeUserPassword**](SignedInUserApi.md#ChangeUserPassword) | **Put** /user/password | Change Password.
[**ClearHelpFlags**](SignedInUserApi.md#ClearHelpFlags) | **Get** /user/helpflags/clear | Clear user help flag.
[**GetSignedInUser**](SignedInUserApi.md#GetSignedInUser) | **Get** /user | Get signed in User.
[**GetSignedInUserAuthTokens**](SignedInUserApi.md#GetSignedInUserAuthTokens) | **Get** /user/auth-tokens | Auth tokens of the actual User.
[**GetSignedInUserOrgList**](SignedInUserApi.md#GetSignedInUserOrgList) | **Get** /user/orgs | Organizations of the actual User.
[**GetSignedInUserTeamList**](SignedInUserApi.md#GetSignedInUserTeamList) | **Get** /user/teams | Teams that the actual User is member of.
[**GetUserQuotas**](SignedInUserApi.md#GetUserQuotas) | **Get** /user/quotas | Fetch user quota.
[**RevokeSignedINAuthTokenCmd**](SignedInUserApi.md#RevokeSignedINAuthTokenCmd) | **Post** /user/revoke-auth-token | Revoke an auth token of the actual User.
[**SetHelpFlag**](SignedInUserApi.md#SetHelpFlag) | **Put** /user/helpflags/{flag_id} | Set user help flag.
[**StarDashboard**](SignedInUserApi.md#StarDashboard) | **Post** /user/stars/dashboard/{dashboard_id} | Star a dashboard.
[**UnstarDashboard**](SignedInUserApi.md#UnstarDashboard) | **Delete** /user/stars/dashboard/{dashboard_id} | Unstar a dashboard.
[**UpdateSignedInUser**](SignedInUserApi.md#UpdateSignedInUser) | **Put** /user | Update signed in User.
[**UserSetUsingOrg**](SignedInUserApi.md#UserSetUsingOrg) | **Post** /user/using/{org_id} | Switch user context for signed in user.


# **ChangeUserPassword**
> SuccessResponseBodyModel ChangeUserPassword(ctx, body)
Change Password.

Changes the password for the user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**ChangeUserPasswordCommandModel**](ChangeUserPasswordCommandModel.md)| To change the email, name, login, theme, provide another one. | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **ClearHelpFlags**
> InlineResponse20016Model ClearHelpFlags(ctx, )
Clear user help flag.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**InlineResponse20016Model**](inline_response_200_16.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSignedInUser**
> UserProfileDtoModel GetSignedInUser(ctx, )
Get signed in User.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**UserProfileDtoModel**](UserProfileDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSignedInUserAuthTokens**
> []UserTokenModel GetSignedInUserAuthTokens(ctx, )
Auth tokens of the actual User.

Return a list of all auth tokens (devices) that the actual user currently have logged in from.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UserTokenModel**](UserToken.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSignedInUserOrgList**
> []UserOrgDtoModel GetSignedInUserOrgList(ctx, )
Organizations of the actual User.

Return a list of all organizations of the current user.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UserOrgDtoModel**](UserOrgDTO.md)

### Authorization

[basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetSignedInUserTeamList**
> []UserOrgDtoModel GetSignedInUserTeamList(ctx, )
Teams that the actual User is member of.

Return a list of all teams that the current user is member of.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UserOrgDtoModel**](UserOrgDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **GetUserQuotas**
> []UserQuotaDtoModel GetUserQuotas(ctx, )
Fetch user quota.

### Required Parameters
This endpoint does not need any parameter.

### Return type

[**[]UserQuotaDtoModel**](UserQuotaDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **RevokeSignedINAuthTokenCmd**
> SuccessResponseBodyModel RevokeSignedINAuthTokenCmd(ctx, body)
Revoke an auth token of the actual User.

Revokes the given auth token (device) for the actual user. User of issued auth token (device) will no longer be logged in and will be required to authenticate again upon next activity.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**RevokeAuthTokenCmdModel**](RevokeAuthTokenCmdModel.md)|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **SetHelpFlag**
> InlineResponse20016Model SetHelpFlag(ctx, flagId)
Set user help flag.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **flagId** | **string**|  | 

### Return type

[**InlineResponse20016Model**](inline_response_200_16.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **StarDashboard**
> SuccessResponseBodyModel StarDashboard(ctx, dashboardId)
Star a dashboard.

Stars the given Dashboard for the actual user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardId** | **string**|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UnstarDashboard**
> SuccessResponseBodyModel UnstarDashboard(ctx, dashboardId)
Unstar a dashboard.

Deletes the starring of the given Dashboard for the actual user.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **dashboardId** | **string**|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UpdateSignedInUser**
> UserProfileDtoModel UpdateSignedInUser(ctx, body)
Update signed in User.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **body** | [**UpdateUserCommandModel**](UpdateUserCommandModel.md)| To change the email, name, login, theme, provide another one. | 

### Return type

[**UserProfileDtoModel**](UserProfileDTO.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **UserSetUsingOrg**
> SuccessResponseBodyModel UserSetUsingOrg(ctx, orgId)
Switch user context for signed in user.

Switch user context to the given organization.

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **orgId** | **int64**|  | 

### Return type

[**SuccessResponseBodyModel**](SuccessResponseBody.md)

### Authorization

[api_key](../README.md#api_key), [basic](../README.md#basic)

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

