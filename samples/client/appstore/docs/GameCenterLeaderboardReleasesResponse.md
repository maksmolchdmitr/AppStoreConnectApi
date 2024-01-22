# GameCenterLeaderboardReleasesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterLeaderboardRelease**](GameCenterLeaderboardRelease.md) |  | 
**Included** | Pointer to [**[]GameCenterLeaderboardReleasesResponseIncludedInner**](GameCenterLeaderboardReleasesResponseIncludedInner.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterLeaderboardReleasesResponse

`func NewGameCenterLeaderboardReleasesResponse(data []GameCenterLeaderboardRelease, links PagedDocumentLinks, ) *GameCenterLeaderboardReleasesResponse`

NewGameCenterLeaderboardReleasesResponse instantiates a new GameCenterLeaderboardReleasesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardReleasesResponseWithDefaults

`func NewGameCenterLeaderboardReleasesResponseWithDefaults() *GameCenterLeaderboardReleasesResponse`

NewGameCenterLeaderboardReleasesResponseWithDefaults instantiates a new GameCenterLeaderboardReleasesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterLeaderboardReleasesResponse) GetData() []GameCenterLeaderboardRelease`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterLeaderboardReleasesResponse) GetDataOk() (*[]GameCenterLeaderboardRelease, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterLeaderboardReleasesResponse) SetData(v []GameCenterLeaderboardRelease)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterLeaderboardReleasesResponse) GetIncluded() []GameCenterLeaderboardReleasesResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterLeaderboardReleasesResponse) GetIncludedOk() (*[]GameCenterLeaderboardReleasesResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterLeaderboardReleasesResponse) SetIncluded(v []GameCenterLeaderboardReleasesResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterLeaderboardReleasesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardReleasesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardReleasesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardReleasesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterLeaderboardReleasesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterLeaderboardReleasesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterLeaderboardReleasesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterLeaderboardReleasesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


