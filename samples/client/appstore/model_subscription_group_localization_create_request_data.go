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

// checks if the SubscriptionGroupLocalizationCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionGroupLocalizationCreateRequestData{}

// SubscriptionGroupLocalizationCreateRequestData struct for SubscriptionGroupLocalizationCreateRequestData
type SubscriptionGroupLocalizationCreateRequestData struct {
	Type string `json:"type"`
	Attributes SubscriptionGroupLocalizationCreateRequestDataAttributes `json:"attributes"`
	Relationships SubscriptionGroupLocalizationCreateRequestDataRelationships `json:"relationships"`
}

// NewSubscriptionGroupLocalizationCreateRequestData instantiates a new SubscriptionGroupLocalizationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionGroupLocalizationCreateRequestData(type_ string, attributes SubscriptionGroupLocalizationCreateRequestDataAttributes, relationships SubscriptionGroupLocalizationCreateRequestDataRelationships) *SubscriptionGroupLocalizationCreateRequestData {
	this := SubscriptionGroupLocalizationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewSubscriptionGroupLocalizationCreateRequestDataWithDefaults instantiates a new SubscriptionGroupLocalizationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionGroupLocalizationCreateRequestDataWithDefaults() *SubscriptionGroupLocalizationCreateRequestData {
	this := SubscriptionGroupLocalizationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *SubscriptionGroupLocalizationCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubscriptionGroupLocalizationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *SubscriptionGroupLocalizationCreateRequestData) GetAttributes() SubscriptionGroupLocalizationCreateRequestDataAttributes {
	if o == nil {
		var ret SubscriptionGroupLocalizationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationCreateRequestData) GetAttributesOk() (*SubscriptionGroupLocalizationCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *SubscriptionGroupLocalizationCreateRequestData) SetAttributes(v SubscriptionGroupLocalizationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *SubscriptionGroupLocalizationCreateRequestData) GetRelationships() SubscriptionGroupLocalizationCreateRequestDataRelationships {
	if o == nil {
		var ret SubscriptionGroupLocalizationCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *SubscriptionGroupLocalizationCreateRequestData) GetRelationshipsOk() (*SubscriptionGroupLocalizationCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *SubscriptionGroupLocalizationCreateRequestData) SetRelationships(v SubscriptionGroupLocalizationCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o SubscriptionGroupLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionGroupLocalizationCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableSubscriptionGroupLocalizationCreateRequestData struct {
	value *SubscriptionGroupLocalizationCreateRequestData
	isSet bool
}

func (v NullableSubscriptionGroupLocalizationCreateRequestData) Get() *SubscriptionGroupLocalizationCreateRequestData {
	return v.value
}

func (v *NullableSubscriptionGroupLocalizationCreateRequestData) Set(val *SubscriptionGroupLocalizationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionGroupLocalizationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionGroupLocalizationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionGroupLocalizationCreateRequestData(val *SubscriptionGroupLocalizationCreateRequestData) *NullableSubscriptionGroupLocalizationCreateRequestData {
	return &NullableSubscriptionGroupLocalizationCreateRequestData{value: val, isSet: true}
}

func (v NullableSubscriptionGroupLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionGroupLocalizationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


