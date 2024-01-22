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

// checks if the GameCenterDetailRelationshipsGameCenterLeaderboardSets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterDetailRelationshipsGameCenterLeaderboardSets{}

// GameCenterDetailRelationshipsGameCenterLeaderboardSets struct for GameCenterDetailRelationshipsGameCenterLeaderboardSets
type GameCenterDetailRelationshipsGameCenterLeaderboardSets struct {
	Links *AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks       `json:"links,omitempty"`
	Meta  *PagingInformation                                                `json:"meta,omitempty"`
	Data  []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner `json:"data,omitempty"`
}

// NewGameCenterDetailRelationshipsGameCenterLeaderboardSets instantiates a new GameCenterDetailRelationshipsGameCenterLeaderboardSets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterDetailRelationshipsGameCenterLeaderboardSets() *GameCenterDetailRelationshipsGameCenterLeaderboardSets {
	this := GameCenterDetailRelationshipsGameCenterLeaderboardSets{}
	return &this
}

// NewGameCenterDetailRelationshipsGameCenterLeaderboardSetsWithDefaults instantiates a new GameCenterDetailRelationshipsGameCenterLeaderboardSets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterDetailRelationshipsGameCenterLeaderboardSetsWithDefaults() *GameCenterDetailRelationshipsGameCenterLeaderboardSets {
	this := GameCenterDetailRelationshipsGameCenterLeaderboardSets{}
	return &this
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSets) GetLinks() AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks {
	if o == nil || IsNil(o.Links) {
		var ret AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSets) GetLinksOk() (*AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSets) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks and assigns it to the Links field.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSets) SetLinks(v AppAvailabilityV2RelationshipsTerritoryAvailabilitiesLinks) {
	o.Links = &v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSets) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSets) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSets) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSets) SetMeta(v PagingInformation) {
	o.Meta = &v
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSets) GetData() []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner {
	if o == nil || IsNil(o.Data) {
		var ret []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSets) GetDataOk() ([]GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSets) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner and assigns it to the Data field.
func (o *GameCenterDetailRelationshipsGameCenterLeaderboardSets) SetData(v []GameCenterDetailRelationshipsGameCenterLeaderboardSetsDataInner) {
	o.Data = v
}

func (o GameCenterDetailRelationshipsGameCenterLeaderboardSets) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterDetailRelationshipsGameCenterLeaderboardSets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableGameCenterDetailRelationshipsGameCenterLeaderboardSets struct {
	value *GameCenterDetailRelationshipsGameCenterLeaderboardSets
	isSet bool
}

func (v NullableGameCenterDetailRelationshipsGameCenterLeaderboardSets) Get() *GameCenterDetailRelationshipsGameCenterLeaderboardSets {
	return v.value
}

func (v *NullableGameCenterDetailRelationshipsGameCenterLeaderboardSets) Set(val *GameCenterDetailRelationshipsGameCenterLeaderboardSets) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterDetailRelationshipsGameCenterLeaderboardSets) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterDetailRelationshipsGameCenterLeaderboardSets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterDetailRelationshipsGameCenterLeaderboardSets(val *GameCenterDetailRelationshipsGameCenterLeaderboardSets) *NullableGameCenterDetailRelationshipsGameCenterLeaderboardSets {
	return &NullableGameCenterDetailRelationshipsGameCenterLeaderboardSets{value: val, isSet: true}
}

func (v NullableGameCenterDetailRelationshipsGameCenterLeaderboardSets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterDetailRelationshipsGameCenterLeaderboardSets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
