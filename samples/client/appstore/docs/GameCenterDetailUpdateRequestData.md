# GameCenterDetailUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**GameCenterDetailCreateRequestDataAttributes**](GameCenterDetailCreateRequestDataAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GameCenterDetailUpdateRequestDataRelationships**](GameCenterDetailUpdateRequestDataRelationships.md) |  | [optional] 

## Methods

### NewGameCenterDetailUpdateRequestData

`func NewGameCenterDetailUpdateRequestData(type_ string, id string, ) *GameCenterDetailUpdateRequestData`

NewGameCenterDetailUpdateRequestData instantiates a new GameCenterDetailUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterDetailUpdateRequestDataWithDefaults

`func NewGameCenterDetailUpdateRequestDataWithDefaults() *GameCenterDetailUpdateRequestData`

NewGameCenterDetailUpdateRequestDataWithDefaults instantiates a new GameCenterDetailUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterDetailUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterDetailUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterDetailUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterDetailUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterDetailUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterDetailUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GameCenterDetailUpdateRequestData) GetAttributes() GameCenterDetailCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterDetailUpdateRequestData) GetAttributesOk() (*GameCenterDetailCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterDetailUpdateRequestData) SetAttributes(v GameCenterDetailCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterDetailUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GameCenterDetailUpdateRequestData) GetRelationships() GameCenterDetailUpdateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterDetailUpdateRequestData) GetRelationshipsOk() (*GameCenterDetailUpdateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterDetailUpdateRequestData) SetRelationships(v GameCenterDetailUpdateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterDetailUpdateRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


