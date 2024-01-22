# GameCenterDetailCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | Pointer to [**GameCenterDetailCreateRequestDataAttributes**](GameCenterDetailCreateRequestDataAttributes.md) |  | [optional] 
**Relationships** | [**AppEventCreateRequestDataRelationships**](AppEventCreateRequestDataRelationships.md) |  | 

## Methods

### NewGameCenterDetailCreateRequestData

`func NewGameCenterDetailCreateRequestData(type_ string, relationships AppEventCreateRequestDataRelationships, ) *GameCenterDetailCreateRequestData`

NewGameCenterDetailCreateRequestData instantiates a new GameCenterDetailCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterDetailCreateRequestDataWithDefaults

`func NewGameCenterDetailCreateRequestDataWithDefaults() *GameCenterDetailCreateRequestData`

NewGameCenterDetailCreateRequestDataWithDefaults instantiates a new GameCenterDetailCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterDetailCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterDetailCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterDetailCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *GameCenterDetailCreateRequestData) GetAttributes() GameCenterDetailCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterDetailCreateRequestData) GetAttributesOk() (*GameCenterDetailCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterDetailCreateRequestData) SetAttributes(v GameCenterDetailCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterDetailCreateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GameCenterDetailCreateRequestData) GetRelationships() AppEventCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterDetailCreateRequestData) GetRelationshipsOk() (*AppEventCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterDetailCreateRequestData) SetRelationships(v AppEventCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


