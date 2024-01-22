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

// checks if the AppStoreVersionPhasedReleaseCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionPhasedReleaseCreateRequest{}

// AppStoreVersionPhasedReleaseCreateRequest struct for AppStoreVersionPhasedReleaseCreateRequest
type AppStoreVersionPhasedReleaseCreateRequest struct {
	Data AppStoreVersionPhasedReleaseCreateRequestData `json:"data"`
}

// NewAppStoreVersionPhasedReleaseCreateRequest instantiates a new AppStoreVersionPhasedReleaseCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionPhasedReleaseCreateRequest(data AppStoreVersionPhasedReleaseCreateRequestData) *AppStoreVersionPhasedReleaseCreateRequest {
	this := AppStoreVersionPhasedReleaseCreateRequest{}
	this.Data = data
	return &this
}

// NewAppStoreVersionPhasedReleaseCreateRequestWithDefaults instantiates a new AppStoreVersionPhasedReleaseCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionPhasedReleaseCreateRequestWithDefaults() *AppStoreVersionPhasedReleaseCreateRequest {
	this := AppStoreVersionPhasedReleaseCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionPhasedReleaseCreateRequest) GetData() AppStoreVersionPhasedReleaseCreateRequestData {
	if o == nil {
		var ret AppStoreVersionPhasedReleaseCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionPhasedReleaseCreateRequest) GetDataOk() (*AppStoreVersionPhasedReleaseCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionPhasedReleaseCreateRequest) SetData(v AppStoreVersionPhasedReleaseCreateRequestData) {
	o.Data = v
}

func (o AppStoreVersionPhasedReleaseCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionPhasedReleaseCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableAppStoreVersionPhasedReleaseCreateRequest struct {
	value *AppStoreVersionPhasedReleaseCreateRequest
	isSet bool
}

func (v NullableAppStoreVersionPhasedReleaseCreateRequest) Get() *AppStoreVersionPhasedReleaseCreateRequest {
	return v.value
}

func (v *NullableAppStoreVersionPhasedReleaseCreateRequest) Set(val *AppStoreVersionPhasedReleaseCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionPhasedReleaseCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionPhasedReleaseCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionPhasedReleaseCreateRequest(val *AppStoreVersionPhasedReleaseCreateRequest) *NullableAppStoreVersionPhasedReleaseCreateRequest {
	return &NullableAppStoreVersionPhasedReleaseCreateRequest{value: val, isSet: true}
}

func (v NullableAppStoreVersionPhasedReleaseCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionPhasedReleaseCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


