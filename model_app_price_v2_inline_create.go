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

// checks if the AppPriceV2InlineCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPriceV2InlineCreate{}

// AppPriceV2InlineCreate struct for AppPriceV2InlineCreate
type AppPriceV2InlineCreate struct {
	Type string  `json:"type"`
	Id   *string `json:"id,omitempty"`
}

// NewAppPriceV2InlineCreate instantiates a new AppPriceV2InlineCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPriceV2InlineCreate(type_ string) *AppPriceV2InlineCreate {
	this := AppPriceV2InlineCreate{}
	this.Type = type_
	return &this
}

// NewAppPriceV2InlineCreateWithDefaults instantiates a new AppPriceV2InlineCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPriceV2InlineCreateWithDefaults() *AppPriceV2InlineCreate {
	this := AppPriceV2InlineCreate{}
	return &this
}

// GetType returns the Type field value
func (o *AppPriceV2InlineCreate) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppPriceV2InlineCreate) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppPriceV2InlineCreate) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AppPriceV2InlineCreate) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPriceV2InlineCreate) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AppPriceV2InlineCreate) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AppPriceV2InlineCreate) SetId(v string) {
	o.Id = &v
}

func (o AppPriceV2InlineCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPriceV2InlineCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableAppPriceV2InlineCreate struct {
	value *AppPriceV2InlineCreate
	isSet bool
}

func (v NullableAppPriceV2InlineCreate) Get() *AppPriceV2InlineCreate {
	return v.value
}

func (v *NullableAppPriceV2InlineCreate) Set(val *AppPriceV2InlineCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPriceV2InlineCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPriceV2InlineCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPriceV2InlineCreate(val *AppPriceV2InlineCreate) *NullableAppPriceV2InlineCreate {
	return &NullableAppPriceV2InlineCreate{value: val, isSet: true}
}

func (v NullableAppPriceV2InlineCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPriceV2InlineCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
