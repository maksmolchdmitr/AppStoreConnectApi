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

// checks if the PromotedPurchaseImageCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PromotedPurchaseImageCreateRequestData{}

// PromotedPurchaseImageCreateRequestData struct for PromotedPurchaseImageCreateRequestData
type PromotedPurchaseImageCreateRequestData struct {
	Type          string                                                    `json:"type"`
	Attributes    AppClipAdvancedExperienceImageCreateRequestDataAttributes `json:"attributes"`
	Relationships PromotedPurchaseImageCreateRequestDataRelationships       `json:"relationships"`
}

// NewPromotedPurchaseImageCreateRequestData instantiates a new PromotedPurchaseImageCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPromotedPurchaseImageCreateRequestData(type_ string, attributes AppClipAdvancedExperienceImageCreateRequestDataAttributes, relationships PromotedPurchaseImageCreateRequestDataRelationships) *PromotedPurchaseImageCreateRequestData {
	this := PromotedPurchaseImageCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewPromotedPurchaseImageCreateRequestDataWithDefaults instantiates a new PromotedPurchaseImageCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPromotedPurchaseImageCreateRequestDataWithDefaults() *PromotedPurchaseImageCreateRequestData {
	this := PromotedPurchaseImageCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *PromotedPurchaseImageCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseImageCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PromotedPurchaseImageCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *PromotedPurchaseImageCreateRequestData) GetAttributes() AppClipAdvancedExperienceImageCreateRequestDataAttributes {
	if o == nil {
		var ret AppClipAdvancedExperienceImageCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseImageCreateRequestData) GetAttributesOk() (*AppClipAdvancedExperienceImageCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *PromotedPurchaseImageCreateRequestData) SetAttributes(v AppClipAdvancedExperienceImageCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *PromotedPurchaseImageCreateRequestData) GetRelationships() PromotedPurchaseImageCreateRequestDataRelationships {
	if o == nil {
		var ret PromotedPurchaseImageCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *PromotedPurchaseImageCreateRequestData) GetRelationshipsOk() (*PromotedPurchaseImageCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *PromotedPurchaseImageCreateRequestData) SetRelationships(v PromotedPurchaseImageCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o PromotedPurchaseImageCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PromotedPurchaseImageCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullablePromotedPurchaseImageCreateRequestData struct {
	value *PromotedPurchaseImageCreateRequestData
	isSet bool
}

func (v NullablePromotedPurchaseImageCreateRequestData) Get() *PromotedPurchaseImageCreateRequestData {
	return v.value
}

func (v *NullablePromotedPurchaseImageCreateRequestData) Set(val *PromotedPurchaseImageCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullablePromotedPurchaseImageCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullablePromotedPurchaseImageCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePromotedPurchaseImageCreateRequestData(val *PromotedPurchaseImageCreateRequestData) *NullablePromotedPurchaseImageCreateRequestData {
	return &NullablePromotedPurchaseImageCreateRequestData{value: val, isSet: true}
}

func (v NullablePromotedPurchaseImageCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePromotedPurchaseImageCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
