# BetaTestersWithoutIncludesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]Build**](Build.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBetaTestersWithoutIncludesResponse

`func NewBetaTestersWithoutIncludesResponse(data []Build, links PagedDocumentLinks, ) *BetaTestersWithoutIncludesResponse`

NewBetaTestersWithoutIncludesResponse instantiates a new BetaTestersWithoutIncludesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaTestersWithoutIncludesResponseWithDefaults

`func NewBetaTestersWithoutIncludesResponseWithDefaults() *BetaTestersWithoutIncludesResponse`

NewBetaTestersWithoutIncludesResponseWithDefaults instantiates a new BetaTestersWithoutIncludesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaTestersWithoutIncludesResponse) GetData() []Build`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaTestersWithoutIncludesResponse) GetDataOk() (*[]Build, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaTestersWithoutIncludesResponse) SetData(v []Build)`

SetData sets Data field to given value.


### GetLinks

`func (o *BetaTestersWithoutIncludesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaTestersWithoutIncludesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaTestersWithoutIncludesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BetaTestersWithoutIncludesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BetaTestersWithoutIncludesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BetaTestersWithoutIncludesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BetaTestersWithoutIncludesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


