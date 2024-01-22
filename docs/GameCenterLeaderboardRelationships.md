# GameCenterLeaderboardRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GameCenterDetail** | Pointer to [**AppRelationshipsGameCenterDetail**](AppRelationshipsGameCenterDetail.md) |  | [optional] 
**GameCenterGroup** | Pointer to [**GameCenterAchievementRelationshipsGameCenterGroup**](GameCenterAchievementRelationshipsGameCenterGroup.md) |  | [optional] 
**GroupLeaderboard** | Pointer to [**GameCenterDetailRelationshipsDefaultLeaderboard**](GameCenterDetailRelationshipsDefaultLeaderboard.md) |  | [optional] 
**GameCenterLeaderboardSets** | Pointer to [**GameCenterDetailRelationshipsGameCenterLeaderboardSets**](GameCenterDetailRelationshipsGameCenterLeaderboardSets.md) |  | [optional] 
**Localizations** | Pointer to [**GameCenterLeaderboardRelationshipsLocalizations**](GameCenterLeaderboardRelationshipsLocalizations.md) |  | [optional] 
**Releases** | Pointer to [**GameCenterDetailRelationshipsLeaderboardReleases**](GameCenterDetailRelationshipsLeaderboardReleases.md) |  | [optional] 

## Methods

### NewGameCenterLeaderboardRelationships

`func NewGameCenterLeaderboardRelationships() *GameCenterLeaderboardRelationships`

NewGameCenterLeaderboardRelationships instantiates a new GameCenterLeaderboardRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterLeaderboardRelationshipsWithDefaults

`func NewGameCenterLeaderboardRelationshipsWithDefaults() *GameCenterLeaderboardRelationships`

NewGameCenterLeaderboardRelationshipsWithDefaults instantiates a new GameCenterLeaderboardRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGameCenterDetail

`func (o *GameCenterLeaderboardRelationships) GetGameCenterDetail() AppRelationshipsGameCenterDetail`

GetGameCenterDetail returns the GameCenterDetail field if non-nil, zero value otherwise.

### GetGameCenterDetailOk

`func (o *GameCenterLeaderboardRelationships) GetGameCenterDetailOk() (*AppRelationshipsGameCenterDetail, bool)`

GetGameCenterDetailOk returns a tuple with the GameCenterDetail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameCenterDetail

`func (o *GameCenterLeaderboardRelationships) SetGameCenterDetail(v AppRelationshipsGameCenterDetail)`

SetGameCenterDetail sets GameCenterDetail field to given value.

### HasGameCenterDetail

`func (o *GameCenterLeaderboardRelationships) HasGameCenterDetail() bool`

HasGameCenterDetail returns a boolean if a field has been set.

### GetGameCenterGroup

`func (o *GameCenterLeaderboardRelationships) GetGameCenterGroup() GameCenterAchievementRelationshipsGameCenterGroup`

GetGameCenterGroup returns the GameCenterGroup field if non-nil, zero value otherwise.

### GetGameCenterGroupOk

`func (o *GameCenterLeaderboardRelationships) GetGameCenterGroupOk() (*GameCenterAchievementRelationshipsGameCenterGroup, bool)`

GetGameCenterGroupOk returns a tuple with the GameCenterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameCenterGroup

`func (o *GameCenterLeaderboardRelationships) SetGameCenterGroup(v GameCenterAchievementRelationshipsGameCenterGroup)`

SetGameCenterGroup sets GameCenterGroup field to given value.

### HasGameCenterGroup

`func (o *GameCenterLeaderboardRelationships) HasGameCenterGroup() bool`

HasGameCenterGroup returns a boolean if a field has been set.

### GetGroupLeaderboard

`func (o *GameCenterLeaderboardRelationships) GetGroupLeaderboard() GameCenterDetailRelationshipsDefaultLeaderboard`

GetGroupLeaderboard returns the GroupLeaderboard field if non-nil, zero value otherwise.

### GetGroupLeaderboardOk

`func (o *GameCenterLeaderboardRelationships) GetGroupLeaderboardOk() (*GameCenterDetailRelationshipsDefaultLeaderboard, bool)`

GetGroupLeaderboardOk returns a tuple with the GroupLeaderboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupLeaderboard

`func (o *GameCenterLeaderboardRelationships) SetGroupLeaderboard(v GameCenterDetailRelationshipsDefaultLeaderboard)`

SetGroupLeaderboard sets GroupLeaderboard field to given value.

### HasGroupLeaderboard

`func (o *GameCenterLeaderboardRelationships) HasGroupLeaderboard() bool`

HasGroupLeaderboard returns a boolean if a field has been set.

### GetGameCenterLeaderboardSets

`func (o *GameCenterLeaderboardRelationships) GetGameCenterLeaderboardSets() GameCenterDetailRelationshipsGameCenterLeaderboardSets`

GetGameCenterLeaderboardSets returns the GameCenterLeaderboardSets field if non-nil, zero value otherwise.

### GetGameCenterLeaderboardSetsOk

`func (o *GameCenterLeaderboardRelationships) GetGameCenterLeaderboardSetsOk() (*GameCenterDetailRelationshipsGameCenterLeaderboardSets, bool)`

GetGameCenterLeaderboardSetsOk returns a tuple with the GameCenterLeaderboardSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameCenterLeaderboardSets

`func (o *GameCenterLeaderboardRelationships) SetGameCenterLeaderboardSets(v GameCenterDetailRelationshipsGameCenterLeaderboardSets)`

SetGameCenterLeaderboardSets sets GameCenterLeaderboardSets field to given value.

### HasGameCenterLeaderboardSets

`func (o *GameCenterLeaderboardRelationships) HasGameCenterLeaderboardSets() bool`

HasGameCenterLeaderboardSets returns a boolean if a field has been set.

### GetLocalizations

`func (o *GameCenterLeaderboardRelationships) GetLocalizations() GameCenterLeaderboardRelationshipsLocalizations`

GetLocalizations returns the Localizations field if non-nil, zero value otherwise.

### GetLocalizationsOk

`func (o *GameCenterLeaderboardRelationships) GetLocalizationsOk() (*GameCenterLeaderboardRelationshipsLocalizations, bool)`

GetLocalizationsOk returns a tuple with the Localizations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocalizations

`func (o *GameCenterLeaderboardRelationships) SetLocalizations(v GameCenterLeaderboardRelationshipsLocalizations)`

SetLocalizations sets Localizations field to given value.

### HasLocalizations

`func (o *GameCenterLeaderboardRelationships) HasLocalizations() bool`

HasLocalizations returns a boolean if a field has been set.

### GetReleases

`func (o *GameCenterLeaderboardRelationships) GetReleases() GameCenterDetailRelationshipsLeaderboardReleases`

GetReleases returns the Releases field if non-nil, zero value otherwise.

### GetReleasesOk

`func (o *GameCenterLeaderboardRelationships) GetReleasesOk() (*GameCenterDetailRelationshipsLeaderboardReleases, bool)`

GetReleasesOk returns a tuple with the Releases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReleases

`func (o *GameCenterLeaderboardRelationships) SetReleases(v GameCenterDetailRelationshipsLeaderboardReleases)`

SetReleases sets Releases field to given value.

### HasReleases

`func (o *GameCenterLeaderboardRelationships) HasReleases() bool`

HasReleases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


