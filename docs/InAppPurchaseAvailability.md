# InAppPurchaseAvailability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppAvailabilityV2Attributes**](AppAvailabilityV2Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**InAppPurchaseAvailabilityRelationships**](InAppPurchaseAvailabilityRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewInAppPurchaseAvailability

`func NewInAppPurchaseAvailability(type_ string, id string, ) *InAppPurchaseAvailability`

NewInAppPurchaseAvailability instantiates a new InAppPurchaseAvailability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInAppPurchaseAvailabilityWithDefaults

`func NewInAppPurchaseAvailabilityWithDefaults() *InAppPurchaseAvailability`

NewInAppPurchaseAvailabilityWithDefaults instantiates a new InAppPurchaseAvailability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *InAppPurchaseAvailability) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *InAppPurchaseAvailability) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *InAppPurchaseAvailability) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *InAppPurchaseAvailability) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *InAppPurchaseAvailability) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *InAppPurchaseAvailability) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *InAppPurchaseAvailability) GetAttributes() AppAvailabilityV2Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *InAppPurchaseAvailability) GetAttributesOk() (*AppAvailabilityV2Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *InAppPurchaseAvailability) SetAttributes(v AppAvailabilityV2Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *InAppPurchaseAvailability) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *InAppPurchaseAvailability) GetRelationships() InAppPurchaseAvailabilityRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *InAppPurchaseAvailability) GetRelationshipsOk() (*InAppPurchaseAvailabilityRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *InAppPurchaseAvailability) SetRelationships(v InAppPurchaseAvailabilityRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *InAppPurchaseAvailability) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *InAppPurchaseAvailability) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *InAppPurchaseAvailability) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *InAppPurchaseAvailability) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *InAppPurchaseAvailability) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


