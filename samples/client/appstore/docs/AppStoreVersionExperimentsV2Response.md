# AppStoreVersionExperimentsV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppStoreVersionExperimentV2**](AppStoreVersionExperimentV2.md) |  | 
**Included** | Pointer to [**[]AppStoreVersionExperimentsV2ResponseIncludedInner**](AppStoreVersionExperimentsV2ResponseIncludedInner.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAppStoreVersionExperimentsV2Response

`func NewAppStoreVersionExperimentsV2Response(data []AppStoreVersionExperimentV2, links PagedDocumentLinks, ) *AppStoreVersionExperimentsV2Response`

NewAppStoreVersionExperimentsV2Response instantiates a new AppStoreVersionExperimentsV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionExperimentsV2ResponseWithDefaults

`func NewAppStoreVersionExperimentsV2ResponseWithDefaults() *AppStoreVersionExperimentsV2Response`

NewAppStoreVersionExperimentsV2ResponseWithDefaults instantiates a new AppStoreVersionExperimentsV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppStoreVersionExperimentsV2Response) GetData() []AppStoreVersionExperimentV2`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppStoreVersionExperimentsV2Response) GetDataOk() (*[]AppStoreVersionExperimentV2, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppStoreVersionExperimentsV2Response) SetData(v []AppStoreVersionExperimentV2)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppStoreVersionExperimentsV2Response) GetIncluded() []AppStoreVersionExperimentsV2ResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppStoreVersionExperimentsV2Response) GetIncludedOk() (*[]AppStoreVersionExperimentsV2ResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppStoreVersionExperimentsV2Response) SetIncluded(v []AppStoreVersionExperimentsV2ResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppStoreVersionExperimentsV2Response) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppStoreVersionExperimentsV2Response) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppStoreVersionExperimentsV2Response) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppStoreVersionExperimentsV2Response) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AppStoreVersionExperimentsV2Response) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppStoreVersionExperimentsV2Response) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppStoreVersionExperimentsV2Response) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppStoreVersionExperimentsV2Response) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


