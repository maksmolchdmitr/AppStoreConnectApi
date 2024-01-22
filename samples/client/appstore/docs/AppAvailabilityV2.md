# AppAvailabilityV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppAvailabilityV2Attributes**](AppAvailabilityV2Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppAvailabilityV2Relationships**](AppAvailabilityV2Relationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewAppAvailabilityV2

`func NewAppAvailabilityV2(type_ string, id string, ) *AppAvailabilityV2`

NewAppAvailabilityV2 instantiates a new AppAvailabilityV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppAvailabilityV2WithDefaults

`func NewAppAvailabilityV2WithDefaults() *AppAvailabilityV2`

NewAppAvailabilityV2WithDefaults instantiates a new AppAvailabilityV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppAvailabilityV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppAvailabilityV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppAvailabilityV2) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AppAvailabilityV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppAvailabilityV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppAvailabilityV2) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AppAvailabilityV2) GetAttributes() AppAvailabilityV2Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppAvailabilityV2) GetAttributesOk() (*AppAvailabilityV2Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppAvailabilityV2) SetAttributes(v AppAvailabilityV2Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppAvailabilityV2) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AppAvailabilityV2) GetRelationships() AppAvailabilityV2Relationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppAvailabilityV2) GetRelationshipsOk() (*AppAvailabilityV2Relationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppAvailabilityV2) SetRelationships(v AppAvailabilityV2Relationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AppAvailabilityV2) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *AppAvailabilityV2) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppAvailabilityV2) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppAvailabilityV2) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppAvailabilityV2) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


