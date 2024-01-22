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

// checks if the BetaBuildLocalizationCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaBuildLocalizationCreateRequestData{}

// BetaBuildLocalizationCreateRequestData struct for BetaBuildLocalizationCreateRequestData
type BetaBuildLocalizationCreateRequestData struct {
	Type string `json:"type"`
	Attributes BetaBuildLocalizationCreateRequestDataAttributes `json:"attributes"`
	Relationships BetaAppReviewSubmissionCreateRequestDataRelationships `json:"relationships"`
}

// NewBetaBuildLocalizationCreateRequestData instantiates a new BetaBuildLocalizationCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaBuildLocalizationCreateRequestData(type_ string, attributes BetaBuildLocalizationCreateRequestDataAttributes, relationships BetaAppReviewSubmissionCreateRequestDataRelationships) *BetaBuildLocalizationCreateRequestData {
	this := BetaBuildLocalizationCreateRequestData{}
	this.Type = type_
	this.Attributes = attributes
	this.Relationships = relationships
	return &this
}

// NewBetaBuildLocalizationCreateRequestDataWithDefaults instantiates a new BetaBuildLocalizationCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaBuildLocalizationCreateRequestDataWithDefaults() *BetaBuildLocalizationCreateRequestData {
	this := BetaBuildLocalizationCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *BetaBuildLocalizationCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BetaBuildLocalizationCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BetaBuildLocalizationCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *BetaBuildLocalizationCreateRequestData) GetAttributes() BetaBuildLocalizationCreateRequestDataAttributes {
	if o == nil {
		var ret BetaBuildLocalizationCreateRequestDataAttributes
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *BetaBuildLocalizationCreateRequestData) GetAttributesOk() (*BetaBuildLocalizationCreateRequestDataAttributes, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Attributes, true
}

// SetAttributes sets field value
func (o *BetaBuildLocalizationCreateRequestData) SetAttributes(v BetaBuildLocalizationCreateRequestDataAttributes) {
	o.Attributes = v
}

// GetRelationships returns the Relationships field value
func (o *BetaBuildLocalizationCreateRequestData) GetRelationships() BetaAppReviewSubmissionCreateRequestDataRelationships {
	if o == nil {
		var ret BetaAppReviewSubmissionCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *BetaBuildLocalizationCreateRequestData) GetRelationshipsOk() (*BetaAppReviewSubmissionCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *BetaBuildLocalizationCreateRequestData) SetRelationships(v BetaAppReviewSubmissionCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o BetaBuildLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaBuildLocalizationCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["attributes"] = o.Attributes
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableBetaBuildLocalizationCreateRequestData struct {
	value *BetaBuildLocalizationCreateRequestData
	isSet bool
}

func (v NullableBetaBuildLocalizationCreateRequestData) Get() *BetaBuildLocalizationCreateRequestData {
	return v.value
}

func (v *NullableBetaBuildLocalizationCreateRequestData) Set(val *BetaBuildLocalizationCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaBuildLocalizationCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaBuildLocalizationCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaBuildLocalizationCreateRequestData(val *BetaBuildLocalizationCreateRequestData) *NullableBetaBuildLocalizationCreateRequestData {
	return &NullableBetaBuildLocalizationCreateRequestData{value: val, isSet: true}
}

func (v NullableBetaBuildLocalizationCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaBuildLocalizationCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


