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

// checks if the AppCategoriesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppCategoriesResponse{}

// AppCategoriesResponse struct for AppCategoriesResponse
type AppCategoriesResponse struct {
	Data     []AppCategory      `json:"data"`
	Included []AppCategory      `json:"included,omitempty"`
	Links    PagedDocumentLinks `json:"links"`
	Meta     *PagingInformation `json:"meta,omitempty"`
}

// NewAppCategoriesResponse instantiates a new AppCategoriesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppCategoriesResponse(data []AppCategory, links PagedDocumentLinks) *AppCategoriesResponse {
	this := AppCategoriesResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppCategoriesResponseWithDefaults instantiates a new AppCategoriesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppCategoriesResponseWithDefaults() *AppCategoriesResponse {
	this := AppCategoriesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppCategoriesResponse) GetData() []AppCategory {
	if o == nil {
		var ret []AppCategory
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppCategoriesResponse) GetDataOk() ([]AppCategory, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppCategoriesResponse) SetData(v []AppCategory) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppCategoriesResponse) GetIncluded() []AppCategory {
	if o == nil || IsNil(o.Included) {
		var ret []AppCategory
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCategoriesResponse) GetIncludedOk() ([]AppCategory, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppCategoriesResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppCategory and assigns it to the Included field.
func (o *AppCategoriesResponse) SetIncluded(v []AppCategory) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppCategoriesResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppCategoriesResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppCategoriesResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppCategoriesResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppCategoriesResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppCategoriesResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppCategoriesResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppCategoriesResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppCategoriesResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAppCategoriesResponse struct {
	value *AppCategoriesResponse
	isSet bool
}

func (v NullableAppCategoriesResponse) Get() *AppCategoriesResponse {
	return v.value
}

func (v *NullableAppCategoriesResponse) Set(val *AppCategoriesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppCategoriesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppCategoriesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppCategoriesResponse(val *AppCategoriesResponse) *NullableAppCategoriesResponse {
	return &NullableAppCategoriesResponse{value: val, isSet: true}
}

func (v NullableAppCategoriesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppCategoriesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
