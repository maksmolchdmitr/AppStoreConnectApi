# GameCenterAchievement

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**GameCenterAchievementAttributes**](GameCenterAchievementAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GameCenterAchievementRelationships**](GameCenterAchievementRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGameCenterAchievement

`func NewGameCenterAchievement(type_ string, id string, ) *GameCenterAchievement`

NewGameCenterAchievement instantiates a new GameCenterAchievement object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterAchievementWithDefaults

`func NewGameCenterAchievementWithDefaults() *GameCenterAchievement`

NewGameCenterAchievementWithDefaults instantiates a new GameCenterAchievement object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterAchievement) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterAchievement) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterAchievement) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterAchievement) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterAchievement) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterAchievement) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GameCenterAchievement) GetAttributes() GameCenterAchievementAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterAchievement) GetAttributesOk() (*GameCenterAchievementAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterAchievement) SetAttributes(v GameCenterAchievementAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterAchievement) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GameCenterAchievement) GetRelationships() GameCenterAchievementRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterAchievement) GetRelationshipsOk() (*GameCenterAchievementRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterAchievement) SetRelationships(v GameCenterAchievementRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterAchievement) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterAchievement) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterAchievement) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterAchievement) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GameCenterAchievement) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


