# BetaTesterUsagesV1MetricResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BetaTesterUsagesV1MetricResponseDataInner**](BetaTesterUsagesV1MetricResponseDataInner.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBetaTesterUsagesV1MetricResponse

`func NewBetaTesterUsagesV1MetricResponse(data []BetaTesterUsagesV1MetricResponseDataInner, links PagedDocumentLinks, ) *BetaTesterUsagesV1MetricResponse`

NewBetaTesterUsagesV1MetricResponse instantiates a new BetaTesterUsagesV1MetricResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaTesterUsagesV1MetricResponseWithDefaults

`func NewBetaTesterUsagesV1MetricResponseWithDefaults() *BetaTesterUsagesV1MetricResponse`

NewBetaTesterUsagesV1MetricResponseWithDefaults instantiates a new BetaTesterUsagesV1MetricResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaTesterUsagesV1MetricResponse) GetData() []BetaTesterUsagesV1MetricResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaTesterUsagesV1MetricResponse) GetDataOk() (*[]BetaTesterUsagesV1MetricResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaTesterUsagesV1MetricResponse) SetData(v []BetaTesterUsagesV1MetricResponseDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *BetaTesterUsagesV1MetricResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaTesterUsagesV1MetricResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaTesterUsagesV1MetricResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BetaTesterUsagesV1MetricResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BetaTesterUsagesV1MetricResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BetaTesterUsagesV1MetricResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BetaTesterUsagesV1MetricResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


