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

// checks if the GameCenterDetail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterDetail{}

// GameCenterDetail struct for GameCenterDetail
type GameCenterDetail struct {
	Type          string                         `json:"type"`
	Id            string                         `json:"id"`
	Attributes    *GameCenterDetailAttributes    `json:"attributes,omitempty"`
	Relationships *GameCenterDetailRelationships `json:"relationships,omitempty"`
	Links         *ResourceLinks                 `json:"links,omitempty"`
}

// NewGameCenterDetail instantiates a new GameCenterDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterDetail(type_ string, id string) *GameCenterDetail {
	this := GameCenterDetail{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewGameCenterDetailWithDefaults instantiates a new GameCenterDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterDetailWithDefaults() *GameCenterDetail {
	this := GameCenterDetail{}
	return &this
}

// GetType returns the Type field value
func (o *GameCenterDetail) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *GameCenterDetail) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *GameCenterDetail) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *GameCenterDetail) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *GameCenterDetail) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *GameCenterDetail) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *GameCenterDetail) GetAttributes() GameCenterDetailAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret GameCenterDetailAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetail) GetAttributesOk() (*GameCenterDetailAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *GameCenterDetail) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given GameCenterDetailAttributes and assigns it to the Attributes field.
func (o *GameCenterDetail) SetAttributes(v GameCenterDetailAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value if set, zero value otherwise.
func (o *GameCenterDetail) GetRelationships() GameCenterDetailRelationships {
	if o == nil || IsNil(o.Relationships) {
		var ret GameCenterDetailRelationships
		return ret
	}
	return *o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetail) GetRelationshipsOk() (*GameCenterDetailRelationships, bool) {
	if o == nil || IsNil(o.Relationships) {
		return nil, false
	}
	return o.Relationships, true
}

// HasRelationships returns a boolean if a field has been set.
func (o *GameCenterDetail) HasRelationships() bool {
	if o != nil && !IsNil(o.Relationships) {
		return true
	}

	return false
}

// SetRelationships gets a reference to the given GameCenterDetailRelationships and assigns it to the Relationships field.
func (o *GameCenterDetail) SetRelationships(v GameCenterDetailRelationships) {
	o.Relationships = &v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *GameCenterDetail) GetLinks() ResourceLinks {
	if o == nil || IsNil(o.Links) {
		var ret ResourceLinks
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterDetail) GetLinksOk() (*ResourceLinks, bool) {
	if o == nil || IsNil(o.Links) {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *GameCenterDetail) HasLinks() bool {
	if o != nil && !IsNil(o.Links) {
		return true
	}

	return false
}

// SetLinks gets a reference to the given ResourceLinks and assigns it to the Links field.
func (o *GameCenterDetail) SetLinks(v ResourceLinks) {
	o.Links = &v
}

func (o GameCenterDetail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterDetail) ToMap() (map[string]interface{}, error) {
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

type NullableGameCenterDetail struct {
	value *GameCenterDetail
	isSet bool
}

func (v NullableGameCenterDetail) Get() *GameCenterDetail {
	return v.value
}

func (v *NullableGameCenterDetail) Set(val *GameCenterDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterDetail(val *GameCenterDetail) *NullableGameCenterDetail {
	return &NullableGameCenterDetail{value: val, isSet: true}
}

func (v NullableGameCenterDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
