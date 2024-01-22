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

// checks if the AppScreenshotSetAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppScreenshotSetAttributes{}

// AppScreenshotSetAttributes struct for AppScreenshotSetAttributes
type AppScreenshotSetAttributes struct {
	ScreenshotDisplayType *ScreenshotDisplayType `json:"screenshotDisplayType,omitempty"`
}

// NewAppScreenshotSetAttributes instantiates a new AppScreenshotSetAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotSetAttributes() *AppScreenshotSetAttributes {
	this := AppScreenshotSetAttributes{}
	return &this
}

// NewAppScreenshotSetAttributesWithDefaults instantiates a new AppScreenshotSetAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotSetAttributesWithDefaults() *AppScreenshotSetAttributes {
	this := AppScreenshotSetAttributes{}
	return &this
}

// GetScreenshotDisplayType returns the ScreenshotDisplayType field value if set, zero value otherwise.
func (o *AppScreenshotSetAttributes) GetScreenshotDisplayType() ScreenshotDisplayType {
	if o == nil || IsNil(o.ScreenshotDisplayType) {
		var ret ScreenshotDisplayType
		return ret
	}
	return *o.ScreenshotDisplayType
}

// GetScreenshotDisplayTypeOk returns a tuple with the ScreenshotDisplayType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetAttributes) GetScreenshotDisplayTypeOk() (*ScreenshotDisplayType, bool) {
	if o == nil || IsNil(o.ScreenshotDisplayType) {
		return nil, false
	}
	return o.ScreenshotDisplayType, true
}

// HasScreenshotDisplayType returns a boolean if a field has been set.
func (o *AppScreenshotSetAttributes) HasScreenshotDisplayType() bool {
	if o != nil && !IsNil(o.ScreenshotDisplayType) {
		return true
	}

	return false
}

// SetScreenshotDisplayType gets a reference to the given ScreenshotDisplayType and assigns it to the ScreenshotDisplayType field.
func (o *AppScreenshotSetAttributes) SetScreenshotDisplayType(v ScreenshotDisplayType) {
	o.ScreenshotDisplayType = &v
}

func (o AppScreenshotSetAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppScreenshotSetAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ScreenshotDisplayType) {
		toSerialize["screenshotDisplayType"] = o.ScreenshotDisplayType
	}
	return toSerialize, nil
}

type NullableAppScreenshotSetAttributes struct {
	value *AppScreenshotSetAttributes
	isSet bool
}

func (v NullableAppScreenshotSetAttributes) Get() *AppScreenshotSetAttributes {
	return v.value
}

func (v *NullableAppScreenshotSetAttributes) Set(val *AppScreenshotSetAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotSetAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotSetAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotSetAttributes(val *AppScreenshotSetAttributes) *NullableAppScreenshotSetAttributes {
	return &NullableAppScreenshotSetAttributes{value: val, isSet: true}
}

func (v NullableAppScreenshotSetAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotSetAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
