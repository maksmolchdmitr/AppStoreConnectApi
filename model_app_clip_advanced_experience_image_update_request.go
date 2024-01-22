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

// checks if the AppClipAdvancedExperienceImageUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceImageUpdateRequest{}

// AppClipAdvancedExperienceImageUpdateRequest struct for AppClipAdvancedExperienceImageUpdateRequest
type AppClipAdvancedExperienceImageUpdateRequest struct {
	Data AppClipAdvancedExperienceImageUpdateRequestData `json:"data"`
}

// NewAppClipAdvancedExperienceImageUpdateRequest instantiates a new AppClipAdvancedExperienceImageUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceImageUpdateRequest(data AppClipAdvancedExperienceImageUpdateRequestData) *AppClipAdvancedExperienceImageUpdateRequest {
	this := AppClipAdvancedExperienceImageUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppClipAdvancedExperienceImageUpdateRequestWithDefaults instantiates a new AppClipAdvancedExperienceImageUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceImageUpdateRequestWithDefaults() *AppClipAdvancedExperienceImageUpdateRequest {
	this := AppClipAdvancedExperienceImageUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipAdvancedExperienceImageUpdateRequest) GetData() AppClipAdvancedExperienceImageUpdateRequestData {
	if o == nil {
		var ret AppClipAdvancedExperienceImageUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceImageUpdateRequest) GetDataOk() (*AppClipAdvancedExperienceImageUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipAdvancedExperienceImageUpdateRequest) SetData(v AppClipAdvancedExperienceImageUpdateRequestData) {
	o.Data = v
}

func (o AppClipAdvancedExperienceImageUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceImageUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceImageUpdateRequest struct {
	value *AppClipAdvancedExperienceImageUpdateRequest
	isSet bool
}

func (v NullableAppClipAdvancedExperienceImageUpdateRequest) Get() *AppClipAdvancedExperienceImageUpdateRequest {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceImageUpdateRequest) Set(val *AppClipAdvancedExperienceImageUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceImageUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceImageUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceImageUpdateRequest(val *AppClipAdvancedExperienceImageUpdateRequest) *NullableAppClipAdvancedExperienceImageUpdateRequest {
	return &NullableAppClipAdvancedExperienceImageUpdateRequest{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceImageUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceImageUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
