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

// checks if the AppEventRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppEventRelationships{}

// AppEventRelationships struct for AppEventRelationships
type AppEventRelationships struct {
	Localizations *AppEventRelationshipsLocalizations `json:"localizations,omitempty"`
}

// NewAppEventRelationships instantiates a new AppEventRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppEventRelationships() *AppEventRelationships {
	this := AppEventRelationships{}
	return &this
}

// NewAppEventRelationshipsWithDefaults instantiates a new AppEventRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppEventRelationshipsWithDefaults() *AppEventRelationships {
	this := AppEventRelationships{}
	return &this
}

// GetLocalizations returns the Localizations field value if set, zero value otherwise.
func (o *AppEventRelationships) GetLocalizations() AppEventRelationshipsLocalizations {
	if o == nil || IsNil(o.Localizations) {
		var ret AppEventRelationshipsLocalizations
		return ret
	}
	return *o.Localizations
}

// GetLocalizationsOk returns a tuple with the Localizations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppEventRelationships) GetLocalizationsOk() (*AppEventRelationshipsLocalizations, bool) {
	if o == nil || IsNil(o.Localizations) {
		return nil, false
	}
	return o.Localizations, true
}

// HasLocalizations returns a boolean if a field has been set.
func (o *AppEventRelationships) HasLocalizations() bool {
	if o != nil && !IsNil(o.Localizations) {
		return true
	}

	return false
}

// SetLocalizations gets a reference to the given AppEventRelationshipsLocalizations and assigns it to the Localizations field.
func (o *AppEventRelationships) SetLocalizations(v AppEventRelationshipsLocalizations) {
	o.Localizations = &v
}

func (o AppEventRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppEventRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Localizations) {
		toSerialize["localizations"] = o.Localizations
	}
	return toSerialize, nil
}

type NullableAppEventRelationships struct {
	value *AppEventRelationships
	isSet bool
}

func (v NullableAppEventRelationships) Get() *AppEventRelationships {
	return v.value
}

func (v *NullableAppEventRelationships) Set(val *AppEventRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableAppEventRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableAppEventRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppEventRelationships(val *AppEventRelationships) *NullableAppEventRelationships {
	return &NullableAppEventRelationships{value: val, isSet: true}
}

func (v NullableAppEventRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppEventRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


