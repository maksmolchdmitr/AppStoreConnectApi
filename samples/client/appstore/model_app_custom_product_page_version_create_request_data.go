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

// checks if the AppCustomProductPageVersionCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCustomProductPageVersionCreateRequestData{}

// AppCustomProductPageVersionCreateRequestData struct for AppCustomProductPageVersionCreateRequestData
type AppCustomProductPageVersionCreateRequestData struct {
	Type string `json:"type"`
	Relationships AppCustomProductPageVersionCreateRequestDataRelationships `json:"relationships"`
}

// NewAppCustomProductPageVersionCreateRequestData instantiates a new AppCustomProductPageVersionCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCustomProductPageVersionCreateRequestData(type_ string, relationships AppCustomProductPageVersionCreateRequestDataRelationships) *AppCustomProductPageVersionCreateRequestData {
	this := AppCustomProductPageVersionCreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewAppCustomProductPageVersionCreateRequestDataWithDefaults instantiates a new AppCustomProductPageVersionCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCustomProductPageVersionCreateRequestDataWithDefaults() *AppCustomProductPageVersionCreateRequestData {
	this := AppCustomProductPageVersionCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppCustomProductPageVersionCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppCustomProductPageVersionCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetRelationships returns the Relationships field value
func (o *AppCustomProductPageVersionCreateRequestData) GetRelationships() AppCustomProductPageVersionCreateRequestDataRelationships {
	if o == nil {
		var ret AppCustomProductPageVersionCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppCustomProductPageVersionCreateRequestData) GetRelationshipsOk() (*AppCustomProductPageVersionCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppCustomProductPageVersionCreateRequestData) SetRelationships(v AppCustomProductPageVersionCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppCustomProductPageVersionCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCustomProductPageVersionCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppCustomProductPageVersionCreateRequestData struct {
	value *AppCustomProductPageVersionCreateRequestData
	isSet bool
}

func (v NullableAppCustomProductPageVersionCreateRequestData) Get() *AppCustomProductPageVersionCreateRequestData {
	return v.value
}

func (v *NullableAppCustomProductPageVersionCreateRequestData) Set(val *AppCustomProductPageVersionCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCustomProductPageVersionCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCustomProductPageVersionCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCustomProductPageVersionCreateRequestData(val *AppCustomProductPageVersionCreateRequestData) *NullableAppCustomProductPageVersionCreateRequestData {
	return &NullableAppCustomProductPageVersionCreateRequestData{value: val, isSet: true}
}

func (v NullableAppCustomProductPageVersionCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCustomProductPageVersionCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

