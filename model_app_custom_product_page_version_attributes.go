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

// checks if the AppCustomProductPageVersionAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageVersionAttributes{}

// AppCustomProductPageVersionAttributes struct for AppCustomProductPageVersionAttributes
type AppCustomProductPageVersionAttributes struct {
	Version *string `json:"version,omitempty"`
	State   *string `json:"state,omitempty"`
}

// NewAppCustomProductPageVersionAttributes instantiates a new AppCustomProductPageVersionAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageVersionAttributes() *AppCustomProductPageVersionAttributes {
	this := AppCustomProductPageVersionAttributes{}
	return &this
}

// NewAppCustomProductPageVersionAttributesWithDefaults instantiates a new AppCustomProductPageVersionAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageVersionAttributesWithDefaults() *AppCustomProductPageVersionAttributes {
	this := AppCustomProductPageVersionAttributes{}
	return &this
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *AppCustomProductPageVersionAttributes) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionAttributes) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *AppCustomProductPageVersionAttributes) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *AppCustomProductPageVersionAttributes) SetVersion(v string) {
	o.Version = &v
}

// GetState returns the State field value if set, zero value otherwise.
func (o *AppCustomProductPageVersionAttributes) GetState() string {
	if o == nil || IsNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionAttributes) GetStateOk() (*string, bool) {
	if o == nil || IsNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *AppCustomProductPageVersionAttributes) HasState() bool {
	if o != nil && !IsNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *AppCustomProductPageVersionAttributes) SetState(v string) {
	o.State = &v
}

func (o AppCustomProductPageVersionAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageVersionAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	if !IsNil(o.State) {
		toSerialize["state"] = o.State
	}
	return toSerialize, nil
}

type NullableAppCustomProductPageVersionAttributes struct {
	value *AppCustomProductPageVersionAttributes
	isSet bool
}

func (v NullableAppCustomProductPageVersionAttributes) Get() *AppCustomProductPageVersionAttributes {
	return v.value
}

func (v *NullableAppCustomProductPageVersionAttributes) Set(val *AppCustomProductPageVersionAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageVersionAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageVersionAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageVersionAttributes(val *AppCustomProductPageVersionAttributes) *NullableAppCustomProductPageVersionAttributes {
	return &NullableAppCustomProductPageVersionAttributes{value: val, isSet: true}
}

func (v NullableAppCustomProductPageVersionAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageVersionAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
