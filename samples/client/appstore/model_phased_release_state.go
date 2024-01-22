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

// PhasedReleaseState the model 'PhasedReleaseState'
type PhasedReleaseState string

// List of PhasedReleaseState
const (
	INACTIVE PhasedReleaseState = "INACTIVE"
	ACTIVE PhasedReleaseState = "ACTIVE"
	PAUSED PhasedReleaseState = "PAUSED"
	COMPLETE PhasedReleaseState = "COMPLETE"
)

// All allowed values of PhasedReleaseState enum
var AllowedPhasedReleaseStateEnumValues = []PhasedReleaseState{
	"INACTIVE",
	"ACTIVE",
	"PAUSED",
	"COMPLETE",
}

func (v *PhasedReleaseState) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := PhasedReleaseState(value)
	for _, existing := range AllowedPhasedReleaseStateEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid PhasedReleaseState", value)
}

// NewPhasedReleaseStateFromValue returns a pointer to a valid PhasedReleaseState
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewPhasedReleaseStateFromValue(v string) (*PhasedReleaseState, error) {
	ev := PhasedReleaseState(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for PhasedReleaseState: valid values are %v", v, AllowedPhasedReleaseStateEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v PhasedReleaseState) IsValid() bool {
	for _, existing := range AllowedPhasedReleaseStateEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to PhasedReleaseState value
func (v PhasedReleaseState) Ptr() *PhasedReleaseState {
	return &v
}

type NullablePhasedReleaseState struct {
	value *PhasedReleaseState
	isSet bool
}

func (v NullablePhasedReleaseState) Get() *PhasedReleaseState {
	return v.value
}

func (v *NullablePhasedReleaseState) Set(val *PhasedReleaseState) {
	v.value = val
	v.isSet = true
}

func (v NullablePhasedReleaseState) IsSet() bool {
	return v.isSet
}

func (v *NullablePhasedReleaseState) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePhasedReleaseState(val *PhasedReleaseState) *NullablePhasedReleaseState {
	return &NullablePhasedReleaseState{value: val, isSet: true}
}

func (v NullablePhasedReleaseState) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePhasedReleaseState) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

