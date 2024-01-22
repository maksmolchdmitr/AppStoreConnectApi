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

// checks if the GameCenterMatchmakingRuleSetCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingRuleSetCreateRequestData{}

// GameCenterMatchmakingRuleSetCreateRequestData struct for GameCenterMatchmakingRuleSetCreateRequestData
type GameCenterMatchmakingRuleSetCreateRequestData struct {
	Type       string                                                  `json:"type"`
	Attributes GameCenterMatchmakingRuleSetCreateRequestDataAttributes `json:"attributes"`
}

// NewGameCenterMatchmakingRuleSetCreateRequestData instantiates a new GameCenterMatchmakingRuleSetCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingRuleSetCreateRequestData(type_ string, attributes GameCenterMatchmakingRuleSetCreateRequestDataAttributes) *GameCenterMatchmakingRuleSetCreateRequestData {
	this := GameCenterMatchmakingRuleSetCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewGameCenterMatchmakingRuleSetCreateRequestDataWithDefaults instantiates a new GameCenterMatchmakingRuleSetCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingRuleSetCreateRequestDataWithDefaults() *GameCenterMatchmakingRuleSetCreateRequestData {
	this := GameCenterMatchmakingRuleSetCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterMatchmakingRuleSetCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterMatchmakingRuleSetCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *GameCenterMatchmakingRuleSetCreateRequestData) GetAttributes() GameCenterMatchmakingRuleSetCreateRequestDataAttributes {
	if o == nil {
		var ret GameCenterMatchmakingRuleSetCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetCreateRequestData) GetAttributesOk() (*GameCenterMatchmakingRuleSetCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GameCenterMatchmakingRuleSetCreateRequestData) SetAttributes(v GameCenterMatchmakingRuleSetCreateRequestDataAttributes) {
	o.Attributes = v
}

func (o GameCenterMatchmakingRuleSetCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingRuleSetCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

type NullableGameCenterMatchmakingRuleSetCreateRequestData struct {
	value *GameCenterMatchmakingRuleSetCreateRequestData
	isSet bool
}

func (v NullableGameCenterMatchmakingRuleSetCreateRequestData) Get() *GameCenterMatchmakingRuleSetCreateRequestData {
	return v.value
}

func (v *NullableGameCenterMatchmakingRuleSetCreateRequestData) Set(val *GameCenterMatchmakingRuleSetCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingRuleSetCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingRuleSetCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingRuleSetCreateRequestData(val *GameCenterMatchmakingRuleSetCreateRequestData) *NullableGameCenterMatchmakingRuleSetCreateRequestData {
	return &NullableGameCenterMatchmakingRuleSetCreateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingRuleSetCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingRuleSetCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
