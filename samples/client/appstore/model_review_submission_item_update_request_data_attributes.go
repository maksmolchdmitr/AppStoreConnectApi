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

// checks if the ReviewSubmissionItemUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSubmissionItemUpdateRequestDataAttributes{}

// ReviewSubmissionItemUpdateRequestDataAttributes struct for ReviewSubmissionItemUpdateRequestDataAttributes
type ReviewSubmissionItemUpdateRequestDataAttributes struct {
	Resolved *bool `json:"resolved,omitempty"`
	Removed *bool `json:"removed,omitempty"`
}

// NewReviewSubmissionItemUpdateRequestDataAttributes instantiates a new ReviewSubmissionItemUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSubmissionItemUpdateRequestDataAttributes() *ReviewSubmissionItemUpdateRequestDataAttributes {
	this := ReviewSubmissionItemUpdateRequestDataAttributes{}
	return &this
}

// NewReviewSubmissionItemUpdateRequestDataAttributesWithDefaults instantiates a new ReviewSubmissionItemUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSubmissionItemUpdateRequestDataAttributesWithDefaults() *ReviewSubmissionItemUpdateRequestDataAttributes {
	this := ReviewSubmissionItemUpdateRequestDataAttributes{}
	return &this
}

// GetResolved returns the Resolved field value if set, zero value otherwise.
func (o *ReviewSubmissionItemUpdateRequestDataAttributes) GetResolved() bool {
	if o == nil || IsNil(o.Resolved) {
		var ret bool
		return ret
	}
	return *o.Resolved
}

// GetResolvedOk returns a tuple with the Resolved field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemUpdateRequestDataAttributes) GetResolvedOk() (*bool, bool) {
	if o == nil || IsNil(o.Resolved) {
		return nil, false
	}
	return o.Resolved, true
}

// HasResolved returns a boolean if a field has been set.
func (o *ReviewSubmissionItemUpdateRequestDataAttributes) HasResolved() bool {
	if o != nil && !IsNil(o.Resolved) {
		return true
	}

	return false
}

// SetResolved gets a reference to the given bool and assigns it to the Resolved field.
func (o *ReviewSubmissionItemUpdateRequestDataAttributes) SetResolved(v bool) {
	o.Resolved = &v
}

// GetRemoved returns the Removed field value if set, zero value otherwise.
func (o *ReviewSubmissionItemUpdateRequestDataAttributes) GetRemoved() bool {
	if o == nil || IsNil(o.Removed) {
		var ret bool
		return ret
	}
	return *o.Removed
}

// GetRemovedOk returns a tuple with the Removed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionItemUpdateRequestDataAttributes) GetRemovedOk() (*bool, bool) {
	if o == nil || IsNil(o.Removed) {
		return nil, false
	}
	return o.Removed, true
}

// HasRemoved returns a boolean if a field has been set.
func (o *ReviewSubmissionItemUpdateRequestDataAttributes) HasRemoved() bool {
	if o != nil && !IsNil(o.Removed) {
		return true
	}

	return false
}

// SetRemoved gets a reference to the given bool and assigns it to the Removed field.
func (o *ReviewSubmissionItemUpdateRequestDataAttributes) SetRemoved(v bool) {
	o.Removed = &v
}

func (o ReviewSubmissionItemUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSubmissionItemUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Resolved) {
		toSerialize["resolved"] = o.Resolved
	}
	if !IsNil(o.Removed) {
		toSerialize["removed"] = o.Removed
	}
	return toSerialize, nil
}

type NullableReviewSubmissionItemUpdateRequestDataAttributes struct {
	value *ReviewSubmissionItemUpdateRequestDataAttributes
	isSet bool
}

func (v NullableReviewSubmissionItemUpdateRequestDataAttributes) Get() *ReviewSubmissionItemUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableReviewSubmissionItemUpdateRequestDataAttributes) Set(val *ReviewSubmissionItemUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSubmissionItemUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSubmissionItemUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSubmissionItemUpdateRequestDataAttributes(val *ReviewSubmissionItemUpdateRequestDataAttributes) *NullableReviewSubmissionItemUpdateRequestDataAttributes {
	return &NullableReviewSubmissionItemUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableReviewSubmissionItemUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSubmissionItemUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

