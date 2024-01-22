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

// checks if the SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes{}

// SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes struct for SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes
type SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes struct {
	Active *bool `json:"active,omitempty"`
}

// NewSubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes instantiates a new SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes() *SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes {
	this := SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes{}
	return &this
}

// NewSubscriptionOfferCodeCustomCodeUpdateRequestDataAttributesWithDefaults instantiates a new SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeCustomCodeUpdateRequestDataAttributesWithDefaults() *SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes {
	this := SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes{}
	return &this
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) SetActive(v bool) {
	o.Active = &v
}

func (o SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes struct {
	value *SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes
	isSet bool
}

func (v NullableSubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) Get() *SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableSubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) Set(val *SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes(val *SubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) *NullableSubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes {
	return &NullableSubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeCustomCodeUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

