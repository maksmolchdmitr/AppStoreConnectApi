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

// checks if the BuildBetaDetailResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildBetaDetailResponse{}

// BuildBetaDetailResponse struct for BuildBetaDetailResponse
type BuildBetaDetailResponse struct {
	Data BuildBetaDetail `json:"data"`
	Included []Build `json:"included,omitempty"`
	Links DocumentLinks `json:"links"`
}

// NewBuildBetaDetailResponse instantiates a new BuildBetaDetailResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBetaDetailResponse(data BuildBetaDetail, links DocumentLinks) *BuildBetaDetailResponse {
	this := BuildBetaDetailResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBuildBetaDetailResponseWithDefaults instantiates a new BuildBetaDetailResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBetaDetailResponseWithDefaults() *BuildBetaDetailResponse {
	this := BuildBetaDetailResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BuildBetaDetailResponse) GetData() BuildBetaDetail {
	if o == nil {
		var ret BuildBetaDetail
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailResponse) GetDataOk() (*BuildBetaDetail, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BuildBetaDetailResponse) SetData(v BuildBetaDetail) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BuildBetaDetailResponse) GetIncluded() []Build {
	if o == nil || IsNil(o.Included) {
		var ret []Build
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailResponse) GetIncludedOk() ([]Build, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BuildBetaDetailResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []Build and assigns it to the Included field.
func (o *BuildBetaDetailResponse) SetIncluded(v []Build) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *BuildBetaDetailResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BuildBetaDetailResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o BuildBetaDetailResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildBetaDetailResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableBuildBetaDetailResponse struct {
	value *BuildBetaDetailResponse
	isSet bool
}

func (v NullableBuildBetaDetailResponse) Get() *BuildBetaDetailResponse {
	return v.value
}

func (v *NullableBuildBetaDetailResponse) Set(val *BuildBetaDetailResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBetaDetailResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBetaDetailResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBetaDetailResponse(val *BuildBetaDetailResponse) *NullableBuildBetaDetailResponse {
	return &NullableBuildBetaDetailResponse{value: val, isSet: true}
}

func (v NullableBuildBetaDetailResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBetaDetailResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

