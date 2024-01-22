# GameCenterMatchmakingQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**GameCenterGroupAttributes**](GameCenterGroupAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GameCenterMatchmakingQueueRelationships**](GameCenterMatchmakingQueueRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGameCenterMatchmakingQueue

`func NewGameCenterMatchmakingQueue(type_ string, id string, ) *GameCenterMatchmakingQueue`

NewGameCenterMatchmakingQueue instantiates a new GameCenterMatchmakingQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterMatchmakingQueueWithDefaults

`func NewGameCenterMatchmakingQueueWithDefaults() *GameCenterMatchmakingQueue`

NewGameCenterMatchmakingQueueWithDefaults instantiates a new GameCenterMatchmakingQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterMatchmakingQueue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterMatchmakingQueue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterMatchmakingQueue) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterMatchmakingQueue) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterMatchmakingQueue) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterMatchmakingQueue) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GameCenterMatchmakingQueue) GetAttributes() GameCenterGroupAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterMatchmakingQueue) GetAttributesOk() (*GameCenterGroupAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterMatchmakingQueue) SetAttributes(v GameCenterGroupAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterMatchmakingQueue) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GameCenterMatchmakingQueue) GetRelationships() GameCenterMatchmakingQueueRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterMatchmakingQueue) GetRelationshipsOk() (*GameCenterMatchmakingQueueRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterMatchmakingQueue) SetRelationships(v GameCenterMatchmakingQueueRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterMatchmakingQueue) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterMatchmakingQueue) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterMatchmakingQueue) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterMatchmakingQueue) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GameCenterMatchmakingQueue) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


