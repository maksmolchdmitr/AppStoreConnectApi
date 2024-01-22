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

// checks if the AppStoreVersionRelationshipsRoutingAppCoverageData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionRelationshipsRoutingAppCoverageData{}

// AppStoreVersionRelationshipsRoutingAppCoverageData struct for AppStoreVersionRelationshipsRoutingAppCoverageData
type AppStoreVersionRelationshipsRoutingAppCoverageData struct {
	Type string `json:"type"`
	Id   string `json:"id"`
}

// NewAppStoreVersionRelationshipsRoutingAppCoverageData instantiates a new AppStoreVersionRelationshipsRoutingAppCoverageData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionRelationshipsRoutingAppCoverageData(type_ string, id string) *AppStoreVersionRelationshipsRoutingAppCoverageData {
	this := AppStoreVersionRelationshipsRoutingAppCoverageData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppStoreVersionRelationshipsRoutingAppCoverageDataWithDefaults instantiates a new AppStoreVersionRelationshipsRoutingAppCoverageData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionRelationshipsRoutingAppCoverageDataWithDefaults() *AppStoreVersionRelationshipsRoutingAppCoverageData {
	this := AppStoreVersionRelationshipsRoutingAppCoverageData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionRelationshipsRoutingAppCoverageData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationshipsRoutingAppCoverageData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionRelationshipsRoutingAppCoverageData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreVersionRelationshipsRoutingAppCoverageData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationshipsRoutingAppCoverageData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreVersionRelationshipsRoutingAppCoverageData) SetId(v string) {
	o.Id = v
}

func (o AppStoreVersionRelationshipsRoutingAppCoverageData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionRelationshipsRoutingAppCoverageData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppStoreVersionRelationshipsRoutingAppCoverageData struct {
	value *AppStoreVersionRelationshipsRoutingAppCoverageData
	isSet bool
}

func (v NullableAppStoreVersionRelationshipsRoutingAppCoverageData) Get() *AppStoreVersionRelationshipsRoutingAppCoverageData {
	return v.value
}

func (v *NullableAppStoreVersionRelationshipsRoutingAppCoverageData) Set(val *AppStoreVersionRelationshipsRoutingAppCoverageData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionRelationshipsRoutingAppCoverageData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionRelationshipsRoutingAppCoverageData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionRelationshipsRoutingAppCoverageData(val *AppStoreVersionRelationshipsRoutingAppCoverageData) *NullableAppStoreVersionRelationshipsRoutingAppCoverageData {
	return &NullableAppStoreVersionRelationshipsRoutingAppCoverageData{value: val, isSet: true}
}

func (v NullableAppStoreVersionRelationshipsRoutingAppCoverageData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionRelationshipsRoutingAppCoverageData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
