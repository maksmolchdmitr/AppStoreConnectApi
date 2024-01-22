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

// checks if the SubscriptionOfferCodeOneTimeUseCodeUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeOneTimeUseCodeUpdateRequest{}

// SubscriptionOfferCodeOneTimeUseCodeUpdateRequest struct for SubscriptionOfferCodeOneTimeUseCodeUpdateRequest
type SubscriptionOfferCodeOneTimeUseCodeUpdateRequest struct {
	Data SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData `json:"data"`
}

// NewSubscriptionOfferCodeOneTimeUseCodeUpdateRequest instantiates a new SubscriptionOfferCodeOneTimeUseCodeUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeOneTimeUseCodeUpdateRequest(data SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) *SubscriptionOfferCodeOneTimeUseCodeUpdateRequest {
	this := SubscriptionOfferCodeOneTimeUseCodeUpdateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionOfferCodeOneTimeUseCodeUpdateRequestWithDefaults instantiates a new SubscriptionOfferCodeOneTimeUseCodeUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeOneTimeUseCodeUpdateRequestWithDefaults() *SubscriptionOfferCodeOneTimeUseCodeUpdateRequest {
	this := SubscriptionOfferCodeOneTimeUseCodeUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionOfferCodeOneTimeUseCodeUpdateRequest) GetData() SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData {
	if o == nil {
		var ret SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeOneTimeUseCodeUpdateRequest) GetDataOk() (*SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionOfferCodeOneTimeUseCodeUpdateRequest) SetData(v SubscriptionOfferCodeOneTimeUseCodeUpdateRequestData) {
	o.Data = v
}

func (o SubscriptionOfferCodeOneTimeUseCodeUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeOneTimeUseCodeUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequest struct {
	value *SubscriptionOfferCodeOneTimeUseCodeUpdateRequest
	isSet bool
}

func (v NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequest) Get() *SubscriptionOfferCodeOneTimeUseCodeUpdateRequest {
	return v.value
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequest) Set(val *SubscriptionOfferCodeOneTimeUseCodeUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequest(val *SubscriptionOfferCodeOneTimeUseCodeUpdateRequest) *NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequest {
	return &NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeOneTimeUseCodeUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}