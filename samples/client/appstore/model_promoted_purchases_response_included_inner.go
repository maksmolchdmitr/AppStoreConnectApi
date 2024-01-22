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

// PromotedPurchasesResponseIncludedInner - struct for PromotedPurchasesResponseIncludedInner
type PromotedPurchasesResponseIncludedInner struct {
	InAppPurchaseV2 *InAppPurchaseV2
	PromotedPurchaseImage *PromotedPurchaseImage
	Subscription *Subscription
}

// InAppPurchaseV2AsPromotedPurchasesResponseIncludedInner is a convenience function that returns InAppPurchaseV2 wrapped in PromotedPurchasesResponseIncludedInner
func InAppPurchaseV2AsPromotedPurchasesResponseIncludedInner(v *InAppPurchaseV2) PromotedPurchasesResponseIncludedInner {
	return PromotedPurchasesResponseIncludedInner{
		InAppPurchaseV2: v,
	}
}

// PromotedPurchaseImageAsPromotedPurchasesResponseIncludedInner is a convenience function that returns PromotedPurchaseImage wrapped in PromotedPurchasesResponseIncludedInner
func PromotedPurchaseImageAsPromotedPurchasesResponseIncludedInner(v *PromotedPurchaseImage) PromotedPurchasesResponseIncludedInner {
	return PromotedPurchasesResponseIncludedInner{
		PromotedPurchaseImage: v,
	}
}

// SubscriptionAsPromotedPurchasesResponseIncludedInner is a convenience function that returns Subscription wrapped in PromotedPurchasesResponseIncludedInner
func SubscriptionAsPromotedPurchasesResponseIncludedInner(v *Subscription) PromotedPurchasesResponseIncludedInner {
	return PromotedPurchasesResponseIncludedInner{
		Subscription: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *PromotedPurchasesResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into InAppPurchaseV2
	err = newStrictDecoder(data).Decode(&dst.InAppPurchaseV2)
	if err == nil {
		jsonInAppPurchaseV2, _ := json.Marshal(dst.InAppPurchaseV2)
		if string(jsonInAppPurchaseV2) == "{}" { // empty struct
			dst.InAppPurchaseV2 = nil
		} else {
			match++
		}
	} else {
		dst.InAppPurchaseV2 = nil
	}

	// try to unmarshal data into PromotedPurchaseImage
	err = newStrictDecoder(data).Decode(&dst.PromotedPurchaseImage)
	if err == nil {
		jsonPromotedPurchaseImage, _ := json.Marshal(dst.PromotedPurchaseImage)
		if string(jsonPromotedPurchaseImage) == "{}" { // empty struct
			dst.PromotedPurchaseImage = nil
		} else {
			match++
		}
	} else {
		dst.PromotedPurchaseImage = nil
	}

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

	if match > 1 { // more than 1 match
		// reset to nil
		dst.InAppPurchaseV2 = nil
		dst.PromotedPurchaseImage = nil
		dst.Subscription = nil

		return fmt.Errorf("data matches more than one schema in oneOf(PromotedPurchasesResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(PromotedPurchasesResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src PromotedPurchasesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.InAppPurchaseV2 != nil {
		return json.Marshal(&src.InAppPurchaseV2)
	}

	if src.PromotedPurchaseImage != nil {
		return json.Marshal(&src.PromotedPurchaseImage)
	}

	if src.Subscription != nil {
		return json.Marshal(&src.Subscription)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *PromotedPurchasesResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.InAppPurchaseV2 != nil {
		return obj.InAppPurchaseV2
	}

	if obj.PromotedPurchaseImage != nil {
		return obj.PromotedPurchaseImage
	}

	if obj.Subscription != nil {
		return obj.Subscription
	}

	// all schemas are nil
	return nil
}

type NullablePromotedPurchasesResponseIncludedInner struct {
	value *PromotedPurchasesResponseIncludedInner
	isSet bool
}

func (v NullablePromotedPurchasesResponseIncludedInner) Get() *PromotedPurchasesResponseIncludedInner {
	return v.value
}

func (v *NullablePromotedPurchasesResponseIncludedInner) Set(val *PromotedPurchasesResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchasesResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchasesResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchasesResponseIncludedInner(val *PromotedPurchasesResponseIncludedInner) *NullablePromotedPurchasesResponseIncludedInner {
	return &NullablePromotedPurchasesResponseIncludedInner{value: val, isSet: true}
}

func (v NullablePromotedPurchasesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchasesResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

