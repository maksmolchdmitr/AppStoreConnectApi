# GameCenterLeaderboardSetRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**GameCenterAchievementReleaseAttributes**](GameCenterAchievementReleaseAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GameCenterLeaderboardSetReleaseRelationships**](GameCenterLeaderboardSetReleaseRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGameCenterLeaderboardSetRelease

`func NewGameCenterLeaderboardSetRelease(type_ string, id string, ) *GameCenterLeaderboardSetRelease`

NewGameCenterLeaderboardSetRelease instantiates a new GameCenterLeaderboardSetRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardSetReleaseWithDefaults

`func NewGameCenterLeaderboardSetReleaseWithDefaults() *GameCenterLeaderboardSetRelease`

NewGameCenterLeaderboardSetReleaseWithDefaults instantiates a new GameCenterLeaderboardSetRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterLeaderboardSetRelease) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterLeaderboardSetRelease) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterLeaderboardSetRelease) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterLeaderboardSetRelease) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterLeaderboardSetRelease) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterLeaderboardSetRelease) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GameCenterLeaderboardSetRelease) GetAttributes() GameCenterAchievementReleaseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterLeaderboardSetRelease) GetAttributesOk() (*GameCenterAchievementReleaseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterLeaderboardSetRelease) SetAttributes(v GameCenterAchievementReleaseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterLeaderboardSetRelease) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GameCenterLeaderboardSetRelease) GetRelationships() GameCenterLeaderboardSetReleaseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterLeaderboardSetRelease) GetRelationshipsOk() (*GameCenterLeaderboardSetReleaseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterLeaderboardSetRelease) SetRelationships(v GameCenterLeaderboardSetReleaseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterLeaderboardSetRelease) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardSetRelease) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardSetRelease) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardSetRelease) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GameCenterLeaderboardSetRelease) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


