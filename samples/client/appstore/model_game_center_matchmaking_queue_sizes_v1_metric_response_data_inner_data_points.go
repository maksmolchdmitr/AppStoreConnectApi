/*
App Store Connect API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 3.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"time"
)

// checks if the GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints{}

// GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints struct for GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints
type GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints struct {
	Start *time.Time `json:"start,omitempty"`
	End *time.Time `json:"end,omitempty"`
	Values *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues `json:"values,omitempty"`
}

// NewGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints instantiates a new GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints() *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints {
	this := GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints{}
	return &this
}

// NewGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsWithDefaults instantiates a new GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsWithDefaults() *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints {
	this := GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints{}
	return &this
}

// GetStart returns the Start field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) GetStart() time.Time {
	if o == nil || IsNil(o.Start) {
		var ret time.Time
		return ret
	}
	return *o.Start
}

// GetStartOk returns a tuple with the Start field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) GetStartOk() (*time.Time, bool) {
	if o == nil || IsNil(o.Start) {
		return nil, false
	}
	return o.Start, true
}

// HasStart returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) HasStart() bool {
	if o != nil && !IsNil(o.Start) {
		return true
	}

	return false
}

// SetStart gets a reference to the given time.Time and assigns it to the Start field.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) SetStart(v time.Time) {
	o.Start = &v
}

// GetEnd returns the End field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) GetEnd() time.Time {
	if o == nil || IsNil(o.End) {
		var ret time.Time
		return ret
	}
	return *o.End
}

// GetEndOk returns a tuple with the End field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) GetEndOk() (*time.Time, bool) {
	if o == nil || IsNil(o.End) {
		return nil, false
	}
	return o.End, true
}

// HasEnd returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) HasEnd() bool {
	if o != nil && !IsNil(o.End) {
		return true
	}

	return false
}

// SetEnd gets a reference to the given time.Time and assigns it to the End field.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) SetEnd(v time.Time) {
	o.End = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) GetValues() GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues {
	if o == nil || IsNil(o.Values) {
		var ret GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues
		return ret
	}
	return *o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) GetValuesOk() (*GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues and assigns it to the Values field.
func (o *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) SetValues(v GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPointsValues) {
	o.Values = &v
}

func (o GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Start) {
		toSerialize["start"] = o.Start
	}
	if !IsNil(o.End) {
		toSerialize["end"] = o.End
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints struct {
	value *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints
	isSet bool
}

func (v NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) Get() *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints {
	return v.value
}

func (v *NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) Set(val *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints(val *GameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) *NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints {
	return &NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints{value: val, isSet: true}
}

func (v NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterMatchmakingQueueSizesV1MetricResponseDataInnerDataPoints) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


