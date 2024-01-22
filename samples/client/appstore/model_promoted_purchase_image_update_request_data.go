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

// checks if the PromotedPurchaseImageUpdateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseImageUpdateRequestData{}

// PromotedPurchaseImageUpdateRequestData struct for PromotedPurchaseImageUpdateRequestData
type PromotedPurchaseImageUpdateRequestData struct {
	Type string `json:"type"`
	Id string `json:"id"`
	Attributes *AppClipAdvancedExperienceImageUpdateRequestDataAttributes `json:"attributes,omitempty"`
}

// NewPromotedPurchaseImageUpdateRequestData instantiates a new PromotedPurchaseImageUpdateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseImageUpdateRequestData(type_ string, id string) *PromotedPurchaseImageUpdateRequestData {
	this := PromotedPurchaseImageUpdateRequestData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewPromotedPurchaseImageUpdateRequestDataWithDefaults instantiates a new PromotedPurchaseImageUpdateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseImageUpdateRequestDataWithDefaults() *PromotedPurchaseImageUpdateRequestData {
	this := PromotedPurchaseImageUpdateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *PromotedPurchaseImageUpdateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseImageUpdateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PromotedPurchaseImageUpdateRequestData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *PromotedPurchaseImageUpdateRequestData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseImageUpdateRequestData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PromotedPurchaseImageUpdateRequestData) SetId(v string) {
	o.Id = v
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PromotedPurchaseImageUpdateRequestData) GetAttributes() AppClipAdvancedExperienceImageUpdateRequestDataAttributes {
	if o == nil || IsNil(o.Attributes) {
		var ret AppClipAdvancedExperienceImageUpdateRequestDataAttributes
		return ret
	}
	return *o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseImageUpdateRequestData) GetAttributesOk() (*AppClipAdvancedExperienceImageUpdateRequestDataAttributes, bool) {
	if o == nil || IsNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PromotedPurchaseImageUpdateRequestData) HasAttributes() bool {
	if o != nil && !IsNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given AppClipAdvancedExperienceImageUpdateRequestDataAttributes and assigns it to the Attributes field.
func (o *PromotedPurchaseImageUpdateRequestData) SetAttributes(v AppClipAdvancedExperienceImageUpdateRequestDataAttributes) {
	o.Attributes = &v
}

func (o PromotedPurchaseImageUpdateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseImageUpdateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	if !IsNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}
	return toSerialize, nil
}

type NullablePromotedPurchaseImageUpdateRequestData struct {
	value *PromotedPurchaseImageUpdateRequestData
	isSet bool
}

func (v NullablePromotedPurchaseImageUpdateRequestData) Get() *PromotedPurchaseImageUpdateRequestData {
	return v.value
}

func (v *NullablePromotedPurchaseImageUpdateRequestData) Set(val *PromotedPurchaseImageUpdateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseImageUpdateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseImageUpdateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseImageUpdateRequestData(val *PromotedPurchaseImageUpdateRequestData) *NullablePromotedPurchaseImageUpdateRequestData {
	return &NullablePromotedPurchaseImageUpdateRequestData{value: val, isSet: true}
}

func (v NullablePromotedPurchaseImageUpdateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseImageUpdateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


