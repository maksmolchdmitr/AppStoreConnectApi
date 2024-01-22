# GameCenterLeaderboardSetCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**GameCenterLeaderboardSetCreateRequestDataAttributes**](GameCenterLeaderboardSetCreateRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**GameCenterLeaderboardSetCreateRequestDataRelationships**](GameCenterLeaderboardSetCreateRequestDataRelationships.md) |  | [optional] 

## Methods

### NewGameCenterLeaderboardSetCreateRequestData

`func NewGameCenterLeaderboardSetCreateRequestData(type_ string, attributes GameCenterLeaderboardSetCreateRequestDataAttributes, ) *GameCenterLeaderboardSetCreateRequestData`

NewGameCenterLeaderboardSetCreateRequestData instantiates a new GameCenterLeaderboardSetCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardSetCreateRequestDataWithDefaults

`func NewGameCenterLeaderboardSetCreateRequestDataWithDefaults() *GameCenterLeaderboardSetCreateRequestData`

NewGameCenterLeaderboardSetCreateRequestDataWithDefaults instantiates a new GameCenterLeaderboardSetCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterLeaderboardSetCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterLeaderboardSetCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterLeaderboardSetCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *GameCenterLeaderboardSetCreateRequestData) GetAttributes() GameCenterLeaderboardSetCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterLeaderboardSetCreateRequestData) GetAttributesOk() (*GameCenterLeaderboardSetCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterLeaderboardSetCreateRequestData) SetAttributes(v GameCenterLeaderboardSetCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *GameCenterLeaderboardSetCreateRequestData) GetRelationships() GameCenterLeaderboardSetCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterLeaderboardSetCreateRequestData) GetRelationshipsOk() (*GameCenterLeaderboardSetCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterLeaderboardSetCreateRequestData) SetRelationships(v GameCenterLeaderboardSetCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterLeaderboardSetCreateRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


