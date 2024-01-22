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

// checks if the AppStoreVersionCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionCreateRequestData{}

// AppStoreVersionCreateRequestData struct for AppStoreVersionCreateRequestData
type AppStoreVersionCreateRequestData struct {
	Type string `json:"type"`
	Attributes AppStoreVersionCreateRequestDataAttributes `json:"attributes"`
	Relationships AppStoreVersionCreateRequestDataRelationships `json:"relationships"`
}

// NewAppStoreVersionCreateRequestData instantiates a new AppStoreVersionCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionCreateRequestData(type_ string, attributes AppStoreVersionCreateRequestDataAttributes, relationships AppStoreVersionCreateRequestDataRelationships) *AppStoreVersionCreateRequestData {
	this := AppStoreVersionCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewAppStoreVersionCreateRequestDataWithDefaults instantiates a new AppStoreVersionCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionCreateRequestDataWithDefaults() *AppStoreVersionCreateRequestData {
	this := AppStoreVersionCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *AppStoreVersionCreateRequestData) GetAttributes() AppStoreVersionCreateRequestDataAttributes {
	if o == nil {
		var ret AppStoreVersionCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionCreateRequestData) GetAttributesOk() (*AppStoreVersionCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *AppStoreVersionCreateRequestData) SetAttributes(v AppStoreVersionCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *AppStoreVersionCreateRequestData) GetRelationships() AppStoreVersionCreateRequestDataRelationships {
	if o == nil {
		var ret AppStoreVersionCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionCreateRequestData) GetRelationshipsOk() (*AppStoreVersionCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppStoreVersionCreateRequestData) SetRelationships(v AppStoreVersionCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppStoreVersionCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppStoreVersionCreateRequestData struct {
	value *AppStoreVersionCreateRequestData
	isSet bool
}

func (v NullableAppStoreVersionCreateRequestData) Get() *AppStoreVersionCreateRequestData {
	return v.value
}

func (v *NullableAppStoreVersionCreateRequestData) Set(val *AppStoreVersionCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionCreateRequestData(val *AppStoreVersionCreateRequestData) *NullableAppStoreVersionCreateRequestData {
	return &NullableAppStoreVersionCreateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreVersionCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

