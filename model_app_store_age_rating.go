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

// AppStoreAgeRating the model 'AppStoreAgeRating'
type AppStoreAgeRating string

// List of AppStoreAgeRating
const (
	FOUR_PLUS      AppStoreAgeRating = "FOUR_PLUS"
	NINE_PLUS      AppStoreAgeRating = "NINE_PLUS"
	TWELVE_PLUS    AppStoreAgeRating = "TWELVE_PLUS"
	SEVENTEEN_PLUS AppStoreAgeRating = "SEVENTEEN_PLUS"
)

// All allowed values of AppStoreAgeRating enum
var AllowedAppStoreAgeRatingEnumValues = []AppStoreAgeRating{
	"FOUR_PLUS",
	"NINE_PLUS",
	"TWELVE_PLUS",
	"SEVENTEEN_PLUS",
}

func (v *AppStoreAgeRating) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AppStoreAgeRating(value)
	for _, existing := range AllowedAppStoreAgeRatingEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AppStoreAgeRating", value)
}

// NewAppStoreAgeRatingFromValue returns a pointer to a valid AppStoreAgeRating
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAppStoreAgeRatingFromValue(v string) (*AppStoreAgeRating, error) {
	ev := AppStoreAgeRating(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AppStoreAgeRating: valid values are %v", v, AllowedAppStoreAgeRatingEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AppStoreAgeRating) IsValid() bool {
	for _, existing := range AllowedAppStoreAgeRatingEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AppStoreAgeRating value
func (v AppStoreAgeRating) Ptr() *AppStoreAgeRating {
	return &v
}

type NullableAppStoreAgeRating struct {
	value *AppStoreAgeRating
	isSet bool
}

func (v NullableAppStoreAgeRating) Get() *AppStoreAgeRating {
	return v.value
}

func (v *NullableAppStoreAgeRating) Set(val *AppStoreAgeRating) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreAgeRating) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreAgeRating) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreAgeRating(val *AppStoreAgeRating) *NullableAppStoreAgeRating {
	return &NullableAppStoreAgeRating{value: val, isSet: true}
}

func (v NullableAppStoreAgeRating) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreAgeRating) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
