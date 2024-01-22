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

// checks if the InAppPurchasePriceAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchasePriceAttributes{}

// InAppPurchasePriceAttributes struct for InAppPurchasePriceAttributes
type InAppPurchasePriceAttributes struct {
	StartDate *string `json:"startDate,omitempty"`
	EndDate   *string `json:"endDate,omitempty"`
	Manual    *bool   `json:"manual,omitempty"`
}

// NewInAppPurchasePriceAttributes instantiates a new InAppPurchasePriceAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchasePriceAttributes() *InAppPurchasePriceAttributes {
	this := InAppPurchasePriceAttributes{}
	return &this
}

// NewInAppPurchasePriceAttributesWithDefaults instantiates a new InAppPurchasePriceAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchasePriceAttributesWithDefaults() *InAppPurchasePriceAttributes {
	this := InAppPurchasePriceAttributes{}
	return &this
}

// GetStartDate returns the StartDate field value if set, zero value otherwise.
func (o *InAppPurchasePriceAttributes) GetStartDate() string {
	if o == nil || IsNil(o.StartDate) {
		var ret string
		return ret
	}
	return *o.StartDate
}

// GetStartDateOk returns a tuple with the StartDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceAttributes) GetStartDateOk() (*string, bool) {
	if o == nil || IsNil(o.StartDate) {
		return nil, false
	}
	return o.StartDate, true
}

// HasStartDate returns a boolean if a field has been set.
func (o *InAppPurchasePriceAttributes) HasStartDate() bool {
	if o != nil && !IsNil(o.StartDate) {
		return true
	}

	return false
}

// SetStartDate gets a reference to the given string and assigns it to the StartDate field.
func (o *InAppPurchasePriceAttributes) SetStartDate(v string) {
	o.StartDate = &v
}

// GetEndDate returns the EndDate field value if set, zero value otherwise.
func (o *InAppPurchasePriceAttributes) GetEndDate() string {
	if o == nil || IsNil(o.EndDate) {
		var ret string
		return ret
	}
	return *o.EndDate
}

// GetEndDateOk returns a tuple with the EndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceAttributes) GetEndDateOk() (*string, bool) {
	if o == nil || IsNil(o.EndDate) {
		return nil, false
	}
	return o.EndDate, true
}

// HasEndDate returns a boolean if a field has been set.
func (o *InAppPurchasePriceAttributes) HasEndDate() bool {
	if o != nil && !IsNil(o.EndDate) {
		return true
	}

	return false
}

// SetEndDate gets a reference to the given string and assigns it to the EndDate field.
func (o *InAppPurchasePriceAttributes) SetEndDate(v string) {
	o.EndDate = &v
}

// GetManual returns the Manual field value if set, zero value otherwise.
func (o *InAppPurchasePriceAttributes) GetManual() bool {
	if o == nil || IsNil(o.Manual) {
		var ret bool
		return ret
	}
	return *o.Manual
}

// GetManualOk returns a tuple with the Manual field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InAppPurchasePriceAttributes) GetManualOk() (*bool, bool) {
	if o == nil || IsNil(o.Manual) {
		return nil, false
	}
	return o.Manual, true
}

// HasManual returns a boolean if a field has been set.
func (o *InAppPurchasePriceAttributes) HasManual() bool {
	if o != nil && !IsNil(o.Manual) {
		return true
	}

	return false
}

// SetManual gets a reference to the given bool and assigns it to the Manual field.
func (o *InAppPurchasePriceAttributes) SetManual(v bool) {
	o.Manual = &v
}

func (o InAppPurchasePriceAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchasePriceAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.StartDate) {
		toSerialize["startDate"] = o.StartDate
	}
	if !IsNil(o.EndDate) {
		toSerialize["endDate"] = o.EndDate
	}
	if !IsNil(o.Manual) {
		toSerialize["manual"] = o.Manual
	}
	return toSerialize, nil
}

type NullableInAppPurchasePriceAttributes struct {
	value *InAppPurchasePriceAttributes
	isSet bool
}

func (v NullableInAppPurchasePriceAttributes) Get() *InAppPurchasePriceAttributes {
	return v.value
}

func (v *NullableInAppPurchasePriceAttributes) Set(val *InAppPurchasePriceAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchasePriceAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchasePriceAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchasePriceAttributes(val *InAppPurchasePriceAttributes) *NullableInAppPurchasePriceAttributes {
	return &NullableInAppPurchasePriceAttributes{value: val, isSet: true}
}

func (v NullableInAppPurchasePriceAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchasePriceAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}