# GameCenterMatchmakingRuleSetsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterMatchmakingRuleSet**](GameCenterMatchmakingRuleSet.md) |  | 
**Included** | Pointer to [**[]GameCenterMatchmakingRuleSetsResponseIncludedInner**](GameCenterMatchmakingRuleSetsResponseIncludedInner.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterMatchmakingRuleSetsResponse

`func NewGameCenterMatchmakingRuleSetsResponse(data []GameCenterMatchmakingRuleSet, links PagedDocumentLinks, ) *GameCenterMatchmakingRuleSetsResponse`

NewGameCenterMatchmakingRuleSetsResponse instantiates a new GameCenterMatchmakingRuleSetsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterMatchmakingRuleSetsResponseWithDefaults

`func NewGameCenterMatchmakingRuleSetsResponseWithDefaults() *GameCenterMatchmakingRuleSetsResponse`

NewGameCenterMatchmakingRuleSetsResponseWithDefaults instantiates a new GameCenterMatchmakingRuleSetsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterMatchmakingRuleSetsResponse) GetData() []GameCenterMatchmakingRuleSet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterMatchmakingRuleSetsResponse) GetDataOk() (*[]GameCenterMatchmakingRuleSet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterMatchmakingRuleSetsResponse) SetData(v []GameCenterMatchmakingRuleSet)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterMatchmakingRuleSetsResponse) GetIncluded() []GameCenterMatchmakingRuleSetsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterMatchmakingRuleSetsResponse) GetIncludedOk() (*[]GameCenterMatchmakingRuleSetsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterMatchmakingRuleSetsResponse) SetIncluded(v []GameCenterMatchmakingRuleSetsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterMatchmakingRuleSetsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterMatchmakingRuleSetsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterMatchmakingRuleSetsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterMatchmakingRuleSetsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterMatchmakingRuleSetsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterMatchmakingRuleSetsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterMatchmakingRuleSetsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterMatchmakingRuleSetsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


