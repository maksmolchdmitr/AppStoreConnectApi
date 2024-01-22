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

// SubscriptionOfferCodePricesResponseIncludedInner - struct for SubscriptionOfferCodePricesResponseIncludedInner
type SubscriptionOfferCodePricesResponseIncludedInner struct {
	SubscriptionPricePoint *SubscriptionPricePoint
	Territory *Territory
}

// SubscriptionPricePointAsSubscriptionOfferCodePricesResponseIncludedInner is a convenience function that returns SubscriptionPricePoint wrapped in SubscriptionOfferCodePricesResponseIncludedInner
func SubscriptionPricePointAsSubscriptionOfferCodePricesResponseIncludedInner(v *SubscriptionPricePoint) SubscriptionOfferCodePricesResponseIncludedInner {
	return SubscriptionOfferCodePricesResponseIncludedInner{
		SubscriptionPricePoint: v,
	}
}

// TerritoryAsSubscriptionOfferCodePricesResponseIncludedInner is a convenience function that returns Territory wrapped in SubscriptionOfferCodePricesResponseIncludedInner
func TerritoryAsSubscriptionOfferCodePricesResponseIncludedInner(v *Territory) SubscriptionOfferCodePricesResponseIncludedInner {
	return SubscriptionOfferCodePricesResponseIncludedInner{
		Territory: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *SubscriptionOfferCodePricesResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
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
		dst.SubscriptionPricePoint = nil
		dst.Territory = nil

		return fmt.Errorf("data matches more than one schema in oneOf(SubscriptionOfferCodePricesResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(SubscriptionOfferCodePricesResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src SubscriptionOfferCodePricesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.SubscriptionPricePoint != nil {
		return json.Marshal(&src.SubscriptionPricePoint)
	}

	if src.Territory != nil {
		return json.Marshal(&src.Territory)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *SubscriptionOfferCodePricesResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
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

type NullableSubscriptionOfferCodePricesResponseIncludedInner struct {
	value *SubscriptionOfferCodePricesResponseIncludedInner
	isSet bool
}

func (v NullableSubscriptionOfferCodePricesResponseIncludedInner) Get() *SubscriptionOfferCodePricesResponseIncludedInner {
	return v.value
}

func (v *NullableSubscriptionOfferCodePricesResponseIncludedInner) Set(val *SubscriptionOfferCodePricesResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionOfferCodePricesResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionOfferCodePricesResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionOfferCodePricesResponseIncludedInner(val *SubscriptionOfferCodePricesResponseIncludedInner) *NullableSubscriptionOfferCodePricesResponseIncludedInner {
	return &NullableSubscriptionOfferCodePricesResponseIncludedInner{value: val, isSet: true}
}

func (v NullableSubscriptionOfferCodePricesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionOfferCodePricesResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


