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

// checks if the AppCustomProductPageUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageUpdateRequest{}

// AppCustomProductPageUpdateRequest struct for AppCustomProductPageUpdateRequest
type AppCustomProductPageUpdateRequest struct {
	Data AppCustomProductPageUpdateRequestData `json:"data"`
}

// NewAppCustomProductPageUpdateRequest instantiates a new AppCustomProductPageUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageUpdateRequest(data AppCustomProductPageUpdateRequestData) *AppCustomProductPageUpdateRequest {
	this := AppCustomProductPageUpdateRequest{}
	this.Data = data
	return &this
}

// NewAppCustomProductPageUpdateRequestWithDefaults instantiates a new AppCustomProductPageUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageUpdateRequestWithDefaults() *AppCustomProductPageUpdateRequest {
	this := AppCustomProductPageUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppCustomProductPageUpdateRequest) GetData() AppCustomProductPageUpdateRequestData {
	if o == nil {
		var ret AppCustomProductPageUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageUpdateRequest) GetDataOk() (*AppCustomProductPageUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppCustomProductPageUpdateRequest) SetData(v AppCustomProductPageUpdateRequestData) {
	o.Data = v
}

func (o AppCustomProductPageUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppCustomProductPageUpdateRequest struct {
	value *AppCustomProductPageUpdateRequest
	isSet bool
}

func (v NullableAppCustomProductPageUpdateRequest) Get() *AppCustomProductPageUpdateRequest {
	return v.value
}

func (v *NullableAppCustomProductPageUpdateRequest) Set(val *AppCustomProductPageUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageUpdateRequest(val *AppCustomProductPageUpdateRequest) *NullableAppCustomProductPageUpdateRequest {
	return &NullableAppCustomProductPageUpdateRequest{value: val, isSet: true}
}

func (v NullableAppCustomProductPageUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


