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

// checks if the GameCenterAppVersionUpdateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterAppVersionUpdateRequest{}

// GameCenterAppVersionUpdateRequest struct for GameCenterAppVersionUpdateRequest
type GameCenterAppVersionUpdateRequest struct {
	Data GameCenterAppVersionUpdateRequestData `json:"data"`
}

// NewGameCenterAppVersionUpdateRequest instantiates a new GameCenterAppVersionUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterAppVersionUpdateRequest(data GameCenterAppVersionUpdateRequestData) *GameCenterAppVersionUpdateRequest {
	this := GameCenterAppVersionUpdateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterAppVersionUpdateRequestWithDefaults instantiates a new GameCenterAppVersionUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterAppVersionUpdateRequestWithDefaults() *GameCenterAppVersionUpdateRequest {
	this := GameCenterAppVersionUpdateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterAppVersionUpdateRequest) GetData() GameCenterAppVersionUpdateRequestData {
	if o == nil {
		var ret GameCenterAppVersionUpdateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterAppVersionUpdateRequest) GetDataOk() (*GameCenterAppVersionUpdateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterAppVersionUpdateRequest) SetData(v GameCenterAppVersionUpdateRequestData) {
	o.Data = v
}

func (o GameCenterAppVersionUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterAppVersionUpdateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGameCenterAppVersionUpdateRequest struct {
	value *GameCenterAppVersionUpdateRequest
	isSet bool
}

func (v NullableGameCenterAppVersionUpdateRequest) Get() *GameCenterAppVersionUpdateRequest {
	return v.value
}

func (v *NullableGameCenterAppVersionUpdateRequest) Set(val *GameCenterAppVersionUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterAppVersionUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterAppVersionUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterAppVersionUpdateRequest(val *GameCenterAppVersionUpdateRequest) *NullableGameCenterAppVersionUpdateRequest {
	return &NullableGameCenterAppVersionUpdateRequest{value: val, isSet: true}
}

func (v NullableGameCenterAppVersionUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterAppVersionUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

