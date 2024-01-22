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

// checks if the AppCustomProductPageVersionCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageVersionCreateRequest{}

// AppCustomProductPageVersionCreateRequest struct for AppCustomProductPageVersionCreateRequest
type AppCustomProductPageVersionCreateRequest struct {
	Data AppCustomProductPageVersionCreateRequestData `json:"data"`
}

// NewAppCustomProductPageVersionCreateRequest instantiates a new AppCustomProductPageVersionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageVersionCreateRequest(data AppCustomProductPageVersionCreateRequestData) *AppCustomProductPageVersionCreateRequest {
	this := AppCustomProductPageVersionCreateRequest{}
	this.Data = data
	return &this
}

// NewAppCustomProductPageVersionCreateRequestWithDefaults instantiates a new AppCustomProductPageVersionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageVersionCreateRequestWithDefaults() *AppCustomProductPageVersionCreateRequest {
	this := AppCustomProductPageVersionCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppCustomProductPageVersionCreateRequest) GetData() AppCustomProductPageVersionCreateRequestData {
	if o == nil {
		var ret AppCustomProductPageVersionCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionCreateRequest) GetDataOk() (*AppCustomProductPageVersionCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppCustomProductPageVersionCreateRequest) SetData(v AppCustomProductPageVersionCreateRequestData) {
	o.Data = v
}

func (o AppCustomProductPageVersionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageVersionCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppCustomProductPageVersionCreateRequest struct {
	value *AppCustomProductPageVersionCreateRequest
	isSet bool
}

func (v NullableAppCustomProductPageVersionCreateRequest) Get() *AppCustomProductPageVersionCreateRequest {
	return v.value
}

func (v *NullableAppCustomProductPageVersionCreateRequest) Set(val *AppCustomProductPageVersionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageVersionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageVersionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageVersionCreateRequest(val *AppCustomProductPageVersionCreateRequest) *NullableAppCustomProductPageVersionCreateRequest {
	return &NullableAppCustomProductPageVersionCreateRequest{value: val, isSet: true}
}

func (v NullableAppCustomProductPageVersionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageVersionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


