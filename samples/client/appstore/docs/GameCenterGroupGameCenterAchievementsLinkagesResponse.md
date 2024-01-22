# GameCenterGroupGameCenterAchievementsLinkagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData**](GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterGroupGameCenterAchievementsLinkagesResponse

`func NewGameCenterGroupGameCenterAchievementsLinkagesResponse(data []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData, links PagedDocumentLinks, ) *GameCenterGroupGameCenterAchievementsLinkagesResponse`

NewGameCenterGroupGameCenterAchievementsLinkagesResponse instantiates a new GameCenterGroupGameCenterAchievementsLinkagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterGroupGameCenterAchievementsLinkagesResponseWithDefaults

`func NewGameCenterGroupGameCenterAchievementsLinkagesResponseWithDefaults() *GameCenterGroupGameCenterAchievementsLinkagesResponse`

NewGameCenterGroupGameCenterAchievementsLinkagesResponseWithDefaults instantiates a new GameCenterGroupGameCenterAchievementsLinkagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) GetData() []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) GetDataOk() (*[]GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) SetData(v []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData)`

SetData sets Data field to given value.


### GetLinks

`func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterGroupGameCenterAchievementsLinkagesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


