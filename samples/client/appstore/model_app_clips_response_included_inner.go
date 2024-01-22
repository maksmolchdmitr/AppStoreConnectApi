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

// AppClipsResponseIncludedInner - struct for AppClipsResponseIncludedInner
type AppClipsResponseIncludedInner struct {
	App *App
	AppClipDefaultExperience *AppClipDefaultExperience
}

// AppAsAppClipsResponseIncludedInner is a convenience function that returns App wrapped in AppClipsResponseIncludedInner
func AppAsAppClipsResponseIncludedInner(v *App) AppClipsResponseIncludedInner {
	return AppClipsResponseIncludedInner{
		App: v,
	}
}

// AppClipDefaultExperienceAsAppClipsResponseIncludedInner is a convenience function that returns AppClipDefaultExperience wrapped in AppClipsResponseIncludedInner
func AppClipDefaultExperienceAsAppClipsResponseIncludedInner(v *AppClipDefaultExperience) AppClipsResponseIncludedInner {
	return AppClipsResponseIncludedInner{
		AppClipDefaultExperience: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppClipsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into App
	err = newStrictDecoder(data).Decode(&dst.App)
	if err == nil {
		jsonApp, _ := json.Marshal(dst.App)
		if string(jsonApp) == "{}" { // empty struct
			dst.App = nil
		} else {
			match++
		}
	} else {
		dst.App = nil
	}

	// try to unmarshal data into AppClipDefaultExperience
	err = newStrictDecoder(data).Decode(&dst.AppClipDefaultExperience)
	if err == nil {
		jsonAppClipDefaultExperience, _ := json.Marshal(dst.AppClipDefaultExperience)
		if string(jsonAppClipDefaultExperience) == "{}" { // empty struct
			dst.AppClipDefaultExperience = nil
		} else {
			match++
		}
	} else {
		dst.AppClipDefaultExperience = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.App = nil
		dst.AppClipDefaultExperience = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppClipsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppClipsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppClipsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.App != nil {
		return json.Marshal(&src.App)
	}

	if src.AppClipDefaultExperience != nil {
		return json.Marshal(&src.AppClipDefaultExperience)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppClipsResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.App != nil {
		return obj.App
	}

	if obj.AppClipDefaultExperience != nil {
		return obj.AppClipDefaultExperience
	}

	// all schemas are nil
	return nil
}

type NullableAppClipsResponseIncludedInner struct {
	value *AppClipsResponseIncludedInner
	isSet bool
}

func (v NullableAppClipsResponseIncludedInner) Get() *AppClipsResponseIncludedInner {
	return v.value
}

func (v *NullableAppClipsResponseIncludedInner) Set(val *AppClipsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipsResponseIncludedInner(val *AppClipsResponseIncludedInner) *NullableAppClipsResponseIncludedInner {
	return &NullableAppClipsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAppClipsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


