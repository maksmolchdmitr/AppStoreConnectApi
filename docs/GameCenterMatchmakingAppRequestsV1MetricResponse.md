# GameCenterMatchmakingAppRequestsV1MetricResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterMatchmakingAppRequestsV1MetricResponseDataInner**](GameCenterMatchmakingAppRequestsV1MetricResponseDataInner.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterMatchmakingAppRequestsV1MetricResponse

`func NewGameCenterMatchmakingAppRequestsV1MetricResponse(data []GameCenterMatchmakingAppRequestsV1MetricResponseDataInner, links PagedDocumentLinks, ) *GameCenterMatchmakingAppRequestsV1MetricResponse`

NewGameCenterMatchmakingAppRequestsV1MetricResponse instantiates a new GameCenterMatchmakingAppRequestsV1MetricResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterMatchmakingAppRequestsV1MetricResponseWithDefaults

`func NewGameCenterMatchmakingAppRequestsV1MetricResponseWithDefaults() *GameCenterMatchmakingAppRequestsV1MetricResponse`

NewGameCenterMatchmakingAppRequestsV1MetricResponseWithDefaults instantiates a new GameCenterMatchmakingAppRequestsV1MetricResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterMatchmakingAppRequestsV1MetricResponse) GetData() []GameCenterMatchmakingAppRequestsV1MetricResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterMatchmakingAppRequestsV1MetricResponse) GetDataOk() (*[]GameCenterMatchmakingAppRequestsV1MetricResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterMatchmakingAppRequestsV1MetricResponse) SetData(v []GameCenterMatchmakingAppRequestsV1MetricResponseDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GameCenterMatchmakingAppRequestsV1MetricResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterMatchmakingAppRequestsV1MetricResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterMatchmakingAppRequestsV1MetricResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterMatchmakingAppRequestsV1MetricResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterMatchmakingAppRequestsV1MetricResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterMatchmakingAppRequestsV1MetricResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterMatchmakingAppRequestsV1MetricResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


