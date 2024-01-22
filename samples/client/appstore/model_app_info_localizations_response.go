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

// checks if the AppInfoLocalizationsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AppInfoLocalizationsResponse{}

// AppInfoLocalizationsResponse struct for AppInfoLocalizationsResponse
type AppInfoLocalizationsResponse struct {
	Data []AppInfoLocalization `json:"data"`
	Included []AppInfo `json:"included,omitempty"`
	Links PagedDocumentLinks `json:"links"`
	Meta *PagingInformation `json:"meta,omitempty"`
}

// NewAppInfoLocalizationsResponse instantiates a new AppInfoLocalizationsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAppInfoLocalizationsResponse(data []AppInfoLocalization, links PagedDocumentLinks) *AppInfoLocalizationsResponse {
	this := AppInfoLocalizationsResponse{}
	this.Data = data
	this.Links = links
	return &this
}

// NewAppInfoLocalizationsResponseWithDefaults instantiates a new AppInfoLocalizationsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAppInfoLocalizationsResponseWithDefaults() *AppInfoLocalizationsResponse {
	this := AppInfoLocalizationsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *AppInfoLocalizationsResponse) GetData() []AppInfoLocalization {
	if o == nil {
		var ret []AppInfoLocalization
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationsResponse) GetDataOk() ([]AppInfoLocalization, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *AppInfoLocalizationsResponse) SetData(v []AppInfoLocalization) {
	o.Data = v
}

// GetIncluded returns the Included field value if set, zero value otherwise.
func (o *AppInfoLocalizationsResponse) GetIncluded() []AppInfo {
	if o == nil || IsNil(o.Included) {
		var ret []AppInfo
		return ret
	}
	return o.Included
}

// GetIncludedOk returns a tuple with the Included field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationsResponse) GetIncludedOk() ([]AppInfo, bool) {
	if o == nil || IsNil(o.Included) {
		return nil, false
	}
	return o.Included, true
}

// HasIncluded returns a boolean if a field has been set.
func (o *AppInfoLocalizationsResponse) HasIncluded() bool {
	if o != nil && !IsNil(o.Included) {
		return true
	}

	return false
}

// SetIncluded gets a reference to the given []AppInfo and assigns it to the Included field.
func (o *AppInfoLocalizationsResponse) SetIncluded(v []AppInfo) {
	o.Included = v
}

// GetLinks returns the Links field value
func (o *AppInfoLocalizationsResponse) GetLinks() PagedDocumentLinks {
	if o == nil {
		var ret PagedDocumentLinks
		return ret
	}

	return o.Links
}

// GetLinksOk returns a tuple with the Links field value
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationsResponse) GetLinksOk() (*PagedDocumentLinks, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Links, true
}

// SetLinks sets field value
func (o *AppInfoLocalizationsResponse) SetLinks(v PagedDocumentLinks) {
	o.Links = v
}

// GetMeta returns the Meta field value if set, zero value otherwise.
func (o *AppInfoLocalizationsResponse) GetMeta() PagingInformation {
	if o == nil || IsNil(o.Meta) {
		var ret PagingInformation
		return ret
	}
	return *o.Meta
}

// GetMetaOk returns a tuple with the Meta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AppInfoLocalizationsResponse) GetMetaOk() (*PagingInformation, bool) {
	if o == nil || IsNil(o.Meta) {
		return nil, false
	}
	return o.Meta, true
}

// HasMeta returns a boolean if a field has been set.
func (o *AppInfoLocalizationsResponse) HasMeta() bool {
	if o != nil && !IsNil(o.Meta) {
		return true
	}

	return false
}

// SetMeta gets a reference to the given PagingInformation and assigns it to the Meta field.
func (o *AppInfoLocalizationsResponse) SetMeta(v PagingInformation) {
	o.Meta = &v
}

func (o AppInfoLocalizationsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AppInfoLocalizationsResponse) ToMap() (map[string]interface{}, error) {
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

type NullableAppInfoLocalizationsResponse struct {
	value *AppInfoLocalizationsResponse
	isSet bool
}

func (v NullableAppInfoLocalizationsResponse) Get() *AppInfoLocalizationsResponse {
	return v.value
}

func (v *NullableAppInfoLocalizationsResponse) Set(val *AppInfoLocalizationsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAppInfoLocalizationsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAppInfoLocalizationsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAppInfoLocalizationsResponse(val *AppInfoLocalizationsResponse) *NullableAppInfoLocalizationsResponse {
	return &NullableAppInfoLocalizationsResponse{value: val, isSet: true}
}

func (v NullableAppInfoLocalizationsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAppInfoLocalizationsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


