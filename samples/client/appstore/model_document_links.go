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

// checks if the DocumentLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DocumentLinks{}

// DocumentLinks struct for DocumentLinks
type DocumentLinks struct {
	Self string `json:"self"`
}

// NewDocumentLinks instantiates a new DocumentLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDocumentLinks(self string) *DocumentLinks {
	this := DocumentLinks{}
	this.Self = self
	return &this
}

// NewDocumentLinksWithDefaults instantiates a new DocumentLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDocumentLinksWithDefaults() *DocumentLinks {
	this := DocumentLinks{}
	return &this
}

// GetSelf returns the Self field value
func (o *DocumentLinks) GetSelf() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Self
}

// GetSelfOk returns a tuple with the Self field value
// and a boolean to check if the value has been set.
func (o *DocumentLinks) GetSelfOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Self, true
}

// SetSelf sets field value
func (o *DocumentLinks) SetSelf(v string) {
	o.Self = v
}

func (o DocumentLinks) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DocumentLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["self"] = o.Self
	return toSerialize, nil
}

type NullableDocumentLinks struct {
	value *DocumentLinks
	isSet bool
}

func (v NullableDocumentLinks) Get() *DocumentLinks {
	return v.value
}

func (v *NullableDocumentLinks) Set(val *DocumentLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableDocumentLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableDocumentLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDocumentLinks(val *DocumentLinks) *NullableDocumentLinks {
	return &NullableDocumentLinks{value: val, isSet: true}
}

func (v NullableDocumentLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDocumentLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


