# GameCenterGroupResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterGroup**](GameCenterGroup.md) |  | 
**Included** | Pointer to [**[]GameCenterGroupsResponseIncludedInner**](GameCenterGroupsResponseIncludedInner.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterGroupResponse

`func NewGameCenterGroupResponse(data GameCenterGroup, links DocumentLinks, ) *GameCenterGroupResponse`

NewGameCenterGroupResponse instantiates a new GameCenterGroupResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterGroupResponseWithDefaults

`func NewGameCenterGroupResponseWithDefaults() *GameCenterGroupResponse`

NewGameCenterGroupResponseWithDefaults instantiates a new GameCenterGroupResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterGroupResponse) GetData() GameCenterGroup`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterGroupResponse) GetDataOk() (*GameCenterGroup, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterGroupResponse) SetData(v GameCenterGroup)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterGroupResponse) GetIncluded() []GameCenterGroupsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterGroupResponse) GetIncludedOk() (*[]GameCenterGroupsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterGroupResponse) SetIncluded(v []GameCenterGroupsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterGroupResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterGroupResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterGroupResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterGroupResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


