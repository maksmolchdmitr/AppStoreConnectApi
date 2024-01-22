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

// checks if the RoutingAppCoverageUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingAppCoverageUpdateRequestData{}

// RoutingAppCoverageUpdateRequestData struct for RoutingAppCoverageUpdateRequestData
type RoutingAppCoverageUpdateRequestData struct {
	Type       string                                                     `json:"type"`
	Id         string                                                     `json:"id"`
	Attributes *AppClipAdvancedExperienceImageUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewRoutingAppCoverageUpdateRequestData instantiates a new RoutingAppCoverageUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingAppCoverageUpdateRequestData(type_ string, id string) *RoutingAppCoverageUpdateRequestData {
	this := RoutingAppCoverageUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewRoutingAppCoverageUpdateRequestDataWithDefaults instantiates a new RoutingAppCoverageUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingAppCoverageUpdateRequestDataWithDefaults() *RoutingAppCoverageUpdateRequestData {
	this := RoutingAppCoverageUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *RoutingAppCoverageUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoutingAppCoverageUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RoutingAppCoverageUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *RoutingAppCoverageUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *RoutingAppCoverageUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *RoutingAppCoverageUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *RoutingAppCoverageUpdateRequestData) GetAttributes() AppClipAdvancedExperienceImageUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipAdvancedExperienceImageUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RoutingAppCoverageUpdateRequestData) GetAttributesOk() (*AppClipAdvancedExperienceImageUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *RoutingAppCoverageUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipAdvancedExperienceImageUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *RoutingAppCoverageUpdateRequestData) SetAttributes(v AppClipAdvancedExperienceImageUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o RoutingAppCoverageUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingAppCoverageUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullableRoutingAppCoverageUpdateRequestData struct {
	value *RoutingAppCoverageUpdateRequestData
	isSet bool
}

func (v NullableRoutingAppCoverageUpdateRequestData) Get() *RoutingAppCoverageUpdateRequestData {
	return v.value
}

func (v *NullableRoutingAppCoverageUpdateRequestData) Set(val *RoutingAppCoverageUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingAppCoverageUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingAppCoverageUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingAppCoverageUpdateRequestData(val *RoutingAppCoverageUpdateRequestData) *NullableRoutingAppCoverageUpdateRequestData {
	return &NullableRoutingAppCoverageUpdateRequestData{value: val, isSet: true}
}

func (v NullableRoutingAppCoverageUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingAppCoverageUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
