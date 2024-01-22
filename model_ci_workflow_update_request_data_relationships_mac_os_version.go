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

// checks if the CiWorkflowUpdateRequestDataRelationshipsMacOsVersion type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiWorkflowUpdateRequestDataRelationshipsMacOsVersion{}

// CiWorkflowUpdateRequestDataRelationshipsMacOsVersion struct for CiWorkflowUpdateRequestDataRelationshipsMacOsVersion
type CiWorkflowUpdateRequestDataRelationshipsMacOsVersion struct {
	Data *CiWorkflowRelationshipsMacOsVersionData `json:"data,omitempty"`
}

// NewCiWorkflowUpdateRequestDataRelationshipsMacOsVersion instantiates a new CiWorkflowUpdateRequestDataRelationshipsMacOsVersion object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiWorkflowUpdateRequestDataRelationshipsMacOsVersion() *CiWorkflowUpdateRequestDataRelationshipsMacOsVersion {
	this := CiWorkflowUpdateRequestDataRelationshipsMacOsVersion{}
	return &this
}

// NewCiWorkflowUpdateRequestDataRelationshipsMacOsVersionWithDefaults instantiates a new CiWorkflowUpdateRequestDataRelationshipsMacOsVersion object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiWorkflowUpdateRequestDataRelationshipsMacOsVersionWithDefaults() *CiWorkflowUpdateRequestDataRelationshipsMacOsVersion {
	this := CiWorkflowUpdateRequestDataRelationshipsMacOsVersion{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CiWorkflowUpdateRequestDataRelationshipsMacOsVersion) GetData() CiWorkflowRelationshipsMacOsVersionData {
	if o == nil || IsNil(o.Data) {
		var ret CiWorkflowRelationshipsMacOsVersionData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiWorkflowUpdateRequestDataRelationshipsMacOsVersion) GetDataOk() (*CiWorkflowRelationshipsMacOsVersionData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CiWorkflowUpdateRequestDataRelationshipsMacOsVersion) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CiWorkflowRelationshipsMacOsVersionData and assigns it to the Data field.
func (o *CiWorkflowUpdateRequestDataRelationshipsMacOsVersion) SetData(v CiWorkflowRelationshipsMacOsVersionData) {
	o.Data = &v
}

func (o CiWorkflowUpdateRequestDataRelationshipsMacOsVersion) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiWorkflowUpdateRequestDataRelationshipsMacOsVersion) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCiWorkflowUpdateRequestDataRelationshipsMacOsVersion struct {
	value *CiWorkflowUpdateRequestDataRelationshipsMacOsVersion
	isSet bool
}

func (v NullableCiWorkflowUpdateRequestDataRelationshipsMacOsVersion) Get() *CiWorkflowUpdateRequestDataRelationshipsMacOsVersion {
	return v.value
}

func (v *NullableCiWorkflowUpdateRequestDataRelationshipsMacOsVersion) Set(val *CiWorkflowUpdateRequestDataRelationshipsMacOsVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableCiWorkflowUpdateRequestDataRelationshipsMacOsVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableCiWorkflowUpdateRequestDataRelationshipsMacOsVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiWorkflowUpdateRequestDataRelationshipsMacOsVersion(val *CiWorkflowUpdateRequestDataRelationshipsMacOsVersion) *NullableCiWorkflowUpdateRequestDataRelationshipsMacOsVersion {
	return &NullableCiWorkflowUpdateRequestDataRelationshipsMacOsVersion{value: val, isSet: true}
}

func (v NullableCiWorkflowUpdateRequestDataRelationshipsMacOsVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiWorkflowUpdateRequestDataRelationshipsMacOsVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
