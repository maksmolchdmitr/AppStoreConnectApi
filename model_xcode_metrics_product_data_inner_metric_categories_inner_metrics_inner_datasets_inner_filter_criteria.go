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

// checks if the XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria{}

// XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria struct for XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria
type XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria struct {
	Percentile          *string `json:"percentile,omitempty"`
	Device              *string `json:"device,omitempty"`
	DeviceMarketingName *string `json:"deviceMarketingName,omitempty"`
}

// NewXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria instantiates a new XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria() *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria {
	this := XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria{}
	return &this
}

// NewXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteriaWithDefaults instantiates a new XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteriaWithDefaults() *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria {
	this := XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria{}
	return &this
}

// GetPercentile returns the Percentile field value if set, zero value otherwise.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) GetPercentile() string {
	if o == nil || IsNil(o.Percentile) {
		var ret string
		return ret
	}
	return *o.Percentile
}

// GetPercentileOk returns a tuple with the Percentile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) GetPercentileOk() (*string, bool) {
	if o == nil || IsNil(o.Percentile) {
		return nil, false
	}
	return o.Percentile, true
}

// HasPercentile returns a boolean if a field has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) HasPercentile() bool {
	if o != nil && !IsNil(o.Percentile) {
		return true
	}

	return false
}

// SetPercentile gets a reference to the given string and assigns it to the Percentile field.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) SetPercentile(v string) {
	o.Percentile = &v
}

// GetDevice returns the Device field value if set, zero value otherwise.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) GetDevice() string {
	if o == nil || IsNil(o.Device) {
		var ret string
		return ret
	}
	return *o.Device
}

// GetDeviceOk returns a tuple with the Device field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) GetDeviceOk() (*string, bool) {
	if o == nil || IsNil(o.Device) {
		return nil, false
	}
	return o.Device, true
}

// HasDevice returns a boolean if a field has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) HasDevice() bool {
	if o != nil && !IsNil(o.Device) {
		return true
	}

	return false
}

// SetDevice gets a reference to the given string and assigns it to the Device field.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) SetDevice(v string) {
	o.Device = &v
}

// GetDeviceMarketingName returns the DeviceMarketingName field value if set, zero value otherwise.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) GetDeviceMarketingName() string {
	if o == nil || IsNil(o.DeviceMarketingName) {
		var ret string
		return ret
	}
	return *o.DeviceMarketingName
}

// GetDeviceMarketingNameOk returns a tuple with the DeviceMarketingName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) GetDeviceMarketingNameOk() (*string, bool) {
	if o == nil || IsNil(o.DeviceMarketingName) {
		return nil, false
	}
	return o.DeviceMarketingName, true
}

// HasDeviceMarketingName returns a boolean if a field has been set.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) HasDeviceMarketingName() bool {
	if o != nil && !IsNil(o.DeviceMarketingName) {
		return true
	}

	return false
}

// SetDeviceMarketingName gets a reference to the given string and assigns it to the DeviceMarketingName field.
func (o *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) SetDeviceMarketingName(v string) {
	o.DeviceMarketingName = &v
}

func (o XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Percentile) {
		toSerialize["percentile"] = o.Percentile
	}
	if !IsNil(o.Device) {
		toSerialize["device"] = o.Device
	}
	if !IsNil(o.DeviceMarketingName) {
		toSerialize["deviceMarketingName"] = o.DeviceMarketingName
	}
	return toSerialize, nil
}

type NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria struct {
	value *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria
	isSet bool
}

func (v NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) Get() *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria {
	return v.value
}

func (v *NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) Set(val *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria(val *XcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) *NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria {
	return &NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria{value: val, isSet: true}
}

func (v NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableXcodeMetricsProductDataInnerMetricCategoriesInnerMetricsInnerDatasetsInnerFilterCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
