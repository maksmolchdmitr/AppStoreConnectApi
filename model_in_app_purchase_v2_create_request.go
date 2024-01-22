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

// checks if the InAppPurchaseV2CreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseV2CreateRequest{}

// InAppPurchaseV2CreateRequest struct for InAppPurchaseV2CreateRequest
type InAppPurchaseV2CreateRequest struct {
	Data InAppPurchaseV2CreateRequestData `json:"data"`
}

// NewInAppPurchaseV2CreateRequest instantiates a new InAppPurchaseV2CreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseV2CreateRequest(data InAppPurchaseV2CreateRequestData) *InAppPurchaseV2CreateRequest {
	this := InAppPurchaseV2CreateRequest{}
	this.Data = data
	return &this
}

// NewInAppPurchaseV2CreateRequestWithDefaults instantiates a new InAppPurchaseV2CreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseV2CreateRequestWithDefaults() *InAppPurchaseV2CreateRequest {
	this := InAppPurchaseV2CreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *InAppPurchaseV2CreateRequest) GetData() InAppPurchaseV2CreateRequestData {
	if o == nil {
		var ret InAppPurchaseV2CreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseV2CreateRequest) GetDataOk() (*InAppPurchaseV2CreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *InAppPurchaseV2CreateRequest) SetData(v InAppPurchaseV2CreateRequestData) {
	o.Data = v
}

func (o InAppPurchaseV2CreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseV2CreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableInAppPurchaseV2CreateRequest struct {
	value *InAppPurchaseV2CreateRequest
	isSet bool
}

func (v NullableInAppPurchaseV2CreateRequest) Get() *InAppPurchaseV2CreateRequest {
	return v.value
}

func (v *NullableInAppPurchaseV2CreateRequest) Set(val *InAppPurchaseV2CreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseV2CreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseV2CreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseV2CreateRequest(val *InAppPurchaseV2CreateRequest) *NullableInAppPurchaseV2CreateRequest {
	return &NullableInAppPurchaseV2CreateRequest{value: val, isSet: true}
}

func (v NullableInAppPurchaseV2CreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseV2CreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
