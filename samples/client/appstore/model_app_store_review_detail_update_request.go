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

// checks if the AppStoreReviewDetailUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreReviewDetailUpdateRequest{}

// AppStoreReviewDetailUpdateRequest struct for AppStoreReviewDetailUpdateRequest
type AppStoreReviewDetailUpdateRequest struct {
	Data AppStoreReviewDetailUpdateRequestData `json:"data"`
}

// NewAppStoreReviewDetailUpdateRequest instantiates a new AppStoreReviewDetailUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewDetailUpdateRequest(data AppStoreReviewDetailUpdateRequestData) *AppStoreReviewDetailUpdateRequest {
	this := AppStoreReviewDetailUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreReviewDetailUpdateRequestWithDefaults instantiates a new AppStoreReviewDetailUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewDetailUpdateRequestWithDefaults() *AppStoreReviewDetailUpdateRequest {
	this := AppStoreReviewDetailUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreReviewDetailUpdateRequest) GetData() AppStoreReviewDetailUpdateRequestData {
	if o == nil {
		var ret AppStoreReviewDetailUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailUpdateRequest) GetDataOk() (*AppStoreReviewDetailUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreReviewDetailUpdateRequest) SetData(v AppStoreReviewDetailUpdateRequestData) {
	o.Data = v
}

func (o AppStoreReviewDetailUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreReviewDetailUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppStoreReviewDetailUpdateRequest struct {
	value *AppStoreReviewDetailUpdateRequest
	isSet bool
}

func (v NullableAppStoreReviewDetailUpdateRequest) Get() *AppStoreReviewDetailUpdateRequest {
	return v.value
}

func (v *NullableAppStoreReviewDetailUpdateRequest) Set(val *AppStoreReviewDetailUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewDetailUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewDetailUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewDetailUpdateRequest(val *AppStoreReviewDetailUpdateRequest) *NullableAppStoreReviewDetailUpdateRequest {
	return &NullableAppStoreReviewDetailUpdateRequest{value: val, isSet: true}
}

func (v NullableAppStoreReviewDetailUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewDetailUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


