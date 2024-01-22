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

// checks if the SubscriptionGroupSubmissionCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGroupSubmissionCreateRequest{}

// SubscriptionGroupSubmissionCreateRequest struct for SubscriptionGroupSubmissionCreateRequest
type SubscriptionGroupSubmissionCreateRequest struct {
	Data SubscriptionGroupSubmissionCreateRequestData `json:"data"`
}

// NewSubscriptionGroupSubmissionCreateRequest instantiates a new SubscriptionGroupSubmissionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGroupSubmissionCreateRequest(data SubscriptionGroupSubmissionCreateRequestData) *SubscriptionGroupSubmissionCreateRequest {
	this := SubscriptionGroupSubmissionCreateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionGroupSubmissionCreateRequestWithDefaults instantiates a new SubscriptionGroupSubmissionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGroupSubmissionCreateRequestWithDefaults() *SubscriptionGroupSubmissionCreateRequest {
	this := SubscriptionGroupSubmissionCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionGroupSubmissionCreateRequest) GetData() SubscriptionGroupSubmissionCreateRequestData {
	if o == nil {
		var ret SubscriptionGroupSubmissionCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupSubmissionCreateRequest) GetDataOk() (*SubscriptionGroupSubmissionCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionGroupSubmissionCreateRequest) SetData(v SubscriptionGroupSubmissionCreateRequestData) {
	o.Data = v
}

func (o SubscriptionGroupSubmissionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGroupSubmissionCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSubscriptionGroupSubmissionCreateRequest struct {
	value *SubscriptionGroupSubmissionCreateRequest
	isSet bool
}

func (v NullableSubscriptionGroupSubmissionCreateRequest) Get() *SubscriptionGroupSubmissionCreateRequest {
	return v.value
}

func (v *NullableSubscriptionGroupSubmissionCreateRequest) Set(val *SubscriptionGroupSubmissionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGroupSubmissionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGroupSubmissionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGroupSubmissionCreateRequest(val *SubscriptionGroupSubmissionCreateRequest) *NullableSubscriptionGroupSubmissionCreateRequest {
	return &NullableSubscriptionGroupSubmissionCreateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionGroupSubmissionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGroupSubmissionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


