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

// checks if the PromotedPurchaseUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseUpdateRequest{}

// PromotedPurchaseUpdateRequest struct for PromotedPurchaseUpdateRequest
type PromotedPurchaseUpdateRequest struct {
	Data PromotedPurchaseUpdateRequestData `json:"data"`
}

// NewPromotedPurchaseUpdateRequest instantiates a new PromotedPurchaseUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseUpdateRequest(data PromotedPurchaseUpdateRequestData) *PromotedPurchaseUpdateRequest {
	this := PromotedPurchaseUpdateRequest{}
	this.Data = data
	return &this
}

// NewPromotedPurchaseUpdateRequestWithDefaults instantiates a new PromotedPurchaseUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseUpdateRequestWithDefaults() *PromotedPurchaseUpdateRequest {
	this := PromotedPurchaseUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *PromotedPurchaseUpdateRequest) GetData() PromotedPurchaseUpdateRequestData {
	if o == nil {
		var ret PromotedPurchaseUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseUpdateRequest) GetDataOk() (*PromotedPurchaseUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PromotedPurchaseUpdateRequest) SetData(v PromotedPurchaseUpdateRequestData) {
	o.Data = v
}

func (o PromotedPurchaseUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullablePromotedPurchaseUpdateRequest struct {
	value *PromotedPurchaseUpdateRequest
	isSet bool
}

func (v NullablePromotedPurchaseUpdateRequest) Get() *PromotedPurchaseUpdateRequest {
	return v.value
}

func (v *NullablePromotedPurchaseUpdateRequest) Set(val *PromotedPurchaseUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseUpdateRequest(val *PromotedPurchaseUpdateRequest) *NullablePromotedPurchaseUpdateRequest {
	return &NullablePromotedPurchaseUpdateRequest{value: val, isSet: true}
}

func (v NullablePromotedPurchaseUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
