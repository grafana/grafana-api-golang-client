# RouteModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Continue_** | **bool** |  | [optional] [default to null]
**GroupBy** | **[]string** |  | [optional] [default to null]
**GroupInterval** | [***DurationModel**](Duration.md) |  | [optional] [default to null]
**GroupWait** | [***DurationModel**](Duration.md) |  | [optional] [default to null]
**Match** | **map[string]string** | Deprecated. Remove before v1.0 release. | [optional] [default to null]
**MatchRe** | [***MatchRegexpsModel**](MatchRegexps.md) |  | [optional] [default to null]
**Matchers** | [***MatchersModel**](Matchers.md) |  | [optional] [default to null]
**MuteTimeIntervals** | **[]string** |  | [optional] [default to null]
**ObjectMatchers** | [***ObjectMatchersModel**](ObjectMatchers.md) |  | [optional] [default to null]
**Receiver** | **string** |  | [optional] [default to null]
**RepeatInterval** | [***DurationModel**](Duration.md) |  | [optional] [default to null]
**Routes** | [**[]RouteModel**](Route.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


