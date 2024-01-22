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

// checks if the CiBuildRunCreateRequestDataRelationshipsBuildRun type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiBuildRunCreateRequestDataRelationshipsBuildRun{}

// CiBuildRunCreateRequestDataRelationshipsBuildRun struct for CiBuildRunCreateRequestDataRelationshipsBuildRun
type CiBuildRunCreateRequestDataRelationshipsBuildRun struct {
	Data *CiBuildActionRelationshipsBuildRunData `json:"data,omitempty"`
}

// NewCiBuildRunCreateRequestDataRelationshipsBuildRun instantiates a new CiBuildRunCreateRequestDataRelationshipsBuildRun object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiBuildRunCreateRequestDataRelationshipsBuildRun() *CiBuildRunCreateRequestDataRelationshipsBuildRun {
	this := CiBuildRunCreateRequestDataRelationshipsBuildRun{}
	return &this
}

// NewCiBuildRunCreateRequestDataRelationshipsBuildRunWithDefaults instantiates a new CiBuildRunCreateRequestDataRelationshipsBuildRun object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiBuildRunCreateRequestDataRelationshipsBuildRunWithDefaults() *CiBuildRunCreateRequestDataRelationshipsBuildRun {
	this := CiBuildRunCreateRequestDataRelationshipsBuildRun{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *CiBuildRunCreateRequestDataRelationshipsBuildRun) GetData() CiBuildActionRelationshipsBuildRunData {
	if o == nil || IsNil(o.Data) {
		var ret CiBuildActionRelationshipsBuildRunData
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiBuildRunCreateRequestDataRelationshipsBuildRun) GetDataOk() (*CiBuildActionRelationshipsBuildRunData, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *CiBuildRunCreateRequestDataRelationshipsBuildRun) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given CiBuildActionRelationshipsBuildRunData and assigns it to the Data field.
func (o *CiBuildRunCreateRequestDataRelationshipsBuildRun) SetData(v CiBuildActionRelationshipsBuildRunData) {
	o.Data = &v
}

func (o CiBuildRunCreateRequestDataRelationshipsBuildRun) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiBuildRunCreateRequestDataRelationshipsBuildRun) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableCiBuildRunCreateRequestDataRelationshipsBuildRun struct {
	value *CiBuildRunCreateRequestDataRelationshipsBuildRun
	isSet bool
}

func (v NullableCiBuildRunCreateRequestDataRelationshipsBuildRun) Get() *CiBuildRunCreateRequestDataRelationshipsBuildRun {
	return v.value
}

func (v *NullableCiBuildRunCreateRequestDataRelationshipsBuildRun) Set(val *CiBuildRunCreateRequestDataRelationshipsBuildRun) {
	v.value = val
	v.isSet = true
}

func (v NullableCiBuildRunCreateRequestDataRelationshipsBuildRun) IsSet() bool {
	return v.isSet
}

func (v *NullableCiBuildRunCreateRequestDataRelationshipsBuildRun) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiBuildRunCreateRequestDataRelationshipsBuildRun(val *CiBuildRunCreateRequestDataRelationshipsBuildRun) *NullableCiBuildRunCreateRequestDataRelationshipsBuildRun {
	return &NullableCiBuildRunCreateRequestDataRelationshipsBuildRun{value: val, isSet: true}
}

func (v NullableCiBuildRunCreateRequestDataRelationshipsBuildRun) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiBuildRunCreateRequestDataRelationshipsBuildRun) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

