# GameCenterLeaderboardSetReleaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterLeaderboardSetRelease**](GameCenterLeaderboardSetRelease.md) |  | 
**Included** | Pointer to [**[]GameCenterLeaderboardSetReleasesResponseIncludedInner**](GameCenterLeaderboardSetReleasesResponseIncludedInner.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterLeaderboardSetReleaseResponse

`func NewGameCenterLeaderboardSetReleaseResponse(data GameCenterLeaderboardSetRelease, links DocumentLinks, ) *GameCenterLeaderboardSetReleaseResponse`

NewGameCenterLeaderboardSetReleaseResponse instantiates a new GameCenterLeaderboardSetReleaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardSetReleaseResponseWithDefaults

`func NewGameCenterLeaderboardSetReleaseResponseWithDefaults() *GameCenterLeaderboardSetReleaseResponse`

NewGameCenterLeaderboardSetReleaseResponseWithDefaults instantiates a new GameCenterLeaderboardSetReleaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterLeaderboardSetReleaseResponse) GetData() GameCenterLeaderboardSetRelease`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterLeaderboardSetReleaseResponse) GetDataOk() (*GameCenterLeaderboardSetRelease, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterLeaderboardSetReleaseResponse) SetData(v GameCenterLeaderboardSetRelease)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterLeaderboardSetReleaseResponse) GetIncluded() []GameCenterLeaderboardSetReleasesResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterLeaderboardSetReleaseResponse) GetIncludedOk() (*[]GameCenterLeaderboardSetReleasesResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterLeaderboardSetReleaseResponse) SetIncluded(v []GameCenterLeaderboardSetReleasesResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterLeaderboardSetReleaseResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardSetReleaseResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardSetReleaseResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardSetReleaseResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


