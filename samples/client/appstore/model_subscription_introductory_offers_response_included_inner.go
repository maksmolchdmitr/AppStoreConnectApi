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

// SubscriptionIntroductoryOffersResponseIncludedInner - struct for SubscriptionIntroductoryOffersResponseIncludedInner
type SubscriptionIntroductoryOffersResponseIncludedInner struct {
	Subscription *Subscription
	SubscriptionPricePoint *SubscriptionPricePoint
	Territory *Territory
}

// SubscriptionAsSubscriptionIntroductoryOffersResponseIncludedInner is a convenience function that returns Subscription wrapped in SubscriptionIntroductoryOffersResponseIncludedInner
func SubscriptionAsSubscriptionIntroductoryOffersResponseIncludedInner(v *Subscription) SubscriptionIntroductoryOffersResponseIncludedInner {
	return SubscriptionIntroductoryOffersResponseIncludedInner{
		Subscription: v,
	}
}

// SubscriptionPricePointAsSubscriptionIntroductoryOffersResponseIncludedInner is a convenience function that returns SubscriptionPricePoint wrapped in SubscriptionIntroductoryOffersResponseIncludedInner
func SubscriptionPricePointAsSubscriptionIntroductoryOffersResponseIncludedInner(v *SubscriptionPricePoint) SubscriptionIntroductoryOffersResponseIncludedInner {
	return SubscriptionIntroductoryOffersResponseIncludedInner{
		SubscriptionPricePoint: v,
	}
}

// TerritoryAsSubscriptionIntroductoryOffersResponseIncludedInner is a convenience function that returns Territory wrapped in SubscriptionIntroductoryOffersResponseIncludedInner
func TerritoryAsSubscriptionIntroductoryOffersResponseIncludedInner(v *Territory) SubscriptionIntroductoryOffersResponseIncludedInner {
	return SubscriptionIntroductoryOffersResponseIncludedInner{
		Territory: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubscriptionIntroductoryOffersResponseIncludedInner) UnmarshalJSON(data []byte) error {
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

	// try to unmarshal data into SubscriptionPricePoint
	err = newStrictDecoder(data).Decode(&dst.SubscriptionPricePoint)
	if err == nil {
		jsonSubscriptionPricePoint, _ := json.Marshal(dst.SubscriptionPricePoint)
		if string(jsonSubscriptionPricePoint) == "{}" { // empty struct
			dst.SubscriptionPricePoint = nil
		} else {
			match++
		}
	} else {
		dst.SubscriptionPricePoint = nil
	}

	// try to unmarshal data into Territory
	err = newStrictDecoder(data).Decode(&dst.Territory)
	if err == nil {
		jsonTerritory, _ := json.Marshal(dst.Territory)
		if string(jsonTerritory) == "{}" { // empty struct
			dst.Territory = nil
		} else {
			match++
		}
	} else {
		dst.Territory = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.Subscription = nil
		dst.SubscriptionPricePoint = nil
		dst.Territory = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SubscriptionIntroductoryOffersResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SubscriptionIntroductoryOffersResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubscriptionIntroductoryOffersResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.Subscription != nil {
		return json.Marshal(&src.Subscription)
	}

	if src.SubscriptionPricePoint != nil {
		return json.Marshal(&src.SubscriptionPricePoint)
	}

	if src.Territory != nil {
		return json.Marshal(&src.Territory)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubscriptionIntroductoryOffersResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.Subscription != nil {
		return obj.Subscription
	}

	if obj.SubscriptionPricePoint != nil {
		return obj.SubscriptionPricePoint
	}

	if obj.Territory != nil {
		return obj.Territory
	}

	// all schemas are nil
	return nil
}

type NullableSubscriptionIntroductoryOffersResponseIncludedInner struct {
	value *SubscriptionIntroductoryOffersResponseIncludedInner
	isSet bool
}

func (v NullableSubscriptionIntroductoryOffersResponseIncludedInner) Get() *SubscriptionIntroductoryOffersResponseIncludedInner {
	return v.value
}

func (v *NullableSubscriptionIntroductoryOffersResponseIncludedInner) Set(val *SubscriptionIntroductoryOffersResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionIntroductoryOffersResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionIntroductoryOffersResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionIntroductoryOffersResponseIncludedInner(val *SubscriptionIntroductoryOffersResponseIncludedInner) *NullableSubscriptionIntroductoryOffersResponseIncludedInner {
	return &NullableSubscriptionIntroductoryOffersResponseIncludedInner{value: val, isSet: true}
}

func (v NullableSubscriptionIntroductoryOffersResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionIntroductoryOffersResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


