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

// checks if the GameCenterAchievementLocalizationRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementLocalizationRelationships{}

// GameCenterAchievementLocalizationRelationships struct for GameCenterAchievementLocalizationRelationships
type GameCenterAchievementLocalizationRelationships struct {
	GameCenterAchievement *GameCenterAchievementLocalizationRelationshipsGameCenterAchievement `json:"gameCenterAchievement,omitempty"`
	GameCenterAchievementImage *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage `json:"gameCenterAchievementImage,omitempty"`
}

// NewGameCenterAchievementLocalizationRelationships instantiates a new GameCenterAchievementLocalizationRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementLocalizationRelationships() *GameCenterAchievementLocalizationRelationships {
	this := GameCenterAchievementLocalizationRelationships{}
	return &this
}

// NewGameCenterAchievementLocalizationRelationshipsWithDefaults instantiates a new GameCenterAchievementLocalizationRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementLocalizationRelationshipsWithDefaults() *GameCenterAchievementLocalizationRelationships {
	this := GameCenterAchievementLocalizationRelationships{}
	return &this
}

// GetGameCenterAchievement returns the GameCenterAchievement field value if set, zero value otherwise.
func (o *GameCenterAchievementLocalizationRelationships) GetGameCenterAchievement() GameCenterAchievementLocalizationRelationshipsGameCenterAchievement {
	if o == nil || IsNil(o.GameCenterAchievement) {
		var ret GameCenterAchievementLocalizationRelationshipsGameCenterAchievement
		return ret
	}
	return *o.GameCenterAchievement
}

// GetGameCenterAchievementOk returns a tuple with the GameCenterAchievement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationRelationships) GetGameCenterAchievementOk() (*GameCenterAchievementLocalizationRelationshipsGameCenterAchievement, bool) {
	if o == nil || IsNil(o.GameCenterAchievement) {
		return nil, false
	}
	return o.GameCenterAchievement, true
}

// HasGameCenterAchievement returns a boolean if a field has been set.
func (o *GameCenterAchievementLocalizationRelationships) HasGameCenterAchievement() bool {
	if o != nil && !IsNil(o.GameCenterAchievement) {
		return true
	}

	return false
}

// SetGameCenterAchievement gets a reference to the given GameCenterAchievementLocalizationRelationshipsGameCenterAchievement and assigns it to the GameCenterAchievement field.
func (o *GameCenterAchievementLocalizationRelationships) SetGameCenterAchievement(v GameCenterAchievementLocalizationRelationshipsGameCenterAchievement) {
	o.GameCenterAchievement = &v
}

// GetGameCenterAchievementImage returns the GameCenterAchievementImage field value if set, zero value otherwise.
func (o *GameCenterAchievementLocalizationRelationships) GetGameCenterAchievementImage() GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage {
	if o == nil || IsNil(o.GameCenterAchievementImage) {
		var ret GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage
		return ret
	}
	return *o.GameCenterAchievementImage
}

// GetGameCenterAchievementImageOk returns a tuple with the GameCenterAchievementImage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationRelationships) GetGameCenterAchievementImageOk() (*GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage, bool) {
	if o == nil || IsNil(o.GameCenterAchievementImage) {
		return nil, false
	}
	return o.GameCenterAchievementImage, true
}

// HasGameCenterAchievementImage returns a boolean if a field has been set.
func (o *GameCenterAchievementLocalizationRelationships) HasGameCenterAchievementImage() bool {
	if o != nil && !IsNil(o.GameCenterAchievementImage) {
		return true
	}

	return false
}

// SetGameCenterAchievementImage gets a reference to the given GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage and assigns it to the GameCenterAchievementImage field.
func (o *GameCenterAchievementLocalizationRelationships) SetGameCenterAchievementImage(v GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) {
	o.GameCenterAchievementImage = &v
}

func (o GameCenterAchievementLocalizationRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementLocalizationRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GameCenterAchievement) {
		toSerialize["gameCenterAchievement"] = o.GameCenterAchievement
	}
	if !IsNil(o.GameCenterAchievementImage) {
		toSerialize["gameCenterAchievementImage"] = o.GameCenterAchievementImage
	}
	return toSerialize, nil
}

type NullableGameCenterAchievementLocalizationRelationships struct {
	value *GameCenterAchievementLocalizationRelationships
	isSet bool
}

func (v NullableGameCenterAchievementLocalizationRelationships) Get() *GameCenterAchievementLocalizationRelationships {
	return v.value
}

func (v *NullableGameCenterAchievementLocalizationRelationships) Set(val *GameCenterAchievementLocalizationRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementLocalizationRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementLocalizationRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementLocalizationRelationships(val *GameCenterAchievementLocalizationRelationships) *NullableGameCenterAchievementLocalizationRelationships {
	return &NullableGameCenterAchievementLocalizationRelationships{value: val, isSet: true}
}

func (v NullableGameCenterAchievementLocalizationRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementLocalizationRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


