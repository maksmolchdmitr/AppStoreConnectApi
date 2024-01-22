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

// checks if the BetaAppClipInvocationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaAppClipInvocationsResponse{}

// BetaAppClipInvocationsResponse struct for BetaAppClipInvocationsResponse
type BetaAppClipInvocationsResponse struct {
	Data []BetaAppClipInvocation `json:"data"`
	Included []BetaAppClipInvocationLocalization `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewBetaAppClipInvocationsResponse instantiates a new BetaAppClipInvocationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaAppClipInvocationsResponse(data []BetaAppClipInvocation, links PagedDocumentLinks) *BetaAppClipInvocationsResponse {
	this := BetaAppClipInvocationsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaAppClipInvocationsResponseWithDefaults instantiates a new BetaAppClipInvocationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaAppClipInvocationsResponseWithDefaults() *BetaAppClipInvocationsResponse {
	this := BetaAppClipInvocationsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaAppClipInvocationsResponse) GetData() []BetaAppClipInvocation {
	if o == nil {
		var ret []BetaAppClipInvocation
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationsResponse) GetDataOk() ([]BetaAppClipInvocation, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BetaAppClipInvocationsResponse) SetData(v []BetaAppClipInvocation) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BetaAppClipInvocationsResponse) GetIncluded() []BetaAppClipInvocationLocalization {
	if o == nil || IsNil(o.Included) {
		var ret []BetaAppClipInvocationLocalization
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationsResponse) GetIncludedOk() ([]BetaAppClipInvocationLocalization, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BetaAppClipInvocationsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []BetaAppClipInvocationLocalization and assigns it to the Included field.
func (o *BetaAppClipInvocationsResponse) SetIncluded(v []BetaAppClipInvocationLocalization) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *BetaAppClipInvocationsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaAppClipInvocationsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BetaAppClipInvocationsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaAppClipInvocationsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BetaAppClipInvocationsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BetaAppClipInvocationsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o BetaAppClipInvocationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaAppClipInvocationsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	if !IsNil(o.Included) {
		toSerialize["included"] = o.Included
	}
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableBetaAppClipInvocationsResponse struct {
	value *BetaAppClipInvocationsResponse
	isSet bool
}

func (v NullableBetaAppClipInvocationsResponse) Get() *BetaAppClipInvocationsResponse {
	return v.value
}

func (v *NullableBetaAppClipInvocationsResponse) Set(val *BetaAppClipInvocationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaAppClipInvocationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaAppClipInvocationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaAppClipInvocationsResponse(val *BetaAppClipInvocationsResponse) *NullableBetaAppClipInvocationsResponse {
	return &NullableBetaAppClipInvocationsResponse{value: val, isSet: true}
}

func (v NullableBetaAppClipInvocationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaAppClipInvocationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


