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

// checks if the GameCenterLeaderboardSetUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetUpdateRequest{}

// GameCenterLeaderboardSetUpdateRequest struct for GameCenterLeaderboardSetUpdateRequest
type GameCenterLeaderboardSetUpdateRequest struct {
	Data GameCenterLeaderboardSetUpdateRequestData `json:"data"`
}

// NewGameCenterLeaderboardSetUpdateRequest instantiates a new GameCenterLeaderboardSetUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetUpdateRequest(data GameCenterLeaderboardSetUpdateRequestData) *GameCenterLeaderboardSetUpdateRequest {
	this := GameCenterLeaderboardSetUpdateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterLeaderboardSetUpdateRequestWithDefaults instantiates a new GameCenterLeaderboardSetUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetUpdateRequestWithDefaults() *GameCenterLeaderboardSetUpdateRequest {
	this := GameCenterLeaderboardSetUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardSetUpdateRequest) GetData() GameCenterLeaderboardSetUpdateRequestData {
	if o == nil {
		var ret GameCenterLeaderboardSetUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetUpdateRequest) GetDataOk() (*GameCenterLeaderboardSetUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardSetUpdateRequest) SetData(v GameCenterLeaderboardSetUpdateRequestData) {
	o.Data = v
}

func (o GameCenterLeaderboardSetUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGameCenterLeaderboardSetUpdateRequest struct {
	value *GameCenterLeaderboardSetUpdateRequest
	isSet bool
}

func (v NullableGameCenterLeaderboardSetUpdateRequest) Get() *GameCenterLeaderboardSetUpdateRequest {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetUpdateRequest) Set(val *GameCenterLeaderboardSetUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetUpdateRequest(val *GameCenterLeaderboardSetUpdateRequest) *NullableGameCenterLeaderboardSetUpdateRequest {
	return &NullableGameCenterLeaderboardSetUpdateRequest{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


