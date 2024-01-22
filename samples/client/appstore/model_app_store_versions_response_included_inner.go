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

// AppStoreVersionsResponseIncludedInner - struct for AppStoreVersionsResponseIncludedInner
type AppStoreVersionsResponseIncludedInner struct {
	AgeRatingDeclaration *AgeRatingDeclaration
	App *App
	AppClipDefaultExperience *AppClipDefaultExperience
	AppStoreReviewDetail *AppStoreReviewDetail
	AppStoreVersionExperiment *AppStoreVersionExperiment
	AppStoreVersionExperimentV2 *AppStoreVersionExperimentV2
	AppStoreVersionLocalization *AppStoreVersionLocalization
	AppStoreVersionPhasedRelease *AppStoreVersionPhasedRelease
	AppStoreVersionSubmission *AppStoreVersionSubmission
	Build *Build
	RoutingAppCoverage *RoutingAppCoverage
}

// AgeRatingDeclarationAsAppStoreVersionsResponseIncludedInner is a convenience function that returns AgeRatingDeclaration wrapped in AppStoreVersionsResponseIncludedInner
func AgeRatingDeclarationAsAppStoreVersionsResponseIncludedInner(v *AgeRatingDeclaration) AppStoreVersionsResponseIncludedInner {
	return AppStoreVersionsResponseIncludedInner{
		AgeRatingDeclaration: v,
	}
}

// AppAsAppStoreVersionsResponseIncludedInner is a convenience function that returns App wrapped in AppStoreVersionsResponseIncludedInner
func AppAsAppStoreVersionsResponseIncludedInner(v *App) AppStoreVersionsResponseIncludedInner {
	return AppStoreVersionsResponseIncludedInner{
		App: v,
	}
}

// AppClipDefaultExperienceAsAppStoreVersionsResponseIncludedInner is a convenience function that returns AppClipDefaultExperience wrapped in AppStoreVersionsResponseIncludedInner
func AppClipDefaultExperienceAsAppStoreVersionsResponseIncludedInner(v *AppClipDefaultExperience) AppStoreVersionsResponseIncludedInner {
	return AppStoreVersionsResponseIncludedInner{
		AppClipDefaultExperience: v,
	}
}

// AppStoreReviewDetailAsAppStoreVersionsResponseIncludedInner is a convenience function that returns AppStoreReviewDetail wrapped in AppStoreVersionsResponseIncludedInner
func AppStoreReviewDetailAsAppStoreVersionsResponseIncludedInner(v *AppStoreReviewDetail) AppStoreVersionsResponseIncludedInner {
	return AppStoreVersionsResponseIncludedInner{
		AppStoreReviewDetail: v,
	}
}

// AppStoreVersionExperimentAsAppStoreVersionsResponseIncludedInner is a convenience function that returns AppStoreVersionExperiment wrapped in AppStoreVersionsResponseIncludedInner
func AppStoreVersionExperimentAsAppStoreVersionsResponseIncludedInner(v *AppStoreVersionExperiment) AppStoreVersionsResponseIncludedInner {
	return AppStoreVersionsResponseIncludedInner{
		AppStoreVersionExperiment: v,
	}
}

// AppStoreVersionExperimentV2AsAppStoreVersionsResponseIncludedInner is a convenience function that returns AppStoreVersionExperimentV2 wrapped in AppStoreVersionsResponseIncludedInner
func AppStoreVersionExperimentV2AsAppStoreVersionsResponseIncludedInner(v *AppStoreVersionExperimentV2) AppStoreVersionsResponseIncludedInner {
	return AppStoreVersionsResponseIncludedInner{
		AppStoreVersionExperimentV2: v,
	}
}

// AppStoreVersionLocalizationAsAppStoreVersionsResponseIncludedInner is a convenience function that returns AppStoreVersionLocalization wrapped in AppStoreVersionsResponseIncludedInner
func AppStoreVersionLocalizationAsAppStoreVersionsResponseIncludedInner(v *AppStoreVersionLocalization) AppStoreVersionsResponseIncludedInner {
	return AppStoreVersionsResponseIncludedInner{
		AppStoreVersionLocalization: v,
	}
}

// AppStoreVersionPhasedReleaseAsAppStoreVersionsResponseIncludedInner is a convenience function that returns AppStoreVersionPhasedRelease wrapped in AppStoreVersionsResponseIncludedInner
func AppStoreVersionPhasedReleaseAsAppStoreVersionsResponseIncludedInner(v *AppStoreVersionPhasedRelease) AppStoreVersionsResponseIncludedInner {
	return AppStoreVersionsResponseIncludedInner{
		AppStoreVersionPhasedRelease: v,
	}
}

// AppStoreVersionSubmissionAsAppStoreVersionsResponseIncludedInner is a convenience function that returns AppStoreVersionSubmission wrapped in AppStoreVersionsResponseIncludedInner
func AppStoreVersionSubmissionAsAppStoreVersionsResponseIncludedInner(v *AppStoreVersionSubmission) AppStoreVersionsResponseIncludedInner {
	return AppStoreVersionsResponseIncludedInner{
		AppStoreVersionSubmission: v,
	}
}

// BuildAsAppStoreVersionsResponseIncludedInner is a convenience function that returns Build wrapped in AppStoreVersionsResponseIncludedInner
func BuildAsAppStoreVersionsResponseIncludedInner(v *Build) AppStoreVersionsResponseIncludedInner {
	return AppStoreVersionsResponseIncludedInner{
		Build: v,
	}
}

// RoutingAppCoverageAsAppStoreVersionsResponseIncludedInner is a convenience function that returns RoutingAppCoverage wrapped in AppStoreVersionsResponseIncludedInner
func RoutingAppCoverageAsAppStoreVersionsResponseIncludedInner(v *RoutingAppCoverage) AppStoreVersionsResponseIncludedInner {
	return AppStoreVersionsResponseIncludedInner{
		RoutingAppCoverage: v,
	}
}


// Unmarshal JSON data into one of the pointers in the struct
func (dst *AppStoreVersionsResponseIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into AgeRatingDeclaration
	err = newStrictDecoder(data).Decode(&dst.AgeRatingDeclaration)
	if err == nil {
		jsonAgeRatingDeclaration, _ := json.Marshal(dst.AgeRatingDeclaration)
		if string(jsonAgeRatingDeclaration) == "{}" { // empty struct
			dst.AgeRatingDeclaration = nil
		} else {
			match++
		}
	} else {
		dst.AgeRatingDeclaration = nil
	}

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

	// try to unmarshal data into AppStoreReviewDetail
	err = newStrictDecoder(data).Decode(&dst.AppStoreReviewDetail)
	if err == nil {
		jsonAppStoreReviewDetail, _ := json.Marshal(dst.AppStoreReviewDetail)
		if string(jsonAppStoreReviewDetail) == "{}" { // empty struct
			dst.AppStoreReviewDetail = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreReviewDetail = nil
	}

	// try to unmarshal data into AppStoreVersionExperiment
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersionExperiment)
	if err == nil {
		jsonAppStoreVersionExperiment, _ := json.Marshal(dst.AppStoreVersionExperiment)
		if string(jsonAppStoreVersionExperiment) == "{}" { // empty struct
			dst.AppStoreVersionExperiment = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersionExperiment = nil
	}

	// try to unmarshal data into AppStoreVersionExperimentV2
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersionExperimentV2)
	if err == nil {
		jsonAppStoreVersionExperimentV2, _ := json.Marshal(dst.AppStoreVersionExperimentV2)
		if string(jsonAppStoreVersionExperimentV2) == "{}" { // empty struct
			dst.AppStoreVersionExperimentV2 = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersionExperimentV2 = nil
	}

	// try to unmarshal data into AppStoreVersionLocalization
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersionLocalization)
	if err == nil {
		jsonAppStoreVersionLocalization, _ := json.Marshal(dst.AppStoreVersionLocalization)
		if string(jsonAppStoreVersionLocalization) == "{}" { // empty struct
			dst.AppStoreVersionLocalization = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersionLocalization = nil
	}

	// try to unmarshal data into AppStoreVersionPhasedRelease
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersionPhasedRelease)
	if err == nil {
		jsonAppStoreVersionPhasedRelease, _ := json.Marshal(dst.AppStoreVersionPhasedRelease)
		if string(jsonAppStoreVersionPhasedRelease) == "{}" { // empty struct
			dst.AppStoreVersionPhasedRelease = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersionPhasedRelease = nil
	}

	// try to unmarshal data into AppStoreVersionSubmission
	err = newStrictDecoder(data).Decode(&dst.AppStoreVersionSubmission)
	if err == nil {
		jsonAppStoreVersionSubmission, _ := json.Marshal(dst.AppStoreVersionSubmission)
		if string(jsonAppStoreVersionSubmission) == "{}" { // empty struct
			dst.AppStoreVersionSubmission = nil
		} else {
			match++
		}
	} else {
		dst.AppStoreVersionSubmission = nil
	}

	// try to unmarshal data into Build
	err = newStrictDecoder(data).Decode(&dst.Build)
	if err == nil {
		jsonBuild, _ := json.Marshal(dst.Build)
		if string(jsonBuild) == "{}" { // empty struct
			dst.Build = nil
		} else {
			match++
		}
	} else {
		dst.Build = nil
	}

	// try to unmarshal data into RoutingAppCoverage
	err = newStrictDecoder(data).Decode(&dst.RoutingAppCoverage)
	if err == nil {
		jsonRoutingAppCoverage, _ := json.Marshal(dst.RoutingAppCoverage)
		if string(jsonRoutingAppCoverage) == "{}" { // empty struct
			dst.RoutingAppCoverage = nil
		} else {
			match++
		}
	} else {
		dst.RoutingAppCoverage = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.AgeRatingDeclaration = nil
		dst.App = nil
		dst.AppClipDefaultExperience = nil
		dst.AppStoreReviewDetail = nil
		dst.AppStoreVersionExperiment = nil
		dst.AppStoreVersionExperimentV2 = nil
		dst.AppStoreVersionLocalization = nil
		dst.AppStoreVersionPhasedRelease = nil
		dst.AppStoreVersionSubmission = nil
		dst.Build = nil
		dst.RoutingAppCoverage = nil

		return fmt.Errorf("data matches more than one schema in oneOf(AppStoreVersionsResponseIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(AppStoreVersionsResponseIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src AppStoreVersionsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	if src.AgeRatingDeclaration != nil {
		return json.Marshal(&src.AgeRatingDeclaration)
	}

	if src.App != nil {
		return json.Marshal(&src.App)
	}

	if src.AppClipDefaultExperience != nil {
		return json.Marshal(&src.AppClipDefaultExperience)
	}

	if src.AppStoreReviewDetail != nil {
		return json.Marshal(&src.AppStoreReviewDetail)
	}

	if src.AppStoreVersionExperiment != nil {
		return json.Marshal(&src.AppStoreVersionExperiment)
	}

	if src.AppStoreVersionExperimentV2 != nil {
		return json.Marshal(&src.AppStoreVersionExperimentV2)
	}

	if src.AppStoreVersionLocalization != nil {
		return json.Marshal(&src.AppStoreVersionLocalization)
	}

	if src.AppStoreVersionPhasedRelease != nil {
		return json.Marshal(&src.AppStoreVersionPhasedRelease)
	}

	if src.AppStoreVersionSubmission != nil {
		return json.Marshal(&src.AppStoreVersionSubmission)
	}

	if src.Build != nil {
		return json.Marshal(&src.Build)
	}

	if src.RoutingAppCoverage != nil {
		return json.Marshal(&src.RoutingAppCoverage)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *AppStoreVersionsResponseIncludedInner) GetActualInstance() (interface{}) {
	if obj == nil {
		return nil
	}
	if obj.AgeRatingDeclaration != nil {
		return obj.AgeRatingDeclaration
	}

	if obj.App != nil {
		return obj.App
	}

	if obj.AppClipDefaultExperience != nil {
		return obj.AppClipDefaultExperience
	}

	if obj.AppStoreReviewDetail != nil {
		return obj.AppStoreReviewDetail
	}

	if obj.AppStoreVersionExperiment != nil {
		return obj.AppStoreVersionExperiment
	}

	if obj.AppStoreVersionExperimentV2 != nil {
		return obj.AppStoreVersionExperimentV2
	}

	if obj.AppStoreVersionLocalization != nil {
		return obj.AppStoreVersionLocalization
	}

	if obj.AppStoreVersionPhasedRelease != nil {
		return obj.AppStoreVersionPhasedRelease
	}

	if obj.AppStoreVersionSubmission != nil {
		return obj.AppStoreVersionSubmission
	}

	if obj.Build != nil {
		return obj.Build
	}

	if obj.RoutingAppCoverage != nil {
		return obj.RoutingAppCoverage
	}

	// all schemas are nil
	return nil
}

type NullableAppStoreVersionsResponseIncludedInner struct {
	value *AppStoreVersionsResponseIncludedInner
	isSet bool
}

func (v NullableAppStoreVersionsResponseIncludedInner) Get() *AppStoreVersionsResponseIncludedInner {
	return v.value
}

func (v *NullableAppStoreVersionsResponseIncludedInner) Set(val *AppStoreVersionsResponseIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionsResponseIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionsResponseIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionsResponseIncludedInner(val *AppStoreVersionsResponseIncludedInner) *NullableAppStoreVersionsResponseIncludedInner {
	return &NullableAppStoreVersionsResponseIncludedInner{value: val, isSet: true}
}

func (v NullableAppStoreVersionsResponseIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionsResponseIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


