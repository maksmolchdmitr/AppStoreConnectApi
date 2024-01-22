# GameCenterAchievementRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**GameCenterAchievementReleaseAttributes**](GameCenterAchievementReleaseAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GameCenterAchievementReleaseRelationships**](GameCenterAchievementReleaseRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGameCenterAchievementRelease

`func NewGameCenterAchievementRelease(type_ string, id string, ) *GameCenterAchievementRelease`

NewGameCenterAchievementRelease instantiates a new GameCenterAchievementRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterAchievementReleaseWithDefaults

`func NewGameCenterAchievementReleaseWithDefaults() *GameCenterAchievementRelease`

NewGameCenterAchievementReleaseWithDefaults instantiates a new GameCenterAchievementRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterAchievementRelease) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterAchievementRelease) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterAchievementRelease) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterAchievementRelease) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterAchievementRelease) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterAchievementRelease) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GameCenterAchievementRelease) GetAttributes() GameCenterAchievementReleaseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterAchievementRelease) GetAttributesOk() (*GameCenterAchievementReleaseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterAchievementRelease) SetAttributes(v GameCenterAchievementReleaseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterAchievementRelease) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GameCenterAchievementRelease) GetRelationships() GameCenterAchievementReleaseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterAchievementRelease) GetRelationshipsOk() (*GameCenterAchievementReleaseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterAchievementRelease) SetRelationships(v GameCenterAchievementReleaseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterAchievementRelease) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterAchievementRelease) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterAchievementRelease) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterAchievementRelease) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GameCenterAchievementRelease) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


