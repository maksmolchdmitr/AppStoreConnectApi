# ScmRepository

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**ScmRepositoryAttributes**](ScmRepositoryAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**ScmRepositoryRelationships**](ScmRepositoryRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewScmRepository

`func NewScmRepository(type_ string, id string, ) *ScmRepository`

NewScmRepository instantiates a new ScmRepository object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScmRepositoryWithDefaults

`func NewScmRepositoryWithDefaults() *ScmRepository`

NewScmRepositoryWithDefaults instantiates a new ScmRepository object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *ScmRepository) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ScmRepository) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ScmRepository) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *ScmRepository) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ScmRepository) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ScmRepository) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *ScmRepository) GetAttributes() ScmRepositoryAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *ScmRepository) GetAttributesOk() (*ScmRepositoryAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *ScmRepository) SetAttributes(v ScmRepositoryAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *ScmRepository) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *ScmRepository) GetRelationships() ScmRepositoryRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *ScmRepository) GetRelationshipsOk() (*ScmRepositoryRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *ScmRepository) SetRelationships(v ScmRepositoryRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *ScmRepository) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *ScmRepository) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *ScmRepository) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *ScmRepository) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *ScmRepository) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


