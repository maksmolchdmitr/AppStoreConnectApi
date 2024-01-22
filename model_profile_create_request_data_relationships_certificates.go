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

// checks if the ProfileCreateRequestDataRelationshipsCertificates type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProfileCreateRequestDataRelationshipsCertificates{}

// ProfileCreateRequestDataRelationshipsCertificates struct for ProfileCreateRequestDataRelationshipsCertificates
type ProfileCreateRequestDataRelationshipsCertificates struct {
	Data []ProfileRelationshipsCertificatesDataInner `json:"data"`
}

// NewProfileCreateRequestDataRelationshipsCertificates instantiates a new ProfileCreateRequestDataRelationshipsCertificates object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProfileCreateRequestDataRelationshipsCertificates(data []ProfileRelationshipsCertificatesDataInner) *ProfileCreateRequestDataRelationshipsCertificates {
	this := ProfileCreateRequestDataRelationshipsCertificates{}
	this.Data = data
	return &this
}

// NewProfileCreateRequestDataRelationshipsCertificatesWithDefaults instantiates a new ProfileCreateRequestDataRelationshipsCertificates object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProfileCreateRequestDataRelationshipsCertificatesWithDefaults() *ProfileCreateRequestDataRelationshipsCertificates {
	this := ProfileCreateRequestDataRelationshipsCertificates{}
	return &this
}

// GetData returns the Data field value
func (o *ProfileCreateRequestDataRelationshipsCertificates) GetData() []ProfileRelationshipsCertificatesDataInner {
	if o == nil {
		var ret []ProfileRelationshipsCertificatesDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ProfileCreateRequestDataRelationshipsCertificates) GetDataOk() ([]ProfileRelationshipsCertificatesDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ProfileCreateRequestDataRelationshipsCertificates) SetData(v []ProfileRelationshipsCertificatesDataInner) {
	o.Data = v
}

func (o ProfileCreateRequestDataRelationshipsCertificates) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProfileCreateRequestDataRelationshipsCertificates) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableProfileCreateRequestDataRelationshipsCertificates struct {
	value *ProfileCreateRequestDataRelationshipsCertificates
	isSet bool
}

func (v NullableProfileCreateRequestDataRelationshipsCertificates) Get() *ProfileCreateRequestDataRelationshipsCertificates {
	return v.value
}

func (v *NullableProfileCreateRequestDataRelationshipsCertificates) Set(val *ProfileCreateRequestDataRelationshipsCertificates) {
	v.value = val
	v.isSet = true
}

func (v NullableProfileCreateRequestDataRelationshipsCertificates) IsSet() bool {
	return v.isSet
}

func (v *NullableProfileCreateRequestDataRelationshipsCertificates) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProfileCreateRequestDataRelationshipsCertificates(val *ProfileCreateRequestDataRelationshipsCertificates) *NullableProfileCreateRequestDataRelationshipsCertificates {
	return &NullableProfileCreateRequestDataRelationshipsCertificates{value: val, isSet: true}
}

func (v NullableProfileCreateRequestDataRelationshipsCertificates) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProfileCreateRequestDataRelationshipsCertificates) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
