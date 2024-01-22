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

// checks if the UserVisibleAppsLinkagesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserVisibleAppsLinkagesResponse{}

// UserVisibleAppsLinkagesResponse struct for UserVisibleAppsLinkagesResponse
type UserVisibleAppsLinkagesResponse struct {
	Data  []AppAvailabilityV2CreateRequestDataRelationshipsAppData `json:"data"`
	Links PagedDocumentLinks                                       `json:"links"`
	Meta  *PagingInformation                                       `json:"meta,omitempty"`
}

// NewUserVisibleAppsLinkagesResponse instantiates a new UserVisibleAppsLinkagesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserVisibleAppsLinkagesResponse(data []AppAvailabilityV2CreateRequestDataRelationshipsAppData, links PagedDocumentLinks) *UserVisibleAppsLinkagesResponse {
	this := UserVisibleAppsLinkagesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewUserVisibleAppsLinkagesResponseWithDefaults instantiates a new UserVisibleAppsLinkagesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserVisibleAppsLinkagesResponseWithDefaults() *UserVisibleAppsLinkagesResponse {
	this := UserVisibleAppsLinkagesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *UserVisibleAppsLinkagesResponse) GetData() []AppAvailabilityV2CreateRequestDataRelationshipsAppData {
	if o == nil {
		var ret []AppAvailabilityV2CreateRequestDataRelationshipsAppData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *UserVisibleAppsLinkagesResponse) GetDataOk() ([]AppAvailabilityV2CreateRequestDataRelationshipsAppData, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *UserVisibleAppsLinkagesResponse) SetData(v []AppAvailabilityV2CreateRequestDataRelationshipsAppData) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *UserVisibleAppsLinkagesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *UserVisibleAppsLinkagesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *UserVisibleAppsLinkagesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *UserVisibleAppsLinkagesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserVisibleAppsLinkagesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *UserVisibleAppsLinkagesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *UserVisibleAppsLinkagesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o UserVisibleAppsLinkagesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserVisibleAppsLinkagesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableUserVisibleAppsLinkagesResponse struct {
	value *UserVisibleAppsLinkagesResponse
	isSet bool
}

func (v NullableUserVisibleAppsLinkagesResponse) Get() *UserVisibleAppsLinkagesResponse {
	return v.value
}

func (v *NullableUserVisibleAppsLinkagesResponse) Set(val *UserVisibleAppsLinkagesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableUserVisibleAppsLinkagesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableUserVisibleAppsLinkagesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserVisibleAppsLinkagesResponse(val *UserVisibleAppsLinkagesResponse) *NullableUserVisibleAppsLinkagesResponse {
	return &NullableUserVisibleAppsLinkagesResponse{value: val, isSet: true}
}

func (v NullableUserVisibleAppsLinkagesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserVisibleAppsLinkagesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
