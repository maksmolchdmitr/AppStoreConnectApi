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

// GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner - struct for GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner
type GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner struct {
	GameCenterMatchmakingTestPlayerPropertyInlineCreate *GameCenterMatchmakingTestPlayerPropertyInlineCreate
	GameCenterMatchmakingTestRequestInlineCreate        *GameCenterMatchmakingTestRequestInlineCreate
}

// GameCenterMatchmakingTestPlayerPropertyInlineCreateAsGameCenterMatchmakingRuleSetTestCreateRequestIncludedInner is a convenience function that returns GameCenterMatchmakingTestPlayerPropertyInlineCreate wrapped in GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner
func GameCenterMatchmakingTestPlayerPropertyInlineCreateAsGameCenterMatchmakingRuleSetTestCreateRequestIncludedInner(v *GameCenterMatchmakingTestPlayerPropertyInlineCreate) GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner {
	return GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner{
		GameCenterMatchmakingTestPlayerPropertyInlineCreate: v,
	}
}

// GameCenterMatchmakingTestRequestInlineCreateAsGameCenterMatchmakingRuleSetTestCreateRequestIncludedInner is a convenience function that returns GameCenterMatchmakingTestRequestInlineCreate wrapped in GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner
func GameCenterMatchmakingTestRequestInlineCreateAsGameCenterMatchmakingRuleSetTestCreateRequestIncludedInner(v *GameCenterMatchmakingTestRequestInlineCreate) GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner {
	return GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner{
		GameCenterMatchmakingTestRequestInlineCreate: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into GameCenterMatchmakingTestPlayerPropertyInlineCreate
	err = newStrictDecoder(data).Decode(&dst.GameCenterMatchmakingTestPlayerPropertyInlineCreate)
	if err == nil {
		jsonGameCenterMatchmakingTestPlayerPropertyInlineCreate, _ := json.Marshal(dst.GameCenterMatchmakingTestPlayerPropertyInlineCreate)
		if string(jsonGameCenterMatchmakingTestPlayerPropertyInlineCreate) == "{}" { // empty struct
			dst.GameCenterMatchmakingTestPlayerPropertyInlineCreate = nil
		} else {
			match++
		}
	} else {
		dst.GameCenterMatchmakingTestPlayerPropertyInlineCreate = nil
	}

	// try to unmarshal data into GameCenterMatchmakingTestRequestInlineCreate
	err = newStrictDecoder(data).Decode(&dst.GameCenterMatchmakingTestRequestInlineCreate)
	if err == nil {
		jsonGameCenterMatchmakingTestRequestInlineCreate, _ := json.Marshal(dst.GameCenterMatchmakingTestRequestInlineCreate)
		if string(jsonGameCenterMatchmakingTestRequestInlineCreate) == "{}" { // empty struct
			dst.GameCenterMatchmakingTestRequestInlineCreate = nil
		} else {
			match++
		}
	} else {
		dst.GameCenterMatchmakingTestRequestInlineCreate = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.GameCenterMatchmakingTestPlayerPropertyInlineCreate = nil
		dst.GameCenterMatchmakingTestRequestInlineCreate = nil

		return fmt.Errorf("data matches more than one schema in oneOf(GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("data failed to match schemas in oneOf(GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner) MarshalJSON() ([]byte, error) {
	if src.GameCenterMatchmakingTestPlayerPropertyInlineCreate != nil {
		return json.Marshal(&src.GameCenterMatchmakingTestPlayerPropertyInlineCreate)
	}

	if src.GameCenterMatchmakingTestRequestInlineCreate != nil {
		return json.Marshal(&src.GameCenterMatchmakingTestRequestInlineCreate)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.GameCenterMatchmakingTestPlayerPropertyInlineCreate != nil {
		return obj.GameCenterMatchmakingTestPlayerPropertyInlineCreate
	}

	if obj.GameCenterMatchmakingTestRequestInlineCreate != nil {
		return obj.GameCenterMatchmakingTestRequestInlineCreate
	}

	// all schemas are nil
	return nil
}

type NullableGameCenterMatchmakingRuleSetTestCreateRequestIncludedInner struct {
	value *GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner
	isSet bool
}

func (v NullableGameCenterMatchmakingRuleSetTestCreateRequestIncludedInner) Get() *GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner {
	return v.value
}

func (v *NullableGameCenterMatchmakingRuleSetTestCreateRequestIncludedInner) Set(val *GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingRuleSetTestCreateRequestIncludedInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingRuleSetTestCreateRequestIncludedInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingRuleSetTestCreateRequestIncludedInner(val *GameCenterMatchmakingRuleSetTestCreateRequestIncludedInner) *NullableGameCenterMatchmakingRuleSetTestCreateRequestIncludedInner {
	return &NullableGameCenterMatchmakingRuleSetTestCreateRequestIncludedInner{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingRuleSetTestCreateRequestIncludedInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingRuleSetTestCreateRequestIncludedInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
