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

// CiTestDestinationKind the model 'CiTestDestinationKind'
type CiTestDestinationKind string

// List of CiTestDestinationKind
const (
	SIMULATOR CiTestDestinationKind = "SIMULATOR"
	MAC CiTestDestinationKind = "MAC"
)

// All allowed values of CiTestDestinationKind enum
var AllowedCiTestDestinationKindEnumValues = []CiTestDestinationKind{
	"SIMULATOR",
	"MAC",
}

func (v *CiTestDestinationKind) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CiTestDestinationKind(value)
	for _, existing := range AllowedCiTestDestinationKindEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CiTestDestinationKind", value)
}

// NewCiTestDestinationKindFromValue returns a pointer to a valid CiTestDestinationKind
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCiTestDestinationKindFromValue(v string) (*CiTestDestinationKind, error) {
	ev := CiTestDestinationKind(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CiTestDestinationKind: valid values are %v", v, AllowedCiTestDestinationKindEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CiTestDestinationKind) IsValid() bool {
	for _, existing := range AllowedCiTestDestinationKindEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CiTestDestinationKind value
func (v CiTestDestinationKind) Ptr() *CiTestDestinationKind {
	return &v
}

type NullableCiTestDestinationKind struct {
	value *CiTestDestinationKind
	isSet bool
}

func (v NullableCiTestDestinationKind) Get() *CiTestDestinationKind {
	return v.value
}

func (v *NullableCiTestDestinationKind) Set(val *CiTestDestinationKind) {
	v.value = val
	v.isSet = true
}

func (v NullableCiTestDestinationKind) IsSet() bool {
	return v.isSet
}

func (v *NullableCiTestDestinationKind) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiTestDestinationKind(val *CiTestDestinationKind) *NullableCiTestDestinationKind {
	return &NullableCiTestDestinationKind{value: val, isSet: true}
}

func (v NullableCiTestDestinationKind) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiTestDestinationKind) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

