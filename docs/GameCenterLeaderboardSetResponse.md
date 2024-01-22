# GameCenterLeaderboardSetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterLeaderboardSet**](GameCenterLeaderboardSet.md) |  | 
**Included** | Pointer to [**[]GameCenterLeaderboardSetsResponseIncludedInner**](GameCenterLeaderboardSetsResponseIncludedInner.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterLeaderboardSetResponse

`func NewGameCenterLeaderboardSetResponse(data GameCenterLeaderboardSet, links DocumentLinks, ) *GameCenterLeaderboardSetResponse`

NewGameCenterLeaderboardSetResponse instantiates a new GameCenterLeaderboardSetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardSetResponseWithDefaults

`func NewGameCenterLeaderboardSetResponseWithDefaults() *GameCenterLeaderboardSetResponse`

NewGameCenterLeaderboardSetResponseWithDefaults instantiates a new GameCenterLeaderboardSetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterLeaderboardSetResponse) GetData() GameCenterLeaderboardSet`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterLeaderboardSetResponse) GetDataOk() (*GameCenterLeaderboardSet, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterLeaderboardSetResponse) SetData(v GameCenterLeaderboardSet)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterLeaderboardSetResponse) GetIncluded() []GameCenterLeaderboardSetsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterLeaderboardSetResponse) GetIncludedOk() (*[]GameCenterLeaderboardSetsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterLeaderboardSetResponse) SetIncluded(v []GameCenterLeaderboardSetsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterLeaderboardSetResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardSetResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardSetResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardSetResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


