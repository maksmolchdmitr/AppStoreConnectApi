# GameCenterAchievementResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterAchievement**](GameCenterAchievement.md) |  | 
**Included** | Pointer to [**[]GameCenterAchievementsResponseIncludedInner**](GameCenterAchievementsResponseIncludedInner.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterAchievementResponse

`func NewGameCenterAchievementResponse(data GameCenterAchievement, links DocumentLinks, ) *GameCenterAchievementResponse`

NewGameCenterAchievementResponse instantiates a new GameCenterAchievementResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterAchievementResponseWithDefaults

`func NewGameCenterAchievementResponseWithDefaults() *GameCenterAchievementResponse`

NewGameCenterAchievementResponseWithDefaults instantiates a new GameCenterAchievementResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterAchievementResponse) GetData() GameCenterAchievement`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterAchievementResponse) GetDataOk() (*GameCenterAchievement, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterAchievementResponse) SetData(v GameCenterAchievement)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterAchievementResponse) GetIncluded() []GameCenterAchievementsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterAchievementResponse) GetIncludedOk() (*[]GameCenterAchievementsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterAchievementResponse) SetIncluded(v []GameCenterAchievementsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterAchievementResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterAchievementResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterAchievementResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterAchievementResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


