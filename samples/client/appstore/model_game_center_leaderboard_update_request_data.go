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

// checks if the GameCenterLeaderboardUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardUpdateRequestData{}

// GameCenterLeaderboardUpdateRequestData struct for GameCenterLeaderboardUpdateRequestData
type GameCenterLeaderboardUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *GameCenterLeaderboardUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewGameCenterLeaderboardUpdateRequestData instantiates a new GameCenterLeaderboardUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardUpdateRequestData(type_ string, id string) *GameCenterLeaderboardUpdateRequestData {
	this := GameCenterLeaderboardUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterLeaderboardUpdateRequestDataWithDefaults instantiates a new GameCenterLeaderboardUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardUpdateRequestDataWithDefaults() *GameCenterLeaderboardUpdateRequestData {
	this := GameCenterLeaderboardUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterLeaderboardUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterLeaderboardUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterLeaderboardUpdateRequestData) GetAttributes() GameCenterLeaderboardUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterLeaderboardUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardUpdateRequestData) GetAttributesOk() (*GameCenterLeaderboardUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterLeaderboardUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterLeaderboardUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *GameCenterLeaderboardUpdateRequestData) SetAttributes(v GameCenterLeaderboardUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o GameCenterLeaderboardUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardUpdateRequestData struct {
	value *GameCenterLeaderboardUpdateRequestData
	isSet bool
}

func (v NullableGameCenterLeaderboardUpdateRequestData) Get() *GameCenterLeaderboardUpdateRequestData {
	return v.value
}

func (v *NullableGameCenterLeaderboardUpdateRequestData) Set(val *GameCenterLeaderboardUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardUpdateRequestData(val *GameCenterLeaderboardUpdateRequestData) *NullableGameCenterLeaderboardUpdateRequestData {
	return &NullableGameCenterLeaderboardUpdateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


