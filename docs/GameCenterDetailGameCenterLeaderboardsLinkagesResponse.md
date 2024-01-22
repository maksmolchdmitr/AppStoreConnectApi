# GameCenterDetailGameCenterLeaderboardsLinkagesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner**](GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterDetailGameCenterLeaderboardsLinkagesResponse

`func NewGameCenterDetailGameCenterLeaderboardsLinkagesResponse(data []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner, links PagedDocumentLinks, ) *GameCenterDetailGameCenterLeaderboardsLinkagesResponse`

NewGameCenterDetailGameCenterLeaderboardsLinkagesResponse instantiates a new GameCenterDetailGameCenterLeaderboardsLinkagesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterDetailGameCenterLeaderboardsLinkagesResponseWithDefaults

`func NewGameCenterDetailGameCenterLeaderboardsLinkagesResponseWithDefaults() *GameCenterDetailGameCenterLeaderboardsLinkagesResponse`

NewGameCenterDetailGameCenterLeaderboardsLinkagesResponseWithDefaults instantiates a new GameCenterDetailGameCenterLeaderboardsLinkagesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterDetailGameCenterLeaderboardsLinkagesResponse) GetData() []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterDetailGameCenterLeaderboardsLinkagesResponse) GetDataOk() (*[]GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterDetailGameCenterLeaderboardsLinkagesResponse) SetData(v []GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GameCenterDetailGameCenterLeaderboardsLinkagesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterDetailGameCenterLeaderboardsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterDetailGameCenterLeaderboardsLinkagesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterDetailGameCenterLeaderboardsLinkagesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterDetailGameCenterLeaderboardsLinkagesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterDetailGameCenterLeaderboardsLinkagesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterDetailGameCenterLeaderboardsLinkagesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


