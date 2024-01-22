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

// checks if the GameCenterGroupCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterGroupCreateRequestData{}

// GameCenterGroupCreateRequestData struct for GameCenterGroupCreateRequestData
type GameCenterGroupCreateRequestData struct {
	Type       string                     `json:"type"`
	Attributes *GameCenterGroupAttributes `json:"attributes,omitempty"`
}

// NewGameCenterGroupCreateRequestData instantiates a new GameCenterGroupCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterGroupCreateRequestData(type_ string) *GameCenterGroupCreateRequestData {
	this := GameCenterGroupCreateRequestData{}
	this.Type = type_
	return &this
}

// NewGameCenterGroupCreateRequestDataWithDefaults instantiates a new GameCenterGroupCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterGroupCreateRequestDataWithDefaults() *GameCenterGroupCreateRequestData {
	this := GameCenterGroupCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterGroupCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterGroupCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterGroupCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterGroupCreateRequestData) GetAttributes() GameCenterGroupAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterGroupAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterGroupCreateRequestData) GetAttributesOk() (*GameCenterGroupAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterGroupCreateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterGroupAttributes and assigns it to the Attributes field.
func (o *GameCenterGroupCreateRequestData) SetAttributes(v GameCenterGroupAttributes) {
	o.Attributes = &v
}

func (o GameCenterGroupCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterGroupCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableGameCenterGroupCreateRequestData struct {
	value *GameCenterGroupCreateRequestData
	isSet bool
}

func (v NullableGameCenterGroupCreateRequestData) Get() *GameCenterGroupCreateRequestData {
	return v.value
}

func (v *NullableGameCenterGroupCreateRequestData) Set(val *GameCenterGroupCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterGroupCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterGroupCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterGroupCreateRequestData(val *GameCenterGroupCreateRequestData) *NullableGameCenterGroupCreateRequestData {
	return &NullableGameCenterGroupCreateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterGroupCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterGroupCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
