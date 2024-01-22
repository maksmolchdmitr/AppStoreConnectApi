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

// checks if the AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage{}

// AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage struct for AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage
type AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage struct {
	Data AppClipAdvancedExperienceRelationshipsHeaderImageData `json:"data"`
}

// NewAppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage instantiates a new AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage(data AppClipAdvancedExperienceRelationshipsHeaderImageData) *AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage {
	this := AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage{}
	this.Data = data
	return &this
}

// NewAppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImageWithDefaults instantiates a new AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImageWithDefaults() *AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage {
	this := AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage) GetData() AppClipAdvancedExperienceRelationshipsHeaderImageData {
	if o == nil {
		var ret AppClipAdvancedExperienceRelationshipsHeaderImageData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage) GetDataOk() (*AppClipAdvancedExperienceRelationshipsHeaderImageData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage) SetData(v AppClipAdvancedExperienceRelationshipsHeaderImageData) {
	o.Data = v
}

func (o AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage struct {
	value *AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage
	isSet bool
}

func (v NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage) Get() *AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage) Set(val *AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage(val *AppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage) *NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage {
	return &NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceCreateRequestDataRelationshipsHeaderImage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


