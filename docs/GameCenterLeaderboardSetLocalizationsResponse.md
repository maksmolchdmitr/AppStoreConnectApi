# GameCenterLeaderboardSetLocalizationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterLeaderboardSetLocalization**](GameCenterLeaderboardSetLocalization.md) |  | 
**Included** | Pointer to [**[]GameCenterLeaderboardSetLocalizationsResponseIncludedInner**](GameCenterLeaderboardSetLocalizationsResponseIncludedInner.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterLeaderboardSetLocalizationsResponse

`func NewGameCenterLeaderboardSetLocalizationsResponse(data []GameCenterLeaderboardSetLocalization, links PagedDocumentLinks, ) *GameCenterLeaderboardSetLocalizationsResponse`

NewGameCenterLeaderboardSetLocalizationsResponse instantiates a new GameCenterLeaderboardSetLocalizationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardSetLocalizationsResponseWithDefaults

`func NewGameCenterLeaderboardSetLocalizationsResponseWithDefaults() *GameCenterLeaderboardSetLocalizationsResponse`

NewGameCenterLeaderboardSetLocalizationsResponseWithDefaults instantiates a new GameCenterLeaderboardSetLocalizationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterLeaderboardSetLocalizationsResponse) GetData() []GameCenterLeaderboardSetLocalization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterLeaderboardSetLocalizationsResponse) GetDataOk() (*[]GameCenterLeaderboardSetLocalization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterLeaderboardSetLocalizationsResponse) SetData(v []GameCenterLeaderboardSetLocalization)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterLeaderboardSetLocalizationsResponse) GetIncluded() []GameCenterLeaderboardSetLocalizationsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterLeaderboardSetLocalizationsResponse) GetIncludedOk() (*[]GameCenterLeaderboardSetLocalizationsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterLeaderboardSetLocalizationsResponse) SetIncluded(v []GameCenterLeaderboardSetLocalizationsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterLeaderboardSetLocalizationsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardSetLocalizationsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardSetLocalizationsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardSetLocalizationsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterLeaderboardSetLocalizationsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterLeaderboardSetLocalizationsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterLeaderboardSetLocalizationsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterLeaderboardSetLocalizationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


