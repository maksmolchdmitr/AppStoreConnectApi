# GameCenterAchievementLocalizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterAchievementLocalization**](GameCenterAchievementLocalization.md) |  | 
**Included** | Pointer to [**[]GameCenterAchievementLocalizationsResponseIncludedInner**](GameCenterAchievementLocalizationsResponseIncludedInner.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterAchievementLocalizationResponse

`func NewGameCenterAchievementLocalizationResponse(data GameCenterAchievementLocalization, links DocumentLinks, ) *GameCenterAchievementLocalizationResponse`

NewGameCenterAchievementLocalizationResponse instantiates a new GameCenterAchievementLocalizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterAchievementLocalizationResponseWithDefaults

`func NewGameCenterAchievementLocalizationResponseWithDefaults() *GameCenterAchievementLocalizationResponse`

NewGameCenterAchievementLocalizationResponseWithDefaults instantiates a new GameCenterAchievementLocalizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterAchievementLocalizationResponse) GetData() GameCenterAchievementLocalization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterAchievementLocalizationResponse) GetDataOk() (*GameCenterAchievementLocalization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterAchievementLocalizationResponse) SetData(v GameCenterAchievementLocalization)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterAchievementLocalizationResponse) GetIncluded() []GameCenterAchievementLocalizationsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterAchievementLocalizationResponse) GetIncludedOk() (*[]GameCenterAchievementLocalizationsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterAchievementLocalizationResponse) SetIncluded(v []GameCenterAchievementLocalizationsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterAchievementLocalizationResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterAchievementLocalizationResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterAchievementLocalizationResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterAchievementLocalizationResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


