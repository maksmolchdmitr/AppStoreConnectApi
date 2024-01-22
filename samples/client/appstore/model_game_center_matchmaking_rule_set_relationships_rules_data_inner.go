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

// checks if the GameCenterMatchmakingRuleSetRelationshipsRulesDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingRuleSetRelationshipsRulesDataInner{}

// GameCenterMatchmakingRuleSetRelationshipsRulesDataInner struct for GameCenterMatchmakingRuleSetRelationshipsRulesDataInner
type GameCenterMatchmakingRuleSetRelationshipsRulesDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewGameCenterMatchmakingRuleSetRelationshipsRulesDataInner instantiates a new GameCenterMatchmakingRuleSetRelationshipsRulesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingRuleSetRelationshipsRulesDataInner(type_ string, id string) *GameCenterMatchmakingRuleSetRelationshipsRulesDataInner {
	this := GameCenterMatchmakingRuleSetRelationshipsRulesDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterMatchmakingRuleSetRelationshipsRulesDataInnerWithDefaults instantiates a new GameCenterMatchmakingRuleSetRelationshipsRulesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingRuleSetRelationshipsRulesDataInnerWithDefaults() *GameCenterMatchmakingRuleSetRelationshipsRulesDataInner {
	this := GameCenterMatchmakingRuleSetRelationshipsRulesDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterMatchmakingRuleSetRelationshipsRulesDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetRelationshipsRulesDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterMatchmakingRuleSetRelationshipsRulesDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterMatchmakingRuleSetRelationshipsRulesDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetRelationshipsRulesDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterMatchmakingRuleSetRelationshipsRulesDataInner) SetId(v string) {
	o.Id = v
}

func (o GameCenterMatchmakingRuleSetRelationshipsRulesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingRuleSetRelationshipsRulesDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableGameCenterMatchmakingRuleSetRelationshipsRulesDataInner struct {
	value *GameCenterMatchmakingRuleSetRelationshipsRulesDataInner
	isSet bool
}

func (v NullableGameCenterMatchmakingRuleSetRelationshipsRulesDataInner) Get() *GameCenterMatchmakingRuleSetRelationshipsRulesDataInner {
	return v.value
}

func (v *NullableGameCenterMatchmakingRuleSetRelationshipsRulesDataInner) Set(val *GameCenterMatchmakingRuleSetRelationshipsRulesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingRuleSetRelationshipsRulesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingRuleSetRelationshipsRulesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingRuleSetRelationshipsRulesDataInner(val *GameCenterMatchmakingRuleSetRelationshipsRulesDataInner) *NullableGameCenterMatchmakingRuleSetRelationshipsRulesDataInner {
	return &NullableGameCenterMatchmakingRuleSetRelationshipsRulesDataInner{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingRuleSetRelationshipsRulesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingRuleSetRelationshipsRulesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


