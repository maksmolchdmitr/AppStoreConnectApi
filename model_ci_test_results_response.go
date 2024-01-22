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

// checks if the CiTestResultsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CiTestResultsResponse{}

// CiTestResultsResponse struct for CiTestResultsResponse
type CiTestResultsResponse struct {
	Data  []CiTestResult     `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// NewCiTestResultsResponse instantiates a new CiTestResultsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCiTestResultsResponse(data []CiTestResult, links PagedDocumentLinks) *CiTestResultsResponse {
	this := CiTestResultsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewCiTestResultsResponseWithDefaults instantiates a new CiTestResultsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCiTestResultsResponseWithDefaults() *CiTestResultsResponse {
	this := CiTestResultsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *CiTestResultsResponse) GetData() []CiTestResult {
	if o == nil {
		var ret []CiTestResult
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *CiTestResultsResponse) GetDataOk() ([]CiTestResult, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *CiTestResultsResponse) SetData(v []CiTestResult) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *CiTestResultsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *CiTestResultsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *CiTestResultsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *CiTestResultsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CiTestResultsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *CiTestResultsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *CiTestResultsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o CiTestResultsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CiTestResultsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableCiTestResultsResponse struct {
	value *CiTestResultsResponse
	isSet bool
}

func (v NullableCiTestResultsResponse) Get() *CiTestResultsResponse {
	return v.value
}

func (v *NullableCiTestResultsResponse) Set(val *CiTestResultsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCiTestResultsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCiTestResultsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCiTestResultsResponse(val *CiTestResultsResponse) *NullableCiTestResultsResponse {
	return &NullableCiTestResultsResponse{value: val, isSet: true}
}

func (v NullableCiTestResultsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCiTestResultsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
