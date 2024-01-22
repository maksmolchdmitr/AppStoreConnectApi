# AppAvailabilityV2Response

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppAvailabilityV2**](AppAvailabilityV2.md) |  | 
**Included** | Pointer to [**[]TerritoryAvailability**](TerritoryAvailability.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewAppAvailabilityV2Response

`func NewAppAvailabilityV2Response(data AppAvailabilityV2, links DocumentLinks, ) *AppAvailabilityV2Response`

NewAppAvailabilityV2Response instantiates a new AppAvailabilityV2Response object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAvailabilityV2ResponseWithDefaults

`func NewAppAvailabilityV2ResponseWithDefaults() *AppAvailabilityV2Response`

NewAppAvailabilityV2ResponseWithDefaults instantiates a new AppAvailabilityV2Response object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppAvailabilityV2Response) GetData() AppAvailabilityV2`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppAvailabilityV2Response) GetDataOk() (*AppAvailabilityV2, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppAvailabilityV2Response) SetData(v AppAvailabilityV2)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppAvailabilityV2Response) GetIncluded() []TerritoryAvailability`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppAvailabilityV2Response) GetIncludedOk() (*[]TerritoryAvailability, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppAvailabilityV2Response) SetIncluded(v []TerritoryAvailability)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppAvailabilityV2Response) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *AppAvailabilityV2Response) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppAvailabilityV2Response) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppAvailabilityV2Response) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


