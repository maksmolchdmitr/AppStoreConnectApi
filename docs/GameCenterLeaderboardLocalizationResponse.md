# GameCenterLeaderboardLocalizationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterLeaderboardLocalization**](GameCenterLeaderboardLocalization.md) |  | 
**Included** | Pointer to [**[]GameCenterLeaderboardLocalizationsResponseIncludedInner**](GameCenterLeaderboardLocalizationsResponseIncludedInner.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterLeaderboardLocalizationResponse

`func NewGameCenterLeaderboardLocalizationResponse(data GameCenterLeaderboardLocalization, links DocumentLinks, ) *GameCenterLeaderboardLocalizationResponse`

NewGameCenterLeaderboardLocalizationResponse instantiates a new GameCenterLeaderboardLocalizationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardLocalizationResponseWithDefaults

`func NewGameCenterLeaderboardLocalizationResponseWithDefaults() *GameCenterLeaderboardLocalizationResponse`

NewGameCenterLeaderboardLocalizationResponseWithDefaults instantiates a new GameCenterLeaderboardLocalizationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterLeaderboardLocalizationResponse) GetData() GameCenterLeaderboardLocalization`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterLeaderboardLocalizationResponse) GetDataOk() (*GameCenterLeaderboardLocalization, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterLeaderboardLocalizationResponse) SetData(v GameCenterLeaderboardLocalization)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterLeaderboardLocalizationResponse) GetIncluded() []GameCenterLeaderboardLocalizationsResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterLeaderboardLocalizationResponse) GetIncludedOk() (*[]GameCenterLeaderboardLocalizationsResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterLeaderboardLocalizationResponse) SetIncluded(v []GameCenterLeaderboardLocalizationsResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterLeaderboardLocalizationResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardLocalizationResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardLocalizationResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardLocalizationResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


