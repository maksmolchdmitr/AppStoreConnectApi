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

// checks if the AppStoreVersionLocalizationsWithoutIncludesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppStoreVersionLocalizationsWithoutIncludesResponse{}

// AppStoreVersionLocalizationsWithoutIncludesResponse struct for AppStoreVersionLocalizationsWithoutIncludesResponse
type AppStoreVersionLocalizationsWithoutIncludesResponse struct {
	Data  []AppStoreVersion  `json:"data"`
	Links PagedDocumentLinks `json:"links"`
	Meta  *PagingInformation `json:"meta,omitempty"`
}

// NewAppStoreVersionLocalizationsWithoutIncludesResponse instantiates a new AppStoreVersionLocalizationsWithoutIncludesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppStoreVersionLocalizationsWithoutIncludesResponse(data []AppStoreVersion, links PagedDocumentLinks) *AppStoreVersionLocalizationsWithoutIncludesResponse {
	this := AppStoreVersionLocalizationsWithoutIncludesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppStoreVersionLocalizationsWithoutIncludesResponseWithDefaults instantiates a new AppStoreVersionLocalizationsWithoutIncludesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppStoreVersionLocalizationsWithoutIncludesResponseWithDefaults() *AppStoreVersionLocalizationsWithoutIncludesResponse {
	this := AppStoreVersionLocalizationsWithoutIncludesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) GetData() []AppStoreVersion {
	if o == nil {
		var ret []AppStoreVersion
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) GetDataOk() ([]AppStoreVersion, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) SetData(v []AppStoreVersion) {
	o.Data = v
}

// GetLinks returns the Links field value
func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppStoreVersionLocalizationsWithoutIncludesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppStoreVersionLocalizationsWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppStoreVersionLocalizationsWithoutIncludesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["links"] = o.Links
	if !IsNil(o.Meta) {
		toSerialize["meta"] = o.Meta
	}
	return toSerialize, nil
}

type NullableAppStoreVersionLocalizationsWithoutIncludesResponse struct {
	value *AppStoreVersionLocalizationsWithoutIncludesResponse
	isSet bool
}

func (v NullableAppStoreVersionLocalizationsWithoutIncludesResponse) Get() *AppStoreVersionLocalizationsWithoutIncludesResponse {
	return v.value
}

func (v *NullableAppStoreVersionLocalizationsWithoutIncludesResponse) Set(val *AppStoreVersionLocalizationsWithoutIncludesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppStoreVersionLocalizationsWithoutIncludesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppStoreVersionLocalizationsWithoutIncludesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppStoreVersionLocalizationsWithoutIncludesResponse(val *AppStoreVersionLocalizationsWithoutIncludesResponse) *NullableAppStoreVersionLocalizationsWithoutIncludesResponse {
	return &NullableAppStoreVersionLocalizationsWithoutIncludesResponse{value: val, isSet: true}
}

func (v NullableAppStoreVersionLocalizationsWithoutIncludesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppStoreVersionLocalizationsWithoutIncludesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
