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

// BuildAudienceType the model 'BuildAudienceType'
type BuildAudienceType string

// List of BuildAudienceType
const (
	INTERNAL_ONLY      BuildAudienceType = "INTERNAL_ONLY"
	APP_STORE_ELIGIBLE BuildAudienceType = "APP_STORE_ELIGIBLE"
)

// All allowed values of BuildAudienceType enum
var AllowedBuildAudienceTypeEnumValues = []BuildAudienceType{
	"INTERNAL_ONLY",
	"APP_STORE_ELIGIBLE",
}

func (v *BuildAudienceType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := BuildAudienceType(value)
	for _, existing := range AllowedBuildAudienceTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid BuildAudienceType", value)
}

// NewBuildAudienceTypeFromValue returns a pointer to a valid BuildAudienceType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewBuildAudienceTypeFromValue(v string) (*BuildAudienceType, error) {
	ev := BuildAudienceType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for BuildAudienceType: valid values are %v", v, AllowedBuildAudienceTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v BuildAudienceType) IsValid() bool {
	for _, existing := range AllowedBuildAudienceTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to BuildAudienceType value
func (v BuildAudienceType) Ptr() *BuildAudienceType {
	return &v
}

type NullableBuildAudienceType struct {
	value *BuildAudienceType
	isSet bool
}

func (v NullableBuildAudienceType) Get() *BuildAudienceType {
	return v.value
}

func (v *NullableBuildAudienceType) Set(val *BuildAudienceType) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildAudienceType) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildAudienceType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildAudienceType(val *BuildAudienceType) *NullableBuildAudienceType {
	return &NullableBuildAudienceType{value: val, isSet: true}
}

func (v NullableBuildAudienceType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildAudienceType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
