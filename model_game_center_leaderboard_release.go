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

// checks if the GameCenterLeaderboardRelease type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterLeaderboardRelease{}

// GameCenterLeaderboardRelease struct for GameCenterLeaderboardRelease
type GameCenterLeaderboardRelease struct {
	Type          string                                     `json:"type"`
	Id            string                                     `json:"id"`
	Attributes    *GameCenterAchievementReleaseAttributes    `json:"attributes,omitempty"`
	Relationships *GameCenterLeaderboardReleaseRelationships `json:"relationships,omitempty"`
	Links         *ResourceLinks                             `json:"links,omitempty"`
}

// NewGameCenterLeaderboardRelease instantiates a new GameCenterLeaderboardRelease object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterLeaderboardRelease(type_ string, id string) *GameCenterLeaderboardRelease {
	this := GameCenterLeaderboardRelease{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterLeaderboardReleaseWithDefaults instantiates a new GameCenterLeaderboardRelease object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterLeaderboardReleaseWithDefaults() *GameCenterLeaderboardRelease {
	this := GameCenterLeaderboardRelease{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterLeaderboardRelease) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardRelease) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterLeaderboardRelease) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterLeaderboardRelease) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardRelease) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterLeaderboardRelease) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterLeaderboardRelease) GetAttributes() GameCenterAchievementReleaseAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterAchievementReleaseAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardRelease) GetAttributesOk() (*GameCenterAchievementReleaseAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterLeaderboardRelease) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterAchievementReleaseAttributes and assigns it to the Attributes field.
func (o *GameCenterLeaderboardRelease) SetAttributes(v GameCenterAchievementReleaseAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GameCenterLeaderboardRelease) GetRelationships() GameCenterLeaderboardReleaseRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GameCenterLeaderboardReleaseRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardRelease) GetRelationshipsOk() (*GameCenterLeaderboardReleaseRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GameCenterLeaderboardRelease) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GameCenterLeaderboardReleaseRelationships and assigns it to the Relationships field.
func (o *GameCenterLeaderboardRelease) SetRelationships(v GameCenterLeaderboardReleaseRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterLeaderboardRelease) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterLeaderboardRelease) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterLeaderboardRelease) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GameCenterLeaderboardRelease) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GameCenterLeaderboardRelease) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterLeaderboardRelease) ToMap() (map[string]interface{}, error) {
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

type NullableGameCenterLeaderboardRelease struct {
	value *GameCenterLeaderboardRelease
	isSet bool
}

func (v NullableGameCenterLeaderboardRelease) Get() *GameCenterLeaderboardRelease {
	return v.value
}

func (v *NullableGameCenterLeaderboardRelease) Set(val *GameCenterLeaderboardRelease) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterLeaderboardRelease) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterLeaderboardRelease) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterLeaderboardRelease(val *GameCenterLeaderboardRelease) *NullableGameCenterLeaderboardRelease {
	return &NullableGameCenterLeaderboardRelease{value: val, isSet: true}
}

func (v NullableGameCenterLeaderboardRelease) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterLeaderboardRelease) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
