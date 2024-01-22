# GameCenterLeaderboardSetImage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**GameCenterAchievementImageAttributes**](GameCenterAchievementImageAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GameCenterLeaderboardSetImageRelationships**](GameCenterLeaderboardSetImageRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGameCenterLeaderboardSetImage

`func NewGameCenterLeaderboardSetImage(type_ string, id string, ) *GameCenterLeaderboardSetImage`

NewGameCenterLeaderboardSetImage instantiates a new GameCenterLeaderboardSetImage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardSetImageWithDefaults

`func NewGameCenterLeaderboardSetImageWithDefaults() *GameCenterLeaderboardSetImage`

NewGameCenterLeaderboardSetImageWithDefaults instantiates a new GameCenterLeaderboardSetImage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterLeaderboardSetImage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterLeaderboardSetImage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterLeaderboardSetImage) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterLeaderboardSetImage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterLeaderboardSetImage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterLeaderboardSetImage) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GameCenterLeaderboardSetImage) GetAttributes() GameCenterAchievementImageAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterLeaderboardSetImage) GetAttributesOk() (*GameCenterAchievementImageAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterLeaderboardSetImage) SetAttributes(v GameCenterAchievementImageAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterLeaderboardSetImage) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GameCenterLeaderboardSetImage) GetRelationships() GameCenterLeaderboardSetImageRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterLeaderboardSetImage) GetRelationshipsOk() (*GameCenterLeaderboardSetImageRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterLeaderboardSetImage) SetRelationships(v GameCenterLeaderboardSetImageRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterLeaderboardSetImage) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardSetImage) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardSetImage) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardSetImage) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GameCenterLeaderboardSetImage) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


