# GameCenterMatchmakingRuleSetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterMatchmakingRuleSet**](GameCenterMatchmakingRuleSet.md) |  | 
**Included** | Pointer to [**[]GameCenterMatchmakingRuleSetsResponseIncludedInner**](GameCenterMatchmakingRuleSetsResponseIncludedInner.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterMatchmakingRuleSetResponse

`func NewGameCenterMatchmakingRuleSetResponse(data GameCenterMatchmakingRuleSet, links DocumentLinks, ) *GameCenterMatchmakingRuleSetResponse`

NewGameCenterMatchmakingRuleSetResponse instantiates a new GameCenterMatchmakingRuleSetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterMatchmakingRuleSetResponseWithDefaults

`func NewGameCenterMatchmakingRuleSetResponseWithDefaults() *GameCenterMatchmakingRuleSetResponse`

NewGameCenterMatchmakingRuleSetResponseWithDefaults instantiates a new GameCenterMatchmakingRuleSetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterMatchmakingRuleSetResponse) GetData() GameCenterMatchmakingRuleSet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterMatchmakingRuleSetResponse) GetDataOk() (*GameCenterMatchmakingRuleSet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterMatchmakingRuleSetResponse) SetData(v GameCenterMatchmakingRuleSet)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterMatchmakingRuleSetResponse) GetIncluded() []GameCenterMatchmakingRuleSetsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterMatchmakingRuleSetResponse) GetIncludedOk() (*[]GameCenterMatchmakingRuleSetsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterMatchmakingRuleSetResponse) SetIncluded(v []GameCenterMatchmakingRuleSetsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterMatchmakingRuleSetResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterMatchmakingRuleSetResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterMatchmakingRuleSetResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterMatchmakingRuleSetResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


