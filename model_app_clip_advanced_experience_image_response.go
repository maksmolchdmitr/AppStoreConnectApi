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

// checks if the AppClipAdvancedExperienceImageResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAdvancedExperienceImageResponse{}

// AppClipAdvancedExperienceImageResponse struct for AppClipAdvancedExperienceImageResponse
type AppClipAdvancedExperienceImageResponse struct {
	Data  AppClipAdvancedExperienceImage `json:"data"`
	Links DocumentLinks                  `json:"links"`
}

// NewAppClipAdvancedExperienceImageResponse instantiates a new AppClipAdvancedExperienceImageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAdvancedExperienceImageResponse(data AppClipAdvancedExperienceImage, links DocumentLinks) *AppClipAdvancedExperienceImageResponse {
	this := AppClipAdvancedExperienceImageResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppClipAdvancedExperienceImageResponseWithDefaults instantiates a new AppClipAdvancedExperienceImageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAdvancedExperienceImageResponseWithDefaults() *AppClipAdvancedExperienceImageResponse {
	this := AppClipAdvancedExperienceImageResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipAdvancedExperienceImageResponse) GetData() AppClipAdvancedExperienceImage {
	if o == nil {
		var ret AppClipAdvancedExperienceImage
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceImageResponse) GetDataOk() (*AppClipAdvancedExperienceImage, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipAdvancedExperienceImageResponse) SetData(v AppClipAdvancedExperienceImage) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppClipAdvancedExperienceImageResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppClipAdvancedExperienceImageResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppClipAdvancedExperienceImageResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppClipAdvancedExperienceImageResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAdvancedExperienceImageResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppClipAdvancedExperienceImageResponse struct {
	value *AppClipAdvancedExperienceImageResponse
	isSet bool
}

func (v NullableAppClipAdvancedExperienceImageResponse) Get() *AppClipAdvancedExperienceImageResponse {
	return v.value
}

func (v *NullableAppClipAdvancedExperienceImageResponse) Set(val *AppClipAdvancedExperienceImageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAdvancedExperienceImageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAdvancedExperienceImageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAdvancedExperienceImageResponse(val *AppClipAdvancedExperienceImageResponse) *NullableAppClipAdvancedExperienceImageResponse {
	return &NullableAppClipAdvancedExperienceImageResponse{value: val, isSet: true}
}

func (v NullableAppClipAdvancedExperienceImageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAdvancedExperienceImageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
