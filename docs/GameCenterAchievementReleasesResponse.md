# GameCenterAchievementReleasesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterAchievementRelease**](GameCenterAchievementRelease.md) |  | 
**Included** | Pointer to [**[]GameCenterAchievementReleasesResponseIncludedInner**](GameCenterAchievementReleasesResponseIncludedInner.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterAchievementReleasesResponse

`func NewGameCenterAchievementReleasesResponse(data []GameCenterAchievementRelease, links PagedDocumentLinks, ) *GameCenterAchievementReleasesResponse`

NewGameCenterAchievementReleasesResponse instantiates a new GameCenterAchievementReleasesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterAchievementReleasesResponseWithDefaults

`func NewGameCenterAchievementReleasesResponseWithDefaults() *GameCenterAchievementReleasesResponse`

NewGameCenterAchievementReleasesResponseWithDefaults instantiates a new GameCenterAchievementReleasesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterAchievementReleasesResponse) GetData() []GameCenterAchievementRelease`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterAchievementReleasesResponse) GetDataOk() (*[]GameCenterAchievementRelease, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterAchievementReleasesResponse) SetData(v []GameCenterAchievementRelease)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterAchievementReleasesResponse) GetIncluded() []GameCenterAchievementReleasesResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterAchievementReleasesResponse) GetIncludedOk() (*[]GameCenterAchievementReleasesResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterAchievementReleasesResponse) SetIncluded(v []GameCenterAchievementReleasesResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterAchievementReleasesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterAchievementReleasesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterAchievementReleasesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterAchievementReleasesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterAchievementReleasesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterAchievementReleasesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterAchievementReleasesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterAchievementReleasesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


