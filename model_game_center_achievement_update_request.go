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

// checks if the GameCenterAchievementUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementUpdateRequest{}

// GameCenterAchievementUpdateRequest struct for GameCenterAchievementUpdateRequest
type GameCenterAchievementUpdateRequest struct {
	Data GameCenterAchievementUpdateRequestData `json:"data"`
}

// NewGameCenterAchievementUpdateRequest instantiates a new GameCenterAchievementUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementUpdateRequest(data GameCenterAchievementUpdateRequestData) *GameCenterAchievementUpdateRequest {
	this := GameCenterAchievementUpdateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterAchievementUpdateRequestWithDefaults instantiates a new GameCenterAchievementUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementUpdateRequestWithDefaults() *GameCenterAchievementUpdateRequest {
	this := GameCenterAchievementUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterAchievementUpdateRequest) GetData() GameCenterAchievementUpdateRequestData {
	if o == nil {
		var ret GameCenterAchievementUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementUpdateRequest) GetDataOk() (*GameCenterAchievementUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterAchievementUpdateRequest) SetData(v GameCenterAchievementUpdateRequestData) {
	o.Data = v
}

func (o GameCenterAchievementUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGameCenterAchievementUpdateRequest struct {
	value *GameCenterAchievementUpdateRequest
	isSet bool
}

func (v NullableGameCenterAchievementUpdateRequest) Get() *GameCenterAchievementUpdateRequest {
	return v.value
}

func (v *NullableGameCenterAchievementUpdateRequest) Set(val *GameCenterAchievementUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementUpdateRequest(val *GameCenterAchievementUpdateRequest) *NullableGameCenterAchievementUpdateRequest {
	return &NullableGameCenterAchievementUpdateRequest{value: val, isSet: true}
}

func (v NullableGameCenterAchievementUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}