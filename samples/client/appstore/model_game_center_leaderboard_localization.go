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

// checks if the GameCenterLeaderboardLocalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardLocalization{}

// GameCenterLeaderboardLocalization struct for GameCenterLeaderboardLocalization
type GameCenterLeaderboardLocalization struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *GameCenterLeaderboardLocalizationAttributes `json:"attributes,omitempty"`
	Relationships *GameCenterLeaderboardLocalizationRelationships `json:"relationships,omitempty"`
	Links *ResourceLinks `json:"links,omitempty"`
}

// NewGameCenterLeaderboardLocalization instantiates a new GameCenterLeaderboardLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardLocalization(type_ string, id string) *GameCenterLeaderboardLocalization {
	this := GameCenterLeaderboardLocalization{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterLeaderboardLocalizationWithDefaults instantiates a new GameCenterLeaderboardLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardLocalizationWithDefaults() *GameCenterLeaderboardLocalization {
	this := GameCenterLeaderboardLocalization{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardLocalization) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalization) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardLocalization) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterLeaderboardLocalization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalization) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterLeaderboardLocalization) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterLeaderboardLocalization) GetAttributes() GameCenterLeaderboardLocalizationAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterLeaderboardLocalizationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalization) GetAttributesOk() (*GameCenterLeaderboardLocalizationAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterLeaderboardLocalization) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterLeaderboardLocalizationAttributes and assigns it to the Attributes field.
func (o *GameCenterLeaderboardLocalization) SetAttributes(v GameCenterLeaderboardLocalizationAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GameCenterLeaderboardLocalization) GetRelationships() GameCenterLeaderboardLocalizationRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GameCenterLeaderboardLocalizationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalization) GetRelationshipsOk() (*GameCenterLeaderboardLocalizationRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GameCenterLeaderboardLocalization) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GameCenterLeaderboardLocalizationRelationships and assigns it to the Relationships field.
func (o *GameCenterLeaderboardLocalization) SetRelationships(v GameCenterLeaderboardLocalizationRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterLeaderboardLocalization) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardLocalization) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterLeaderboardLocalization) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GameCenterLeaderboardLocalization) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GameCenterLeaderboardLocalization) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardLocalization) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	if !IsNil(o.Relationships) {
		toSerialize["relationships"] = o.Relationships
	}
	if !IsNil(o.Links) {
		toSerialize["links"] = o.Links
	}
	return toSerialize, nil
}

type NullableGameCenterLeaderboardLocalization struct {
	value *GameCenterLeaderboardLocalization
	isSet bool
}

func (v NullableGameCenterLeaderboardLocalization) Get() *GameCenterLeaderboardLocalization {
	return v.value
}

func (v *NullableGameCenterLeaderboardLocalization) Set(val *GameCenterLeaderboardLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardLocalization(val *GameCenterLeaderboardLocalization) *NullableGameCenterLeaderboardLocalization {
	return &NullableGameCenterLeaderboardLocalization{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

