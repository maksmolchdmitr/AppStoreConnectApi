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

// checks if the BundleIdsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BundleIdsResponse{}

// BundleIdsResponse struct for BundleIdsResponse
type BundleIdsResponse struct {
	Data     []BundleId                       `json:"data"`
	Included []BundleIdsResponseIncludedInner `json:"included,omitempty"`
	Links    PagedDocumentLinks               `json:"links"`
	Meta     *PagingInformation               `json:"meta,omitempty"`
}

// NewBundleIdsResponse instantiates a new BundleIdsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBundleIdsResponse(data []BundleId, links PagedDocumentLinks) *BundleIdsResponse {
	this := BundleIdsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewBundleIdsResponseWithDefaults instantiates a new BundleIdsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBundleIdsResponseWithDefaults() *BundleIdsResponse {
	this := BundleIdsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BundleIdsResponse) GetData() []BundleId {
	if o == nil {
		var ret []BundleId
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BundleIdsResponse) GetDataOk() ([]BundleId, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BundleIdsResponse) SetData(v []BundleId) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *BundleIdsResponse) GetIncluded() []BundleIdsResponseIncludedInner {
	if o == nil || IsNil(o.Included) {
		var ret []BundleIdsResponseIncludedInner
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdsResponse) GetIncludedOk() ([]BundleIdsResponseIncludedInner, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *BundleIdsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []BundleIdsResponseIncludedInner and assigns it to the Included field.
func (o *BundleIdsResponse) SetIncluded(v []BundleIdsResponseIncludedInner) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *BundleIdsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *BundleIdsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *BundleIdsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *BundleIdsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BundleIdsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *BundleIdsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *BundleIdsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o BundleIdsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BundleIdsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableBundleIdsResponse struct {
	value *BundleIdsResponse
	isSet bool
}

func (v NullableBundleIdsResponse) Get() *BundleIdsResponse {
	return v.value
}

func (v *NullableBundleIdsResponse) Set(val *BundleIdsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBundleIdsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBundleIdsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBundleIdsResponse(val *BundleIdsResponse) *NullableBundleIdsResponse {
	return &NullableBundleIdsResponse{value: val, isSet: true}
}

func (v NullableBundleIdsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBundleIdsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
