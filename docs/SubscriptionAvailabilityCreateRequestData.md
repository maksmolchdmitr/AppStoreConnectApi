# SubscriptionAvailabilityCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**AppAvailabilityV2CreateRequestDataAttributes**](AppAvailabilityV2CreateRequestDataAttributes.md) |  | 
**Relationships** | [**SubscriptionAvailabilityCreateRequestDataRelationships**](SubscriptionAvailabilityCreateRequestDataRelationships.md) |  | 

## Methods

### NewSubscriptionAvailabilityCreateRequestData

`func NewSubscriptionAvailabilityCreateRequestData(type_ string, attributes AppAvailabilityV2CreateRequestDataAttributes, relationships SubscriptionAvailabilityCreateRequestDataRelationships, ) *SubscriptionAvailabilityCreateRequestData`

NewSubscriptionAvailabilityCreateRequestData instantiates a new SubscriptionAvailabilityCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionAvailabilityCreateRequestDataWithDefaults

`func NewSubscriptionAvailabilityCreateRequestDataWithDefaults() *SubscriptionAvailabilityCreateRequestData`

NewSubscriptionAvailabilityCreateRequestDataWithDefaults instantiates a new SubscriptionAvailabilityCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *SubscriptionAvailabilityCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionAvailabilityCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionAvailabilityCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *SubscriptionAvailabilityCreateRequestData) GetAttributes() AppAvailabilityV2CreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *SubscriptionAvailabilityCreateRequestData) GetAttributesOk() (*AppAvailabilityV2CreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *SubscriptionAvailabilityCreateRequestData) SetAttributes(v AppAvailabilityV2CreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *SubscriptionAvailabilityCreateRequestData) GetRelationships() SubscriptionAvailabilityCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *SubscriptionAvailabilityCreateRequestData) GetRelationshipsOk() (*SubscriptionAvailabilityCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *SubscriptionAvailabilityCreateRequestData) SetRelationships(v SubscriptionAvailabilityCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


