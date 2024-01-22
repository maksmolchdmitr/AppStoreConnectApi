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

// checks if the GameCenterAchievementLocalizationUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementLocalizationUpdateRequestData{}

// GameCenterAchievementLocalizationUpdateRequestData struct for GameCenterAchievementLocalizationUpdateRequestData
type GameCenterAchievementLocalizationUpdateRequestData struct {
	Type       string                                                        `json:"type"`
	Id         string                                                        `json:"id"`
	Attributes *GameCenterAchievementLocalizationUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewGameCenterAchievementLocalizationUpdateRequestData instantiates a new GameCenterAchievementLocalizationUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementLocalizationUpdateRequestData(type_ string, id string) *GameCenterAchievementLocalizationUpdateRequestData {
	this := GameCenterAchievementLocalizationUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterAchievementLocalizationUpdateRequestDataWithDefaults instantiates a new GameCenterAchievementLocalizationUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementLocalizationUpdateRequestDataWithDefaults() *GameCenterAchievementLocalizationUpdateRequestData {
	this := GameCenterAchievementLocalizationUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterAchievementLocalizationUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterAchievementLocalizationUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterAchievementLocalizationUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterAchievementLocalizationUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterAchievementLocalizationUpdateRequestData) GetAttributes() GameCenterAchievementLocalizationUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterAchievementLocalizationUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationUpdateRequestData) GetAttributesOk() (*GameCenterAchievementLocalizationUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterAchievementLocalizationUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterAchievementLocalizationUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *GameCenterAchievementLocalizationUpdateRequestData) SetAttributes(v GameCenterAchievementLocalizationUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o GameCenterAchievementLocalizationUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementLocalizationUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableGameCenterAchievementLocalizationUpdateRequestData struct {
	value *GameCenterAchievementLocalizationUpdateRequestData
	isSet bool
}

func (v NullableGameCenterAchievementLocalizationUpdateRequestData) Get() *GameCenterAchievementLocalizationUpdateRequestData {
	return v.value
}

func (v *NullableGameCenterAchievementLocalizationUpdateRequestData) Set(val *GameCenterAchievementLocalizationUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementLocalizationUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementLocalizationUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementLocalizationUpdateRequestData(val *GameCenterAchievementLocalizationUpdateRequestData) *NullableGameCenterAchievementLocalizationUpdateRequestData {
	return &NullableGameCenterAchievementLocalizationUpdateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterAchievementLocalizationUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementLocalizationUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
