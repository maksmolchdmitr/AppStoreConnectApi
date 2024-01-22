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

// checks if the AppClipAppStoreReviewDetailAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipAppStoreReviewDetailAttributes{}

// AppClipAppStoreReviewDetailAttributes struct for AppClipAppStoreReviewDetailAttributes
type AppClipAppStoreReviewDetailAttributes struct {
	InvocationUrls []string `json:"invocationUrls,omitempty"`
}

// NewAppClipAppStoreReviewDetailAttributes instantiates a new AppClipAppStoreReviewDetailAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipAppStoreReviewDetailAttributes() *AppClipAppStoreReviewDetailAttributes {
	this := AppClipAppStoreReviewDetailAttributes{}
	return &this
}

// NewAppClipAppStoreReviewDetailAttributesWithDefaults instantiates a new AppClipAppStoreReviewDetailAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipAppStoreReviewDetailAttributesWithDefaults() *AppClipAppStoreReviewDetailAttributes {
	this := AppClipAppStoreReviewDetailAttributes{}
	return &this
}

// GetInvocationUrls returns the InvocationUrls field value if set, zero value otherwise.
func (o *AppClipAppStoreReviewDetailAttributes) GetInvocationUrls() []string {
	if o == nil || IsNil(o.InvocationUrls) {
		var ret []string
		return ret
	}
	return o.InvocationUrls
}

// GetInvocationUrlsOk returns a tuple with the InvocationUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppClipAppStoreReviewDetailAttributes) GetInvocationUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.InvocationUrls) {
		return nil, false
	}
	return o.InvocationUrls, true
}

// HasInvocationUrls returns a boolean if a field has been set.
func (o *AppClipAppStoreReviewDetailAttributes) HasInvocationUrls() bool {
	if o != nil && !IsNil(o.InvocationUrls) {
		return true
	}

	return false
}

// SetInvocationUrls gets a reference to the given []string and assigns it to the InvocationUrls field.
func (o *AppClipAppStoreReviewDetailAttributes) SetInvocationUrls(v []string) {
	o.InvocationUrls = v
}

func (o AppClipAppStoreReviewDetailAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipAppStoreReviewDetailAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.InvocationUrls) {
		toSerialize["invocationUrls"] = o.InvocationUrls
	}
	return toSerialize, nil
}

type NullableAppClipAppStoreReviewDetailAttributes struct {
	value *AppClipAppStoreReviewDetailAttributes
	isSet bool
}

func (v NullableAppClipAppStoreReviewDetailAttributes) Get() *AppClipAppStoreReviewDetailAttributes {
	return v.value
}

func (v *NullableAppClipAppStoreReviewDetailAttributes) Set(val *AppClipAppStoreReviewDetailAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipAppStoreReviewDetailAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipAppStoreReviewDetailAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipAppStoreReviewDetailAttributes(val *AppClipAppStoreReviewDetailAttributes) *NullableAppClipAppStoreReviewDetailAttributes {
	return &NullableAppClipAppStoreReviewDetailAttributes{value: val, isSet: true}
}

func (v NullableAppClipAppStoreReviewDetailAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipAppStoreReviewDetailAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


