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

// checks if the AppCustomProductPageLocalizationAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageLocalizationAttributes{}

// AppCustomProductPageLocalizationAttributes struct for AppCustomProductPageLocalizationAttributes
type AppCustomProductPageLocalizationAttributes struct {
	Locale          *string `json:"locale,omitempty"`
	PromotionalText *string `json:"promotionalText,omitempty"`
}

// NewAppCustomProductPageLocalizationAttributes instantiates a new AppCustomProductPageLocalizationAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageLocalizationAttributes() *AppCustomProductPageLocalizationAttributes {
	this := AppCustomProductPageLocalizationAttributes{}
	return &this
}

// NewAppCustomProductPageLocalizationAttributesWithDefaults instantiates a new AppCustomProductPageLocalizationAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageLocalizationAttributesWithDefaults() *AppCustomProductPageLocalizationAttributes {
	this := AppCustomProductPageLocalizationAttributes{}
	return &this
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *AppCustomProductPageLocalizationAttributes) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationAttributes) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *AppCustomProductPageLocalizationAttributes) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *AppCustomProductPageLocalizationAttributes) SetLocale(v string) {
	o.Locale = &v
}

// GetPromotionalText returns the PromotionalText field value if set, zero value otherwise.
func (o *AppCustomProductPageLocalizationAttributes) GetPromotionalText() string {
	if o == nil || IsNil(o.PromotionalText) {
		var ret string
		return ret
	}
	return *o.PromotionalText
}

// GetPromotionalTextOk returns a tuple with the PromotionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationAttributes) GetPromotionalTextOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionalText) {
		return nil, false
	}
	return o.PromotionalText, true
}

// HasPromotionalText returns a boolean if a field has been set.
func (o *AppCustomProductPageLocalizationAttributes) HasPromotionalText() bool {
	if o != nil && !IsNil(o.PromotionalText) {
		return true
	}

	return false
}

// SetPromotionalText gets a reference to the given string and assigns it to the PromotionalText field.
func (o *AppCustomProductPageLocalizationAttributes) SetPromotionalText(v string) {
	o.PromotionalText = &v
}

func (o AppCustomProductPageLocalizationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageLocalizationAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	if !IsNil(o.PromotionalText) {
		toSerialize["promotionalText"] = o.PromotionalText
	}
	return toSerialize, nil
}

type NullableAppCustomProductPageLocalizationAttributes struct {
	value *AppCustomProductPageLocalizationAttributes
	isSet bool
}

func (v NullableAppCustomProductPageLocalizationAttributes) Get() *AppCustomProductPageLocalizationAttributes {
	return v.value
}

func (v *NullableAppCustomProductPageLocalizationAttributes) Set(val *AppCustomProductPageLocalizationAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageLocalizationAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageLocalizationAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageLocalizationAttributes(val *AppCustomProductPageLocalizationAttributes) *NullableAppCustomProductPageLocalizationAttributes {
	return &NullableAppCustomProductPageLocalizationAttributes{value: val, isSet: true}
}

func (v NullableAppCustomProductPageLocalizationAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageLocalizationAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}