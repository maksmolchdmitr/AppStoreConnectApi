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

// checks if the GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner{}

// GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner struct for GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner
type GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewGameCenterMatchmakingRuleSetRelationshipsTeamsDataInner instantiates a new GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingRuleSetRelationshipsTeamsDataInner(type_ string, id string) *GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner {
	this := GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterMatchmakingRuleSetRelationshipsTeamsDataInnerWithDefaults instantiates a new GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingRuleSetRelationshipsTeamsDataInnerWithDefaults() *GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner {
	this := GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) SetId(v string) {
	o.Id = v
}

func (o GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableGameCenterMatchmakingRuleSetRelationshipsTeamsDataInner struct {
	value *GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner
	isSet bool
}

func (v NullableGameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) Get() *GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner {
	return v.value
}

func (v *NullableGameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) Set(val *GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingRuleSetRelationshipsTeamsDataInner(val *GameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) *NullableGameCenterMatchmakingRuleSetRelationshipsTeamsDataInner {
	return &NullableGameCenterMatchmakingRuleSetRelationshipsTeamsDataInner{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingRuleSetRelationshipsTeamsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

