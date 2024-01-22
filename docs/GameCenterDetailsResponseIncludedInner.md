# GameCenterDetailsResponseIncludedInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**GameCenterAchievementReleaseAttributes**](GameCenterAchievementReleaseAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GameCenterLeaderboardSetReleaseRelationships**](GameCenterLeaderboardSetReleaseRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGameCenterDetailsResponseIncludedInner

`func NewGameCenterDetailsResponseIncludedInner(type_ string, id string, ) *GameCenterDetailsResponseIncludedInner`

NewGameCenterDetailsResponseIncludedInner instantiates a new GameCenterDetailsResponseIncludedInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterDetailsResponseIncludedInnerWithDefaults

`func NewGameCenterDetailsResponseIncludedInnerWithDefaults() *GameCenterDetailsResponseIncludedInner`

NewGameCenterDetailsResponseIncludedInnerWithDefaults instantiates a new GameCenterDetailsResponseIncludedInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterDetailsResponseIncludedInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterDetailsResponseIncludedInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterDetailsResponseIncludedInner) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterDetailsResponseIncludedInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterDetailsResponseIncludedInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterDetailsResponseIncludedInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GameCenterDetailsResponseIncludedInner) GetAttributes() GameCenterAchievementReleaseAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterDetailsResponseIncludedInner) GetAttributesOk() (*GameCenterAchievementReleaseAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterDetailsResponseIncludedInner) SetAttributes(v GameCenterAchievementReleaseAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterDetailsResponseIncludedInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GameCenterDetailsResponseIncludedInner) GetRelationships() GameCenterLeaderboardSetReleaseRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterDetailsResponseIncludedInner) GetRelationshipsOk() (*GameCenterLeaderboardSetReleaseRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterDetailsResponseIncludedInner) SetRelationships(v GameCenterLeaderboardSetReleaseRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterDetailsResponseIncludedInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterDetailsResponseIncludedInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterDetailsResponseIncludedInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterDetailsResponseIncludedInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GameCenterDetailsResponseIncludedInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


