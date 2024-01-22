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

// checks if the GameCenterAchievementLocalizationCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementLocalizationCreateRequestData{}

// GameCenterAchievementLocalizationCreateRequestData struct for GameCenterAchievementLocalizationCreateRequestData
type GameCenterAchievementLocalizationCreateRequestData struct {
	Type          string                                                          `json:"type"`
	Attributes    GameCenterAchievementLocalizationCreateRequestDataAttributes    `json:"attributes"`
	Relationships GameCenterAchievementLocalizationCreateRequestDataRelationships `json:"relationships"`
}

// NewGameCenterAchievementLocalizationCreateRequestData instantiates a new GameCenterAchievementLocalizationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementLocalizationCreateRequestData(type_ string, attributes GameCenterAchievementLocalizationCreateRequestDataAttributes, relationships GameCenterAchievementLocalizationCreateRequestDataRelationships) *GameCenterAchievementLocalizationCreateRequestData {
	this := GameCenterAchievementLocalizationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewGameCenterAchievementLocalizationCreateRequestDataWithDefaults instantiates a new GameCenterAchievementLocalizationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementLocalizationCreateRequestDataWithDefaults() *GameCenterAchievementLocalizationCreateRequestData {
	this := GameCenterAchievementLocalizationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterAchievementLocalizationCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterAchievementLocalizationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *GameCenterAchievementLocalizationCreateRequestData) GetAttributes() GameCenterAchievementLocalizationCreateRequestDataAttributes {
	if o == nil {
		var ret GameCenterAchievementLocalizationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationCreateRequestData) GetAttributesOk() (*GameCenterAchievementLocalizationCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GameCenterAchievementLocalizationCreateRequestData) SetAttributes(v GameCenterAchievementLocalizationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *GameCenterAchievementLocalizationCreateRequestData) GetRelationships() GameCenterAchievementLocalizationCreateRequestDataRelationships {
	if o == nil {
		var ret GameCenterAchievementLocalizationCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationCreateRequestData) GetRelationshipsOk() (*GameCenterAchievementLocalizationCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *GameCenterAchievementLocalizationCreateRequestData) SetRelationships(v GameCenterAchievementLocalizationCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o GameCenterAchievementLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementLocalizationCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableGameCenterAchievementLocalizationCreateRequestData struct {
	value *GameCenterAchievementLocalizationCreateRequestData
	isSet bool
}

func (v NullableGameCenterAchievementLocalizationCreateRequestData) Get() *GameCenterAchievementLocalizationCreateRequestData {
	return v.value
}

func (v *NullableGameCenterAchievementLocalizationCreateRequestData) Set(val *GameCenterAchievementLocalizationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementLocalizationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementLocalizationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementLocalizationCreateRequestData(val *GameCenterAchievementLocalizationCreateRequestData) *NullableGameCenterAchievementLocalizationCreateRequestData {
	return &NullableGameCenterAchievementLocalizationCreateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterAchievementLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementLocalizationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
