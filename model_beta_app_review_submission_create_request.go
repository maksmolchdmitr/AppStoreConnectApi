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

// checks if the BetaAppReviewSubmissionCreateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppReviewSubmissionCreateRequest{}

// BetaAppReviewSubmissionCreateRequest struct for BetaAppReviewSubmissionCreateRequest
type BetaAppReviewSubmissionCreateRequest struct {
	Data BetaAppReviewSubmissionCreateRequestData `json:"data"`
}

// NewBetaAppReviewSubmissionCreateRequest instantiates a new BetaAppReviewSubmissionCreateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppReviewSubmissionCreateRequest(data BetaAppReviewSubmissionCreateRequestData) *BetaAppReviewSubmissionCreateRequest {
	this := BetaAppReviewSubmissionCreateRequest{}
	this.Data = data
	return &this
}

// NewBetaAppReviewSubmissionCreateRequestWithDefaults instantiates a new BetaAppReviewSubmissionCreateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppReviewSubmissionCreateRequestWithDefaults() *BetaAppReviewSubmissionCreateRequest {
	this := BetaAppReviewSubmissionCreateRequest{}
	return &this
}

// GetData returns the Data field value
func (o *BetaAppReviewSubmissionCreateRequest) GetData() BetaAppReviewSubmissionCreateRequestData {
	if o == nil {
		var ret BetaAppReviewSubmissionCreateRequestData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaAppReviewSubmissionCreateRequest) GetDataOk() (*BetaAppReviewSubmissionCreateRequestData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BetaAppReviewSubmissionCreateRequest) SetData(v BetaAppReviewSubmissionCreateRequestData) {
	o.Data = v
}

func (o BetaAppReviewSubmissionCreateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppReviewSubmissionCreateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

type NullableBetaAppReviewSubmissionCreateRequest struct {
	value *BetaAppReviewSubmissionCreateRequest
	isSet bool
}

func (v NullableBetaAppReviewSubmissionCreateRequest) Get() *BetaAppReviewSubmissionCreateRequest {
	return v.value
}

func (v *NullableBetaAppReviewSubmissionCreateRequest) Set(val *BetaAppReviewSubmissionCreateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppReviewSubmissionCreateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppReviewSubmissionCreateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppReviewSubmissionCreateRequest(val *BetaAppReviewSubmissionCreateRequest) *NullableBetaAppReviewSubmissionCreateRequest {
	return &NullableBetaAppReviewSubmissionCreateRequest{value: val, isSet: true}
}

func (v NullableBetaAppReviewSubmissionCreateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppReviewSubmissionCreateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
