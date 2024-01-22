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

// checks if the AppClipAppStoreReviewDetailCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAppStoreReviewDetailCreateRequest{}

// AppClipAppStoreReviewDetailCreateRequest struct for AppClipAppStoreReviewDetailCreateRequest
type AppClipAppStoreReviewDetailCreateRequest struct {
	Data AppClipAppStoreReviewDetailCreateRequestData `json:"data"`
}

// NewAppClipAppStoreReviewDetailCreateRequest instantiates a new AppClipAppStoreReviewDetailCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAppStoreReviewDetailCreateRequest(data AppClipAppStoreReviewDetailCreateRequestData) *AppClipAppStoreReviewDetailCreateRequest {
	this := AppClipAppStoreReviewDetailCreateRequest{}
	this.Data = data
	return &this
}

// NewAppClipAppStoreReviewDetailCreateRequestWithDefaults instantiates a new AppClipAppStoreReviewDetailCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAppStoreReviewDetailCreateRequestWithDefaults() *AppClipAppStoreReviewDetailCreateRequest {
	this := AppClipAppStoreReviewDetailCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipAppStoreReviewDetailCreateRequest) GetData() AppClipAppStoreReviewDetailCreateRequestData {
	if o == nil {
		var ret AppClipAppStoreReviewDetailCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipAppStoreReviewDetailCreateRequest) GetDataOk() (*AppClipAppStoreReviewDetailCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipAppStoreReviewDetailCreateRequest) SetData(v AppClipAppStoreReviewDetailCreateRequestData) {
	o.Data = v
}

func (o AppClipAppStoreReviewDetailCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAppStoreReviewDetailCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppClipAppStoreReviewDetailCreateRequest struct {
	value *AppClipAppStoreReviewDetailCreateRequest
	isSet bool
}

func (v NullableAppClipAppStoreReviewDetailCreateRequest) Get() *AppClipAppStoreReviewDetailCreateRequest {
	return v.value
}

func (v *NullableAppClipAppStoreReviewDetailCreateRequest) Set(val *AppClipAppStoreReviewDetailCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAppStoreReviewDetailCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAppStoreReviewDetailCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAppStoreReviewDetailCreateRequest(val *AppClipAppStoreReviewDetailCreateRequest) *NullableAppClipAppStoreReviewDetailCreateRequest {
	return &NullableAppClipAppStoreReviewDetailCreateRequest{value: val, isSet: true}
}

func (v NullableAppClipAppStoreReviewDetailCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAppStoreReviewDetailCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
