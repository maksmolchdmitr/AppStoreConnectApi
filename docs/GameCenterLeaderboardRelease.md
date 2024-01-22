# GameCenterLeaderboardRelease

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**GameCenterAchievementReleaseAttributes**](GameCenterAchievementReleaseAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GameCenterLeaderboardReleaseRelationships**](GameCenterLeaderboardReleaseRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGameCenterLeaderboardRelease

`func NewGameCenterLeaderboardRelease(type_ string, id string, ) *GameCenterLeaderboardRelease`

NewGameCenterLeaderboardRelease instantiates a new GameCenterLeaderboardRelease object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardReleaseWithDefaults

`func NewGameCenterLeaderboardReleaseWithDefaults() *GameCenterLeaderboardRelease`

NewGameCenterLeaderboardReleaseWithDefaults instantiates a new GameCenterLeaderboardRelease object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterLeaderboardRelease) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterLeaderboardRelease) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterLeaderboardRelease) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterLeaderboardRelease) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterLeaderboardRelease) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterLeaderboardRelease) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GameCenterLeaderboardRelease) GetAttributes() GameCenterAchievementReleaseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterLeaderboardRelease) GetAttributesOk() (*GameCenterAchievementReleaseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterLeaderboardRelease) SetAttributes(v GameCenterAchievementReleaseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterLeaderboardRelease) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GameCenterLeaderboardRelease) GetRelationships() GameCenterLeaderboardReleaseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterLeaderboardRelease) GetRelationshipsOk() (*GameCenterLeaderboardReleaseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterLeaderboardRelease) SetRelationships(v GameCenterLeaderboardReleaseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterLeaderboardRelease) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardRelease) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardRelease) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardRelease) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GameCenterLeaderboardRelease) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


