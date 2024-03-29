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

// checks if the GameCenterAchievementImageUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementImageUpdateRequest{}

// GameCenterAchievementImageUpdateRequest struct for GameCenterAchievementImageUpdateRequest
type GameCenterAchievementImageUpdateRequest struct {
	Data GameCenterAchievementImageUpdateRequestData `json:"data"`
}

// NewGameCenterAchievementImageUpdateRequest instantiates a new GameCenterAchievementImageUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementImageUpdateRequest(data GameCenterAchievementImageUpdateRequestData) *GameCenterAchievementImageUpdateRequest {
	this := GameCenterAchievementImageUpdateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterAchievementImageUpdateRequestWithDefaults instantiates a new GameCenterAchievementImageUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementImageUpdateRequestWithDefaults() *GameCenterAchievementImageUpdateRequest {
	this := GameCenterAchievementImageUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterAchievementImageUpdateRequest) GetData() GameCenterAchievementImageUpdateRequestData {
	if o == nil {
		var ret GameCenterAchievementImageUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementImageUpdateRequest) GetDataOk() (*GameCenterAchievementImageUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterAchievementImageUpdateRequest) SetData(v GameCenterAchievementImageUpdateRequestData) {
	o.Data = v
}

func (o GameCenterAchievementImageUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementImageUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGameCenterAchievementImageUpdateRequest struct {
	value *GameCenterAchievementImageUpdateRequest
	isSet bool
}

func (v NullableGameCenterAchievementImageUpdateRequest) Get() *GameCenterAchievementImageUpdateRequest {
	return v.value
}

func (v *NullableGameCenterAchievementImageUpdateRequest) Set(val *GameCenterAchievementImageUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementImageUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementImageUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementImageUpdateRequest(val *GameCenterAchievementImageUpdateRequest) *NullableGameCenterAchievementImageUpdateRequest {
	return &NullableGameCenterAchievementImageUpdateRequest{value: val, isSet: true}
}

func (v NullableGameCenterAchievementImageUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementImageUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
