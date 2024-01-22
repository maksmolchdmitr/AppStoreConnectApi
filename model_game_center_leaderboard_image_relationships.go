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

// checks if the GameCenterLeaderboardImageRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardImageRelationships{}

// GameCenterLeaderboardImageRelationships struct for GameCenterLeaderboardImageRelationships
type GameCenterLeaderboardImageRelationships struct {
	GameCenterLeaderboardLocalization *GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalization `json:"gameCenterLeaderboardLocalization,omitempty"`
}

// NewGameCenterLeaderboardImageRelationships instantiates a new GameCenterLeaderboardImageRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardImageRelationships() *GameCenterLeaderboardImageRelationships {
	this := GameCenterLeaderboardImageRelationships{}
	return &this
}

// NewGameCenterLeaderboardImageRelationshipsWithDefaults instantiates a new GameCenterLeaderboardImageRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardImageRelationshipsWithDefaults() *GameCenterLeaderboardImageRelationships {
	this := GameCenterLeaderboardImageRelationships{}
	return &this
}

// GetGameCenterLeaderboardLocalization returns the GameCenterLeaderboardLocalization field value if set, zero value otherwise.
func (o *GameCenterLeaderboardImageRelationships) GetGameCenterLeaderboardLocalization() GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalization {
	if o == nil || IsNil(o.GameCenterLeaderboardLocalization) {
		var ret GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalization
		return ret
	}
	return *o.GameCenterLeaderboardLocalization
}

// GetGameCenterLeaderboardLocalizationOk returns a tuple with the GameCenterLeaderboardLocalization field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardImageRelationships) GetGameCenterLeaderboardLocalizationOk() (*GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalization, bool) {
	if o == nil || IsNil(o.GameCenterLeaderboardLocalization) {
		return nil, false
	}
	return o.GameCenterLeaderboardLocalization, true
}

// HasGameCenterLeaderboardLocalization returns a boolean if a field has been set.
func (o *GameCenterLeaderboardImageRelationships) HasGameCenterLeaderboardLocalization() bool {
	if o != nil && !IsNil(o.GameCenterLeaderboardLocalization) {
		return true
	}

	return false
}

// SetGameCenterLeaderboardLocalization gets a reference to the given GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalization and assigns it to the GameCenterLeaderboardLocalization field.
func (o *GameCenterLeaderboardImageRelationships) SetGameCenterLeaderboardLocalization(v GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalization) {
	o.GameCenterLeaderboardLocalization = &v
}

func (o GameCenterLeaderboardImageRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardImageRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GameCenterLeaderboardLocalization) {
		toSerialize["gameCenterLeaderboardLocalization"] = o.GameCenterLeaderboardLocalization
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardImageRelationships struct {
	value *GameCenterLeaderboardImageRelationships
	isSet bool
}

func (v NullableGameCenterLeaderboardImageRelationships) Get() *GameCenterLeaderboardImageRelationships {
	return v.value
}

func (v *NullableGameCenterLeaderboardImageRelationships) Set(val *GameCenterLeaderboardImageRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardImageRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardImageRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardImageRelationships(val *GameCenterLeaderboardImageRelationships) *NullableGameCenterLeaderboardImageRelationships {
	return &NullableGameCenterLeaderboardImageRelationships{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardImageRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardImageRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
