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

// checks if the CiMacOsVersionRelationshipsXcodeVersionsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiMacOsVersionRelationshipsXcodeVersionsDataInner{}

// CiMacOsVersionRelationshipsXcodeVersionsDataInner struct for CiMacOsVersionRelationshipsXcodeVersionsDataInner
type CiMacOsVersionRelationshipsXcodeVersionsDataInner struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewCiMacOsVersionRelationshipsXcodeVersionsDataInner instantiates a new CiMacOsVersionRelationshipsXcodeVersionsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiMacOsVersionRelationshipsXcodeVersionsDataInner(type_ string, id string) *CiMacOsVersionRelationshipsXcodeVersionsDataInner {
	this := CiMacOsVersionRelationshipsXcodeVersionsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewCiMacOsVersionRelationshipsXcodeVersionsDataInnerWithDefaults instantiates a new CiMacOsVersionRelationshipsXcodeVersionsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiMacOsVersionRelationshipsXcodeVersionsDataInnerWithDefaults() *CiMacOsVersionRelationshipsXcodeVersionsDataInner {
	this := CiMacOsVersionRelationshipsXcodeVersionsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *CiMacOsVersionRelationshipsXcodeVersionsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CiMacOsVersionRelationshipsXcodeVersionsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CiMacOsVersionRelationshipsXcodeVersionsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *CiMacOsVersionRelationshipsXcodeVersionsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *CiMacOsVersionRelationshipsXcodeVersionsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *CiMacOsVersionRelationshipsXcodeVersionsDataInner) SetId(v string) {
	o.Id = v
}

func (o CiMacOsVersionRelationshipsXcodeVersionsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiMacOsVersionRelationshipsXcodeVersionsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableCiMacOsVersionRelationshipsXcodeVersionsDataInner struct {
	value *CiMacOsVersionRelationshipsXcodeVersionsDataInner
	isSet bool
}

func (v NullableCiMacOsVersionRelationshipsXcodeVersionsDataInner) Get() *CiMacOsVersionRelationshipsXcodeVersionsDataInner {
	return v.value
}

func (v *NullableCiMacOsVersionRelationshipsXcodeVersionsDataInner) Set(val *CiMacOsVersionRelationshipsXcodeVersionsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableCiMacOsVersionRelationshipsXcodeVersionsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableCiMacOsVersionRelationshipsXcodeVersionsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiMacOsVersionRelationshipsXcodeVersionsDataInner(val *CiMacOsVersionRelationshipsXcodeVersionsDataInner) *NullableCiMacOsVersionRelationshipsXcodeVersionsDataInner {
	return &NullableCiMacOsVersionRelationshipsXcodeVersionsDataInner{value: val, isSet: true}
}

func (v NullableCiMacOsVersionRelationshipsXcodeVersionsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiMacOsVersionRelationshipsXcodeVersionsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
