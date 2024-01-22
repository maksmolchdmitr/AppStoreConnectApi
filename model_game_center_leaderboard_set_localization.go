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

// checks if the GameCenterLeaderboardSetLocalization type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardSetLocalization{}

// GameCenterLeaderboardSetLocalization struct for GameCenterLeaderboardSetLocalization
type GameCenterLeaderboardSetLocalization struct {
	Type          string                                             `json:"type"`
	Id            string                                             `json:"id"`
	Attributes    *GameCenterLeaderboardSetLocalizationAttributes    `json:"attributes,omitempty"`
	Relationships *GameCenterLeaderboardSetLocalizationRelationships `json:"relationships,omitempty"`
	Links         *ResourceLinks                                     `json:"links,omitempty"`
}

// NewGameCenterLeaderboardSetLocalization instantiates a new GameCenterLeaderboardSetLocalization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardSetLocalization(type_ string, id string) *GameCenterLeaderboardSetLocalization {
	this := GameCenterLeaderboardSetLocalization{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterLeaderboardSetLocalizationWithDefaults instantiates a new GameCenterLeaderboardSetLocalization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardSetLocalizationWithDefaults() *GameCenterLeaderboardSetLocalization {
	this := GameCenterLeaderboardSetLocalization{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardSetLocalization) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalization) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardSetLocalization) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterLeaderboardSetLocalization) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalization) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterLeaderboardSetLocalization) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetLocalization) GetAttributes() GameCenterLeaderboardSetLocalizationAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterLeaderboardSetLocalizationAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalization) GetAttributesOk() (*GameCenterLeaderboardSetLocalizationAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetLocalization) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterLeaderboardSetLocalizationAttributes and assigns it to the Attributes field.
func (o *GameCenterLeaderboardSetLocalization) SetAttributes(v GameCenterLeaderboardSetLocalizationAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetLocalization) GetRelationships() GameCenterLeaderboardSetLocalizationRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GameCenterLeaderboardSetLocalizationRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalization) GetRelationshipsOk() (*GameCenterLeaderboardSetLocalizationRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetLocalization) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GameCenterLeaderboardSetLocalizationRelationships and assigns it to the Relationships field.
func (o *GameCenterLeaderboardSetLocalization) SetRelationships(v GameCenterLeaderboardSetLocalizationRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterLeaderboardSetLocalization) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardSetLocalization) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterLeaderboardSetLocalization) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GameCenterLeaderboardSetLocalization) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GameCenterLeaderboardSetLocalization) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardSetLocalization) ToMap() (map[string]interface{}, error) {
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

type NullableGameCenterLeaderboardSetLocalization struct {
	value *GameCenterLeaderboardSetLocalization
	isSet bool
}

func (v NullableGameCenterLeaderboardSetLocalization) Get() *GameCenterLeaderboardSetLocalization {
	return v.value
}

func (v *NullableGameCenterLeaderboardSetLocalization) Set(val *GameCenterLeaderboardSetLocalization) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardSetLocalization) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardSetLocalization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardSetLocalization(val *GameCenterLeaderboardSetLocalization) *NullableGameCenterLeaderboardSetLocalization {
	return &NullableGameCenterLeaderboardSetLocalization{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardSetLocalization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardSetLocalization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
