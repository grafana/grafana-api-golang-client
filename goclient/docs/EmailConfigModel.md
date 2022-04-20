# EmailConfigModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthIdentity** | **string** |  | [optional] [default to null]
**AuthPassword** | [***SecretModel**](Secret.md) |  | [optional] [default to null]
**AuthSecret** | [***SecretModel**](Secret.md) |  | [optional] [default to null]
**AuthUsername** | **string** |  | [optional] [default to null]
**From** | **string** |  | [optional] [default to null]
**Headers** | **map[string]string** |  | [optional] [default to null]
**Hello** | **string** |  | [optional] [default to null]
**Html** | **string** |  | [optional] [default to null]
**RequireTls** | **bool** |  | [optional] [default to null]
**SendResolved** | **bool** |  | [optional] [default to null]
**Smarthost** | [***HostPortModel**](HostPort.md) |  | [optional] [default to null]
**Text** | **string** |  | [optional] [default to null]
**TlsConfig** | [***TlsConfigModel**](TLSConfig.md) |  | [optional] [default to null]
**To** | **string** | Email address to notify. | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


