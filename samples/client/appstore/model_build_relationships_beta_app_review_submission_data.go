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

// checks if the BuildRelationshipsBetaAppReviewSubmissionData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BuildRelationshipsBetaAppReviewSubmissionData{}

// BuildRelationshipsBetaAppReviewSubmissionData struct for BuildRelationshipsBetaAppReviewSubmissionData
type BuildRelationshipsBetaAppReviewSubmissionData struct {
	Type string `json:"type"`
	Id string `json:"id"`
}

// NewBuildRelationshipsBetaAppReviewSubmissionData instantiates a new BuildRelationshipsBetaAppReviewSubmissionData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildRelationshipsBetaAppReviewSubmissionData(type_ string, id string) *BuildRelationshipsBetaAppReviewSubmissionData {
	this := BuildRelationshipsBetaAppReviewSubmissionData{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewBuildRelationshipsBetaAppReviewSubmissionDataWithDefaults instantiates a new BuildRelationshipsBetaAppReviewSubmissionData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildRelationshipsBetaAppReviewSubmissionDataWithDefaults() *BuildRelationshipsBetaAppReviewSubmissionData {
	this := BuildRelationshipsBetaAppReviewSubmissionData{}
	return &this
}

// GetType returns the Type field value
func (o *BuildRelationshipsBetaAppReviewSubmissionData) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *BuildRelationshipsBetaAppReviewSubmissionData) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *BuildRelationshipsBetaAppReviewSubmissionData) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *BuildRelationshipsBetaAppReviewSubmissionData) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *BuildRelationshipsBetaAppReviewSubmissionData) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *BuildRelationshipsBetaAppReviewSubmissionData) SetId(v string) {
	o.Id = v
}

func (o BuildRelationshipsBetaAppReviewSubmissionData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BuildRelationshipsBetaAppReviewSubmissionData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

type NullableBuildRelationshipsBetaAppReviewSubmissionData struct {
	value *BuildRelationshipsBetaAppReviewSubmissionData
	isSet bool
}

func (v NullableBuildRelationshipsBetaAppReviewSubmissionData) Get() *BuildRelationshipsBetaAppReviewSubmissionData {
	return v.value
}

func (v *NullableBuildRelationshipsBetaAppReviewSubmissionData) Set(val *BuildRelationshipsBetaAppReviewSubmissionData) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildRelationshipsBetaAppReviewSubmissionData) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildRelationshipsBetaAppReviewSubmissionData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildRelationshipsBetaAppReviewSubmissionData(val *BuildRelationshipsBetaAppReviewSubmissionData) *NullableBuildRelationshipsBetaAppReviewSubmissionData {
	return &NullableBuildRelationshipsBetaAppReviewSubmissionData{value: val, isSet: true}
}

func (v NullableBuildRelationshipsBetaAppReviewSubmissionData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildRelationshipsBetaAppReviewSubmissionData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


