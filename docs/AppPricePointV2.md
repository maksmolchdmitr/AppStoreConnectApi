# AppPricePointV2

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**AppPricePointV2Attributes**](AppPricePointV2Attributes.md) |  | [optional] 
**Relationships** | Pointer to [**AppPricePointV2Relationships**](AppPricePointV2Relationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewAppPricePointV2

`func NewAppPricePointV2(type_ string, id string, ) *AppPricePointV2`

NewAppPricePointV2 instantiates a new AppPricePointV2 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAppPricePointV2WithDefaults

`func NewAppPricePointV2WithDefaults() *AppPricePointV2`

NewAppPricePointV2WithDefaults instantiates a new AppPricePointV2 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *AppPricePointV2) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AppPricePointV2) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AppPricePointV2) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *AppPricePointV2) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AppPricePointV2) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AppPricePointV2) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *AppPricePointV2) GetAttributes() AppPricePointV2Attributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *AppPricePointV2) GetAttributesOk() (*AppPricePointV2Attributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *AppPricePointV2) SetAttributes(v AppPricePointV2Attributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *AppPricePointV2) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *AppPricePointV2) GetRelationships() AppPricePointV2Relationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *AppPricePointV2) GetRelationshipsOk() (*AppPricePointV2Relationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *AppPricePointV2) SetRelationships(v AppPricePointV2Relationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *AppPricePointV2) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *AppPricePointV2) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *AppPricePointV2) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *AppPricePointV2) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *AppPricePointV2) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


