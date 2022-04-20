# HttpClientConfigModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Authorization** | [***AuthorizationModel**](Authorization.md) |  | [optional] [default to null]
**BasicAuth** | [***BasicAuthModel**](BasicAuth.md) |  | [optional] [default to null]
**BearerToken** | [***SecretModel**](Secret.md) |  | [optional] [default to null]
**BearerTokenFile** | **string** | The bearer token file for the targets. Deprecated in favour of Authorization.CredentialsFile. | [optional] [default to null]
**FollowRedirects** | **bool** | FollowRedirects specifies whether the client should follow HTTP 3xx redirects. The omitempty flag is not set, because it would be hidden from the marshalled configuration when set to false. | [optional] [default to null]
**Oauth2** | [***OAuth2Model**](OAuth2.md) |  | [optional] [default to null]
**ProxyUrl** | [***UrlModel**](URL.md) |  | [optional] [default to null]
**TlsConfig** | [***TlsConfigModel**](TLSConfig.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


