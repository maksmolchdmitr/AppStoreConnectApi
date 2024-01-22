# GameCenterMatchmakingSessionsV1MetricResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterMatchmakingSessionsV1MetricResponseDataInner**](GameCenterMatchmakingSessionsV1MetricResponseDataInner.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterMatchmakingSessionsV1MetricResponse

`func NewGameCenterMatchmakingSessionsV1MetricResponse(data []GameCenterMatchmakingSessionsV1MetricResponseDataInner, links PagedDocumentLinks, ) *GameCenterMatchmakingSessionsV1MetricResponse`

NewGameCenterMatchmakingSessionsV1MetricResponse instantiates a new GameCenterMatchmakingSessionsV1MetricResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterMatchmakingSessionsV1MetricResponseWithDefaults

`func NewGameCenterMatchmakingSessionsV1MetricResponseWithDefaults() *GameCenterMatchmakingSessionsV1MetricResponse`

NewGameCenterMatchmakingSessionsV1MetricResponseWithDefaults instantiates a new GameCenterMatchmakingSessionsV1MetricResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterMatchmakingSessionsV1MetricResponse) GetData() []GameCenterMatchmakingSessionsV1MetricResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterMatchmakingSessionsV1MetricResponse) GetDataOk() (*[]GameCenterMatchmakingSessionsV1MetricResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterMatchmakingSessionsV1MetricResponse) SetData(v []GameCenterMatchmakingSessionsV1MetricResponseDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GameCenterMatchmakingSessionsV1MetricResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterMatchmakingSessionsV1MetricResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterMatchmakingSessionsV1MetricResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterMatchmakingSessionsV1MetricResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterMatchmakingSessionsV1MetricResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterMatchmakingSessionsV1MetricResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterMatchmakingSessionsV1MetricResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


