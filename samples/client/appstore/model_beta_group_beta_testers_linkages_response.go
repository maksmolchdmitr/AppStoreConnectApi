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

// checks if the BetaGroupBetaTestersLinkagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BetaGroupBetaTestersLinkagesResponse{}

// BetaGroupBetaTestersLinkagesResponse struct for BetaGroupBetaTestersLinkagesResponse
type BetaGroupBetaTestersLinkagesResponse struct {
	Data []BetaGroupRelationshipsBetaTestersDataInner `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewBetaGroupBetaTestersLinkagesResponse instantiates a new BetaGroupBetaTestersLinkagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBetaGroupBetaTestersLinkagesResponse(data []BetaGroupRelationshipsBetaTestersDataInner, links PagedDocumentLinks) *BetaGroupBetaTestersLinkagesResponse {
	this := BetaGroupBetaTestersLinkagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBetaGroupBetaTestersLinkagesResponseWithDefaults instantiates a new BetaGroupBetaTestersLinkagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBetaGroupBetaTestersLinkagesResponseWithDefaults() *BetaGroupBetaTestersLinkagesResponse {
	this := BetaGroupBetaTestersLinkagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BetaGroupBetaTestersLinkagesResponse) GetData() []BetaGroupRelationshipsBetaTestersDataInner {
	if o == nil {
		var ret []BetaGroupRelationshipsBetaTestersDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BetaGroupBetaTestersLinkagesResponse) GetDataOk() ([]BetaGroupRelationshipsBetaTestersDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BetaGroupBetaTestersLinkagesResponse) SetData(v []BetaGroupRelationshipsBetaTestersDataInner) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *BetaGroupBetaTestersLinkagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BetaGroupBetaTestersLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BetaGroupBetaTestersLinkagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BetaGroupBetaTestersLinkagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BetaGroupBetaTestersLinkagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BetaGroupBetaTestersLinkagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BetaGroupBetaTestersLinkagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o BetaGroupBetaTestersLinkagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BetaGroupBetaTestersLinkagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableBetaGroupBetaTestersLinkagesResponse struct {
	value *BetaGroupBetaTestersLinkagesResponse
	isSet bool
}

func (v NullableBetaGroupBetaTestersLinkagesResponse) Get() *BetaGroupBetaTestersLinkagesResponse {
	return v.value
}

func (v *NullableBetaGroupBetaTestersLinkagesResponse) Set(val *BetaGroupBetaTestersLinkagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBetaGroupBetaTestersLinkagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBetaGroupBetaTestersLinkagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBetaGroupBetaTestersLinkagesResponse(val *BetaGroupBetaTestersLinkagesResponse) *NullableBetaGroupBetaTestersLinkagesResponse {
	return &NullableBetaGroupBetaTestersLinkagesResponse{value: val, isSet: true}
}

func (v NullableBetaGroupBetaTestersLinkagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBetaGroupBetaTestersLinkagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


