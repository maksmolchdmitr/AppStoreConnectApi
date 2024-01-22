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

// checks if the AppScreenshotSetCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppScreenshotSetCreateRequestDataAttributes{}

// AppScreenshotSetCreateRequestDataAttributes struct for AppScreenshotSetCreateRequestDataAttributes
type AppScreenshotSetCreateRequestDataAttributes struct {
	ScreenshotDisplayType ScreenshotDisplayType `json:"screenshotDisplayType"`
}

// NewAppScreenshotSetCreateRequestDataAttributes instantiates a new AppScreenshotSetCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppScreenshotSetCreateRequestDataAttributes(screenshotDisplayType ScreenshotDisplayType) *AppScreenshotSetCreateRequestDataAttributes {
	this := AppScreenshotSetCreateRequestDataAttributes{}
	this.ScreenshotDisplayType = screenshotDisplayType
	return &this
}

// NewAppScreenshotSetCreateRequestDataAttributesWithDefaults instantiates a new AppScreenshotSetCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppScreenshotSetCreateRequestDataAttributesWithDefaults() *AppScreenshotSetCreateRequestDataAttributes {
	this := AppScreenshotSetCreateRequestDataAttributes{}
	return &this
}

// GetScreenshotDisplayType returns the ScreenshotDisplayType field value
func (o *AppScreenshotSetCreateRequestDataAttributes) GetScreenshotDisplayType() ScreenshotDisplayType {
	if o == nil {
		var ret ScreenshotDisplayType
		return ret
	}

	return o.ScreenshotDisplayType
}

// GetScreenshotDisplayTypeOk returns a tuple with the ScreenshotDisplayType field value
// and a boolean to check if the value has been set.
func (o *AppScreenshotSetCreateRequestDataAttributes) GetScreenshotDisplayTypeOk() (*ScreenshotDisplayType, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ScreenshotDisplayType, true
}

// SetScreenshotDisplayType sets field value
func (o *AppScreenshotSetCreateRequestDataAttributes) SetScreenshotDisplayType(v ScreenshotDisplayType) {
	o.ScreenshotDisplayType = v
}

func (o AppScreenshotSetCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppScreenshotSetCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["screenshotDisplayType"] = o.ScreenshotDisplayType
	return toSerialize, nil
}

type NullableAppScreenshotSetCreateRequestDataAttributes struct {
	value *AppScreenshotSetCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppScreenshotSetCreateRequestDataAttributes) Get() *AppScreenshotSetCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppScreenshotSetCreateRequestDataAttributes) Set(val *AppScreenshotSetCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppScreenshotSetCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppScreenshotSetCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppScreenshotSetCreateRequestDataAttributes(val *AppScreenshotSetCreateRequestDataAttributes) *NullableAppScreenshotSetCreateRequestDataAttributes {
	return &NullableAppScreenshotSetCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppScreenshotSetCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppScreenshotSetCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


