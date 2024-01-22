# AppsBetaTesterUsagesV1MetricResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppsBetaTesterUsagesV1MetricResponseDataInner**](AppsBetaTesterUsagesV1MetricResponseDataInner.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 
**Included** | Pointer to [**[]BetaTester**](BetaTester.md) |  | [optional] 

## Methods

### NewAppsBetaTesterUsagesV1MetricResponse

`func NewAppsBetaTesterUsagesV1MetricResponse(data []AppsBetaTesterUsagesV1MetricResponseDataInner, links PagedDocumentLinks, ) *AppsBetaTesterUsagesV1MetricResponse`

NewAppsBetaTesterUsagesV1MetricResponse instantiates a new AppsBetaTesterUsagesV1MetricResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppsBetaTesterUsagesV1MetricResponseWithDefaults

`func NewAppsBetaTesterUsagesV1MetricResponseWithDefaults() *AppsBetaTesterUsagesV1MetricResponse`

NewAppsBetaTesterUsagesV1MetricResponseWithDefaults instantiates a new AppsBetaTesterUsagesV1MetricResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppsBetaTesterUsagesV1MetricResponse) GetData() []AppsBetaTesterUsagesV1MetricResponseDataInner`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppsBetaTesterUsagesV1MetricResponse) GetDataOk() (*[]AppsBetaTesterUsagesV1MetricResponseDataInner, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppsBetaTesterUsagesV1MetricResponse) SetData(v []AppsBetaTesterUsagesV1MetricResponseDataInner)`

SetData sets Data field to given value.


### GetLinks

`func (o *AppsBetaTesterUsagesV1MetricResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppsBetaTesterUsagesV1MetricResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppsBetaTesterUsagesV1MetricResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AppsBetaTesterUsagesV1MetricResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppsBetaTesterUsagesV1MetricResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppsBetaTesterUsagesV1MetricResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppsBetaTesterUsagesV1MetricResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.

### GetIncluded

`func (o *AppsBetaTesterUsagesV1MetricResponse) GetIncluded() []BetaTester`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppsBetaTesterUsagesV1MetricResponse) GetIncludedOk() (*[]BetaTester, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppsBetaTesterUsagesV1MetricResponse) SetIncluded(v []BetaTester)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppsBetaTesterUsagesV1MetricResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


