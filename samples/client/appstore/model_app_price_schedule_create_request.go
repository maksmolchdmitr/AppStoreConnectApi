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

// checks if the AppPriceScheduleCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPriceScheduleCreateRequest{}

// AppPriceScheduleCreateRequest struct for AppPriceScheduleCreateRequest
type AppPriceScheduleCreateRequest struct {
	Data AppPriceScheduleCreateRequestData `json:"data"`
	Included []AppPriceScheduleCreateRequestIncludedInner `json:"included,omitempty"`
}

// NewAppPriceScheduleCreateRequest instantiates a new AppPriceScheduleCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPriceScheduleCreateRequest(data AppPriceScheduleCreateRequestData) *AppPriceScheduleCreateRequest {
	this := AppPriceScheduleCreateRequest{}
	this.Data = data
	return &this
}

// NewAppPriceScheduleCreateRequestWithDefaults instantiates a new AppPriceScheduleCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPriceScheduleCreateRequestWithDefaults() *AppPriceScheduleCreateRequest {
	this := AppPriceScheduleCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppPriceScheduleCreateRequest) GetData() AppPriceScheduleCreateRequestData {
	if o == nil {
		var ret AppPriceScheduleCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppPriceScheduleCreateRequest) GetDataOk() (*AppPriceScheduleCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppPriceScheduleCreateRequest) SetData(v AppPriceScheduleCreateRequestData) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppPriceScheduleCreateRequest) GetIncluded() []AppPriceScheduleCreateRequestIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []AppPriceScheduleCreateRequestIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPriceScheduleCreateRequest) GetIncludedOk() ([]AppPriceScheduleCreateRequestIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppPriceScheduleCreateRequest) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppPriceScheduleCreateRequestIncludedInner and assigns it to the Included field.
func (o *AppPriceScheduleCreateRequest) SetIncluded(v []AppPriceScheduleCreateRequestIncludedInner) {
	o.Included = v
}

func (o AppPriceScheduleCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPriceScheduleCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	return toSerialize, nil
}

type NullableAppPriceScheduleCreateRequest struct {
	value *AppPriceScheduleCreateRequest
	isSet bool
}

func (v NullableAppPriceScheduleCreateRequest) Get() *AppPriceScheduleCreateRequest {
	return v.value
}

func (v *NullableAppPriceScheduleCreateRequest) Set(val *AppPriceScheduleCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPriceScheduleCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPriceScheduleCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPriceScheduleCreateRequest(val *AppPriceScheduleCreateRequest) *NullableAppPriceScheduleCreateRequest {
	return &NullableAppPriceScheduleCreateRequest{value: val, isSet: true}
}

func (v NullableAppPriceScheduleCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPriceScheduleCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


