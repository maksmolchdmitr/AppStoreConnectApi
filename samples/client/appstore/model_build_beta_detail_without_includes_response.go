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

// checks if the BuildBetaDetailWithoutIncludesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildBetaDetailWithoutIncludesResponse{}

// BuildBetaDetailWithoutIncludesResponse struct for BuildBetaDetailWithoutIncludesResponse
type BuildBetaDetailWithoutIncludesResponse struct {
	Data Build `json:"data"`
	Links DocumentLinks `json:"links"`
}

// NewBuildBetaDetailWithoutIncludesResponse instantiates a new BuildBetaDetailWithoutIncludesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBetaDetailWithoutIncludesResponse(data Build, links DocumentLinks) *BuildBetaDetailWithoutIncludesResponse {
	this := BuildBetaDetailWithoutIncludesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBuildBetaDetailWithoutIncludesResponseWithDefaults instantiates a new BuildBetaDetailWithoutIncludesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBetaDetailWithoutIncludesResponseWithDefaults() *BuildBetaDetailWithoutIncludesResponse {
	this := BuildBetaDetailWithoutIncludesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BuildBetaDetailWithoutIncludesResponse) GetData() Build {
	if o == nil {
		var ret Build
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailWithoutIncludesResponse) GetDataOk() (*Build, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BuildBetaDetailWithoutIncludesResponse) SetData(v Build) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BuildBetaDetailWithoutIncludesResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailWithoutIncludesResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BuildBetaDetailWithoutIncludesResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BuildBetaDetailWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildBetaDetailWithoutIncludesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableBuildBetaDetailWithoutIncludesResponse struct {
	value *BuildBetaDetailWithoutIncludesResponse
	isSet bool
}

func (v NullableBuildBetaDetailWithoutIncludesResponse) Get() *BuildBetaDetailWithoutIncludesResponse {
	return v.value
}

func (v *NullableBuildBetaDetailWithoutIncludesResponse) Set(val *BuildBetaDetailWithoutIncludesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBetaDetailWithoutIncludesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBetaDetailWithoutIncludesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBetaDetailWithoutIncludesResponse(val *BuildBetaDetailWithoutIncludesResponse) *NullableBuildBetaDetailWithoutIncludesResponse {
	return &NullableBuildBetaDetailWithoutIncludesResponse{value: val, isSet: true}
}

func (v NullableBuildBetaDetailWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBetaDetailWithoutIncludesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


