# AppStoreVersionLocalizationsWithoutIncludesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]AppStoreVersion**](AppStoreVersion.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewAppStoreVersionLocalizationsWithoutIncludesResponse

`func NewAppStoreVersionLocalizationsWithoutIncludesResponse(data []AppStoreVersion, links PagedDocumentLinks, ) *AppStoreVersionLocalizationsWithoutIncludesResponse`

NewAppStoreVersionLocalizationsWithoutIncludesResponse instantiates a new AppStoreVersionLocalizationsWithoutIncludesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppStoreVersionLocalizationsWithoutIncludesResponseWithDefaults

`func NewAppStoreVersionLocalizationsWithoutIncludesResponseWithDefaults() *AppStoreVersionLocalizationsWithoutIncludesResponse`

NewAppStoreVersionLocalizationsWithoutIncludesResponseWithDefaults instantiates a new AppStoreVersionLocalizationsWithoutIncludesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) GetData() []AppStoreVersion`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) GetDataOk() (*[]AppStoreVersion, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) SetData(v []AppStoreVersion)`

SetData sets Data field to given value.


### GetLinks

`func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


