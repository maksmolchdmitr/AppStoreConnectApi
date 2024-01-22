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

// checks if the GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData{}

// GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData struct for GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData
type GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewGameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData instantiates a new GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData(type_ string, id string) *GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData {
	this := GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationDataWithDefaults instantiates a new GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationDataWithDefaults() *GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData {
	this := GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) SetId(v string) {
	o.Id = v
}

func (o GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableGameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData struct {
	value *GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData
	isSet bool
}

func (v NullableGameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) Get() *GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData {
	return v.value
}

func (v *NullableGameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) Set(val *GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData(val *GameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) *NullableGameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData {
	return &NullableGameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardImageRelationshipsGameCenterLeaderboardLocalizationData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
