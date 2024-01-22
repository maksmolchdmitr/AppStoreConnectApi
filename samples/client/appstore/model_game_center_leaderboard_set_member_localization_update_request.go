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

// checks if the GameCenterLeaderboardSetMemberLocalizationUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetMemberLocalizationUpdateRequest{}

// GameCenterLeaderboardSetMemberLocalizationUpdateRequest struct for GameCenterLeaderboardSetMemberLocalizationUpdateRequest
type GameCenterLeaderboardSetMemberLocalizationUpdateRequest struct {
	Data GameCenterLeaderboardSetMemberLocalizationUpdateRequestData `json:"data"`
}

// NewGameCenterLeaderboardSetMemberLocalizationUpdateRequest instantiates a new GameCenterLeaderboardSetMemberLocalizationUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetMemberLocalizationUpdateRequest(data GameCenterLeaderboardSetMemberLocalizationUpdateRequestData) *GameCenterLeaderboardSetMemberLocalizationUpdateRequest {
	this := GameCenterLeaderboardSetMemberLocalizationUpdateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterLeaderboardSetMemberLocalizationUpdateRequestWithDefaults instantiates a new GameCenterLeaderboardSetMemberLocalizationUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetMemberLocalizationUpdateRequestWithDefaults() *GameCenterLeaderboardSetMemberLocalizationUpdateRequest {
	this := GameCenterLeaderboardSetMemberLocalizationUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardSetMemberLocalizationUpdateRequest) GetData() GameCenterLeaderboardSetMemberLocalizationUpdateRequestData {
	if o == nil {
		var ret GameCenterLeaderboardSetMemberLocalizationUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetMemberLocalizationUpdateRequest) GetDataOk() (*GameCenterLeaderboardSetMemberLocalizationUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardSetMemberLocalizationUpdateRequest) SetData(v GameCenterLeaderboardSetMemberLocalizationUpdateRequestData) {
	o.Data = v
}

func (o GameCenterLeaderboardSetMemberLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetMemberLocalizationUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGameCenterLeaderboardSetMemberLocalizationUpdateRequest struct {
	value *GameCenterLeaderboardSetMemberLocalizationUpdateRequest
	isSet bool
}

func (v NullableGameCenterLeaderboardSetMemberLocalizationUpdateRequest) Get() *GameCenterLeaderboardSetMemberLocalizationUpdateRequest {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetMemberLocalizationUpdateRequest) Set(val *GameCenterLeaderboardSetMemberLocalizationUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetMemberLocalizationUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetMemberLocalizationUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetMemberLocalizationUpdateRequest(val *GameCenterLeaderboardSetMemberLocalizationUpdateRequest) *NullableGameCenterLeaderboardSetMemberLocalizationUpdateRequest {
	return &NullableGameCenterLeaderboardSetMemberLocalizationUpdateRequest{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetMemberLocalizationUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetMemberLocalizationUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

