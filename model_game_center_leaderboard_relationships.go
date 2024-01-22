/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GameCenterLeaderboardRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardRelationships{}

// GameCenterLeaderboardRelationships struct for GameCenterLeaderboardRelationships
type GameCenterLeaderboardRelationships struct {
	GameCenterDetail          *AppRelationshipsGameCenterDetail                       `json:"gameCenterDetail,omitempty"`
	GameCenterGroup           *GameCenterAchievementRelationshipsGameCenterGroup      `json:"gameCenterGroup,omitempty"`
	GroupLeaderboard          *GameCenterDetailRelationshipsDefaultLeaderboard        `json:"groupLeaderboard,omitempty"`
	GameCenterLeaderboardSets *GameCenterDetailRelationshipsGameCenterLeaderboardSets `json:"gameCenterLeaderboardSets,omitempty"`
	Localizations             *GameCenterLeaderboardRelationshipsLocalizations        `json:"localizations,omitempty"`
	Releases                  *GameCenterDetailRelationshipsLeaderboardReleases       `json:"releases,omitempty"`
}

// NewGameCenterLeaderboardRelationships instantiates a new GameCenterLeaderboardRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardRelationships() *GameCenterLeaderboardRelationships {
	this := GameCenterLeaderboardRelationships{}
	return &this
}

// NewGameCenterLeaderboardRelationshipsWithDefaults instantiates a new GameCenterLeaderboardRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardRelationshipsWithDefaults() *GameCenterLeaderboardRelationships {
	this := GameCenterLeaderboardRelationships{}
	return &this
}

// GetGameCenterDetail returns the GameCenterDetail field value if set, zero value otherwise.
func (o *GameCenterLeaderboardRelationships) GetGameCenterDetail() AppRelationshipsGameCenterDetail {
	if o == nil || IsNil(o.GameCenterDetail) {
		var ret AppRelationshipsGameCenterDetail
		return ret
	}
	return *o.GameCenterDetail
}

// GetGameCenterDetailOk returns a tuple with the GameCenterDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardRelationships) GetGameCenterDetailOk() (*AppRelationshipsGameCenterDetail, bool) {
	if o == nil || IsNil(o.GameCenterDetail) {
		return nil, false
	}
	return o.GameCenterDetail, true
}

// HasGameCenterDetail returns a boolean if a field has been set.
func (o *GameCenterLeaderboardRelationships) HasGameCenterDetail() bool {
	if o != nil && !IsNil(o.GameCenterDetail) {
		return true
	}

	return false
}

// SetGameCenterDetail gets a reference to the given AppRelationshipsGameCenterDetail and assigns it to the GameCenterDetail field.
func (o *GameCenterLeaderboardRelationships) SetGameCenterDetail(v AppRelationshipsGameCenterDetail) {
	o.GameCenterDetail = &v
}

// GetGameCenterGroup returns the GameCenterGroup field value if set, zero value otherwise.
func (o *GameCenterLeaderboardRelationships) GetGameCenterGroup() GameCenterAchievementRelationshipsGameCenterGroup {
	if o == nil || IsNil(o.GameCenterGroup) {
		var ret GameCenterAchievementRelationshipsGameCenterGroup
		return ret
	}
	return *o.GameCenterGroup
}

// GetGameCenterGroupOk returns a tuple with the GameCenterGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardRelationships) GetGameCenterGroupOk() (*GameCenterAchievementRelationshipsGameCenterGroup, bool) {
	if o == nil || IsNil(o.GameCenterGroup) {
		return nil, false
	}
	return o.GameCenterGroup, true
}

// HasGameCenterGroup returns a boolean if a field has been set.
func (o *GameCenterLeaderboardRelationships) HasGameCenterGroup() bool {
	if o != nil && !IsNil(o.GameCenterGroup) {
		return true
	}

	return false
}

// SetGameCenterGroup gets a reference to the given GameCenterAchievementRelationshipsGameCenterGroup and assigns it to the GameCenterGroup field.
func (o *GameCenterLeaderboardRelationships) SetGameCenterGroup(v GameCenterAchievementRelationshipsGameCenterGroup) {
	o.GameCenterGroup = &v
}

// GetGroupLeaderboard returns the GroupLeaderboard field value if set, zero value otherwise.
func (o *GameCenterLeaderboardRelationships) GetGroupLeaderboard() GameCenterDetailRelationshipsDefaultLeaderboard {
	if o == nil || IsNil(o.GroupLeaderboard) {
		var ret GameCenterDetailRelationshipsDefaultLeaderboard
		return ret
	}
	return *o.GroupLeaderboard
}

// GetGroupLeaderboardOk returns a tuple with the GroupLeaderboard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardRelationships) GetGroupLeaderboardOk() (*GameCenterDetailRelationshipsDefaultLeaderboard, bool) {
	if o == nil || IsNil(o.GroupLeaderboard) {
		return nil, false
	}
	return o.GroupLeaderboard, true
}

// HasGroupLeaderboard returns a boolean if a field has been set.
func (o *GameCenterLeaderboardRelationships) HasGroupLeaderboard() bool {
	if o != nil && !IsNil(o.GroupLeaderboard) {
		return true
	}

	return false
}

// SetGroupLeaderboard gets a reference to the given GameCenterDetailRelationshipsDefaultLeaderboard and assigns it to the GroupLeaderboard field.
func (o *GameCenterLeaderboardRelationships) SetGroupLeaderboard(v GameCenterDetailRelationshipsDefaultLeaderboard) {
	o.GroupLeaderboard = &v
}

// GetGameCenterLeaderboardSets returns the GameCenterLeaderboardSets field value if set, zero value otherwise.
func (o *GameCenterLeaderboardRelationships) GetGameCenterLeaderboardSets() GameCenterDetailRelationshipsGameCenterLeaderboardSets {
	if o == nil || IsNil(o.GameCenterLeaderboardSets) {
		var ret GameCenterDetailRelationshipsGameCenterLeaderboardSets
		return ret
	}
	return *o.GameCenterLeaderboardSets
}

// GetGameCenterLeaderboardSetsOk returns a tuple with the GameCenterLeaderboardSets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardRelationships) GetGameCenterLeaderboardSetsOk() (*GameCenterDetailRelationshipsGameCenterLeaderboardSets, bool) {
	if o == nil || IsNil(o.GameCenterLeaderboardSets) {
		return nil, false
	}
	return o.GameCenterLeaderboardSets, true
}

// HasGameCenterLeaderboardSets returns a boolean if a field has been set.
func (o *GameCenterLeaderboardRelationships) HasGameCenterLeaderboardSets() bool {
	if o != nil && !IsNil(o.GameCenterLeaderboardSets) {
		return true
	}

	return false
}

// SetGameCenterLeaderboardSets gets a reference to the given GameCenterDetailRelationshipsGameCenterLeaderboardSets and assigns it to the GameCenterLeaderboardSets field.
func (o *GameCenterLeaderboardRelationships) SetGameCenterLeaderboardSets(v GameCenterDetailRelationshipsGameCenterLeaderboardSets) {
	o.GameCenterLeaderboardSets = &v
}

// GetLocalizations returns the Localizations field value if set, zero value otherwise.
func (o *GameCenterLeaderboardRelationships) GetLocalizations() GameCenterLeaderboardRelationshipsLocalizations {
	if o == nil || IsNil(o.Localizations) {
		var ret GameCenterLeaderboardRelationshipsLocalizations
		return ret
	}
	return *o.Localizations
}

// GetLocalizationsOk returns a tuple with the Localizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardRelationships) GetLocalizationsOk() (*GameCenterLeaderboardRelationshipsLocalizations, bool) {
	if o == nil || IsNil(o.Localizations) {
		return nil, false
	}
	return o.Localizations, true
}

// HasLocalizations returns a boolean if a field has been set.
func (o *GameCenterLeaderboardRelationships) HasLocalizations() bool {
	if o != nil && !IsNil(o.Localizations) {
		return true
	}

	return false
}

// SetLocalizations gets a reference to the given GameCenterLeaderboardRelationshipsLocalizations and assigns it to the Localizations field.
func (o *GameCenterLeaderboardRelationships) SetLocalizations(v GameCenterLeaderboardRelationshipsLocalizations) {
	o.Localizations = &v
}

// GetReleases returns the Releases field value if set, zero value otherwise.
func (o *GameCenterLeaderboardRelationships) GetReleases() GameCenterDetailRelationshipsLeaderboardReleases {
	if o == nil || IsNil(o.Releases) {
		var ret GameCenterDetailRelationshipsLeaderboardReleases
		return ret
	}
	return *o.Releases
}

// GetReleasesOk returns a tuple with the Releases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardRelationships) GetReleasesOk() (*GameCenterDetailRelationshipsLeaderboardReleases, bool) {
	if o == nil || IsNil(o.Releases) {
		return nil, false
	}
	return o.Releases, true
}

// HasReleases returns a boolean if a field has been set.
func (o *GameCenterLeaderboardRelationships) HasReleases() bool {
	if o != nil && !IsNil(o.Releases) {
		return true
	}

	return false
}

// SetReleases gets a reference to the given GameCenterDetailRelationshipsLeaderboardReleases and assigns it to the Releases field.
func (o *GameCenterLeaderboardRelationships) SetReleases(v GameCenterDetailRelationshipsLeaderboardReleases) {
	o.Releases = &v
}

func (o GameCenterLeaderboardRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GameCenterDetail) {
		toSerialize["gameCenterDetail"] = o.GameCenterDetail
	}
	if !IsNil(o.GameCenterGroup) {
		toSerialize["gameCenterGroup"] = o.GameCenterGroup
	}
	if !IsNil(o.GroupLeaderboard) {
		toSerialize["groupLeaderboard"] = o.GroupLeaderboard
	}
	if !IsNil(o.GameCenterLeaderboardSets) {
		toSerialize["gameCenterLeaderboardSets"] = o.GameCenterLeaderboardSets
	}
	if !IsNil(o.Localizations) {
		toSerialize["localizations"] = o.Localizations
	}
	if !IsNil(o.Releases) {
		toSerialize["releases"] = o.Releases
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardRelationships struct {
	value *GameCenterLeaderboardRelationships
	isSet bool
}

func (v NullableGameCenterLeaderboardRelationships) Get() *GameCenterLeaderboardRelationships {
	return v.value
}

func (v *NullableGameCenterLeaderboardRelationships) Set(val *GameCenterLeaderboardRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardRelationships(val *GameCenterLeaderboardRelationships) *NullableGameCenterLeaderboardRelationships {
	return &NullableGameCenterLeaderboardRelationships{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
