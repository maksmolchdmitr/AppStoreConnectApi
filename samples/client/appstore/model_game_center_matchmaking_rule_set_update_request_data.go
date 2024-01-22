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

// checks if the GameCenterMatchmakingRuleSetUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingRuleSetUpdateRequestData{}

// GameCenterMatchmakingRuleSetUpdateRequestData struct for GameCenterMatchmakingRuleSetUpdateRequestData
type GameCenterMatchmakingRuleSetUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *GameCenterMatchmakingRuleSetUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewGameCenterMatchmakingRuleSetUpdateRequestData instantiates a new GameCenterMatchmakingRuleSetUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingRuleSetUpdateRequestData(type_ string, id string) *GameCenterMatchmakingRuleSetUpdateRequestData {
	this := GameCenterMatchmakingRuleSetUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterMatchmakingRuleSetUpdateRequestDataWithDefaults instantiates a new GameCenterMatchmakingRuleSetUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingRuleSetUpdateRequestDataWithDefaults() *GameCenterMatchmakingRuleSetUpdateRequestData {
	this := GameCenterMatchmakingRuleSetUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterMatchmakingRuleSetUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterMatchmakingRuleSetUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterMatchmakingRuleSetUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterMatchmakingRuleSetUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterMatchmakingRuleSetUpdateRequestData) GetAttributes() GameCenterMatchmakingRuleSetUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterMatchmakingRuleSetUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetUpdateRequestData) GetAttributesOk() (*GameCenterMatchmakingRuleSetUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterMatchmakingRuleSetUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterMatchmakingRuleSetUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *GameCenterMatchmakingRuleSetUpdateRequestData) SetAttributes(v GameCenterMatchmakingRuleSetUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o GameCenterMatchmakingRuleSetUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingRuleSetUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingRuleSetUpdateRequestData struct {
	value *GameCenterMatchmakingRuleSetUpdateRequestData
	isSet bool
}

func (v NullableGameCenterMatchmakingRuleSetUpdateRequestData) Get() *GameCenterMatchmakingRuleSetUpdateRequestData {
	return v.value
}

func (v *NullableGameCenterMatchmakingRuleSetUpdateRequestData) Set(val *GameCenterMatchmakingRuleSetUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingRuleSetUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingRuleSetUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingRuleSetUpdateRequestData(val *GameCenterMatchmakingRuleSetUpdateRequestData) *NullableGameCenterMatchmakingRuleSetUpdateRequestData {
	return &NullableGameCenterMatchmakingRuleSetUpdateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingRuleSetUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingRuleSetUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

