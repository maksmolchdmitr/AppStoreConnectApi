# GameCenterGroupUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**GameCenterGroupAttributes**](GameCenterGroupAttributes.md) |  | [optional] 

## Methods

### NewGameCenterGroupUpdateRequestData

`func NewGameCenterGroupUpdateRequestData(type_ string, id string, ) *GameCenterGroupUpdateRequestData`

NewGameCenterGroupUpdateRequestData instantiates a new GameCenterGroupUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterGroupUpdateRequestDataWithDefaults

`func NewGameCenterGroupUpdateRequestDataWithDefaults() *GameCenterGroupUpdateRequestData`

NewGameCenterGroupUpdateRequestDataWithDefaults instantiates a new GameCenterGroupUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterGroupUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterGroupUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterGroupUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterGroupUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterGroupUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterGroupUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GameCenterGroupUpdateRequestData) GetAttributes() GameCenterGroupAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterGroupUpdateRequestData) GetAttributesOk() (*GameCenterGroupAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterGroupUpdateRequestData) SetAttributes(v GameCenterGroupAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterGroupUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


