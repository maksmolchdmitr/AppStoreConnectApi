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

// checks if the GameCenterAchievementLocalizationRelationshipsGameCenterAchievement type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementLocalizationRelationshipsGameCenterAchievement{}

// GameCenterAchievementLocalizationRelationshipsGameCenterAchievement struct for GameCenterAchievementLocalizationRelationshipsGameCenterAchievement
type GameCenterAchievementLocalizationRelationshipsGameCenterAchievement struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks              `json:"links,omitempty"`
	Data  *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData `json:"data,omitempty"`
}

// NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievement instantiates a new GameCenterAchievementLocalizationRelationshipsGameCenterAchievement object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievement() *GameCenterAchievementLocalizationRelationshipsGameCenterAchievement {
	this := GameCenterAchievementLocalizationRelationshipsGameCenterAchievement{}
	return &this
}

// NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievementWithDefaults instantiates a new GameCenterAchievementLocalizationRelationshipsGameCenterAchievement object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievementWithDefaults() *GameCenterAchievementLocalizationRelationshipsGameCenterAchievement {
	this := GameCenterAchievementLocalizationRelationshipsGameCenterAchievement{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievement) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievement) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievement) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievement) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievement) GetData() GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData {
	if o == nil || IsNil(o.Data) {
		var ret GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievement) GetDataOk() (*GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievement) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData and assigns it to the Data field.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievement) SetData(v GameCenterAchievementLocalizationRelationshipsGameCenterAchievementData) {
	o.Data = &v
}

func (o GameCenterAchievementLocalizationRelationshipsGameCenterAchievement) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementLocalizationRelationshipsGameCenterAchievement) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievement struct {
	value *GameCenterAchievementLocalizationRelationshipsGameCenterAchievement
	isSet bool
}

func (v NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievement) Get() *GameCenterAchievementLocalizationRelationshipsGameCenterAchievement {
	return v.value
}

func (v *NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievement) Set(val *GameCenterAchievementLocalizationRelationshipsGameCenterAchievement) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievement) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievement) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievement(val *GameCenterAchievementLocalizationRelationshipsGameCenterAchievement) *NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievement {
	return &NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievement{value: val, isSet: true}
}

func (v NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievement) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievement) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
