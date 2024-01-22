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

// checks if the CustomerReviewResponseV1Relationships type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CustomerReviewResponseV1Relationships{}

// CustomerReviewResponseV1Relationships struct for CustomerReviewResponseV1Relationships
type CustomerReviewResponseV1Relationships struct {
	Review *CustomerReviewResponseV1RelationshipsReview `json:"review,omitempty"`
}

// NewCustomerReviewResponseV1Relationships instantiates a new CustomerReviewResponseV1Relationships object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerReviewResponseV1Relationships() *CustomerReviewResponseV1Relationships {
	this := CustomerReviewResponseV1Relationships{}
	return &this
}

// NewCustomerReviewResponseV1RelationshipsWithDefaults instantiates a new CustomerReviewResponseV1Relationships object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerReviewResponseV1RelationshipsWithDefaults() *CustomerReviewResponseV1Relationships {
	this := CustomerReviewResponseV1Relationships{}
	return &this
}

// GetReview returns the Review field value if set, zero value otherwise.
func (o *CustomerReviewResponseV1Relationships) GetReview() CustomerReviewResponseV1RelationshipsReview {
	if o == nil || IsNil(o.Review) {
		var ret CustomerReviewResponseV1RelationshipsReview
		return ret
	}
	return *o.Review
}

// GetReviewOk returns a tuple with the Review field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CustomerReviewResponseV1Relationships) GetReviewOk() (*CustomerReviewResponseV1RelationshipsReview, bool) {
	if o == nil || IsNil(o.Review) {
		return nil, false
	}
	return o.Review, true
}

// HasReview returns a boolean if a field has been set.
func (o *CustomerReviewResponseV1Relationships) HasReview() bool {
	if o != nil && !IsNil(o.Review) {
		return true
	}

	return false
}

// SetReview gets a reference to the given CustomerReviewResponseV1RelationshipsReview and assigns it to the Review field.
func (o *CustomerReviewResponseV1Relationships) SetReview(v CustomerReviewResponseV1RelationshipsReview) {
	o.Review = &v
}

func (o CustomerReviewResponseV1Relationships) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CustomerReviewResponseV1Relationships) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Review) {
		toSerialize["review"] = o.Review
	}
	return toSerialize, nil
}

type NullableCustomerReviewResponseV1Relationships struct {
	value *CustomerReviewResponseV1Relationships
	isSet bool
}

func (v NullableCustomerReviewResponseV1Relationships) Get() *CustomerReviewResponseV1Relationships {
	return v.value
}

func (v *NullableCustomerReviewResponseV1Relationships) Set(val *CustomerReviewResponseV1Relationships) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerReviewResponseV1Relationships) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerReviewResponseV1Relationships) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerReviewResponseV1Relationships(val *CustomerReviewResponseV1Relationships) *NullableCustomerReviewResponseV1Relationships {
	return &NullableCustomerReviewResponseV1Relationships{value: val, isSet: true}
}

func (v NullableCustomerReviewResponseV1Relationships) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerReviewResponseV1Relationships) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
