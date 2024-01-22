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

// checks if the GameCenterAchievementLocalizationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementLocalizationCreateRequest{}

// GameCenterAchievementLocalizationCreateRequest struct for GameCenterAchievementLocalizationCreateRequest
type GameCenterAchievementLocalizationCreateRequest struct {
	Data GameCenterAchievementLocalizationCreateRequestData `json:"data"`
}

// NewGameCenterAchievementLocalizationCreateRequest instantiates a new GameCenterAchievementLocalizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementLocalizationCreateRequest(data GameCenterAchievementLocalizationCreateRequestData) *GameCenterAchievementLocalizationCreateRequest {
	this := GameCenterAchievementLocalizationCreateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterAchievementLocalizationCreateRequestWithDefaults instantiates a new GameCenterAchievementLocalizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementLocalizationCreateRequestWithDefaults() *GameCenterAchievementLocalizationCreateRequest {
	this := GameCenterAchievementLocalizationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterAchievementLocalizationCreateRequest) GetData() GameCenterAchievementLocalizationCreateRequestData {
	if o == nil {
		var ret GameCenterAchievementLocalizationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationCreateRequest) GetDataOk() (*GameCenterAchievementLocalizationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterAchievementLocalizationCreateRequest) SetData(v GameCenterAchievementLocalizationCreateRequestData) {
	o.Data = v
}

func (o GameCenterAchievementLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementLocalizationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGameCenterAchievementLocalizationCreateRequest struct {
	value *GameCenterAchievementLocalizationCreateRequest
	isSet bool
}

func (v NullableGameCenterAchievementLocalizationCreateRequest) Get() *GameCenterAchievementLocalizationCreateRequest {
	return v.value
}

func (v *NullableGameCenterAchievementLocalizationCreateRequest) Set(val *GameCenterAchievementLocalizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementLocalizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementLocalizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementLocalizationCreateRequest(val *GameCenterAchievementLocalizationCreateRequest) *NullableGameCenterAchievementLocalizationCreateRequest {
	return &NullableGameCenterAchievementLocalizationCreateRequest{value: val, isSet: true}
}

func (v NullableGameCenterAchievementLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementLocalizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
