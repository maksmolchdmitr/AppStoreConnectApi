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

// checks if the GameCenterLeaderboardGroupLeaderboardLinkageRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardGroupLeaderboardLinkageRequest{}

// GameCenterLeaderboardGroupLeaderboardLinkageRequest struct for GameCenterLeaderboardGroupLeaderboardLinkageRequest
type GameCenterLeaderboardGroupLeaderboardLinkageRequest struct {
	Data GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner `json:"data"`
}

// NewGameCenterLeaderboardGroupLeaderboardLinkageRequest instantiates a new GameCenterLeaderboardGroupLeaderboardLinkageRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardGroupLeaderboardLinkageRequest(data GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) *GameCenterLeaderboardGroupLeaderboardLinkageRequest {
	this := GameCenterLeaderboardGroupLeaderboardLinkageRequest{}
	this.Data = data
	return &this
}

// NewGameCenterLeaderboardGroupLeaderboardLinkageRequestWithDefaults instantiates a new GameCenterLeaderboardGroupLeaderboardLinkageRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardGroupLeaderboardLinkageRequestWithDefaults() *GameCenterLeaderboardGroupLeaderboardLinkageRequest {
	this := GameCenterLeaderboardGroupLeaderboardLinkageRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterLeaderboardGroupLeaderboardLinkageRequest) GetData() GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner {
	if o == nil {
		var ret GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardGroupLeaderboardLinkageRequest) GetDataOk() (*GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterLeaderboardGroupLeaderboardLinkageRequest) SetData(v GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) {
	o.Data = v
}

func (o GameCenterLeaderboardGroupLeaderboardLinkageRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardGroupLeaderboardLinkageRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGameCenterLeaderboardGroupLeaderboardLinkageRequest struct {
	value *GameCenterLeaderboardGroupLeaderboardLinkageRequest
	isSet bool
}

func (v NullableGameCenterLeaderboardGroupLeaderboardLinkageRequest) Get() *GameCenterLeaderboardGroupLeaderboardLinkageRequest {
	return v.value
}

func (v *NullableGameCenterLeaderboardGroupLeaderboardLinkageRequest) Set(val *GameCenterLeaderboardGroupLeaderboardLinkageRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardGroupLeaderboardLinkageRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardGroupLeaderboardLinkageRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardGroupLeaderboardLinkageRequest(val *GameCenterLeaderboardGroupLeaderboardLinkageRequest) *NullableGameCenterLeaderboardGroupLeaderboardLinkageRequest {
	return &NullableGameCenterLeaderboardGroupLeaderboardLinkageRequest{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardGroupLeaderboardLinkageRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardGroupLeaderboardLinkageRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

