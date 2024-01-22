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

// checks if the SubscriptionAvailabilityRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionAvailabilityRelationships{}

// SubscriptionAvailabilityRelationships struct for SubscriptionAvailabilityRelationships
type SubscriptionAvailabilityRelationships struct {
	Subscription *PromotedPurchaseRelationshipsSubscription `json:"subscription,omitempty"`
	AvailableTerritories *AppAvailabilityRelationshipsAvailableTerritories `json:"availableTerritories,omitempty"`
}

// NewSubscriptionAvailabilityRelationships instantiates a new SubscriptionAvailabilityRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionAvailabilityRelationships() *SubscriptionAvailabilityRelationships {
	this := SubscriptionAvailabilityRelationships{}
	return &this
}

// NewSubscriptionAvailabilityRelationshipsWithDefaults instantiates a new SubscriptionAvailabilityRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionAvailabilityRelationshipsWithDefaults() *SubscriptionAvailabilityRelationships {
	this := SubscriptionAvailabilityRelationships{}
	return &this
}

// GetSubscription returns the Subscription field value if set, zero value otherwise.
func (o *SubscriptionAvailabilityRelationships) GetSubscription() PromotedPurchaseRelationshipsSubscription {
	if o == nil || IsNil(o.Subscription) {
		var ret PromotedPurchaseRelationshipsSubscription
		return ret
	}
	return *o.Subscription
}

// GetSubscriptionOk returns a tuple with the Subscription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAvailabilityRelationships) GetSubscriptionOk() (*PromotedPurchaseRelationshipsSubscription, bool) {
	if o == nil || IsNil(o.Subscription) {
		return nil, false
	}
	return o.Subscription, true
}

// HasSubscription returns a boolean if a field has been set.
func (o *SubscriptionAvailabilityRelationships) HasSubscription() bool {
	if o != nil && !IsNil(o.Subscription) {
		return true
	}

	return false
}

// SetSubscription gets a reference to the given PromotedPurchaseRelationshipsSubscription and assigns it to the Subscription field.
func (o *SubscriptionAvailabilityRelationships) SetSubscription(v PromotedPurchaseRelationshipsSubscription) {
	o.Subscription = &v
}

// GetAvailableTerritories returns the AvailableTerritories field value if set, zero value otherwise.
func (o *SubscriptionAvailabilityRelationships) GetAvailableTerritories() AppAvailabilityRelationshipsAvailableTerritories {
	if o == nil || IsNil(o.AvailableTerritories) {
		var ret AppAvailabilityRelationshipsAvailableTerritories
		return ret
	}
	return *o.AvailableTerritories
}

// GetAvailableTerritoriesOk returns a tuple with the AvailableTerritories field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionAvailabilityRelationships) GetAvailableTerritoriesOk() (*AppAvailabilityRelationshipsAvailableTerritories, bool) {
	if o == nil || IsNil(o.AvailableTerritories) {
		return nil, false
	}
	return o.AvailableTerritories, true
}

// HasAvailableTerritories returns a boolean if a field has been set.
func (o *SubscriptionAvailabilityRelationships) HasAvailableTerritories() bool {
	if o != nil && !IsNil(o.AvailableTerritories) {
		return true
	}

	return false
}

// SetAvailableTerritories gets a reference to the given AppAvailabilityRelationshipsAvailableTerritories and assigns it to the AvailableTerritories field.
func (o *SubscriptionAvailabilityRelationships) SetAvailableTerritories(v AppAvailabilityRelationshipsAvailableTerritories) {
	o.AvailableTerritories = &v
}

func (o SubscriptionAvailabilityRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionAvailabilityRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Subscription) {
		toSerialize["subscription"] = o.Subscription
	}
	if !IsNil(o.AvailableTerritories) {
		toSerialize["availableTerritories"] = o.AvailableTerritories
	}
	return toSerialize, nil
}

type NullableSubscriptionAvailabilityRelationships struct {
	value *SubscriptionAvailabilityRelationships
	isSet bool
}

func (v NullableSubscriptionAvailabilityRelationships) Get() *SubscriptionAvailabilityRelationships {
	return v.value
}

func (v *NullableSubscriptionAvailabilityRelationships) Set(val *SubscriptionAvailabilityRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionAvailabilityRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionAvailabilityRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionAvailabilityRelationships(val *SubscriptionAvailabilityRelationships) *NullableSubscriptionAvailabilityRelationships {
	return &NullableSubscriptionAvailabilityRelationships{value: val, isSet: true}
}

func (v NullableSubscriptionAvailabilityRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionAvailabilityRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


