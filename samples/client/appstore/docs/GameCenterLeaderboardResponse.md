# GameCenterLeaderboardResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterLeaderboard**](GameCenterLeaderboard.md) |  | 
**Included** | Pointer to [**[]GameCenterLeaderboardsResponseIncludedInner**](GameCenterLeaderboardsResponseIncludedInner.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterLeaderboardResponse

`func NewGameCenterLeaderboardResponse(data GameCenterLeaderboard, links DocumentLinks, ) *GameCenterLeaderboardResponse`

NewGameCenterLeaderboardResponse instantiates a new GameCenterLeaderboardResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardResponseWithDefaults

`func NewGameCenterLeaderboardResponseWithDefaults() *GameCenterLeaderboardResponse`

NewGameCenterLeaderboardResponseWithDefaults instantiates a new GameCenterLeaderboardResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterLeaderboardResponse) GetData() GameCenterLeaderboard`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterLeaderboardResponse) GetDataOk() (*GameCenterLeaderboard, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterLeaderboardResponse) SetData(v GameCenterLeaderboard)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterLeaderboardResponse) GetIncluded() []GameCenterLeaderboardsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterLeaderboardResponse) GetIncludedOk() (*[]GameCenterLeaderboardsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterLeaderboardResponse) SetIncluded(v []GameCenterLeaderboardsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterLeaderboardResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


