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

// checks if the GameCenterGroupGameCenterAchievementsLinkagesRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterGroupGameCenterAchievementsLinkagesRequest{}

// GameCenterGroupGameCenterAchievementsLinkagesRequest struct for GameCenterGroupGameCenterAchievementsLinkagesRequest
type GameCenterGroupGameCenterAchievementsLinkagesRequest struct {
	Data []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData `json:"data"`
}

// NewGameCenterGroupGameCenterAchievementsLinkagesRequest instantiates a new GameCenterGroupGameCenterAchievementsLinkagesRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterGroupGameCenterAchievementsLinkagesRequest(data []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData) *GameCenterGroupGameCenterAchievementsLinkagesRequest {
	this := GameCenterGroupGameCenterAchievementsLinkagesRequest{}
	this.Data = data
	return &this
}

// NewGameCenterGroupGameCenterAchievementsLinkagesRequestWithDefaults instantiates a new GameCenterGroupGameCenterAchievementsLinkagesRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterGroupGameCenterAchievementsLinkagesRequestWithDefaults() *GameCenterGroupGameCenterAchievementsLinkagesRequest {
	this := GameCenterGroupGameCenterAchievementsLinkagesRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterGroupGameCenterAchievementsLinkagesRequest) GetData() []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData {
	if o == nil {
		var ret []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterGroupGameCenterAchievementsLinkagesRequest) GetDataOk() ([]GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GameCenterGroupGameCenterAchievementsLinkagesRequest) SetData(v []GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData) {
	o.Data = v
}

func (o GameCenterGroupGameCenterAchievementsLinkagesRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterGroupGameCenterAchievementsLinkagesRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGameCenterGroupGameCenterAchievementsLinkagesRequest struct {
	value *GameCenterGroupGameCenterAchievementsLinkagesRequest
	isSet bool
}

func (v NullableGameCenterGroupGameCenterAchievementsLinkagesRequest) Get() *GameCenterGroupGameCenterAchievementsLinkagesRequest {
	return v.value
}

func (v *NullableGameCenterGroupGameCenterAchievementsLinkagesRequest) Set(val *GameCenterGroupGameCenterAchievementsLinkagesRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterGroupGameCenterAchievementsLinkagesRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterGroupGameCenterAchievementsLinkagesRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterGroupGameCenterAchievementsLinkagesRequest(val *GameCenterGroupGameCenterAchievementsLinkagesRequest) *NullableGameCenterGroupGameCenterAchievementsLinkagesRequest {
	return &NullableGameCenterGroupGameCenterAchievementsLinkagesRequest{value: val, isSet: true}
}

func (v NullableGameCenterGroupGameCenterAchievementsLinkagesRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterGroupGameCenterAchievementsLinkagesRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
