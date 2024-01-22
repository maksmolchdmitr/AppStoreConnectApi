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

// checks if the GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets{}

// GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets struct for GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets
type GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets struct {
	Data []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner `json:"data,omitempty"`
}

// NewGameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets instantiates a new GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets() *GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets {
	this := GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets{}
	return &this
}

// NewGameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSetsWithDefaults instantiates a new GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSetsWithDefaults() *GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets {
	this := GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets) GetData() []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets) GetDataOk() ([]GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner and assigns it to the Data field.
func (o *GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets) SetData(v []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) {
	o.Data = v
}

func (o GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets struct {
	value *GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets
	isSet bool
}

func (v NullableGameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets) Get() *GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets {
	return v.value
}

func (v *NullableGameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets) Set(val *GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets(val *GameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets) *NullableGameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets {
	return &NullableGameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardCreateRequestDataRelationshipsGameCenterLeaderboardSets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}