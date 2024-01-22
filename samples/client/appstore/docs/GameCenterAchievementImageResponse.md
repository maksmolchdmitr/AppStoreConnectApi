# GameCenterAchievementImageResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterAchievementImage**](GameCenterAchievementImage.md) |  | 
**Included** | Pointer to [**[]GameCenterAchievementLocalization**](GameCenterAchievementLocalization.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterAchievementImageResponse

`func NewGameCenterAchievementImageResponse(data GameCenterAchievementImage, links DocumentLinks, ) *GameCenterAchievementImageResponse`

NewGameCenterAchievementImageResponse instantiates a new GameCenterAchievementImageResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterAchievementImageResponseWithDefaults

`func NewGameCenterAchievementImageResponseWithDefaults() *GameCenterAchievementImageResponse`

NewGameCenterAchievementImageResponseWithDefaults instantiates a new GameCenterAchievementImageResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterAchievementImageResponse) GetData() GameCenterAchievementImage`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterAchievementImageResponse) GetDataOk() (*GameCenterAchievementImage, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterAchievementImageResponse) SetData(v GameCenterAchievementImage)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterAchievementImageResponse) GetIncluded() []GameCenterAchievementLocalization`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterAchievementImageResponse) GetIncludedOk() (*[]GameCenterAchievementLocalization, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterAchievementImageResponse) SetIncluded(v []GameCenterAchievementLocalization)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterAchievementImageResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterAchievementImageResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterAchievementImageResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterAchievementImageResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


