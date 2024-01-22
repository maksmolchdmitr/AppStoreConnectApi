/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues{}

// GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues struct for GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues
type GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues struct {
	Count *int32 `json:"count,omitempty"`
	AverageResult *float32 `json:"averageResult,omitempty"`
	P50Result *float32 `json:"p50Result,omitempty"`
	P95Result *float32 `json:"p95Result,omitempty"`
}

// NewGameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues instantiates a new GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues() *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues {
	this := GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues{}
	return &this
}

// NewGameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValuesWithDefaults instantiates a new GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValuesWithDefaults() *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues {
	this := GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues{}
	return &this
}

// GetCount returns the Count field value if set, zero value otherwise.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) GetCount() int32 {
	if o == nil || IsNil(o.Count) {
		var ret int32
		return ret
	}
	return *o.Count
}

// GetCountOk returns a tuple with the Count field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) GetCountOk() (*int32, bool) {
	if o == nil || IsNil(o.Count) {
		return nil, false
	}
	return o.Count, true
}

// HasCount returns a boolean if a field has been set.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) HasCount() bool {
	if o != nil && !IsNil(o.Count) {
		return true
	}

	return false
}

// SetCount gets a reference to the given int32 and assigns it to the Count field.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) SetCount(v int32) {
	o.Count = &v
}

// GetAverageResult returns the AverageResult field value if set, zero value otherwise.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) GetAverageResult() float32 {
	if o == nil || IsNil(o.AverageResult) {
		var ret float32
		return ret
	}
	return *o.AverageResult
}

// GetAverageResultOk returns a tuple with the AverageResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) GetAverageResultOk() (*float32, bool) {
	if o == nil || IsNil(o.AverageResult) {
		return nil, false
	}
	return o.AverageResult, true
}

// HasAverageResult returns a boolean if a field has been set.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) HasAverageResult() bool {
	if o != nil && !IsNil(o.AverageResult) {
		return true
	}

	return false
}

// SetAverageResult gets a reference to the given float32 and assigns it to the AverageResult field.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) SetAverageResult(v float32) {
	o.AverageResult = &v
}

// GetP50Result returns the P50Result field value if set, zero value otherwise.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) GetP50Result() float32 {
	if o == nil || IsNil(o.P50Result) {
		var ret float32
		return ret
	}
	return *o.P50Result
}

// GetP50ResultOk returns a tuple with the P50Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) GetP50ResultOk() (*float32, bool) {
	if o == nil || IsNil(o.P50Result) {
		return nil, false
	}
	return o.P50Result, true
}

// HasP50Result returns a boolean if a field has been set.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) HasP50Result() bool {
	if o != nil && !IsNil(o.P50Result) {
		return true
	}

	return false
}

// SetP50Result gets a reference to the given float32 and assigns it to the P50Result field.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) SetP50Result(v float32) {
	o.P50Result = &v
}

// GetP95Result returns the P95Result field value if set, zero value otherwise.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) GetP95Result() float32 {
	if o == nil || IsNil(o.P95Result) {
		var ret float32
		return ret
	}
	return *o.P95Result
}

// GetP95ResultOk returns a tuple with the P95Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) GetP95ResultOk() (*float32, bool) {
	if o == nil || IsNil(o.P95Result) {
		return nil, false
	}
	return o.P95Result, true
}

// HasP95Result returns a boolean if a field has been set.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) HasP95Result() bool {
	if o != nil && !IsNil(o.P95Result) {
		return true
	}

	return false
}

// SetP95Result gets a reference to the given float32 and assigns it to the P95Result field.
func (o *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) SetP95Result(v float32) {
	o.P95Result = &v
}

func (o GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Count) {
		toSerialize["count"] = o.Count
	}
	if !IsNil(o.AverageResult) {
		toSerialize["averageResult"] = o.AverageResult
	}
	if !IsNil(o.P50Result) {
		toSerialize["p50Result"] = o.P50Result
	}
	if !IsNil(o.P95Result) {
		toSerialize["p95Result"] = o.P95Result
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues struct {
	value *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues
	isSet bool
}

func (v NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) Get() *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues {
	return v.value
}

func (v *NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) Set(val *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues(val *GameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) *NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues {
	return &NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingNumberRuleResultsV1MetricResponseDataInnerDataPointsValues) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

