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

// checks if the AppPreOrderCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppPreOrderCreateRequestData{}

// AppPreOrderCreateRequestData struct for AppPreOrderCreateRequestData
type AppPreOrderCreateRequestData struct {
	Type string `json:"type"`
	Attributes *AppPreOrderCreateRequestDataAttributes `json:"attributes,omitempty"`
	Relationships AppEventCreateRequestDataRelationships `json:"relationships"`
}

// NewAppPreOrderCreateRequestData instantiates a new AppPreOrderCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppPreOrderCreateRequestData(type_ string, relationships AppEventCreateRequestDataRelationships) *AppPreOrderCreateRequestData {
	this := AppPreOrderCreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewAppPreOrderCreateRequestDataWithDefaults instantiates a new AppPreOrderCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppPreOrderCreateRequestDataWithDefaults() *AppPreOrderCreateRequestData {
	this := AppPreOrderCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppPreOrderCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppPreOrderCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppPreOrderCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *AppPreOrderCreateRequestData) GetAttributes() AppPreOrderCreateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppPreOrderCreateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppPreOrderCreateRequestData) GetAttributesOk() (*AppPreOrderCreateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *AppPreOrderCreateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppPreOrderCreateRequestDataAttributes and assigns it to the Attributes field.
func (o *AppPreOrderCreateRequestData) SetAttributes(v AppPreOrderCreateRequestDataAttributes) {
	o.Attributes = &v
}

// GetRelationships returns the Relationships field value
func (o *AppPreOrderCreateRequestData) GetRelationships() AppEventCreateRequestDataRelationships {
	if o == nil {
		var ret AppEventCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppPreOrderCreateRequestData) GetRelationshipsOk() (*AppEventCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppPreOrderCreateRequestData) SetRelationships(v AppEventCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppPreOrderCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppPreOrderCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppPreOrderCreateRequestData struct {
	value *AppPreOrderCreateRequestData
	isSet bool
}

func (v NullableAppPreOrderCreateRequestData) Get() *AppPreOrderCreateRequestData {
	return v.value
}

func (v *NullableAppPreOrderCreateRequestData) Set(val *AppPreOrderCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppPreOrderCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppPreOrderCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppPreOrderCreateRequestData(val *AppPreOrderCreateRequestData) *NullableAppPreOrderCreateRequestData {
	return &NullableAppPreOrderCreateRequestData{value: val, isSet: true}
}

func (v NullableAppPreOrderCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppPreOrderCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


