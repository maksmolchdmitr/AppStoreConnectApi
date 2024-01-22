# AppAvailabilityV2CreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AppAvailabilityV2CreateRequestData**](AppAvailabilityV2CreateRequestData.md) |  | 
**Included** | Pointer to [**[]TerritoryAvailabilityInlineCreate**](TerritoryAvailabilityInlineCreate.md) |  | [optional] 

## Methods

### NewAppAvailabilityV2CreateRequest

`func NewAppAvailabilityV2CreateRequest(data AppAvailabilityV2CreateRequestData, ) *AppAvailabilityV2CreateRequest`

NewAppAvailabilityV2CreateRequest instantiates a new AppAvailabilityV2CreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAvailabilityV2CreateRequestWithDefaults

`func NewAppAvailabilityV2CreateRequestWithDefaults() *AppAvailabilityV2CreateRequest`

NewAppAvailabilityV2CreateRequestWithDefaults instantiates a new AppAvailabilityV2CreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AppAvailabilityV2CreateRequest) GetData() AppAvailabilityV2CreateRequestData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AppAvailabilityV2CreateRequest) GetDataOk() (*AppAvailabilityV2CreateRequestData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AppAvailabilityV2CreateRequest) SetData(v AppAvailabilityV2CreateRequestData)`

SetData sets Data field to given value.


### GetIncluded

`func (o *AppAvailabilityV2CreateRequest) GetIncluded() []TerritoryAvailabilityInlineCreate`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *AppAvailabilityV2CreateRequest) GetIncludedOk() (*[]TerritoryAvailabilityInlineCreate, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *AppAvailabilityV2CreateRequest) SetIncluded(v []TerritoryAvailabilityInlineCreate)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *AppAvailabilityV2CreateRequest) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


