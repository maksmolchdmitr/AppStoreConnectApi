# TerritoryAvailabilityResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**TerritoryAvailability**](TerritoryAvailability.md) |  | 
**Included** | Pointer to [**[]Territory**](Territory.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewTerritoryAvailabilityResponse

`func NewTerritoryAvailabilityResponse(data TerritoryAvailability, links DocumentLinks, ) *TerritoryAvailabilityResponse`

NewTerritoryAvailabilityResponse instantiates a new TerritoryAvailabilityResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTerritoryAvailabilityResponseWithDefaults

`func NewTerritoryAvailabilityResponseWithDefaults() *TerritoryAvailabilityResponse`

NewTerritoryAvailabilityResponseWithDefaults instantiates a new TerritoryAvailabilityResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *TerritoryAvailabilityResponse) GetData() TerritoryAvailability`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *TerritoryAvailabilityResponse) GetDataOk() (*TerritoryAvailability, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *TerritoryAvailabilityResponse) SetData(v TerritoryAvailability)`

SetData sets Data field to given value.


### GetIncluded

`func (o *TerritoryAvailabilityResponse) GetIncluded() []Territory`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *TerritoryAvailabilityResponse) GetIncludedOk() (*[]Territory, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *TerritoryAvailabilityResponse) SetIncluded(v []Territory)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *TerritoryAvailabilityResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *TerritoryAvailabilityResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *TerritoryAvailabilityResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *TerritoryAvailabilityResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


