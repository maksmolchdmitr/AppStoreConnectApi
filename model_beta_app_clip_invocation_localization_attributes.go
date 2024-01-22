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

// checks if the BetaAppClipInvocationLocalizationAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocationLocalizationAttributes{}

// BetaAppClipInvocationLocalizationAttributes struct for BetaAppClipInvocationLocalizationAttributes
type BetaAppClipInvocationLocalizationAttributes struct {
	Title  *string `json:"title,omitempty"`
	Locale *string `json:"locale,omitempty"`
}

// NewBetaAppClipInvocationLocalizationAttributes instantiates a new BetaAppClipInvocationLocalizationAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocationLocalizationAttributes() *BetaAppClipInvocationLocalizationAttributes {
	this := BetaAppClipInvocationLocalizationAttributes{}
	return &this
}

// NewBetaAppClipInvocationLocalizationAttributesWithDefaults instantiates a new BetaAppClipInvocationLocalizationAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationLocalizationAttributesWithDefaults() *BetaAppClipInvocationLocalizationAttributes {
	this := BetaAppClipInvocationLocalizationAttributes{}
	return &this
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *BetaAppClipInvocationLocalizationAttributes) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationLocalizationAttributes) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *BetaAppClipInvocationLocalizationAttributes) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *BetaAppClipInvocationLocalizationAttributes) SetTitle(v string) {
	o.Title = &v
}

// GetLocale returns the Locale field value if set, zero value otherwise.
func (o *BetaAppClipInvocationLocalizationAttributes) GetLocale() string {
	if o == nil || IsNil(o.Locale) {
		var ret string
		return ret
	}
	return *o.Locale
}

// GetLocaleOk returns a tuple with the Locale field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationLocalizationAttributes) GetLocaleOk() (*string, bool) {
	if o == nil || IsNil(o.Locale) {
		return nil, false
	}
	return o.Locale, true
}

// HasLocale returns a boolean if a field has been set.
func (o *BetaAppClipInvocationLocalizationAttributes) HasLocale() bool {
	if o != nil && !IsNil(o.Locale) {
		return true
	}

	return false
}

// SetLocale gets a reference to the given string and assigns it to the Locale field.
func (o *BetaAppClipInvocationLocalizationAttributes) SetLocale(v string) {
	o.Locale = &v
}

func (o BetaAppClipInvocationLocalizationAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocationLocalizationAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Locale) {
		toSerialize["locale"] = o.Locale
	}
	return toSerialize, nil
}

type NullableBetaAppClipInvocationLocalizationAttributes struct {
	value *BetaAppClipInvocationLocalizationAttributes
	isSet bool
}

func (v NullableBetaAppClipInvocationLocalizationAttributes) Get() *BetaAppClipInvocationLocalizationAttributes {
	return v.value
}

func (v *NullableBetaAppClipInvocationLocalizationAttributes) Set(val *BetaAppClipInvocationLocalizationAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocationLocalizationAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocationLocalizationAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocationLocalizationAttributes(val *BetaAppClipInvocationLocalizationAttributes) *NullableBetaAppClipInvocationLocalizationAttributes {
	return &NullableBetaAppClipInvocationLocalizationAttributes{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocationLocalizationAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocationLocalizationAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
