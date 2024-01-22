# BuildIconsWithoutIncludesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Build**](Build.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBuildIconsWithoutIncludesResponse

`func NewBuildIconsWithoutIncludesResponse(data []Build, links PagedDocumentLinks, ) *BuildIconsWithoutIncludesResponse`

NewBuildIconsWithoutIncludesResponse instantiates a new BuildIconsWithoutIncludesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBuildIconsWithoutIncludesResponseWithDefaults

`func NewBuildIconsWithoutIncludesResponseWithDefaults() *BuildIconsWithoutIncludesResponse`

NewBuildIconsWithoutIncludesResponseWithDefaults instantiates a new BuildIconsWithoutIncludesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BuildIconsWithoutIncludesResponse) GetData() []Build`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BuildIconsWithoutIncludesResponse) GetDataOk() (*[]Build, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BuildIconsWithoutIncludesResponse) SetData(v []Build)`

SetData sets Data field to given value.


### GetLinks

`func (o *BuildIconsWithoutIncludesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BuildIconsWithoutIncludesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BuildIconsWithoutIncludesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BuildIconsWithoutIncludesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BuildIconsWithoutIncludesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BuildIconsWithoutIncludesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BuildIconsWithoutIncludesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


