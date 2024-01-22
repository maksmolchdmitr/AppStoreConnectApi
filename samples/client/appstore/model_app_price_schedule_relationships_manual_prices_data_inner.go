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

// checks if the AppPriceScheduleRelationshipsManualPricesDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPriceScheduleRelationshipsManualPricesDataInner{}

// AppPriceScheduleRelationshipsManualPricesDataInner struct for AppPriceScheduleRelationshipsManualPricesDataInner
type AppPriceScheduleRelationshipsManualPricesDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppPriceScheduleRelationshipsManualPricesDataInner instantiates a new AppPriceScheduleRelationshipsManualPricesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPriceScheduleRelationshipsManualPricesDataInner(type_ string, id string) *AppPriceScheduleRelationshipsManualPricesDataInner {
	this := AppPriceScheduleRelationshipsManualPricesDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppPriceScheduleRelationshipsManualPricesDataInnerWithDefaults instantiates a new AppPriceScheduleRelationshipsManualPricesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPriceScheduleRelationshipsManualPricesDataInnerWithDefaults() *AppPriceScheduleRelationshipsManualPricesDataInner {
	this := AppPriceScheduleRelationshipsManualPricesDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *AppPriceScheduleRelationshipsManualPricesDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppPriceScheduleRelationshipsManualPricesDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppPriceScheduleRelationshipsManualPricesDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppPriceScheduleRelationshipsManualPricesDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppPriceScheduleRelationshipsManualPricesDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppPriceScheduleRelationshipsManualPricesDataInner) SetId(v string) {
	o.Id = v
}

func (o AppPriceScheduleRelationshipsManualPricesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPriceScheduleRelationshipsManualPricesDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppPriceScheduleRelationshipsManualPricesDataInner struct {
	value *AppPriceScheduleRelationshipsManualPricesDataInner
	isSet bool
}

func (v NullableAppPriceScheduleRelationshipsManualPricesDataInner) Get() *AppPriceScheduleRelationshipsManualPricesDataInner {
	return v.value
}

func (v *NullableAppPriceScheduleRelationshipsManualPricesDataInner) Set(val *AppPriceScheduleRelationshipsManualPricesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPriceScheduleRelationshipsManualPricesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPriceScheduleRelationshipsManualPricesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPriceScheduleRelationshipsManualPricesDataInner(val *AppPriceScheduleRelationshipsManualPricesDataInner) *NullableAppPriceScheduleRelationshipsManualPricesDataInner {
	return &NullableAppPriceScheduleRelationshipsManualPricesDataInner{value: val, isSet: true}
}

func (v NullableAppPriceScheduleRelationshipsManualPricesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPriceScheduleRelationshipsManualPricesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


