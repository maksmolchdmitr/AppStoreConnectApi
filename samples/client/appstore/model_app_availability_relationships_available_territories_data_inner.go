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

// checks if the AppAvailabilityRelationshipsAvailableTerritoriesDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppAvailabilityRelationshipsAvailableTerritoriesDataInner{}

// AppAvailabilityRelationshipsAvailableTerritoriesDataInner struct for AppAvailabilityRelationshipsAvailableTerritoriesDataInner
type AppAvailabilityRelationshipsAvailableTerritoriesDataInner struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppAvailabilityRelationshipsAvailableTerritoriesDataInner instantiates a new AppAvailabilityRelationshipsAvailableTerritoriesDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppAvailabilityRelationshipsAvailableTerritoriesDataInner(type_ string, id string) *AppAvailabilityRelationshipsAvailableTerritoriesDataInner {
	this := AppAvailabilityRelationshipsAvailableTerritoriesDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppAvailabilityRelationshipsAvailableTerritoriesDataInnerWithDefaults instantiates a new AppAvailabilityRelationshipsAvailableTerritoriesDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppAvailabilityRelationshipsAvailableTerritoriesDataInnerWithDefaults() *AppAvailabilityRelationshipsAvailableTerritoriesDataInner {
	this := AppAvailabilityRelationshipsAvailableTerritoriesDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *AppAvailabilityRelationshipsAvailableTerritoriesDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityRelationshipsAvailableTerritoriesDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppAvailabilityRelationshipsAvailableTerritoriesDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppAvailabilityRelationshipsAvailableTerritoriesDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppAvailabilityRelationshipsAvailableTerritoriesDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppAvailabilityRelationshipsAvailableTerritoriesDataInner) SetId(v string) {
	o.Id = v
}

func (o AppAvailabilityRelationshipsAvailableTerritoriesDataInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppAvailabilityRelationshipsAvailableTerritoriesDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppAvailabilityRelationshipsAvailableTerritoriesDataInner struct {
	value *AppAvailabilityRelationshipsAvailableTerritoriesDataInner
	isSet bool
}

func (v NullableAppAvailabilityRelationshipsAvailableTerritoriesDataInner) Get() *AppAvailabilityRelationshipsAvailableTerritoriesDataInner {
	return v.value
}

func (v *NullableAppAvailabilityRelationshipsAvailableTerritoriesDataInner) Set(val *AppAvailabilityRelationshipsAvailableTerritoriesDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppAvailabilityRelationshipsAvailableTerritoriesDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppAvailabilityRelationshipsAvailableTerritoriesDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppAvailabilityRelationshipsAvailableTerritoriesDataInner(val *AppAvailabilityRelationshipsAvailableTerritoriesDataInner) *NullableAppAvailabilityRelationshipsAvailableTerritoriesDataInner {
	return &NullableAppAvailabilityRelationshipsAvailableTerritoriesDataInner{value: val, isSet: true}
}

func (v NullableAppAvailabilityRelationshipsAvailableTerritoriesDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppAvailabilityRelationshipsAvailableTerritoriesDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


