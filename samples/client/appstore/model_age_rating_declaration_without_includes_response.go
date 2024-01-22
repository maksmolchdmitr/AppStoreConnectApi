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

// checks if the AgeRatingDeclarationWithoutIncludesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgeRatingDeclarationWithoutIncludesResponse{}

// AgeRatingDeclarationWithoutIncludesResponse struct for AgeRatingDeclarationWithoutIncludesResponse
type AgeRatingDeclarationWithoutIncludesResponse struct {
	Data AppStoreVersion `json:"data"`
	Links DocumentLinks `json:"links"`
}

// NewAgeRatingDeclarationWithoutIncludesResponse instantiates a new AgeRatingDeclarationWithoutIncludesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgeRatingDeclarationWithoutIncludesResponse(data AppStoreVersion, links DocumentLinks) *AgeRatingDeclarationWithoutIncludesResponse {
	this := AgeRatingDeclarationWithoutIncludesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAgeRatingDeclarationWithoutIncludesResponseWithDefaults instantiates a new AgeRatingDeclarationWithoutIncludesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgeRatingDeclarationWithoutIncludesResponseWithDefaults() *AgeRatingDeclarationWithoutIncludesResponse {
	this := AgeRatingDeclarationWithoutIncludesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AgeRatingDeclarationWithoutIncludesResponse) GetData() AppStoreVersion {
	if o == nil {
		var ret AppStoreVersion
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationWithoutIncludesResponse) GetDataOk() (*AppStoreVersion, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AgeRatingDeclarationWithoutIncludesResponse) SetData(v AppStoreVersion) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AgeRatingDeclarationWithoutIncludesResponse) GetLinks() DocumentLinks {
	if o == nil {
		var ret DocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AgeRatingDeclarationWithoutIncludesResponse) GetLinksOk() (*DocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AgeRatingDeclarationWithoutIncludesResponse) SetLinks(v DocumentLinks) {
	o.Links = v
}

func (o AgeRatingDeclarationWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgeRatingDeclarationWithoutIncludesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	return toSerialize, nil
}

type NullableAgeRatingDeclarationWithoutIncludesResponse struct {
	value *AgeRatingDeclarationWithoutIncludesResponse
	isSet bool
}

func (v NullableAgeRatingDeclarationWithoutIncludesResponse) Get() *AgeRatingDeclarationWithoutIncludesResponse {
	return v.value
}

func (v *NullableAgeRatingDeclarationWithoutIncludesResponse) Set(val *AgeRatingDeclarationWithoutIncludesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAgeRatingDeclarationWithoutIncludesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAgeRatingDeclarationWithoutIncludesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgeRatingDeclarationWithoutIncludesResponse(val *AgeRatingDeclarationWithoutIncludesResponse) *NullableAgeRatingDeclarationWithoutIncludesResponse {
	return &NullableAgeRatingDeclarationWithoutIncludesResponse{value: val, isSet: true}
}

func (v NullableAgeRatingDeclarationWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgeRatingDeclarationWithoutIncludesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


