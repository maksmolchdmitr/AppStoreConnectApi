# GameCenterMatchmakingRuleSetsResponseIncludedInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**GameCenterGroupAttributes**](GameCenterGroupAttributes.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 
**Relationships** | Pointer to [**GameCenterMatchmakingQueueRelationships**](GameCenterMatchmakingQueueRelationships.md) |  | [optional] 

## Methods

### NewGameCenterMatchmakingRuleSetsResponseIncludedInner

`func NewGameCenterMatchmakingRuleSetsResponseIncludedInner(type_ string, id string, ) *GameCenterMatchmakingRuleSetsResponseIncludedInner`

NewGameCenterMatchmakingRuleSetsResponseIncludedInner instantiates a new GameCenterMatchmakingRuleSetsResponseIncludedInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterMatchmakingRuleSetsResponseIncludedInnerWithDefaults

`func NewGameCenterMatchmakingRuleSetsResponseIncludedInnerWithDefaults() *GameCenterMatchmakingRuleSetsResponseIncludedInner`

NewGameCenterMatchmakingRuleSetsResponseIncludedInnerWithDefaults instantiates a new GameCenterMatchmakingRuleSetsResponseIncludedInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) GetAttributes() GameCenterGroupAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) GetAttributesOk() (*GameCenterGroupAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) SetAttributes(v GameCenterGroupAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetRelationships

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) GetRelationships() GameCenterMatchmakingQueueRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) GetRelationshipsOk() (*GameCenterMatchmakingQueueRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) SetRelationships(v GameCenterMatchmakingQueueRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterMatchmakingRuleSetsResponseIncludedInner) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


