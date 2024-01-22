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

// checks if the ReviewSubmissionUpdateRequestDataAttributes type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ReviewSubmissionUpdateRequestDataAttributes{}

// ReviewSubmissionUpdateRequestDataAttributes struct for ReviewSubmissionUpdateRequestDataAttributes
type ReviewSubmissionUpdateRequestDataAttributes struct {
	Submitted *bool `json:"submitted,omitempty"`
	Canceled  *bool `json:"canceled,omitempty"`
}

// NewReviewSubmissionUpdateRequestDataAttributes instantiates a new ReviewSubmissionUpdateRequestDataAttributes object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewSubmissionUpdateRequestDataAttributes() *ReviewSubmissionUpdateRequestDataAttributes {
	this := ReviewSubmissionUpdateRequestDataAttributes{}
	return &this
}

// NewReviewSubmissionUpdateRequestDataAttributesWithDefaults instantiates a new ReviewSubmissionUpdateRequestDataAttributes object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewSubmissionUpdateRequestDataAttributesWithDefaults() *ReviewSubmissionUpdateRequestDataAttributes {
	this := ReviewSubmissionUpdateRequestDataAttributes{}
	return &this
}

// GetSubmitted returns the Submitted field value if set, zero value otherwise.
func (o *ReviewSubmissionUpdateRequestDataAttributes) GetSubmitted() bool {
	if o == nil || IsNil(o.Submitted) {
		var ret bool
		return ret
	}
	return *o.Submitted
}

// GetSubmittedOk returns a tuple with the Submitted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionUpdateRequestDataAttributes) GetSubmittedOk() (*bool, bool) {
	if o == nil || IsNil(o.Submitted) {
		return nil, false
	}
	return o.Submitted, true
}

// HasSubmitted returns a boolean if a field has been set.
func (o *ReviewSubmissionUpdateRequestDataAttributes) HasSubmitted() bool {
	if o != nil && !IsNil(o.Submitted) {
		return true
	}

	return false
}

// SetSubmitted gets a reference to the given bool and assigns it to the Submitted field.
func (o *ReviewSubmissionUpdateRequestDataAttributes) SetSubmitted(v bool) {
	o.Submitted = &v
}

// GetCanceled returns the Canceled field value if set, zero value otherwise.
func (o *ReviewSubmissionUpdateRequestDataAttributes) GetCanceled() bool {
	if o == nil || IsNil(o.Canceled) {
		var ret bool
		return ret
	}
	return *o.Canceled
}

// GetCanceledOk returns a tuple with the Canceled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewSubmissionUpdateRequestDataAttributes) GetCanceledOk() (*bool, bool) {
	if o == nil || IsNil(o.Canceled) {
		return nil, false
	}
	return o.Canceled, true
}

// HasCanceled returns a boolean if a field has been set.
func (o *ReviewSubmissionUpdateRequestDataAttributes) HasCanceled() bool {
	if o != nil && !IsNil(o.Canceled) {
		return true
	}

	return false
}

// SetCanceled gets a reference to the given bool and assigns it to the Canceled field.
func (o *ReviewSubmissionUpdateRequestDataAttributes) SetCanceled(v bool) {
	o.Canceled = &v
}

func (o ReviewSubmissionUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ReviewSubmissionUpdateRequestDataAttributes) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Submitted) {
		toSerialize["submitted"] = o.Submitted
	}
	if !IsNil(o.Canceled) {
		toSerialize["canceled"] = o.Canceled
	}
	return toSerialize, nil
}

type NullableReviewSubmissionUpdateRequestDataAttributes struct {
	value *ReviewSubmissionUpdateRequestDataAttributes
	isSet bool
}

func (v NullableReviewSubmissionUpdateRequestDataAttributes) Get() *ReviewSubmissionUpdateRequestDataAttributes {
	return v.value
}

func (v *NullableReviewSubmissionUpdateRequestDataAttributes) Set(val *ReviewSubmissionUpdateRequestDataAttributes) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewSubmissionUpdateRequestDataAttributes) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewSubmissionUpdateRequestDataAttributes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewSubmissionUpdateRequestDataAttributes(val *ReviewSubmissionUpdateRequestDataAttributes) *NullableReviewSubmissionUpdateRequestDataAttributes {
	return &NullableReviewSubmissionUpdateRequestDataAttributes{value: val, isSet: true}
}

func (v NullableReviewSubmissionUpdateRequestDataAttributes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewSubmissionUpdateRequestDataAttributes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
