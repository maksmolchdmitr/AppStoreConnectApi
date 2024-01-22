# GameCenterLeaderboardCreateRequestData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Attributes** | [**GameCenterLeaderboardCreateRequestDataAttributes**](GameCenterLeaderboardCreateRequestDataAttributes.md) |  | 
**Relationships** | Pointer to [**GameCenterLeaderboardCreateRequestDataRelationships**](GameCenterLeaderboardCreateRequestDataRelationships.md) |  | [optional] 

## Methods

### NewGameCenterLeaderboardCreateRequestData

`func NewGameCenterLeaderboardCreateRequestData(type_ string, attributes GameCenterLeaderboardCreateRequestDataAttributes, ) *GameCenterLeaderboardCreateRequestData`

NewGameCenterLeaderboardCreateRequestData instantiates a new GameCenterLeaderboardCreateRequestData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardCreateRequestDataWithDefaults

`func NewGameCenterLeaderboardCreateRequestDataWithDefaults() *GameCenterLeaderboardCreateRequestData`

NewGameCenterLeaderboardCreateRequestDataWithDefaults instantiates a new GameCenterLeaderboardCreateRequestData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterLeaderboardCreateRequestData) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterLeaderboardCreateRequestData) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterLeaderboardCreateRequestData) SetType(v string)`

SetType sets Type field to given value.


### GetAttributes

`func (o *GameCenterLeaderboardCreateRequestData) GetAttributes() GameCenterLeaderboardCreateRequestDataAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterLeaderboardCreateRequestData) GetAttributesOk() (*GameCenterLeaderboardCreateRequestDataAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterLeaderboardCreateRequestData) SetAttributes(v GameCenterLeaderboardCreateRequestDataAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *GameCenterLeaderboardCreateRequestData) GetRelationships() GameCenterLeaderboardCreateRequestDataRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterLeaderboardCreateRequestData) GetRelationshipsOk() (*GameCenterLeaderboardCreateRequestDataRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterLeaderboardCreateRequestData) SetRelationships(v GameCenterLeaderboardCreateRequestDataRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterLeaderboardCreateRequestData) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


