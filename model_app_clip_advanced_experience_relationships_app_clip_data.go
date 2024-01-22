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

// checks if the AppClipAdvancedExperienceRelationshipsAppClipData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceRelationshipsAppClipData{}

// AppClipAdvancedExperienceRelationshipsAppClipData struct for AppClipAdvancedExperienceRelationshipsAppClipData
type AppClipAdvancedExperienceRelationshipsAppClipData struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewAppClipAdvancedExperienceRelationshipsAppClipData instantiates a new AppClipAdvancedExperienceRelationshipsAppClipData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceRelationshipsAppClipData(type_ string, id string) *AppClipAdvancedExperienceRelationshipsAppClipData {
	this := AppClipAdvancedExperienceRelationshipsAppClipData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppClipAdvancedExperienceRelationshipsAppClipDataWithDefaults instantiates a new AppClipAdvancedExperienceRelationshipsAppClipData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceRelationshipsAppClipDataWithDefaults() *AppClipAdvancedExperienceRelationshipsAppClipData {
	this := AppClipAdvancedExperienceRelationshipsAppClipData{}
	return &this
}

// GetType returns the Type field value
func (o *AppClipAdvancedExperienceRelationshipsAppClipData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceRelationshipsAppClipData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppClipAdvancedExperienceRelationshipsAppClipData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppClipAdvancedExperienceRelationshipsAppClipData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceRelationshipsAppClipData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppClipAdvancedExperienceRelationshipsAppClipData) SetId(v string) {
	o.Id = v
}

func (o AppClipAdvancedExperienceRelationshipsAppClipData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceRelationshipsAppClipData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceRelationshipsAppClipData struct {
	value *AppClipAdvancedExperienceRelationshipsAppClipData
	isSet bool
}

func (v NullableAppClipAdvancedExperienceRelationshipsAppClipData) Get() *AppClipAdvancedExperienceRelationshipsAppClipData {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceRelationshipsAppClipData) Set(val *AppClipAdvancedExperienceRelationshipsAppClipData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceRelationshipsAppClipData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceRelationshipsAppClipData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceRelationshipsAppClipData(val *AppClipAdvancedExperienceRelationshipsAppClipData) *NullableAppClipAdvancedExperienceRelationshipsAppClipData {
	return &NullableAppClipAdvancedExperienceRelationshipsAppClipData{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceRelationshipsAppClipData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceRelationshipsAppClipData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
