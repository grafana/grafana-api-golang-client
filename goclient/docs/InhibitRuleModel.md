# InhibitRuleModel

## Properties
Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Equal** | [***LabelNamesModel**](LabelNames.md) |  | [optional] [default to null]
**SourceMatch** | **map[string]string** | SourceMatch defines a set of labels that have to equal the given value for source alerts. Deprecated. Remove before v1.0 release. | [optional] [default to null]
**SourceMatchRe** | [***MatchRegexpsModel**](MatchRegexps.md) |  | [optional] [default to null]
**SourceMatchers** | [***MatchersModel**](Matchers.md) |  | [optional] [default to null]
**TargetMatch** | **map[string]string** | TargetMatch defines a set of labels that have to equal the given value for target alerts. Deprecated. Remove before v1.0 release. | [optional] [default to null]
**TargetMatchRe** | [***MatchRegexpsModel**](MatchRegexps.md) |  | [optional] [default to null]
**TargetMatchers** | [***MatchersModel**](Matchers.md) |  | [optional] [default to null]

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


