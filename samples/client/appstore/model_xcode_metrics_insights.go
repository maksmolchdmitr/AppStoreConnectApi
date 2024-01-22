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

// checks if the XcodeMetricsInsights type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &XcodeMetricsInsights{}

// XcodeMetricsInsights struct for XcodeMetricsInsights
type XcodeMetricsInsights struct {
	TrendingUp []MetricsInsight `json:"trendingUp,omitempty"`
	Regressions []MetricsInsight `json:"regressions,omitempty"`
}

// NewXcodeMetricsInsights instantiates a new XcodeMetricsInsights object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXcodeMetricsInsights() *XcodeMetricsInsights {
	this := XcodeMetricsInsights{}
	return &this
}

// NewXcodeMetricsInsightsWithDefaults instantiates a new XcodeMetricsInsights object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXcodeMetricsInsightsWithDefaults() *XcodeMetricsInsights {
	this := XcodeMetricsInsights{}
	return &this
}

// GetTrendingUp returns the TrendingUp field value if set, zero value otherwise.
func (o *XcodeMetricsInsights) GetTrendingUp() []MetricsInsight {
	if o == nil || IsNil(o.TrendingUp) {
		var ret []MetricsInsight
		return ret
	}
	return o.TrendingUp
}

// GetTrendingUpOk returns a tuple with the TrendingUp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XcodeMetricsInsights) GetTrendingUpOk() ([]MetricsInsight, bool) {
	if o == nil || IsNil(o.TrendingUp) {
		return nil, false
	}
	return o.TrendingUp, true
}

// HasTrendingUp returns a boolean if a field has been set.
func (o *XcodeMetricsInsights) HasTrendingUp() bool {
	if o != nil && !IsNil(o.TrendingUp) {
		return true
	}

	return false
}

// SetTrendingUp gets a reference to the given []MetricsInsight and assigns it to the TrendingUp field.
func (o *XcodeMetricsInsights) SetTrendingUp(v []MetricsInsight) {
	o.TrendingUp = v
}

// GetRegressions returns the Regressions field value if set, zero value otherwise.
func (o *XcodeMetricsInsights) GetRegressions() []MetricsInsight {
	if o == nil || IsNil(o.Regressions) {
		var ret []MetricsInsight
		return ret
	}
	return o.Regressions
}

// GetRegressionsOk returns a tuple with the Regressions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XcodeMetricsInsights) GetRegressionsOk() ([]MetricsInsight, bool) {
	if o == nil || IsNil(o.Regressions) {
		return nil, false
	}
	return o.Regressions, true
}

// HasRegressions returns a boolean if a field has been set.
func (o *XcodeMetricsInsights) HasRegressions() bool {
	if o != nil && !IsNil(o.Regressions) {
		return true
	}

	return false
}

// SetRegressions gets a reference to the given []MetricsInsight and assigns it to the Regressions field.
func (o *XcodeMetricsInsights) SetRegressions(v []MetricsInsight) {
	o.Regressions = v
}

func (o XcodeMetricsInsights) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o XcodeMetricsInsights) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TrendingUp) {
		toSerialize["trendingUp"] = o.TrendingUp
	}
	if !IsNil(o.Regressions) {
		toSerialize["regressions"] = o.Regressions
	}
	return toSerialize, nil
}

type NullableXcodeMetricsInsights struct {
	value *XcodeMetricsInsights
	isSet bool
}

func (v NullableXcodeMetricsInsights) Get() *XcodeMetricsInsights {
	return v.value
}

func (v *NullableXcodeMetricsInsights) Set(val *XcodeMetricsInsights) {
	v.value = val
	v.isSet = true
}

func (v NullableXcodeMetricsInsights) IsSet() bool {
	return v.isSet
}

func (v *NullableXcodeMetricsInsights) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXcodeMetricsInsights(val *XcodeMetricsInsights) *NullableXcodeMetricsInsights {
	return &NullableXcodeMetricsInsights{value: val, isSet: true}
}

func (v NullableXcodeMetricsInsights) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXcodeMetricsInsights) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


