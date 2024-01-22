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

// checks if the GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage{}

// GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage struct for GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage
type GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks `json:"links,omitempty"`
	Data *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData `json:"data,omitempty"`
}

// NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage instantiates a new GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage() *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage {
	this := GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage{}
	return &this
}

// NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageWithDefaults instantiates a new GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageWithDefaults() *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage {
	this := GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) GetData() GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData {
	if o == nil || IsNil(o.Data) {
		var ret GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) GetDataOk() (*GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData and assigns it to the Data field.
func (o *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) SetData(v GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImageData) {
	o.Data = &v
}

func (o GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage struct {
	value *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage
	isSet bool
}

func (v NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) Get() *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage {
	return v.value
}

func (v *NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) Set(val *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage(val *GameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) *NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage {
	return &NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage{value: val, isSet: true}
}

func (v NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAchievementLocalizationRelationshipsGameCenterAchievementImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


