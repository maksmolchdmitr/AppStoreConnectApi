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

// checks if the AppCustomProductPageLocalizationInlineCreateAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageLocalizationInlineCreateAttributes{}

// AppCustomProductPageLocalizationInlineCreateAttributes struct for AppCustomProductPageLocalizationInlineCreateAttributes
type AppCustomProductPageLocalizationInlineCreateAttributes struct {
	Locale          string  `json:"locale"`
	PromotionalText *string `json:"promotionalText,omitempty"`
}

// NewAppCustomProductPageLocalizationInlineCreateAttributes instantiates a new AppCustomProductPageLocalizationInlineCreateAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageLocalizationInlineCreateAttributes(locale string) *AppCustomProductPageLocalizationInlineCreateAttributes {
	this := AppCustomProductPageLocalizationInlineCreateAttributes{}
	this.Locale = locale
	return &this
}

// NewAppCustomProductPageLocalizationInlineCreateAttributesWithDefaults instantiates a new AppCustomProductPageLocalizationInlineCreateAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageLocalizationInlineCreateAttributesWithDefaults() *AppCustomProductPageLocalizationInlineCreateAttributes {
	this := AppCustomProductPageLocalizationInlineCreateAttributes{}
	return &this
}

// GetLocale returns the Locale field value
func (o *AppCustomProductPageLocalizationInlineCreateAttributes) GetLocale() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationInlineCreateAttributes) GetLocaleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Locale, true
}

// SetLocale sets field value
func (o *AppCustomProductPageLocalizationInlineCreateAttributes) SetLocale(v string) {
	o.Locale = v
}

// GetPromotionalText returns the PromotionalText field value if set, zero value otherwise.
func (o *AppCustomProductPageLocalizationInlineCreateAttributes) GetPromotionalText() string {
	if o == nil || IsNil(o.PromotionalText) {
		var ret string
		return ret
	}
	return *o.PromotionalText
}

// GetPromotionalTextOk returns a tuple with the PromotionalText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationInlineCreateAttributes) GetPromotionalTextOk() (*string, bool) {
	if o == nil || IsNil(o.PromotionalText) {
		return nil, false
	}
	return o.PromotionalText, true
}

// HasPromotionalText returns a boolean if a field has been set.
func (o *AppCustomProductPageLocalizationInlineCreateAttributes) HasPromotionalText() bool {
	if o != nil && !IsNil(o.PromotionalText) {
		return true
	}

	return false
}

// SetPromotionalText gets a reference to the given string and assigns it to the PromotionalText field.
func (o *AppCustomProductPageLocalizationInlineCreateAttributes) SetPromotionalText(v string) {
	o.PromotionalText = &v
}

func (o AppCustomProductPageLocalizationInlineCreateAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageLocalizationInlineCreateAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["locale"] = o.Locale
	if !IsNil(o.PromotionalText) {
		toSerialize["promotionalText"] = o.PromotionalText
	}
	return toSerialize, nil
}

type NullableAppCustomProductPageLocalizationInlineCreateAttributes struct {
	value *AppCustomProductPageLocalizationInlineCreateAttributes
	isSet bool
}

func (v NullableAppCustomProductPageLocalizationInlineCreateAttributes) Get() *AppCustomProductPageLocalizationInlineCreateAttributes {
	return v.value
}

func (v *NullableAppCustomProductPageLocalizationInlineCreateAttributes) Set(val *AppCustomProductPageLocalizationInlineCreateAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageLocalizationInlineCreateAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageLocalizationInlineCreateAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageLocalizationInlineCreateAttributes(val *AppCustomProductPageLocalizationInlineCreateAttributes) *NullableAppCustomProductPageLocalizationInlineCreateAttributes {
	return &NullableAppCustomProductPageLocalizationInlineCreateAttributes{value: val, isSet: true}
}

func (v NullableAppCustomProductPageLocalizationInlineCreateAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageLocalizationInlineCreateAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}