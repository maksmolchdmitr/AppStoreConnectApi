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

// checks if the AppCustomProductPageCreateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageCreateRequestDataAttributes{}

// AppCustomProductPageCreateRequestDataAttributes struct for AppCustomProductPageCreateRequestDataAttributes
type AppCustomProductPageCreateRequestDataAttributes struct {
	Name string `json:"name"`
}

// NewAppCustomProductPageCreateRequestDataAttributes instantiates a new AppCustomProductPageCreateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageCreateRequestDataAttributes(name string) *AppCustomProductPageCreateRequestDataAttributes {
	this := AppCustomProductPageCreateRequestDataAttributes{}
	this.Name = name
	return &this
}

// NewAppCustomProductPageCreateRequestDataAttributesWithDefaults instantiates a new AppCustomProductPageCreateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageCreateRequestDataAttributesWithDefaults() *AppCustomProductPageCreateRequestDataAttributes {
	this := AppCustomProductPageCreateRequestDataAttributes{}
	return &this
}

// GetName returns the Name field value
func (o *AppCustomProductPageCreateRequestDataAttributes) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageCreateRequestDataAttributes) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *AppCustomProductPageCreateRequestDataAttributes) SetName(v string) {
	o.Name = v
}

func (o AppCustomProductPageCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageCreateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableAppCustomProductPageCreateRequestDataAttributes struct {
	value *AppCustomProductPageCreateRequestDataAttributes
	isSet bool
}

func (v NullableAppCustomProductPageCreateRequestDataAttributes) Get() *AppCustomProductPageCreateRequestDataAttributes {
	return v.value
}

func (v *NullableAppCustomProductPageCreateRequestDataAttributes) Set(val *AppCustomProductPageCreateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageCreateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageCreateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageCreateRequestDataAttributes(val *AppCustomProductPageCreateRequestDataAttributes) *NullableAppCustomProductPageCreateRequestDataAttributes {
	return &NullableAppCustomProductPageCreateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableAppCustomProductPageCreateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageCreateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
