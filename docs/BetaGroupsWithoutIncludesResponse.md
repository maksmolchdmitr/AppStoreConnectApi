# BetaGroupsWithoutIncludesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]BetaTester**](BetaTester.md) |  | 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewBetaGroupsWithoutIncludesResponse

`func NewBetaGroupsWithoutIncludesResponse(data []BetaTester, links PagedDocumentLinks, ) *BetaGroupsWithoutIncludesResponse`

NewBetaGroupsWithoutIncludesResponse instantiates a new BetaGroupsWithoutIncludesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBetaGroupsWithoutIncludesResponseWithDefaults

`func NewBetaGroupsWithoutIncludesResponseWithDefaults() *BetaGroupsWithoutIncludesResponse`

NewBetaGroupsWithoutIncludesResponseWithDefaults instantiates a new BetaGroupsWithoutIncludesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *BetaGroupsWithoutIncludesResponse) GetData() []BetaTester`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *BetaGroupsWithoutIncludesResponse) GetDataOk() (*[]BetaTester, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *BetaGroupsWithoutIncludesResponse) SetData(v []BetaTester)`

SetData sets Data field to given value.


### GetLinks

`func (o *BetaGroupsWithoutIncludesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *BetaGroupsWithoutIncludesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *BetaGroupsWithoutIncludesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *BetaGroupsWithoutIncludesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *BetaGroupsWithoutIncludesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *BetaGroupsWithoutIncludesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *BetaGroupsWithoutIncludesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


