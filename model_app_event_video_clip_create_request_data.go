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

// checks if the AppEventVideoClipCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventVideoClipCreateRequestData{}

// AppEventVideoClipCreateRequestData struct for AppEventVideoClipCreateRequestData
type AppEventVideoClipCreateRequestData struct {
	Type          string                                           `json:"type"`
	Attributes    AppEventVideoClipCreateRequestDataAttributes     `json:"attributes"`
	Relationships AppEventScreenshotCreateRequestDataRelationships `json:"relationships"`
}

// NewAppEventVideoClipCreateRequestData instantiates a new AppEventVideoClipCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventVideoClipCreateRequestData(type_ string, attributes AppEventVideoClipCreateRequestDataAttributes, relationships AppEventScreenshotCreateRequestDataRelationships) *AppEventVideoClipCreateRequestData {
	this := AppEventVideoClipCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppEventVideoClipCreateRequestDataWithDefaults instantiates a new AppEventVideoClipCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventVideoClipCreateRequestDataWithDefaults() *AppEventVideoClipCreateRequestData {
	this := AppEventVideoClipCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppEventVideoClipCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppEventVideoClipCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppEventVideoClipCreateRequestData) GetAttributes() AppEventVideoClipCreateRequestDataAttributes {
	if o == nil {
		var ret AppEventVideoClipCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipCreateRequestData) GetAttributesOk() (*AppEventVideoClipCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppEventVideoClipCreateRequestData) SetAttributes(v AppEventVideoClipCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppEventVideoClipCreateRequestData) GetRelationships() AppEventScreenshotCreateRequestDataRelationships {
	if o == nil {
		var ret AppEventScreenshotCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppEventVideoClipCreateRequestData) GetRelationshipsOk() (*AppEventScreenshotCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppEventVideoClipCreateRequestData) SetRelationships(v AppEventScreenshotCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppEventVideoClipCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventVideoClipCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppEventVideoClipCreateRequestData struct {
	value *AppEventVideoClipCreateRequestData
	isSet bool
}

func (v NullableAppEventVideoClipCreateRequestData) Get() *AppEventVideoClipCreateRequestData {
	return v.value
}

func (v *NullableAppEventVideoClipCreateRequestData) Set(val *AppEventVideoClipCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventVideoClipCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventVideoClipCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventVideoClipCreateRequestData(val *AppEventVideoClipCreateRequestData) *NullableAppEventVideoClipCreateRequestData {
	return &NullableAppEventVideoClipCreateRequestData{value: val, isSet: true}
}

func (v NullableAppEventVideoClipCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventVideoClipCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
