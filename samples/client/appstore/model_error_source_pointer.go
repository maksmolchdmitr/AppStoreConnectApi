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

// checks if the ErrorSourcePointer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ErrorSourcePointer{}

// ErrorSourcePointer struct for ErrorSourcePointer
type ErrorSourcePointer struct {
	Pointer *string `json:"pointer,omitempty"`
}

// NewErrorSourcePointer instantiates a new ErrorSourcePointer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorSourcePointer() *ErrorSourcePointer {
	this := ErrorSourcePointer{}
	return &this
}

// NewErrorSourcePointerWithDefaults instantiates a new ErrorSourcePointer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorSourcePointerWithDefaults() *ErrorSourcePointer {
	this := ErrorSourcePointer{}
	return &this
}

// GetPointer returns the Pointer field value if set, zero value otherwise.
func (o *ErrorSourcePointer) GetPointer() string {
	if o == nil || IsNil(o.Pointer) {
		var ret string
		return ret
	}
	return *o.Pointer
}

// GetPointerOk returns a tuple with the Pointer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ErrorSourcePointer) GetPointerOk() (*string, bool) {
	if o == nil || IsNil(o.Pointer) {
		return nil, false
	}
	return o.Pointer, true
}

// HasPointer returns a boolean if a field has been set.
func (o *ErrorSourcePointer) HasPointer() bool {
	if o != nil && !IsNil(o.Pointer) {
		return true
	}

	return false
}

// SetPointer gets a reference to the given string and assigns it to the Pointer field.
func (o *ErrorSourcePointer) SetPointer(v string) {
	o.Pointer = &v
}

func (o ErrorSourcePointer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ErrorSourcePointer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Pointer) {
		toSerialize["pointer"] = o.Pointer
	}
	return toSerialize, nil
}

type NullableErrorSourcePointer struct {
	value *ErrorSourcePointer
	isSet bool
}

func (v NullableErrorSourcePointer) Get() *ErrorSourcePointer {
	return v.value
}

func (v *NullableErrorSourcePointer) Set(val *ErrorSourcePointer) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorSourcePointer) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorSourcePointer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorSourcePointer(val *ErrorSourcePointer) *NullableErrorSourcePointer {
	return &NullableErrorSourcePointer{value: val, isSet: true}
}

func (v NullableErrorSourcePointer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorSourcePointer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


