# GameCenterAchievementLocalizationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterAchievementLocalization**](GameCenterAchievementLocalization.md) |  | 
**Included** | Pointer to [**[]GameCenterAchievementLocalizationsResponseIncludedInner**](GameCenterAchievementLocalizationsResponseIncludedInner.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterAchievementLocalizationsResponse

`func NewGameCenterAchievementLocalizationsResponse(data []GameCenterAchievementLocalization, links PagedDocumentLinks, ) *GameCenterAchievementLocalizationsResponse`

NewGameCenterAchievementLocalizationsResponse instantiates a new GameCenterAchievementLocalizationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterAchievementLocalizationsResponseWithDefaults

`func NewGameCenterAchievementLocalizationsResponseWithDefaults() *GameCenterAchievementLocalizationsResponse`

NewGameCenterAchievementLocalizationsResponseWithDefaults instantiates a new GameCenterAchievementLocalizationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterAchievementLocalizationsResponse) GetData() []GameCenterAchievementLocalization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterAchievementLocalizationsResponse) GetDataOk() (*[]GameCenterAchievementLocalization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterAchievementLocalizationsResponse) SetData(v []GameCenterAchievementLocalization)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterAchievementLocalizationsResponse) GetIncluded() []GameCenterAchievementLocalizationsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterAchievementLocalizationsResponse) GetIncludedOk() (*[]GameCenterAchievementLocalizationsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterAchievementLocalizationsResponse) SetIncluded(v []GameCenterAchievementLocalizationsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterAchievementLocalizationsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterAchievementLocalizationsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterAchievementLocalizationsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterAchievementLocalizationsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterAchievementLocalizationsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterAchievementLocalizationsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterAchievementLocalizationsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterAchievementLocalizationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


