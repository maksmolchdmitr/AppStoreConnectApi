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

// checks if the InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2 type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2{}

// InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2 struct for InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2
type InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2 struct {
	Data *AppRelationshipsInAppPurchasesDataInner `json:"data,omitempty"`
}

// NewInAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2 instantiates a new InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2() *InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2 {
	this := InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2{}
	return &this
}

// NewInAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2WithDefaults instantiates a new InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2WithDefaults() *InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2 {
	this := InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2) GetData() AppRelationshipsInAppPurchasesDataInner {
	if o == nil || IsNil(o.Data) {
		var ret AppRelationshipsInAppPurchasesDataInner
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2) GetDataOk() (*AppRelationshipsInAppPurchasesDataInner, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppRelationshipsInAppPurchasesDataInner and assigns it to the Data field.
func (o *InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2) SetData(v AppRelationshipsInAppPurchasesDataInner) {
	o.Data = &v
}

func (o InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2 struct {
	value *InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2
	isSet bool
}

func (v NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2) Get() *InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2 {
	return v.value
}

func (v *NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2) Set(val *InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2(val *InAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2) *NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2 {
	return &NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2{value: val, isSet: true}
}

func (v NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchasePriceInlineCreateRelationshipsInAppPurchaseV2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
