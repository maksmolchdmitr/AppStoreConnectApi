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

// checks if the SubscriptionLocalizationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionLocalizationCreateRequest{}

// SubscriptionLocalizationCreateRequest struct for SubscriptionLocalizationCreateRequest
type SubscriptionLocalizationCreateRequest struct {
	Data SubscriptionLocalizationCreateRequestData `json:"data"`
}

// NewSubscriptionLocalizationCreateRequest instantiates a new SubscriptionLocalizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionLocalizationCreateRequest(data SubscriptionLocalizationCreateRequestData) *SubscriptionLocalizationCreateRequest {
	this := SubscriptionLocalizationCreateRequest{}
	this.Data = data
	return &this
}

// NewSubscriptionLocalizationCreateRequestWithDefaults instantiates a new SubscriptionLocalizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionLocalizationCreateRequestWithDefaults() *SubscriptionLocalizationCreateRequest {
	this := SubscriptionLocalizationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *SubscriptionLocalizationCreateRequest) GetData() SubscriptionLocalizationCreateRequestData {
	if o == nil {
		var ret SubscriptionLocalizationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SubscriptionLocalizationCreateRequest) GetDataOk() (*SubscriptionLocalizationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SubscriptionLocalizationCreateRequest) SetData(v SubscriptionLocalizationCreateRequestData) {
	o.Data = v
}

func (o SubscriptionLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionLocalizationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableSubscriptionLocalizationCreateRequest struct {
	value *SubscriptionLocalizationCreateRequest
	isSet bool
}

func (v NullableSubscriptionLocalizationCreateRequest) Get() *SubscriptionLocalizationCreateRequest {
	return v.value
}

func (v *NullableSubscriptionLocalizationCreateRequest) Set(val *SubscriptionLocalizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionLocalizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionLocalizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionLocalizationCreateRequest(val *SubscriptionLocalizationCreateRequest) *NullableSubscriptionLocalizationCreateRequest {
	return &NullableSubscriptionLocalizationCreateRequest{value: val, isSet: true}
}

func (v NullableSubscriptionLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionLocalizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


