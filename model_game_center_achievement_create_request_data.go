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

// checks if the GameCenterAchievementCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementCreateRequestData{}

// GameCenterAchievementCreateRequestData struct for GameCenterAchievementCreateRequestData
type GameCenterAchievementCreateRequestData struct {
	Type          string                                               `json:"type"`
	Attributes    GameCenterAchievementCreateRequestDataAttributes     `json:"attributes"`
	Relationships *GameCenterAchievementCreateRequestDataRelationships `json:"relationships,omitempty"`
}

// NewGameCenterAchievementCreateRequestData instantiates a new GameCenterAchievementCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementCreateRequestData(type_ string, attributes GameCenterAchievementCreateRequestDataAttributes) *GameCenterAchievementCreateRequestData {
	this := GameCenterAchievementCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewGameCenterAchievementCreateRequestDataWithDefaults instantiates a new GameCenterAchievementCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementCreateRequestDataWithDefaults() *GameCenterAchievementCreateRequestData {
	this := GameCenterAchievementCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterAchievementCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterAchievementCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *GameCenterAchievementCreateRequestData) GetAttributes() GameCenterAchievementCreateRequestDataAttributes {
	if o == nil {
		var ret GameCenterAchievementCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementCreateRequestData) GetAttributesOk() (*GameCenterAchievementCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GameCenterAchievementCreateRequestData) SetAttributes(v GameCenterAchievementCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GameCenterAchievementCreateRequestData) GetRelationships() GameCenterAchievementCreateRequestDataRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GameCenterAchievementCreateRequestDataRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementCreateRequestData) GetRelationshipsOk() (*GameCenterAchievementCreateRequestDataRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GameCenterAchievementCreateRequestData) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GameCenterAchievementCreateRequestDataRelationships and assigns it to the Relationships field.
func (o *GameCenterAchievementCreateRequestData) SetRelationships(v GameCenterAchievementCreateRequestDataRelationships) {
	o.Relationships = &v
}

func (o GameCenterAchievementCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	return toSerialize, nil
}

type NullableGameCenterAchievementCreateRequestData struct {
	value *GameCenterAchievementCreateRequestData
	isSet bool
}

func (v NullableGameCenterAchievementCreateRequestData) Get() *GameCenterAchievementCreateRequestData {
	return v.value
}

func (v *NullableGameCenterAchievementCreateRequestData) Set(val *GameCenterAchievementCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementCreateRequestData(val *GameCenterAchievementCreateRequestData) *NullableGameCenterAchievementCreateRequestData {
	return &NullableGameCenterAchievementCreateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterAchievementCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
