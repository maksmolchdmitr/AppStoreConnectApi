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

// checks if the AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData{}

// AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData struct for AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData
type AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData instantiates a new AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData(type_ string, id string) *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData {
	this := AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionDataWithDefaults instantiates a new AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionDataWithDefaults() *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData {
	this := AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData{}
	return &this
}

// GetType returns the Type field value
func (o *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) SetId(v string) {
	o.Id = v
}

func (o AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData struct {
	value *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData
	isSet bool
}

func (v NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) Get() *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData {
	return v.value
}

func (v *NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) Set(val *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData(val *AppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) *NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData {
	return &NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData{value: val, isSet: true}
}

func (v NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageLocalizationRelationshipsAppCustomProductPageVersionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


