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

// checks if the BetaBuildUsagesV1MetricResponseDataInnerDimensions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaBuildUsagesV1MetricResponseDataInnerDimensions{}

// BetaBuildUsagesV1MetricResponseDataInnerDimensions struct for BetaBuildUsagesV1MetricResponseDataInnerDimensions
type BetaBuildUsagesV1MetricResponseDataInnerDimensions struct {
	BundleIds *BetaBuildUsagesV1MetricResponseDataInnerDimensionsBundleIds `json:"bundleIds,omitempty"`
}

// NewBetaBuildUsagesV1MetricResponseDataInnerDimensions instantiates a new BetaBuildUsagesV1MetricResponseDataInnerDimensions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaBuildUsagesV1MetricResponseDataInnerDimensions() *BetaBuildUsagesV1MetricResponseDataInnerDimensions {
	this := BetaBuildUsagesV1MetricResponseDataInnerDimensions{}
	return &this
}

// NewBetaBuildUsagesV1MetricResponseDataInnerDimensionsWithDefaults instantiates a new BetaBuildUsagesV1MetricResponseDataInnerDimensions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaBuildUsagesV1MetricResponseDataInnerDimensionsWithDefaults() *BetaBuildUsagesV1MetricResponseDataInnerDimensions {
	this := BetaBuildUsagesV1MetricResponseDataInnerDimensions{}
	return &this
}

// GetBundleIds returns the BundleIds field value if set, zero value otherwise.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDimensions) GetBundleIds() BetaBuildUsagesV1MetricResponseDataInnerDimensionsBundleIds {
	if o == nil || IsNil(o.BundleIds) {
		var ret BetaBuildUsagesV1MetricResponseDataInnerDimensionsBundleIds
		return ret
	}
	return *o.BundleIds
}

// GetBundleIdsOk returns a tuple with the BundleIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDimensions) GetBundleIdsOk() (*BetaBuildUsagesV1MetricResponseDataInnerDimensionsBundleIds, bool) {
	if o == nil || IsNil(o.BundleIds) {
		return nil, false
	}
	return o.BundleIds, true
}

// HasBundleIds returns a boolean if a field has been set.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDimensions) HasBundleIds() bool {
	if o != nil && !IsNil(o.BundleIds) {
		return true
	}

	return false
}

// SetBundleIds gets a reference to the given BetaBuildUsagesV1MetricResponseDataInnerDimensionsBundleIds and assigns it to the BundleIds field.
func (o *BetaBuildUsagesV1MetricResponseDataInnerDimensions) SetBundleIds(v BetaBuildUsagesV1MetricResponseDataInnerDimensionsBundleIds) {
	o.BundleIds = &v
}

func (o BetaBuildUsagesV1MetricResponseDataInnerDimensions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaBuildUsagesV1MetricResponseDataInnerDimensions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BundleIds) {
		toSerialize["bundleIds"] = o.BundleIds
	}
	return toSerialize, nil
}

type NullableBetaBuildUsagesV1MetricResponseDataInnerDimensions struct {
	value *BetaBuildUsagesV1MetricResponseDataInnerDimensions
	isSet bool
}

func (v NullableBetaBuildUsagesV1MetricResponseDataInnerDimensions) Get() *BetaBuildUsagesV1MetricResponseDataInnerDimensions {
	return v.value
}

func (v *NullableBetaBuildUsagesV1MetricResponseDataInnerDimensions) Set(val *BetaBuildUsagesV1MetricResponseDataInnerDimensions) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaBuildUsagesV1MetricResponseDataInnerDimensions) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaBuildUsagesV1MetricResponseDataInnerDimensions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaBuildUsagesV1MetricResponseDataInnerDimensions(val *BetaBuildUsagesV1MetricResponseDataInnerDimensions) *NullableBetaBuildUsagesV1MetricResponseDataInnerDimensions {
	return &NullableBetaBuildUsagesV1MetricResponseDataInnerDimensions{value: val, isSet: true}
}

func (v NullableBetaBuildUsagesV1MetricResponseDataInnerDimensions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaBuildUsagesV1MetricResponseDataInnerDimensions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}