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

// checks if the AppClipAdvancedExperienceImageCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceImageCreateRequestData{}

// AppClipAdvancedExperienceImageCreateRequestData struct for AppClipAdvancedExperienceImageCreateRequestData
type AppClipAdvancedExperienceImageCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppClipAdvancedExperienceImageCreateRequestDataAttributes `json:"attributes"`
}

// NewAppClipAdvancedExperienceImageCreateRequestData instantiates a new AppClipAdvancedExperienceImageCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceImageCreateRequestData(type_ string, attributes AppClipAdvancedExperienceImageCreateRequestDataAttributes) *AppClipAdvancedExperienceImageCreateRequestData {
	this := AppClipAdvancedExperienceImageCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	return &this
}

// NewAppClipAdvancedExperienceImageCreateRequestDataWithDefaults instantiates a new AppClipAdvancedExperienceImageCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceImageCreateRequestDataWithDefaults() *AppClipAdvancedExperienceImageCreateRequestData {
	this := AppClipAdvancedExperienceImageCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipAdvancedExperienceImageCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceImageCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipAdvancedExperienceImageCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppClipAdvancedExperienceImageCreateRequestData) GetAttributes() AppClipAdvancedExperienceImageCreateRequestDataAttributes {
	if o == nil {
		var ret AppClipAdvancedExperienceImageCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceImageCreateRequestData) GetAttributesOk() (*AppClipAdvancedExperienceImageCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppClipAdvancedExperienceImageCreateRequestData) SetAttributes(v AppClipAdvancedExperienceImageCreateRequestDataAttributes) {
	o.Attributes = v
}

func (o AppClipAdvancedExperienceImageCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceImageCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceImageCreateRequestData struct {
	value *AppClipAdvancedExperienceImageCreateRequestData
	isSet bool
}

func (v NullableAppClipAdvancedExperienceImageCreateRequestData) Get() *AppClipAdvancedExperienceImageCreateRequestData {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceImageCreateRequestData) Set(val *AppClipAdvancedExperienceImageCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceImageCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceImageCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceImageCreateRequestData(val *AppClipAdvancedExperienceImageCreateRequestData) *NullableAppClipAdvancedExperienceImageCreateRequestData {
	return &NullableAppClipAdvancedExperienceImageCreateRequestData{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceImageCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceImageCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

