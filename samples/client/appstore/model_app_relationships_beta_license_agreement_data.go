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

// checks if the AppRelationshipsBetaLicenseAgreementData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationshipsBetaLicenseAgreementData{}

// AppRelationshipsBetaLicenseAgreementData struct for AppRelationshipsBetaLicenseAgreementData
type AppRelationshipsBetaLicenseAgreementData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppRelationshipsBetaLicenseAgreementData instantiates a new AppRelationshipsBetaLicenseAgreementData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsBetaLicenseAgreementData(type_ string, id string) *AppRelationshipsBetaLicenseAgreementData {
	this := AppRelationshipsBetaLicenseAgreementData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppRelationshipsBetaLicenseAgreementDataWithDefaults instantiates a new AppRelationshipsBetaLicenseAgreementData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsBetaLicenseAgreementDataWithDefaults() *AppRelationshipsBetaLicenseAgreementData {
	this := AppRelationshipsBetaLicenseAgreementData{}
	return &this
}

// GetType returns the Type field value
func (o *AppRelationshipsBetaLicenseAgreementData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsBetaLicenseAgreementData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppRelationshipsBetaLicenseAgreementData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppRelationshipsBetaLicenseAgreementData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsBetaLicenseAgreementData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppRelationshipsBetaLicenseAgreementData) SetId(v string) {
	o.Id = v
}

func (o AppRelationshipsBetaLicenseAgreementData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationshipsBetaLicenseAgreementData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppRelationshipsBetaLicenseAgreementData struct {
	value *AppRelationshipsBetaLicenseAgreementData
	isSet bool
}

func (v NullableAppRelationshipsBetaLicenseAgreementData) Get() *AppRelationshipsBetaLicenseAgreementData {
	return v.value
}

func (v *NullableAppRelationshipsBetaLicenseAgreementData) Set(val *AppRelationshipsBetaLicenseAgreementData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsBetaLicenseAgreementData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsBetaLicenseAgreementData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsBetaLicenseAgreementData(val *AppRelationshipsBetaLicenseAgreementData) *NullableAppRelationshipsBetaLicenseAgreementData {
	return &NullableAppRelationshipsBetaLicenseAgreementData{value: val, isSet: true}
}

func (v NullableAppRelationshipsBetaLicenseAgreementData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsBetaLicenseAgreementData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

