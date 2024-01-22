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

// checks if the BetaTesterUsagesV1MetricResponseDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaTesterUsagesV1MetricResponseDataInner{}

// BetaTesterUsagesV1MetricResponseDataInner struct for BetaTesterUsagesV1MetricResponseDataInner
type BetaTesterUsagesV1MetricResponseDataInner struct {
	DataPoints *AppsBetaTesterUsagesV1MetricResponseDataInnerDataPoints `json:"dataPoints,omitempty"`
	Dimensions *BetaTesterUsagesV1MetricResponseDataInnerDimensions     `json:"dimensions,omitempty"`
}

// NewBetaTesterUsagesV1MetricResponseDataInner instantiates a new BetaTesterUsagesV1MetricResponseDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaTesterUsagesV1MetricResponseDataInner() *BetaTesterUsagesV1MetricResponseDataInner {
	this := BetaTesterUsagesV1MetricResponseDataInner{}
	return &this
}

// NewBetaTesterUsagesV1MetricResponseDataInnerWithDefaults instantiates a new BetaTesterUsagesV1MetricResponseDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaTesterUsagesV1MetricResponseDataInnerWithDefaults() *BetaTesterUsagesV1MetricResponseDataInner {
	this := BetaTesterUsagesV1MetricResponseDataInner{}
	return &this
}

// GetDataPoints returns the DataPoints field value if set, zero value otherwise.
func (o *BetaTesterUsagesV1MetricResponseDataInner) GetDataPoints() AppsBetaTesterUsagesV1MetricResponseDataInnerDataPoints {
	if o == nil || IsNil(o.DataPoints) {
		var ret AppsBetaTesterUsagesV1MetricResponseDataInnerDataPoints
		return ret
	}
	return *o.DataPoints
}

// GetDataPointsOk returns a tuple with the DataPoints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterUsagesV1MetricResponseDataInner) GetDataPointsOk() (*AppsBetaTesterUsagesV1MetricResponseDataInnerDataPoints, bool) {
	if o == nil || IsNil(o.DataPoints) {
		return nil, false
	}
	return o.DataPoints, true
}

// HasDataPoints returns a boolean if a field has been set.
func (o *BetaTesterUsagesV1MetricResponseDataInner) HasDataPoints() bool {
	if o != nil && !IsNil(o.DataPoints) {
		return true
	}

	return false
}

// SetDataPoints gets a reference to the given AppsBetaTesterUsagesV1MetricResponseDataInnerDataPoints and assigns it to the DataPoints field.
func (o *BetaTesterUsagesV1MetricResponseDataInner) SetDataPoints(v AppsBetaTesterUsagesV1MetricResponseDataInnerDataPoints) {
	o.DataPoints = &v
}

// GetDimensions returns the Dimensions field value if set, zero value otherwise.
func (o *BetaTesterUsagesV1MetricResponseDataInner) GetDimensions() BetaTesterUsagesV1MetricResponseDataInnerDimensions {
	if o == nil || IsNil(o.Dimensions) {
		var ret BetaTesterUsagesV1MetricResponseDataInnerDimensions
		return ret
	}
	return *o.Dimensions
}

// GetDimensionsOk returns a tuple with the Dimensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaTesterUsagesV1MetricResponseDataInner) GetDimensionsOk() (*BetaTesterUsagesV1MetricResponseDataInnerDimensions, bool) {
	if o == nil || IsNil(o.Dimensions) {
		return nil, false
	}
	return o.Dimensions, true
}

// HasDimensions returns a boolean if a field has been set.
func (o *BetaTesterUsagesV1MetricResponseDataInner) HasDimensions() bool {
	if o != nil && !IsNil(o.Dimensions) {
		return true
	}

	return false
}

// SetDimensions gets a reference to the given BetaTesterUsagesV1MetricResponseDataInnerDimensions and assigns it to the Dimensions field.
func (o *BetaTesterUsagesV1MetricResponseDataInner) SetDimensions(v BetaTesterUsagesV1MetricResponseDataInnerDimensions) {
	o.Dimensions = &v
}

func (o BetaTesterUsagesV1MetricResponseDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaTesterUsagesV1MetricResponseDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DataPoints) {
		toSerialize["dataPoints"] = o.DataPoints
	}
	if !IsNil(o.Dimensions) {
		toSerialize["dimensions"] = o.Dimensions
	}
	return toSerialize, nil
}

type NullableBetaTesterUsagesV1MetricResponseDataInner struct {
	value *BetaTesterUsagesV1MetricResponseDataInner
	isSet bool
}

func (v NullableBetaTesterUsagesV1MetricResponseDataInner) Get() *BetaTesterUsagesV1MetricResponseDataInner {
	return v.value
}

func (v *NullableBetaTesterUsagesV1MetricResponseDataInner) Set(val *BetaTesterUsagesV1MetricResponseDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaTesterUsagesV1MetricResponseDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaTesterUsagesV1MetricResponseDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaTesterUsagesV1MetricResponseDataInner(val *BetaTesterUsagesV1MetricResponseDataInner) *NullableBetaTesterUsagesV1MetricResponseDataInner {
	return &NullableBetaTesterUsagesV1MetricResponseDataInner{value: val, isSet: true}
}

func (v NullableBetaTesterUsagesV1MetricResponseDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaTesterUsagesV1MetricResponseDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
