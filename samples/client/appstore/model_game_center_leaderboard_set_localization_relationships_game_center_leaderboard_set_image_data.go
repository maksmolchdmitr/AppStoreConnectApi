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

// checks if the GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData{}

// GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData struct for GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData
type GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData instantiates a new GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData(type_ string, id string) *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData {
	this := GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageDataWithDefaults instantiates a new GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageDataWithDefaults() *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData {
	this := GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) SetId(v string) {
	o.Id = v
}

func (o GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData struct {
	value *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData
	isSet bool
}

func (v NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) Get() *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) Set(val *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData(val *GameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) *NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData {
	return &NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetLocalizationRelationshipsGameCenterLeaderboardSetImageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


