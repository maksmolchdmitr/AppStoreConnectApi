# GameCenterMatchmakingTestRequestInlineCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | Pointer to **string** |  | [optional] 
**Attributes** | [**GameCenterMatchmakingTestRequestInlineCreateAttributes**](GameCenterMatchmakingTestRequestInlineCreateAttributes.md) |  | 
**Relationships** | Pointer to [**GameCenterMatchmakingTestRequestInlineCreateRelationships**](GameCenterMatchmakingTestRequestInlineCreateRelationships.md) |  | [optional] 

## Methods

### NewGameCenterMatchmakingTestRequestInlineCreate

`func NewGameCenterMatchmakingTestRequestInlineCreate(type_ string, attributes GameCenterMatchmakingTestRequestInlineCreateAttributes, ) *GameCenterMatchmakingTestRequestInlineCreate`

NewGameCenterMatchmakingTestRequestInlineCreate instantiates a new GameCenterMatchmakingTestRequestInlineCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterMatchmakingTestRequestInlineCreateWithDefaults

`func NewGameCenterMatchmakingTestRequestInlineCreateWithDefaults() *GameCenterMatchmakingTestRequestInlineCreate`

NewGameCenterMatchmakingTestRequestInlineCreateWithDefaults instantiates a new GameCenterMatchmakingTestRequestInlineCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterMatchmakingTestRequestInlineCreate) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterMatchmakingTestRequestInlineCreate) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterMatchmakingTestRequestInlineCreate) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterMatchmakingTestRequestInlineCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterMatchmakingTestRequestInlineCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterMatchmakingTestRequestInlineCreate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GameCenterMatchmakingTestRequestInlineCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetAttributes

`func (o *GameCenterMatchmakingTestRequestInlineCreate) GetAttributes() GameCenterMatchmakingTestRequestInlineCreateAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterMatchmakingTestRequestInlineCreate) GetAttributesOk() (*GameCenterMatchmakingTestRequestInlineCreateAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterMatchmakingTestRequestInlineCreate) SetAttributes(v GameCenterMatchmakingTestRequestInlineCreateAttributes)`

SetAttributes sets Attributes field to given value.


### GetRelationships

`func (o *GameCenterMatchmakingTestRequestInlineCreate) GetRelationships() GameCenterMatchmakingTestRequestInlineCreateRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterMatchmakingTestRequestInlineCreate) GetRelationshipsOk() (*GameCenterMatchmakingTestRequestInlineCreateRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterMatchmakingTestRequestInlineCreate) SetRelationships(v GameCenterMatchmakingTestRequestInlineCreateRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterMatchmakingTestRequestInlineCreate) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


