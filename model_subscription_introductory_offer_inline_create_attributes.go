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

// checks if the SubscriptionIntroductoryOfferInlineCreateAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionIntroductoryOfferInlineCreateAttributes{}

// SubscriptionIntroductoryOfferInlineCreateAttributes struct for SubscriptionIntroductoryOfferInlineCreateAttributes
type SubscriptionIntroductoryOfferInlineCreateAttributes struct {
	StartDate       *string                   `json:"startDate,omitempty"`
	EndDate         *string                   `json:"endDate,omitempty"`
	Duration        SubscriptionOfferDuration `json:"duration"`
	OfferMode       SubscriptionOfferMode     `json:"offerMode"`
	NumberOfPeriods int32                     `json:"numberOfPeriods"`
}

// NewSubscriptionIntroductoryOfferInlineCreateAttributes instantiates a new SubscriptionIntroductoryOfferInlineCreateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionIntroductoryOfferInlineCreateAttributes(duration SubscriptionOfferDuration, offerMode SubscriptionOfferMode, numberOfPeriods int32) *SubscriptionIntroductoryOfferInlineCreateAttributes {
	this := SubscriptionIntroductoryOfferInlineCreateAttributes{}
	this.Duration = duration
	this.OfferMode = offerMode
	this.NumberOfPeriods = numberOfPeriods
	return &this
}

// NewSubscriptionIntroductoryOfferInlineCreateAttributesWithDefaults instantiates a new SubscriptionIntroductoryOfferInlineCreateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionIntroductoryOfferInlineCreateAttributesWithDefaults() *SubscriptionIntroductoryOfferInlineCreateAttributes {
	this := SubscriptionIntroductoryOfferInlineCreateAttributes{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) SetEndDate(v string) {
	o.EndDate = &v
}

// GetDuration returns the Duration field value
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) GetDuration() SubscriptionOfferDuration {
	if o == nil {
		var ret SubscriptionOfferDuration
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) GetDurationOk() (*SubscriptionOfferDuration, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) SetDuration(v SubscriptionOfferDuration) {
	o.Duration = v
}

// GetOfferMode returns the OfferMode field value
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) GetOfferMode() SubscriptionOfferMode {
	if o == nil {
		var ret SubscriptionOfferMode
		return ret
	}

	return o.OfferMode
}

// GetOfferModeOk returns a tuple with the OfferMode field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) GetOfferModeOk() (*SubscriptionOfferMode, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OfferMode, true
}

// SetOfferMode sets field value
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) SetOfferMode(v SubscriptionOfferMode) {
	o.OfferMode = v
}

// GetNumberOfPeriods returns the NumberOfPeriods field value
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) GetNumberOfPeriods() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.NumberOfPeriods
}

// GetNumberOfPeriodsOk returns a tuple with the NumberOfPeriods field value
// and a boolean to check if the value has been set.
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) GetNumberOfPeriodsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NumberOfPeriods, true
}

// SetNumberOfPeriods sets field value
func (o *SubscriptionIntroductoryOfferInlineCreateAttributes) SetNumberOfPeriods(v int32) {
	o.NumberOfPeriods = v
}

func (o SubscriptionIntroductoryOfferInlineCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionIntroductoryOfferInlineCreateAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	toSerialize["duration"] = o.Duration
	toSerialize["offerMode"] = o.OfferMode
	toSerialize["numberOfPeriods"] = o.NumberOfPeriods
	return toSerialize, nil
}

type NullableSubscriptionIntroductoryOfferInlineCreateAttributes struct {
	value *SubscriptionIntroductoryOfferInlineCreateAttributes
	isSet bool
}

func (v NullableSubscriptionIntroductoryOfferInlineCreateAttributes) Get() *SubscriptionIntroductoryOfferInlineCreateAttributes {
	return v.value
}

func (v *NullableSubscriptionIntroductoryOfferInlineCreateAttributes) Set(val *SubscriptionIntroductoryOfferInlineCreateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionIntroductoryOfferInlineCreateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionIntroductoryOfferInlineCreateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionIntroductoryOfferInlineCreateAttributes(val *SubscriptionIntroductoryOfferInlineCreateAttributes) *NullableSubscriptionIntroductoryOfferInlineCreateAttributes {
	return &NullableSubscriptionIntroductoryOfferInlineCreateAttributes{value: val, isSet: true}
}

func (v NullableSubscriptionIntroductoryOfferInlineCreateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionIntroductoryOfferInlineCreateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}