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

// checks if the AppPreviewRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreviewRelationships{}

// AppPreviewRelationships struct for AppPreviewRelationships
type AppPreviewRelationships struct {
	AppPreviewSet *AppPreviewRelationshipsAppPreviewSet `json:"appPreviewSet,omitempty"`
}

// NewAppPreviewRelationships instantiates a new AppPreviewRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreviewRelationships() *AppPreviewRelationships {
	this := AppPreviewRelationships{}
	return &this
}

// NewAppPreviewRelationshipsWithDefaults instantiates a new AppPreviewRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreviewRelationshipsWithDefaults() *AppPreviewRelationships {
	this := AppPreviewRelationships{}
	return &this
}

// GetAppPreviewSet returns the AppPreviewSet field value if set, zero value otherwise.
func (o *AppPreviewRelationships) GetAppPreviewSet() AppPreviewRelationshipsAppPreviewSet {
	if o == nil || IsNil(o.AppPreviewSet) {
		var ret AppPreviewRelationshipsAppPreviewSet
		return ret
	}
	return *o.AppPreviewSet
}

// GetAppPreviewSetOk returns a tuple with the AppPreviewSet field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreviewRelationships) GetAppPreviewSetOk() (*AppPreviewRelationshipsAppPreviewSet, bool) {
	if o == nil || IsNil(o.AppPreviewSet) {
		return nil, false
	}
	return o.AppPreviewSet, true
}

// HasAppPreviewSet returns a boolean if a field has been set.
func (o *AppPreviewRelationships) HasAppPreviewSet() bool {
	if o != nil && !IsNil(o.AppPreviewSet) {
		return true
	}

	return false
}

// SetAppPreviewSet gets a reference to the given AppPreviewRelationshipsAppPreviewSet and assigns it to the AppPreviewSet field.
func (o *AppPreviewRelationships) SetAppPreviewSet(v AppPreviewRelationshipsAppPreviewSet) {
	o.AppPreviewSet = &v
}

func (o AppPreviewRelationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreviewRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AppPreviewSet) {
		toSerialize["appPreviewSet"] = o.AppPreviewSet
	}
	return toSerialize, nil
}

type NullableAppPreviewRelationships struct {
	value *AppPreviewRelationships
	isSet bool
}

func (v NullableAppPreviewRelationships) Get() *AppPreviewRelationships {
	return v.value
}

func (v *NullableAppPreviewRelationships) Set(val *AppPreviewRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreviewRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreviewRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreviewRelationships(val *AppPreviewRelationships) *NullableAppPreviewRelationships {
	return &NullableAppPreviewRelationships{value: val, isSet: true}
}

func (v NullableAppPreviewRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreviewRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
