/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// SubscriptionGroupsResponseIncludedInner - struct for SubscriptionGroupsResponseIncludedInner
type SubscriptionGroupsResponseIncludedInner struct {
	Subscription                  *Subscription
	SubscriptionGroupLocalization *SubscriptionGroupLocalization
}

// SubscriptionAsSubscriptionGroupsResponseIncludedInner is a convenience function that returns Subscription wrapped in SubscriptionGroupsResponseIncludedInner
func SubscriptionAsSubscriptionGroupsResponseIncludedInner(v *Subscription) SubscriptionGroupsResponseIncludedInner {
	return SubscriptionGroupsResponseIncludedInner{
		Subscription: v,
	}
}

// SubscriptionGroupLocalizationAsSubscriptionGroupsResponseIncludedInner is a convenience function that returns SubscriptionGroupLocalization wrapped in SubscriptionGroupsResponseIncludedInner
func SubscriptionGroupLocalizationAsSubscriptionGroupsResponseIncludedInner(v *SubscriptionGroupLocalization) SubscriptionGroupsResponseIncludedInner {
	return SubscriptionGroupsResponseIncludedInner{
		SubscriptionGroupLocalization: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubscriptionGroupsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into Subscription
	err = newStrictDecoder(data).Decode(&dst.Subscription)
	if err == nil {
		jsonSubscription, _ := json.Marshal(dst.Subscription)
		if string(jsonSubscription) == "{}" { // empty struct
			dst.Subscription = nil
		} else {
			match++
		}
	} else {
		dst.Subscription = nil
	}

	// try to unmarshal data into SubscriptionGroupLocalization
	err = newStrictDecoder(data).Decode(&dst.SubscriptionGroupLocalization)
	if err == nil {
		jsonSubscriptionGroupLocalization, _ := json.Marshal(dst.SubscriptionGroupLocalization)
		if string(jsonSubscriptionGroupLocalization) == "{}" { // empty struct
			dst.SubscriptionGroupLocalization = nil
		} else {
			match++
		}
	} else {
		dst.SubscriptionGroupLocalization = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Subscription = nil
		dst.SubscriptionGroupLocalization = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SubscriptionGroupsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SubscriptionGroupsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubscriptionGroupsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.Subscription != nil {
		return json.Marshal(&src.Subscription)
	}

	if src.SubscriptionGroupLocalization != nil {
		return json.Marshal(&src.SubscriptionGroupLocalization)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubscriptionGroupsResponseIncludedInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.Subscription != nil {
		return obj.Subscription
	}

	if obj.SubscriptionGroupLocalization != nil {
		return obj.SubscriptionGroupLocalization
	}

	// all schemas are nil
	return nil
}

type NullableSubscriptionGroupsResponseIncludedInner struct {
	value *SubscriptionGroupsResponseIncludedInner
	isSet bool
}

func (v NullableSubscriptionGroupsResponseIncludedInner) Get() *SubscriptionGroupsResponseIncludedInner {
	return v.value
}

func (v *NullableSubscriptionGroupsResponseIncludedInner) Set(val *SubscriptionGroupsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGroupsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGroupsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGroupsResponseIncludedInner(val *SubscriptionGroupsResponseIncludedInner) *NullableSubscriptionGroupsResponseIncludedInner {
	return &NullableSubscriptionGroupsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableSubscriptionGroupsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGroupsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}