# GameCenterMatchmakingTeamsResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterMatchmakingTeam**](GameCenterMatchmakingTeam.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterMatchmakingTeamsResponse

`func NewGameCenterMatchmakingTeamsResponse(data []GameCenterMatchmakingTeam, links PagedDocumentLinks, ) *GameCenterMatchmakingTeamsResponse`

NewGameCenterMatchmakingTeamsResponse instantiates a new GameCenterMatchmakingTeamsResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterMatchmakingTeamsResponseWithDefaults

`func NewGameCenterMatchmakingTeamsResponseWithDefaults() *GameCenterMatchmakingTeamsResponse`

NewGameCenterMatchmakingTeamsResponseWithDefaults instantiates a new GameCenterMatchmakingTeamsResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterMatchmakingTeamsResponse) GetData() []GameCenterMatchmakingTeam`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterMatchmakingTeamsResponse) GetDataOk() (*[]GameCenterMatchmakingTeam, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterMatchmakingTeamsResponse) SetData(v []GameCenterMatchmakingTeam)`

SetData sets Data field to given value.


### GetLinks

`func (o *GameCenterMatchmakingTeamsResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterMatchmakingTeamsResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterMatchmakingTeamsResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterMatchmakingTeamsResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterMatchmakingTeamsResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterMatchmakingTeamsResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterMatchmakingTeamsResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


