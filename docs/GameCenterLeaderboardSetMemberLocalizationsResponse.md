# GameCenterLeaderboardSetMemberLocalizationsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterLeaderboardSetMemberLocalization**](GameCenterLeaderboardSetMemberLocalization.md) |  | 
**Included** | Pointer to [**[]GameCenterLeaderboardSetMemberLocalizationsResponseIncludedInner**](GameCenterLeaderboardSetMemberLocalizationsResponseIncludedInner.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterLeaderboardSetMemberLocalizationsResponse

`func NewGameCenterLeaderboardSetMemberLocalizationsResponse(data []GameCenterLeaderboardSetMemberLocalization, links PagedDocumentLinks, ) *GameCenterLeaderboardSetMemberLocalizationsResponse`

NewGameCenterLeaderboardSetMemberLocalizationsResponse instantiates a new GameCenterLeaderboardSetMemberLocalizationsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardSetMemberLocalizationsResponseWithDefaults

`func NewGameCenterLeaderboardSetMemberLocalizationsResponseWithDefaults() *GameCenterLeaderboardSetMemberLocalizationsResponse`

NewGameCenterLeaderboardSetMemberLocalizationsResponseWithDefaults instantiates a new GameCenterLeaderboardSetMemberLocalizationsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterLeaderboardSetMemberLocalizationsResponse) GetData() []GameCenterLeaderboardSetMemberLocalization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterLeaderboardSetMemberLocalizationsResponse) GetDataOk() (*[]GameCenterLeaderboardSetMemberLocalization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterLeaderboardSetMemberLocalizationsResponse) SetData(v []GameCenterLeaderboardSetMemberLocalization)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterLeaderboardSetMemberLocalizationsResponse) GetIncluded() []GameCenterLeaderboardSetMemberLocalizationsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterLeaderboardSetMemberLocalizationsResponse) GetIncludedOk() (*[]GameCenterLeaderboardSetMemberLocalizationsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterLeaderboardSetMemberLocalizationsResponse) SetIncluded(v []GameCenterLeaderboardSetMemberLocalizationsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterLeaderboardSetMemberLocalizationsResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardSetMemberLocalizationsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardSetMemberLocalizationsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardSetMemberLocalizationsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterLeaderboardSetMemberLocalizationsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterLeaderboardSetMemberLocalizationsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterLeaderboardSetMemberLocalizationsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterLeaderboardSetMemberLocalizationsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


