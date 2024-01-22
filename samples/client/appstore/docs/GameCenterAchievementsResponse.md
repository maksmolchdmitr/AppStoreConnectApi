# GameCenterAchievementsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterAchievement**](GameCenterAchievement.md) |  | 
**Included** | Pointer to [**[]GameCenterAchievementsResponseIncludedInner**](GameCenterAchievementsResponseIncludedInner.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterAchievementsResponse

`func NewGameCenterAchievementsResponse(data []GameCenterAchievement, links PagedDocumentLinks, ) *GameCenterAchievementsResponse`

NewGameCenterAchievementsResponse instantiates a new GameCenterAchievementsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterAchievementsResponseWithDefaults

`func NewGameCenterAchievementsResponseWithDefaults() *GameCenterAchievementsResponse`

NewGameCenterAchievementsResponseWithDefaults instantiates a new GameCenterAchievementsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterAchievementsResponse) GetData() []GameCenterAchievement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterAchievementsResponse) GetDataOk() (*[]GameCenterAchievement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterAchievementsResponse) SetData(v []GameCenterAchievement)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterAchievementsResponse) GetIncluded() []GameCenterAchievementsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterAchievementsResponse) GetIncludedOk() (*[]GameCenterAchievementsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterAchievementsResponse) SetIncluded(v []GameCenterAchievementsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterAchievementsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterAchievementsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterAchievementsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterAchievementsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterAchievementsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterAchievementsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterAchievementsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterAchievementsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


