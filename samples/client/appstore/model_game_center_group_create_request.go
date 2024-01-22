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

// checks if the GameCenterGroupCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GameCenterGroupCreateRequest{}

// GameCenterGroupCreateRequest struct for GameCenterGroupCreateRequest
type GameCenterGroupCreateRequest struct {
	Data GameCenterGroupCreateRequestData `json:"data"`
}

// NewGameCenterGroupCreateRequest instantiates a new GameCenterGroupCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGameCenterGroupCreateRequest(data GameCenterGroupCreateRequestData) *GameCenterGroupCreateRequest {
	this := GameCenterGroupCreateRequest{}
	this.Data = data
	return &this
}

// NewGameCenterGroupCreateRequestWithDefaults instantiates a new GameCenterGroupCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGameCenterGroupCreateRequestWithDefaults() *GameCenterGroupCreateRequest {
	this := GameCenterGroupCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *GameCenterGroupCreateRequest) GetData() GameCenterGroupCreateRequestData {
	if o == nil {
		var ret GameCenterGroupCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GameCenterGroupCreateRequest) GetDataOk() (*GameCenterGroupCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GameCenterGroupCreateRequest) SetData(v GameCenterGroupCreateRequestData) {
	o.Data = v
}

func (o GameCenterGroupCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GameCenterGroupCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableGameCenterGroupCreateRequest struct {
	value *GameCenterGroupCreateRequest
	isSet bool
}

func (v NullableGameCenterGroupCreateRequest) Get() *GameCenterGroupCreateRequest {
	return v.value
}

func (v *NullableGameCenterGroupCreateRequest) Set(val *GameCenterGroupCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGameCenterGroupCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGameCenterGroupCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGameCenterGroupCreateRequest(val *GameCenterGroupCreateRequest) *NullableGameCenterGroupCreateRequest {
	return &NullableGameCenterGroupCreateRequest{value: val, isSet: true}
}

func (v NullableGameCenterGroupCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGameCenterGroupCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


