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

// checks if the GameCenterMatchmakingQueueCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingQueueCreateRequestData{}

// GameCenterMatchmakingQueueCreateRequestData struct for GameCenterMatchmakingQueueCreateRequestData
type GameCenterMatchmakingQueueCreateRequestData struct {
	Type string `json:"type"`
	Attributes GameCenterMatchmakingQueueCreateRequestDataAttributes `json:"attributes"`
	Relationships GameCenterMatchmakingQueueCreateRequestDataRelationships `json:"relationships"`
}

// NewGameCenterMatchmakingQueueCreateRequestData instantiates a new GameCenterMatchmakingQueueCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingQueueCreateRequestData(type_ string, attributes GameCenterMatchmakingQueueCreateRequestDataAttributes, relationships GameCenterMatchmakingQueueCreateRequestDataRelationships) *GameCenterMatchmakingQueueCreateRequestData {
	this := GameCenterMatchmakingQueueCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewGameCenterMatchmakingQueueCreateRequestDataWithDefaults instantiates a new GameCenterMatchmakingQueueCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingQueueCreateRequestDataWithDefaults() *GameCenterMatchmakingQueueCreateRequestData {
	this := GameCenterMatchmakingQueueCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterMatchmakingQueueCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterMatchmakingQueueCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *GameCenterMatchmakingQueueCreateRequestData) GetAttributes() GameCenterMatchmakingQueueCreateRequestDataAttributes {
	if o == nil {
		var ret GameCenterMatchmakingQueueCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueCreateRequestData) GetAttributesOk() (*GameCenterMatchmakingQueueCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *GameCenterMatchmakingQueueCreateRequestData) SetAttributes(v GameCenterMatchmakingQueueCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *GameCenterMatchmakingQueueCreateRequestData) GetRelationships() GameCenterMatchmakingQueueCreateRequestDataRelationships {
	if o == nil {
		var ret GameCenterMatchmakingQueueCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueCreateRequestData) GetRelationshipsOk() (*GameCenterMatchmakingQueueCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *GameCenterMatchmakingQueueCreateRequestData) SetRelationships(v GameCenterMatchmakingQueueCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o GameCenterMatchmakingQueueCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingQueueCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableGameCenterMatchmakingQueueCreateRequestData struct {
	value *GameCenterMatchmakingQueueCreateRequestData
	isSet bool
}

func (v NullableGameCenterMatchmakingQueueCreateRequestData) Get() *GameCenterMatchmakingQueueCreateRequestData {
	return v.value
}

func (v *NullableGameCenterMatchmakingQueueCreateRequestData) Set(val *GameCenterMatchmakingQueueCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingQueueCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingQueueCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingQueueCreateRequestData(val *GameCenterMatchmakingQueueCreateRequestData) *NullableGameCenterMatchmakingQueueCreateRequestData {
	return &NullableGameCenterMatchmakingQueueCreateRequestData{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingQueueCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingQueueCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


