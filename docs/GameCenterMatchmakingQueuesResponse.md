# GameCenterMatchmakingQueuesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterMatchmakingQueue**](GameCenterMatchmakingQueue.md) |  | 
**Included** | Pointer to [**[]GameCenterMatchmakingRuleSet**](GameCenterMatchmakingRuleSet.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterMatchmakingQueuesResponse

`func NewGameCenterMatchmakingQueuesResponse(data []GameCenterMatchmakingQueue, links PagedDocumentLinks, ) *GameCenterMatchmakingQueuesResponse`

NewGameCenterMatchmakingQueuesResponse instantiates a new GameCenterMatchmakingQueuesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterMatchmakingQueuesResponseWithDefaults

`func NewGameCenterMatchmakingQueuesResponseWithDefaults() *GameCenterMatchmakingQueuesResponse`

NewGameCenterMatchmakingQueuesResponseWithDefaults instantiates a new GameCenterMatchmakingQueuesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterMatchmakingQueuesResponse) GetData() []GameCenterMatchmakingQueue`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterMatchmakingQueuesResponse) GetDataOk() (*[]GameCenterMatchmakingQueue, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterMatchmakingQueuesResponse) SetData(v []GameCenterMatchmakingQueue)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterMatchmakingQueuesResponse) GetIncluded() []GameCenterMatchmakingRuleSet`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterMatchmakingQueuesResponse) GetIncludedOk() (*[]GameCenterMatchmakingRuleSet, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterMatchmakingQueuesResponse) SetIncluded(v []GameCenterMatchmakingRuleSet)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterMatchmakingQueuesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterMatchmakingQueuesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterMatchmakingQueuesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterMatchmakingQueuesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterMatchmakingQueuesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterMatchmakingQueuesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterMatchmakingQueuesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterMatchmakingQueuesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


