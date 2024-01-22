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

// checks if the AppClipAppStoreReviewDetailUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAppStoreReviewDetailUpdateRequest{}

// AppClipAppStoreReviewDetailUpdateRequest struct for AppClipAppStoreReviewDetailUpdateRequest
type AppClipAppStoreReviewDetailUpdateRequest struct {
	Data AppClipAppStoreReviewDetailUpdateRequestData `json:"data"`
}

// NewAppClipAppStoreReviewDetailUpdateRequest instantiates a new AppClipAppStoreReviewDetailUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAppStoreReviewDetailUpdateRequest(data AppClipAppStoreReviewDetailUpdateRequestData) *AppClipAppStoreReviewDetailUpdateRequest {
	this := AppClipAppStoreReviewDetailUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppClipAppStoreReviewDetailUpdateRequestWithDefaults instantiates a new AppClipAppStoreReviewDetailUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAppStoreReviewDetailUpdateRequestWithDefaults() *AppClipAppStoreReviewDetailUpdateRequest {
	this := AppClipAppStoreReviewDetailUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipAppStoreReviewDetailUpdateRequest) GetData() AppClipAppStoreReviewDetailUpdateRequestData {
	if o == nil {
		var ret AppClipAppStoreReviewDetailUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipAppStoreReviewDetailUpdateRequest) GetDataOk() (*AppClipAppStoreReviewDetailUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipAppStoreReviewDetailUpdateRequest) SetData(v AppClipAppStoreReviewDetailUpdateRequestData) {
	o.Data = v
}

func (o AppClipAppStoreReviewDetailUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAppStoreReviewDetailUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppClipAppStoreReviewDetailUpdateRequest struct {
	value *AppClipAppStoreReviewDetailUpdateRequest
	isSet bool
}

func (v NullableAppClipAppStoreReviewDetailUpdateRequest) Get() *AppClipAppStoreReviewDetailUpdateRequest {
	return v.value
}

func (v *NullableAppClipAppStoreReviewDetailUpdateRequest) Set(val *AppClipAppStoreReviewDetailUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAppStoreReviewDetailUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAppStoreReviewDetailUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAppStoreReviewDetailUpdateRequest(val *AppClipAppStoreReviewDetailUpdateRequest) *NullableAppClipAppStoreReviewDetailUpdateRequest {
	return &NullableAppClipAppStoreReviewDetailUpdateRequest{value: val, isSet: true}
}

func (v NullableAppClipAppStoreReviewDetailUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAppStoreReviewDetailUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}