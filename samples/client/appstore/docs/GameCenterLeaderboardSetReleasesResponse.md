# GameCenterLeaderboardSetReleasesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterLeaderboardSetRelease**](GameCenterLeaderboardSetRelease.md) |  | 
**Included** | Pointer to [**[]GameCenterLeaderboardSetReleasesResponseIncludedInner**](GameCenterLeaderboardSetReleasesResponseIncludedInner.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterLeaderboardSetReleasesResponse

`func NewGameCenterLeaderboardSetReleasesResponse(data []GameCenterLeaderboardSetRelease, links PagedDocumentLinks, ) *GameCenterLeaderboardSetReleasesResponse`

NewGameCenterLeaderboardSetReleasesResponse instantiates a new GameCenterLeaderboardSetReleasesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardSetReleasesResponseWithDefaults

`func NewGameCenterLeaderboardSetReleasesResponseWithDefaults() *GameCenterLeaderboardSetReleasesResponse`

NewGameCenterLeaderboardSetReleasesResponseWithDefaults instantiates a new GameCenterLeaderboardSetReleasesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterLeaderboardSetReleasesResponse) GetData() []GameCenterLeaderboardSetRelease`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterLeaderboardSetReleasesResponse) GetDataOk() (*[]GameCenterLeaderboardSetRelease, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterLeaderboardSetReleasesResponse) SetData(v []GameCenterLeaderboardSetRelease)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterLeaderboardSetReleasesResponse) GetIncluded() []GameCenterLeaderboardSetReleasesResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterLeaderboardSetReleasesResponse) GetIncludedOk() (*[]GameCenterLeaderboardSetReleasesResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterLeaderboardSetReleasesResponse) SetIncluded(v []GameCenterLeaderboardSetReleasesResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterLeaderboardSetReleasesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardSetReleasesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardSetReleasesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardSetReleasesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterLeaderboardSetReleasesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterLeaderboardSetReleasesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterLeaderboardSetReleasesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterLeaderboardSetReleasesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


