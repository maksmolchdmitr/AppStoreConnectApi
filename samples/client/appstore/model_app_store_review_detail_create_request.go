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

// checks if the AppStoreReviewDetailCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreReviewDetailCreateRequest{}

// AppStoreReviewDetailCreateRequest struct for AppStoreReviewDetailCreateRequest
type AppStoreReviewDetailCreateRequest struct {
	Data AppStoreReviewDetailCreateRequestData `json:"data"`
}

// NewAppStoreReviewDetailCreateRequest instantiates a new AppStoreReviewDetailCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreReviewDetailCreateRequest(data AppStoreReviewDetailCreateRequestData) *AppStoreReviewDetailCreateRequest {
	this := AppStoreReviewDetailCreateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreReviewDetailCreateRequestWithDefaults instantiates a new AppStoreReviewDetailCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreReviewDetailCreateRequestWithDefaults() *AppStoreReviewDetailCreateRequest {
	this := AppStoreReviewDetailCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreReviewDetailCreateRequest) GetData() AppStoreReviewDetailCreateRequestData {
	if o == nil {
		var ret AppStoreReviewDetailCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreReviewDetailCreateRequest) GetDataOk() (*AppStoreReviewDetailCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreReviewDetailCreateRequest) SetData(v AppStoreReviewDetailCreateRequestData) {
	o.Data = v
}

func (o AppStoreReviewDetailCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreReviewDetailCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppStoreReviewDetailCreateRequest struct {
	value *AppStoreReviewDetailCreateRequest
	isSet bool
}

func (v NullableAppStoreReviewDetailCreateRequest) Get() *AppStoreReviewDetailCreateRequest {
	return v.value
}

func (v *NullableAppStoreReviewDetailCreateRequest) Set(val *AppStoreReviewDetailCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreReviewDetailCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreReviewDetailCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreReviewDetailCreateRequest(val *AppStoreReviewDetailCreateRequest) *NullableAppStoreReviewDetailCreateRequest {
	return &NullableAppStoreReviewDetailCreateRequest{value: val, isSet: true}
}

func (v NullableAppStoreReviewDetailCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreReviewDetailCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


