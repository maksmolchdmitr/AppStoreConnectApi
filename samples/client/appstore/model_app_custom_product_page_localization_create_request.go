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

// checks if the AppCustomProductPageLocalizationCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageLocalizationCreateRequest{}

// AppCustomProductPageLocalizationCreateRequest struct for AppCustomProductPageLocalizationCreateRequest
type AppCustomProductPageLocalizationCreateRequest struct {
	Data AppCustomProductPageLocalizationCreateRequestData `json:"data"`
}

// NewAppCustomProductPageLocalizationCreateRequest instantiates a new AppCustomProductPageLocalizationCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageLocalizationCreateRequest(data AppCustomProductPageLocalizationCreateRequestData) *AppCustomProductPageLocalizationCreateRequest {
	this := AppCustomProductPageLocalizationCreateRequest{}
	this.Data = data
	return &this
}

// NewAppCustomProductPageLocalizationCreateRequestWithDefaults instantiates a new AppCustomProductPageLocalizationCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageLocalizationCreateRequestWithDefaults() *AppCustomProductPageLocalizationCreateRequest {
	this := AppCustomProductPageLocalizationCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppCustomProductPageLocalizationCreateRequest) GetData() AppCustomProductPageLocalizationCreateRequestData {
	if o == nil {
		var ret AppCustomProductPageLocalizationCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationCreateRequest) GetDataOk() (*AppCustomProductPageLocalizationCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppCustomProductPageLocalizationCreateRequest) SetData(v AppCustomProductPageLocalizationCreateRequestData) {
	o.Data = v
}

func (o AppCustomProductPageLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageLocalizationCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppCustomProductPageLocalizationCreateRequest struct {
	value *AppCustomProductPageLocalizationCreateRequest
	isSet bool
}

func (v NullableAppCustomProductPageLocalizationCreateRequest) Get() *AppCustomProductPageLocalizationCreateRequest {
	return v.value
}

func (v *NullableAppCustomProductPageLocalizationCreateRequest) Set(val *AppCustomProductPageLocalizationCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageLocalizationCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageLocalizationCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageLocalizationCreateRequest(val *AppCustomProductPageLocalizationCreateRequest) *NullableAppCustomProductPageLocalizationCreateRequest {
	return &NullableAppCustomProductPageLocalizationCreateRequest{value: val, isSet: true}
}

func (v NullableAppCustomProductPageLocalizationCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageLocalizationCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


