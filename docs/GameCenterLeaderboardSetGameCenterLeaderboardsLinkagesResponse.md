# GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner**](GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse

`func NewGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse(data []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner, links PagedDocumentLinks, ) *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse`

NewGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse instantiates a new GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponseWithDefaults

`func NewGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponseWithDefaults() *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse`

NewGameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponseWithDefaults instantiates a new GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) GetData() []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) GetDataOk() (*[]GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) SetData(v []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterLeaderboardSetGameCenterLeaderboardsLinkagesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


