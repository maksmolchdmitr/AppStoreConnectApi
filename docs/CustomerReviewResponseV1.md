# CustomerReviewResponseV1

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**CustomerReviewResponseV1Attributes**](CustomerReviewResponseV1Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**CustomerReviewResponseV1Relationships**](CustomerReviewResponseV1Relationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewCustomerReviewResponseV1

`func NewCustomerReviewResponseV1(type_ string, id string, ) *CustomerReviewResponseV1`

NewCustomerReviewResponseV1 instantiates a new CustomerReviewResponseV1 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCustomerReviewResponseV1WithDefaults

`func NewCustomerReviewResponseV1WithDefaults() *CustomerReviewResponseV1`

NewCustomerReviewResponseV1WithDefaults instantiates a new CustomerReviewResponseV1 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *CustomerReviewResponseV1) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CustomerReviewResponseV1) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CustomerReviewResponseV1) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *CustomerReviewResponseV1) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CustomerReviewResponseV1) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CustomerReviewResponseV1) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *CustomerReviewResponseV1) GetAttributes() CustomerReviewResponseV1Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *CustomerReviewResponseV1) GetAttributesOk() (*CustomerReviewResponseV1Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *CustomerReviewResponseV1) SetAttributes(v CustomerReviewResponseV1Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *CustomerReviewResponseV1) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *CustomerReviewResponseV1) GetRelationships() CustomerReviewResponseV1Relationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *CustomerReviewResponseV1) GetRelationshipsOk() (*CustomerReviewResponseV1Relationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *CustomerReviewResponseV1) SetRelationships(v CustomerReviewResponseV1Relationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *CustomerReviewResponseV1) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *CustomerReviewResponseV1) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *CustomerReviewResponseV1) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *CustomerReviewResponseV1) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *CustomerReviewResponseV1) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


