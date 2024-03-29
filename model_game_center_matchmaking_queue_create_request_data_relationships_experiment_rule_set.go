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

// checks if the GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet{}

// GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet struct for GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet
type GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet struct {
	Data *GameCenterMatchmakingQueueRelationshipsRuleSetData `json:"data,omitempty"`
}

// NewGameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet instantiates a new GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet() *GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet {
	this := GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet{}
	return &this
}

// NewGameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSetWithDefaults instantiates a new GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSetWithDefaults() *GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet {
	this := GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet) GetData() GameCenterMatchmakingQueueRelationshipsRuleSetData {
	if o == nil || IsNil(o.Data) {
		var ret GameCenterMatchmakingQueueRelationshipsRuleSetData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet) GetDataOk() (*GameCenterMatchmakingQueueRelationshipsRuleSetData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GameCenterMatchmakingQueueRelationshipsRuleSetData and assigns it to the Data field.
func (o *GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet) SetData(v GameCenterMatchmakingQueueRelationshipsRuleSetData) {
	o.Data = &v
}

func (o GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet struct {
	value *GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet
	isSet bool
}

func (v NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet) Get() *GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet {
	return v.value
}

func (v *NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet) Set(val *GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet(val *GameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet) *NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet {
	return &NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingQueueCreateRequestDataRelationshipsExperimentRuleSet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
