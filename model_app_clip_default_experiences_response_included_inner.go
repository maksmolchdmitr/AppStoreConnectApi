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

// AppClipDefaultExperiencesResponseIncludedInner - struct for AppClipDefaultExperiencesResponseIncludedInner
type AppClipDefaultExperiencesResponseIncludedInner struct {
	AppClip                              *AppClip
	AppClipAppStoreReviewDetail          *AppClipAppStoreReviewDetail
	AppClipDefaultExperienceLocalization *AppClipDefaultExperienceLocalization
	AppStoreVersion                      *AppStoreVersion
}

// AppClipAsAppClipDefaultExperiencesResponseIncludedInner is a convenience function that returns AppClip wrapped in AppClipDefaultExperiencesResponseIncludedInner
func AppClipAsAppClipDefaultExperiencesResponseIncludedInner(v *AppClip) AppClipDefaultExperiencesResponseIncludedInner {
	return AppClipDefaultExperiencesResponseIncludedInner{
		AppClip: v,
	}
}

// AppClipAppStoreReviewDetailAsAppClipDefaultExperiencesResponseIncludedInner is a convenience function that returns AppClipAppStoreReviewDetail wrapped in AppClipDefaultExperiencesResponseIncludedInner
func AppClipAppStoreReviewDetailAsAppClipDefaultExperiencesResponseIncludedInner(v *AppClipAppStoreReviewDetail) AppClipDefaultExperiencesResponseIncludedInner {
	return AppClipDefaultExperiencesResponseIncludedInner{
		AppClipAppStoreReviewDetail: v,
	}
}

// AppClipDefaultExperienceLocalizationAsAppClipDefaultExperiencesResponseIncludedInner is a convenience function that returns AppClipDefaultExperienceLocalization wrapped in AppClipDefaultExperiencesResponseIncludedInner
func AppClipDefaultExperienceLocalizationAsAppClipDefaultExperiencesResponseIncludedInner(v *AppClipDefaultExperienceLocalization) AppClipDefaultExperiencesResponseIncludedInner {
	return AppClipDefaultExperiencesResponseIncludedInner{
		AppClipDefaultExperienceLocalization: v,
	}
}

// AppStoreVersionAsAppClipDefaultExperiencesResponseIncludedInner is a convenience function that returns AppStoreVersion wrapped in AppClipDefaultExperiencesResponseIncludedInner
func AppStoreVersionAsAppClipDefaultExperiencesResponseIncludedInner(v *AppStoreVersion) AppClipDefaultExperiencesResponseIncludedInner {
	return AppClipDefaultExperiencesResponseIncludedInner{
		AppStoreVersion: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppClipDefaultExperiencesResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AppClip
	err = newStrictDecoder(data).Decode(&dst.AppClip)
	if err == nil {
		jsonAppClip, _ := json.Marshal(dst.AppClip)
		if string(jsonAppClip) == "{}" { // empty struct
			dst.AppClip = nil
		} else {
			match++
		}
	} else {
		dst.AppClip = nil
	}

	// try to unmarshal data into AppClipAppStoreReviewDetail
	err = newStrictDecoder(data).Decode(&dst.AppClipAppStoreReviewDetail)
	if err == nil {
		jsonAppClipAppStoreReviewDetail, _ := json.Marshal(dst.AppClipAppStoreReviewDetail)
		if string(jsonAppClipAppStoreReviewDetail) == "{}" { // empty struct
			dst.AppClipAppStoreReviewDetail = nil
		} else {
			match++
		}
	} else {
		dst.AppClipAppStoreReviewDetail = nil
	}

	// try to unmarshal data into AppClipDefaultExperienceLocalization
	err = newStrictDecoder(data).Decode(&dst.AppClipDefaultExperienceLocalization)
	if err == nil {
		jsonAppClipDefaultExperienceLocalization, _ := json.Marshal(dst.AppClipDefaultExperienceLocalization)
		if string(jsonAppClipDefaultExperienceLocalization) == "{}" { // empty struct
			dst.AppClipDefaultExperienceLocalization = nil
		} else {
			match++
		}
	} else {
		dst.AppClipDefaultExperienceLocalization = nil
	}

	// try to unmarshal data into AppStoreVersion
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersion)
	if err == nil {
		jsonAppStoreVersion, _ := json.Marshal(dst.AppStoreVersion)
		if string(jsonAppStoreVersion) == "{}" { // empty struct
			dst.AppStoreVersion = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersion = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AppClip = nil
		dst.AppClipAppStoreReviewDetail = nil
		dst.AppClipDefaultExperienceLocalization = nil
		dst.AppStoreVersion = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppClipDefaultExperiencesResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppClipDefaultExperiencesResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppClipDefaultExperiencesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AppClip != nil {
		return json.Marshal(&src.AppClip)
	}

	if src.AppClipAppStoreReviewDetail != nil {
		return json.Marshal(&src.AppClipAppStoreReviewDetail)
	}

	if src.AppClipDefaultExperienceLocalization != nil {
		return json.Marshal(&src.AppClipDefaultExperienceLocalization)
	}

	if src.AppStoreVersion != nil {
		return json.Marshal(&src.AppStoreVersion)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppClipDefaultExperiencesResponseIncludedInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.AppClip != nil {
		return obj.AppClip
	}

	if obj.AppClipAppStoreReviewDetail != nil {
		return obj.AppClipAppStoreReviewDetail
	}

	if obj.AppClipDefaultExperienceLocalization != nil {
		return obj.AppClipDefaultExperienceLocalization
	}

	if obj.AppStoreVersion != nil {
		return obj.AppStoreVersion
	}

	// all schemas are nil
	return nil
}

type NullableAppClipDefaultExperiencesResponseIncludedInner struct {
	value *AppClipDefaultExperiencesResponseIncludedInner
	isSet bool
}

func (v NullableAppClipDefaultExperiencesResponseIncludedInner) Get() *AppClipDefaultExperiencesResponseIncludedInner {
	return v.value
}

func (v *NullableAppClipDefaultExperiencesResponseIncludedInner) Set(val *AppClipDefaultExperiencesResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipDefaultExperiencesResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipDefaultExperiencesResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipDefaultExperiencesResponseIncludedInner(val *AppClipDefaultExperiencesResponseIncludedInner) *NullableAppClipDefaultExperiencesResponseIncludedInner {
	return &NullableAppClipDefaultExperiencesResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAppClipDefaultExperiencesResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipDefaultExperiencesResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
