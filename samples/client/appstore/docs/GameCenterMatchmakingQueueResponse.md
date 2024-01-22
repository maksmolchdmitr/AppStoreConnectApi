# GameCenterMatchmakingQueueResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterMatchmakingQueue**](GameCenterMatchmakingQueue.md) |  | 
**Included** | Pointer to [**[]GameCenterMatchmakingRuleSet**](GameCenterMatchmakingRuleSet.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterMatchmakingQueueResponse

`func NewGameCenterMatchmakingQueueResponse(data GameCenterMatchmakingQueue, links DocumentLinks, ) *GameCenterMatchmakingQueueResponse`

NewGameCenterMatchmakingQueueResponse instantiates a new GameCenterMatchmakingQueueResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterMatchmakingQueueResponseWithDefaults

`func NewGameCenterMatchmakingQueueResponseWithDefaults() *GameCenterMatchmakingQueueResponse`

NewGameCenterMatchmakingQueueResponseWithDefaults instantiates a new GameCenterMatchmakingQueueResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterMatchmakingQueueResponse) GetData() GameCenterMatchmakingQueue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterMatchmakingQueueResponse) GetDataOk() (*GameCenterMatchmakingQueue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterMatchmakingQueueResponse) SetData(v GameCenterMatchmakingQueue)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterMatchmakingQueueResponse) GetIncluded() []GameCenterMatchmakingRuleSet`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterMatchmakingQueueResponse) GetIncludedOk() (*[]GameCenterMatchmakingRuleSet, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterMatchmakingQueueResponse) SetIncluded(v []GameCenterMatchmakingRuleSet)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterMatchmakingQueueResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterMatchmakingQueueResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterMatchmakingQueueResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterMatchmakingQueueResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


