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

// checks if the RoutingAppCoverageCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RoutingAppCoverageCreateRequestData{}

// RoutingAppCoverageCreateRequestData struct for RoutingAppCoverageCreateRequestData
type RoutingAppCoverageCreateRequestData struct {
	Type          string                                                    `json:"type"`
	Attributes    AppClipAdvancedExperienceImageCreateRequestDataAttributes `json:"attributes"`
	Relationships AppStoreReviewDetailCreateRequestDataRelationships        `json:"relationships"`
}

// NewRoutingAppCoverageCreateRequestData instantiates a new RoutingAppCoverageCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoutingAppCoverageCreateRequestData(type_ string, attributes AppClipAdvancedExperienceImageCreateRequestDataAttributes, relationships AppStoreReviewDetailCreateRequestDataRelationships) *RoutingAppCoverageCreateRequestData {
	this := RoutingAppCoverageCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewRoutingAppCoverageCreateRequestDataWithDefaults instantiates a new RoutingAppCoverageCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoutingAppCoverageCreateRequestDataWithDefaults() *RoutingAppCoverageCreateRequestData {
	this := RoutingAppCoverageCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *RoutingAppCoverageCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *RoutingAppCoverageCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *RoutingAppCoverageCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *RoutingAppCoverageCreateRequestData) GetAttributes() AppClipAdvancedExperienceImageCreateRequestDataAttributes {
	if o == nil {
		var ret AppClipAdvancedExperienceImageCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *RoutingAppCoverageCreateRequestData) GetAttributesOk() (*AppClipAdvancedExperienceImageCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *RoutingAppCoverageCreateRequestData) SetAttributes(v AppClipAdvancedExperienceImageCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *RoutingAppCoverageCreateRequestData) GetRelationships() AppStoreReviewDetailCreateRequestDataRelationships {
	if o == nil {
		var ret AppStoreReviewDetailCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *RoutingAppCoverageCreateRequestData) GetRelationshipsOk() (*AppStoreReviewDetailCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *RoutingAppCoverageCreateRequestData) SetRelationships(v AppStoreReviewDetailCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o RoutingAppCoverageCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RoutingAppCoverageCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableRoutingAppCoverageCreateRequestData struct {
	value *RoutingAppCoverageCreateRequestData
	isSet bool
}

func (v NullableRoutingAppCoverageCreateRequestData) Get() *RoutingAppCoverageCreateRequestData {
	return v.value
}

func (v *NullableRoutingAppCoverageCreateRequestData) Set(val *RoutingAppCoverageCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableRoutingAppCoverageCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableRoutingAppCoverageCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoutingAppCoverageCreateRequestData(val *RoutingAppCoverageCreateRequestData) *NullableRoutingAppCoverageCreateRequestData {
	return &NullableRoutingAppCoverageCreateRequestData{value: val, isSet: true}
}

func (v NullableRoutingAppCoverageCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoutingAppCoverageCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
