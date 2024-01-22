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

// checks if the SubscriptionSubmissionCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionSubmissionCreateRequest{}

// SubscriptionSubmissionCreateRequest struct for SubscriptionSubmissionCreateRequest
type SubscriptionSubmissionCreateRequest struct {
	Data SubscriptionSubmissionCreateRequestData `json:"data"`
}

// NewSubscriptionSubmissionCreateRequest instantiates a new SubscriptionSubmissionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionSubmissionCreateRequest(data SubscriptionSubmissionCreateRequestData) *SubscriptionSubmissionCreateRequest {
	this := SubscriptionSubmissionCreateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionSubmissionCreateRequestWithDefaults instantiates a new SubscriptionSubmissionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionSubmissionCreateRequestWithDefaults() *SubscriptionSubmissionCreateRequest {
	this := SubscriptionSubmissionCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionSubmissionCreateRequest) GetData() SubscriptionSubmissionCreateRequestData {
	if o == nil {
		var ret SubscriptionSubmissionCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionSubmissionCreateRequest) GetDataOk() (*SubscriptionSubmissionCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionSubmissionCreateRequest) SetData(v SubscriptionSubmissionCreateRequestData) {
	o.Data = v
}

func (o SubscriptionSubmissionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionSubmissionCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSubscriptionSubmissionCreateRequest struct {
	value *SubscriptionSubmissionCreateRequest
	isSet bool
}

func (v NullableSubscriptionSubmissionCreateRequest) Get() *SubscriptionSubmissionCreateRequest {
	return v.value
}

func (v *NullableSubscriptionSubmissionCreateRequest) Set(val *SubscriptionSubmissionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionSubmissionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionSubmissionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionSubmissionCreateRequest(val *SubscriptionSubmissionCreateRequest) *NullableSubscriptionSubmissionCreateRequest {
	return &NullableSubscriptionSubmissionCreateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionSubmissionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionSubmissionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


