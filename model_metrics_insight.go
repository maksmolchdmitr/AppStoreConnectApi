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

// checks if the MetricsInsight type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MetricsInsight{}

// MetricsInsight struct for MetricsInsight
type MetricsInsight struct {
	MetricCategory        *MetricCategory                  `json:"metricCategory,omitempty"`
	LatestVersion         *string                          `json:"latestVersion,omitempty"`
	Metric                *string                          `json:"metric,omitempty"`
	SummaryString         *string                          `json:"summaryString,omitempty"`
	ReferenceVersions     *string                          `json:"referenceVersions,omitempty"`
	MaxLatestVersionValue *float32                         `json:"maxLatestVersionValue,omitempty"`
	SubSystemLabel        *string                          `json:"subSystemLabel,omitempty"`
	HighImpact            *bool                            `json:"highImpact,omitempty"`
	Populations           []MetricsInsightPopulationsInner `json:"populations,omitempty"`
}

// NewMetricsInsight instantiates a new MetricsInsight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMetricsInsight() *MetricsInsight {
	this := MetricsInsight{}
	return &this
}

// NewMetricsInsightWithDefaults instantiates a new MetricsInsight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMetricsInsightWithDefaults() *MetricsInsight {
	this := MetricsInsight{}
	return &this
}

// GetMetricCategory returns the MetricCategory field value if set, zero value otherwise.
func (o *MetricsInsight) GetMetricCategory() MetricCategory {
	if o == nil || IsNil(o.MetricCategory) {
		var ret MetricCategory
		return ret
	}
	return *o.MetricCategory
}

// GetMetricCategoryOk returns a tuple with the MetricCategory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsInsight) GetMetricCategoryOk() (*MetricCategory, bool) {
	if o == nil || IsNil(o.MetricCategory) {
		return nil, false
	}
	return o.MetricCategory, true
}

// HasMetricCategory returns a boolean if a field has been set.
func (o *MetricsInsight) HasMetricCategory() bool {
	if o != nil && !IsNil(o.MetricCategory) {
		return true
	}

	return false
}

// SetMetricCategory gets a reference to the given MetricCategory and assigns it to the MetricCategory field.
func (o *MetricsInsight) SetMetricCategory(v MetricCategory) {
	o.MetricCategory = &v
}

// GetLatestVersion returns the LatestVersion field value if set, zero value otherwise.
func (o *MetricsInsight) GetLatestVersion() string {
	if o == nil || IsNil(o.LatestVersion) {
		var ret string
		return ret
	}
	return *o.LatestVersion
}

// GetLatestVersionOk returns a tuple with the LatestVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsInsight) GetLatestVersionOk() (*string, bool) {
	if o == nil || IsNil(o.LatestVersion) {
		return nil, false
	}
	return o.LatestVersion, true
}

// HasLatestVersion returns a boolean if a field has been set.
func (o *MetricsInsight) HasLatestVersion() bool {
	if o != nil && !IsNil(o.LatestVersion) {
		return true
	}

	return false
}

// SetLatestVersion gets a reference to the given string and assigns it to the LatestVersion field.
func (o *MetricsInsight) SetLatestVersion(v string) {
	o.LatestVersion = &v
}

// GetMetric returns the Metric field value if set, zero value otherwise.
func (o *MetricsInsight) GetMetric() string {
	if o == nil || IsNil(o.Metric) {
		var ret string
		return ret
	}
	return *o.Metric
}

// GetMetricOk returns a tuple with the Metric field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsInsight) GetMetricOk() (*string, bool) {
	if o == nil || IsNil(o.Metric) {
		return nil, false
	}
	return o.Metric, true
}

// HasMetric returns a boolean if a field has been set.
func (o *MetricsInsight) HasMetric() bool {
	if o != nil && !IsNil(o.Metric) {
		return true
	}

	return false
}

// SetMetric gets a reference to the given string and assigns it to the Metric field.
func (o *MetricsInsight) SetMetric(v string) {
	o.Metric = &v
}

// GetSummaryString returns the SummaryString field value if set, zero value otherwise.
func (o *MetricsInsight) GetSummaryString() string {
	if o == nil || IsNil(o.SummaryString) {
		var ret string
		return ret
	}
	return *o.SummaryString
}

// GetSummaryStringOk returns a tuple with the SummaryString field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsInsight) GetSummaryStringOk() (*string, bool) {
	if o == nil || IsNil(o.SummaryString) {
		return nil, false
	}
	return o.SummaryString, true
}

// HasSummaryString returns a boolean if a field has been set.
func (o *MetricsInsight) HasSummaryString() bool {
	if o != nil && !IsNil(o.SummaryString) {
		return true
	}

	return false
}

// SetSummaryString gets a reference to the given string and assigns it to the SummaryString field.
func (o *MetricsInsight) SetSummaryString(v string) {
	o.SummaryString = &v
}

// GetReferenceVersions returns the ReferenceVersions field value if set, zero value otherwise.
func (o *MetricsInsight) GetReferenceVersions() string {
	if o == nil || IsNil(o.ReferenceVersions) {
		var ret string
		return ret
	}
	return *o.ReferenceVersions
}

// GetReferenceVersionsOk returns a tuple with the ReferenceVersions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsInsight) GetReferenceVersionsOk() (*string, bool) {
	if o == nil || IsNil(o.ReferenceVersions) {
		return nil, false
	}
	return o.ReferenceVersions, true
}

// HasReferenceVersions returns a boolean if a field has been set.
func (o *MetricsInsight) HasReferenceVersions() bool {
	if o != nil && !IsNil(o.ReferenceVersions) {
		return true
	}

	return false
}

// SetReferenceVersions gets a reference to the given string and assigns it to the ReferenceVersions field.
func (o *MetricsInsight) SetReferenceVersions(v string) {
	o.ReferenceVersions = &v
}

// GetMaxLatestVersionValue returns the MaxLatestVersionValue field value if set, zero value otherwise.
func (o *MetricsInsight) GetMaxLatestVersionValue() float32 {
	if o == nil || IsNil(o.MaxLatestVersionValue) {
		var ret float32
		return ret
	}
	return *o.MaxLatestVersionValue
}

// GetMaxLatestVersionValueOk returns a tuple with the MaxLatestVersionValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsInsight) GetMaxLatestVersionValueOk() (*float32, bool) {
	if o == nil || IsNil(o.MaxLatestVersionValue) {
		return nil, false
	}
	return o.MaxLatestVersionValue, true
}

// HasMaxLatestVersionValue returns a boolean if a field has been set.
func (o *MetricsInsight) HasMaxLatestVersionValue() bool {
	if o != nil && !IsNil(o.MaxLatestVersionValue) {
		return true
	}

	return false
}

// SetMaxLatestVersionValue gets a reference to the given float32 and assigns it to the MaxLatestVersionValue field.
func (o *MetricsInsight) SetMaxLatestVersionValue(v float32) {
	o.MaxLatestVersionValue = &v
}

// GetSubSystemLabel returns the SubSystemLabel field value if set, zero value otherwise.
func (o *MetricsInsight) GetSubSystemLabel() string {
	if o == nil || IsNil(o.SubSystemLabel) {
		var ret string
		return ret
	}
	return *o.SubSystemLabel
}

// GetSubSystemLabelOk returns a tuple with the SubSystemLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsInsight) GetSubSystemLabelOk() (*string, bool) {
	if o == nil || IsNil(o.SubSystemLabel) {
		return nil, false
	}
	return o.SubSystemLabel, true
}

// HasSubSystemLabel returns a boolean if a field has been set.
func (o *MetricsInsight) HasSubSystemLabel() bool {
	if o != nil && !IsNil(o.SubSystemLabel) {
		return true
	}

	return false
}

// SetSubSystemLabel gets a reference to the given string and assigns it to the SubSystemLabel field.
func (o *MetricsInsight) SetSubSystemLabel(v string) {
	o.SubSystemLabel = &v
}

// GetHighImpact returns the HighImpact field value if set, zero value otherwise.
func (o *MetricsInsight) GetHighImpact() bool {
	if o == nil || IsNil(o.HighImpact) {
		var ret bool
		return ret
	}
	return *o.HighImpact
}

// GetHighImpactOk returns a tuple with the HighImpact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsInsight) GetHighImpactOk() (*bool, bool) {
	if o == nil || IsNil(o.HighImpact) {
		return nil, false
	}
	return o.HighImpact, true
}

// HasHighImpact returns a boolean if a field has been set.
func (o *MetricsInsight) HasHighImpact() bool {
	if o != nil && !IsNil(o.HighImpact) {
		return true
	}

	return false
}

// SetHighImpact gets a reference to the given bool and assigns it to the HighImpact field.
func (o *MetricsInsight) SetHighImpact(v bool) {
	o.HighImpact = &v
}

// GetPopulations returns the Populations field value if set, zero value otherwise.
func (o *MetricsInsight) GetPopulations() []MetricsInsightPopulationsInner {
	if o == nil || IsNil(o.Populations) {
		var ret []MetricsInsightPopulationsInner
		return ret
	}
	return o.Populations
}

// GetPopulationsOk returns a tuple with the Populations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MetricsInsight) GetPopulationsOk() ([]MetricsInsightPopulationsInner, bool) {
	if o == nil || IsNil(o.Populations) {
		return nil, false
	}
	return o.Populations, true
}

// HasPopulations returns a boolean if a field has been set.
func (o *MetricsInsight) HasPopulations() bool {
	if o != nil && !IsNil(o.Populations) {
		return true
	}

	return false
}

// SetPopulations gets a reference to the given []MetricsInsightPopulationsInner and assigns it to the Populations field.
func (o *MetricsInsight) SetPopulations(v []MetricsInsightPopulationsInner) {
	o.Populations = v
}

func (o MetricsInsight) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MetricsInsight) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MetricCategory) {
		toSerialize["metricCategory"] = o.MetricCategory
	}
	if !IsNil(o.LatestVersion) {
		toSerialize["latestVersion"] = o.LatestVersion
	}
	if !IsNil(o.Metric) {
		toSerialize["metric"] = o.Metric
	}
	if !IsNil(o.SummaryString) {
		toSerialize["summaryString"] = o.SummaryString
	}
	if !IsNil(o.ReferenceVersions) {
		toSerialize["referenceVersions"] = o.ReferenceVersions
	}
	if !IsNil(o.MaxLatestVersionValue) {
		toSerialize["maxLatestVersionValue"] = o.MaxLatestVersionValue
	}
	if !IsNil(o.SubSystemLabel) {
		toSerialize["subSystemLabel"] = o.SubSystemLabel
	}
	if !IsNil(o.HighImpact) {
		toSerialize["highImpact"] = o.HighImpact
	}
	if !IsNil(o.Populations) {
		toSerialize["populations"] = o.Populations
	}
	return toSerialize, nil
}

type NullableMetricsInsight struct {
	value *MetricsInsight
	isSet bool
}

func (v NullableMetricsInsight) Get() *MetricsInsight {
	return v.value
}

func (v *NullableMetricsInsight) Set(val *MetricsInsight) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricsInsight) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricsInsight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricsInsight(val *MetricsInsight) *NullableMetricsInsight {
	return &NullableMetricsInsight{value: val, isSet: true}
}

func (v NullableMetricsInsight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricsInsight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}