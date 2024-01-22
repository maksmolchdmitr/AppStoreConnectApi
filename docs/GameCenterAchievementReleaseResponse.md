# GameCenterAchievementReleaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterAchievementRelease**](GameCenterAchievementRelease.md) |  | 
**Included** | Pointer to [**[]GameCenterAchievementReleasesResponseIncludedInner**](GameCenterAchievementReleasesResponseIncludedInner.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterAchievementReleaseResponse

`func NewGameCenterAchievementReleaseResponse(data GameCenterAchievementRelease, links DocumentLinks, ) *GameCenterAchievementReleaseResponse`

NewGameCenterAchievementReleaseResponse instantiates a new GameCenterAchievementReleaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterAchievementReleaseResponseWithDefaults

`func NewGameCenterAchievementReleaseResponseWithDefaults() *GameCenterAchievementReleaseResponse`

NewGameCenterAchievementReleaseResponseWithDefaults instantiates a new GameCenterAchievementReleaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterAchievementReleaseResponse) GetData() GameCenterAchievementRelease`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterAchievementReleaseResponse) GetDataOk() (*GameCenterAchievementRelease, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterAchievementReleaseResponse) SetData(v GameCenterAchievementRelease)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterAchievementReleaseResponse) GetIncluded() []GameCenterAchievementReleasesResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterAchievementReleaseResponse) GetIncludedOk() (*[]GameCenterAchievementReleasesResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterAchievementReleaseResponse) SetIncluded(v []GameCenterAchievementReleasesResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterAchievementReleaseResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterAchievementReleaseResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterAchievementReleaseResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterAchievementReleaseResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


