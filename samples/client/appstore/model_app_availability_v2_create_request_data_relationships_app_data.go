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

// checks if the AppAvailabilityV2CreateRequestDataRelationshipsAppData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAvailabilityV2CreateRequestDataRelationshipsAppData{}

// AppAvailabilityV2CreateRequestDataRelationshipsAppData struct for AppAvailabilityV2CreateRequestDataRelationshipsAppData
type AppAvailabilityV2CreateRequestDataRelationshipsAppData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppAvailabilityV2CreateRequestDataRelationshipsAppData instantiates a new AppAvailabilityV2CreateRequestDataRelationshipsAppData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAvailabilityV2CreateRequestDataRelationshipsAppData(type_ string, id string) *AppAvailabilityV2CreateRequestDataRelationshipsAppData {
	this := AppAvailabilityV2CreateRequestDataRelationshipsAppData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppAvailabilityV2CreateRequestDataRelationshipsAppDataWithDefaults instantiates a new AppAvailabilityV2CreateRequestDataRelationshipsAppData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAvailabilityV2CreateRequestDataRelationshipsAppDataWithDefaults() *AppAvailabilityV2CreateRequestDataRelationshipsAppData {
	this := AppAvailabilityV2CreateRequestDataRelationshipsAppData{}
	return &this
}

// GetType returns the Type field value
func (o *AppAvailabilityV2CreateRequestDataRelationshipsAppData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2CreateRequestDataRelationshipsAppData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppAvailabilityV2CreateRequestDataRelationshipsAppData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppAvailabilityV2CreateRequestDataRelationshipsAppData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityV2CreateRequestDataRelationshipsAppData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppAvailabilityV2CreateRequestDataRelationshipsAppData) SetId(v string) {
	o.Id = v
}

func (o AppAvailabilityV2CreateRequestDataRelationshipsAppData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAvailabilityV2CreateRequestDataRelationshipsAppData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppAvailabilityV2CreateRequestDataRelationshipsAppData struct {
	value *AppAvailabilityV2CreateRequestDataRelationshipsAppData
	isSet bool
}

func (v NullableAppAvailabilityV2CreateRequestDataRelationshipsAppData) Get() *AppAvailabilityV2CreateRequestDataRelationshipsAppData {
	return v.value
}

func (v *NullableAppAvailabilityV2CreateRequestDataRelationshipsAppData) Set(val *AppAvailabilityV2CreateRequestDataRelationshipsAppData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAvailabilityV2CreateRequestDataRelationshipsAppData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAvailabilityV2CreateRequestDataRelationshipsAppData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAvailabilityV2CreateRequestDataRelationshipsAppData(val *AppAvailabilityV2CreateRequestDataRelationshipsAppData) *NullableAppAvailabilityV2CreateRequestDataRelationshipsAppData {
	return &NullableAppAvailabilityV2CreateRequestDataRelationshipsAppData{value: val, isSet: true}
}

func (v NullableAppAvailabilityV2CreateRequestDataRelationshipsAppData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAvailabilityV2CreateRequestDataRelationshipsAppData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


