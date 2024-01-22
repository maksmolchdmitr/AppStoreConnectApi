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

// checks if the GameCenterAchievementCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementCreateRequestDataRelationships{}

// GameCenterAchievementCreateRequestDataRelationships struct for GameCenterAchievementCreateRequestDataRelationships
type GameCenterAchievementCreateRequestDataRelationships struct {
	GameCenterDetail *GameCenterAchievementCreateRequestDataRelationshipsGameCenterDetail `json:"gameCenterDetail,omitempty"`
	GameCenterGroup *GameCenterAchievementCreateRequestDataRelationshipsGameCenterGroup `json:"gameCenterGroup,omitempty"`
}

// NewGameCenterAchievementCreateRequestDataRelationships instantiates a new GameCenterAchievementCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementCreateRequestDataRelationships() *GameCenterAchievementCreateRequestDataRelationships {
	this := GameCenterAchievementCreateRequestDataRelationships{}
	return &this
}

// NewGameCenterAchievementCreateRequestDataRelationshipsWithDefaults instantiates a new GameCenterAchievementCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementCreateRequestDataRelationshipsWithDefaults() *GameCenterAchievementCreateRequestDataRelationships {
	this := GameCenterAchievementCreateRequestDataRelationships{}
	return &this
}

// GetGameCenterDetail returns the GameCenterDetail field value if set, zero value otherwise.
func (o *GameCenterAchievementCreateRequestDataRelationships) GetGameCenterDetail() GameCenterAchievementCreateRequestDataRelationshipsGameCenterDetail {
	if o == nil || IsNil(o.GameCenterDetail) {
		var ret GameCenterAchievementCreateRequestDataRelationshipsGameCenterDetail
		return ret
	}
	return *o.GameCenterDetail
}

// GetGameCenterDetailOk returns a tuple with the GameCenterDetail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementCreateRequestDataRelationships) GetGameCenterDetailOk() (*GameCenterAchievementCreateRequestDataRelationshipsGameCenterDetail, bool) {
	if o == nil || IsNil(o.GameCenterDetail) {
		return nil, false
	}
	return o.GameCenterDetail, true
}

// HasGameCenterDetail returns a boolean if a field has been set.
func (o *GameCenterAchievementCreateRequestDataRelationships) HasGameCenterDetail() bool {
	if o != nil && !IsNil(o.GameCenterDetail) {
		return true
	}

	return false
}

// SetGameCenterDetail gets a reference to the given GameCenterAchievementCreateRequestDataRelationshipsGameCenterDetail and assigns it to the GameCenterDetail field.
func (o *GameCenterAchievementCreateRequestDataRelationships) SetGameCenterDetail(v GameCenterAchievementCreateRequestDataRelationshipsGameCenterDetail) {
	o.GameCenterDetail = &v
}

// GetGameCenterGroup returns the GameCenterGroup field value if set, zero value otherwise.
func (o *GameCenterAchievementCreateRequestDataRelationships) GetGameCenterGroup() GameCenterAchievementCreateRequestDataRelationshipsGameCenterGroup {
	if o == nil || IsNil(o.GameCenterGroup) {
		var ret GameCenterAchievementCreateRequestDataRelationshipsGameCenterGroup
		return ret
	}
	return *o.GameCenterGroup
}

// GetGameCenterGroupOk returns a tuple with the GameCenterGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementCreateRequestDataRelationships) GetGameCenterGroupOk() (*GameCenterAchievementCreateRequestDataRelationshipsGameCenterGroup, bool) {
	if o == nil || IsNil(o.GameCenterGroup) {
		return nil, false
	}
	return o.GameCenterGroup, true
}

// HasGameCenterGroup returns a boolean if a field has been set.
func (o *GameCenterAchievementCreateRequestDataRelationships) HasGameCenterGroup() bool {
	if o != nil && !IsNil(o.GameCenterGroup) {
		return true
	}

	return false
}

// SetGameCenterGroup gets a reference to the given GameCenterAchievementCreateRequestDataRelationshipsGameCenterGroup and assigns it to the GameCenterGroup field.
func (o *GameCenterAchievementCreateRequestDataRelationships) SetGameCenterGroup(v GameCenterAchievementCreateRequestDataRelationshipsGameCenterGroup) {
	o.GameCenterGroup = &v
}

func (o GameCenterAchievementCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GameCenterDetail) {
		toSerialize["gameCenterDetail"] = o.GameCenterDetail
	}
	if !IsNil(o.GameCenterGroup) {
		toSerialize["gameCenterGroup"] = o.GameCenterGroup
	}
	return toSerialize, nil
}

type NullableGameCenterAchievementCreateRequestDataRelationships struct {
	value *GameCenterAchievementCreateRequestDataRelationships
	isSet bool
}

func (v NullableGameCenterAchievementCreateRequestDataRelationships) Get() *GameCenterAchievementCreateRequestDataRelationships {
	return v.value
}

func (v *NullableGameCenterAchievementCreateRequestDataRelationships) Set(val *GameCenterAchievementCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementCreateRequestDataRelationships(val *GameCenterAchievementCreateRequestDataRelationships) *NullableGameCenterAchievementCreateRequestDataRelationships {
	return &NullableGameCenterAchievementCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableGameCenterAchievementCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


