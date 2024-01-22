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

// checks if the AppClipAdvancedExperienceCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceCreateRequestData{}

// AppClipAdvancedExperienceCreateRequestData struct for AppClipAdvancedExperienceCreateRequestData
type AppClipAdvancedExperienceCreateRequestData struct {
	Type          string                                                  `json:"type"`
	Attributes    AppClipAdvancedExperienceCreateRequestDataAttributes    `json:"attributes"`
	Relationships AppClipAdvancedExperienceCreateRequestDataRelationships `json:"relationships"`
}

// NewAppClipAdvancedExperienceCreateRequestData instantiates a new AppClipAdvancedExperienceCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceCreateRequestData(type_ string, attributes AppClipAdvancedExperienceCreateRequestDataAttributes, relationships AppClipAdvancedExperienceCreateRequestDataRelationships) *AppClipAdvancedExperienceCreateRequestData {
	this := AppClipAdvancedExperienceCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppClipAdvancedExperienceCreateRequestDataWithDefaults instantiates a new AppClipAdvancedExperienceCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceCreateRequestDataWithDefaults() *AppClipAdvancedExperienceCreateRequestData {
	this := AppClipAdvancedExperienceCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipAdvancedExperienceCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipAdvancedExperienceCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppClipAdvancedExperienceCreateRequestData) GetAttributes() AppClipAdvancedExperienceCreateRequestDataAttributes {
	if o == nil {
		var ret AppClipAdvancedExperienceCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceCreateRequestData) GetAttributesOk() (*AppClipAdvancedExperienceCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppClipAdvancedExperienceCreateRequestData) SetAttributes(v AppClipAdvancedExperienceCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppClipAdvancedExperienceCreateRequestData) GetRelationships() AppClipAdvancedExperienceCreateRequestDataRelationships {
	if o == nil {
		var ret AppClipAdvancedExperienceCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceCreateRequestData) GetRelationshipsOk() (*AppClipAdvancedExperienceCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppClipAdvancedExperienceCreateRequestData) SetRelationships(v AppClipAdvancedExperienceCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppClipAdvancedExperienceCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceCreateRequestData struct {
	value *AppClipAdvancedExperienceCreateRequestData
	isSet bool
}

func (v NullableAppClipAdvancedExperienceCreateRequestData) Get() *AppClipAdvancedExperienceCreateRequestData {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceCreateRequestData) Set(val *AppClipAdvancedExperienceCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceCreateRequestData(val *AppClipAdvancedExperienceCreateRequestData) *NullableAppClipAdvancedExperienceCreateRequestData {
	return &NullableAppClipAdvancedExperienceCreateRequestData{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
