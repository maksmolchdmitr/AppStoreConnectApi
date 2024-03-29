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

// checks if the AppWithoutIncludesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppWithoutIncludesResponse{}

// AppWithoutIncludesResponse struct for AppWithoutIncludesResponse
type AppWithoutIncludesResponse struct {
	Data  PrereleaseVersion `json:"data"`
	Links DocumentLinks     `json:"links"`
}

// NewAppWithoutIncludesResponse instantiates a new AppWithoutIncludesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppWithoutIncludesResponse(data PrereleaseVersion, links DocumentLinks) *AppWithoutIncludesResponse {
	this := AppWithoutIncludesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppWithoutIncludesResponseWithDefaults instantiates a new AppWithoutIncludesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppWithoutIncludesResponseWithDefaults() *AppWithoutIncludesResponse {
	this := AppWithoutIncludesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppWithoutIncludesResponse) GetData() PrereleaseVersion {
	if o == nil {
		var ret PrereleaseVersion
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppWithoutIncludesResponse) GetDataOk() (*PrereleaseVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppWithoutIncludesResponse) SetData(v PrereleaseVersion) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppWithoutIncludesResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppWithoutIncludesResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppWithoutIncludesResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AppWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppWithoutIncludesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAppWithoutIncludesResponse struct {
	value *AppWithoutIncludesResponse
	isSet bool
}

func (v NullableAppWithoutIncludesResponse) Get() *AppWithoutIncludesResponse {
	return v.value
}

func (v *NullableAppWithoutIncludesResponse) Set(val *AppWithoutIncludesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppWithoutIncludesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppWithoutIncludesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppWithoutIncludesResponse(val *AppWithoutIncludesResponse) *NullableAppWithoutIncludesResponse {
	return &NullableAppWithoutIncludesResponse{value: val, isSet: true}
}

func (v NullableAppWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppWithoutIncludesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
