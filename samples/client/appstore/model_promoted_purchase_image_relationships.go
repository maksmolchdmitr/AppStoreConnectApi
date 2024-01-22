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

// checks if the PromotedPurchaseImageRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseImageRelationships{}

// PromotedPurchaseImageRelationships struct for PromotedPurchaseImageRelationships
type PromotedPurchaseImageRelationships struct {
	PromotedPurchase *InAppPurchaseV2RelationshipsPromotedPurchase `json:"promotedPurchase,omitempty"`
}

// NewPromotedPurchaseImageRelationships instantiates a new PromotedPurchaseImageRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseImageRelationships() *PromotedPurchaseImageRelationships {
	this := PromotedPurchaseImageRelationships{}
	return &this
}

// NewPromotedPurchaseImageRelationshipsWithDefaults instantiates a new PromotedPurchaseImageRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseImageRelationshipsWithDefaults() *PromotedPurchaseImageRelationships {
	this := PromotedPurchaseImageRelationships{}
	return &this
}

// GetPromotedPurchase returns the PromotedPurchase field value if set, zero value otherwise.
func (o *PromotedPurchaseImageRelationships) GetPromotedPurchase() InAppPurchaseV2RelationshipsPromotedPurchase {
	if o == nil || IsNil(o.PromotedPurchase) {
		var ret InAppPurchaseV2RelationshipsPromotedPurchase
		return ret
	}
	return *o.PromotedPurchase
}

// GetPromotedPurchaseOk returns a tuple with the PromotedPurchase field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseImageRelationships) GetPromotedPurchaseOk() (*InAppPurchaseV2RelationshipsPromotedPurchase, bool) {
	if o == nil || IsNil(o.PromotedPurchase) {
		return nil, false
	}
	return o.PromotedPurchase, true
}

// HasPromotedPurchase returns a boolean if a field has been set.
func (o *PromotedPurchaseImageRelationships) HasPromotedPurchase() bool {
	if o != nil && !IsNil(o.PromotedPurchase) {
		return true
	}

	return false
}

// SetPromotedPurchase gets a reference to the given InAppPurchaseV2RelationshipsPromotedPurchase and assigns it to the PromotedPurchase field.
func (o *PromotedPurchaseImageRelationships) SetPromotedPurchase(v InAppPurchaseV2RelationshipsPromotedPurchase) {
	o.PromotedPurchase = &v
}

func (o PromotedPurchaseImageRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseImageRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PromotedPurchase) {
		toSerialize["promotedPurchase"] = o.PromotedPurchase
	}
	return toSerialize, nil
}

type NullablePromotedPurchaseImageRelationships struct {
	value *PromotedPurchaseImageRelationships
	isSet bool
}

func (v NullablePromotedPurchaseImageRelationships) Get() *PromotedPurchaseImageRelationships {
	return v.value
}

func (v *NullablePromotedPurchaseImageRelationships) Set(val *PromotedPurchaseImageRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseImageRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseImageRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseImageRelationships(val *PromotedPurchaseImageRelationships) *NullablePromotedPurchaseImageRelationships {
	return &NullablePromotedPurchaseImageRelationships{value: val, isSet: true}
}

func (v NullablePromotedPurchaseImageRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseImageRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


