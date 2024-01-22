# GameCenterMatchmakingQueueRequestsV1MetricResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner**](GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewGameCenterMatchmakingQueueRequestsV1MetricResponse

`func NewGameCenterMatchmakingQueueRequestsV1MetricResponse(data []GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner, links PagedDocumentLinks, ) *GameCenterMatchmakingQueueRequestsV1MetricResponse`

NewGameCenterMatchmakingQueueRequestsV1MetricResponse instantiates a new GameCenterMatchmakingQueueRequestsV1MetricResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterMatchmakingQueueRequestsV1MetricResponseWithDefaults

`func NewGameCenterMatchmakingQueueRequestsV1MetricResponseWithDefaults() *GameCenterMatchmakingQueueRequestsV1MetricResponse`

NewGameCenterMatchmakingQueueRequestsV1MetricResponseWithDefaults instantiates a new GameCenterMatchmakingQueueRequestsV1MetricResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) GetData() []GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) GetDataOk() (*[]GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) SetData(v []GameCenterMatchmakingQueueRequestsV1MetricResponseDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *GameCenterMatchmakingQueueRequestsV1MetricResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


