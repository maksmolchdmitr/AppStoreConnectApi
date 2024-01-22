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

// checks if the GameCenterMatchmakingRuleSetAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingRuleSetAttributes{}

// GameCenterMatchmakingRuleSetAttributes struct for GameCenterMatchmakingRuleSetAttributes
type GameCenterMatchmakingRuleSetAttributes struct {
	ReferenceName *string `json:"referenceName,omitempty"`
	RuleLanguageVersion *int32 `json:"ruleLanguageVersion,omitempty"`
	MinPlayers *int32 `json:"minPlayers,omitempty"`
	MaxPlayers *int32 `json:"maxPlayers,omitempty"`
}

// NewGameCenterMatchmakingRuleSetAttributes instantiates a new GameCenterMatchmakingRuleSetAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingRuleSetAttributes() *GameCenterMatchmakingRuleSetAttributes {
	this := GameCenterMatchmakingRuleSetAttributes{}
	return &this
}

// NewGameCenterMatchmakingRuleSetAttributesWithDefaults instantiates a new GameCenterMatchmakingRuleSetAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingRuleSetAttributesWithDefaults() *GameCenterMatchmakingRuleSetAttributes {
	this := GameCenterMatchmakingRuleSetAttributes{}
	return &this
}

// GetReferenceName returns the ReferenceName field value if set, zero value otherwise.
func (o *GameCenterMatchmakingRuleSetAttributes) GetReferenceName() string {
	if o == nil || IsNil(o.ReferenceName) {
		var ret string
		return ret
	}
	return *o.ReferenceName
}

// GetReferenceNameOk returns a tuple with the ReferenceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetAttributes) GetReferenceNameOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceName) {
		return nil, false
	}
	return o.ReferenceName, true
}

// HasReferenceName returns a boolean if a field has been set.
func (o *GameCenterMatchmakingRuleSetAttributes) HasReferenceName() bool {
	if o != nil && !IsNil(o.ReferenceName) {
		return true
	}

	return false
}

// SetReferenceName gets a reference to the given string and assigns it to the ReferenceName field.
func (o *GameCenterMatchmakingRuleSetAttributes) SetReferenceName(v string) {
	o.ReferenceName = &v
}

// GetRuleLanguageVersion returns the RuleLanguageVersion field value if set, zero value otherwise.
func (o *GameCenterMatchmakingRuleSetAttributes) GetRuleLanguageVersion() int32 {
	if o == nil || IsNil(o.RuleLanguageVersion) {
		var ret int32
		return ret
	}
	return *o.RuleLanguageVersion
}

// GetRuleLanguageVersionOk returns a tuple with the RuleLanguageVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetAttributes) GetRuleLanguageVersionOk() (*int32, bool) {
	if o == nil || IsNil(o.RuleLanguageVersion) {
		return nil, false
	}
	return o.RuleLanguageVersion, true
}

// HasRuleLanguageVersion returns a boolean if a field has been set.
func (o *GameCenterMatchmakingRuleSetAttributes) HasRuleLanguageVersion() bool {
	if o != nil && !IsNil(o.RuleLanguageVersion) {
		return true
	}

	return false
}

// SetRuleLanguageVersion gets a reference to the given int32 and assigns it to the RuleLanguageVersion field.
func (o *GameCenterMatchmakingRuleSetAttributes) SetRuleLanguageVersion(v int32) {
	o.RuleLanguageVersion = &v
}

// GetMinPlayers returns the MinPlayers field value if set, zero value otherwise.
func (o *GameCenterMatchmakingRuleSetAttributes) GetMinPlayers() int32 {
	if o == nil || IsNil(o.MinPlayers) {
		var ret int32
		return ret
	}
	return *o.MinPlayers
}

// GetMinPlayersOk returns a tuple with the MinPlayers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetAttributes) GetMinPlayersOk() (*int32, bool) {
	if o == nil || IsNil(o.MinPlayers) {
		return nil, false
	}
	return o.MinPlayers, true
}

// HasMinPlayers returns a boolean if a field has been set.
func (o *GameCenterMatchmakingRuleSetAttributes) HasMinPlayers() bool {
	if o != nil && !IsNil(o.MinPlayers) {
		return true
	}

	return false
}

// SetMinPlayers gets a reference to the given int32 and assigns it to the MinPlayers field.
func (o *GameCenterMatchmakingRuleSetAttributes) SetMinPlayers(v int32) {
	o.MinPlayers = &v
}

// GetMaxPlayers returns the MaxPlayers field value if set, zero value otherwise.
func (o *GameCenterMatchmakingRuleSetAttributes) GetMaxPlayers() int32 {
	if o == nil || IsNil(o.MaxPlayers) {
		var ret int32
		return ret
	}
	return *o.MaxPlayers
}

// GetMaxPlayersOk returns a tuple with the MaxPlayers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingRuleSetAttributes) GetMaxPlayersOk() (*int32, bool) {
	if o == nil || IsNil(o.MaxPlayers) {
		return nil, false
	}
	return o.MaxPlayers, true
}

// HasMaxPlayers returns a boolean if a field has been set.
func (o *GameCenterMatchmakingRuleSetAttributes) HasMaxPlayers() bool {
	if o != nil && !IsNil(o.MaxPlayers) {
		return true
	}

	return false
}

// SetMaxPlayers gets a reference to the given int32 and assigns it to the MaxPlayers field.
func (o *GameCenterMatchmakingRuleSetAttributes) SetMaxPlayers(v int32) {
	o.MaxPlayers = &v
}

func (o GameCenterMatchmakingRuleSetAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingRuleSetAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ReferenceName) {
		toSerialize["referenceName"] = o.ReferenceName
	}
	if !IsNil(o.RuleLanguageVersion) {
		toSerialize["ruleLanguageVersion"] = o.RuleLanguageVersion
	}
	if !IsNil(o.MinPlayers) {
		toSerialize["minPlayers"] = o.MinPlayers
	}
	if !IsNil(o.MaxPlayers) {
		toSerialize["maxPlayers"] = o.MaxPlayers
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingRuleSetAttributes struct {
	value *GameCenterMatchmakingRuleSetAttributes
	isSet bool
}

func (v NullableGameCenterMatchmakingRuleSetAttributes) Get() *GameCenterMatchmakingRuleSetAttributes {
	return v.value
}

func (v *NullableGameCenterMatchmakingRuleSetAttributes) Set(val *GameCenterMatchmakingRuleSetAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingRuleSetAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingRuleSetAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingRuleSetAttributes(val *GameCenterMatchmakingRuleSetAttributes) *NullableGameCenterMatchmakingRuleSetAttributes {
	return &NullableGameCenterMatchmakingRuleSetAttributes{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingRuleSetAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingRuleSetAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


