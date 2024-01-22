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

// checks if the AppPriceScheduleCreateRequestDataRelationshipsManualPrices type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPriceScheduleCreateRequestDataRelationshipsManualPrices{}

// AppPriceScheduleCreateRequestDataRelationshipsManualPrices struct for AppPriceScheduleCreateRequestDataRelationshipsManualPrices
type AppPriceScheduleCreateRequestDataRelationshipsManualPrices struct {
	Data []AppPriceScheduleRelationshipsManualPricesDataInner `json:"data"`
}

// NewAppPriceScheduleCreateRequestDataRelationshipsManualPrices instantiates a new AppPriceScheduleCreateRequestDataRelationshipsManualPrices object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPriceScheduleCreateRequestDataRelationshipsManualPrices(data []AppPriceScheduleRelationshipsManualPricesDataInner) *AppPriceScheduleCreateRequestDataRelationshipsManualPrices {
	this := AppPriceScheduleCreateRequestDataRelationshipsManualPrices{}
	this.Data = data
	return &this
}

// NewAppPriceScheduleCreateRequestDataRelationshipsManualPricesWithDefaults instantiates a new AppPriceScheduleCreateRequestDataRelationshipsManualPrices object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPriceScheduleCreateRequestDataRelationshipsManualPricesWithDefaults() *AppPriceScheduleCreateRequestDataRelationshipsManualPrices {
	this := AppPriceScheduleCreateRequestDataRelationshipsManualPrices{}
	return &this
}

// GetData returns the Data field value
func (o *AppPriceScheduleCreateRequestDataRelationshipsManualPrices) GetData() []AppPriceScheduleRelationshipsManualPricesDataInner {
	if o == nil {
		var ret []AppPriceScheduleRelationshipsManualPricesDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppPriceScheduleCreateRequestDataRelationshipsManualPrices) GetDataOk() ([]AppPriceScheduleRelationshipsManualPricesDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppPriceScheduleCreateRequestDataRelationshipsManualPrices) SetData(v []AppPriceScheduleRelationshipsManualPricesDataInner) {
	o.Data = v
}

func (o AppPriceScheduleCreateRequestDataRelationshipsManualPrices) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPriceScheduleCreateRequestDataRelationshipsManualPrices) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppPriceScheduleCreateRequestDataRelationshipsManualPrices struct {
	value *AppPriceScheduleCreateRequestDataRelationshipsManualPrices
	isSet bool
}

func (v NullableAppPriceScheduleCreateRequestDataRelationshipsManualPrices) Get() *AppPriceScheduleCreateRequestDataRelationshipsManualPrices {
	return v.value
}

func (v *NullableAppPriceScheduleCreateRequestDataRelationshipsManualPrices) Set(val *AppPriceScheduleCreateRequestDataRelationshipsManualPrices) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPriceScheduleCreateRequestDataRelationshipsManualPrices) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPriceScheduleCreateRequestDataRelationshipsManualPrices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPriceScheduleCreateRequestDataRelationshipsManualPrices(val *AppPriceScheduleCreateRequestDataRelationshipsManualPrices) *NullableAppPriceScheduleCreateRequestDataRelationshipsManualPrices {
	return &NullableAppPriceScheduleCreateRequestDataRelationshipsManualPrices{value: val, isSet: true}
}

func (v NullableAppPriceScheduleCreateRequestDataRelationshipsManualPrices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPriceScheduleCreateRequestDataRelationshipsManualPrices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
