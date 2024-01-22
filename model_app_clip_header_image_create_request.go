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

// checks if the AppClipHeaderImageCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppClipHeaderImageCreateRequest{}

// AppClipHeaderImageCreateRequest struct for AppClipHeaderImageCreateRequest
type AppClipHeaderImageCreateRequest struct {
	Data AppClipHeaderImageCreateRequestData `json:"data"`
}

// NewAppClipHeaderImageCreateRequest instantiates a new AppClipHeaderImageCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppClipHeaderImageCreateRequest(data AppClipHeaderImageCreateRequestData) *AppClipHeaderImageCreateRequest {
	this := AppClipHeaderImageCreateRequest{}
	this.Data = data
	return &this
}

// NewAppClipHeaderImageCreateRequestWithDefaults instantiates a new AppClipHeaderImageCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppClipHeaderImageCreateRequestWithDefaults() *AppClipHeaderImageCreateRequest {
	this := AppClipHeaderImageCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppClipHeaderImageCreateRequest) GetData() AppClipHeaderImageCreateRequestData {
	if o == nil {
		var ret AppClipHeaderImageCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppClipHeaderImageCreateRequest) GetDataOk() (*AppClipHeaderImageCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppClipHeaderImageCreateRequest) SetData(v AppClipHeaderImageCreateRequestData) {
	o.Data = v
}

func (o AppClipHeaderImageCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppClipHeaderImageCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppClipHeaderImageCreateRequest struct {
	value *AppClipHeaderImageCreateRequest
	isSet bool
}

func (v NullableAppClipHeaderImageCreateRequest) Get() *AppClipHeaderImageCreateRequest {
	return v.value
}

func (v *NullableAppClipHeaderImageCreateRequest) Set(val *AppClipHeaderImageCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppClipHeaderImageCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppClipHeaderImageCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppClipHeaderImageCreateRequest(val *AppClipHeaderImageCreateRequest) *NullableAppClipHeaderImageCreateRequest {
	return &NullableAppClipHeaderImageCreateRequest{value: val, isSet: true}
}

func (v NullableAppClipHeaderImageCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppClipHeaderImageCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
