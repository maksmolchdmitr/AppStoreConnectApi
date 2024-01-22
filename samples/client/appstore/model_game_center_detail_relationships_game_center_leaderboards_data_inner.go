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

// checks if the GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner{}

// GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner struct for GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner
type GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner instantiates a new GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner(type_ string, id string) *GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner {
	this := GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterDetailRelationshipsGameCenterLeaderboardsDataInnerWithDefaults instantiates a new GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterDetailRelationshipsGameCenterLeaderboardsDataInnerWithDefaults() *GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner {
	this := GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) SetId(v string) {
	o.Id = v
}

func (o GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner struct {
	value *GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner
	isSet bool
}

func (v NullableGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) Get() *GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner {
	return v.value
}

func (v *NullableGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) Set(val *GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner(val *GameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) *NullableGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner {
	return &NullableGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner{value: val, isSet: true}
}

func (v NullableGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterDetailRelationshipsGameCenterLeaderboardsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


