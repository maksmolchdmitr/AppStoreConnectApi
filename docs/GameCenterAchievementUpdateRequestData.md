# GameCenterAchievementUpdateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**GameCenterAchievementUpdateRequestDataAttributes**](GameCenterAchievementUpdateRequestDataAttributes.md) |  | [optional] 

## Methods

### NewGameCenterAchievementUpdateRequestData

`func NewGameCenterAchievementUpdateRequestData(type_ string, id string, ) *GameCenterAchievementUpdateRequestData`

NewGameCenterAchievementUpdateRequestData instantiates a new GameCenterAchievementUpdateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterAchievementUpdateRequestDataWithDefaults

`func NewGameCenterAchievementUpdateRequestDataWithDefaults() *GameCenterAchievementUpdateRequestData`

NewGameCenterAchievementUpdateRequestDataWithDefaults instantiates a new GameCenterAchievementUpdateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterAchievementUpdateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterAchievementUpdateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterAchievementUpdateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterAchievementUpdateRequestData) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterAchievementUpdateRequestData) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterAchievementUpdateRequestData) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GameCenterAchievementUpdateRequestData) GetAttributes() GameCenterAchievementUpdateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterAchievementUpdateRequestData) GetAttributesOk() (*GameCenterAchievementUpdateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterAchievementUpdateRequestData) SetAttributes(v GameCenterAchievementUpdateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterAchievementUpdateRequestData) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


