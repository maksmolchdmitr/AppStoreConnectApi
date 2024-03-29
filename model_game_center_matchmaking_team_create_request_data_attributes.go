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

// checks if the GameCenterMatchmakingTeamCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingTeamCreateRequestDataAttributes{}

// GameCenterMatchmakingTeamCreateRequestDataAttributes struct for GameCenterMatchmakingTeamCreateRequestDataAttributes
type GameCenterMatchmakingTeamCreateRequestDataAttributes struct {
	ReferenceName string `json:"referenceName"`
	MinPlayers    int32  `json:"minPlayers"`
	MaxPlayers    int32  `json:"maxPlayers"`
}

// NewGameCenterMatchmakingTeamCreateRequestDataAttributes instantiates a new GameCenterMatchmakingTeamCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingTeamCreateRequestDataAttributes(referenceName string, minPlayers int32, maxPlayers int32) *GameCenterMatchmakingTeamCreateRequestDataAttributes {
	this := GameCenterMatchmakingTeamCreateRequestDataAttributes{}
	this.ReferenceName = referenceName
	this.MinPlayers = minPlayers
	this.MaxPlayers = maxPlayers
	return &this
}

// NewGameCenterMatchmakingTeamCreateRequestDataAttributesWithDefaults instantiates a new GameCenterMatchmakingTeamCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingTeamCreateRequestDataAttributesWithDefaults() *GameCenterMatchmakingTeamCreateRequestDataAttributes {
	this := GameCenterMatchmakingTeamCreateRequestDataAttributes{}
	return &this
}

// GetReferenceName returns the ReferenceName field value
func (o *GameCenterMatchmakingTeamCreateRequestDataAttributes) GetReferenceName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTeamCreateRequestDataAttributes) GetReferenceNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReferenceName, true
}

// SetReferenceName sets field value
func (o *GameCenterMatchmakingTeamCreateRequestDataAttributes) SetReferenceName(v string) {
	o.ReferenceName = v
}

// GetMinPlayers returns the MinPlayers field value
func (o *GameCenterMatchmakingTeamCreateRequestDataAttributes) GetMinPlayers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MinPlayers
}

// GetMinPlayersOk returns a tuple with the MinPlayers field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTeamCreateRequestDataAttributes) GetMinPlayersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MinPlayers, true
}

// SetMinPlayers sets field value
func (o *GameCenterMatchmakingTeamCreateRequestDataAttributes) SetMinPlayers(v int32) {
	o.MinPlayers = v
}

// GetMaxPlayers returns the MaxPlayers field value
func (o *GameCenterMatchmakingTeamCreateRequestDataAttributes) GetMaxPlayers() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.MaxPlayers
}

// GetMaxPlayersOk returns a tuple with the MaxPlayers field value
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingTeamCreateRequestDataAttributes) GetMaxPlayersOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.MaxPlayers, true
}

// SetMaxPlayers sets field value
func (o *GameCenterMatchmakingTeamCreateRequestDataAttributes) SetMaxPlayers(v int32) {
	o.MaxPlayers = v
}

func (o GameCenterMatchmakingTeamCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingTeamCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["referenceName"] = o.ReferenceName
	toSerialize["minPlayers"] = o.MinPlayers
	toSerialize["maxPlayers"] = o.MaxPlayers
	return toSerialize, nil
}

type NullableGameCenterMatchmakingTeamCreateRequestDataAttributes struct {
	value *GameCenterMatchmakingTeamCreateRequestDataAttributes
	isSet bool
}

func (v NullableGameCenterMatchmakingTeamCreateRequestDataAttributes) Get() *GameCenterMatchmakingTeamCreateRequestDataAttributes {
	return v.value
}

func (v *NullableGameCenterMatchmakingTeamCreateRequestDataAttributes) Set(val *GameCenterMatchmakingTeamCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingTeamCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingTeamCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingTeamCreateRequestDataAttributes(val *GameCenterMatchmakingTeamCreateRequestDataAttributes) *NullableGameCenterMatchmakingTeamCreateRequestDataAttributes {
	return &NullableGameCenterMatchmakingTeamCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingTeamCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingTeamCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
