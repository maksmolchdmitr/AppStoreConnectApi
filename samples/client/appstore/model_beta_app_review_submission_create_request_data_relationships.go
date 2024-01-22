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

// checks if the BetaAppReviewSubmissionCreateRequestDataRelationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppReviewSubmissionCreateRequestDataRelationships{}

// BetaAppReviewSubmissionCreateRequestDataRelationships struct for BetaAppReviewSubmissionCreateRequestDataRelationships
type BetaAppReviewSubmissionCreateRequestDataRelationships struct {
	Build BetaAppReviewSubmissionCreateRequestDataRelationshipsBuild `json:"build"`
}

// NewBetaAppReviewSubmissionCreateRequestDataRelationships instantiates a new BetaAppReviewSubmissionCreateRequestDataRelationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppReviewSubmissionCreateRequestDataRelationships(build BetaAppReviewSubmissionCreateRequestDataRelationshipsBuild) *BetaAppReviewSubmissionCreateRequestDataRelationships {
	this := BetaAppReviewSubmissionCreateRequestDataRelationships{}
	this.Build = build
	return &this
}

// NewBetaAppReviewSubmissionCreateRequestDataRelationshipsWithDefaults instantiates a new BetaAppReviewSubmissionCreateRequestDataRelationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppReviewSubmissionCreateRequestDataRelationshipsWithDefaults() *BetaAppReviewSubmissionCreateRequestDataRelationships {
	this := BetaAppReviewSubmissionCreateRequestDataRelationships{}
	return &this
}

// GetBuild returns the Build field value
func (o *BetaAppReviewSubmissionCreateRequestDataRelationships) GetBuild() BetaAppReviewSubmissionCreateRequestDataRelationshipsBuild {
	if o == nil {
		var ret BetaAppReviewSubmissionCreateRequestDataRelationshipsBuild
		return ret
	}

	return o.Build
}

// GetBuildOk returns a tuple with the Build field value
// and a boolean to check if the value has been set.
func (o *BetaAppReviewSubmissionCreateRequestDataRelationships) GetBuildOk() (*BetaAppReviewSubmissionCreateRequestDataRelationshipsBuild, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Build, true
}

// SetBuild sets field value
func (o *BetaAppReviewSubmissionCreateRequestDataRelationships) SetBuild(v BetaAppReviewSubmissionCreateRequestDataRelationshipsBuild) {
	o.Build = v
}

func (o BetaAppReviewSubmissionCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppReviewSubmissionCreateRequestDataRelationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["build"] = o.Build
	return toSerialize, nil
}

type NullableBetaAppReviewSubmissionCreateRequestDataRelationships struct {
	value *BetaAppReviewSubmissionCreateRequestDataRelationships
	isSet bool
}

func (v NullableBetaAppReviewSubmissionCreateRequestDataRelationships) Get() *BetaAppReviewSubmissionCreateRequestDataRelationships {
	return v.value
}

func (v *NullableBetaAppReviewSubmissionCreateRequestDataRelationships) Set(val *BetaAppReviewSubmissionCreateRequestDataRelationships) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppReviewSubmissionCreateRequestDataRelationships) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppReviewSubmissionCreateRequestDataRelationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppReviewSubmissionCreateRequestDataRelationships(val *BetaAppReviewSubmissionCreateRequestDataRelationships) *NullableBetaAppReviewSubmissionCreateRequestDataRelationships {
	return &NullableBetaAppReviewSubmissionCreateRequestDataRelationships{value: val, isSet: true}
}

func (v NullableBetaAppReviewSubmissionCreateRequestDataRelationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppReviewSubmissionCreateRequestDataRelationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


