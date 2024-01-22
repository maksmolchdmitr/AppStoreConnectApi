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

// checks if the SubscriptionOfferCodeCustomCodeRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionOfferCodeCustomCodeRelationships{}

// SubscriptionOfferCodeCustomCodeRelationships struct for SubscriptionOfferCodeCustomCodeRelationships
type SubscriptionOfferCodeCustomCodeRelationships struct {
	OfferCode *SubscriptionOfferCodeCustomCodeRelationshipsOfferCode `json:"offerCode,omitempty"`
}

// NewSubscriptionOfferCodeCustomCodeRelationships instantiates a new SubscriptionOfferCodeCustomCodeRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionOfferCodeCustomCodeRelationships() *SubscriptionOfferCodeCustomCodeRelationships {
	this := SubscriptionOfferCodeCustomCodeRelationships{}
	return &this
}

// NewSubscriptionOfferCodeCustomCodeRelationshipsWithDefaults instantiates a new SubscriptionOfferCodeCustomCodeRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionOfferCodeCustomCodeRelationshipsWithDefaults() *SubscriptionOfferCodeCustomCodeRelationships {
	this := SubscriptionOfferCodeCustomCodeRelationships{}
	return &this
}

// GetOfferCode returns the OfferCode field value if set, zero value otherwise.
func (o *SubscriptionOfferCodeCustomCodeRelationships) GetOfferCode() SubscriptionOfferCodeCustomCodeRelationshipsOfferCode {
	if o == nil || IsNil(o.OfferCode) {
		var ret SubscriptionOfferCodeCustomCodeRelationshipsOfferCode
		return ret
	}
	return *o.OfferCode
}

// GetOfferCodeOk returns a tuple with the OfferCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionOfferCodeCustomCodeRelationships) GetOfferCodeOk() (*SubscriptionOfferCodeCustomCodeRelationshipsOfferCode, bool) {
	if o == nil || IsNil(o.OfferCode) {
		return nil, false
	}
	return o.OfferCode, true
}

// HasOfferCode returns a boolean if a field has been set.
func (o *SubscriptionOfferCodeCustomCodeRelationships) HasOfferCode() bool {
	if o != nil && !IsNil(o.OfferCode) {
		return true
	}

	return false
}

// SetOfferCode gets a reference to the given SubscriptionOfferCodeCustomCodeRelationshipsOfferCode and assigns it to the OfferCode field.
func (o *SubscriptionOfferCodeCustomCodeRelationships) SetOfferCode(v SubscriptionOfferCodeCustomCodeRelationshipsOfferCode) {
	o.OfferCode = &v
}

func (o SubscriptionOfferCodeCustomCodeRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionOfferCodeCustomCodeRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.OfferCode) {
		toSerialize["offerCode"] = o.OfferCode
	}
	return toSerialize, nil
}

type NullableSubscriptionOfferCodeCustomCodeRelationships struct {
	value *SubscriptionOfferCodeCustomCodeRelationships
	isSet bool
}

func (v NullableSubscriptionOfferCodeCustomCodeRelationships) Get() *SubscriptionOfferCodeCustomCodeRelationships {
	return v.value
}

func (v *NullableSubscriptionOfferCodeCustomCodeRelationships) Set(val *SubscriptionOfferCodeCustomCodeRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodeCustomCodeRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodeCustomCodeRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodeCustomCodeRelationships(val *SubscriptionOfferCodeCustomCodeRelationships) *NullableSubscriptionOfferCodeCustomCodeRelationships {
	return &NullableSubscriptionOfferCodeCustomCodeRelationships{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodeCustomCodeRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodeCustomCodeRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


