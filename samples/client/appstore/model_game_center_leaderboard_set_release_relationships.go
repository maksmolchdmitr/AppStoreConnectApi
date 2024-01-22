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

// checks if the GameCenterLeaderboardSetReleaseRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetReleaseRelationships{}

// GameCenterLeaderboardSetReleaseRelationships struct for GameCenterLeaderboardSetReleaseRelationships
type GameCenterLeaderboardSetReleaseRelationships struct {
	GameCenterDetail *AppRelationshipsGameCenterDetail `json:"gameCenterDetail,omitempty"`
	GameCenterLeaderboardSet *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSet `json:"gameCenterLeaderboardSet,omitempty"`
}

// NewGameCenterLeaderboardSetReleaseRelationships instantiates a new GameCenterLeaderboardSetReleaseRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetReleaseRelationships() *GameCenterLeaderboardSetReleaseRelationships {
	this := GameCenterLeaderboardSetReleaseRelationships{}
	return &this
}

// NewGameCenterLeaderboardSetReleaseRelationshipsWithDefaults instantiates a new GameCenterLeaderboardSetReleaseRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetReleaseRelationshipsWithDefaults() *GameCenterLeaderboardSetReleaseRelationships {
	this := GameCenterLeaderboardSetReleaseRelationships{}
	return &this
}

// GetGameCenterDetail returns the GameCenterDetail field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetReleaseRelationships) GetGameCenterDetail() AppRelationshipsGameCenterDetail {
	if o == nil || IsNil(o.GameCenterDetail) {
		var ret AppRelationshipsGameCenterDetail
		return ret
	}
	return *o.GameCenterDetail
}

// GetGameCenterDetailOk returns a tuple with the GameCenterDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetReleaseRelationships) GetGameCenterDetailOk() (*AppRelationshipsGameCenterDetail, bool) {
	if o == nil || IsNil(o.GameCenterDetail) {
		return nil, false
	}
	return o.GameCenterDetail, true
}

// HasGameCenterDetail returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetReleaseRelationships) HasGameCenterDetail() bool {
	if o != nil && !IsNil(o.GameCenterDetail) {
		return true
	}

	return false
}

// SetGameCenterDetail gets a reference to the given AppRelationshipsGameCenterDetail and assigns it to the GameCenterDetail field.
func (o *GameCenterLeaderboardSetReleaseRelationships) SetGameCenterDetail(v AppRelationshipsGameCenterDetail) {
	o.GameCenterDetail = &v
}

// GetGameCenterLeaderboardSet returns the GameCenterLeaderboardSet field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetReleaseRelationships) GetGameCenterLeaderboardSet() GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSet {
	if o == nil || IsNil(o.GameCenterLeaderboardSet) {
		var ret GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSet
		return ret
	}
	return *o.GameCenterLeaderboardSet
}

// GetGameCenterLeaderboardSetOk returns a tuple with the GameCenterLeaderboardSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetReleaseRelationships) GetGameCenterLeaderboardSetOk() (*GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSet, bool) {
	if o == nil || IsNil(o.GameCenterLeaderboardSet) {
		return nil, false
	}
	return o.GameCenterLeaderboardSet, true
}

// HasGameCenterLeaderboardSet returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetReleaseRelationships) HasGameCenterLeaderboardSet() bool {
	if o != nil && !IsNil(o.GameCenterLeaderboardSet) {
		return true
	}

	return false
}

// SetGameCenterLeaderboardSet gets a reference to the given GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSet and assigns it to the GameCenterLeaderboardSet field.
func (o *GameCenterLeaderboardSetReleaseRelationships) SetGameCenterLeaderboardSet(v GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSet) {
	o.GameCenterLeaderboardSet = &v
}

func (o GameCenterLeaderboardSetReleaseRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetReleaseRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GameCenterDetail) {
		toSerialize["gameCenterDetail"] = o.GameCenterDetail
	}
	if !IsNil(o.GameCenterLeaderboardSet) {
		toSerialize["gameCenterLeaderboardSet"] = o.GameCenterLeaderboardSet
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardSetReleaseRelationships struct {
	value *GameCenterLeaderboardSetReleaseRelationships
	isSet bool
}

func (v NullableGameCenterLeaderboardSetReleaseRelationships) Get() *GameCenterLeaderboardSetReleaseRelationships {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetReleaseRelationships) Set(val *GameCenterLeaderboardSetReleaseRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetReleaseRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetReleaseRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetReleaseRelationships(val *GameCenterLeaderboardSetReleaseRelationships) *NullableGameCenterLeaderboardSetReleaseRelationships {
	return &NullableGameCenterLeaderboardSetReleaseRelationships{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetReleaseRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetReleaseRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


