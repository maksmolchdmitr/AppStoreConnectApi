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

// checks if the AppRelationshipsBetaGroupsDataInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppRelationshipsBetaGroupsDataInner{}

// AppRelationshipsBetaGroupsDataInner struct for AppRelationshipsBetaGroupsDataInner
type AppRelationshipsBetaGroupsDataInner struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewAppRelationshipsBetaGroupsDataInner instantiates a new AppRelationshipsBetaGroupsDataInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppRelationshipsBetaGroupsDataInner(type_ string, id string) *AppRelationshipsBetaGroupsDataInner {
	this := AppRelationshipsBetaGroupsDataInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppRelationshipsBetaGroupsDataInnerWithDefaults instantiates a new AppRelationshipsBetaGroupsDataInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppRelationshipsBetaGroupsDataInnerWithDefaults() *AppRelationshipsBetaGroupsDataInner {
	this := AppRelationshipsBetaGroupsDataInner{}
	return &this
}

// GetType returns the Type field value
func (o *AppRelationshipsBetaGroupsDataInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsBetaGroupsDataInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppRelationshipsBetaGroupsDataInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppRelationshipsBetaGroupsDataInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppRelationshipsBetaGroupsDataInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppRelationshipsBetaGroupsDataInner) SetId(v string) {
	o.Id = v
}

func (o AppRelationshipsBetaGroupsDataInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppRelationshipsBetaGroupsDataInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppRelationshipsBetaGroupsDataInner struct {
	value *AppRelationshipsBetaGroupsDataInner
	isSet bool
}

func (v NullableAppRelationshipsBetaGroupsDataInner) Get() *AppRelationshipsBetaGroupsDataInner {
	return v.value
}

func (v *NullableAppRelationshipsBetaGroupsDataInner) Set(val *AppRelationshipsBetaGroupsDataInner) {
	v.value = val
	v.isSet = true
}

func (v NullableAppRelationshipsBetaGroupsDataInner) IsSet() bool {
	return v.isSet
}

func (v *NullableAppRelationshipsBetaGroupsDataInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppRelationshipsBetaGroupsDataInner(val *AppRelationshipsBetaGroupsDataInner) *NullableAppRelationshipsBetaGroupsDataInner {
	return &NullableAppRelationshipsBetaGroupsDataInner{value: val, isSet: true}
}

func (v NullableAppRelationshipsBetaGroupsDataInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppRelationshipsBetaGroupsDataInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}