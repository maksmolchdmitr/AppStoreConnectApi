# TerritoryAvailabilitiesResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**[]TerritoryAvailability**](TerritoryAvailability.md) |  | 
**Included** | Pointer to [**[]Territory**](Territory.md) |  | [optional] 
**Links** | [**PagedDocumentLinks**](PagedDocumentLinks.md) |  | 
**Meta** | Pointer to [**PagingInformation**](PagingInformation.md) |  | [optional] 

## Methods

### NewTerritoryAvailabilitiesResponse

`func NewTerritoryAvailabilitiesResponse(data []TerritoryAvailability, links PagedDocumentLinks, ) *TerritoryAvailabilitiesResponse`

NewTerritoryAvailabilitiesResponse instantiates a new TerritoryAvailabilitiesResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerritoryAvailabilitiesResponseWithDefaults

`func NewTerritoryAvailabilitiesResponseWithDefaults() *TerritoryAvailabilitiesResponse`

NewTerritoryAvailabilitiesResponseWithDefaults instantiates a new TerritoryAvailabilitiesResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TerritoryAvailabilitiesResponse) GetData() []TerritoryAvailability`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TerritoryAvailabilitiesResponse) GetDataOk() (*[]TerritoryAvailability, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TerritoryAvailabilitiesResponse) SetData(v []TerritoryAvailability)`

SetData sets Data field to given value.


### GetIncluded

`func (o *TerritoryAvailabilitiesResponse) GetIncluded() []Territory`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *TerritoryAvailabilitiesResponse) GetIncludedOk() (*[]Territory, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *TerritoryAvailabilitiesResponse) SetIncluded(v []Territory)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *TerritoryAvailabilitiesResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *TerritoryAvailabilitiesResponse) GetLinks() PagedDocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TerritoryAvailabilitiesResponse) GetLinksOk() (*PagedDocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TerritoryAvailabilitiesResponse) SetLinks(v PagedDocumentLinks)`

SetLinks sets Links field to given value.


### GetMeta

`func (o *TerritoryAvailabilitiesResponse) GetMeta() PagingInformation`

GetMeta returns the Meta field if non-nil, zero value otherwise.

### GetMetaOk

`func (o *TerritoryAvailabilitiesResponse) GetMetaOk() (*PagingInformation, bool)`

GetMetaOk returns a tuple with the Meta field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMeta

`func (o *TerritoryAvailabilitiesResponse) SetMeta(v PagingInformation)`

SetMeta sets Meta field to given value.

### HasMeta

`func (o *TerritoryAvailabilitiesResponse) HasMeta() bool`

HasMeta returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


