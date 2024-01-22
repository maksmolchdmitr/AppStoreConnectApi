# GameCenterLeaderboardSetLocalizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterLeaderboardSetLocalization**](GameCenterLeaderboardSetLocalization.md) |  | 
**Included** | Pointer to [**[]GameCenterLeaderboardSetLocalizationsResponseIncludedInner**](GameCenterLeaderboardSetLocalizationsResponseIncludedInner.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterLeaderboardSetLocalizationResponse

`func NewGameCenterLeaderboardSetLocalizationResponse(data GameCenterLeaderboardSetLocalization, links DocumentLinks, ) *GameCenterLeaderboardSetLocalizationResponse`

NewGameCenterLeaderboardSetLocalizationResponse instantiates a new GameCenterLeaderboardSetLocalizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardSetLocalizationResponseWithDefaults

`func NewGameCenterLeaderboardSetLocalizationResponseWithDefaults() *GameCenterLeaderboardSetLocalizationResponse`

NewGameCenterLeaderboardSetLocalizationResponseWithDefaults instantiates a new GameCenterLeaderboardSetLocalizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterLeaderboardSetLocalizationResponse) GetData() GameCenterLeaderboardSetLocalization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterLeaderboardSetLocalizationResponse) GetDataOk() (*GameCenterLeaderboardSetLocalization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterLeaderboardSetLocalizationResponse) SetData(v GameCenterLeaderboardSetLocalization)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterLeaderboardSetLocalizationResponse) GetIncluded() []GameCenterLeaderboardSetLocalizationsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterLeaderboardSetLocalizationResponse) GetIncludedOk() (*[]GameCenterLeaderboardSetLocalizationsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterLeaderboardSetLocalizationResponse) SetIncluded(v []GameCenterLeaderboardSetLocalizationsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterLeaderboardSetLocalizationResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardSetLocalizationResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardSetLocalizationResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardSetLocalizationResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


