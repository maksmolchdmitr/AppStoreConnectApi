# GameCenterDetailRelationships

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**App** | Pointer to [**AppAvailabilityRelationshipsApp**](AppAvailabilityRelationshipsApp.md) |  | [optional] 
**GameCenterAppVersions** | Pointer to [**GameCenterAppVersionRelationshipsCompatibilityVersions**](GameCenterAppVersionRelationshipsCompatibilityVersions.md) |  | [optional] 
**GameCenterGroup** | Pointer to [**GameCenterAchievementRelationshipsGameCenterGroup**](GameCenterAchievementRelationshipsGameCenterGroup.md) |  | [optional] 
**GameCenterLeaderboards** | Pointer to [**GameCenterDetailRelationshipsGameCenterLeaderboards**](GameCenterDetailRelationshipsGameCenterLeaderboards.md) |  | [optional] 
**GameCenterLeaderboardSets** | Pointer to [**GameCenterDetailRelationshipsGameCenterLeaderboardSets**](GameCenterDetailRelationshipsGameCenterLeaderboardSets.md) |  | [optional] 
**GameCenterAchievements** | Pointer to [**GameCenterDetailRelationshipsGameCenterAchievements**](GameCenterDetailRelationshipsGameCenterAchievements.md) |  | [optional] 
**DefaultLeaderboard** | Pointer to [**GameCenterDetailRelationshipsDefaultLeaderboard**](GameCenterDetailRelationshipsDefaultLeaderboard.md) |  | [optional] 
**DefaultGroupLeaderboard** | Pointer to [**GameCenterDetailRelationshipsDefaultLeaderboard**](GameCenterDetailRelationshipsDefaultLeaderboard.md) |  | [optional] 
**AchievementReleases** | Pointer to [**GameCenterAchievementRelationshipsReleases**](GameCenterAchievementRelationshipsReleases.md) |  | [optional] 
**LeaderboardReleases** | Pointer to [**GameCenterDetailRelationshipsLeaderboardReleases**](GameCenterDetailRelationshipsLeaderboardReleases.md) |  | [optional] 
**LeaderboardSetReleases** | Pointer to [**GameCenterDetailRelationshipsLeaderboardSetReleases**](GameCenterDetailRelationshipsLeaderboardSetReleases.md) |  | [optional] 

## Methods

### NewGameCenterDetailRelationships

`func NewGameCenterDetailRelationships() *GameCenterDetailRelationships`

NewGameCenterDetailRelationships instantiates a new GameCenterDetailRelationships object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterDetailRelationshipsWithDefaults

`func NewGameCenterDetailRelationshipsWithDefaults() *GameCenterDetailRelationships`

NewGameCenterDetailRelationshipsWithDefaults instantiates a new GameCenterDetailRelationships object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApp

`func (o *GameCenterDetailRelationships) GetApp() AppAvailabilityRelationshipsApp`

GetApp returns the App field if non-nil, zero value otherwise.

### GetAppOk

`func (o *GameCenterDetailRelationships) GetAppOk() (*AppAvailabilityRelationshipsApp, bool)`

GetAppOk returns a tuple with the App field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApp

`func (o *GameCenterDetailRelationships) SetApp(v AppAvailabilityRelationshipsApp)`

SetApp sets App field to given value.

### HasApp

`func (o *GameCenterDetailRelationships) HasApp() bool`

HasApp returns a boolean if a field has been set.

### GetGameCenterAppVersions

`func (o *GameCenterDetailRelationships) GetGameCenterAppVersions() GameCenterAppVersionRelationshipsCompatibilityVersions`

GetGameCenterAppVersions returns the GameCenterAppVersions field if non-nil, zero value otherwise.

### GetGameCenterAppVersionsOk

`func (o *GameCenterDetailRelationships) GetGameCenterAppVersionsOk() (*GameCenterAppVersionRelationshipsCompatibilityVersions, bool)`

GetGameCenterAppVersionsOk returns a tuple with the GameCenterAppVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameCenterAppVersions

`func (o *GameCenterDetailRelationships) SetGameCenterAppVersions(v GameCenterAppVersionRelationshipsCompatibilityVersions)`

SetGameCenterAppVersions sets GameCenterAppVersions field to given value.

### HasGameCenterAppVersions

`func (o *GameCenterDetailRelationships) HasGameCenterAppVersions() bool`

HasGameCenterAppVersions returns a boolean if a field has been set.

### GetGameCenterGroup

`func (o *GameCenterDetailRelationships) GetGameCenterGroup() GameCenterAchievementRelationshipsGameCenterGroup`

GetGameCenterGroup returns the GameCenterGroup field if non-nil, zero value otherwise.

### GetGameCenterGroupOk

`func (o *GameCenterDetailRelationships) GetGameCenterGroupOk() (*GameCenterAchievementRelationshipsGameCenterGroup, bool)`

GetGameCenterGroupOk returns a tuple with the GameCenterGroup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameCenterGroup

`func (o *GameCenterDetailRelationships) SetGameCenterGroup(v GameCenterAchievementRelationshipsGameCenterGroup)`

SetGameCenterGroup sets GameCenterGroup field to given value.

### HasGameCenterGroup

`func (o *GameCenterDetailRelationships) HasGameCenterGroup() bool`

HasGameCenterGroup returns a boolean if a field has been set.

### GetGameCenterLeaderboards

`func (o *GameCenterDetailRelationships) GetGameCenterLeaderboards() GameCenterDetailRelationshipsGameCenterLeaderboards`

GetGameCenterLeaderboards returns the GameCenterLeaderboards field if non-nil, zero value otherwise.

### GetGameCenterLeaderboardsOk

`func (o *GameCenterDetailRelationships) GetGameCenterLeaderboardsOk() (*GameCenterDetailRelationshipsGameCenterLeaderboards, bool)`

GetGameCenterLeaderboardsOk returns a tuple with the GameCenterLeaderboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameCenterLeaderboards

`func (o *GameCenterDetailRelationships) SetGameCenterLeaderboards(v GameCenterDetailRelationshipsGameCenterLeaderboards)`

SetGameCenterLeaderboards sets GameCenterLeaderboards field to given value.

### HasGameCenterLeaderboards

`func (o *GameCenterDetailRelationships) HasGameCenterLeaderboards() bool`

HasGameCenterLeaderboards returns a boolean if a field has been set.

### GetGameCenterLeaderboardSets

`func (o *GameCenterDetailRelationships) GetGameCenterLeaderboardSets() GameCenterDetailRelationshipsGameCenterLeaderboardSets`

GetGameCenterLeaderboardSets returns the GameCenterLeaderboardSets field if non-nil, zero value otherwise.

### GetGameCenterLeaderboardSetsOk

`func (o *GameCenterDetailRelationships) GetGameCenterLeaderboardSetsOk() (*GameCenterDetailRelationshipsGameCenterLeaderboardSets, bool)`

GetGameCenterLeaderboardSetsOk returns a tuple with the GameCenterLeaderboardSets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameCenterLeaderboardSets

`func (o *GameCenterDetailRelationships) SetGameCenterLeaderboardSets(v GameCenterDetailRelationshipsGameCenterLeaderboardSets)`

SetGameCenterLeaderboardSets sets GameCenterLeaderboardSets field to given value.

### HasGameCenterLeaderboardSets

`func (o *GameCenterDetailRelationships) HasGameCenterLeaderboardSets() bool`

HasGameCenterLeaderboardSets returns a boolean if a field has been set.

### GetGameCenterAchievements

`func (o *GameCenterDetailRelationships) GetGameCenterAchievements() GameCenterDetailRelationshipsGameCenterAchievements`

GetGameCenterAchievements returns the GameCenterAchievements field if non-nil, zero value otherwise.

### GetGameCenterAchievementsOk

`func (o *GameCenterDetailRelationships) GetGameCenterAchievementsOk() (*GameCenterDetailRelationshipsGameCenterAchievements, bool)`

GetGameCenterAchievementsOk returns a tuple with the GameCenterAchievements field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGameCenterAchievements

`func (o *GameCenterDetailRelationships) SetGameCenterAchievements(v GameCenterDetailRelationshipsGameCenterAchievements)`

SetGameCenterAchievements sets GameCenterAchievements field to given value.

### HasGameCenterAchievements

`func (o *GameCenterDetailRelationships) HasGameCenterAchievements() bool`

HasGameCenterAchievements returns a boolean if a field has been set.

### GetDefaultLeaderboard

`func (o *GameCenterDetailRelationships) GetDefaultLeaderboard() GameCenterDetailRelationshipsDefaultLeaderboard`

GetDefaultLeaderboard returns the DefaultLeaderboard field if non-nil, zero value otherwise.

### GetDefaultLeaderboardOk

`func (o *GameCenterDetailRelationships) GetDefaultLeaderboardOk() (*GameCenterDetailRelationshipsDefaultLeaderboard, bool)`

GetDefaultLeaderboardOk returns a tuple with the DefaultLeaderboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultLeaderboard

`func (o *GameCenterDetailRelationships) SetDefaultLeaderboard(v GameCenterDetailRelationshipsDefaultLeaderboard)`

SetDefaultLeaderboard sets DefaultLeaderboard field to given value.

### HasDefaultLeaderboard

`func (o *GameCenterDetailRelationships) HasDefaultLeaderboard() bool`

HasDefaultLeaderboard returns a boolean if a field has been set.

### GetDefaultGroupLeaderboard

`func (o *GameCenterDetailRelationships) GetDefaultGroupLeaderboard() GameCenterDetailRelationshipsDefaultLeaderboard`

GetDefaultGroupLeaderboard returns the DefaultGroupLeaderboard field if non-nil, zero value otherwise.

### GetDefaultGroupLeaderboardOk

`func (o *GameCenterDetailRelationships) GetDefaultGroupLeaderboardOk() (*GameCenterDetailRelationshipsDefaultLeaderboard, bool)`

GetDefaultGroupLeaderboardOk returns a tuple with the DefaultGroupLeaderboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultGroupLeaderboard

`func (o *GameCenterDetailRelationships) SetDefaultGroupLeaderboard(v GameCenterDetailRelationshipsDefaultLeaderboard)`

SetDefaultGroupLeaderboard sets DefaultGroupLeaderboard field to given value.

### HasDefaultGroupLeaderboard

`func (o *GameCenterDetailRelationships) HasDefaultGroupLeaderboard() bool`

HasDefaultGroupLeaderboard returns a boolean if a field has been set.

### GetAchievementReleases

`func (o *GameCenterDetailRelationships) GetAchievementReleases() GameCenterAchievementRelationshipsReleases`

GetAchievementReleases returns the AchievementReleases field if non-nil, zero value otherwise.

### GetAchievementReleasesOk

`func (o *GameCenterDetailRelationships) GetAchievementReleasesOk() (*GameCenterAchievementRelationshipsReleases, bool)`

GetAchievementReleasesOk returns a tuple with the AchievementReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchievementReleases

`func (o *GameCenterDetailRelationships) SetAchievementReleases(v GameCenterAchievementRelationshipsReleases)`

SetAchievementReleases sets AchievementReleases field to given value.

### HasAchievementReleases

`func (o *GameCenterDetailRelationships) HasAchievementReleases() bool`

HasAchievementReleases returns a boolean if a field has been set.

### GetLeaderboardReleases

`func (o *GameCenterDetailRelationships) GetLeaderboardReleases() GameCenterDetailRelationshipsLeaderboardReleases`

GetLeaderboardReleases returns the LeaderboardReleases field if non-nil, zero value otherwise.

### GetLeaderboardReleasesOk

`func (o *GameCenterDetailRelationships) GetLeaderboardReleasesOk() (*GameCenterDetailRelationshipsLeaderboardReleases, bool)`

GetLeaderboardReleasesOk returns a tuple with the LeaderboardReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderboardReleases

`func (o *GameCenterDetailRelationships) SetLeaderboardReleases(v GameCenterDetailRelationshipsLeaderboardReleases)`

SetLeaderboardReleases sets LeaderboardReleases field to given value.

### HasLeaderboardReleases

`func (o *GameCenterDetailRelationships) HasLeaderboardReleases() bool`

HasLeaderboardReleases returns a boolean if a field has been set.

### GetLeaderboardSetReleases

`func (o *GameCenterDetailRelationships) GetLeaderboardSetReleases() GameCenterDetailRelationshipsLeaderboardSetReleases`

GetLeaderboardSetReleases returns the LeaderboardSetReleases field if non-nil, zero value otherwise.

### GetLeaderboardSetReleasesOk

`func (o *GameCenterDetailRelationships) GetLeaderboardSetReleasesOk() (*GameCenterDetailRelationshipsLeaderboardSetReleases, bool)`

GetLeaderboardSetReleasesOk returns a tuple with the LeaderboardSetReleases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLeaderboardSetReleases

`func (o *GameCenterDetailRelationships) SetLeaderboardSetReleases(v GameCenterDetailRelationshipsLeaderboardSetReleases)`

SetLeaderboardSetReleases sets LeaderboardSetReleases field to given value.

### HasLeaderboardSetReleases

`func (o *GameCenterDetailRelationships) HasLeaderboardSetReleases() bool`

HasLeaderboardSetReleases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


