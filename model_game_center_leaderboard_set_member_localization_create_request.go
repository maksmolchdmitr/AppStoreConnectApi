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

// checks if the GameCenterLeaderboardSetMemberLocalizationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetMemberLocalizationCreateRequest{}

// GameCenterLeaderboardSetMemberLocalizationCreateRequest struct for GameCenterLeaderboardSetMemberLocalizationCreateRequest
type GameCenterLeaderboardSetMemberLocalizationCreateRequest struct {
	Data GameCenterLeaderboardSetMemberLocalizationCreateRequestData `json:"data"`
}

// NewGameCenterLeaderboardSetMemberLocalizationCreateRequest instantiates a new GameCenterLeaderboardSetMemberLocalizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetMemberLocalizationCreateRequest(data GameCenterLeaderboardSetMemberLocalizationCreateRequestData) *GameCenterLeaderboardSetMemberLocalizationCreateRequest {
	this := GameCenterLeaderboardSetMemberLocalizationCreateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterLeaderboardSetMemberLocalizationCreateRequestWithDefaults instantiates a new GameCenterLeaderboardSetMemberLocalizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetMemberLocalizationCreateRequestWithDefaults() *GameCenterLeaderboardSetMemberLocalizationCreateRequest {
	this := GameCenterLeaderboardSetMemberLocalizationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardSetMemberLocalizationCreateRequest) GetData() GameCenterLeaderboardSetMemberLocalizationCreateRequestData {
	if o == nil {
		var ret GameCenterLeaderboardSetMemberLocalizationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetMemberLocalizationCreateRequest) GetDataOk() (*GameCenterLeaderboardSetMemberLocalizationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardSetMemberLocalizationCreateRequest) SetData(v GameCenterLeaderboardSetMemberLocalizationCreateRequestData) {
	o.Data = v
}

func (o GameCenterLeaderboardSetMemberLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetMemberLocalizationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGameCenterLeaderboardSetMemberLocalizationCreateRequest struct {
	value *GameCenterLeaderboardSetMemberLocalizationCreateRequest
	isSet bool
}

func (v NullableGameCenterLeaderboardSetMemberLocalizationCreateRequest) Get() *GameCenterLeaderboardSetMemberLocalizationCreateRequest {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetMemberLocalizationCreateRequest) Set(val *GameCenterLeaderboardSetMemberLocalizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetMemberLocalizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetMemberLocalizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetMemberLocalizationCreateRequest(val *GameCenterLeaderboardSetMemberLocalizationCreateRequest) *NullableGameCenterLeaderboardSetMemberLocalizationCreateRequest {
	return &NullableGameCenterLeaderboardSetMemberLocalizationCreateRequest{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetMemberLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetMemberLocalizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
