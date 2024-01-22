# GameCenterLeaderboardReleaseResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**GameCenterLeaderboardRelease**](GameCenterLeaderboardRelease.md) |  | 
**Included** | Pointer to [**[]GameCenterLeaderboardReleasesResponseIncludedInner**](GameCenterLeaderboardReleasesResponseIncludedInner.md) |  | [optional] 
**Links** | [**DocumentLinks**](DocumentLinks.md) |  | 

## Methods

### NewGameCenterLeaderboardReleaseResponse

`func NewGameCenterLeaderboardReleaseResponse(data GameCenterLeaderboardRelease, links DocumentLinks, ) *GameCenterLeaderboardReleaseResponse`

NewGameCenterLeaderboardReleaseResponse instantiates a new GameCenterLeaderboardReleaseResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardReleaseResponseWithDefaults

`func NewGameCenterLeaderboardReleaseResponseWithDefaults() *GameCenterLeaderboardReleaseResponse`

NewGameCenterLeaderboardReleaseResponseWithDefaults instantiates a new GameCenterLeaderboardReleaseResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *GameCenterLeaderboardReleaseResponse) GetData() GameCenterLeaderboardRelease`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *GameCenterLeaderboardReleaseResponse) GetDataOk() (*GameCenterLeaderboardRelease, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *GameCenterLeaderboardReleaseResponse) SetData(v GameCenterLeaderboardRelease)`

SetData sets Data field to given value.


### GetIncluded

`func (o *GameCenterLeaderboardReleaseResponse) GetIncluded() []GameCenterLeaderboardReleasesResponseIncludedInner`

GetIncluded returns the Included field if non-nil, zero value otherwise.

### GetIncludedOk

`func (o *GameCenterLeaderboardReleaseResponse) GetIncludedOk() (*[]GameCenterLeaderboardReleasesResponseIncludedInner, bool)`

GetIncludedOk returns a tuple with the Included field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncluded

`func (o *GameCenterLeaderboardReleaseResponse) SetIncluded(v []GameCenterLeaderboardReleasesResponseIncludedInner)`

SetIncluded sets Included field to given value.

### HasIncluded

`func (o *GameCenterLeaderboardReleaseResponse) HasIncluded() bool`

HasIncluded returns a boolean if a field has been set.

### GetLinks

`func (o *GameCenterLeaderboardReleaseResponse) GetLinks() DocumentLinks`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GameCenterLeaderboardReleaseResponse) GetLinksOk() (*DocumentLinks, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GameCenterLeaderboardReleaseResponse) SetLinks(v DocumentLinks)`

SetLinks sets Links field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


