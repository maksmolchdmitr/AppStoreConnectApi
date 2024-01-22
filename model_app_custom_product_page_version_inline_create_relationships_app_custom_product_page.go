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

// checks if the AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage{}

// AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage struct for AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage
type AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage struct {
	Data *AppCustomProductPageVersionRelationshipsAppCustomProductPageData `json:"data,omitempty"`
}

// NewAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage instantiates a new AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage() *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage {
	this := AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage{}
	return &this
}

// NewAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageWithDefaults instantiates a new AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPageWithDefaults() *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage {
	this := AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) GetData() AppCustomProductPageVersionRelationshipsAppCustomProductPageData {
	if o == nil || IsNil(o.Data) {
		var ret AppCustomProductPageVersionRelationshipsAppCustomProductPageData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) GetDataOk() (*AppCustomProductPageVersionRelationshipsAppCustomProductPageData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given AppCustomProductPageVersionRelationshipsAppCustomProductPageData and assigns it to the Data field.
func (o *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) SetData(v AppCustomProductPageVersionRelationshipsAppCustomProductPageData) {
	o.Data = &v
}

func (o AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage struct {
	value *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage
	isSet bool
}

func (v NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) Get() *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage {
	return v.value
}

func (v *NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) Set(val *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage(val *AppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) *NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage {
	return &NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage{value: val, isSet: true}
}

func (v NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageVersionInlineCreateRelationshipsAppCustomProductPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
