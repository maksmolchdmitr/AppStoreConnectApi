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

// checks if the ResourceLinks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ResourceLinks{}

// ResourceLinks struct for ResourceLinks
type ResourceLinks struct {
	Self *string `json:"self,omitempty"`
}

// NewResourceLinks instantiates a new ResourceLinks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewResourceLinks() *ResourceLinks {
	this := ResourceLinks{}
	return &this
}

// NewResourceLinksWithDefaults instantiates a new ResourceLinks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceLinksWithDefaults() *ResourceLinks {
	this := ResourceLinks{}
	return &this
}

// GetSelf returns the Self field value if set, zero value otherwise.
func (o *ResourceLinks) GetSelf() string {
	if o == nil || IsNil(o.Self) {
		var ret string
		return ret
	}
	return *o.Self
}

// GetSelfOk returns a tuple with the Self field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceLinks) GetSelfOk() (*string, bool) {
	if o == nil || IsNil(o.Self) {
		return nil, false
	}
	return o.Self, true
}

// HasSelf returns a boolean if a field has been set.
func (o *ResourceLinks) HasSelf() bool {
	if o != nil && !IsNil(o.Self) {
		return true
	}

	return false
}

// SetSelf gets a reference to the given string and assigns it to the Self field.
func (o *ResourceLinks) SetSelf(v string) {
	o.Self = &v
}

func (o ResourceLinks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceLinks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Self) {
		toSerialize["self"] = o.Self
	}
	return toSerialize, nil
}

type NullableResourceLinks struct {
	value *ResourceLinks
	isSet bool
}

func (v NullableResourceLinks) Get() *ResourceLinks {
	return v.value
}

func (v *NullableResourceLinks) Set(val *ResourceLinks) {
	v.value = val
	v.isSet = true
}

func (v NullableResourceLinks) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceLinks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceLinks(val *ResourceLinks) *NullableResourceLinks {
	return &NullableResourceLinks{value: val, isSet: true}
}

func (v NullableResourceLinks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceLinks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}