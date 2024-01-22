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

// checks if the InAppPurchaseSubmissionCreateRequestData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &InAppPurchaseSubmissionCreateRequestData{}

// InAppPurchaseSubmissionCreateRequestData struct for InAppPurchaseSubmissionCreateRequestData
type InAppPurchaseSubmissionCreateRequestData struct {
	Type string `json:"type"`
	Relationships InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships `json:"relationships"`
}

// NewInAppPurchaseSubmissionCreateRequestData instantiates a new InAppPurchaseSubmissionCreateRequestData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInAppPurchaseSubmissionCreateRequestData(type_ string, relationships InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships) *InAppPurchaseSubmissionCreateRequestData {
	this := InAppPurchaseSubmissionCreateRequestData{}
	this.Type = type_
	this.Relationships = relationships
	return &this
}

// NewInAppPurchaseSubmissionCreateRequestDataWithDefaults instantiates a new InAppPurchaseSubmissionCreateRequestData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInAppPurchaseSubmissionCreateRequestDataWithDefaults() *InAppPurchaseSubmissionCreateRequestData {
	this := InAppPurchaseSubmissionCreateRequestData{}
	return &this
}

// GetType returns the Type field value
func (o *InAppPurchaseSubmissionCreateRequestData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseSubmissionCreateRequestData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *InAppPurchaseSubmissionCreateRequestData) SetType(v string) {
	o.Type = v
}

// GetRelationships returns the Relationships field value
func (o *InAppPurchaseSubmissionCreateRequestData) GetRelationships() InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships {
	if o == nil {
		var ret InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships
		return ret
	}

	return o.Relationships
}

// GetRelationshipsOk returns a tuple with the Relationships field value
// and a boolean to check if the value has been set.
func (o *InAppPurchaseSubmissionCreateRequestData) GetRelationshipsOk() (*InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Relationships, true
}

// SetRelationships sets field value
func (o *InAppPurchaseSubmissionCreateRequestData) SetRelationships(v InAppPurchaseAppStoreReviewScreenshotCreateRequestDataRelationships) {
	o.Relationships = v
}

func (o InAppPurchaseSubmissionCreateRequestData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o InAppPurchaseSubmissionCreateRequestData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["relationships"] = o.Relationships
	return toSerialize, nil
}

type NullableInAppPurchaseSubmissionCreateRequestData struct {
	value *InAppPurchaseSubmissionCreateRequestData
	isSet bool
}

func (v NullableInAppPurchaseSubmissionCreateRequestData) Get() *InAppPurchaseSubmissionCreateRequestData {
	return v.value
}

func (v *NullableInAppPurchaseSubmissionCreateRequestData) Set(val *InAppPurchaseSubmissionCreateRequestData) {
	v.value = val
	v.isSet = true
}

func (v NullableInAppPurchaseSubmissionCreateRequestData) IsSet() bool {
	return v.isSet
}

func (v *NullableInAppPurchaseSubmissionCreateRequestData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInAppPurchaseSubmissionCreateRequestData(val *InAppPurchaseSubmissionCreateRequestData) *NullableInAppPurchaseSubmissionCreateRequestData {
	return &NullableInAppPurchaseSubmissionCreateRequestData{value: val, isSet: true}
}

func (v NullableInAppPurchaseSubmissionCreateRequestData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInAppPurchaseSubmissionCreateRequestData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


