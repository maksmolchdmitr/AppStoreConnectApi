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

// checks if the GameCenterAchievementReleaseCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementReleaseCreateRequest{}

// GameCenterAchievementReleaseCreateRequest struct for GameCenterAchievementReleaseCreateRequest
type GameCenterAchievementReleaseCreateRequest struct {
	Data GameCenterAchievementReleaseCreateRequestData `json:"data"`
}

// NewGameCenterAchievementReleaseCreateRequest instantiates a new GameCenterAchievementReleaseCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementReleaseCreateRequest(data GameCenterAchievementReleaseCreateRequestData) *GameCenterAchievementReleaseCreateRequest {
	this := GameCenterAchievementReleaseCreateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterAchievementReleaseCreateRequestWithDefaults instantiates a new GameCenterAchievementReleaseCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementReleaseCreateRequestWithDefaults() *GameCenterAchievementReleaseCreateRequest {
	this := GameCenterAchievementReleaseCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterAchievementReleaseCreateRequest) GetData() GameCenterAchievementReleaseCreateRequestData {
	if o == nil {
		var ret GameCenterAchievementReleaseCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementReleaseCreateRequest) GetDataOk() (*GameCenterAchievementReleaseCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterAchievementReleaseCreateRequest) SetData(v GameCenterAchievementReleaseCreateRequestData) {
	o.Data = v
}

func (o GameCenterAchievementReleaseCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementReleaseCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGameCenterAchievementReleaseCreateRequest struct {
	value *GameCenterAchievementReleaseCreateRequest
	isSet bool
}

func (v NullableGameCenterAchievementReleaseCreateRequest) Get() *GameCenterAchievementReleaseCreateRequest {
	return v.value
}

func (v *NullableGameCenterAchievementReleaseCreateRequest) Set(val *GameCenterAchievementReleaseCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementReleaseCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementReleaseCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementReleaseCreateRequest(val *GameCenterAchievementReleaseCreateRequest) *NullableGameCenterAchievementReleaseCreateRequest {
	return &NullableGameCenterAchievementReleaseCreateRequest{value: val, isSet: true}
}

func (v NullableGameCenterAchievementReleaseCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementReleaseCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


