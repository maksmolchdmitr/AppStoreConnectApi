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

// checks if the BuildBetaDetailUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildBetaDetailUpdateRequestData{}

// BuildBetaDetailUpdateRequestData struct for BuildBetaDetailUpdateRequestData
type BuildBetaDetailUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *BuildBetaDetailUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewBuildBetaDetailUpdateRequestData instantiates a new BuildBetaDetailUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildBetaDetailUpdateRequestData(type_ string, id string) *BuildBetaDetailUpdateRequestData {
	this := BuildBetaDetailUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBuildBetaDetailUpdateRequestDataWithDefaults instantiates a new BuildBetaDetailUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildBetaDetailUpdateRequestDataWithDefaults() *BuildBetaDetailUpdateRequestData {
	this := BuildBetaDetailUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *BuildBetaDetailUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BuildBetaDetailUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BuildBetaDetailUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BuildBetaDetailUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *BuildBetaDetailUpdateRequestData) GetAttributes() BuildBetaDetailUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret BuildBetaDetailUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildBetaDetailUpdateRequestData) GetAttributesOk() (*BuildBetaDetailUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *BuildBetaDetailUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given BuildBetaDetailUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *BuildBetaDetailUpdateRequestData) SetAttributes(v BuildBetaDetailUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o BuildBetaDetailUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildBetaDetailUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableBuildBetaDetailUpdateRequestData struct {
	value *BuildBetaDetailUpdateRequestData
	isSet bool
}

func (v NullableBuildBetaDetailUpdateRequestData) Get() *BuildBetaDetailUpdateRequestData {
	return v.value
}

func (v *NullableBuildBetaDetailUpdateRequestData) Set(val *BuildBetaDetailUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildBetaDetailUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildBetaDetailUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildBetaDetailUpdateRequestData(val *BuildBetaDetailUpdateRequestData) *NullableBuildBetaDetailUpdateRequestData {
	return &NullableBuildBetaDetailUpdateRequestData{value: val, isSet: true}
}

func (v NullableBuildBetaDetailUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildBetaDetailUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

