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

// checks if the GameCenterMatchmakingRuleCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingRuleCreateRequestData{}

// GameCenterMatchmakingRuleCreateRequestData struct for GameCenterMatchmakingRuleCreateRequestData
type GameCenterMatchmakingRuleCreateRequestData struct {
	Type          string                                                  `json:"type"`
	Attributes    GameCenterMatchmakingRuleCreateRequestDataAttributes    `json:"attributes"`
	Relationships GameCenterMatchmakingRuleCreateRequestDataRelationships `json:"relationships"`
}

// NewGameCenterMatchmakingRuleCreateRequestData instantiates a new GameCenterMatchmakingRuleCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingRuleCreateRequestData(type_ string, attributes GameCenterMatchmakingRuleCreateRequestDataAttributes, relationships GameCenterMatchmakingRuleCreateRequestDataRelationships) *GameCenterMatchmakingRuleCreateRequestData {
	this := GameCenterMatchmakingRuleCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewGameCenterMatchmakingRuleCreateRequestDataWithDefaults instantiates a new GameCenterMatchmakingRuleCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingRuleCreateRequestDataWithDefaults() *GameCenterMatchmakingRuleCreateRequestData {
	this := GameCenterMatchmakingRuleCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterMatchmakingRuleCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterMatchmakingRuleCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *GameCenterMatchmakingRuleCreateRequestData) GetAttributes() GameCenterMatchmakingRuleCreateRequestDataAttributes {
	if o == nil {
		var ret GameCenterMatchmakingRuleCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleCreateRequestData) GetAttributesOk() (*GameCenterMatchmakingRuleCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GameCenterMatchmakingRuleCreateRequestData) SetAttributes(v GameCenterMatchmakingRuleCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *GameCenterMatchmakingRuleCreateRequestData) GetRelationships() GameCenterMatchmakingRuleCreateRequestDataRelationships {
	if o == nil {
		var ret GameCenterMatchmakingRuleCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleCreateRequestData) GetRelationshipsOk() (*GameCenterMatchmakingRuleCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *GameCenterMatchmakingRuleCreateRequestData) SetRelationships(v GameCenterMatchmakingRuleCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o GameCenterMatchmakingRuleCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingRuleCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableGameCenterMatchmakingRuleCreateRequestData struct {
	value *GameCenterMatchmakingRuleCreateRequestData
	isSet bool
}

func (v NullableGameCenterMatchmakingRuleCreateRequestData) Get() *GameCenterMatchmakingRuleCreateRequestData {
	return v.value
}

func (v *NullableGameCenterMatchmakingRuleCreateRequestData) Set(val *GameCenterMatchmakingRuleCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingRuleCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingRuleCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingRuleCreateRequestData(val *GameCenterMatchmakingRuleCreateRequestData) *NullableGameCenterMatchmakingRuleCreateRequestData {
	return &NullableGameCenterMatchmakingRuleCreateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingRuleCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingRuleCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
