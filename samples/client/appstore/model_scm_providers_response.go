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

// checks if the ScmProvidersResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScmProvidersResponse{}

// ScmProvidersResponse struct for ScmProvidersResponse
type ScmProvidersResponse struct {
	Data []ScmProvider `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewScmProvidersResponse instantiates a new ScmProvidersResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScmProvidersResponse(data []ScmProvider, links PagedDocumentLinks) *ScmProvidersResponse {
	this := ScmProvidersResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewScmProvidersResponseWithDefaults instantiates a new ScmProvidersResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScmProvidersResponseWithDefaults() *ScmProvidersResponse {
	this := ScmProvidersResponse{}
	return &this
}

// GetData returns the Data field value
func (o *ScmProvidersResponse) GetData() []ScmProvider {
	if o == nil {
		var ret []ScmProvider
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ScmProvidersResponse) GetDataOk() ([]ScmProvider, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *ScmProvidersResponse) SetData(v []ScmProvider) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *ScmProvidersResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *ScmProvidersResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *ScmProvidersResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *ScmProvidersResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScmProvidersResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *ScmProvidersResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *ScmProvidersResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o ScmProvidersResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScmProvidersResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableScmProvidersResponse struct {
	value *ScmProvidersResponse
	isSet bool
}

func (v NullableScmProvidersResponse) Get() *ScmProvidersResponse {
	return v.value
}

func (v *NullableScmProvidersResponse) Set(val *ScmProvidersResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableScmProvidersResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableScmProvidersResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScmProvidersResponse(val *ScmProvidersResponse) *NullableScmProvidersResponse {
	return &NullableScmProvidersResponse{value: val, isSet: true}
}

func (v NullableScmProvidersResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScmProvidersResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

