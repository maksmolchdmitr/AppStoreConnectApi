# GameCenterAchievementCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**GameCenterAchievementCreateRequestDataAttributes**](GameCenterAchievementCreateRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**GameCenterAchievementCreateRequestDataRelationships**](GameCenterAchievementCreateRequestDataRelationships.md) |  | [optional] 

## Methods

### NewGameCenterAchievementCreateRequestData

`func NewGameCenterAchievementCreateRequestData(type_ string, attributes GameCenterAchievementCreateRequestDataAttributes, ) *GameCenterAchievementCreateRequestData`

NewGameCenterAchievementCreateRequestData instantiates a new GameCenterAchievementCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterAchievementCreateRequestDataWithDefaults

`func NewGameCenterAchievementCreateRequestDataWithDefaults() *GameCenterAchievementCreateRequestData`

NewGameCenterAchievementCreateRequestDataWithDefaults instantiates a new GameCenterAchievementCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterAchievementCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterAchievementCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterAchievementCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *GameCenterAchievementCreateRequestData) GetAttributes() GameCenterAchievementCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterAchievementCreateRequestData) GetAttributesOk() (*GameCenterAchievementCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterAchievementCreateRequestData) SetAttributes(v GameCenterAchievementCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *GameCenterAchievementCreateRequestData) GetRelationships() GameCenterAchievementCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterAchievementCreateRequestData) GetRelationshipsOk() (*GameCenterAchievementCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterAchievementCreateRequestData) SetRelationships(v GameCenterAchievementCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterAchievementCreateRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


