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

// checks if the SubscriptionAppStoreReviewScreenshotUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionAppStoreReviewScreenshotUpdateRequest{}

// SubscriptionAppStoreReviewScreenshotUpdateRequest struct for SubscriptionAppStoreReviewScreenshotUpdateRequest
type SubscriptionAppStoreReviewScreenshotUpdateRequest struct {
	Data SubscriptionAppStoreReviewScreenshotUpdateRequestData `json:"data"`
}

// NewSubscriptionAppStoreReviewScreenshotUpdateRequest instantiates a new SubscriptionAppStoreReviewScreenshotUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionAppStoreReviewScreenshotUpdateRequest(data SubscriptionAppStoreReviewScreenshotUpdateRequestData) *SubscriptionAppStoreReviewScreenshotUpdateRequest {
	this := SubscriptionAppStoreReviewScreenshotUpdateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionAppStoreReviewScreenshotUpdateRequestWithDefaults instantiates a new SubscriptionAppStoreReviewScreenshotUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionAppStoreReviewScreenshotUpdateRequestWithDefaults() *SubscriptionAppStoreReviewScreenshotUpdateRequest {
	this := SubscriptionAppStoreReviewScreenshotUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionAppStoreReviewScreenshotUpdateRequest) GetData() SubscriptionAppStoreReviewScreenshotUpdateRequestData {
	if o == nil {
		var ret SubscriptionAppStoreReviewScreenshotUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionAppStoreReviewScreenshotUpdateRequest) GetDataOk() (*SubscriptionAppStoreReviewScreenshotUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionAppStoreReviewScreenshotUpdateRequest) SetData(v SubscriptionAppStoreReviewScreenshotUpdateRequestData) {
	o.Data = v
}

func (o SubscriptionAppStoreReviewScreenshotUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionAppStoreReviewScreenshotUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSubscriptionAppStoreReviewScreenshotUpdateRequest struct {
	value *SubscriptionAppStoreReviewScreenshotUpdateRequest
	isSet bool
}

func (v NullableSubscriptionAppStoreReviewScreenshotUpdateRequest) Get() *SubscriptionAppStoreReviewScreenshotUpdateRequest {
	return v.value
}

func (v *NullableSubscriptionAppStoreReviewScreenshotUpdateRequest) Set(val *SubscriptionAppStoreReviewScreenshotUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionAppStoreReviewScreenshotUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionAppStoreReviewScreenshotUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionAppStoreReviewScreenshotUpdateRequest(val *SubscriptionAppStoreReviewScreenshotUpdateRequest) *NullableSubscriptionAppStoreReviewScreenshotUpdateRequest {
	return &NullableSubscriptionAppStoreReviewScreenshotUpdateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionAppStoreReviewScreenshotUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionAppStoreReviewScreenshotUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

