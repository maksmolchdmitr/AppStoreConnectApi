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

// AppCustomProductPageCreateRequestIncludedInner - struct for AppCustomProductPageCreateRequestIncludedInner
type AppCustomProductPageCreateRequestIncludedInner struct {
	AppCustomProductPageLocalizationInlineCreate *AppCustomProductPageLocalizationInlineCreate
	AppCustomProductPageVersionInlineCreate *AppCustomProductPageVersionInlineCreate
}

// AppCustomProductPageLocalizationInlineCreateAsAppCustomProductPageCreateRequestIncludedInner is a convenience function that returns AppCustomProductPageLocalizationInlineCreate wrapped in AppCustomProductPageCreateRequestIncludedInner
func AppCustomProductPageLocalizationInlineCreateAsAppCustomProductPageCreateRequestIncludedInner(v *AppCustomProductPageLocalizationInlineCreate) AppCustomProductPageCreateRequestIncludedInner {
	return AppCustomProductPageCreateRequestIncludedInner{
		AppCustomProductPageLocalizationInlineCreate: v,
	}
}

// AppCustomProductPageVersionInlineCreateAsAppCustomProductPageCreateRequestIncludedInner is a convenience function that returns AppCustomProductPageVersionInlineCreate wrapped in AppCustomProductPageCreateRequestIncludedInner
func AppCustomProductPageVersionInlineCreateAsAppCustomProductPageCreateRequestIncludedInner(v *AppCustomProductPageVersionInlineCreate) AppCustomProductPageCreateRequestIncludedInner {
	return AppCustomProductPageCreateRequestIncludedInner{
		AppCustomProductPageVersionInlineCreate: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppCustomProductPageCreateRequestIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppCustomProductPageLocalizationInlineCreate
	err = newStrictDecoder(data).Decode(&dst.AppCustomProductPageLocalizationInlineCreate)
	if err == nil {
		jsonAppCustomProductPageLocalizationInlineCreate, _ := json.Marshal(dst.AppCustomProductPageLocalizationInlineCreate)
		if string(jsonAppCustomProductPageLocalizationInlineCreate) == "{}" { // empty struct
			dst.AppCustomProductPageLocalizationInlineCreate = nil
		} else {
			match++
		}
	} else {
		dst.AppCustomProductPageLocalizationInlineCreate = nil
	}

	// try to unmarshal data into AppCustomProductPageVersionInlineCreate
	err = newStrictDecoder(data).Decode(&dst.AppCustomProductPageVersionInlineCreate)
	if err == nil {
		jsonAppCustomProductPageVersionInlineCreate, _ := json.Marshal(dst.AppCustomProductPageVersionInlineCreate)
		if string(jsonAppCustomProductPageVersionInlineCreate) == "{}" { // empty struct
			dst.AppCustomProductPageVersionInlineCreate = nil
		} else {
			match++
		}
	} else {
		dst.AppCustomProductPageVersionInlineCreate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppCustomProductPageLocalizationInlineCreate = nil
		dst.AppCustomProductPageVersionInlineCreate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppCustomProductPageCreateRequestIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppCustomProductPageCreateRequestIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppCustomProductPageCreateRequestIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AppCustomProductPageLocalizationInlineCreate != nil {
		return json.Marshal(&src.AppCustomProductPageLocalizationInlineCreate)
	}

	if src.AppCustomProductPageVersionInlineCreate != nil {
		return json.Marshal(&src.AppCustomProductPageVersionInlineCreate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppCustomProductPageCreateRequestIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AppCustomProductPageLocalizationInlineCreate != nil {
		return obj.AppCustomProductPageLocalizationInlineCreate
	}

	if obj.AppCustomProductPageVersionInlineCreate != nil {
		return obj.AppCustomProductPageVersionInlineCreate
	}

	// all schemas are nil
	return nil
}

type NullableAppCustomProductPageCreateRequestIncludedInner struct {
	value *AppCustomProductPageCreateRequestIncludedInner
	isSet bool
}

func (v NullableAppCustomProductPageCreateRequestIncludedInner) Get() *AppCustomProductPageCreateRequestIncludedInner {
	return v.value
}

func (v *NullableAppCustomProductPageCreateRequestIncludedInner) Set(val *AppCustomProductPageCreateRequestIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageCreateRequestIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageCreateRequestIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageCreateRequestIncludedInner(val *AppCustomProductPageCreateRequestIncludedInner) *NullableAppCustomProductPageCreateRequestIncludedInner {
	return &NullableAppCustomProductPageCreateRequestIncludedInner{value: val, isSet: true}
}

func (v NullableAppCustomProductPageCreateRequestIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageCreateRequestIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


