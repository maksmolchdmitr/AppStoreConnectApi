# GameCenterMatchmakingTestRequestInlineCreateAttributes

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestName** | **string** |  | 
**SecondsInQueue** | **int32** |  | 
**Locale** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**MinPlayers** | Pointer to **int32** |  | [optional] 
**MaxPlayers** | Pointer to **int32** |  | [optional] 
**PlayerCount** | Pointer to **int32** |  | [optional] 
**BundleId** | **string** |  | 
**Platform** | [**Platform**](Platform.md) |  | 
**AppVersion** | **string** |  | 

## Methods

### NewGameCenterMatchmakingTestRequestInlineCreateAttributes

`func NewGameCenterMatchmakingTestRequestInlineCreateAttributes(requestName string, secondsInQueue int32, bundleId string, platform Platform, appVersion string, ) *GameCenterMatchmakingTestRequestInlineCreateAttributes`

NewGameCenterMatchmakingTestRequestInlineCreateAttributes instantiates a new GameCenterMatchmakingTestRequestInlineCreateAttributes object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGameCenterMatchmakingTestRequestInlineCreateAttributesWithDefaults

`func NewGameCenterMatchmakingTestRequestInlineCreateAttributesWithDefaults() *GameCenterMatchmakingTestRequestInlineCreateAttributes`

NewGameCenterMatchmakingTestRequestInlineCreateAttributesWithDefaults instantiates a new GameCenterMatchmakingTestRequestInlineCreateAttributes object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestName

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetRequestName() string`

GetRequestName returns the RequestName field if non-nil, zero value otherwise.

### GetRequestNameOk

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetRequestNameOk() (*string, bool)`

GetRequestNameOk returns a tuple with the RequestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestName

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetRequestName(v string)`

SetRequestName sets RequestName field to given value.


### GetSecondsInQueue

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetSecondsInQueue() int32`

GetSecondsInQueue returns the SecondsInQueue field if non-nil, zero value otherwise.

### GetSecondsInQueueOk

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetSecondsInQueueOk() (*int32, bool)`

GetSecondsInQueueOk returns a tuple with the SecondsInQueue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecondsInQueue

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetSecondsInQueue(v int32)`

SetSecondsInQueue sets SecondsInQueue field to given value.


### GetLocale

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetLocale() string`

GetLocale returns the Locale field if non-nil, zero value otherwise.

### GetLocaleOk

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetLocaleOk() (*string, bool)`

GetLocaleOk returns a tuple with the Locale field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocale

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetLocale(v string)`

SetLocale sets Locale field to given value.

### HasLocale

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) HasLocale() bool`

HasLocale returns a boolean if a field has been set.

### GetLocation

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMinPlayers

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetMinPlayers() int32`

GetMinPlayers returns the MinPlayers field if non-nil, zero value otherwise.

### GetMinPlayersOk

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetMinPlayersOk() (*int32, bool)`

GetMinPlayersOk returns a tuple with the MinPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinPlayers

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetMinPlayers(v int32)`

SetMinPlayers sets MinPlayers field to given value.

### HasMinPlayers

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) HasMinPlayers() bool`

HasMinPlayers returns a boolean if a field has been set.

### GetMaxPlayers

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetMaxPlayers() int32`

GetMaxPlayers returns the MaxPlayers field if non-nil, zero value otherwise.

### GetMaxPlayersOk

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetMaxPlayersOk() (*int32, bool)`

GetMaxPlayersOk returns a tuple with the MaxPlayers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxPlayers

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetMaxPlayers(v int32)`

SetMaxPlayers sets MaxPlayers field to given value.

### HasMaxPlayers

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) HasMaxPlayers() bool`

HasMaxPlayers returns a boolean if a field has been set.

### GetPlayerCount

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetPlayerCount() int32`

GetPlayerCount returns the PlayerCount field if non-nil, zero value otherwise.

### GetPlayerCountOk

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetPlayerCountOk() (*int32, bool)`

GetPlayerCountOk returns a tuple with the PlayerCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlayerCount

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetPlayerCount(v int32)`

SetPlayerCount sets PlayerCount field to given value.

### HasPlayerCount

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) HasPlayerCount() bool`

HasPlayerCount returns a boolean if a field has been set.

### GetBundleId

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetBundleId() string`

GetBundleId returns the BundleId field if non-nil, zero value otherwise.

### GetBundleIdOk

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetBundleIdOk() (*string, bool)`

GetBundleIdOk returns a tuple with the BundleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBundleId

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetBundleId(v string)`

SetBundleId sets BundleId field to given value.


### GetPlatform

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetPlatform() Platform`

GetPlatform returns the Platform field if non-nil, zero value otherwise.

### GetPlatformOk

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetPlatformOk() (*Platform, bool)`

GetPlatformOk returns a tuple with the Platform field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlatform

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetPlatform(v Platform)`

SetPlatform sets Platform field to given value.


### GetAppVersion

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetAppVersion() string`

GetAppVersion returns the AppVersion field if non-nil, zero value otherwise.

### GetAppVersionOk

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) GetAppVersionOk() (*string, bool)`

GetAppVersionOk returns a tuple with the AppVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppVersion

`func (o *GameCenterMatchmakingTestRequestInlineCreateAttributes) SetAppVersion(v string)`

SetAppVersion sets AppVersion field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


