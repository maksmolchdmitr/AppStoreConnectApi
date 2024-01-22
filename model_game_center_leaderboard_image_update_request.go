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

// checks if the GameCenterLeaderboardImageUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardImageUpdateRequest{}

// GameCenterLeaderboardImageUpdateRequest struct for GameCenterLeaderboardImageUpdateRequest
type GameCenterLeaderboardImageUpdateRequest struct {
	Data GameCenterLeaderboardImageUpdateRequestData `json:"data"`
}

// NewGameCenterLeaderboardImageUpdateRequest instantiates a new GameCenterLeaderboardImageUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardImageUpdateRequest(data GameCenterLeaderboardImageUpdateRequestData) *GameCenterLeaderboardImageUpdateRequest {
	this := GameCenterLeaderboardImageUpdateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterLeaderboardImageUpdateRequestWithDefaults instantiates a new GameCenterLeaderboardImageUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardImageUpdateRequestWithDefaults() *GameCenterLeaderboardImageUpdateRequest {
	this := GameCenterLeaderboardImageUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardImageUpdateRequest) GetData() GameCenterLeaderboardImageUpdateRequestData {
	if o == nil {
		var ret GameCenterLeaderboardImageUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardImageUpdateRequest) GetDataOk() (*GameCenterLeaderboardImageUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardImageUpdateRequest) SetData(v GameCenterLeaderboardImageUpdateRequestData) {
	o.Data = v
}

func (o GameCenterLeaderboardImageUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardImageUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGameCenterLeaderboardImageUpdateRequest struct {
	value *GameCenterLeaderboardImageUpdateRequest
	isSet bool
}

func (v NullableGameCenterLeaderboardImageUpdateRequest) Get() *GameCenterLeaderboardImageUpdateRequest {
	return v.value
}

func (v *NullableGameCenterLeaderboardImageUpdateRequest) Set(val *GameCenterLeaderboardImageUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardImageUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardImageUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardImageUpdateRequest(val *GameCenterLeaderboardImageUpdateRequest) *NullableGameCenterLeaderboardImageUpdateRequest {
	return &NullableGameCenterLeaderboardImageUpdateRequest{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardImageUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardImageUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
