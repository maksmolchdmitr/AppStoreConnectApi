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

// checks if the GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard{}

// GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard struct for GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard
type GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard struct {
	Data GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner `json:"data"`
}

// NewGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard instantiates a new GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard(data GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard {
	this := GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard{}
	this.Data = data
	return &this
}

// NewGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboardWithDefaults instantiates a new GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboardWithDefaults() *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard {
	this := GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) GetData() GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner {
	if o == nil {
		var ret GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) GetDataOk() (*GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) SetData(v GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) {
	o.Data = v
}

func (o GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard struct {
	value *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard
	isSet bool
}

func (v NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) Get() *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard {
	return v.value
}

func (v *NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) Set(val *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard(val *GameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) *NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard {
	return &NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardLocalizationCreateRequestDataRelationshipsGameCenterLeaderboard) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
