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

// checks if the GameCenterLeaderboardLocalizationUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardLocalizationUpdateRequest{}

// GameCenterLeaderboardLocalizationUpdateRequest struct for GameCenterLeaderboardLocalizationUpdateRequest
type GameCenterLeaderboardLocalizationUpdateRequest struct {
	Data GameCenterLeaderboardLocalizationUpdateRequestData `json:"data"`
}

// NewGameCenterLeaderboardLocalizationUpdateRequest instantiates a new GameCenterLeaderboardLocalizationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardLocalizationUpdateRequest(data GameCenterLeaderboardLocalizationUpdateRequestData) *GameCenterLeaderboardLocalizationUpdateRequest {
	this := GameCenterLeaderboardLocalizationUpdateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterLeaderboardLocalizationUpdateRequestWithDefaults instantiates a new GameCenterLeaderboardLocalizationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardLocalizationUpdateRequestWithDefaults() *GameCenterLeaderboardLocalizationUpdateRequest {
	this := GameCenterLeaderboardLocalizationUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardLocalizationUpdateRequest) GetData() GameCenterLeaderboardLocalizationUpdateRequestData {
	if o == nil {
		var ret GameCenterLeaderboardLocalizationUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationUpdateRequest) GetDataOk() (*GameCenterLeaderboardLocalizationUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardLocalizationUpdateRequest) SetData(v GameCenterLeaderboardLocalizationUpdateRequestData) {
	o.Data = v
}

func (o GameCenterLeaderboardLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardLocalizationUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGameCenterLeaderboardLocalizationUpdateRequest struct {
	value *GameCenterLeaderboardLocalizationUpdateRequest
	isSet bool
}

func (v NullableGameCenterLeaderboardLocalizationUpdateRequest) Get() *GameCenterLeaderboardLocalizationUpdateRequest {
	return v.value
}

func (v *NullableGameCenterLeaderboardLocalizationUpdateRequest) Set(val *GameCenterLeaderboardLocalizationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardLocalizationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardLocalizationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardLocalizationUpdateRequest(val *GameCenterLeaderboardLocalizationUpdateRequest) *NullableGameCenterLeaderboardLocalizationUpdateRequest {
	return &NullableGameCenterLeaderboardLocalizationUpdateRequest{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardLocalizationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
