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

// checks if the AppStoreVersionSubmissionCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionSubmissionCreateRequestData{}

// AppStoreVersionSubmissionCreateRequestData struct for AppStoreVersionSubmissionCreateRequestData
type AppStoreVersionSubmissionCreateRequestData struct {
	Type          string                                             `json:"type"`
	Relationships AppStoreReviewDetailCreateRequestDataRelationships `json:"relationships"`
}

// NewAppStoreVersionSubmissionCreateRequestData instantiates a new AppStoreVersionSubmissionCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionSubmissionCreateRequestData(type_ string, relationships AppStoreReviewDetailCreateRequestDataRelationships) *AppStoreVersionSubmissionCreateRequestData {
	this := AppStoreVersionSubmissionCreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewAppStoreVersionSubmissionCreateRequestDataWithDefaults instantiates a new AppStoreVersionSubmissionCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionSubmissionCreateRequestDataWithDefaults() *AppStoreVersionSubmissionCreateRequestData {
	this := AppStoreVersionSubmissionCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *AppStoreVersionSubmissionCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionSubmissionCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AppStoreVersionSubmissionCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetRelationships returns the Relationships field value
func (o *AppStoreVersionSubmissionCreateRequestData) GetRelationships() AppStoreReviewDetailCreateRequestDataRelationships {
	if o == nil {
		var ret AppStoreReviewDetailCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionSubmissionCreateRequestData) GetRelationshipsOk() (*AppStoreReviewDetailCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *AppStoreVersionSubmissionCreateRequestData) SetRelationships(v AppStoreReviewDetailCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o AppStoreVersionSubmissionCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionSubmissionCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableAppStoreVersionSubmissionCreateRequestData struct {
	value *AppStoreVersionSubmissionCreateRequestData
	isSet bool
}

func (v NullableAppStoreVersionSubmissionCreateRequestData) Get() *AppStoreVersionSubmissionCreateRequestData {
	return v.value
}

func (v *NullableAppStoreVersionSubmissionCreateRequestData) Set(val *AppStoreVersionSubmissionCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionSubmissionCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionSubmissionCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionSubmissionCreateRequestData(val *AppStoreVersionSubmissionCreateRequestData) *NullableAppStoreVersionSubmissionCreateRequestData {
	return &NullableAppStoreVersionSubmissionCreateRequestData{value: val, isSet: true}
}

func (v NullableAppStoreVersionSubmissionCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionSubmissionCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
