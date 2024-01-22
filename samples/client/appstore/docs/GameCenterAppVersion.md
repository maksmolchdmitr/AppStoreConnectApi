# GameCenterAppVersion

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** |  | 
**Id** | **string** |  | 
**Attributes** | Pointer to [**GameCenterAppVersionAttributes**](GameCenterAppVersionAttributes.md) |  | [optional] 
**Relationships** | Pointer to [**GameCenterAppVersionRelationships**](GameCenterAppVersionRelationships.md) |  | [optional] 
**Links** | Pointer to [**ResourceLinks**](ResourceLinks.md) |  | [optional] 

## Methods

### NewGameCenterAppVersion

`func NewGameCenterAppVersion(type_ string, id string, ) *GameCenterAppVersion`

NewGameCenterAppVersion instantiates a new GameCenterAppVersion object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterAppVersionWithDefaults

`func NewGameCenterAppVersionWithDefaults() *GameCenterAppVersion`

NewGameCenterAppVersionWithDefaults instantiates a new GameCenterAppVersion object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *GameCenterAppVersion) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *GameCenterAppVersion) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *GameCenterAppVersion) SetType(v string)`

SetType sets Type field to given value.


### GetId

`func (o *GameCenterAppVersion) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GameCenterAppVersion) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GameCenterAppVersion) SetId(v string)`

SetId sets Id field to given value.


### GetAttributes

`func (o *GameCenterAppVersion) GetAttributes() GameCenterAppVersionAttributes`

GetAttributes returns the Attributes field if non-nil, zero value otherwise.

### GetAttributesOk

`func (o *GameCenterAppVersion) GetAttributesOk() (*GameCenterAppVersionAttributes, bool)`

GetAttributesOk returns a tuple with the Attributes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttributes

`func (o *GameCenterAppVersion) SetAttributes(v GameCenterAppVersionAttributes)`

SetAttributes sets Attributes field to given value.

### HasAttributes

`func (o *GameCenterAppVersion) HasAttributes() bool`

HasAttributes returns a boolean if a field has been set.

### GetRelationships

`func (o *GameCenterAppVersion) GetRelationships() GameCenterAppVersionRelationships`

GetRelationships returns the Relationships field if non-nil, zero value otherwise.

### GetRelationshipsOk

`func (o *GameCenterAppVersion) GetRelationshipsOk() (*GameCenterAppVersionRelationships, bool)`

GetRelationshipsOk returns a tuple with the Relationships field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelationships

`func (o *GameCenterAppVersion) SetRelationships(v GameCenterAppVersionRelationships)`

SetRelationships sets Relationships field to given value.

### HasRelationships

`func (o *GameCenterAppVersion) HasRelationships() bool`

HasRelationships returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterAppVersion) GetLinks() ResourceLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterAppVersion) GetLinksOk() (*ResourceLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterAppVersion) SetLinks(v ResourceLinks)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GameCenterAppVersion) HasLinks() bool`

HasLinks returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


