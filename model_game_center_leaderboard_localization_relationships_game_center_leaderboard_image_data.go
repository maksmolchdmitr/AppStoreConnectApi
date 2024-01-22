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

// checks if the GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData{}

// GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData struct for GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData
type GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData instantiates a new GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData(type_ string, id string) *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData {
	this := GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageDataWithDefaults instantiates a new GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageDataWithDefaults() *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData {
	this := GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) SetId(v string) {
	o.Id = v
}

func (o GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData struct {
	value *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData
	isSet bool
}

func (v NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) Get() *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData {
	return v.value
}

func (v *NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) Set(val *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData(val *GameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) *NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData {
	return &NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardLocalizationRelationshipsGameCenterLeaderboardImageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
