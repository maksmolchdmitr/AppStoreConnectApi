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

// checks if the GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes{}

// GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes struct for GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes
type GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes struct {
	PlayerId string `json:"playerId"`
	Properties []Property `json:"properties,omitempty"`
}

// NewGameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes instantiates a new GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes(playerId string) *GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes {
	this := GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes{}
	this.PlayerId = playerId
	return &this
}

// NewGameCenterMatchmakingTestPlayerPropertyInlineCreateAttributesWithDefaults instantiates a new GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingTestPlayerPropertyInlineCreateAttributesWithDefaults() *GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes {
	this := GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes{}
	return &this
}

// GetPlayerId returns the PlayerId field value
func (o *GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) GetPlayerId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PlayerId
}

// GetPlayerIdOk returns a tuple with the PlayerId field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) GetPlayerIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PlayerId, true
}

// SetPlayerId sets field value
func (o *GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) SetPlayerId(v string) {
	o.PlayerId = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) GetProperties() []Property {
	if o == nil || IsNil(o.Properties) {
		var ret []Property
		return ret
	}
	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) GetPropertiesOk() ([]Property, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given []Property and assigns it to the Properties field.
func (o *GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) SetProperties(v []Property) {
	o.Properties = v
}

func (o GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["playerId"] = o.PlayerId
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes struct {
	value *GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes
	isSet bool
}

func (v NullableGameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) Get() *GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes {
	return v.value
}

func (v *NullableGameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) Set(val *GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes(val *GameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) *NullableGameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes {
	return &NullableGameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingTestPlayerPropertyInlineCreateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


