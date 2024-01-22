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

// checks if the AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData{}

// AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData struct for AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData
type AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData instantiates a new AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData(type_ string, id string) *AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData {
	this := AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseDataWithDefaults instantiates a new AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseDataWithDefaults() *AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData {
	this := AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) SetId(v string) {
	o.Id = v
}

func (o AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData struct {
	value *AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData
	isSet bool
}

func (v NullableAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) Get() *AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData {
	return v.value
}

func (v *NullableAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) Set(val *AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData(val *AppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) *NullableAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData {
	return &NullableAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData{value: val, isSet: true}
}

func (v NullableAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionRelationshipsAppStoreVersionPhasedReleaseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


